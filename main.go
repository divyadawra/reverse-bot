package main

import (
	"log"
	"reverseGifBot/utils"

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
		c := make(chan string, 1)
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		if update.Message.Animation == nil {
			continue
		}

		filePath, _ := bot.GetFileDirectURL(update.Message.Animation.FileID)

		// log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		c <- utils.Reverse(filePath)

		// msg := tgbotapi.NewAnimationUpload(update.Message.Chat.ID, gifFilePath)
		// msg := tgbotapi.NewMessage(update.Message.Chat.ID, "hello "+update.Message.From.FirstName)

		msg := tgbotapi.NewDocumentUpload(update.Message.Chat.ID, <-c)
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)

	}
}
