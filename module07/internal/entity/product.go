package entity

import (
	"errors"
	"time"

	"goexpert.com/module07/pkg/entity"
)
var (
	ErrIDIsRequired = errors.New("id is required")
	ErrInvalidID = errors.New("invalid id")
	ErrNameIsRequered = errors.New("name is requered")
	ErrPriceIsRequired = errors.New("price is requeride")
	ErrInvalidPrice = errors.New(" invalid price")
)

type Product struct{
	
	ID entity.ID `json:"id"`
	Name string `json:"name"`
	Price int `json:"price"`
	Created_at time.Time `json:"created_at"`
}

func NewProduct(name string, price int) (*Product,error){
	product := &Product{
		ID: entity.NewID(),
		Name: name,
		Price: price,
		Created_at: time.Now(),
	}
	err:= product.Validate()
	if err !=  nil{
		return nil, err
	}
	return product, nil
}

func(p *Product) Validate()error{
 if p.ID.String() ==""{
	return ErrIDIsRequired
 }
 if p.Name == ""{
	return ErrNameIsRequered
 }
 if p.Price == 0{
	return ErrPriceIsRequired
 }
 if p.Price < 0{
	return ErrInvalidPrice
 }
 if p.Price ==0{
	return ErrIDIsRequired
 }
 if _ ,err := entity.ParseID(p.ID.String()); err != nil{
	return ErrInvalidID
 }
 
 return nil
}