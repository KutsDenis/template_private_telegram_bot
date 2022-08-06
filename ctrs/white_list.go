package ctrs

import (
	"goBot/services"
	"strings"
)

func GetWhiteList(username string) bool {
	path := "data/users/" + username + ".json"

	usr := services.ReadJSON(path)

	if usr.WhiteList == true {
		return true
	}
	return false
}

func RequestForWhiteList(id int64, firstName, lastName, username string) {
	SendToAdmin("Пользователь @" + username + " запрашивает доступ.\n/approving")

	pathUsers := "data/users/" + username + ".json"
	pathConsider := "data/for_consider"

	services.WriteUserToJSON(id, pathUsers, firstName, lastName, username, false)

	tempData := services.ReadFile(pathConsider)
	tempData += username + " "

	services.WriteFile(pathConsider, tempData)

}

func ConsiderWhiteList() {
	pathConsider := "data/for_consider"

	data := services.ReadFile(pathConsider)
	if data == "" {
		SendToAdmin("Пустой список")
	} else {
		usernames := strings.Fields(data)

		SendKeyboardToAdmin(approveKeyboard, "Одобрить доступ @"+usernames[0]+"?")
	}
}

func ApproveWhiteList(approve bool) {
	var newData string

	pathConsider := "data/for_consider"

	data := services.ReadFile(pathConsider)
	usernames := strings.Fields(data)

	path := "data/users/" + usernames[0] + ".json"

	usr := services.ReadJSON(path)

	for i := 1; i < len(usernames); i++ {
		newData = newData + usernames[i] + " "
	}

	services.WriteFile(pathConsider, newData)

	services.WriteUserToJSON(usr.Id, path, usr.FirstName, usr.LastName, usr.UserName, approve)

	if approve {
		SendToAdmin("@" + usr.UserName + " теперь в белом списке")
		SendMessage(usr.Id, "Вам одобрили доступ")
	} else {
		SendToAdmin("Доступ для " + "@" + usr.UserName + " отклонен")
		SendMessage(usr.Id, "Вам отказали в доступе")
	}
}
