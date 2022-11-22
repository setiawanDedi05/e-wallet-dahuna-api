package model

import (
	"math/rand"
	"mydiary_api/database"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	Balance     uint64 `json:"balance"`
	Pin         int    `json:"-"`
	UserID      uint   `json:"user_id"`
	User        User   `gorm:"foreignKey:UserID" json:"user"`
	CardNnumber string `json:"card_number" gorm:"unique"`
}

func (wallet *Wallet) Save() (*Wallet, error) {
	wallet.CardNnumber = generateCardNumber()
	err := database.Database.Create(&wallet).Error

	if err != nil {
		return &Wallet{}, err
	}
	return wallet, nil
}

func FindAllWallet() ([]Wallet, error) {
	var wallets []Wallet
	err := database.Database.Preload("User").Find(&wallets).Error

	if err != nil {
		return []Wallet{}, err
	}

	return wallets, nil
}

func generateCardNumber() string {
	t := time.Now()
	secondString := strconv.Itoa(t.Second())
	minutesString := strconv.Itoa(t.Minute())
	hourString := strconv.Itoa(t.Hour())
	yearString := strconv.Itoa(t.Year())
	monthString := strconv.Itoa(int(t.Month()))
	dayString := strconv.Itoa(t.Day())
	randomNumberString := strconv.Itoa(rand.Intn(999-100) + 100)

	return randomNumberString + yearString + monthString + dayString + hourString + minutesString + secondString
}
