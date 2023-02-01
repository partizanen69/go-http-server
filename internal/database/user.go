package database

import (
	"errors"
	"time"
)

type User struct {
	CreatedAt time.Time `json:"createdAt"`
	Email string `json:"email"`
	Password string `json:"password"`
	Name string `json:"name"`
	Age int `json:"age"`
}

func (c Client) CreateUser(email, password, name string, age int) (User, error) {
	data, err := c.readDb()
	if err != nil {
		return User{}, err
	}
	data.Users[email] = User{
		Email: email,
		Password: password,
		Name: name,
		Age: age,
		CreatedAt: time.Now().UTC(),
	}
	err = c.updateDb(data)
	if err != nil {
		return User{}, err
	}
	return data.Users[email], nil
}

func (c Client) UpdateUser(email, password, name string, age int) (User, error) {
	data, err := c.readDb()
	if err != nil {
		return User{}, err
	}
	user, ok := data.Users[email]
	if ! ok {
		return User{}, errors.New("user doesn't exist")
	}

	user.Password = password
	user.Name = name
	user.Age = age
	
	data.Users[email] = user
	err = c.updateDb(data)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (c Client) GetUser(email string) (User, error) {
	data, err := c.readDb()
	if err != nil {
		return User{}, err
	}
	user, ok := data.Users[email]
	if ! ok {
		return User{}, errors.New("user doesn't exist")
	}
	return user, nil
}

func (c Client) DeleteUser(email string) error {
	data, err := c.readDb()
	if err != nil {
		return err
	}
	
	delete(data.Users, email)
	return c.updateDb(data)
}