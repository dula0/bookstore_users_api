package users

// data access object serves as the persistence layer so any interactions with the db occurs here

import (
	"fmt"
	"strings"

	"github.com/dula0/bookstore_users_api/databases/mysql/users_db"
	"github.com/dula0/bookstore_users_api/logger"
	"github.com/dula0/bookstore_users_api/utils/errors"
	"github.com/dula0/bookstore_users_api/utils/mysql_utils"
)

// SQL Query
const (
	insertUserQuery = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?);"

	getUserQuery = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id=?;"

	updateUserQuery = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"

	deleteUserQuery = "DELETE FROM users WHERE id=?;"

	findByStatusQuery = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"

	findByEmailAndPasswordQuery = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE email=? AND password=?;"
)

// Retrieves user by their user ID
func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(getUserQuery)
	if err != nil {
		logger.Error("error with get user method ", err)
		return errors.InternalServerError("database error")
	}
	defer stmt.Close()

	// Holds the sql query result
	result := stmt.QueryRow(user.ID)

	// copy column values into struct fields.
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		logger.Error("error when trying to get user by id", getErr)
		return errors.InternalServerError("database error")
	}

	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(insertUserQuery)
	if err != nil {
		logger.Error("error when trying to save user by", err)
		return errors.InternalServerError("database error")
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)

	if saveErr != nil {
		logger.Error("error when trying to save user", saveErr)
		return errors.InternalServerError("database error")
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error retrieving database generated id after creating the user", saveErr)
		return errors.InternalServerError("database error")
	}
	user.ID = userId
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(updateUserQuery)
	if err != nil {
		logger.Error("error with preparing user update method", err)
		return errors.InternalServerError("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		logger.Error("error with execution of user update method", err)
		return errors.InternalServerError("database error")
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(deleteUserQuery)
	if err != nil {
		logger.Error("error with preparation of user delete method", err)
		return errors.InternalServerError("database error")
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.ID); err != nil {
		logger.Error("error with  execution of user delete method", err)
		return errors.InternalServerError("database error")
	}

	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(findByStatusQuery)
	if err != nil {
		logger.Error("error with  preparation of user find by status method", err)
		return nil, errors.InternalServerError("database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error with  statement query of user find by status method", err)
		return nil, errors.InternalServerError("database error")
	}
	defer rows.Close()

	results := make([]User, 0)

	for rows.Next() { // iterate over the returned rows from our query
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("error while iterating over scanned user rows", err)
			return nil, errors.InternalServerError("database error")
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, errors.NotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return results, nil
}

func (user *User) FindByEmailAndPassword() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(findByEmailAndPasswordQuery)
	if err != nil {
		logger.Error("error with preparing to get user by email and password", err)
		return errors.InternalServerError("database error")
	}
	defer stmt.Close()

	// Holds the sql query result
	result := stmt.QueryRow(user.Email, user.Password)

	// copy column values into struct fields.
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		if strings.Contains(getErr.Error(), mysql_utils.NoRowsError) {
			return errors.NotFoundError("no user found with credentials")
		}
		logger.Error("error when trying to get user by email and password", getErr)
		return errors.InternalServerError("database error")
	}

	return nil
}
