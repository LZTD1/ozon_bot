package course

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *CourseSubdomainCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "/help__education__course — print list of commands\n/get__education__course — get a entity\n/list__education__course — get a list of your entity (💎: with pagination via telegram keyboard)\n/delete__education__course — delete an existing entity\n/new__education__course — create a new entity // not implemented (💎: implement list fields via arguments)\n/edit__education__course — edit a entity      // not implemented")

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("CourseSubdomainCommander.Help: error sending reply message to chat - %v", err)
	}
}
