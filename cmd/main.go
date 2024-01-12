// entry point
package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/theyashwanthsai/Blog-Backend-Go/internals/controller"
)

func main(){
	// define routes and handler functions
	r := mux.NewRouter()


	// routes
	r.HandleFunc("/home", controller.Home).Methods("GET")
	r.HandleFunc("/posts", controller.getAllPosts).Methods("GET")
	r.HandleFunc("/post/{id}", controller.getPostByID).Methods("GET")
	r.HandleFunc("/post", controller.addPost).Methods("POST")
	r.HandleFunc("/post/{id}", controller.updatePost).Methods("PUT")
	r.HandleFunc("/post/{id}", controller.deletePost).Methods("DELETE")


	fmt.Println("Server running on 8080:")
	log.Fatal(http.ListenAndServe(":8080", r))
}