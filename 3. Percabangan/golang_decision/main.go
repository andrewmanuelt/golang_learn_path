package main

import "fmt"

func main() {
	// Define variable first
	var isComplete bool

	a := 10
	b := 20

	var car string = "Honda"

	// if
	if isComplete {
		fmt.Println("Task complete...")
	}

	// if - else
	if isComplete {
		fmt.Println("Task complete...")
	} else {
		fmt.Println("Task not complete yet...")
	}

	// if - else if - else
	if a < b {
		fmt.Println(a, "is less than", b)
	} else if a > b {
		fmt.Println(a, "is greater than", b)
	} else {
		fmt.Println(a, "equals", b)
	}

	// switch - case
	switch a % 2 {
	case 0:
		fmt.Println("Even number")
	case 1, 3, 5, 7:
		fmt.Println("Odd number")
	}

	// switch - case - fallthrough
	switch car {
	case "Toyota":
		fmt.Println("Sedan")
	case "Honda":
		fmt.Println("Motorcycle")
		fallthrough
	case "":
		fmt.Println("Car")
	}
}
