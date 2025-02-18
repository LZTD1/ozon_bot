package course

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"
	"strings"
)

type CallbackListData struct {
	Offset int `json:"offset"`
}

func (c *CourseSubdomainCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}

	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("CourseSubdomainCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}

	list, err := c.service.List(uint64(parsedData.Offset), 10)
	if err != nil {
		log.Printf("CourseSubdomainCommander.CallbackList: error get list messages to chat - %v", err)
	}
	text := strings.Builder{}
	text.WriteString("Here all the courses: \n\n")
	for i, course := range list {
		text.WriteString(fmt.Sprintf("%d. %s\n", i+1, course.Title))
	}

	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, text.String())

	serializedData, _ := json.Marshal(CallbackListData{
		Offset: parsedData.Offset + 10,
	})
	newCallbackPath := path.CallbackPath{
		Domain:       "education",
		Subdomain:    "course",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", newCallbackPath.String()),
		),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("CourseSubdomainCommander.List: error sending reply message to chat - %v", err)
	}
}
