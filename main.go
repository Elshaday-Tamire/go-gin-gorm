package main

import (
	"go_gin_gorm/controllers"
	"go_gin_gorm/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	r := gin.Default()

	r.POST("/products", controllers.ProductsCreate)
	r.GET("/products", controllers.ProductsGetAll)
	r.GET("/products/:id", controllers.ProductsGetById)
	r.PUT("/products/:id", controllers.ProductUpdate)
	r.DELETE("/products/:id", controllers.ProductDelete)
	r.POST("/orders", controllers.OrderCreate)
	r.Run()
}
