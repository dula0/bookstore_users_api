package app

import (
	"github.com/gin-gonic/gin"
)

// Handles all requests & is responsible for creating a go routine for every request
var (
	router = gin.Default()
)

func StartApp() {
	urlMap()
	router.Run()
}
