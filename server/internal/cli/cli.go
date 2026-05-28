package cli

import "fmt"

func Run() {
	var choice int = MainMenu()
	fmt.Println(choice)
}

func MainMenu() int {
	fmt.Println("Development terminal UI for quizer-web")
	fmt.Println("Select a mode:")
	choice := Menu("Editor", "Take a quiz", "Exit")

	switch choice {
	case 0:
		return 1
	case 1:
		return 2
	case 2:
		fmt.Println("Exiting...")
		return 0
	default:
		panic("Unexpected choice")
	}
}
