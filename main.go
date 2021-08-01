package main

import (
	"log"
	"os"
)

const botAPI = "https://api.telegram.org/bot"

var botToken string

func init() {
	botToken = os.Getenv("botToken")
}

func main() {
	offset := 0
	for {
		updates, err := getUpdates(offset)
		if err != nil {
			log.Println(err.Error())
		}
		for _, update := range updates {
			err := sendMessage(update)
			if err != nil {
				log.Println(err.Error())
			}
			offset = update.UpdateId + 1
		}
	}
}
