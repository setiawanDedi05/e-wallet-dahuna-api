package model

import (
	"mydiary_api/database"

	"gorm.io/gorm"
)

type TransferHistory struct {
	gorm.Model
	SenderID        uint64      `json:"sender_id"`
	Sender          User        `gorm:"foreignKey:SenderID"`
	ReceiverID      uint64      `json:"receiver_id"`
	Receiver        User        `gorm:"foreignKey:ReceiverID"`
	TransactionCode uint64      `json:"transaction_code"`
	Transaction     Transaction `gorm:"foreignKey:TransactionCode"`
}

func (transferHistory *TransferHistory) Save() (*TransferHistory, error) {
	err := database.Database.Create(&transferHistory).Error

	if err != nil {
		return &TransferHistory{}, err
	}

	return transferHistory, nil
}

func FindAllTransferHistory() ([]TransferHistory, error) {
	var transferHistory []TransferHistory
	err := database.Database.Preload("Transaction").Find(&transferHistory).Error

	if err != nil {
		return []TransferHistory{}, err
	}

	return transferHistory, nil
}
