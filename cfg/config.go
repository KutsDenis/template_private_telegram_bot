package cfg

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Token    string
	ChatLogs int64
	Admin    int64
}

func LoadConfig() Config {
	configPath := "cfg/config.json"

	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}
