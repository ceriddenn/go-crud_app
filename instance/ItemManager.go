package instance

import (
	"fmt"
	"gorm.io/gorm"
)

type ItemManager struct {
	Db    *gorm.DB
	Items []Item
}

func (im *ItemManager) AddItem(i Item) {
	c := im.Db.Create(&i)
	fmt.Println(c.Error)
	im.Items = append(im.Items, i)
}

func (im *ItemManager) GetItems() []Item {
	var allItems []Item

	im.Db.First(&allItems)
	im.Items = allItems
	return im.Items
}

func (im *ItemManager) GetItem(p int) Item {
	var queriedItem Item
	im.Db.Where("Id = ?", p).Find(&queriedItem)
	return queriedItem
}
