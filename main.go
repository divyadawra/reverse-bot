package main

import (
	"io/ioutil"
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

		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		if update.Message.Animation == nil {
			errMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hey "+update.Message.From.FirstName+", Send a GIF file")
			errMsg.ReplyToMessageID = update.Message.MessageID
			bot.Send(errMsg)
			continue
		}

		mp4filePath, _ := bot.GetFileDirectURL(update.Message.Animation.FileID)

		utils.Reverse(mp4filePath)

		gifBytes, err := ioutil.ReadFile("./files/reversed.gif")
		if err != nil {
			panic(err)
		}
		photoFileBytes := tgbotapi.FileBytes{
			Name:  "./files/reversed.gif",
			Bytes: gifBytes,
		}
		msg := tgbotapi.NewAnimationUpload(update.Message.Chat.ID, photoFileBytes)

		msg.ReplyToMessageID = update.Message.MessageID

		resp, err := bot.Send(msg)
		if err != nil {
			log.Println("error is -----", err)
		}
		log.Println("response is -----", resp)

	}
}
