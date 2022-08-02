package users

import (
	"fmt"
	"strings"

	"github.com/dula0/bookstore_users_api/databases/mysql/users_db"
	"github.com/dula0/bookstore_users_api/utils/date_utils"
	"github.com/dula0/bookstore_users_api/utils/errors"
)

const (
	uniqueEmail     = "email_UNIQUE" // duplicate email error response
	InsertUserQuery = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
)

var (
	usersDB = make(map[int64]*User)
)

// Retrieves user by their user ID or returns an error if there is one
func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
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

	stmt, err := users_db.Client.Prepare(InsertUserQuery)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), uniqueEmail) {
			return errors.BadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.InternalServerError(fmt.Sprintf("error while trying to save user: %s", err.Error()))
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.InternalServerError(fmt.Sprintf("error while trying to save user: %s", err.Error()))
	}
	user.ID = userId
	return nil
}
