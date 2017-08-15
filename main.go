package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/telegram-bot-api.v4"
)

const (
	DB_NAME = "sqlite3"
	DB_PATH = "./data.db"
)

var (
	db  *sql.DB
	bot *tgbotapi.BotAPI
)

func init() {
	var err error
	db, err = sql.Open(DB_NAME, DB_PATH)
	if err != nil {
		err.Error()
	}
}

func main() {

	var err error
	bot, err = tgbotapi.NewBotAPI(TOKEN)

	if err != nil {
		err.Error()
	}

	bot.Debug = true

	log.Printf("Authoriezed on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			command := update.Message.Command()

			switch command {
			case "menu":
				menu(update.Message)
			case "start":
				start(update.Message)
			case "video":
				video(update.Message)
			case "score":
				score(update.Message)
			default:
				fmt.Println("Not found command")
			}

		}

	}

}
