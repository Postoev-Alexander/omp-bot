package customer

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *CustomerCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	id, err := strconv.Atoi(args)
	if err != nil || id < 0 {
		log.Printf("CustomerCommander.Delete "+
			"error parsing id: %v", err)

		c.Reply(
			inputMessage.Chat.ID,
			"Failed to parse customer id! Correct syntax for 'delete' command is:\n"+
				`/delete__buy__customer <id> (id >= 0)`)
		return
	}

	_, err = c.customerService.Remove(uint64(id))
	if err != nil {
		log.Printf("Fail to get product with idx %d: %v", id, err)

		c.Reply(
			inputMessage.Chat.ID,
			fmt.Sprintf(`Failed to delete customer: %v`, err))
		return
	}

	c.Reply(
		inputMessage.Chat.ID,
		fmt.Sprintf(`Successfully deleted customer with id %d`, id))
}