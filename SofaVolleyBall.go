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
		jsonparser.ArrayEach([]byte(scoring), func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			val, _, _, _ := jsonparser.Get(value, "homeScore")
			fmt.Println(string(val))
		})

	}, "sportItem", "tournaments")
	if err != nil {
		Logging(err)
		return
	}

}
