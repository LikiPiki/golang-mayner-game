package main

import (
	"fmt"
	"log"
	"strconv"

	"gopkg.in/telegram-bot-api.v4"
)

func buy(call *tgbotapi.CallbackQuery, number string) {
	num, err := strconv.Atoi(number)
	if err != nil {
		err.Error()
	}
	username := call.From.UserName
	var us User
	db.First(&us, "name = ?", username)
	fmt.Println(us.Money)

	var reply tgbotapi.MessageConfig

	if us.Money >= videos[num].Cost {
		us.Money -= videos[num].Cost
		fmt.Println(us.Money)
		if num == 0 {
			us.Mayner1++
		} else if num == 1 {
			us.Mayner2++
		} else if num == 2 {
			us.Mayner3++
		} else if num == 3 {
			us.Mayner4++
		}
		db.Save(&us)

		reply = tgbotapi.NewMessage(int64(call.From.ID), fmt.Sprintf("Куплено /video"))
	} else {
		reply = tgbotapi.NewMessage(int64(call.From.ID), "Недостаточно денег")
	}

	_, err = bot.Send(reply)
	if err != nil {
		err.Error()
	}
}

func changeValue(call *tgbotapi.CallbackQuery, number string) {
	id, err := strconv.Atoi(number)
	if err != nil {
		log.Println(err)
	}
	var val CryptoValue
	db.First(&val, id)
	fmt.Println(val)
	var us User
	db.First(&us, "name = ?", call.From.UserName)
	us.CryptoValue = val
	db.Save(&us)
	reply := tgbotapi.NewMessage(
		int64(call.From.ID),
		fmt.Sprintf("Майним %s", val.Name),
	)
	_, err = bot.Send(reply)
	if err != nil {
		log.Fatal(err)
	}
}

func sellAll(call *tgbotapi.CallbackQuery) {
	var message string
	message = "Успешно продано!\nBTC - %d\nР - %d"
	username := call.From.UserName

	var us User
	db.First(&us, "name = ?", username)

	var reply tgbotapi.MessageConfig

	if us.Score >= 500 {
		us.Money += int(us.Score / 500)
		us.Score -= (us.Score / 500) * 500
		db.Model(&us).Update(User{
			Score: us.Score,
			Money: us.Money,
		})
		reply = tgbotapi.NewMessage(int64(call.From.ID), fmt.Sprintf(message, us.Score, us.Money))
	} else {
		reply = tgbotapi.NewMessage(int64(call.From.ID), "Слишком мало b чтобы продать!!!")
	}

	_, err := bot.Send(reply)
	if err != nil {
		err.Error()
	}
}
