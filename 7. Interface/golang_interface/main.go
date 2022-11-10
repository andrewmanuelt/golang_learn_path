package main

import "fmt"

type programmingLanguage interface {
	showName() string
	showCreator() string
}

type data struct {
	name    string
	creator string
}

func (d data) showName() string {
	return d.name
}

func (d data) showCreator() string {
	return d.creator
}

func showInfo(pl programmingLanguage) {
	fmt.Println(pl)
	fmt.Println(pl.showName())
	fmt.Println(pl.showCreator())
}

func main() {
	data := data{
		name:    "PHP",
		creator: "Rasmus Leedorf",
	}

	showInfo(data)
}
