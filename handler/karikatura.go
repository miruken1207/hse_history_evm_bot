package handler

import (
	"log"

	"github.com/joho/godotenv"
	"gopkg.in/telebot.v3"
)

func handleKarikatura(bot *telebot.Bot) func(c telebot.Context) error {

	return func(c telebot.Context) error {

		log.Printf("%s: /karikatura", c.Sender().FirstName)

		if err := godotenv.Load(); err != nil {
			log.Fatal("Ошибка загрузки .env файла: ", err)
		}

		album := telebot.Album{
			&telebot.Photo{
				File:    telebot.FromDisk("data/photo_1.jpg"),
				Caption: "Карикатуры про ЭВМ📸",
			},
			&telebot.Photo{
				File: telebot.FromDisk("data/photo_2.jpg"),
			},
			&telebot.Photo{
				File: telebot.FromDisk("data/photo_3.jpg"),
			},
			&telebot.Photo{
				File: telebot.FromDisk("data/photo_4.jpg"),
			},
			&telebot.Photo{
				File: telebot.FromDisk("data/photo_5.jpg"),
			},
		}

		return c.SendAlbum(album)
	}

}
