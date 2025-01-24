package main

import (
	"fmt"
	"log"
	"os"

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

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)
	}
	fmt.Println("qwer")
}
