package main

import (
	"fmt"
	"log"
	"os"

	"database/sql"
	_ "database/sql"

	"github.com/eavehh/marketpl.microserv/internal/catalog/api"
	"github.com/eavehh/marketpl.microserv/internal/catalog/api/handlers"
	"github.com/eavehh/marketpl.microserv/internal/catalog/application/queries"
	"github.com/eavehh/marketpl.microserv/internal/catalog/infrastructure/persistence"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	if err := godotenv.Load("../../.env"); err != nil {
		log.Println(".env file doesnt exist ", err)
	}

	pg_host := os.Getenv("localhost")
	app_port := os.Getenv("9101")
	pg_user := os.Getenv("postgres")
	pg_pass := os.Getenv("12345678")
	pg_db := os.Getenv("catalog_db_dev")
	pg_ssl := os.Getenv("disable")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		pg_host, app_port, pg_user, pg_pass, pg_db, pg_ssl,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err, "sql.Open(postgres, dsn)")
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Fatal(err, "if err = db.Ping()")
	}

	engine := gin.Default()
	engine.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	brands_repo := persistence.New_brand_repo(db)
	brands_query_handler := queries.New_brands_queries(brands_repo)
	brands_api_handler := handlers.New_brands_handler(brands_query_handler)

	api.Register_routes(engine, brands_api_handler)

	if err := engine.Run(":9001"); err != nil {
		log.Fatal(err.Error())
	}
}
