package services

import (
	"fmt"
	"goBot/ctrs"
	"io"
	"os"
)

func WriteFile(path, data string) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			ctrs.SendLogs("Не получилось закрыть файл")
		}
	}(file)

	_, err = file.WriteString(data)
	if err != nil {
		panic(err)
	}
}

func ReadFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			ctrs.SendLogs("Не пполучилось закрыть файл")
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
