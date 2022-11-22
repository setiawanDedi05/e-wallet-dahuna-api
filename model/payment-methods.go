package model

import (
	"mydiary_api/database"

	"gorm.io/gorm"
)

type PaymentMethod struct {
	gorm.Model
	Name      string `json:"name"`
	Code      string `json:"code"`
	Thumbnail string `json:"thumbnail"`
	Status    string `json:"status"`
}

func (paymentMethod *PaymentMethod) Save() (*PaymentMethod, error) {
	err := database.Database.Create(&paymentMethod).Error

	if err != nil {
		return &PaymentMethod{}, err
	}

	return paymentMethod, nil
}

func FindAllPaymentMethod() ([]PaymentMethod, error) {
	var paymentMethod []PaymentMethod
	err := database.Database.Find(&paymentMethod).Error

	if err != nil {
		return []PaymentMethod{}, err
	}

	return paymentMethod, nil

}
