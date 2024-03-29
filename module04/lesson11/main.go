// Making queries with where
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
}
func main(){
	dsn:="root:root@tcp(localhost:3306)/goexpert"

	db, err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err!= nil{
		panic(err)
	}
	db.AutoMigrate(&Product{})
	// db.Create(&Product{
	// 	Name: "Tablet",
	// 	Price: 1000.00,
	// })
	// create bat
	// products := []Product{
	// {Name:"Notebook", Price:100},
	// {Name:"Mouse", Price:100},
	// {Name:"Keyboard", Price:100},
	// }
	// db.Create(&products)

	//select
	// var product Product
	// db.First(&product,3)
	// fmt.Println(product)

	//select by name
	// db.First(&product, "name=?","Mouse")
	// fmt.Println(product)

	//select all
	// var products []Product
	// db.Find(&products)
	// for _, product := range products{
	// 	fmt.Println(product)
	// }
	//select with pagination
	// var products []Product
	// db.Limit(2).Offset(2).Find(&products)
	// for _, product := range products{
	// 	fmt.Println(product)
	// }

	//select with where
	// var products []Product
	// db.Where("price >?",100).Find(&products)
	// for _, product := range products{
	// 	fmt.Println(product)
	// }


	//with like
	var products []Product
	db.Where("name LIKE?","%book%").Find(&products)
	for _, product := range products{
		fmt.Println(product)
	}
}
