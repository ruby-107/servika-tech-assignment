package main

import (
	"fmt"
	"log"
	"net/http"

	"project/controllers"
	"project/db"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize MongoDB connection
	err := db.InitMongoDB("mongodb://localhost:27017", "yourdb")
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	// Define your routes here
	r.HandleFunc("/pets", controllers.ListPets).Methods("GET")
	r.HandleFunc("/pets", controllers.CreatePet).Methods("POST")
	r.HandleFunc("/pets/{id}", controllers.GetPetAndEvents).Methods("GET")
	r.HandleFunc("/pets/{id}", controllers.EditPet).Methods("PUT")
	r.HandleFunc("/pets/{id}", controllers.AddEvent).Methods("POST")
	r.HandleFunc("/pets/{id}", controllers.DeletePet).Methods("DELETE")

	// Add other routes for reading, updating, and deleting pets

	http.Handle("/", r)
	fmt.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
