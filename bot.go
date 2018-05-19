package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
	"io"
)

var (
	bot      *tgbotapi.BotAPI
	config   Configuration
	phrases  map[int]string
	db       *gorm.DB
	keyboard tgbotapi.InlineKeyboardMarkup
)

func main() {
	initLog()
	initConfig()
	initStrings()
	initDB()
	initKeyboard()

	var err error
	bot, err = tgbotapi.NewBotAPI(config.BotToken)
	if err != nil {
		log.Print("ERROR: ")
		log.Panic(err)
	}

	bot.Debug = false
	log.Printf("Authorized on account %s", bot.Self.UserName)


	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			if update.Message.IsCommand() {
				log.Print("It is command")
				command := update.Message.Command()
				switch command {
				case "start":
					log.Print("It is start")
					start(update.Message)
				}
			}
		} else if update.CallbackQuery != nil {
			log.Print("It is callback")
			switch update.CallbackQuery.Data {
			case "check":
				go editCheck(update.CallbackQuery)
			}
			bot.AnswerCallbackQuery(tgbotapi.CallbackConfig{update.CallbackQuery.ID, "", false, "", 0})
		}

	}
}

func initLog() {
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Print("ERROR: ")
		log.Panic(err)
	}
	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)
}

func initConfig() {
	readJson(&config, "config.json")
}

func initStrings() {
	readJson(&phrases, "strings.json")
}

func initDB() {
	var err error
	db, err = gorm.Open("sqlite3", config.DBName+".db")

	if err != nil {
		log.Print("********** ERROR: ")
		log.Panic(err)
	} else {
		log.Print("Opened DB")
	}
	db.LogMode(true)
	log.Print("Set LogMode")
	db.AutoMigrate(&User{})
	log.Print("Migrated")
}

func initKeyboard() {
	keyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL(phrases[12], config.ChatURL),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(phrases[14], "check"),
		),
	)
}
