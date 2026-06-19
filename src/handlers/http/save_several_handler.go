package http

import (
	"file-server-go/handlers/db"
	"file-server-go/handlers/mechanics"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)


func SaveSeveralHandler(c *gin.Context, pool *pgxpool.Pool) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"messsage": "Error! Please check request data!"})
		return
	}
	files := form.File["files"]

	for _, file := range files {
		dst := filepath.Join("./files/", filepath.Base(file.Filename))
		fileType := mechanics.Filter(file.Filename)
		err := c.SaveUploadedFile(file, dst)
		if err != nil {
			log.Printf("Error save file: %s", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error! Try again later!"})
			return
		}

		err = db.InsertHandler(pool, dst, fileType)
			if err != nil {
				log.Printf("Error insert into database: %s", err)
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Error! Try again later!"})
				return
			}

	}

	c.JSON(http.StatusOK, gin.H{"message": "Success!"})
}
