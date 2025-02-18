package course

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"
	"strings"
)

func (c *CourseSubdomainCommander) List(inputMessage *tgbotapi.Message) {

	list, err := c.service.List(0, 10)
	if err != nil {
		log.Printf("CourseSubdomainCommander.List: error get list messages to chat - %v", err)
	}
	text := strings.Builder{}
	text.WriteString("Here all the courses: \n\n")
	for i, course := range list {
		text.WriteString(fmt.Sprintf("%d. %s\n", i+1, course.Title))
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, text.String())

	serializedData, _ := json.Marshal(CallbackListData{
		Offset: 10,
	})
	callbackPath := path.CallbackPath{
		Domain:       "education",
		Subdomain:    "course",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("CourseSubdomainCommander.List: error sending reply message to chat - %v", err)
	}
}
