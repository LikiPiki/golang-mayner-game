package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/telegram-bot-api.v4"
)

const (
	HOUR        = 2
	MONEY_VALUE = 10
)

var (
	// users
	insert_new_user            = "INSERT INTO users (name, mayner1, mayner2, mayner3, mayner4, score, money, time, user_id, active) VALUES (?, 1, 0, 0, 0, 0, 300, ?, ?, 0);"
	find_user                  = "SELECT id FROM users WHERE name=?"
	find_score_by_username     = "SELECT score, time FROM users WHERE name=?"
	update_user_score_and_time = "UPDATE users SET score=?, time=? WHERE name=?"
	get_all_video              = "SELECT mayner1, mayner2, mayner3, mayner4 FROM users WHERE name=?"
	get_new_money              = "SELECT money, score FROM users WHERE name=?"

	// value
	get_values = "SELECT id, name, cost FROM value"
)

func renderScore(username string) (money int64) {
	row := db.QueryRow(
		find_score_by_username,
		username,
	)

	var clock int64

	err := row.Scan(&money, &clock)
	if err != nil {
		log.Println(err)
	}

	var videocarts [4]int
	col := db.QueryRow(
		get_all_video,
		username,
	)

	err = col.Scan(
		&videocarts[0],
		&videocarts[1],
		&videocarts[2],
		&videocarts[3],
	)
	if err != nil {
		log.Println(err)
	}

	timeBefore := clock
	timeNow := time.Now().Unix()

	for i, el := range videos {
		money += (timeNow - timeBefore) * int64(videocarts[i]*el.Power/HOUR)
	}

	_, err = db.Exec(
		update_user_score_and_time,
		money,
		time.Now().Unix(),
		username,
	)
	return
}

func menu(msg *tgbotapi.Message) {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/score"),
			tgbotapi.NewKeyboardButton("/video"),
			tgbotapi.NewKeyboardButton("/shop"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/валюта"),
			tgbotapi.NewKeyboardButton("/stat"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/donate"),
			tgbotapi.NewKeyboardButton("/help"),
			tgbotapi.NewKeyboardButton("/sell"),
		),
	)
	reply := tgbotapi.NewMessage(msg.Chat.ID, "Menu")
	reply.ReplyMarkup = keyboard
	_, err := bot.Send(reply)
	if err != nil {
		log.Println(err)
	}

}

func video(msg *tgbotapi.Message) {
	var videos [4]int

	row := db.QueryRow(
		get_all_video,
		msg.From.UserName,
	)

	err := row.Scan(
		&videos[0],
		&videos[1],
		&videos[2],
		&videos[3],
	)

	if err != nil {
		log.Println(err)
	}

	fmt.Println(videos)

	for i, el := range videos {
		reply := tgbotapi.NewMessage(msg.Chat.ID, fmt.Sprintf("Количество видеокарт %d - %d", i+1, el))
		_, err := bot.Send(reply)
		if err != nil {
			log.Println(err)
		}
	}

}

func start(msg *tgbotapi.Message) {
	var reply tgbotapi.MessageConfig

	if len(msg.From.UserName) > 4 {
		var res sql.NullString
		row := db.QueryRow(find_user, msg.From.UserName)

		err := row.Scan(&res)
		if err != nil {
			log.Println(err)
		}
		if res.Valid {
			// this is temporary code here for ALPHA ONLY
			_, err := db.Exec("UPDATE users SET user_id=? WHERE name=?",
				msg.From.ID,
				msg.From.UserName,
			)
			if err != nil {
				log.Println(err)
			}
			// this is temporary code here
			reply = tgbotapi.NewMessage(msg.Chat.ID, "Ты уже зарегистрирован")
		} else {
			_, err := db.Exec(
				insert_new_user,
				msg.From.UserName,
				time.Now().Unix(),
				msg.From.ID,
			)
			if err != nil {
				log.Println(err)
			}
			reply = tgbotapi.NewMessage(msg.Chat.ID, "Ты регнулся! /help")
		}

	} else {
		reply = tgbotapi.NewMessage(msg.Chat.ID, "Не удается тебя зарегистрировать\nВ настройках установи  usernam\nИ попробуй /start снова")
	}

	_, err := bot.Send(reply)
	if err != nil {
		log.Println(err)
	}
}

func sell(msg *tgbotapi.Message) {
	row := db.QueryRow(get_new_money, msg.From.UserName)
	var money, score int64
	_ = renderScore(msg.From.UserName)
	err := row.Scan(&money, &score)
	if err != nil {
		log.Println(err)
	}

	reply := tgbotapi.NewMessage(msg.Chat.ID, fmt.Sprintf("Сейчас у тебя %dР\nОбменять можно 500b -> 1Р, от 500b\nБаланс: %db", money, score))

	reply.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Продать!", "yes"),
		),
	)
	_, err = bot.Send(reply)
	if err != nil {
		log.Println(err)
	}

}

func score(msg *tgbotapi.Message) {
	money := renderScore(msg.Chat.UserName)
	reply := tgbotapi.NewMessage(msg.Chat.ID, fmt.Sprintf("Твой баланс %d Bitcoins!", money))

	_, err := bot.Send(reply)
	if err != nil {
		log.Println(err)
	}
}

func donate(msg *tgbotapi.Message) {
	reply := tgbotapi.NewMessage(msg.Chat.ID, "Нужно больше Р, чтобы купить видеокарт?\nПиши @likipiki.\nИ за небольшое пожертвование получи бонусы!\nТак же туда можно присылать отзывы и предложения!")
	_, err := bot.Send(reply)
	if err != nil {
		log.Println(err)
	}
}

func shop(msg *tgbotapi.Message) {
	var reply tgbotapi.MessageConfig
	for i, el := range videos {
		reply = tgbotapi.NewMessage(msg.Chat.ID, fmt.Sprintf("%s\n%s", el.Name, el.Desk))

		reply.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("Купить %dР", el.Cost), fmt.Sprintf("video %d", i)),
			),
		)
		_, err := bot.Send(reply)
		if err != nil {
			log.Println(err)
		}
	}
}

func value(msg *tgbotapi.Message) {
	rows, err := db.Query(get_values)
	if err != nil {
		log.Println(err)
	}
	var name string
	var cost int
	var id int
	for rows.Next() {
		err := rows.Scan(&id, &name, &cost)
		if err != nil {
			log.Println(err)
		}

		reply := tgbotapi.NewMessage(
			msg.Chat.ID,
			fmt.Sprintf("%s:\nМайнится в минуту: %d", name, cost),
		)
		reply.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Выбрать!", fmt.Sprintf("value %d", id)),
			),
		)

		_, err = bot.Send(reply)
		if err != nil {
			log.Println(err)
		}
	}
}

func help(msg *tgbotapi.Message) {
	reply := tgbotapi.NewMessage(msg.Chat.ID, helpDesc)
	_, err := bot.Send(reply)
	if err != nil {
		log.Println(err)
	}
}
