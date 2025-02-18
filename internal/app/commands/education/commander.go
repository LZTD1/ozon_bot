package education

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	course2 "github.com/ozonmp/omp-bot/internal/app/commands/education/course"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/education/course"
	"log"
)

//
//type Commander interface {
//	Help(inputMsg *tgbotapi.Message)
//	Get(inputMsg *tgbotapi.Message)
//	List(inputMsg *tgbotapi.Message)
//	Delete(inputMsg *tgbotapi.Message)
//
//	New(inputMsg *tgbotapi.Message)
//	Edit(inputMsg *tgbotapi.Message)
//}

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type CourseCommander struct {
	service            course.Service
	subdomainCommander Commander
	bot                *tgbotapi.BotAPI
}

func NewEduCommander(bot *tgbotapi.BotAPI, service course.Service) *CourseCommander {
	return &CourseCommander{
		bot:                bot,
		service:            service,
		subdomainCommander: course2.NewCourseSubdomainCommander(bot, service),
	}
}

func (c *CourseCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "course":
		c.subdomainCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("CourseCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *CourseCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "course":
		c.subdomainCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("CourseCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
