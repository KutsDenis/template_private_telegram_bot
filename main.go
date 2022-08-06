package main

import (
	"goBot/cfg"
	"goBot/cmd"
	"goBot/ctrs"

	botapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot := cfg.InitBot()

	ctrs.SendLogs("Authorized on account " + bot.Self.UserName)

	u := botapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		var text string

		if update.CallbackQuery != nil {

			if update.CallbackQuery.Data == "approve" {
				go ctrs.ApproveWhiteList(true)
			} else if update.CallbackQuery.Data == "decline" {
				go ctrs.ApproveWhiteList(false)
			}
		}

		if update.Message == nil { // Игнорирование любых обновлений не касающихся сообщений
			continue
		}

		if !update.Message.IsCommand() { // Игнорировать все сообщения кроме команд и управления админа
			continue
		}

		// Команды
		if ctrs.GetWhiteList(update.SentFrom().UserName) {
			usr := update.SentFrom()

			switch update.Message.Command() {
			case "help":
				text = "Не моли о помощи"
			case "start":
				text = "https://memepedia.ru/wp-content/uploads/2017/05/%D1%8F-%D1%81%D0%BA%D0%B0%D0%B7%D0%B0%D0%BB%D0%B0-%D1%81%D1%82%D0%B0%D1%80%D1%82%D1%83%D0%B5%D0%BC.jpg"
			case "approving":
				go ctrs.ConsiderWhiteList()
			default:
				text = "Неизвестная команда"
			}

			if text != "" {
				ctrs.SendMessage(usr.ID, text)
			}
		} else {
			usr := update.SentFrom()

			if update.Message.Command() == "start" {
				go cmd.Start(usr.ID, usr.FirstName, usr.LastName, usr.UserName)
			} else {
				text = "Доступ запрещен"
			}
			if text != "" {
				ctrs.SendMessage(usr.ID, text)
			}
		}
	}
}
