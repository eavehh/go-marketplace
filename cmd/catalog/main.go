package main

import (
	"fmt"
	"log"
	"os"

	"database/sql"
	_ "database/sql"

	"github.com/eavehh/marketpl.microserv/internal/catalog/api"
	"github.com/eavehh/marketpl.microserv/internal/catalog/api/handlers"
	commands "github.com/eavehh/marketpl.microserv/internal/catalog/application/comands"
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
	categories_repo := persistence.New_Category_repo(db)
	items_repo := persistence.New_item_repository(db)

	list_brands := queries.New_brands_handler(brands_repo)
	list_categories := queries.New_categories_handler(categories_repo)
	list_items := queries.New_catalog_items_handler(items_repo)
	item_by_id := queries.New_catalog_item_by_id_handler(items_repo)
	item_by_title := queries.New_catalog_item_by_title_handler(items_repo)
	create_item := commands.New_create_catalog_item_handler(items_repo)

	brands_handler := handlers.New_brands_handler(list_brands)
	categories_handler := handlers.New_categories_handler(list_categories)
	items_handler := handlers.New_catalog_items_handler(
		list_items,
		item_by_id,
		item_by_title,
		create_item,
	)

	api.Register_routes(engine,
		brands_handler,
		categories_handler,
		items_handler,
	)

	if err := engine.Run(":9001"); err != nil {
		log.Fatal(err.Error())
	}
}
