package main

import (
	"log"

	"github.com/LassyB/IAAFScrape/config"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("You've hit main")
}
