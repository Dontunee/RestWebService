package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(user User) (User, error) {
	if user.ID != 0 {
		return User{}, errors.New("user must not include ID")
	}

	user.ID = nextID
	nextID++
	users = append(users, &user)
	return user, nil
}

func GetUserById(id int) (User, error) {
	for _, v := range users {
		if v.ID == id {
			return *v, nil
		}
	}
	return User{}, fmt.Errorf("user with ID '%v' not found", id)
}

func UpdateUser(user User) (User, error) {
	for i, v := range users {
		if v.ID == user.ID {
			users[i] = &user
			return *users[i], nil
		}
	}
	return User{}, fmt.Errorf("user with ID '%v' does not exist", user.ID)
}

func RemoveUserById(id int) error {
	for i, v := range users {
		if v.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("user with id '%v' does not exist", id)
}
