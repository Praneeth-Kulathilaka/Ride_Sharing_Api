package models

type RideStatus string

const (
	StatusActive RideStatus = "active"
	StatusPending RideStatus = "pending"
	StatusCompleted RideStatus = "completed"
)
type Rides struct {
	RideId          int     	`gorm:"primaryKey" json:"ride_id"`
	VehicleType     string  	`json:"vehicle_type"`
	PickupLocation  string  	`json:"pickup_location"`
	DropoffLocation string  	`json:"dropoff_location"`
	Distance        float64 	`json:"distance"`
	EstimatedFare   float64 	`json:"estimated_fare"`
	PaymentId       string  	`gorm:"default:'pending'" json:"payment_id"`
	Status 			RideStatus 	`gorm:"type:VARCHAR(20); default:'pending'" json:"status"`
	CustomerId      int     	`gorm:"index" json:"customer_id"`
	RiderId         *int     	`gorm:"index" json:"rider_id"`
}

type Customers struct {
	CustomerId   int     `gorm:"primaryKey" json:"customer_id"`
	CustomerName string  `json:"customer_name"`
	MobileNo     string  `json:"mobile_no"`
	Rides        []Rides `gorm:"foreignKey:CustomerId" json:"rides"`
}

type Riders struct {
	RiderId       int     `gorm:"primaryKey" json:"rider_id"`
	RiderName     string  `json:"rider_name"`
	ContactNumber string  `json:"contact_number"`
	VehicleNumber string  `json:"vehicle_number"`
	VehicleType   string  `json:"vehicle_type"`
	Rides         []Rides `gorm:"foreignKey:RiderId" json:"rides"`
}
