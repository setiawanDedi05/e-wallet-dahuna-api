package model

import (
	"mydiary_api/database"

	"gorm.io/gorm"
)

type Tip struct {
	gorm.Model
	Title     string `json:"title"`
	Url       string `json:"url"`
	Thumbnail string `json:"thumbnail"`
}

func (tip *Tip) Save() (*Tip, error) {
	err := database.Database.Create(&tip).Error

	if err != nil {
		return &Tip{}, err
	}

	return tip, nil
}

func FindAllTip() ([]Tip, error) {
	var tip []Tip
	err := database.Database.Find(&tip).Error

	if err != nil {
		return []Tip{}, err
	}

	return tip, nil
}
