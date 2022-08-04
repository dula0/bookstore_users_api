package users
// data transfer object

import (
	"strings"

	"github.com/dula0/bookstore_users_api/utils/errors"
)

type User struct {
	ID          int64  `json:"id,omitempty"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (user *User) Validate() *errors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	return nil
}
