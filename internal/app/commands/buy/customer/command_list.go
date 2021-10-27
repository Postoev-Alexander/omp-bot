package customer

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (c *CustomerCommander) List(inputMessage *tgbotapi.Message) {
	c.SendPage(inputMessage.Chat.ID, "Here all the customers: \n\n", 0, 5)
}

func (c *CustomerCommander) SendPage(chatID int64, header string, cursor uint64, limit uint64) {
	customers, err := c.customerService.List(cursor, limit+1)
	if err != nil {
		log.Printf("CustomerCommander.SendPage: error getting customer list - %v", err)
		return
	}

	msg := tgbotapi.NewMessage(chatID, "")

	if len(customers) == int(limit+1) {
		addNextButton(&msg, cursor, limit)
		customers = customers[:len(customers)-1]
	}

	outputMsgText := header
	for _, p := range customers {
		outputMsgText += p.String()
		outputMsgText += "\n"
	}

	msg.Text = outputMsgText

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("CustomerCommander.SendPage: error sending reply message to chat - %v", err)
	}
}

func addNextButton(msg *tgbotapi.MessageConfig, cursor uint64, limit uint64) {
	serializedData, _ := json.Marshal(CallbackListData{
		Cursor: cursor + limit,
		Limit:  limit,
	})

	callbackPath := path.CallbackPath{
		Domain:       "buy",
		Subdomain:    "customer",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		),
	)
}
