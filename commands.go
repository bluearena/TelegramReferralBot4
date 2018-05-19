package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
	"strconv"
	"log"
)

func start(message *tgbotapi.Message) {
	fields := strings.Fields(message.Text)
	if len(fields) == 1 {
		var user User
		db.First(&user, "telegram_id = ?", message.From.ID)
		//err := db.Collection("users").FindOne(bson.M{"telegramid": message.From.ID}, &user)
		if user == (User{}) {
			user := User{TelegramID: message.From.ID,
				Username: message.From.FirstName,
				Token: generateToken(),
			}
			//err = db.Collection("users").Save(&user)
			db.Create(&user)
		}
		sendMessage(message.Chat.ID, phrases[1], keyboard)
	} else if len(fields) == 2 {
		var user User
		db.Find(&user, "token = ?", fields[1])
		//err := db.Collection("users").FindOne(bson.M{"token": fields[1]}, &user)
		if user == (User{}) {
			db.First(&user, "telegram_id = ?", message.From.ID)
			//err = db.Collection("users").FindOne(bson.M{"telegramid": message.From.ID}, &user)
			if user == (User{}) {
				user := User{TelegramID: message.From.ID,
					Username: message.From.FirstName,
					Token: generateToken(),
				}
				db.Create(&user)
				//err = db.Collection("users").Save(&user)
			}
			sendMessage(message.Chat.ID, phrases[1], keyboard)
		} else {
			user2 := User{}
			db.First(&user2, "telegram_id = ?", message.From.ID)
			//err := db.Collection("users").FindOne(bson.M{"telegramid": message.From.ID}, &user2)
			if user2 == (User{}) {

				user.RefCount++
				db.Save(&user)
				//err = db.Collection("users").Save(&user)

				user2 = User{TelegramID: message.From.ID,
					Username: message.From.FirstName,
					Token: generateToken(),
				}
				db.Create(&user2)
				//err = db.Collection("users").Save(&user2)
			}
			sendMessage(message.Chat.ID, phrases[1], keyboard)
		}

	}
}

func editCheck(query *tgbotapi.CallbackQuery) {
	log.Printf("[%s] %s", query.From.FirstName, "clicked Check")
	var user User
	db.First(&user, "telegram_id = ?", query.From.ID)
	//err := db.Collection("users").FindOne(bson.M{"telegramid": query.From.ID}, &user)

	text := phrases[8] + "t.me/" +
		config.BotUsername + "?start=" + user.Token + "\n\n" +
		phrases[9] + strconv.Itoa(user.RefCount)
	if user.RefCount >= 50 {
		text += "\n\nJoin privileged channel: " + config.Link50
	} else if user.RefCount >= 20 {
		text += "\n\nJoin privileged channel: " + config.Link20
	}
	editMessage(query.Message.Chat.ID, query.Message.MessageID, text)
}
