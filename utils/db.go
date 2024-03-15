package utils

import (
	"crud-app/instance"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, errDB := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=root password=mysecretpassword dbname=gorm port=5432 sslmode=disable TimeZone=America/New_York",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	errMigrate := db.AutoMigrate(&instance.Item{})
	if errMigrate != nil {
		fmt.Println("An error occurred with auto migrating.")
	}
	if errDB != nil {
		fmt.Println("An error occurred init the DB.")
	}
	return db
}
