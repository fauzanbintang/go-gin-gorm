package main

import (
	"github.com/fauzanbintang/go-gin-gorm/controllers"
	"github.com/fauzanbintang/go-gin-gorm/models"
	"github.com/gin-contrib/cors" // cors
	"github.com/gin-gonic/gin"    // gin
	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"
)

func main() {
	models.Init() // init db

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		// AllowHeaders:     []string{"Origin"},
		// ExposeHeaders:    []string{"Content-Length"},
		// AllowCredentials: true,
	}))

	// db := models.SetupModels()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "hello"})
	})

	orders := router.Group("/orders")
	{
		orders.GET("/", controllers.GetOrders)
		orders.POST("/", controllers.CreateOrder)
		orders.PUT("/:orderId", controllers.UpdateOrder)
		orders.DELETE("/:orderId", controllers.DeleteOrder)
	}

	items := router.Group("/items")
	{
		items.GET("/", controllers.GetItems)
	}

	router.Run()
}
