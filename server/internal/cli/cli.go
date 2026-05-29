package cli

import "fmt"

func Run() {
	for {
		var choice int = MainMenu()

		switch choice {
		case 0:
			EditorMode()
		case 1:
			// Handle Take a quiz mode
		case 2:
			fmt.Println("Exiting...")
			return
		}
	}
}

func MainMenu() int {
	fmt.Println("Development terminal UI for quizer-web")
	fmt.Println("Select a mode:")
	choice := Menu("Editor", "Take a quiz", "Exit")

	if choice < 0 || choice > 2 {
		panic("Invalid return code: Error in Menu function. Expected 0, 1, or 2, got " + fmt.Sprint(choice) + ".")
	}

	return choice
}
