package models

type RideResponse struct {
	RideId          int     	`json:"ride_id"`
	VehicleType     string  	`json:"vehicle_type"`
	PickupLocation  string  	`json:"pickup_location"`
	DropoffLocation string  	`json:"dropoff_location"`
	Distance        float64 	`json:"distance"`
	EstimatedFare   float64 	`json:"estimated_fare"`
}