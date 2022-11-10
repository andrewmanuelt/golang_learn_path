package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Declaration of variable
	var age int
	var name string
	var isGraduate bool
	var gpa float32
	var toefl rune

	/*
		Value will defined as its own default value

		int = 0
		string = ""
		bool = false
		float = 0
		rune = 0
	*/

	fmt.Println("I'm", age)
	fmt.Println("My name is:", name)
	fmt.Println("I am graduated student ?", isGraduate)
	fmt.Println("My GPA :", gpa)
	fmt.Println("My TOEFL score :", toefl)

	// Modify the variable value
	age = 20
	name = "Andy"
	isGraduate = false
	gpa = 3.2
	toefl = 421

	fmt.Println("I'm", age)
	fmt.Println("My name is:", name)
	fmt.Println("I am graduated student ?", isGraduate)
	fmt.Println("My GPA :", gpa)
	fmt.Println("My TOEFL score :", toefl)

	gpa = 3

	fmt.Println("My new GPA:", gpa)

	// Instant declaration of variable.
	// No need to defined variable type since go can detect and defined byself
	motto := "why u can do now if u can do tomorrow"
	mathScore := 76

	fmt.Println("My motivation is", motto)
	fmt.Println("My recent math score is", mathScore)

	// How to check variable type
	fmt.Println("Type of age is", reflect.TypeOf(age))
	fmt.Println("Type of name is", reflect.TypeOf(name))
	fmt.Println("Type of isGraduate is", reflect.TypeOf(isGraduate))
	fmt.Println("Type of gpa is", reflect.TypeOf(gpa))
	fmt.Println("Type of toefl is", reflect.TypeOf(toefl))
}
