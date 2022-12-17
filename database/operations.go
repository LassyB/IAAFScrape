package database

import (
	"container/list"
	"fmt"
)

func (db Database) Ping() error {
	err := db.SqlDb.PingContext(dbContext)
	if err == nil {
		fmt.Println("Database alive")
	}
	return err
}

func (db Database) GetActiveAthletes() (*list.List, error) {
	sqlStatement := fmt.Sprintf("SELECT Id, Name, IAAF FROM Athletes WHERE IAAF IS NOT NULL AND Published = 1")
	data, err := db.SqlDb.QueryContext(dbContext, sqlStatement)

	list := list.New()
	if err != nil {
		return list, err
	}
	for data.Next() {
		var Id, Name, IAAF string
		err = data.Scan(&Id, &Name, &IAAF)
		if err != nil {
			return list, err
		}
		athlete := Athlete{Id, Name, IAAF}
		list.PushBack(athlete)
	}
	return list, nil
}
