package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/eavehh/marketpl.microserv/internal/checkout/api"
	"github.com/eavehh/marketpl.microserv/internal/checkout/api/handlers"
	"github.com/eavehh/marketpl.microserv/internal/checkout/application/queries"
	"github.com/eavehh/marketpl.microserv/internal/checkout/infrastructure/persistence"
	"github.com/eavehh/marketpl.microserv/internal/shared"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println(".env file does not exist: ", err)
	}

	app_post := os.Getenv("CHECKOUT_APP_PORT")
	migrations_path := os.Getenv("CHECKOUT_MIGRATIONS_PATH")
	dsn := os.Getenv("CHECKOUT_DATABASE_URL")

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err, "CHECKOUT: sql.Open(postgres, dsn)")
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Fatal(err, "CHECKOUT: Cannot connect to the DB; if err = db.Ping()")
	}

	log.Println("CHECKOUT: Successfully connect to the DB ")

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal("CHECKOUT: Automigrate: postgres.WithInstance error: ", err)
	}
	m, err := migrate.NewWithDatabaseInstance(migrations_path, "postgres", driver)
	if err != nil {
		log.Fatal("CHECKOUT: Automigrate: migrate.NewWithDatabaseInstance error: ", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("CHECKOUT: Automigrate: migrate.up error: ", err)
	}
	log.Println("migrate applied")

	repo := persistence.New_order_repository(db)
	order_by_id_handler := queries.New_order_by_id_query_handler(repo)

	order_handler := handlers.New_order_handler(order_by_id_handler)

	r := gin.Default()
	r.Use(shared.Err_handler_middleware())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	api.Register_routes(r, order_handler)

	if err := r.Run(":" + app_post); err != nil {
		log.Fatal(err)
	}
}
