package db

import (
	"context"
	"file-server-go/testutil"
	"testing"

	"github.com/jackc/pgx/v5"
)

//Test from Insert One values:

func TestInsertOne(t *testing.T) {
	conn := testutil.ConnectionBd(t)
	testType := "Test"
	testPath := "Test/Path"
	tx, err := conn.Begin(t.Context())
	defer tx.Rollback(context.Background())
	defer conn.Close(t.Context())
	if err != nil {
		t.Errorf("Error make transaction: %s", err)
	}

	//Test Func: 
	err = InsertHandler(tx, testPath, testType)
	if err != nil {
		t.Errorf("Error insert data: %s", err)
	}

	checkInBd(t, tx, testType, testPath)

}


// Test for GetFilePath func:
func TestGetFilePath(t *testing.T) {	
	testType := "Test"
	testPath := "test/Test"
	testId := 969966996
	conn := testutil.ConnectionBd(t)
	tx, err := conn.Begin(t.Context())
	if err != nil {
		t.Errorf("Error: %s", err)
		conn.Close(t.Context())
	}
	defer tx.Rollback(t.Context())
	defer conn.Close(t.Context())

	//Inserting test value:
	err = insertTestValue(tx, testId,testType, testPath)	
	if err != nil {
		t.Errorf("Error: %s", err)		
		return
	}
	
	//Test func:
	get_path, err := GetFilePath(tx, testId)
	if err != nil {
		t.Errorf("Error: %s", err)		
		return
	}

	if get_path != testPath {
		t.Error("Error! String from DB does not match with string from test!")
	}
}

//Test delete function:
func TestDeleteHandler(t *testing.T) {
	testType := "Test"
	testPath := "test/Test"
	testId := 969966996
	conn := testutil.ConnectionBd(t)
	tx, err := conn.Begin(t.Context())
	defer tx.Rollback(context.Background())
	defer conn.Close(t.Context())
	if err != nil {
		t.Errorf("Error make transaction: %s", err)
	}
	
	//Insert test data:
	err = insertTestValue(tx, testId, testType, testPath)
	if err != nil{
		t.Errorf("Error insert test values: %s", err)
	}

	//Checking whether the test data has been inserted into the table:
	res := checkInBd(t, tx, testType, testPath) 
	if !res{
		t.Errorf("Error! No test data in table!")
	}

	//Start delete:
	err = DeleteHandler(tx, testId)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	res = checkInBd(t, tx, testType, testPath) 
	if res{
		t.Errorf("Error! No test data in table!")
	}
}


//Help functions:
func checkInBd(t *testing.T, conn pgx.Tx, testType, testPath string) bool {
	rows, err := conn.Query(t.Context(), "SELECT EXISTS (SELECT 1 FROM data WHERE type = $1 AND path = $2)", testType, testPath)
	if err != nil {
		t.Errorf("Error making request: %s", err)
	}

	var exists bool
	for rows.Next() {
		if err := rows.Scan(&exists); err != nil {
			t.Errorf("Error: %s", err)
		}
	}

	if err := rows.Err(); err != nil {
		t.Errorf("Error: %s", err)
	}
	return exists
}

func insertTestValue(conn pgx.Tx, id int, testType, testPath string) error {
	_, err := conn.Exec(context.Background(), "INSERT INTO data(id, type, path) VALUES ($1, $2, $3)", id, testType, testPath)

	return err
}


