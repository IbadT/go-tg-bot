package main

import (
	"fmt"
	"log"
	"os"

	"github.com/IbadT/go-tg-bot.git/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла:", err)
	}

	token := os.Getenv("TG_TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)

	// bot, err := tgbotapi.NewBotAPI("7794320942:AAFnaMd2xs20MDGy0Sut99CH8PvZk3Luju4")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	telegramBot := telegram.NewBot(bot)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}
