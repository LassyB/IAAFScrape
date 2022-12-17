package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var (
	config *configStruct
)

type configStruct struct {
	ConnectionString string `json:"connectionString"`
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
	return nil
}
