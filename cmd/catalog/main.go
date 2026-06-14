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

	if err := godotenv.Load(".env"); err != nil {
		log.Println(".env file does not exis:", err)
	}

	pg_host := os.Getenv("CATALOG_PG_HOST")
	pg_port := os.Getenv("CATALOG_PG_PORT")
	pg_user := os.Getenv("CATALOG_PG_USER")
	pg_pass := os.Getenv("CATALOG_PG_PASS")
	pg_db := os.Getenv("CATALOG_PG_DB")
	pg_ssl := os.Getenv("CATALOG_PG_SSL")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		pg_host, pg_port, pg_user, pg_pass, pg_db, pg_ssl,
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

	categories_repo := persistence.New_Category_repo(db)
	categories_query_handler := queries.New_categories_query(categories_repo)
	categories_handler := handlers.New_categories_handler(categories_query_handler)

	api.Register_routes(engine, brands_api_handler, categories_handler)

	if err := engine.Run(":9001"); err != nil {
		log.Fatal(err.Error())
	}
}
