package buy

import (
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
	bot                *tgbotapi.BotAPI
	customerCommander Commander
}

func NewBuyCommander(
	bot *tgbotapi.BotAPI,
) *BuyCommander {
	return &BuyCommander{
		bot: bot,
		// customerCommander
		customerCommander: customer.NewBuyCustomerCommander(bot),
	}
}

func (c *BuyCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Customer {
	case "customer":
		c.customerCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("BuyCommander.HandleCallback: unknown customer - %s", callbackPath.Customer)
	}
}

func (c *BuyCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Customer {
	case "customer":
		c.customerCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("BuyCommander.HandleCommand: unknown customer - %s", commandPath.Customer)
	}
}
