package main

import "fmt"

type car struct {
	carType    string
	brand      string
	yearOfProd int
}

type garage struct {
	bicycle string
	car     car
}

func main() {
	// Single struct declaration
	carStruct := car{
		carType:    "Sedan",
		brand:      "Daihatsu",
		yearOfProd: 2020,
	}

	fmt.Println("Car type:", carStruct.carType)
	fmt.Println("Car brand:", carStruct.brand)
	fmt.Println("Car year production :", carStruct.yearOfProd)
	fmt.Println("")

	// Struct inside struct (car <- carStruct, see garage)
	garageStruct := garage{
		bicycle: "Polygon",
		car:     carStruct,
	}

	fmt.Println("Garage bicycle:", garageStruct.bicycle)
	fmt.Println("Car inside garage type:", garageStruct.car.carType)
	fmt.Println("Car inside garage brand:", garageStruct.car.brand)
	fmt.Println("Car inside garage year production :", garageStruct.car.yearOfProd)
	fmt.Println("")

	// Anonymous struct
	civic := struct {
		brand  string
		engine string
		hp     int
	}{
		brand:  "Toyota",
		engine: "DOHC-V2",
		hp:     2000,
	}

	fmt.Println("Car brand :", civic.brand)
	fmt.Println("Car engine :", civic.engine)
	fmt.Println("Car power :", civic.hp)
	fmt.Println("")

	// Multi value anonymous struct
	jazz := []struct {
		brand      string
		engine     string
		yearOfProd int
		hp         int
	}{
		{
			brand:      "Toyota",
			engine:     "VVTi",
			yearOfProd: 2012,
			hp:         2000,
		},
		{
			brand:      "Toyota",
			engine:     "VVTi-V2",
			yearOfProd: 2018,
			hp:         2000,
		},
	}

	fmt.Println("========================================")

	for _, car := range jazz {
		fmt.Println("Car brand :", car.brand)
		fmt.Println("Car engine :", car.engine)
		fmt.Println("Car year production :", car.yearOfProd)
		fmt.Println("Car power :", car.hp)
		fmt.Println("========================================")
	}
}
