// Package sql performs SQL operations
package sql

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// InitDB initializes the database
func InitDB(dbLocation string) (*sql.DB, error) {
	// Open the MySQL DB at the provided location
	database, err := sql.Open("mysql", dbLocation)
	if err != nil {
		return nil, fmt.Errorf("Error opening MySQL DB File. Error: %v\n", err)
	}
	return database, nil
}

// CreateTable creates the table from the supplied values
func CreateTable(database *sql.DB, tableName string, values []string) error {
	// Structure the headers
	headers := structureHeaders(values)

	// Create the table
	createTableQuery := `CREATE TABLE IF NOT EXISTS ` + tableName + `(` + headers + `)`
	tableQuery, err := database.Prepare(createTableQuery)
	if err != nil {
		return fmt.Errorf("Improper SQL Query. Error: %v\n", err)
	}
	defer tableQuery.Close()
	tableQuery.Exec()
	return nil
}

// InsertDB inserts data into the database
func InsertDB(database *sql.DB, tableName string, headers []string, values [][]string) error {
	// Build insert query
	insertValueQuery := buildInsertQuery(tableName, headers)

	// Insert the data
	tx, err := database.Begin()
	if err != nil {
		return err
	}
	//fmt.Printf("Inserting data into database: %v table: %v\n", dbLocation, tableName)
	rows := 0
	for _, value := range values {
		// Convert slice to interface
		row := make([]interface{}, len(value))
		for id := range value {
			row[id] = value[id]
		}
		insertQuery, err := tx.Prepare(insertValueQuery)
		if err != nil {
			return fmt.Errorf("Error in inserting data into DB. Error: %v\n", err)
		}
		defer insertQuery.Close()
		insertQuery.Exec(row...)
		rows++
	}
	tx.Commit()
	fmt.Printf("%v rows have been inserted into the table: %v\n", rows, tableName)
	return nil
}

func structureHeaders(headers []string) string {
	headersString := strings.Join(headers[:], "` VARCHAR(255), `")
	headersString = "`" + headersString + "` VARCHAR(255)"
	headersString = strings.Replace(headersString, "`CVSS` VARCHAR(255)", "`CVSS` NUMERIC", 1)
	return headersString
}

func buildInsertQuery(tableName string, headers []string) string {
	headersString := strings.Join(headers[:], ", ")
	valueCount := len(headers)
	count := 1
	valueString := "?"
	for count < valueCount {
		valueString = valueString + ", ?"
		count++
	}
	insertValueQuery := fmt.Sprintf("INSERT INTO %v (%v) VALUES (%v)", tableName, headersString, valueString)
	return insertValueQuery
}
