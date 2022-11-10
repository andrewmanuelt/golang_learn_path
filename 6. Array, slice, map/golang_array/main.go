package main

import "fmt"

func main() {
	// Single array without declaration
	var animals [3]string

	animals[0] = "giraffe"
	animals[1] = "lion"
	animals[2] = "ostritch"

	fmt.Println(animals)
	fmt.Println()

	// Single array with declaration
	var flowers = [3]string{"Lily", "Rose", "Jasmine"}

	fmt.Println(flowers)
	fmt.Println()

	// Multi array
	var numbers [3][3]int

	numbers[0][0] = 1
	numbers[0][1] = 3
	numbers[0][2] = 5
	numbers[1][0] = 2
	numbers[1][1] = 4
	numbers[1][2] = 6
	numbers[2][0] = 16
	numbers[2][1] = 32
	numbers[2][2] = 64

	fmt.Println(numbers)
	fmt.Println()

	// Slice
	newAnimals := animals[0:2]

	for _, animal := range newAnimals {
		fmt.Println(animal)
	}

	fmt.Println()

	newNumber := numbers[0:1][0:2]

	for _, number := range newNumber {
		fmt.Println(number)
	}

	fmt.Println()

	subject := map[string]string{}

	subject["monday"] = "math"
	subject["tuesday"] = "history"
	subject["wednesday"] = "biology"
	subject["thursday"] = "physics"
	subject["friday"] = "chemical"

	fmt.Println(subject)
	fmt.Println(subject["thursday"])
	fmt.Println()

	prime := map[int]int{}

	prime[0] = 2
	prime[1] = 3
	prime[2] = 5
	prime[3] = 7

	fmt.Println(prime)
	fmt.Println(prime[0])
}
