package db

import (
	"fmt"

	"github.com/aceberg/rediary/internal/check"
	"github.com/aceberg/rediary/internal/models"
)

// Create - create table if not exists
func Create(path string) {

	sqlStatement := `CREATE TABLE IF NOT EXISTS records (
		"ID"		INTEGER PRIMARY KEY,
		"DATE"		TEXT,
		"TAG"		TEXT,
		"NAME"		TEXT,
		"COLOR"		TEXT,
		"NOTE"		TEXT,
		"MINUS"		TEXT,
		"PLUS"		TEXT,
		"TOTAL"		TEXT
	);`
	err := exec(path, sqlStatement)
	check.IfError(err)
}

// Insert - insert one rec into DB
func Insert(path string, rec models.Record) {

	sqlStatement := `INSERT INTO records (DATE, TAG, NAME, COLOR, NOTE, MINUS, PLUS, TOTAL) 
	VALUES ('%s','%s','%s','%s','%s','%d','%d','%d');`

	rec.Name = quoteStr(rec.Name)
	rec.Tag = quoteStr(rec.Tag)

	sqlStatement = fmt.Sprintf(sqlStatement, rec.Date, rec.Tag, rec.Name, rec.Color, rec.Note, rec.Minus, rec.Plus, rec.Total)

	err := exec(path, sqlStatement)
	check.IfError(err)
}

// Delete - delete one record
func Delete(path string, id int) {

	sqlStatement := `DELETE FROM records WHERE ID='%d';`

	sqlStatement = fmt.Sprintf(sqlStatement, id)

	err := exec(path, sqlStatement)
	check.IfError(err)
}

// Clear - delete all records from table
func Clear(path string) {
	sqlStatement := `DELETE FROM records;`
	err := exec(path, sqlStatement)
	check.IfError(err)
}

// Migrate - add column to table
func Migrate(path string) {

	sqlStatement := `SELECT 1 FROM records WHERE NOTE;`
	err := exec(path, sqlStatement)

	if err != nil {

		sqlStatement = `ALTER TABLE records ADD NOTE TEXT DEFAULT "";`
		err = exec(path, sqlStatement)
		check.IfError(err)
	}
}
