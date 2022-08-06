package cmd

import "goBot/ctrs"

func Start(id int64, firstName, lastName, username string) {
	if username == "" {
		ctrs.SendMessage(id, "Для использования данного бота требуется указать в настройках @UserName")
		ctrs.SendToAdmin("Пользователь " + firstName + " " + lastName + " хотел получить доступ, но у него нет @username")
	} else {
		ctrs.SendMessage(id, "Привет "+firstName+"! Заявка на доступ к боту отправлена, ожидайте ответа рассмотрения.")

		ctrs.RequestForWhiteList(id, firstName, lastName, username)
	}
}
