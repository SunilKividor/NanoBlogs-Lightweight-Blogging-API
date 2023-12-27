package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"go-mongodb/database"
	"go-mongodb/helpers"
	"go-mongodb/models"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type blog struct{}
type Blog models.Blog

func NewBlog() *blog {
	return &blog{}
}

// get blog by userId
func (b *blog) GetAllBlogByUserID(rw http.ResponseWriter, r *http.Request) {

	objId, err := helpers.GetObjectId(r, "id")
	if err != nil {
		http.Error(rw, "Inavlid Id", http.StatusBadRequest)
		return
	}

	filter := bson.D{{Key: "user_id", Value: objId}}
	connection := database.GetCollection("Blogs")

	er := helpers.GetBlogs(connection, filter, rw)
	if er != nil {
		http.Error(rw, "Could not get the blogs", http.StatusInternalServerError)
	}
}

// Get All Blogs
func (b *blog) GetRandomeBlogsForUser(rw http.ResponseWriter, r *http.Request) {

	connectionBlogs := database.GetCollection("Blogs")
	connectionUser := database.GetCollection("Users")
	objId, err := helpers.GetObjectId(r, "id")
	if err != nil {
		http.Error(rw, "Inavlid Id", http.StatusBadRequest)
		return
	}

	//Take blogs according to Users favourite tags
	filter := bson.M{"_id": objId}
	projection := bson.M{"tags": 1}
	var user models.User

	err = connectionUser.FindOne(context.TODO(), filter, options.FindOne().SetProjection(projection)).Decode(&user)
	if err != nil {
		http.Error(rw, "Could not get the blogs", http.StatusInternalServerError)
		return
	}
	favTags := user.Tags
	if favTags == nil {
		log.Println("Could not get the tags")
		return
	}

	filte := bson.M{"tags": bson.M{"$in": favTags}, "user_id": bson.M{"$ne": objId}}
	err = helpers.GetBlogs(connectionBlogs, filte, rw)
	if err != nil {
		log.Println(err)
		http.Error(rw, "oops there were none fav blogs", http.StatusBadRequest)
		return
	}

}

// Post One blog
func (b *blog) PostBlog(rw http.ResponseWriter, r *http.Request) {

	newBlog := models.GetBlog()

	err := json.NewDecoder(r.Body).Decode(&newBlog)
	if err != nil {
		http.Error(rw, "Invalid request body", http.StatusBadRequest)
		return
	}

	if e := newBlog.ValidateBlog(); e != nil {
		http.Error(rw, "Invalid Input", http.StatusBadRequest)
		return
	}

	connection := database.GetCollection("Blogs")

	result, er := connection.InsertOne(context.TODO(), newBlog)
	if er != nil {
		http.Error(rw, "Error inserting Blog", http.StatusBadRequest)
		return
	}

	fmt.Print("Inserted with ID:", result.InsertedID)
}

// update blog by blog id
func (b *blog) UpdateBlog(rw http.ResponseWriter, r *http.Request) {

	objectId, err := helpers.GetObjectId(r, "id")
	if err != nil {
		http.Error(rw, "Invalid ID", http.StatusBadRequest)
		return
	}

	updateBlog := models.GetBlog()
	er := json.NewDecoder(r.Body).Decode(&updateBlog)
	if er != nil {
		http.Error(rw, "Could not decode the JSON", http.StatusNotAcceptable)
		return
	}

	filter := bson.D{{Key: "_id", Value: objectId}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "title", Value: updateBlog.Title}, {Key: "body", Value: updateBlog.Body}}}}

	connection := database.GetCollection("Blogs")
	result, e := connection.UpdateOne(context.TODO(), filter, update)
	if e != nil {
		http.Error(rw, "Error updating the blog", http.StatusBadRequest)
	}

	if result.ModifiedCount == 0 {
		http.Error(rw, "Could not update the Blog", http.StatusNotFound)
		return
	}

	rw.Write([]byte("The Updation was successful"))
}

// delete blog by blog id
func (b *blog) DeleteBlog(rw http.ResponseWriter, r *http.Request) {

	objId, err := helpers.GetObjectId(r, "id")
	if err != nil {
		http.Error(rw, "Invalid ID", http.StatusBadRequest)
		return
	}

	filter := bson.D{{Key: "_id", Value: objId}}

	connection := database.GetDatabaseConnection()

	result, er := connection.Collection("Blogs").DeleteOne(context.TODO(), filter)
	if er != nil {
		http.Error(rw, "Error while Deleting", http.StatusInternalServerError)
		return
	}
	if result.DeletedCount == 0 {
		http.Error(rw, "Could not find the document", http.StatusBadRequest)
		return
	}
	rw.Write([]byte("Successully Deleted"))

}
