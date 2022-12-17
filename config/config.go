package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var (
	Server   string
	Database string
	Username string
	Password string
	config   *configStruct
)

type configStruct struct {
	Server   string `json:"server"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func ReadConfig() error {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal(err)
		return err
	}
	Server = config.Server
	Database = config.Database
	Username = config.Username
	Password = config.Password
	return nil
}
