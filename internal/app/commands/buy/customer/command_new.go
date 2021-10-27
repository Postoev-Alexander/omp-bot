package customer

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/buy"
)

type NewCustomerData struct {
	UserId uint64 `json:"id"`
	Name   string `json:"name"`
	Age    uint32 `json:"age"`
}

func (c *CustomerCommander) New(inputMessage *tgbotapi.Message) {
	data := inputMessage.CommandArguments()

	parsedData := NewCustomerData{}
	err := json.Unmarshal([]byte(data), &parsedData)
	if err != nil {
		log.Printf("CustomerCommander.New: "+
			"error reading json data from "+
			"input string %v - %v", data, err)

		c.Reply(
			inputMessage.Chat.ID,
			"Failed to parse json! Correct syntax for 'new' command is:\n"+
				`/new__buy__customer { "name": "<string>", "age":<number>}`)
		return
	}

	newId, err := c.customerService.Create(
		buy.Customer{
			Name: parsedData.Name,
			Age:  parsedData.Age,
		})

	if err != nil {
		log.Printf("Fail to create customer: %v", err)
		return
	}

	c.Reply(
		inputMessage.Chat.ID,
		fmt.Sprintf("Customer with id %v created", newId),
	)
}
