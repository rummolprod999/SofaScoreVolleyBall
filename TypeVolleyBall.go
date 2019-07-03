package main

import (
	"fmt"
	"src/github.com/buger/jsonparser"
)

type VolleyBall struct {
	homeTeam   string
	homeScore  []byte
	awayTeam   string
	awayScore  []byte
	statusType string
	id         int64
	changeDate string
}

func (m *VolleyBall) printMatch() {
	fmt.Printf("Id game: %d\n", m.id)
	fmt.Printf("Date Change: %s\n", m.changeDate)
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
