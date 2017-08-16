package main

import (
	"database/sql"
	"fmt"
	"strconv"

	"gopkg.in/telegram-bot-api.v4"
)

var (
	check_score            = "SELECT score, money FROM users WHERE name=?"
	update_money_and_score = "UPDATE users SET score=?, money=? WHERE name=?"
	select_money           = "SELECT money FROM users WHERE name=?"
	select_mayner1         = "SELECT mayner1 FROM users WHERE name=?"
	select_mayner2         = "SELECT mayner2 FROM users WHERE name=?"
	select_mayner3         = "SELECT mayner3 FROM users WHERE name=?"
	select_mayner4         = "SELECT mayner4 FROM users WHERE name=?"
	update_mayner1         = "UPDATE users SET mayner1=? WHERE name=?"
	update_mayner2         = "UPDATE users SET mayner2=? WHERE name=?"
	update_mayner3         = "UPDATE users SET mayner3=? WHERE name=?"
	update_mayner4         = "UPDATE users SET mayner4=? WHERE name=?"
	update_money           = "UPDATE users SET money=? WHERE name=?"
)

func buy(call *tgbotapi.CallbackQuery, number string) {
	num, err := strconv.Atoi(number)
	if err != nil {
		err.Error()
	}
	username := call.From.UserName
	var money int

	row := db.QueryRow(select_money, username)
	err = row.Scan(&money)
	if err != nil {
		err.Error()
	}

	var reply tgbotapi.MessageConfig
	var mayner int

	if money >= videos[num].Cost {
		money -= videos[num].Cost
		var col *sql.Row

		if num == 0 {
			col = db.QueryRow(select_mayner1, username)
		} else if num == 1 {
			col = db.QueryRow(select_mayner2, username)
		} else if num == 2 {
			col = db.QueryRow(select_mayner3, username)
		} else if num == 3 {
			col = db.QueryRow(select_mayner4, username)
		}

		err := col.Scan(&mayner)
		if err != nil {
			err.Error()
		}
		mayner++

		if num == 0 {
			_, err := db.Exec(update_mayner1, mayner, username)
			if err != nil {
				err.Error()
			}
		} else if num == 1 {
			_, err := db.Exec(update_mayner2, mayner, username)
			if err != nil {
				err.Error()
			}
		} else if num == 2 {
			_, err := db.Exec(update_mayner3, mayner, username)
			if err != nil {
				err.Error()
			}
		} else if num == 3 {
			_, err := db.Exec(update_mayner4, mayner, username)
			if err != nil {
				err.Error()
			}
		}

		_, err = db.Exec(update_money, money, username)
		if err != nil {
			err.Error()
		}

		reply = tgbotapi.NewMessage(int64(call.From.ID), fmt.Sprintf("Куплено %d", mayner))
	} else {
		reply = tgbotapi.NewMessage(int64(call.From.ID), "Недостаточно денег")
	}

	_, err = bot.Send(reply)
	if err != nil {
		err.Error()
	}
}

func sellAll(call *tgbotapi.CallbackQuery) {
	var message string
	message = "Успешно продано!\nBTC - %d\nР - %d"
	username := call.From.UserName

	row := db.QueryRow(check_score, username)

	var reply tgbotapi.MessageConfig
	var score, money int64

	err := row.Scan(&score, &money)
	if err != nil {
		err.Error()
	}

	if score >= 500 {
		money += score / 500
		score -= (score / 500) * 500
		_, err = db.Exec(
			update_money_and_score,
			score,
			money,
			username,
		)

		if err != nil {
			err.Error()
		}
		reply = tgbotapi.NewMessage(int64(call.From.ID), fmt.Sprintf(message, score, money))
	} else {
		reply = tgbotapi.NewMessage(int64(call.From.ID), "Слишком мало b чтобы продать!!!")
	}

	_, err = bot.Send(reply)
	if err != nil {
		err.Error()
	}
}
