package handler

import (
	"fmt"
	"log"

	"gopkg.in/telebot.v3"
)

func handleStart(bot *telebot.Bot) func(c telebot.Context) error {

	return func(c telebot.Context) error {

		log.Printf("%s: /start", c.Sender().FirstName)

		message := `Привет %s! 👋

Я расскажу тебе, как в СССР создавали первые компьютеры. 🖥️

Чтобы продолжить откройте /menu 📲
`
		inlineStartButton := [][]telebot.InlineButton{
			{
				telebot.InlineButton{
					Unique: "menu",
					Text:   "Меню",
				},
			},
		}

		inlineMarkup := &telebot.ReplyMarkup{}
		inlineMarkup.InlineKeyboard = inlineStartButton

		bot.Handle(&inlineStartButton[0][0], handleMenu(bot))

		return c.Send(fmt.Sprintf(message, c.Sender().FirstName), inlineMarkup)
	}
}
