package course

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/education/course"
	"log"
)

type CourseTextCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)  // return error not implemented
	Edit(inputMsg *tgbotapi.Message) // return error not implemented
}

type CourseCallbackCommander interface {
	CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
}

type CourseSubdomainCommander struct {
	bot     *tgbotapi.BotAPI
	service course.Service
}

func NewCourseSubdomainCommander(bot *tgbotapi.BotAPI, service course.Service) *CourseSubdomainCommander {
	return &CourseSubdomainCommander{
		bot:     bot,
		service: service,
	}
}

func (c *CourseSubdomainCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		c.Help(callback.Message)
	}
	log.Printf("CourseSubdomainCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
}

func (c *CourseSubdomainCommander) HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(message)
	case "list":
		c.List(message)
	case "get":
		c.Get(message)
	case "delete":
		c.Delete(message)
	case "new":
		c.New(message)
	case "edit":
		c.Edit(message)
	default:
		c.Help(message)
	}
	log.Printf("CourseSubdomainCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)

}
