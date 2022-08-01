package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Controller_hello(c *gin.Context) {
	c.String(http.StatusOK, "hello world!\n")
}
