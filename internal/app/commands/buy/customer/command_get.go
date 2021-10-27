package customer

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *CustomerCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	id, err := strconv.Atoi(args)
	if err != nil || id < 0 {
		log.Printf("CustomerCommander.Get "+
			"error parsing id: %v", err)

		c.Reply(
			inputMessage.Chat.ID,
			"Failed to parse customer id! Correct syntax for 'get' command is:\n"+
				`/get__buy__customer <id> (id >= 0)`)
		return
	}

	customer, err := c.customerService.Describe(uint64(id))
	if err != nil {
		log.Printf("fail to get product with idx %d: %v", id, err)

		c.Reply(
			inputMessage.Chat.ID,
			fmt.Sprintf(`Failed to get customer: %v`, err))
		return
	}

	c.Reply(
		inputMessage.Chat.ID,
		customer.String())
}
