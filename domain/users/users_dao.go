package users

import (
	"fmt"

	"github.com/dula0/bookstore_users_api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

// Retrieves user by their user ID or returns an error if there is one
func (user *User) Get() *errors.RestErr {
	result := usersDB[user.ID]
	if result == nil {
		return errors.NotFoundError(fmt.Sprintf("user %d not found", user.ID))
	}
	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

// Method to save user into db
func (user *User) Save() *errors.RestErr {
	current := usersDB[user.ID]
	// Check to see if we have a user with a matching ID in the db
	if current != nil {
		if current.Email == user.Email {
			return errors.BadRequestError(fmt.Sprintf("user email %s already exists", user.Email))
		}
		return errors.BadRequestError(fmt.Sprintf("user %d already exists", user.ID))
	}
	usersDB[user.ID] = user
	return nil
}
