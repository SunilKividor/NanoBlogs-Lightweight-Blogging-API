package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var database = "Medium"
var connection *mongo.Database

func GetDatabaseConnection() *mongo.Database {
	return connection
}

func GetCollection(collName string) *mongo.Collection {
	return connection.Collection(collName)
}

func ConnectToDB() {
	er := godotenv.Load()
	if er != nil {
		print("Error while loading env")
		return
	}
	connectionString := os.Getenv("MONGODB_URI")

	clientOptions := options.Client().ApplyURI(connectionString)

	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Print("Connection to database Successful")

	connection = client.Database(database)

}
