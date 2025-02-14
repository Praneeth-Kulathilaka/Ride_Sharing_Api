package main

import (
	"RideSharingApi/database"
	"RideSharingApi/routers"
	"fmt"
	"net/http"
)

func main() {
	database.ConnectDB()

	r := routers.SetRoutes()
	database.Migrate()
	
	
	port := ":8080"
	fmt.Println("Server is running on port", port)
	http.ListenAndServe(port, r)
}