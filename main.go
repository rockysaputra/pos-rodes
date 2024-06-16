package main

import (
	"log"
	"pos-desy/config"
	"pos-desy/models"
	"pos-desy/routes"
)

func main() {
	config.ConnectDatabase()

	db := config.DB
	// migrate the schema
	db.AutoMigrate(&models.Category{}, &models.Item{}, &models.SellData{}, &models.SellDataSummary{}, &models.User{})

	r := routes.SetupRoute()
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server : ", err)
	}
}
