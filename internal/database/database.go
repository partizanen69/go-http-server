package database

import (
	"encoding/json"
	"errors"
	"os"
)

type Client struct {
	filePath string
}

func NewClient(filePath string) *Client {
	return &Client{
		filePath: filePath,
	}
}

func (c Client) createDb() error {
	data, err := json.Marshal(databaseSchema{
		Users: make(map[string]User),
		Posts: make(map[string]Post),
	})
	if err != nil {
		return err
	}
	err = os.WriteFile(c.filePath, data, 0600)
	if err != nil {
		return err
	}
	return nil
}

func (c Client) EnsureDB() error {
	_, err := os.ReadFile(c.filePath)
	if errors.Is(err, os.ErrNotExist) {
		return c.createDb()
	}
	return err
}

func (c Client) updateDb(db databaseSchema) error {
	data, err := json.Marshal(db)
	if err != nil {
		return err
	}
	 err = os.WriteFile(c.filePath, data, 0600)
	 return err
}

func (c Client) readDb() (databaseSchema, error) {
	byteData, err := os.ReadFile(c.filePath)
	if err != nil {
		return databaseSchema{}, err
	}
	
	var data databaseSchema
	if err = json.Unmarshal(byteData, &data); err != nil {
		return databaseSchema{}, err
	}
	
	return data, nil
}

type databaseSchema struct {
	Users map[string]User `json:"users"`
	Posts map[string]Post `json:"posts"`
}