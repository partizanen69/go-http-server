package database

import (
	"errors"
	"time"

	"github.com/google/uuid"
)



type Post struct {
	ID string `json"id"`
	CreatedAt time.Time `json:"createdAt"`
	UserEmail string `json:"userEmail"`
	Text string `json:"text"`
}


func (c Client) CreatePost(userEmail, text string) (Post, error) {
	data, err := c.readDb()
	if err != nil {
		return Post{}, err
	}
	
	_, ok := data.Users[userEmail]
	if !ok {
		return Post{}, errors.New("user doesn't exist")
	}
	
	post := Post{
		ID: uuid.New().String(),
		CreatedAt: time.Now().UTC(),
		UserEmail: userEmail,
		Text: text,
	}
	data.Posts[post.ID] = post

	err = c.updateDb(data)
	if err != nil {
		return Post{}, err
	}
	
	return post, nil
}
