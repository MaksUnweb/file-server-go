package main

import (
	"context"
	"file-server-go/handlers/http"
	"file-server-go/shared/types"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)


func main() {
	//Get filepath from store data:
	filePath, isExists := os.LookupEnv("FILE_PATH")
	if !isExists {
		log.Fatal("Error! No filepath in system environment!")
	}
	dbUrl, isExists := os.LookupEnv("DB_URL")

	if !isExists {
		log.Fatal("Error! No database url in system environment!")
	}

	router := gin.Default()
	router.MaxMultipartMemory = 256 << 20 // 256MB
	router.Use(types.ErrorHandler())
	pool,err := connect_db(dbUrl)
	if err != nil {
		log.Fatal("Database not work: ", err)
	}

	router.GET("/", http.HomeHandler)
	router.POST("/save-one", func(c *gin.Context) {
			http.SaveOneHandler(c, pool, filePath)
	})
	router.POST("save-several", func(c *gin.Context) {
		http.SaveSeveralHandler(c, pool, filePath)	
	})
	router.DELETE("/delete/:id", func(c *gin.Context) {
			http.DeleteHandler(c, pool)
	})

	router.GET("/download-file", func(c *gin.Context) {
		http.DownloadFile(c, pool)
	})
	
	router.Run(":8080")
}


func connect_db(db_url string) (*pgxpool.Pool, error){
	pool, err := pgxpool.New(context.Background(), db_url)	
	return pool, err
}


