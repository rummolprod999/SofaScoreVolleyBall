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
			Logging(err, "callback workWithResponse")
			return
		}
		scoring, _, _, err := jsonparser.Get(value, "events")
		//fmt.Println(string(scoring))
		if err != nil {
			Logging(err, "scoring")
			return
		}
		_, err = jsonparser.ArrayEach([]byte(scoring), t.VolleyBallMatch)
		if err != nil {
			Logging(err, "VolleyBallMatch")
			return
		}
	}, "sportItem", "tournaments")
	if err != nil {
		Logging(err, "sportItem tournaments")
		return
	}

}

func (t *SofaVolleyBall) VolleyBallMatch(value []byte, dataType jsonparser.ValueType, offset int, err error) {
	homeTeam, _, _, err := jsonparser.Get(value, "homeTeam", "name")
	if err != nil {
		Logging(err, "homeTeam")
		return
	}
	homeScore, _, _, err := jsonparser.Get(value, "homeScore")
	if err != nil {
		Logging(err, "homeScore")
		return
	}
	awayTeam, _, _, err := jsonparser.Get(value, "awayTeam", "name")
	if err != nil {
		Logging(err)
		return
	}
	awayScore, _, _, err := jsonparser.Get(value, "awayScore")
	if err != nil {
		Logging(err, "awayScore")
		return
	}
	statusType, _, _, err := jsonparser.Get(value, "status", "type")
	if err != nil {
		Logging(err, "statusType")
		return
	}
	/*fmt.Println(string(statusType))
	fmt.Println(string(homeTeam))
	fmt.Println(string(homeScore))
	fmt.Println(string(awayTeam))
	fmt.Println(string(awayScore))
	fmt.Println()*/
	volT := VolleyBall{homeTeam: string(homeTeam), homeScore: homeScore, awayTeam: string(awayTeam), awayScore: awayScore, statusType: string(statusType)}
	t.printMatch(volT)
}
func (t *SofaVolleyBall) printMatch(m VolleyBall) {
	fmt.Printf("Status game: %s\n", m.statusType)
	fmt.Printf("Home Team: %s\n", m.homeTeam)
	err := jsonparser.ObjectEach(m.homeScore, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		fmt.Printf("%s: %s\n", string(key), string(value))
		return nil
	})
	if err != nil {
		Logging(err, "printMatch", fmt.Sprintf("%s", string(m.homeScore)))
		return
	}
	fmt.Printf("Away Team: %s\n", m.awayTeam)
	err = jsonparser.ObjectEach(m.awayScore, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		fmt.Printf("%s: %s\n", string(key), string(value))
		return nil
	})
	if err != nil {
		Logging(err, "printMatch", fmt.Sprintf("%s", string(m.awayScore)))
		return
	}
	fmt.Printf("\n\n\n")
}
