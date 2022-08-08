package users

// data access object serves as the persistence layer so any interactions with the db occurs here

import (
	"fmt"

	"github.com/dula0/bookstore_users_api/databases/mysql/users_db"
	"github.com/dula0/bookstore_users_api/utils/errors"
	"github.com/dula0/bookstore_users_api/utils/mysql_utils"
)

// SQL Query
const (
	insertUserQuery = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?);"

	getUserQuery = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id=?;"

	updateUserQuery = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"

	deleteUserQuery = "DELETE FROM users WHERE id=?;"

	findUserByStatusQuery = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
)

// Retrieves user by their user ID
func (user *User) Get() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(getUserQuery)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer stmt.Close()

	// Holds the sql query result
	result := stmt.QueryRow(user.ID)

	// copy column values into struct fields.
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		return mysql_utils.ParseError(getErr)
	}

	return nil
}

func (user *User) Save() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(insertUserQuery)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)

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

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(updateUserQuery)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(deleteUserQuery)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.ID); err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(findUserByStatusQuery)
	if err != nil {
		return nil, errors.InternalServerError(err.Error())
	}
	defer stmt.Close()

	// execute the query
	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.InternalServerError(err.Error())
	}
	defer rows.Close()

	results := make([]User, 0)
	// iterate over the returned rows from our query
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, errors.NotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return results, nil
}
