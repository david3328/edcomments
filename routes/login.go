package routes

import (
	"github.com/david3328/edcomments/controllers"
	"github.com/gorilla/mux"
)

func SetLoginRouter(router *mux.Router) {
	router.HandleFunc("/api/login", controllers.Login).Methods("POST")
}
