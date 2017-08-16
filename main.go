package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/telegram-bot-api.v4"
)

const (
	DB_NAME  = "sqlite3"
	DB_PATH  = "./data.db"
	helpDesc = "Покупай видеокарты, майни Биткоин получай бабки!\nНачать играть /menu\nСоздатель @likipiki\nПри поддержке чатика UwebDesign!\nС начала игры у тя есть немного бабла. Пойди в /shop и купи свою первую видяку!"
)

var (
	db     *sql.DB
	bot    *tgbotapi.BotAPI
	videos = make([]VideoCard, 0)
)

func init() {
	var err error
	db, err = sql.Open(DB_NAME, DB_PATH)
	if err != nil {
		err.Error()
	}

	videos = append(videos, VideoCard{"gtx 460", "Слабая видеокарта", 100, 1})
	videos = append(videos, VideoCard{"gtx 640", "Средняя видеокарта", 300, 4})
	videos = append(videos, VideoCard{"gtx 1080", "Мощная видеокарта", 1000, 7})
	videos = append(videos, VideoCard{"gtx 1080ti", "Топовая видеокарта", 3000, 30})

}

type VideoCard struct {
	Name  string
	Desk  string
	Cost  int
	Power int
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

		switch {
		case update.Message != nil:
			if update.Message.IsCommand() {
				command := update.Message.Command()

				switch command {
				case "menu":
					go menu(update.Message)
				case "start":
					go start(update.Message)
				case "video":
					go video(update.Message)
				case "score":
					go score(update.Message)
				case "sell":
					go sell(update.Message)
				case "shop":
					go shop(update.Message)
				case "help":
					go help(update.Message)
				case "donate":
					go donate(update.Message)
				default:
					fmt.Println("Not found command")
				}
			}
		case update.CallbackQuery != nil:
			if update.CallbackQuery.Data == "yes" {
				go sellAll(update.CallbackQuery)
			} else if strings.Split(update.CallbackQuery.Data, " ")[0] == "video" {
				go buy(
					update.CallbackQuery,
					strings.Split(update.CallbackQuery.Data, " ")[1],
				)
				fmt.Println("BUYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYY")
			}
		default:
			continue
		}

	}

}
