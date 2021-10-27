package customer

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *CustomerCommander) Help(inputMessage *tgbotapi.Message) {
	c.Reply(
		inputMessage.Chat.ID,
		"Available commands:\n"+
			`		/help__buy__customer — print list of commands
		/get__buy__customer <id> (id >= 0) — get a customer
		/list__buy__customer   — get a list of customers
		/delete__buy__customer <id> (id >= 0) — delete an existing customer   
		/new__buy__customer { "name": "<string>", "age":<number>} — place a new customer
		/edit__buy__customer {"id": <number>,"name": "<string>", "age": <number>} — edit an existing customer`,
	)
}
