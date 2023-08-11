package initializers

import (
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(){
	var err error

	DB, err = gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database")
	}

}