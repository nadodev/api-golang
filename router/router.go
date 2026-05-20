package router

import (
	"database/sql"
	"go-project/controller"
	"go-project/db"
	"go-project/repository"
	"go-project/usecase"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()
	dbConnection, _ := db.ConnectDB()
	productRepository := repository.NewProductRepository(dbConnection)
	productUseCase := usecase.NewProductUseCase(productRepository)
	productController := controller.NewProductController(productUseCase)

	router.GET("/products", productController.GetProducts)
	router.GET("/products/:id", productController.GetProductById)
	err := router.Run(":8080")

	if err != nil {
		return
	}

	defer func(dbConnection *sql.DB) {
		err := dbConnection.Close()
		if err != nil {
			panic(err)
		}
	}(dbConnection)
}
