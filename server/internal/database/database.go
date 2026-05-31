package database

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

func insert(database *sql.DB, table string, columns []string, values []any) (int, error) {

	var now string = time.Now().UTC().Format("2006-01-02 15:04:05")
	columns = append(columns, "created_at", "updated_at")
	values = append(values, now, now)

	var err error = checkColumnsAndValues(columns, values)
	if err != nil {
		return 0, err
	}

	placeholders := make([]string, len(values))
	for i := range values {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}

	var query string = "INSERT INTO " + table + " (" + strings.Join(columns, ", ") + ") VALUES (" + strings.Join(placeholders, ", ") + ") RETURNING id"

	fmt.Println("Querying database: " + query)

	var row *sql.Row = database.QueryRow(query, values...)
	var id int
	err = row.Scan(&id)

	return id, err
}

func update(database *sql.DB, table string, id int, columns []string, values []any) (bool, error) {
	var now string = time.Now().UTC().Format("2006-01-02 15:04:05")
	columns = append(columns, "updated_at")
	values = append(values, now)

	var err error = checkColumnsAndValues(columns, values)
	if err != nil {
		return false, err
	}

	var setString string = ""
	for i := range values {
		setString += columns[i] + " = "

		setString += fmt.Sprintf("$%d", i+1)

		if i < len(values)-1 {
			setString += ", "
		}
	}

	var query string = "UPDATE " + table + " SET " + setString + " WHERE id = " + fmt.Sprintf("$%d", len(values)+1)

	fmt.Println("Querying database: " + query)

	_, err = database.Exec(query, append(values, id)...)

	var success bool = false
	if err == nil {
		success = true
	}

	return success, err
}

func checkColumnsAndValues(columns []string, values []any) error {
	if len(columns) != len(values) {
		return fmt.Errorf("Uneven columns and values.")
	} else {
		return nil
	}
}
