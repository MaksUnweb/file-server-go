package http

import (
	"file-server-go/shared/types"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"file-server-go/handlers/db"
)

//Func for return file. If file uknown or server is error, return JSON-message:
func DownloadFile(c *gin.Context, pool *pgxpool.Pool) {
	id := c.Query("fileId")	
	if id == "" {
		c.Error(types.ErrBadRequest)
		return
	}
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.Error(types.ErrBadRequest)
		return
	}

	filePath, err := db.GetFilePath(pool, intId)
	if err != nil {
		c.Error(types.ErrInternalServer)
		return
	}

	c.File(filePath)
}

