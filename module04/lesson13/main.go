// change and remove records
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
	gorm.Model
}
func main(){
	dsn:="root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err!= nil{
		panic(err)
	}
	db.AutoMigrate(&Product{})
	
	// db.Create(&Product{
	// 	Name: "Tablet",
	// 	Price: 1000.00,
	// })

	// var p Product
	// db.First(&p,1)
	// p.Name = "New Mouse"
	// db.Save(&p)
	// fmt.Println(p)

	var p2 Product
	db.First(&p2,1)
	fmt.Println(p2.Name)
	db.Delete(&p2)

}
