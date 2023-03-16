package db

import (
	"github.com/jmoiron/sqlx"

	// Import module for SQLite DB
	_ "modernc.org/sqlite"

	"github.com/aceberg/rediary/internal/check"
	"github.com/aceberg/rediary/internal/models"
)

func connect(path string) *sqlx.DB {
	dbx, err := sqlx.Connect("sqlite", path)
	check.IfError(err)

	return dbx
}

func exec(path string, sqlStatement string) {

	dbx := connect(path)

	_, err := dbx.Exec(sqlStatement)
	check.IfError(err)
}

// Select - select all from DB
func Select(path string) []models.Record {
	var recs []models.Record

	dbx := connect(path)

	err := dbx.Select(&recs, "SELECT * FROM records ORDER BY DATE DESC")
	check.IfError(err)

	return recs
}
