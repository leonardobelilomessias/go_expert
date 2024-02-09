//  bad lock
package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

)

type Product struct{
	ID int `gorm:primaryKey`
	Name string
	Price float64
	//relation many to many
	Categories []Category `gorm:"many2many:products_categories;"`
	SerialNumber SerialNumber
	gorm.Model
}
type Category struct{
	ID int `gorm:primaryKey`
	Name string
	//relation many to many
	Products []Product `gorm:"many2many:products_categories;"`
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
	db.AutoMigrate(&Product{}, &Category{})
	tx := db.Begin()
	var c Category
	err = tx.Debug().Clauses(clause.Locking{Strength:"UPDATE"}).First(&c,1).Error
	if err!= nil{
		panic(err)
	}
	c.Name="Games"
	tx.Debug().Save(&c)
	tx.Commit()

}
