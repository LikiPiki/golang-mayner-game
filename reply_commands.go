package main

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/telegram-bot-api.v4"
)

const (
	HOUR = 4
)

func getUserMoney(username string) (money int) {
	var us User
	db.First(&us, "name = ?", username)
	fmt.Println("Money is", us)
	return us.Money
}

func renderScore(username string) int64 {
	var us User
	var videocarts [4]int

	db.First(&us, "name = ?", username)
	fmt.Println("--------- printed --------")
	fmt.Println("-   ", us, "-")
	fmt.Println("--------- printed  --------")
	var val CryptoValue
	db.Model(&us).Related(&val)
	// testing val
	fmt.Println("--------- DEBUG  val -----")
	fmt.Println("-   ", val, "-")
	fmt.Println("--------- DEBUG  --------")
	// testing val END

	videocarts[0] = us.Mayner1
	videocarts[1] = us.Mayner2
	videocarts[2] = us.Mayner3
	videocarts[3] = us.Mayner4

	timeNow := int(time.Now().Unix())

	// testing acti

	fmt.Println("--------- DEBUG --------")
	fmt.Println("-   ", val, "-")
	fmt.Println("--------- DEBUG  --------")
	// testing us
	// testing us END

	for i, el := range videos {
		us.Score += int64((timeNow-us.Time)*(videocarts[i]*el.Power*val.Cost)) / 2
	}
	us.Time = timeNow
	db.Model(&us).Update(User{Time: us.Time, Score: us.Score})
	return us.Score
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

	var us User
	db.First(&us, "name = ?", msg.From.UserName)

	videos[0] = us.Mayner1
	videos[1] = us.Mayner2
	videos[2] = us.Mayner3
	videos[3] = us.Mayner4
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
	username := msg.From.UserName

	if len(username) > 4 {
		var us User
		db.Where("name = ?", username).First(&us)

		if us.ID != 0 {
			reply = tgbotapi.NewMessage(msg.Chat.ID, "Ты уже зарегистрирован")
		} else {
			us = us.NewDefaultUser()
			us.Name = username
			us.User_id = msg.Chat.ID
			db.Create(&us)
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
	_ = renderScore(msg.From.UserName)
	var us User
	db.First(&us, "name = ?", msg.From.UserName)

	reply := tgbotapi.NewMessage(msg.Chat.ID, fmt.Sprintf("Сейчас у тебя %dР\nОбменять можно 500b -> 1Р, от 500b\nБаланс: %db", us.Money, us.Score))

	reply.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Продать!", "yes"),
		),
	)
	_, err := bot.Send(reply)
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
	currentMoney := getUserMoney(msg.From.UserName)
	reply = tgbotapi.NewMessage(msg.Chat.ID, fmt.Sprintf("Баланс: %d Р", currentMoney))
	_, err := bot.Send(reply)
	if err != nil {
		log.Println(err)
	}
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
	var values []CryptoValue
	db.Find(&values)
	// testing values
	fmt.Println("--------- DEBUG --------")
	fmt.Println("-   ", values, "-")
	fmt.Println("--------- DEBUG  --------")
	for _, val := range values {
		reply := tgbotapi.NewMessage(
			msg.Chat.ID,
			fmt.Sprintf("%s:\nМайнится в минуту: %d", val.Name, val.Cost),
		)
		reply.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Выбрать!", fmt.Sprintf("value %d", val.ID)),
			),
		)

		_, err := bot.Send(reply)
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
