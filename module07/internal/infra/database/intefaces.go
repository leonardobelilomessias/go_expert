package database

import "goexpert.com/module07/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type ProductInterface interface {
	Create(product *entity.Product) error
	FindAll(page int, limit int, sort string) ([]entity.Product, error)
	FindByID(id string) (*entity.Product, error)
	Delete(id string) error
	Update(product *entity.Product) error 

	
}
