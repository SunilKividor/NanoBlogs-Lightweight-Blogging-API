package helpers

import (
	"context"
	"encoding/json"
	"go-mongodb/models"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetObjectId(r *http.Request, idFormat string) (interface{}, error) {
	vars := mux.Vars(r)
	id := vars[idFormat]

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return objId, err
	}
	return objId, nil
}

func GetBlogs(connection *mongo.Collection, filter interface{}, rw io.Writer) error {

	options := options.Find().SetSort(bson.D{{Key: "time", Value: -1}})

	var result []models.Blog
	cur, err := connection.Find(context.TODO(), filter, options)
	if err != nil {
		return err
	}

	ee := cur.All(context.TODO(), &result)
	if ee != nil {
		return ee
	}

	displey, e := json.Marshal(result)
	if e != nil {
		log.Fatal("Could not marshal to JSON")
		return e
	}
	rw.Write(displey)
	return nil
}

func GetTime() string {
	ist, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		log.Fatal(err)
		return time.Now().String()
	}

	currentTime := time.Now().In(ist)

	return currentTime.Format("Mon, 02 Jan 2006 15:04:05 MST")
}
