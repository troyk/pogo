package pogo

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

// LogFunc can be set on the Db instance to allow query logging.
type LogFunc func(query string, args ...interface{})

type DB struct {
	*sql.DB
}

// Open opens a new database connection.
func Open(dataSourceName string) (*DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	pogoDb := &DB{}
	pogoDb.DB = db
	return pogoDb, nil
}

func (db *DB) Query(query string, args ...interface{}) (*Rows, error) {
	db.Log(query, args)
	rows, err := db.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return &Rows{rows}, nil
}

func (db *DB) Log(query string, args ...interface{}) {
	log.Println(query, args)
}
