package main

import (
	"fmt"
	"strings"
)

type VolleyBall struct {
	homeTeam     string
	homeScore    []byte
	homeScoreMap map[string]int
	awayTeam     string
	awayScore    []byte
	awayScoreMap map[string]int
	statusType   string
	id           int64
	changeDate   string
	seasonName   string
}

func (m *VolleyBall) printMatch() {
	if m.statusType == "notstarted" || m.statusType == "finished" || m.statusType == "canceled" {
		return
	}
	fmt.Printf("Id game: %d\n", m.id)
	fmt.Printf("Date Change: %s\n", m.changeDate)
	fmt.Printf("Status game: %s\n", m.statusType)
	fmt.Printf("Home Team: %s\n", m.homeTeam)
	for k, v := range m.homeScoreMap {
		fmt.Printf("%s: %d\n", k, v)
	}
	fmt.Printf("Away Team: %s\n", m.awayTeam)
	for k, v := range m.awayScoreMap {
		fmt.Printf("%s: %d\n", k, v)
	}
	fmt.Printf("\n\n\n")
	//SendToTelegram(m)
}

func (m *VolleyBall) sendMatch() {
	if m.statusType == "notstarted" || m.statusType == "finished" || m.statusType == "canceled" {
		return
	}
	for k, v := range m.awayScoreMap {
		if !strings.Contains(k, "period") {
			continue
		}
		if per, ok := m.homeScoreMap[k]; ok {
			if per == 21 && v == 21 {
				SendToTelegram(m, k)
			}
		}
	}
}
