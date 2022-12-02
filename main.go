package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Patrick564/qr-generator-bot/pkg/qr"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))
	if err != nil {
		log.Fatalf("dolor: %+v", err)
	}
	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates := bot.GetUpdatesChan(u)

	qrc, writer := qr.QRWriter()
	tgbotapi.NewInlineQueryResultPhoto("1231", "https://cdn2.thecatapi.com/images/MTcwNjExMw.jpg")

	for update := range updates {
		err = qr.New(update.Message.Text, writer)
		if err != nil {
			panic(err)
		}

		if update.Message == nil {
			continue
		}

		photoBytes := tgbotapi.FileBytes{
			Name:  "first",
			Bytes: qrc.Bytes(),
		}

		if _, err := bot.Send(tgbotapi.NewPhoto(update.Message.From.ID, photoBytes)); err != nil {
			fmt.Printf("Bytes of photo: %+v\n", photoBytes)
			panic(err)
		}
	}
}
