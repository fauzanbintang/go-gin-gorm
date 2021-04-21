package controllers

import (
	"github.com/fauzanbintang/go-gin-gorm/models"
	"github.com/gin-gonic/gin"
)

func GetOrders(c *gin.Context) {
	db := models.GetDB()

	var Orders []models.Order

	err := db.Debug().Preload("Items").Find(&Orders).Error

	if err != nil {
		c.JSON(500, gin.H{
			"messages": "Internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"orders": Orders,
	})
}

func GetItems(c *gin.Context) {
	db := models.GetDB()

	var Items []models.Item

	err := db.Debug().Find(&Items).Error

	if err != nil {
		c.JSON(500, gin.H{
			"messages": "Internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"items": Items,
	})
}

func CreateOrder(c *gin.Context) {
	db := models.GetDB()

	var Order models.Order
	// var Item = []models.Item{}

	if err := c.ShouldBindJSON(&Order); err != nil {
		c.JSON(500, gin.H{
			"messages": "error parse",
		})
		return
	}

	errOrder := db.Debug().Create(&Order).Error

	if errOrder != nil {
		c.JSON(500, gin.H{
			"messages": "error create",
		})
	}

	c.JSON(201, gin.H{
		"data": Order,
		// "item": Item,
	})
}

func UpdateOrder(c *gin.Context) {
	db := models.GetDB()

	var Order models.Order
	var OrderInput models.Order

	var OrderID = c.Param("orderId")

	if err := c.ShouldBindJSON(&OrderInput); err != nil {
		c.JSON(500, gin.H{
			"messages": "error parse",
		})
		return
	}

	err := db.Debug().Model(&Order).Where("id = ?", OrderID).Updates(models.Order{
		CustomerName: OrderInput.CustomerName,
		OrderedAt:    OrderInput.OrderedAt,
	}).Error

	if err != nil {
		c.JSON(500, gin.H{
			"messages": "error update",
			"err":      err,
		})
		return
	}

	c.JSON(200, gin.H{
		"messages": "success update",
		"data":     Order,
	})
}

func DeleteOrder(c *gin.Context) {
	db := models.GetDB()

	var OrderID = c.Param("orderId")
	var Order models.Order

	err := db.Debug().Delete(&Order, OrderID).Error

	if err != nil {
		c.JSON(500, gin.H{
			"messages": "error delete",
		})
		return
	}

	c.JSON(200, gin.H{
		"messages": "success delete",
		"data":     Order,
	})
}
