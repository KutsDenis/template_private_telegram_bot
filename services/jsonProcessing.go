package services

import (
	"encoding/json"
	"goBot/ctrs"
	"io/ioutil"
	"log"
	"os"
)

func WriteUserToJSON(id int64, path, firstName, lastName, username string, approve bool) {
	usr := ctrs.User{
		Id:        id,
		FirstName: firstName,
		LastName:  lastName,
		UserName:  username,
		WhiteList: approve,
	}
	data, _ := json.Marshal(usr)

	err := ioutil.WriteFile(path, data, 0644)
	if err != nil {
		ctrs.SendLogs("Ошибка записи файла")
	}
}

func ReadJSON(path string) ctrs.User {
	file, err := os.Open(path)

	if err != nil {
		return ctrs.User{}
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

	var usr ctrs.User

	jsonErr := json.Unmarshal(data, &usr)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return usr
}
