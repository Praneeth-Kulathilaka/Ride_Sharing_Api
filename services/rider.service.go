package services

import (
	"RideSharingApi/database"
	"RideSharingApi/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func AddRider(w http.ResponseWriter, r *http.Request) {
	rider := models.Riders{}
	err := json.NewDecoder(r.Body).Decode(&rider)
	if err != nil {
		log.Println("Error getting customer data..",err)
		return
	}
	rider.RiderId = uuid.New().ClockSequence()

	result := database.DB.Create(&rider)

	if result.Error != nil {
		log.Println("Error handling database..",err)
		return
	}
	log.Println("Rider succussfully stored",result)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message":"Rider successfully added"})
}