package customer

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/buy"
)

type EditCustomerData struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
	Age  uint32 `json:"age"`
}

func (c *CustomerCommander) Edit(inputMessage *tgbotapi.Message) {
	data := inputMessage.CommandArguments()

	parsedData := EditCustomerData{}
	err := json.Unmarshal([]byte(data), &parsedData)
	if err != nil {
		log.Printf("CustomerCommander.Edit: "+
			"error reading json data from "+
			"input string %v - %v", data, err)

		c.Reply(
			inputMessage.Chat.ID,
			"Failed to parse json! Correct syntax for 'edit' command is:\n"+
				`/edit__buy__customer {"id": <number>,"name": "<string>", "age": <number>}`)
		return
	}

	err = c.customerService.Update(
		parsedData.Id,
		buy.Customer{
			Name: parsedData.Name,
			Age:  parsedData.Age,
		})

	if err != nil {
		log.Printf("Fail to update customer: %v", err)
		c.Reply(
			inputMessage.Chat.ID,
			fmt.Sprintf("Fail to update customer: %v", err))
		return
	}

	c.Reply(
		inputMessage.Chat.ID,
		fmt.Sprintf("Customer with id %d updated successfully", parsedData.Id))
}
