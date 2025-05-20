package main

import (
	bot "history_project/bot"
	handler "history_project/handler"
	"log"
)

func main() {

	token := bot.GetBotToken()
	bot := bot.InitBot(token)

	handler.RegisterHandlers(bot)

	log.Printf("%s started working!\n\n", bot.Me.FirstName)
	bot.Start()
}
