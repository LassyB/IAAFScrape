package iaaf

import (
	"container/list"
	"fmt"
	"net/http"
)

func GetAthleteResults(iaafCode string, year int) (*list.List, error) {
	list := list.New()
	resp, err := http.Get(fmt.Sprintf("https://www.worldathletics.org/data/GetCompetitorResultsByYearHtml?resultsByYear=%d&resultsByYearOrderBy=date&aaId=%s", year, iaafCode))
	if err != nil {
		return list, err
	}
	defer resp.Body.Close()
	return list, nil
}
