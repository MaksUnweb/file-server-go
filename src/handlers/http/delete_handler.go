package http

import (
	"file-server-go/handlers/db"
	"file-server-go/shared/types"
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
		c.Error(types.ErrBadRequest)
		return
	}

	filePath, err := db.GetFilePath(pool, intId)
	if err != nil {
		c.Error(types.ErrInternalServer)
	}

	err = os.Remove(filePath)
	if err != nil {
		c.Error(types.ErrInternalServer)
		return
	}

	err = db.DeleteHandler(pool, intId)
	if err != nil {
		c.Error(types.ErrInternalServer)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success!"})
}
