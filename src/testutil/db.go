package testutil

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
)

func ConnectionBd(t *testing.T) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), "postgres://Admin:1111@127.0.0.1:5432/file_db")

	if err != nil {
		t.Errorf("Error: %s", err)		
	}
	return conn 
}
