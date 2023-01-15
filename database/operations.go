package database

import (
	"container/list"
	"database/sql"
	"fmt"
	"time"
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

func (db Database) GetAthleteResults(athlete Athlete) (*list.List, error) {
	sqlStatement := fmt.Sprintf("SELECT Id, Name, Mark, Date, AthleteId, Location, Season FROM Results WHERE AthleteId = %s", athlete.Id)
	data, err := db.SqlDb.QueryContext(dbContext, sqlStatement)

	list := list.New()
	if err != nil {
		return list, err
	}
	for data.Next() {
		var Id, AthleteId int
		var Name, Mark, Location, Season sql.NullString
		var Date time.Time
		err = data.Scan(&Id, &Name, &Mark, &Date, &AthleteId, &Location, &Season)
		if err != nil {
			return list, err
		}
		result := Result{Id, Name.String, Mark.String, Date, AthleteId, Location.String, Season.String}
		list.PushBack(result)
	}
	return list, nil
}
