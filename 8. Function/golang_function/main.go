package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	s, _ := strconv.ParseFloat("7", 64)

	circle, square := multiReturnFunction(s)

	fmt.Println("Circle area ", circle)
	fmt.Println("Square area ", square)
}

func circleArea(r float64) float64 {
	phi := 3.17

	return phi * math.Pow(r, 2)
}

func squareArea(s float64) float64 {
	return math.Pow(s, 2)
}

func multiReturnFunction(s float64) (float64, float64) {
	return circleArea(s), squareArea(s)
}
