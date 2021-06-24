package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("1802468177:AAET5EL6gU6fTKAqh05MIravDbdPcDFeFWY")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		mimeType := update.Message.Animation.MimeType
		// fileName := update.Message.Animation.FileName
		fmt.Println(mimeType)
		// switch mimeType {
		// case "video/mp4":
		// 	file := utils.Mp4File{fileName}
		// 	file.Reverse()
		// case "gif":
		// 	file := &GifFile{fileName}
		// 	file.Reverse()
		// }
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "hello "+update.Message.From.FirstName)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
