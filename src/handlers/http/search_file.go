package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"file-server-go/handlers/db"
)

//Func for return file. If file uknown or server is error, return JSON-message:
func SearchFile(c *gin.Context, pool *pgxpool.Pool) {
	id := c.Query("fileId")	
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No argument id!"})
		return
	}
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect arg!"})
		return
	}

	filePath, err := db.GetFilePath(pool, intId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error! Please repeat later!"})
		return
	}

	c.File(filePath)
}

