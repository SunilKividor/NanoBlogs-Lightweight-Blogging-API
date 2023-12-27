package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"go-mongodb/database"
	"go-mongodb/models"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func (b *blog) GetUsers(rw http.ResponseWriter, r *http.Request) {
	fmt.Print("Entered GetUsers block")
	filter := bson.D{{}}

	connection := database.GetCollection("Users")

	cur, err := connection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Could not get the users")
		return
	}
	var result []models.User

	er := cur.All(context.TODO(), &result)
	if er != nil {
		http.Error(rw, "Could not get the users", http.StatusInternalServerError)
		return
	}

	// for _, user := range result {
	// 	err := json.NewEncoder(rw).Encode(user)
	// 	if err != nil {
	// 		log.Fatal("Could not marshal to JSON")
	// 	}
	// }
	displey, err := json.Marshal(result)
	if err != nil {
		http.Error(rw, "Could not Marshal the Users to JSON", http.StatusInternalServerError)
		return
	}
	rw.Write(displey)

}

func (b *blog) PostNewUser(rw http.ResponseWriter, r *http.Request) {
	body := models.GetUser()

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(rw, "Invalid Input", http.StatusBadRequest)
		return
	}

	err = body.ValidateUser()
	if err != nil {
		http.Error(rw, "Invalid Input", http.StatusBadRequest)
		return
	}

	connection := database.GetDatabaseConnection()

	result, er := connection.Collection("Users").InsertOne(context.TODO(), body)
	if er != nil {
		http.Error(rw, "Could not add User", http.StatusInternalServerError)
		return
	}

	fmt.Print(result.InsertedID)
}
