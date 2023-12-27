package routes

import (
	"go-mongodb/handlers"

	"github.com/gorilla/mux"
)

func BlogRoutes(router *mux.Router) {
	blog := handlers.NewBlog()

	router.HandleFunc("/getusers", blog.GetUsers).Methods("GET")
	router.HandleFunc("/{id}", blog.GetRandomeBlogsForUser).Methods("GET")
	router.HandleFunc("/get/{id}", blog.GetAllBlogByUserID).Methods("GET")
	router.HandleFunc("/comments/{id}", blog.GetComments).Methods("GET")
	router.HandleFunc("/newuser", blog.PostNewUser).Methods("POST")
	router.HandleFunc("/post", blog.PostBlog).Methods("POST")
	router.HandleFunc("/comment/{id}", blog.PostComments).Methods("POST")
	router.HandleFunc("/put/{id}", blog.UpdateBlog).Methods("PUT")
	router.HandleFunc("/delete/{id}", blog.DeleteBlog).Methods("DELETE")
}
