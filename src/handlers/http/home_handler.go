package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func HomeHandler(c *gin.Context) {
	value := gin.H{
		"message": "Welcome to the file manager! You can use the following methods to work: File search - /search-file; file deletion - /delete; file upload - /save",
	}
	c.JSON(http.StatusOK, value)
}
