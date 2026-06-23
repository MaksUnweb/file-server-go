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


func SaveSeveralHandler(c *gin.Context, pool *pgxpool.Pool, filePath string) {
	form, err := c.MultipartForm()
	if err != nil {
		c.Error(types.ErrBadRequest)
		return
	}
	files := form.File["files"]

	for _, file := range files {
		dst := filepath.Join(filePath, filepath.Base(file.Filename))
		fileType := mechanics.Filter(file.Filename)
		err := c.SaveUploadedFile(file, dst)
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

	}

	c.JSON(http.StatusOK, gin.H{"message": "Success!"})
}
