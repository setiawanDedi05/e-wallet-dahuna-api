package model

import (
	"mydiary_api/database"

	"gorm.io/gorm"
)

type DataPlan struct {
	gorm.Model
	Name           string        `json:"name"`
	Price          uint64        `json:"price"`
	OperatorCardId uint64        `gorm:"index" json:"operator_card_id"`
	OperatorCard   *OperatorCard `gorm:"foreignKey:OperatorCardId" json:"operator_card"`
}

func (dataplan *DataPlan) Save() (*DataPlan, error) {
	err := database.Database.Create(&dataplan).Error

	if err != nil {
		return &DataPlan{}, err
	}

	return dataplan, nil
}

func FindAllDataPlan() ([]DataPlan, error) {
	var dataPlan []DataPlan
	err := database.Database.Preload("OperatorCard").Find(&dataPlan).Error

	if err != nil {
		return []DataPlan{}, err
	}

	return dataPlan, nil
}
