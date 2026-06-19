package db

import (
	"context"
	"file-server-go/shared/interfaces"
)

//Func from 
func GetFilePath[T interfaces.Queries](pool T, id int) (string, error) {
	var _path string
	err := pool.QueryRow(context.Background(), "SELECT path FROM data WHERE id = $1", id).Scan(&_path)
	if err != nil { return "", err }
	return _path, nil
}
