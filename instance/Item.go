package instance

type Item struct {
	Id       int `gorm:"primaryKey"`
	Name     string
	Category string
	Value    int
}

func (i *Item) UName(n string) {
	i.Name = n
}

func (i *Item) UCategory(c string) {
	i.Category = c
}

func (i *Item) UValue(v int) {
	i.Value = v
}

func (i *Item) Save(im *ItemManager) {
	im.Db.Save(&i)
}
