package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dsn := "root:9655@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	// create
	//db.Create(&Product{Name: "Laptop", Price: 1000})
	//
	// create batch
	//products := []Product{
	//	{Name: "Mouse", Price: 50},
	//	{Name: "Keyboard", Price: 150},
	//}
	//db.Create(&products)

	// read
	var product Product
	db.First(&product, 17)
	fmt.Println(product)

	//db.First(&product, "name = ?", "Laptop")
	//fmt.Println(product)

	// read all
	//var products []Product
	//db.Find(&products)
	//for _, product := range products {
	//	fmt.Println(product)
	//}

	//var products []Product
	//db.Limit(2).Offset(2).Find(&products)
	//for _, product := range products {
	//	fmt.Println(product)
	//}

	// where
	//var products []Product
	//db.Where("price < ?", 100).Find(&products)
	//for _, product := range products {
	//	fmt.Println(product)
	//}

	//var products []Product
	//db.Where("name LIKE ?", "%key%").Find(&products)
	//for _, product := range products {
	//	fmt.Println(product)
	//}

	// update
	//var product Product
	//db.First(&product, 17)
	//product.Name = "New Keyboard"
	//db.Save(&product)
	//
	//var newKeyboard Product
	//db.First(&newKeyboard, 17)
	//fmt.Println(newKeyboard)

	// delete
	//db.Delete(&newKeyboard)
}
