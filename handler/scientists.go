package handler

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/telebot.v3"
)

func handleScientists(bot *telebot.Bot) func(c telebot.Context) error {

	return func(c telebot.Context) error {

		log.Printf("%s: /scientists", c.Sender().FirstName)

		inlineButton := [][]telebot.InlineButton{
			{
				telebot.InlineButton{
					Unique: "lebedev",
					Text:   "С. А. Лебедев",
				},

				telebot.InlineButton{
					Unique: "glushkov",
					Text:   "В. М.Глушков",
				},
			},

			{
				telebot.InlineButton{
					Unique: "bruk",
					Text:   "И. С. Брук",
				},

				telebot.InlineButton{
					Unique: "brusnetsov",
					Text:   "Н. П. Брусенцов",
				},
			},

			{
				telebot.InlineButton{
					Unique: "bazilevskiy",
					Text:   "Ю. Я. Базилевский",
				},
			},

			{
				telebot.InlineButton{
					Unique: "exit2",
					Text:   "◀️ Назад в Меню📲",
				},
			},
		}

		inlineMarkup := &telebot.ReplyMarkup{}
		inlineMarkup.InlineKeyboard = inlineButton

		message := `Личности 👤

		Выберите интересующую вас личность:
		`

		bot.Handle(&inlineButton[0][0], func(c telebot.Context) error {

			log.Printf("%s: С. А. Лебедев", c.Sender().FirstName)

			if err := c.Send(&telebot.Photo{File: telebot.FromDisk("data/lebedev.jpg")}); err != nil {
				return err
			}

			if err := godotenv.Load(); err != nil {
				log.Fatal("Ошибка загрузки .env файла: ", err)
			}
			message, err := os.ReadFile("data/lebedev.txt")
			if err != nil {
				log.Fatalf("Ошибка при чтении файла: %v", err)
			}

			inlineMedia := [][]telebot.InlineButton{
				{
					telebot.InlineButton{
						Unique: "people",
						Text:   "Личности 👤",
					},
				},
			}

			inlineMarkup := &telebot.ReplyMarkup{}
			inlineMarkup.InlineKeyboard = inlineMedia

			bot.Handle(&inlineMedia[0][0], handleScientists(bot))

			return c.Send(string(message), telebot.ModeHTML, inlineMarkup)
		})

		bot.Handle(&inlineButton[0][1], func(c telebot.Context) error {

			log.Printf("%s: В. М.Глушков", c.Sender().FirstName)

			if err := c.Send(&telebot.Photo{File: telebot.FromDisk("data/glushkov.jpg")}); err != nil {
				return err
			}

			if err := godotenv.Load(); err != nil {
				log.Fatal("Ошибка загрузки .env файла: ", err)
			}
			message, err := os.ReadFile("data/glushkov.txt")
			if err != nil {
				log.Fatalf("Ошибка при чтении файла: %v", err)
			}

			inlineMedia := [][]telebot.InlineButton{
				{
					telebot.InlineButton{
						Unique: "people",
						Text:   "Личности 👤",
					},
				},
			}

			inlineMarkup := &telebot.ReplyMarkup{}
			inlineMarkup.InlineKeyboard = inlineMedia

			bot.Handle(&inlineMedia[0][0], handleScientists(bot))

			return c.Send(string(message), telebot.ModeHTML, inlineMarkup)
		})

		bot.Handle(&inlineButton[1][0], func(c telebot.Context) error {

			log.Printf("%s: И. С. Брук", c.Sender().FirstName)

			if err := c.Send(&telebot.Photo{File: telebot.FromDisk("data/bruk.jpg")}); err != nil {
				return err
			}

			if err := godotenv.Load(); err != nil {
				log.Fatal("Ошибка загрузки .env файла: ", err)
			}
			message, err := os.ReadFile("data/bruk.txt")
			if err != nil {
				log.Fatalf("Ошибка при чтении файла: %v", err)
			}

			inlineMedia := [][]telebot.InlineButton{
				{
					telebot.InlineButton{
						Unique: "people",
						Text:   "Личности 👤",
					},
				},
			}

			inlineMarkup := &telebot.ReplyMarkup{}
			inlineMarkup.InlineKeyboard = inlineMedia

			bot.Handle(&inlineMedia[0][0], handleScientists(bot))

			return c.Send(string(message), telebot.ModeHTML, inlineMarkup)
		})

		bot.Handle(&inlineButton[1][1], func(c telebot.Context) error {

			log.Printf("%s: Н. П. Брусенцов", c.Sender().FirstName)

			if err := c.Send(&telebot.Photo{File: telebot.FromDisk("data/brusentsov.jpg")}); err != nil {
				return err
			}

			if err := godotenv.Load(); err != nil {
				log.Fatal("Ошибка загрузки .env файла: ", err)
			}
			message, err := os.ReadFile("data/brusentsov.txt")
			if err != nil {
				log.Fatalf("Ошибка при чтении файла: %v", err)
			}

			inlineMedia := [][]telebot.InlineButton{
				{
					telebot.InlineButton{
						Unique: "people",
						Text:   "Личности 👤",
					},
				},
			}

			inlineMarkup := &telebot.ReplyMarkup{}
			inlineMarkup.InlineKeyboard = inlineMedia

			bot.Handle(&inlineMedia[0][0], handleScientists(bot))

			return c.Send(string(message), telebot.ModeHTML, inlineMarkup)
		})

		bot.Handle(&inlineButton[2][0], func(c telebot.Context) error {

			log.Printf("%s: Ю. Я. Базилевский", c.Sender().FirstName)

			if err := c.Send(&telebot.Photo{File: telebot.FromDisk("data/basilevskiy.jpg")}); err != nil {
				return err
			}

			if err := godotenv.Load(); err != nil {
				log.Fatal("Ошибка загрузки .env файла: ", err)
			}
			message, err := os.ReadFile("data/basilevskiy.txt")
			if err != nil {
				log.Fatalf("Ошибка при чтении файла: %v", err)
			}

			inlineMedia := [][]telebot.InlineButton{
				{
					telebot.InlineButton{
						Unique: "people",
						Text:   "Личности 👤",
					},
				},
			}

			inlineMarkup := &telebot.ReplyMarkup{}
			inlineMarkup.InlineKeyboard = inlineMedia

			bot.Handle(&inlineMedia[0][0], handleScientists(bot))

			return c.Send(string(message), telebot.ModeHTML, inlineMarkup)
		})

		bot.Handle(&inlineButton[3][0], handleMenu(bot))

		return c.Send(message, inlineMarkup)
	}
}
