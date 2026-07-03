package main

import (
	"fmt"
	"log"
	"os"

	"database/sql"
	_ "database/sql"

	"github.com/eavehh/marketpl.microserv/internal/catalog/api"
	"github.com/eavehh/marketpl.microserv/internal/catalog/api/handlers"
	"github.com/eavehh/marketpl.microserv/internal/catalog/application/commands"
	"github.com/eavehh/marketpl.microserv/internal/catalog/application/queries"
	"github.com/eavehh/marketpl.microserv/internal/catalog/infrastructure/persistence"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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
	pg_pass := os.Getenv("CATALOG_PG_PASSWORD")
	pg_db := os.Getenv("CATALOG_PG_DATABASE")
	pg_ssl := os.Getenv("CATALOG_PG_SSLMODE")
	migrations_path := os.Getenv("CATALOG_MIGRATIONS_PATH")
	app_port := ":" + os.Getenv("CATALOG_APP_PORT")
	// app_port := ":9001"
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		pg_host, pg_port, pg_user, pg_pass, pg_db, pg_ssl,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err, "CATALOG: sql.Open(postgres, dsn)")
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Fatal(err, "CATALOG: Cannot connect to the DB; if err = db.Ping()")
	}

	log.Println("Successfully connect to the DB ")

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal("CATALOG: Automigrate: postgres.WithInstance error: ", err)
	}
	m, err := migrate.NewWithDatabaseInstance(migrations_path, "postgres", driver)
	if err != nil {
		log.Fatal("CATALOG: Automigrate: migrate.NewWithDatabaseInstance error: ", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("CATALOG: Automigrate: migrate.up error: ", err)
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
	item_by_brand_title := queries.New_catalog_item_by_brand_title_handler(items_repo)
	create_item := commands.New_create_catalog_item_handler(items_repo)
	update_item := commands.New_update_catalog_item_handler(items_repo)
	delete_item := commands.New_delete_catalog_item_handler(items_repo)
	catalog_items_v2 := queries.New_catalog_items_v2_handler(items_repo)

	brands_handler := handlers.New_brands_handler(list_brands)
	categories_handler := handlers.New_categories_handler(list_categories)
	items_handler := handlers.New_catalog_items_handler(
		list_items,
		item_by_id,
		item_by_title,
		item_by_brand_title,
		create_item,
		update_item,
		delete_item,
	)
	items_handler_v2 := handlers.New_catalog_items_handler_v2(catalog_items_v2)

	api.Register_routes(engine,
		brands_handler,
		categories_handler,
		items_handler,
		items_handler_v2,
	)
	log.Printf("Start server on port %s/n", app_port)

	if err := engine.Run(app_port); err != nil {
		log.Fatal(err.Error())
	}
}
