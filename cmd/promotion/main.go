package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/eavehh/marketpl.microserv/internal/promotion/application/commands"
	"github.com/eavehh/marketpl.microserv/internal/promotion/application/queries"
	promotion_grpc "github.com/eavehh/marketpl.microserv/internal/promotion/grpc"
	pb "github.com/eavehh/marketpl.microserv/internal/promotion/grpc/pb"
	"github.com/eavehh/marketpl.microserv/internal/promotion/infrastructure/persistence"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App_config struct {
	Grpc_port       string
	Database_url    string
	Migrations_path string
}

//  у promotion подход даже более правильный: короткоживущее соединение db_mig открывается, мигрирует,
//  аккуратно закрывается через m.Close() (это чистит миграционный lock в БД — важно, если несколько инстансов
//  сервиса стартуют одновременно), и только потом открывается отдельный, настроенный под нагрузку пул db для самого приложения.
//  Это стандартный паттерн в проде — миграции и рантайм-пул сознательно разделяют, чтобы настройки
//  одного (короткая жизнь, дефолтные лимиты) не мешали другому (долгоживущий пул с MaxOpenConns/MaxIdleConns).

func main() {
	cfg := Load_config(".env")

	db_mig, err := sql.Open("mysql", cfg.Database_url)
	if err != nil {
		log.Fatalf("migration_db open: %v", err)
	}

	if err := run_migrations(db_mig, cfg.Migrations_path); err != nil {
		log.Fatalf("run migrations: %v", err)
	}

	db, err := open_db(cfg.Database_url)
	if err != nil {
		log.Fatalf("open_db error: %v", err)
	}
	defer db.Close()
	log.Print("mysql connected")

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := run_grpc_server(ctx, cfg.Grpc_port, db); err != nil {
		log.Fatalf("run_grpc_server: %v", err)
	}
}

func open_db(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("sql.Open(): %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetConnMaxIdleTime(time.Minute * 2)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("db.Ping: %w", err)
	}
	return db, nil
}

func Load_config(env_path string) App_config {
	if err := godotenv.Load(env_path); err != nil {
		log.Printf("info: .env file not found (%s) using default values", env_path)
	}
	return App_config{
		Grpc_port:       get_env("PROMOTION_GRPC_PORT", "9003"),
		Database_url:    get_env("PROMOTION_DATABASE_URL", "root:123456789@tcp(localhost:9103)/promotion-db?parseTime=true&multiStatements=true"),
		Migrations_path: get_env("PROMOTION_MIGRATIONS_PATH", "file://./migrations/promotion"),
	}
}

func get_env(key string, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func run_migrations(db *sql.DB, migrations_path string) error {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return fmt.Errorf("migrate driver: %w", err)
	}
	m, err := migrate.NewWithDatabaseInstance(migrations_path, "mysql", driver)
	if err != nil {
		return fmt.Errorf("migrate init: %w", err)
	}

	defer m.Close()

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("migrate up: %w", err)
	}
	return nil
}

func run_grpc_server(ctx context.Context, port string, db *sql.DB) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return fmt.Errorf("net.Listen: %w", err)
	}

	grpc_server := grpc.NewServer()
	greet_service := promotion_grpc.New_greeter_service()

	repo := persistence.New_promo_repository(db)
	query_handler := queries.New_get_by_catalog_item_handler(repo)
	create_handler := commands.New_create_promo_handler(repo)
	update_handler := commands.New_update_promo_handler(repo)
	promo_service := promotion_grpc.NewPromotionService(query_handler, create_handler, update_handler)

	pb.RegisterGreeterServer(grpc_server, greet_service)
	pb.RegisterPromotionServiceServer(grpc_server, promo_service)

	reflection.Register(grpc_server)
	// postman (server reflection)

	go func() {
		<-ctx.Done()
		log.Println("shutdown signal received: stopped grpc server...")
		timer := time.AfterFunc(10*time.Second, func() {
			log.Println("timeout exceede, derver stopped")
			grpc_server.Stop()
		})
		defer timer.Stop()
		grpc_server.GracefulStop()
	}()

	log.Printf("grpc server id listening on: %s", port)

	if err := grpc_server.Serve(lis); err != nil {
		return fmt.Errorf("grpc_server.Serve: %w", err)
	}
	return nil
}
