package model

import (
	"mydiary_api/database"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserID            uint64          `json:"user_id"`
	User              User            `gorm:"foreignKey:UserID" json:"user"`
	TransactionTypeID uint64          `json:"transaction_type_id"`
	TransactionType   TransactionType `gorm:"foreignKey:TransactionTypeID" json:"transaction_type"`
	PaymentMethodID   uint64          `json:"payment_method_id"`
	PaymentMethod     PaymentMethod   `gorm:"foreignKey:PaymentMethodID" json:"payment_method"`
	ProductId         uint64          `json:"product_id"`
	Product           Product         `gorm:"foreignKey:ProductId"`
	Amount            uint64          `json:"amount"`
	TransactionCode   uint64          `json:"transaction_code"`
	Description       string          `json:"description"`
	Status            string          `json:"status"`
}

func (transaction *Transaction) Save() (*Transaction, error) {
	err := database.Database.Create(&transaction).Error

	if err != nil {
		return &Transaction{}, err
	}

	return transaction, nil
}

func FindAllTransaction() ([]Transaction, error) {
	var transaction []Transaction
	err := database.Database.Preload("User").Preload("TransactionType").Preload("PaymentMethod").Preload("Product").Find(&transaction).Error

	if err != nil {
		return []Transaction{}, err
	}

	return transaction, nil
}
