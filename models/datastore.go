package models

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Datastore interface {
	SubmitUser(*User)
}

type DB struct {
	*sql.DB
}

func NewDB(dataSource string) *DB {
	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		panic(err)
	}

	return &DB{db}
}
