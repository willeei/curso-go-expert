package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:category_products"`
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:category_products"`
	gorm.Model
}

func main() {
	dsn := "root:9655@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

	//category := Category{Name: "Cozinha"}
	//db.Create(&category)
	//
	//category2 := Category{Name: "Eletronics"}
	//db.Create(&category2)
	//
	//db.Create(&Product{
	//	Name:       "Panela",
	//	Price:      100,
	//	Categories: []Category{category, category2},
	//})

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			fmt.Println(" - ", product.Name)
		}
	}
}
