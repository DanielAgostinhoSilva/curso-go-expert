package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// create category
	//category := Category{Name: "Eletronicos"}
	//db.Create(&category)

	// create procut
	//product := Product{
	//	Name:       "Mouse",
	//	Price:      1000.00,
	//	CategoryID: category.ID,
	//}
	//db.Create(&product)

	// create serial number
	//db.Create(&SerialNumber{
	//	Number:    "12345",
	//	ProductID: product.ID,
	//})

	//var products []Product
	//db.Preload("Category").Preload("SerialNumber").Find(&products)
	//for _, product := range products {
	//	json, _ := json.MarshalIndent(product, "", "  ")
	//	fmt.Println(string(json))
	//}

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println("Categoria: ", category.Name)
		for _, product := range category.Products {
			fmt.Println("Produto: ", product.Name)
		}
		fmt.Println("")
	}
}
