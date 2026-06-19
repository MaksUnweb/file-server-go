package db

import (
	"context"
	"file-server-go/shared/interfaces"
)

func DeleteHandler[T interfaces.Queries](conn T, id int) error {
	_, err := conn.Exec(context.Background(), "DELETE FROM data WHERE id = $1", id)
	
	return err
}
