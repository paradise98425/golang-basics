package repository

import (
	"context"
	"log"

	"../entity"
	firestore "cloud.google.com/go/firestore/apiv1"
)

type PostRepository interface {
	Save(*entity.Post) (*entity.Post, error)
	// FindAll() ([]entity.Post, error)
}

type repo struct{}

const (
	projectId      string = "golang-basics"
	collectionName string = "posts"
)

//NewPostRepository

func NewPostRepository() PostRepository {
	return &repo{}
}

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to Create a Firestore Client: %v", err)
		return nil, err
	}
	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})
}
