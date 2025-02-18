package course

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/education/course"
	"log"
	"strings"
)

func (c *CourseSubdomainCommander) New(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	data := strings.Split(args, " ")
	if len(data) != 2 {
		log.Printf("fail to get data from payload %v", args)
		c.bot.Send(tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("fail to get data from payload %v", args)))
		return
	}

	status, err := c.service.Create(course.Course{
		Title:       data[0],
		Description: data[1],
	})

	if err != nil {
		log.Printf("fail to create course: %v", err)
		return
	}

	botMsg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("Курс успешно создан, его идентификатор %d", status),
	)

	_, err = c.bot.Send(botMsg)
	if err != nil {
		log.Printf("CourseSubdomainCommander.New: error sending reply message to chat - %v", err)
	}
}
