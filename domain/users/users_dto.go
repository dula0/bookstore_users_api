package users

// data transfer object

import (
	"strings"

	"github.com/dula0/bookstore_users_api/utils/errors"
)

const (
	StatusActive = "active"
)

type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

type Users []User

func (user *User) Validate() *errors.RestErr {
	user.FirstName = strings.TrimSpace(strings.ToLower(user.FirstName))
	user.LastName = strings.TrimSpace(strings.ToLower(user.LastName))

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.BadRequestError("Email address cannot be blank")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.BadRequestError("invalid password")
	}
	return nil
}
