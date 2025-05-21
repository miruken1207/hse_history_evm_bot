package handler

import (
	"gopkg.in/telebot.v3"
)

func RegisterHandlers(bot *telebot.Bot) {

	bot.Handle(telebot.OnText, handleText(bot))
	bot.Handle("/start", handleStart(bot))
	bot.Handle("/menu", handleMenu(bot))
}
