package models

type Rides struct {
	RideId int `gorm:"primaryKey"`
	VehicleType string 
	PickupLocation string
	DropoffLocation string
	Distance float64
	EstimatedFare float64
	CustomerId int `gorm:"index"`
	RiderId int `gorm:"index"`
	
}

type Customers struct {
	CustomerId   int    `gorm:"primaryKey"`
	CustomerName string
	MobileNo     string
	Rides        []Rides `gorm:"foreignKey:CustomerId"`
}

type Riders struct {
	RiderId int `gorm:"primaryKey"`
	RiderName string
	ContactNumber string
	VehicleNumber string
	VehicleType string
	Rides []Rides `gorm:"foreignKey:RiderId"`
}