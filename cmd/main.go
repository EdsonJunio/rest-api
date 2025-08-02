package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"rest-api/controller"
	"rest-api/db"
	"rest-api/repository"
	"rest-api/usecase"
)

func main() {
	_ = godotenv.Load()
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// Camada de repository
	productRepository := repository.NewProductRepository(dbConnection)
	// Camada usecase (agora retorna ponteiro)
	productUsecase := usecase.NewProductUsecase(productRepository)
	// Camada de controllers (agora retorna ponteiro)
	productController := controller.NewProductController(productUsecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", productController.GetProducts)
	server.GET("/product/:productId", productController.GetProductByID)
	server.POST("/product", productController.CreateProduct)
	server.PUT("/products/:id", productController.UpdateProductByID)
	server.DELETE("/products/:id", productController.DeleteProductByID)

	server.Run(":8080")
}
