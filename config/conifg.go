package config

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type Config struct {
	Database Database `json:"database"`
}

type Database struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func ReadConfig() (config Config) {
	//获取路径./config.json下的配置选项
	file, err := os.Open("./config.json")
	if err != nil {
		log.Println(err)
		time.Sleep(10 * time.Second)
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Println(err)
		time.Sleep(10 * time.Second)
		panic(err)
	}

	return config
}
