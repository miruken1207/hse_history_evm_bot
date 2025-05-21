package handler

import (
	"log"

	"gopkg.in/telebot.v3"
)

// func handleText(c telebot.Context) error {

// 	userName := c.Sender().FirstName
// 	userMessage := c.Text()
// 	log.Printf("%s: %s", userName, userMessage)

// 	return c.Send("Чтобы получить больше информации, перейдите в /menu 📲")
// }

func handleText(bot *telebot.Bot) func(c telebot.Context) error {

	return func(c telebot.Context) error {
		
		userName := c.Sender().FirstName
		userMessage := c.Text()
		log.Printf("%s: %s", userName, userMessage)

		return c.Send("Чтобы получить больше информации, перейдите в /menu 📲")
	}
}
