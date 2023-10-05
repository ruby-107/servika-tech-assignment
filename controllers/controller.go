package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"project/db"
	models "project/model"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func ListPets(w http.ResponseWriter, r *http.Request) {
	// Get the optional query parameter "species"
	species := r.URL.Query().Get("species")

	filter := bson.M{}
	if species != "" {
		filter["species"] = species
	}

	petCollection := db.Database.Collection("pets")
	cur, err := petCollection.Find(context.TODO(), filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cur.Close(context.TODO())

	var pets []models.Pet
	for cur.Next(context.TODO()) {
		var pet models.Pet
		if err := cur.Decode(&pet); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		pets = append(pets, pet)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pets)
}

func CreatePet(w http.ResponseWriter, r *http.Request) {
	var pet models.Pet
	if err := json.NewDecoder(r.Body).Decode(&pet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	collection := db.Database.Collection("pets")
	_, err := collection.InsertOne(context.TODO(), pet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pet)
}

func GetPetAndEvents(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	pet := models.Pet{}
	collection := db.Database.Collection("pets")
	err := collection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&pet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pet)
}

func EditPet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var updatedPet models.Pet
	if err := json.NewDecoder(r.Body).Decode(&updatedPet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	collection := db.Database.Collection("pets")
	filter := bson.M{"id": id}
	update := bson.M{"$set": updatedPet}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedPet)
}

func AddEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var newEvent models.Event
	if err := json.NewDecoder(r.Body).Decode(&newEvent); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	petCollection := db.Database.Collection("pets")
	petFilter := bson.M{"id": id}
	pet := models.Pet{}
	err := petCollection.FindOne(context.TODO(), petFilter).Decode(&pet)
	if err != nil {
		http.Error(w, "Pet not found", http.StatusNotFound)
		return
	}

	pet.Events = append(pet.Events, newEvent)

	// Update the pet in MongoDB with the new event
	petUpdate := bson.M{"$set": pet}
	_, err = petCollection.UpdateOne(context.TODO(), petFilter, petUpdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the added event
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newEvent)
}

func DeletePet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	collection := db.Database.Collection("pets")
	_, err := collection.DeleteOne(context.TODO(), bson.M{"id": id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Pet deleted successfully"))
}
