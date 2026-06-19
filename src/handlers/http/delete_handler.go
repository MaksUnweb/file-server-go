package http

import (
	"file-server-go/handlers/db"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)


func DeleteHandler(c *gin.Context, pool *pgxpool.Pool) {
	id := c.Param("id")	

	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID!"})
		return
	}

	filePath, err := db.GetFilePath(pool, intId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error! Please later again!"})
	}

	err = os.Remove(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error! Please later again!"})
		return
	}


	err = db.DeleteHandler(pool, intId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error! Please later again!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success!"})
}
