package model

import (
	"mydiary_api/database"

	"gorm.io/gorm"
)

type OperatorCard struct {
	gorm.Model
	Name      string     `json:"name"`
	Status    string     `json:"status"`
	Thumbnail string     `json:"thumbnail"`
	DataPlans []DataPlan `json:"data_plans"`
}

func (operatorCard *OperatorCard) Save() (*OperatorCard, error) {
	err := database.Database.Create(&operatorCard).Error

	if err != nil {
		return &OperatorCard{}, err
	}

	return operatorCard, nil
}

func FindAllOperatorCard() ([]OperatorCard, error) {
	var operatorCards []OperatorCard
	err := database.Database.Preload("DataPlans").Find(&operatorCards).Error

	if err != nil {
		return []OperatorCard{}, err
	}

	return operatorCards, nil

}
