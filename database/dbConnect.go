package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var DB *gorm.DB

var err error

func DatabaseInit(){
	dsn := "root:1qaz@WSX@tcp(127.0.0.1:3306)/senaoapi?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connected")
}
