package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter() // creating the new instance of mux router

	const port string = ":8081" // defining the port

	//	Handler function of mux router to server the http request on path "/"
	r.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Up and running...")
	})
	r.HandleFunc("/posts", getPosts).Methods("GET")
	r.HandleFunc("/posts", addPosts).Methods("POST")
	log.Println("server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, r))
}
