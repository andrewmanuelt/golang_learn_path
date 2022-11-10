package main

import (
	"fmt"
)

func main() {
	// Normal text
	fmt.Println("Hello World")

	// One line multi text
	fmt.Println("My", "first", "golang", "code")

	// One line string with integer
	fmt.Println("Copyright", 2022)

	// Special character
	fmt.Print("\n")
	fmt.Println("Line break")
	fmt.Println("\tTabbed line")
	fmt.Println("No line break")

	// Hold terminal until key pressed
	var input string

	fmt.Scanln(&input)
}
