package routers

import (
	"crud-app/instance"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"state": "healthy"})
}

func GetAllItems(c *gin.Context, i instance.ItemManager) {
	c.JSON(http.StatusOK, gin.H{"data": i.GetItems()})
}

func GetTest(c *gin.Context, i instance.ItemManager) {
	var item = i.GetItem(3243)
	item.UName("Best Item")
	item.Save(&i)
	c.JSON(http.StatusOK, gin.H{"data": item})
}
