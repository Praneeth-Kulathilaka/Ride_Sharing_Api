package routers

import (
	"RideSharingApi/services"

	"github.com/gorilla/mux"
)

func SetRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/customer",services.AddCustomer).Methods("POST")

	return r

}