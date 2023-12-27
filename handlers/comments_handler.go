package handlers

import (
	"context"
	"encoding/json"

	"go-mongodb/database"
	"go-mongodb/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (b *blog) GetComments(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(rw, "Invalid Id", http.StatusBadRequest)
		return
	}

	connection := database.GetDatabaseConnection()

	var result models.BlogComments
	filter := bson.D{{Key: "blog_id", Value: objId}}

	er := connection.Collection("Comments").FindOne(context.TODO(), filter).Decode(&result)
	if er != nil {
		http.Error(rw, "Could not get the comments", http.StatusInternalServerError)
		return
	}

	for _, content := range result.Comments {
		err := json.NewEncoder(rw).Encode(content)
		if err != nil {
			log.Fatal("Could not marhsal to JSON")
		}
	}
}

func (b *blog) PostComments(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(rw, "Invalid Id", http.StatusBadRequest)
		return
	}

	var com models.Comment
	er := json.NewDecoder(r.Body).Decode(&com)
	if er != nil {
		http.Error(rw, "Invalid Comment body", http.StatusBadRequest)
		return
	}

	collection := database.GetDatabaseConnection()

	filter := bson.D{{Key: "blog_id", Value: objId}}
	update := bson.D{{Key: "$push", Value: bson.D{{Key: "comments", Value: com}}}}

	upsert := true

	result, e := collection.Collection("Comments").UpdateOne(context.TODO(), filter, update, &options.UpdateOptions{
		Upsert: &upsert,
	})
	if e != nil {
		http.Error(rw, "Could not add the comment sorry", http.StatusInternalServerError)
		return
	}

	log.Print(result.UpsertedID)

}
