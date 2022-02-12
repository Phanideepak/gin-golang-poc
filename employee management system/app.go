package main

import (
	"fmt"

	"github.com/deepak/ems/api/models"
	"github.com/deepak/ems/api/routes"
	"github.com/deepak/ems/config"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))

	if err != nil {
		fmt.Println(err)
	}
	defer config.DB.Close()
	config.DB.AutoMigrate(&models.Department{})

	r := routes.SetupRouter()
	r.Run()
}
