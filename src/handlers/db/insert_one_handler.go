package db

import (
	"context"
	"file-server-go/interfaces"
)



func InsertHandler[T interfaces.Queries](pool T, filepath, fileType string) error  {
	_, err := pool.Exec(context.Background(), "INSERT INTO data (type, path) VALUES ($1, $2)", 
	fileType, filepath)
	return err
}
