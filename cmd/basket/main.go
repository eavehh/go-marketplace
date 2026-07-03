package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println(".env file does not exist: ", err)
	}

	app_post := os.Getenv("BASKET_APP_PORT")
	migrations_path := os.Getenv("BASKET_MIGRATIONS_PATH")

	dsn := os.Getenv("BASKET_DATABASE_URL")

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

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	if err := r.Run(":" + app_post); err != nil {
		log.Fatal(err)
	}
}
