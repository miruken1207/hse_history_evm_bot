package handler

import (
	"log"

	"gopkg.in/telebot.v3"
)

func handleMenu(bot *telebot.Bot) func(c telebot.Context) error {

	return func(c telebot.Context) error {

		log.Printf("%s: /menu", c.Sender().FirstName)

		inlineButton := [][]telebot.InlineButton{
			{
				telebot.InlineButton{
					Unique: "history_evm",
					Text:   "Вся история ЭВМ 📚",
					Data:   "open",
				},

				telebot.InlineButton{
					Unique: "scientists",
					Text:   "Личности 👤",
				},
			},

			{
				telebot.InlineButton{
					Unique: "karikatura",
					Text:   "Карикатуры 🖼️",
				},

				telebot.InlineButton{
					Unique: "media",
					Text:   "Медиа 🎥",
				},
			},
		}

		inlineMarkup := &telebot.ReplyMarkup{}
		inlineMarkup.InlineKeyboard = inlineButton

		message := `Меню 📲

		Выберите интересующую вас тему:
		`

		bot.Handle(&inlineButton[0][0], handleHistoryEVM(bot))
		bot.Handle(&inlineButton[0][1], handleScientists(bot))
		bot.Handle(&inlineButton[1][0], handleKarikatura(bot))
		bot.Handle(&inlineButton[1][1], handleMedia(bot))

		return c.Send(message, inlineMarkup)
	}
}
