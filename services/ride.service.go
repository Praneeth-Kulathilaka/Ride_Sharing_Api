package services

import (
	"RideSharingApi/database"
	"RideSharingApi/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func AddRide(w http.ResponseWriter, r *http.Request) {
	ride := models.Rides{}
	err := json.NewDecoder(r.Body).Decode(&ride)
	pricePerKm := 30.00
	if err != nil {
		http.Error(w, "An error occured",http.StatusBadRequest)
		log.Println("Error getting ride data", err)
		return
	}
	ride.RideId = uuid.New().ClockSequence()
	ride.EstimatedFare = ride.Distance*pricePerKm
	result := database.DB.Create(&ride)
	if result.Error != nil {
		http.Error(w, "Internal server error",http.StatusInternalServerError)
		log.Println("Error handling database..",err)
		return
	}

	rideResponse := models.RideResponse{
        RideId:         ride.RideId,
        VehicleType:    ride.VehicleType,
        PickupLocation: ride.PickupLocation,
        DropoffLocation: ride.DropoffLocation,
        Distance:       ride.Distance,
        EstimatedFare:  ride.EstimatedFare,
    }
	
	log.Println("Ride succussfully stored",result)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rideResponse)

}