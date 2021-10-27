package customer

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/model/buy"
	"github.com/ozonmp/omp-bot/internal/service/buy/customer"
)

type CustomerService interface {
	Describe(customerID uint64) (*buy.Customer, error)
	List(cursor uint64, limit uint64) ([]buy.Customer, error)
	Create(buy.Customer) (uint64, error)
	Update(customerID uint64, customer buy.Customer) error
	Remove(customerID uint64) (bool, error)
}

type CustomerCommander struct {
	bot             *tgbotapi.BotAPI
	customerService CustomerService
}

func NewCustomerCommander(bot *tgbotapi.BotAPI) *CustomerCommander {
	customerService := customer.NewDummyCustomerService()

	return &CustomerCommander{
		bot:             bot,
		customerService: customerService,
	}
}

func (c *CustomerCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("CustomerCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *CustomerCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	case "new":
		c.New(msg)
	case "delete":
		c.Delete(msg)
	case "edit":
		c.Edit(msg)
	default:
		c.Default(msg)
	}
}

func (c *CustomerCommander) Reply(chatID int64, message string) {
	msg := tgbotapi.NewMessage(chatID, message)
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("CustomerCommander: error sending reply message to chat - %v", err)
	}
}
