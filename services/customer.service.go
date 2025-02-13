package services

import (
	"RideSharingApi/database"
	"RideSharingApi/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func AddCustomer(w http.ResponseWriter, r *http.Request) {
	customer := models.Customers{}
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		log.Println("Error getting customer data..",err)
		return
	}
	customer.CustomerId = uuid.New().ClockSequence()

	result := database.DB.Create(&customer)

	if result.Error != nil {
		log.Println("Error handling database..",err)
		return
	}
	log.Println("Customer succussfully stored",result)
}