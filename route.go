package main

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

var (
	posts []Post // creats a variable named posts which is an array of struct Post
)

func init() {
	// assign a post array
	posts = []Post{Post{Id: 1, Title: "Title 1", Text: "Text"}}
}

func getPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("content-type", "application/json")
	result, err := json.Marshal(posts)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"message": "error marshalling the posts array"}`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	resp.Write(result)
}
