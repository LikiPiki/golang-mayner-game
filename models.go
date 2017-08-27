package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"time"
)

type User struct {
	gorm.Model    `json:"-"`
	Name          string      `json:"name"`
	User_id       int64       `json:"-"`
	Mayner1       int         `json:"video1"`
	Mayner2       int         `json:"video2"`
	Mayner3       int         `json:"video3"`
	Mayner4       int         `json:"video4"`
	Score         int64       `json:"score"`
	Money         int         `json:"money"`
	Time          int         `json:"-"`
	CryptoValue   CryptoValue `json:"-"`
	CryptoValueID int         `json:"-"`
}

func (u User) NewDefaultUser() User {
	u.Time = int(time.Now().Unix())
	u.Money = 300
	return u
}

type CryptoValue struct {
	gorm.Model `json:"-"`
	Name       string `json:"name"`
	Cost       int    `json:"cost"`
}

func addDefaultValue() {
	bitcoin := CryptoValue{Name: "Bitcoin", Cost: 1}
	ethereum := CryptoValue{Name: "Ethereum", Cost: 2}
	bitcoincash := CryptoValue{Name: "Bitcoin Cash", Cost: 0}
	db.Create(&bitcoin)
	db.Create(&ethereum)
	db.Create(&bitcoincash)

}
