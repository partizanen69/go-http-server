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

func (c Client) GetPosts(userEmail string) ([]Post, error) {
	data, err := c.readDb()
	if err != nil {
		return nil, err
	}

	posts := []Post{}
	for _, post := range data.Posts {
		if post.UserEmail == userEmail {
			posts = append(posts, post)
		}
	}

	return posts, nil
}

func (c Client) DeletePost(id string) error {
	data, err := c.readDb()
	if err != nil {
		return err
	}

	delete(data.Posts, id)
	return c.updateDb(data)
}