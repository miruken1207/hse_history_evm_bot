package handler

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/telebot.v3"
)

func handleHistoryEVM(bot *telebot.Bot) func(c telebot.Context) error {

	return func(c telebot.Context) error {

		log.Printf("%s: /history_evm", c.Sender().FirstName)

		if err := godotenv.Load(); err != nil {
			log.Fatal("Ошибка загрузки .env файла: ", err)
		}

		message, err := os.ReadFile(os.Getenv("HISTORY_EVM_PATH"))
		if err != nil {
			log.Fatalf("Ошибка при чтении файла: %v", err)
		}

		return c.Send(string(message), telebot.ModeHTML)
	}

}
