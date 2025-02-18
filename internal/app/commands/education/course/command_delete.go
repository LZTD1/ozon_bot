package course

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *CourseSubdomainCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	status, err := c.service.Remove(uint64(idx))
	if err != nil {
		log.Printf("fail to get course with idx %d: %v", idx, err)
		return
	}
	msg := "Статус проведения операции: "
	if status {
		msg = msg + "успешно"
	} else {
		msg = msg + "не успешно"
	}

	botMsg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		msg,
	)

	_, err = c.bot.Send(botMsg)
	if err != nil {
		log.Printf("DemoSubdomainCommander.Get: error sending reply message to chat - %v", err)
	}
}
