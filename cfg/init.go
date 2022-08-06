package cfg

import (
	botapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func InitBot() *botapi.BotAPI {
	config := LoadConfig()

	BOT, err := botapi.NewBotAPI(config.Token)
	if err != nil {
		log.Panic(err)
	}

	BOT.Debug = true

	return BOT
}
