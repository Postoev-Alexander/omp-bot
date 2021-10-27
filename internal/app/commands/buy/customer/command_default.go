package customer

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *CustomerCommander) Default(inputMessage *tgbotapi.Message) {
	log.Printf("Unknown command from [%s]: %s", inputMessage.From.UserName, inputMessage.Text)
	c.Help(inputMessage)
}
