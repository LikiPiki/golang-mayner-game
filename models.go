package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"time"
)

type User struct {
	gorm.Model
	Name          string
	User_id       int64
	Mayner1       int
	Mayner2       int
	Mayner3       int
	Mayner4       int
	Score         int64
	Money         int
	Time          int
	CryptoValue   CryptoValue
	CryptoValueID int
}

func (u User) NewDefaultUser() User {
	u.Time = int(time.Now().Unix())
	u.Money = 300
	return u
}

type CryptoValue struct {
	gorm.Model
	Name string
	Cost int
}

func addDefaultValue() {
	bitcoin := CryptoValue{Name: "Bitcoin", Cost: 1}
	ethereum := CryptoValue{Name: "Ethereum", Cost: 2}
	bitcoincash := CryptoValue{Name: "Bitcoin Cash", Cost: 0}
	db.Create(&bitcoin)
	db.Create(&ethereum)
	db.Create(&bitcoincash)

}
