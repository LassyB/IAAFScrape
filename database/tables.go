package database

import "time"

type Athlete struct {
	Id   string
	Name string
	IAAF string
}

type Result struct {
	Id        int
	Name      string
	Mark      string
	Date      time.Time
	AthleteId int
	Location  string
	Season    string
}
