package app

import (
	hello "github.com/dula0/bookstore_users_api/controllers/hello_controller"
	users "github.com/dula0/bookstore_users_api/controllers/users_controller"
)

// Defines all the routes in our microservice, and maps it with the respective controller
func urlMap() {
	router.GET("/", hello.Controller_hello)
	router.GET("/users/:user_id", users.GetUser)

	router.POST("/users", users.CreateUser)
}
