package model

import (
	"mydiary_api/database"

	"gorm.io/gorm"
)

type TransactionType struct {
	gorm.Model
	Name      string `json:"name"`
	Code      string `json:"code"`
	Action    string `json:"action"`
	Thumbnail string `json:"thumbnail"`
}

func (transactionType *TransactionType) Save() (*TransactionType, error) {
	err := database.Database.Create(&transactionType).Error

	if err != nil {
		return &TransactionType{}, err
	}

	return transactionType, nil
}

func FindAllTransactionType() ([]TransactionType, error) {
	var transactionType []TransactionType
	err := database.Database.Find(&transactionType).Error

	if err != nil {
		return []TransactionType{}, err
	}

	return transactionType, nil
}
