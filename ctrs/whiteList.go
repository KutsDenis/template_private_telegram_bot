package ctrs

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func GetWhiteList(username string) bool {
	path := "data/users/" + username + ".json"

	usr := readJSON(path)

	if usr.WhiteList == true {
		return true
	}
	return false
}

func RequestForWhiteList(id int64, firstName, lastName, username string) {
	SendToAdmin("Пользователь @" + username + " запрашивает доступ.\n/approving")

	pathUsers := "data/users/" + username + ".json"
	pathConsider := "data/for_consider"

	writeUserToJSON(id, pathUsers, firstName, lastName, username, false)

	tempData := readFile(pathConsider)
	tempData += username + " "

	writeFile(pathConsider, tempData)

}

func ConsiderWhiteList() {
	pathConsider := "data/for_consider"

	data := readFile(pathConsider)
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

	data := readFile(pathConsider)
	usernames := strings.Fields(data)

	path := "data/users/" + usernames[0] + ".json"

	usr := readJSON(path)

	for i := 1; i < len(usernames); i++ {
		newData = newData + usernames[i] + " "
	}

	writeFile(pathConsider, newData)

	writeUserToJSON(usr.Id, path, usr.FirstName, usr.LastName, usr.UserName, approve)

	if approve {
		SendToAdmin("@" + usr.UserName + " теперь в белом списке")
		SendMessage(usr.Id, "Вам одобрили доступ")
	} else {
		SendToAdmin("Доступ для " + "@" + usr.UserName + " отклонен")
		SendMessage(usr.Id, "Вам отказали в доступе")
	}
}

func writeUserToJSON(id int64, path, firstName, lastName, username string, approve bool) {
	usr := User{
		Id:        id,
		FirstName: firstName,
		LastName:  lastName,
		UserName:  username,
		WhiteList: approve,
	}
	data, _ := json.Marshal(usr)

	err := ioutil.WriteFile(path, data, 0644)
	if err != nil {
		SendLogs("Ошибка записи файла")
	}
}

func readJSON(path string) User {
	file, err := os.Open(path)

	if err != nil {
		return User{}
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var usr User

	jsonErr := json.Unmarshal(data, &usr)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return usr
}

func writeFile(path, data string) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			SendLogs("Не получилось закрыть файл")
		}
	}(file)

	_, err = file.WriteString(data)
	if err != nil {
		panic(err)
	}
}

func readFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			SendLogs("Не пполучилось закрыть файл")
		}
	}(file)

	data := make([]byte, 64)

	for {
		n, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		return string(data[:n])
	}
	return ""
}
