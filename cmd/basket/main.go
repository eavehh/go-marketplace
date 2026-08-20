package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/eavehh/marketpl.microserv/internal/basket/api"
	"github.com/eavehh/marketpl.microserv/internal/basket/api/handlers"
	"github.com/eavehh/marketpl.microserv/internal/basket/application/commands"
	"github.com/eavehh/marketpl.microserv/internal/basket/application/interfaces"
	"github.com/eavehh/marketpl.microserv/internal/basket/application/queries"
	"github.com/eavehh/marketpl.microserv/internal/basket/infrastructure/cache"
	"github.com/eavehh/marketpl.microserv/internal/basket/infrastructure/persistence"
	"github.com/eavehh/marketpl.microserv/internal/promotion/grpc/pb"
	"github.com/eavehh/marketpl.microserv/internal/shared"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println(".env file does not exist: ", err)
	}

	app_post := os.Getenv("BASKET_APP_PORT")
	migrations_path := os.Getenv("BASKET_MIGRATIONS_PATH")
	dsn := os.Getenv("BASKET_DATABASE_URL")
	redis_url := os.Getenv("BASKET_REDIS_URL")
	redis_password := os.Getenv("BASKET_REDIS_PASSWORD")

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err, "BASKET: sql.Open(postgres, dsn)")
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Fatal(err, "BASKET: Cannot connect to the DB; if err = db.Ping()")
	}

	log.Println("BASKET: Successfully connect to the DB ")

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal("BASKET: Automigrate: postgres.WithInstance error: ", err)
	}
	m, err := migrate.NewWithDatabaseInstance(migrations_path, "postgres", driver)
	if err != nil {
		log.Fatal("BASKET: Automigrate: migrate.NewWithDatabaseInstance error: ", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("BASKET: Automigrate: migrate.up error: ", err)
	}
	log.Println("migrate applied")

	redis_client := redis.NewClient(&redis.Options{
		Addr:     redis_url,
		Password: redis_password,
		DB:       0,
	})

	if err := redis_client.Ping(context.Background()).Err(); err != nil {
		log.Fatal("redis ping error")
	}
	log.Println("redis connected")

	defer redis_client.Close()

	pg_repo := persistence.New_cart_repository(db)

	var repo interfaces.Cart_repository = cache.New_redis_cart_repository(
		pg_repo,
		redis_client,
	)

	promotion_host := os.Getenv("BASKET_PROMOTION_HOST")
	promotion_port := os.Getenv("BASKET_PROMOTION_PORT")

	promotion_addr := fmt.Sprintf("%s:%s", promotion_host, promotion_port)

	grpc_conn, err := grpc.NewClient(
		promotion_addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()), // для снятия протокола шифрования (только для dev)
	)

	if err != nil {
		log.Printf("warning: faild to connect to promotion service: %s", promotion_addr)
	} else {
		defer grpc_conn.Close()
		log.Printf("promotin grpc client connected to %s", promotion_addr)
	}

	var promo_client pb.PromotionServiceClient
	if grpc_conn != nil {
		promo_client = pb.NewPromotionServiceClient(grpc_conn)
	}

	save_cart_handler := commands.New_save_cart_handler(repo, promo_client)
	get_cart_handler := queries.New_get_cart_handler(repo)
	delete_cart_handler := commands.New_delete_handler(repo)

	cart_handler := handlers.New_cart_handler(
		save_cart_handler,
		get_cart_handler,
		delete_cart_handler,
	)

	r := gin.Default()
	r.Use(shared.Err_handler_middleware())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	api.Register_routes(r, cart_handler)

	if err := r.Run(":" + app_post); err != nil {
		log.Fatal(err)
	}
}
