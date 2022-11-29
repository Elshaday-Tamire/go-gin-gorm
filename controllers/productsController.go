package controllers

import (
	"go_gin_gorm/initializers"
	"go_gin_gorm/models"

	"github.com/gin-gonic/gin"
)

func ProductsCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		Title       string
		Description string
		Base64      string
	}
	c.Bind(&body)

	//create a product
	product := models.Product{Title: body.Title, Description: body.Description, Base64: body.Base64}

	result := initializers.DB.Create(&product)

	if result.Error != nil {
		c.Status(400)
		return
	}
	//Return it
	c.JSON(200, gin.H{
		"product": product,
	})
}
func OrderCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		OrderName string
	}
	c.Bind(&body)

	//create a product
	order := models.Order{OrderName: body.OrderName}

	result := initializers.DB.Create(&order)

	if result.Error != nil {
		c.Status(400)
		return
	}
	//Return it
	c.JSON(200, gin.H{
		"order": order,
	})
}

func ProductsGetAll(c *gin.Context) {
	// Get the products
	var products []models.Product
	initializers.DB.Find(&products)

	//Respond with them
	c.JSON(200, gin.H{
		"products": products,
	})
}

func ProductsGetById(c *gin.Context) {
	// Get id off url
	id := c.Param("id")
	// Get the products
	var product models.Product
	initializers.DB.First(&product, id)

	//Respond with them
	c.JSON(200, gin.H{
		"product": product,
	})
}

func ProductUpdate(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")

	//get the data off req body
	var body struct {
		Title       string
		Description string
		Base64      string
	}
	c.Bind(&body)
	//find the post where updating is to be performed
	var product models.Product
	initializers.DB.First(&product, id)
	//update it
	initializers.DB.Model(&product).Updates(models.Product{
		Title: body.Title, Description: body.Description, Base64: body.Base64,
	})
	//respond with it
	c.JSON(200, gin.H{
		"product": product,
	})
}

func ProductDelete(c *gin.Context) {
	//Get the id off the url
	id := c.Param("id")
	//Delete the products
	initializers.DB.Delete(&models.Product{}, id)
	//Respond
	c.Status(200)
}
