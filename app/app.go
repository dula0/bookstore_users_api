package app

import (
	"github.com/dula0/bookstore_users_api/logger"
	"github.com/gin-gonic/gin"
)

// Handles all requests & is responsible for creating a go routine for every request
var (
	router = gin.Default()
)

func StartApp() {
	urlMap()

	logger.Info("starting the application...")
	router.Run()
}
