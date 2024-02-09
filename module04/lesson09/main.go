//Setting and creating operations
package main

import(
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
	products := []Product{
	{Name:"Notebook", Price:100},
	{Name:"Mouse", Price:100},
	{Name:"Keyboard", Price:100},
	}
	db.Create(&products)
}