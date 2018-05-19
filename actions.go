package main

import (
	"log"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func sendMessage(chatId int64, text string, keyboard interface{}) {
	msg := tgbotapi.NewMessage(chatId, text)
	//msg.ParseMode = tgbotapi.ModeMarkdown
	//msg.DisableWebPagePreview = true
	_, ok := keyboard.(tgbotapi.ReplyKeyboardMarkup)
	if ok {
		msg.ReplyMarkup = keyboard
	} else {
		_, ok = keyboard.(tgbotapi.InlineKeyboardMarkup)
		if ok {
			msg.ReplyMarkup = &keyboard
		} else {
			msg.ReplyMarkup = nil
		}
	}

	_, err := bot.Send(msg)
	if err != nil {
		log.Print(err)
	}
	log.Printf("[Bot] SENT %s TO %v", msg.Text, msg.ChatID)
}

func editMessage(chatId int64, messageId int, text string){
	msg := tgbotapi.NewEditMessageText(chatId, messageId, text)
	//msg.ParseMode = tgbotapi.ModeMarkdown
	msg.ReplyMarkup = &keyboard
	bot.Send(msg)
	log.Printf("[Bot] EDITED %s TO %v", msg.Text, msg.ChatID)
}