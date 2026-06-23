package http

import (
	"file-server-go/handlers/db"
	"file-server-go/handlers/mechanics"
	"file-server-go/shared/types"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)


func SaveOneHandler(c *gin.Context, pool *pgxpool.Pool, filePath string) {

	file, err := c.FormFile("file")
	if err != nil {
		c.Error(types.ErrBadRequest)
		return
	}

	fileType := mechanics.Filter(file.Filename)

	dst := filepath.Join(filePath, filepath.Base(file.Filename))
	err =	c.SaveUploadedFile(file, dst)
	if err != nil {
		log.Printf("Error save file: %s", err)
		c.Error(types.ErrInternalServer)
		return
	}

	err = db.InsertHandler(pool, dst, fileType)
	if err != nil {
		log.Printf("Error insert into database: %s", err)
		c.Error(types.ErrInternalServer)
		return
	}


	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}

