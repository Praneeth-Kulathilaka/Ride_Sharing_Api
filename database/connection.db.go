package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "#Saranga22908"
    dbname   = "ride_service"
)

var DB *gorm.DB

func ConnectDB()  {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Println("Error connecting to database",err)
        return
    }
    DB = db
    log.Println("connected to database successfully...")


}