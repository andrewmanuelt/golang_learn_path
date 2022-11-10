package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


// go get -u gorm.io/gorm
// go get -u gorm.io/driver/mysql

type Colors struct {
	gorm.Model

	Id    int    `gorm:"column:id"`
	Color string `gorm:"color"`
}

func main() {
	var color []Colors

	dsn := "root:@tcp(127.0.0.1:3306)/go_sample?parseTime=True"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&color)

	if err != nil {
		panic(err)
	}

	db.Find(&color)

	fmt.Println("========= Select Result ===========")

	for _, row := range color {
		fmt.Println(row.Id, ".", row.Color)
	}

	colorsData := Colors{
		Color: "Magenta",
	}

	db.Model(&color).Create(&colorsData)

	colorsData = Colors{
		Color: "Blue",
	}

	db.Model(&color).Create(&colorsData)

	colorsData = Colors{
		Color: "Orange",
	}

	db.Model(&color).Create(&colorsData)

	fmt.Println("========= Insert Result ===========")

	db.Find(&color)

	for _, row := range color {
		fmt.Println(row.Id, ".", row.Color)
	}

	updateData := Colors{
		Color: "Cyan",
	}

	db.Model(&color).Where("id", 2).Updates(&updateData)

	fmt.Println("========= Update Result ===========")

	db.Find(&color)

	for _, row := range color {
		fmt.Println(row.Id, ".", row.Color)
	}
}
