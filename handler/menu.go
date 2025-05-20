package handler

import (
	"log"

	"gopkg.in/telebot.v3"
)

func handleMenu() func(c telebot.Context) error {

	return func(c telebot.Context) error {

		log.Printf("%s: %s", c.Sender().FirstName, c.Text())

		inlineButton := [][]telebot.InlineButton{
			{
				telebot.InlineButton{
					Unique: "btn_1",
					Text:   "Вся история ЭВМ 📚",
				},

				telebot.InlineButton{
					Unique: "btn_2",
					Text:   "Личности 👤",
				},
			},

			{
				telebot.InlineButton{
					Unique: "btn_3",
					Text:   "Карикатуры 🖼️",
				},

				telebot.InlineButton{
					Unique: "btn_4",
					Text:   "Медиа 🎥",
				},
			},
		}

		inlineMarkup := &telebot.ReplyMarkup{}
		inlineMarkup.InlineKeyboard = inlineButton

		message := `Меню 📲

		Выберите интересующую вас тему:
		`

		return c.Send(message, inlineMarkup)
	}
}
