package migration

import (
	"fmt"
	"github.com/yujen77300/API-test/database"
	"github.com/yujen77300/API-test/models"
	"log"
)

func Migration() {
	err := database.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Migration completed")
}