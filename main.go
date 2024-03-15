package main

import (
	"crud-app/instance"
	"crud-app/routers"
	"crud-app/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/gorm"
)

func main() {
	var DB *gorm.DB = utils.InitDB()
	var itemManager instance.ItemManager = instance.ItemManager{Db: DB, Items: make([]instance.Item, 0)}

	//itemManager.AddItem(instance.Item{Id: 3243, Name: "Item1", Category: "ItemCat", Value: 100})
	r := gin.Default()
	// api v1
	v1 := r.Group("/api/version/primary")
	v1.GET("health", routers.GetHealth)
	v1.GET("items", utils.ItemManagerWrapper(routers.GetAllItems, itemManager))
	v1.GET("update", utils.ItemManagerWrapper(routers.GetTest, itemManager))

	err := r.Run()
	if err != nil {
		fmt.Println(err)
	}

}
