package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
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
	var user_money int
	row := db.QueryRow(get_money, msg.Chat.UserName)
	err = row.Scan(&user_money)

	if err != nil {
		message = "Ошибка"
	}

	user_money += money
	_, err = db.Exec(update_money, user_money, msg.Chat.UserName)
	if err != nil {
		message = "Ошибка"
	}
	reply = tgbotapi.NewMessage(msg.Chat.ID, message)
	_, err = bot.Send(reply)
	if err != nil {
		log.Println(err)
	}

}

func sendMessageForAll(message string) {
	rows, err := db.Query(select_users_id)
	if err != nil {
		log.Println(err)
	}
	var reply tgbotapi.MessageConfig
	var id sql.NullInt64
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			log.Println(err)
		}
		// this is test alpha code for users without ID, for old users
		if id.Valid {
			reply = tgbotapi.NewMessage(id.Int64, message)
			_, err := bot.Send(reply)
			if err != nil {
				log.Println(err)
			}
		}
	}

}
