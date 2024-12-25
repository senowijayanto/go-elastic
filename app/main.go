package main

import (
	"go-elastic/config"
	"go-elastic/handlers"
	"go-elastic/migrations"
	"go-elastic/repositories"
	"go-elastic/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Setup database
	config.SetupDatabase()

	// Run migrations
	migrations.RunMigrations()

	// Initialize components
	productRepo := repositories.NewProductRepository()
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	// Setup Gin
	r := gin.Default()

	// Routes
	productGroup := r.Group("/products")
	{
		productGroup.POST("/", productHandler.CreateProduct)
		productGroup.GET("/", productHandler.GetAllProducts)
		productGroup.PUT("/:id", productHandler.UpdateProduct)
		productGroup.DELETE("/:id", productHandler.DeleteProduct)
	}

	// Run server
	r.Run(":8080")
}
