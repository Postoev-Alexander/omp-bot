package customer

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/buy/customer"
)

type BuyCustomerCommander struct {
	bot              *tgbotapi.BotAPI
	customerService *customer.Service
}

func NewBuyCustomerCommander(
	bot *tgbotapi.BotAPI,
) *BuyCustomerCommander {
	customerService := customer.NewService()

	return &BuyCustomerCommander{
		bot:              bot,
		customerService: customerService,
	}
}

func (c *BuyCustomerCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("BuyCustomerCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *BuyCustomerCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	default:
		c.Default(msg)
	}
}
