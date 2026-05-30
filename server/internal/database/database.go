package database

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

func Insert(database *sql.DB, table string, columns []string, values []any) (sql.Result, error) { // TODO: Sanitization

	var now string = time.Now().UTC().Format("2006-01-02 15:04:05")
	columns = append(columns, "created_at", "updated_at")
	values = append(values, now, now)

	var valuesString string = ""

	for i, value := range values {
		switch fmt.Sprintf("%T", value) {
		case "string":
			valuesString += "'" + fmt.Sprintf("%v", value) + "'"
		case "int":
			valuesString += fmt.Sprintf("%v", value)
		default:
			valuesString += fmt.Sprintf("%v", value)
		}

		if i < len(values)-1 {
			valuesString += ", "
		}
	}

	var query string = "INSERT INTO " + table + " (" + strings.Join(columns, ", ") + ") VALUES (" + valuesString + ")"

	fmt.Println("Querying database: " + query)

	result, err := database.Exec(query)

	return result, err
}
