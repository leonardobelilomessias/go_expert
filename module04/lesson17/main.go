// trap of has many
package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct{
	ID int `gorm:primaryKey`
	Name string
	Price float64
	CategoryID int
	Category Category
	SerialNumber SerialNumber
	gorm.Model
}
type Category struct{
	ID int `gorm:primaryKey`
	Name string
	Products []Product
}

type SerialNumber struct{
	ID int `gorm:primaryKey`
	Number string
	ProductId int
}

func main(){
	dsn:="root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err!= nil{
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{},&SerialNumber{})
	
	// //create category
	// category := Category{
	// 	Name: "House",
	// }
	// db.Create(&category)

	// //Create  product
	// db.Create(&Product{
	// 	Name: "chair",
	// 	Price: 100.00,
	// 	CategoryID: 1,
	// })

	// db.Create(&SerialNumber{
	// 	Number: "12345",
	// 	ProductId: 1,

	// })
	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumber").Find(&categories).Error
	if err!= nil{
		panic(err)
	}
	for _, category:=range categories{
		fmt.Println(category.Name,":")
		for _, product:= range category.Products{
			println("- ",product.Name, category.Name, "serial number", product.SerialNumber.Number)
		}
	}
}
