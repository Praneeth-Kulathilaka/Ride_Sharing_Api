package database

import (
	// "RideSharingApi/database"
	"RideSharingApi/models"
	"log"
)

func Migrate() {
	err := DB.AutoMigrate(&models.Customers{})

	if err != nil {
		log.Println("Error migrating tables", err)
	}

	err = DB.AutoMigrate(&models.Riders{})

	if err != nil {
		log.Println("Error migrating customer table", err)
	}
	err = DB.AutoMigrate(&models.Rides{})

	if err != nil {
		log.Println("Error migrating rider tables", err)
	}
	log.Println("All the tables successfully migrated")
}