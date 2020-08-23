package models

import (
	"errors"
	"fmt"
)

// User Struct ....
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

// return Users ....
func GetUsers() []*User {
	return users
}

// add User ....
func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("new User must not include id or must have id 0 ")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

// Return a User based on input id ....
func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("User id with %v not found ", id)
}

// Update a user .....
func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User id with %v not found ", u.ID)
}

// remove a user .....
func RemoveUserByID(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User id with %v not found ", id)
}
