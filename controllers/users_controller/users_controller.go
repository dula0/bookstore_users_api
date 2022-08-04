// Main responsibility is to handle the request and do some minimal validation for the required parameters and then to send the validated request to the services package
package controllers

import (
	"net/http"
	"strconv"

	"github.com/dula0/bookstore_users_api/domain/users"
	"github.com/dula0/bookstore_users_api/services"
	"github.com/dula0/bookstore_users_api/utils/errors"
	"github.com/gin-gonic/gin"
)

func getUserID(IdParams string) (int64, *errors.RestErr) {
	userID, userErr := strconv.ParseInt(IdParams, 10, 64)
	if userErr != nil {
		return 0, errors.BadRequestError("User ID should be a number")
	}
	return userID, nil
}

func Create(c *gin.Context) {

	var user users.User
	// ShouldBindJSON will process the request body and deserialize the request json
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.BadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func Get(c *gin.Context) {
	userID, idErr := getUserID(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	user, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func Update(c *gin.Context) {
	userID, idErr := getUserID(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.BadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.ID = userID

	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func Delete(c *gin.Context) {
	userID, idErr := getUserID(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	if err := services.DeleteUser(userID); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

