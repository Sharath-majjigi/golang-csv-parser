package database

import (
    "fmt"
    "log"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "sharath/models"
)

var db *gorm.DB

func Connect() {
    var err error

    db, err = gorm.Open("mysql", "YOUR_DB_CONNECTION_URL")

    if err != nil {
        log.Fatal(err)
    }

	fmt.Println("Database connected Sucessfully !")
    db.AutoMigrate(&models.Person{})
	fmt.Println("Database Migrated Sucessfully !")
}

func GetDB() *gorm.DB {
    return db
}