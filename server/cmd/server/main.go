package main

import (
	"fmt"
	"os"
	"quizer-web/internal/cli"
	"strings"
)

func main() {
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
