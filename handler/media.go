package handler

import (
	"log"

	"gopkg.in/telebot.v3"
)

func handleMedia(bot *telebot.Bot) func(c telebot.Context) error {

	return func(c telebot.Context) error {

		log.Printf("%s: /media", c.Sender().FirstName)

		inlineButton := [][]telebot.InlineButton{
			{
				telebot.InlineButton{
					Unique: "evm1",
					Text:   "ЭВМ \"Стрела\"",
				},

				telebot.InlineButton{
					Unique: "evm2",
					Text:   "БЭСМ-6",
				},
			},

			{
				telebot.InlineButton{
					Unique: "evm3",
					Text:   "ЭВМ \"Днепр\"",
				},

				telebot.InlineButton{
					Unique: "evm4",
					Text:   "ЭВМ \"Корвет\"",
				},
			},

			{
				telebot.InlineButton{
					Unique: "evm5",
					Text:   "МЭСМ",
				},

				telebot.InlineButton{
					Unique: "evm6",
					Text:   "ЭВМ \"Агат\"",
				},
			},

			{
				telebot.InlineButton{
					Unique: "exit",
					Text:   "◀️ Назад в Меню📲",
				},
			},
		}

		inlineMarkup := &telebot.ReplyMarkup{}
		inlineMarkup.InlineKeyboard = inlineButton

		message := `Медиа 🎥

		Выберите интересующий вас ЭВМ:
		`

		bot.Handle(&inlineButton[0][0], func(c telebot.Context) error {

			log.Printf("%s: ЭВМ \"Стрела\"", c.Sender().FirstName)

			return c.Send(&telebot.Photo{
				File:    telebot.FromDisk("data/evm1.jpg"),
				Caption: "ЭВМ \"Стрела\"",
			})
		})

		bot.Handle(&inlineButton[0][1], func(c telebot.Context) error {

			log.Printf("%s: БЭСМ-6", c.Sender().FirstName)

			return c.Send(&telebot.Photo{
				File:    telebot.FromDisk("data/evm2.jpg"),
				Caption: "БЭСМ-6",
			})
		})

		bot.Handle(&inlineButton[1][0], func(c telebot.Context) error {

			log.Printf("%s: ЭВМ \"Днепр\"", c.Sender().FirstName)

			return c.Send(&telebot.Photo{
				File:    telebot.FromDisk("data/evm3.jpg"),
				Caption: "ЭВМ \"Днепр\"",
			})
		})

		bot.Handle(&inlineButton[1][1], func(c telebot.Context) error {

			log.Printf("%s: ЭВМ \"Корвет\"", c.Sender().FirstName)

			return c.Send(&telebot.Photo{
				File:    telebot.FromDisk("data/evm4.jpg"),
				Caption: "ЭВМ \"Корвет\"",
			})
		})

		bot.Handle(&inlineButton[2][0], func(c telebot.Context) error {

			log.Printf("%s: МЭСМ", c.Sender().FirstName)

			return c.Send(&telebot.Photo{
				File:    telebot.FromDisk("data/evm5.jpg"),
				Caption: "МЭСМ",
			})
		})

		bot.Handle(&inlineButton[2][1], func(c telebot.Context) error {

			log.Printf("%s: ЭВМ \"Агат\"", c.Sender().FirstName)

			return c.Send(&telebot.Photo{
				File:    telebot.FromDisk("data/evm6.jpg"),
				Caption: "ЭВМ \"Агат\"",
			})
		})

		bot.Handle(&inlineButton[3][0], handleMenu(bot))

		return c.Send(message, inlineMarkup)
	}
}
