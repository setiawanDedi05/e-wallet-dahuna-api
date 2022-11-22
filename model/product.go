package model

import (
	"mydiary_api/database"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string `json:"name"`
	Thumbnail   string `json:"thumbnail"`
	Price       uint64 `json:"price"`
	Status      string `json:"status"`
	Description string `json:"description"`
}

func (product *Product) Save() (*Product, error) {
	err := database.Database.Create(&product).Error

	if err != nil {
		return &Product{}, err
	}

	return product, nil
}

func FindAllProduct() ([]Product, error) {
	var products []Product
	err := database.Database.Find(&products).Error

	if err != nil {
		return []Product{}, err
	}

	return products, nil
}
