package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:9655@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	// create
	db.Create(&Product{Name: "Laptop", Price: 1000})

	// create batch
	products := []Product{
		{Name: "Mouse", Price: 50},
		{Name: "Keyboard", Price: 150},
	}
	db.Create(&products)

}
