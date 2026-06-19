package main

import (
	"context"
	"file-server-go/handlers/http"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"file-server-go/shared/types"
)


func main() {

	router := gin.Default()
	router.MaxMultipartMemory = 256 << 20 // 256MB
	router.Use(types.ErrorHandler())
	db_url := "postgres://Admin:1111@127.0.0.1:5432/file_db"
	pool,err := connect_db(db_url)
	if err != nil {
		log.Fatal("Database not work: ", err)
	}

	router.GET("/", http.HomeHandler)
	router.POST("/save-one", func(c *gin.Context) {
			http.SaveOneHandler(c, pool)
	})
	router.POST("save-several", func(c *gin.Context) {
		http.SaveSeveralHandler(c, pool)	
	})
	router.DELETE("/delete/:id", func(c *gin.Context) {
			http.DeleteHandler(c, pool)
	})

	router.GET("/search-file", func(c *gin.Context) {
		http.SearchFile(c, pool)
	})
	
	router.Run(":8080")
}


func connect_db(db_url string) (*pgxpool.Pool, error){
	pool, err := pgxpool.New(context.Background(), db_url)	
	return pool, err
}


