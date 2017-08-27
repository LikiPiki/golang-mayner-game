package main

import (
	"fmt"
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"strconv"
)

var (
	get_money       = "SELECT money FROM users WHERE name=?"
	select_users_id = "SELECT user_id FROM users"
)

func addMoney(cost string, name string, msg *tgbotapi.Message) {
	var reply tgbotapi.MessageConfig
	message := "Успешно"
	money, err := strconv.Atoi(cost)
	if err != nil {
		message = "Ошибка"
	}
	var us User
	db.First(&us, "name = ?", name)
	if us.ID != 0 {
		us.Money += money
	} else {
		message = "Ошибка"
	}
	db.Model(&us).Update(User{Money: us.Money})
	reply = tgbotapi.NewMessage(msg.Chat.ID, message)
	_, err = bot.Send(reply)
	if err != nil {
		log.Println(err)
	}

}

func sendMessageForAll(message string) {
	var users []User
	db.Find(&users)
	var reply tgbotapi.MessageConfig
	for _, i := range users {
		reply = tgbotapi.NewMessage(i.User_id, message)
		_, err := bot.Send(reply)
		if err != nil {
			log.Println(err)
		}
	}
}

func backupDb(msg *tgbotapi.Message) {
	reply := tgbotapi.NewDocumentUpload(int64(msg.From.ID), DB_IN_FOLDER_NAME)
	_, err := bot.Send(reply)
	if err != nil {
		fmt.Println("here")
		log.Println(err)
	}
}
