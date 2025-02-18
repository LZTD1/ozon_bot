package course

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/education/course"
	"log"
	"strconv"
	"strings"
)

func (c *CourseSubdomainCommander) Edit(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	data := strings.Split(args, " ")
	if len(data) != 3 {
		log.Printf("fail to get data from payload %v", args)
		c.bot.Send(tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("fail to get data from payload %v", args)))
		return
	}

	id, err := strconv.Atoi(data[0])
	if err != nil {
		log.Printf("fail to get id from payload %v", args)
		return
	}
	err = c.service.Update(uint64(id), course.Course{
		Title:       data[1],
		Description: data[2],
	})
	if err != nil {
		log.Printf("fail to create course: %v", err)
		return
	}

	botMsg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("Курс успешно обновлен!"),
	)

	_, err = c.bot.Send(botMsg)
	if err != nil {
		log.Printf("CourseSubdomainCommander.Edit: error sending reply message to chat - %v", err)
	}
}
