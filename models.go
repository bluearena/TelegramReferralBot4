package main

import (
	"time"
)

type Configuration struct {
	BotToken, BotUsername, DBName string
	ChatURL                       string
	Link20, Link50 				  string
}

type User struct {
	TelegramID      int `gorm:"primary_key"`
	Username, Token string
	RefCount        int
	CreatedAt       time.Time
}
