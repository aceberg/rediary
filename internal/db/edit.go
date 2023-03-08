package db

import (
	"fmt"

	"github.com/aceberg/rediary/internal/models"
)

// Create - create table if not exists
func Create(path string) {

	sqlStatement := `CREATE TABLE IF NOT EXISTS records (
		"ID"		INTEGER PRIMARY KEY,
		"DATE"		TEXT,
		"NAME"		TEXT,
		"MINUS"		TEXT,
		"PLUS"		TEXT,
		"TOTAL"		TEXT
	);`
	exec(path, sqlStatement)
}

// Insert - insert one rec into DB
func Insert(path string, rec models.Record) {

	sqlStatement := `INSERT INTO records (DATE, NAME, MINUS, PLUS, TOTAL) 
	VALUES ('%s','%s','%d','%d','%d');`

	rec.Name = quoteStr(rec.Name)

	sqlStatement = fmt.Sprintf(sqlStatement, rec.Date, rec.Name, rec.Minus, rec.Plus, rec.Total)

	exec(path, sqlStatement)
}

// Delete - delete one record
func Delete(path string, id int) {

	sqlStatement := `DELETE FROM records WHERE ID='%d';`

	sqlStatement = fmt.Sprintf(sqlStatement, id)

	exec(path, sqlStatement)
}

// Clear - delete all records from table
func Clear(path string) {
	sqlStatement := `DELETE FROM records;`
	exec(path, sqlStatement)
}
