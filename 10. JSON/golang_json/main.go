package main

import (
	"encoding/json"
	"fmt"
)

type Fruit struct {
	FruitName string `json:"fruit"`
	Color     string `json:"color"`
	Taste     string `json:"taste"`
}

func main() {
	dataObject := map[string]string{
		"fruit": "Apple",
		"color": "Red",
		"taste": "Sweet",
	}

	dataTropicalFruit := []struct {
		FruitName string `json:"fruit"`
		Color     string `json:"color"`
		Taste     string `json:"taste"`
	}{
		{
			FruitName: "Pineapple",
			Color:     "Yellow to green",
			Taste:     "Sweet and little bit sour",
		},
		{
			FruitName: "Mango",
			Color:     "Green",
			Taste:     "Sweet",
		},
	}

	dataString := `{"fruit":"Apple","color":"Red","taste":"Sweet"}`

	jsonObject, _ := json.Marshal(dataObject)

	fmt.Println(string(jsonObject))

	var jsonStruct Fruit

	err := json.Unmarshal([]byte(dataString), &jsonStruct)

	if err != nil {
		panic(err)
	}

	fmt.Println(jsonStruct.FruitName)
	fmt.Println(jsonStruct.Color)
	fmt.Println(jsonStruct.Taste)

	jsonTropicalFruit, _ := json.Marshal(dataTropicalFruit)

	fmt.Println(string(jsonTropicalFruit))

	stringTropicalFruit := `[{"fruit":"Pineapple", "color":"Yellow to green", "taste":"Sweet and little bit sour"}, {"fruit":"Mango","color":"Green","taste":"Sweet"}]`

	var tropicalFruit []Fruit

	json.Unmarshal([]byte(stringTropicalFruit), &tropicalFruit)

	for loop := range tropicalFruit {
		fmt.Println(tropicalFruit[loop].FruitName)
		fmt.Println(tropicalFruit[loop].Color)
		fmt.Println(tropicalFruit[loop].Taste)
	}
}
