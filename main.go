package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/LassyB/IAAFScrape/config"
	"github.com/LassyB/IAAFScrape/database"
	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
		return
	}
	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s", config.Server, config.Username, config.Password, config.Database)
	sqlObj, err := sql.Open("mssql", connectionString)
	if err != nil {
		fmt.Println(fmt.Errorf("error opening database: %v", err))
	}
	data := database.Database{
		SqlDb: sqlObj,
	}
	list, err := data.GetActiveAthletes()
	if err != nil {
		log.Println(err)
	}
	for e := list.Front(); e != nil; e = e.Next() {
		athlete := database.Athlete(e.Value.(database.Athlete))
		fmt.Println(athlete.Name)
	}
}
