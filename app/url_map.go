package app

import (
	users "github.com/dula0/bookstore_users_api/controllers/users_controller"
)

// Defines all the routes in our microservice, and maps it with the respective controller
func urlMap() {
	
	router.POST("/users", users.Create)
	router.POST("/users/login", users.Login)

	router.GET("/users/:user_id", users.Get)
	router.GET("/internal/users/search", users.Search)

	router.PATCH("/users/:user_id", users.Update)
	router.PUT("/users/:user_id", users.Update)
	
	router.DELETE("/users/:user_id", users.Delete)
}
