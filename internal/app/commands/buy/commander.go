package buy

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/buy/customer"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type BuyCommander struct {
	bot               *tgbotapi.BotAPI
	customerCommander Commander
}

func NewBuyCommander(bot *tgbotapi.BotAPI) *BuyCommander {
	return &BuyCommander{
		bot:               bot,
		customerCommander: customer.NewCustomerCommander(bot),
	}
}

func (c *BuyCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "customer":
		c.customerCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("BuyCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *BuyCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "customer":
		c.customerCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("BuyCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)

		msg := tgbotapi.NewMessage(
			msg.Chat.ID,
			fmt.Sprintf("Unknown subdomain - %s", commandPath.Subdomain))

		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("BuyCommander.HandleCommand: error sending reply message to chat - %v", err)
		}
	}
}