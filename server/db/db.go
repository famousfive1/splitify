package db

import (
	"database/sql"

	_ "github.com/glebarez/go-sqlite"
)

type SqliteDB struct {
	db *sql.DB
	Group GroupRepo
}

// TODO: Use context in all repos to manage timeout and such
func NewDB() (*SqliteDB, error) {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		return nil, err
	}

	gr, err := newGroupRepo(db)
	if err != nil {
		return nil, err
	}

	return &SqliteDB {
		db: db,
		Group: gr,
	}, nil
}
