package main

import "fmt"

func main() {
	// For with condition, limit and itterator
	a := 4

	for b := 1; b < a; b++ {
		fmt.Println(b)
	}

	x := 4
	y := 7

	// For without condition, limit and itterator
	// U can find condition checker inside the block
	for {
		fmt.Println(x)
		x++

		if x == y {
			break
		}
	}

	// Itteration with range. Fit with array data
	var i = []rune{1, 2, 3, 4, 5, 6, 7}

	for _, j := range i {
		fmt.Println(j)
	}
}
