package main

import (
	"database/sql"
	"fmt"
	"os"
	"quizer-web/internal/cli"
	"quizer-web/internal/database"
	"strings"
	"time"
)

func main() {
	fmt.Println("Starting server...")
	fmt.Println("Waiting 0.5 seconds for the database to be ready...")
	time.Sleep(500 * time.Millisecond)

	var _ *sql.DB = database.OpenConnection("Quizzes")
	var _ *sql.DB = database.OpenConnection("Users")
	var doCLI string = strings.ToLower(os.Getenv("DO_CLI"))

	switch doCLI {
	case "true", "":
		cli.Run() // Run the CLI interface
		return
	case "false":
		fmt.Println("CLI mode is off. Nothing will happen.") // TODO: Implement server functionality here
	default:
		panic("Invalid value for DO_CLI environment variable. Expected 'true' or 'false', got: " + doCLI)
	}

	fmt.Println("Server is exiting.")
}
