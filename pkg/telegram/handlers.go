package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const commandStart = "start"

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Я не знаю такой команды")
	// msg := tgbotapi.NewMessage(message.Chat.ID, "Ты ввел команду /start")
	switch message.Command() {
	case commandStart:
		msg.Text = "Ты ввел команду /start"
		// msg.ReplyToMessageID = update.Message.MessageID
		_, err := b.bot.Send(msg)
		return err
	default:
		// msg.ReplyToMessageID = update.Message.MessageID
		_, err := b.bot.Send(msg)
		return err
	}
}

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	// msg.ReplyToMessageID = update.Message.MessageID
	b.bot.Send(msg)
}
