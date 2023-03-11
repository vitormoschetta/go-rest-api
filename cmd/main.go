package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/vitormoschetta/go-rest-api/internal/config"
	"github.com/vitormoschetta/go-rest-api/internal/controllers"
	"github.com/vitormoschetta/go-rest-api/internal/repositories"
)

func main() {
	config.LoadConfig()
	dbConnection := conectDb()
	productRepository := repositories.NewProductRepository(dbConnection)
	productController := controllers.NewProductController(productRepository)

	router := gin.Default()
	router.GET("/products", productController.FindAll)
	router.GET("/products/:id", productController.FindByID)
	router.POST("/products", productController.Create)
	router.PUT("/products/:id", productController.Update)
	router.DELETE("/products/:id", productController.Delete)

	router.Run()
}

func conectDb() *sql.DB {
	var cfg = mysql.Config{
		User:                 os.Getenv("DATABASE_USER"),
		Passwd:               os.Getenv("DATABASE_PASSWORD"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DATABASE_HOST") + ":" + os.Getenv("DATABASE_PORT"),
		DBName:               os.Getenv("DATABASE_NAME"),
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected to database")
	return db
}
