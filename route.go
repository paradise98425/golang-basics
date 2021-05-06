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
		resp.Write([]byte(`{"error": "error marshalling the posts array"}`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	resp.Write(result)
}

// function to add the new post 
func addPosts(resp http.ResponseWriter, req *http.Request) {
	var post Post 	// variable post is of type Post struct 
	err := json.NewDecoder(req.Body).Decode(&posts)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error" : "error unmarshalling the request"}`))
		return
	}
	// creating the post id
	post.Id = len(posts) + 1
	// adding the post to the array of posts
	posts = append(posts, post)
	result, err2 := json.Marshal(post)
	if err2 != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error" : "error marshalling the response"}`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	resp.Write(result)
}	
