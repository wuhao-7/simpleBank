package db

import (
	"log"
	"os"
	"testing"
	"database/sql"
	_ "github.com/lib/pq"
)
const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:pX4f6My@localhost:5432/simple_bank?sslmode=disable"
)
var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver,dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}
