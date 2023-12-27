package main

import (
	"go-mongodb/database"
	"go-mongodb/routes"
	"os"

	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	database.ConnectToDB()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Could not load the environment variables")
		return
	}

	port := os.Getenv("PORT")

	router := mux.NewRouter()

	routes.BlogRoutes(router)

	s := http.Server{
		Addr:         (":" + port),
		Handler:      router,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	log.Print("Port Running on 8080")
	s.ListenAndServe()
}
