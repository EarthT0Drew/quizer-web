package cli

import "fmt"

func Menu(options ...string) int {
	for index := range options {
		fmt.Printf("%v. %v\n", index+1, options[index])
	}

	var choice int

	fmt.Scan(&choice)

	// if choice < 1 || choice > len(options) {
	// 	fmt.Println("Invalid choice, please try again.")
	// 	return Menu(options...)
	// }

	return choice - 1
}
