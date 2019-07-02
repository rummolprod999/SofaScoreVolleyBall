package main

import (
	"fmt"
	"src/github.com/buger/jsonparser"
	"time"
)

type SofaVolleyBall struct {
}

func (t *SofaVolleyBall) run() {
	curDate := time.Now().Format("2006-01-02")
	UrlSofa := fmt.Sprintf("https://www.sofascore.com/volleyball//%s/json?_=%d", curDate, random(1000000000, 9999999999))
	response := DownloadPage(UrlSofa)
	t.workWithResponse(response)
}

func (t *SofaVolleyBall) workWithResponse(s string) {
	defer SaveStack()
	_, err := jsonparser.ArrayEach([]byte(s), func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		if err != nil {
			Logging(err)
			return
		}
		scoring, _, _, err := jsonparser.Get(value, "events")
		//fmt.Println(string(scoring))
		if err != nil {
			Logging(err)
			return
		}
		_, err = jsonparser.ArrayEach([]byte(scoring), t.VolleyBallMatch)
		if err != nil {
			Logging(err)
			return
		}
	}, "sportItem", "tournaments")
	if err != nil {
		Logging(err)
		return
	}

}

func (t *SofaVolleyBall) VolleyBallMatch(value []byte, dataType jsonparser.ValueType, offset int, err error) {
	homeTeam, _, _, err := jsonparser.Get(value, "homeTeam", "name")
	if err != nil {
		Logging(err)
		return
	}
	homeScore, _, _, err := jsonparser.Get(value, "homeScore")
	if err != nil {
		Logging(err)
		return
	}
	awayTeam, _, _, err := jsonparser.Get(value, "awayTeam", "name")
	if err != nil {
		Logging(err)
		return
	}
	awayScore, _, _, err := jsonparser.Get(value, "awayScore")
	if err != nil {
		Logging(err)
		return
	}
	fmt.Println(string(homeTeam))
	fmt.Println(string(homeScore))
	fmt.Println(string(awayTeam))
	fmt.Println(string(awayScore))
	fmt.Println()
}