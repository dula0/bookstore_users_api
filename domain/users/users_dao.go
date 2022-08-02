package users

import (
	"github.com/dula0/bookstore_users_api/databases/mysql/users_db"
	"github.com/dula0/bookstore_users_api/utils/date_utils"
	"github.com/dula0/bookstore_users_api/utils/errors"
	"github.com/dula0/bookstore_users_api/utils/mysql_utils"
)

const (
	insertUserQuery = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	getUserQuery    = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
)

// Retrieves user by their user ID or returns an error if there is one
func (user *User) Get() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(getUserQuery)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer stmt.Close()

	// Holds the sql query result
	result := stmt.QueryRow(user.ID)

	// copy column values into struct fields.
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		return mysql_utils.ParseError(getErr)
	}

	return nil
}

// Method to save user into db
func (user *User) Save() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(insertUserQuery)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)

	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	user.ID = userId
	return nil
}
