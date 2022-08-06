package ctrs

import (
	botapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"goBot/cfg"
	"log"
)

var bot = cfg.InitBot()

var approveKeyboard = botapi.NewInlineKeyboardMarkup(
	botapi.NewInlineKeyboardRow(
		botapi.NewInlineKeyboardButtonData("Одобрить", "approve"),
		botapi.NewInlineKeyboardButtonData("Отклонить", "decline"),
	))

func SendMessage(chatID int64, text string) {

	msg := botapi.NewMessage(chatID, text)

	if _, err := bot.Send(msg); err != nil {
		SendLogs(err.Error())
		log.Panic(err)
	}
}

func SendMarkup(chatID int64, markup botapi.InlineKeyboardMarkup, text string) {
	msg := botapi.NewMessage(chatID, text)

	msg.ReplyMarkup = markup

	if _, err := bot.Send(msg); err != nil {
		SendLogs(err.Error())
		log.Panic(err)
	}
}

func SendLogs(text string) {
	SendMessage(cfg.LoadConfig().ChatLogs, text)
}

func SendToAdmin(text string) {
	SendMessage(cfg.LoadConfig().Admin, text)
}

func SendKeyboardToAdmin(markup botapi.InlineKeyboardMarkup, text string) {
	SendMarkup(cfg.LoadConfig().Admin, markup, text)
}
