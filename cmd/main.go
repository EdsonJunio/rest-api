package main

import (
	"github.com/gin-gonic/gin"
	"rest-api/controller"
	"rest-api/db"
	"rest-api/repository"
	"rest-api/usecase"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	//Camada de repository
	ProductRepository := repository.NewProductRepository(dbConnection)
	//Camada usecase
	ProductUsercase := usecase.NewProductUseCase(ProductRepository)
	//Camadas de controllers
	ProductController := controller.NewproductController(ProductUsercase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)
	server.GET("/product/:productId", ProductController.GetProductById)

	server.Run(":8080")
}
