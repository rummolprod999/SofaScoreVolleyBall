package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

func SendToTelegram(m *VolleyBall, period string, score int) {
	if !CheckIfExist(fmt.Sprintf("%d", m.id), period, score) {
		return
	}
	bot, err := tgbotapi.NewBotAPI(BotToken)
	if err != nil {
		Logging(err)
	}
	msg := tgbotapi.NewMessage(ChannelId, CreateMessage(m, period, score))
	msg.ParseMode = "html"
	_, err = bot.Send(msg)
	if err != nil {
		Logging(err)
	}
	Logging("send message")
}

func CheckIfExist(id_game, period string, score int) bool {
	db, err := DbConnection()
	if err != nil {
		Logging(err)
		return true
	}
	defer db.Close()
	rows, err := db.Query("SELECT id FROM sofa WHERE id_game=$1 AND period=$2 AND score=$3", id_game, period, score)
	if err != nil {
		Logging(err)
		return true
	}
	if rows.Next() {
		rows.Close()
		return false
	}
	rows.Close()
	_, err = db.Exec("INSERT INTO sofa (id, id_game, period, score) VALUES (NULL, $1, $2, $3)", id_game, period, score)
	if err != nil {
		Logging(err)
		return true
	}
	return true
}

func CreateMessage(m *VolleyBall, period string, score int) string {
	message := ""
	seasName, err := strconv.Unquote("\"" + m.seasonName + "\"")
	if err != nil {
		seasName = m.seasonName
	}
	message += fmt.Sprintf("<b>Season name:</b> %s\n", seasName)
	//message += fmt.Sprintf("<b>Date Change:</b> %s\n", m.changeDate)
	//message += fmt.Sprintf("<b>Status game:</b> %s\n", m.statusType)
	message += fmt.Sprintf("\n")
	homeTeam, err := strconv.Unquote("\"" + m.homeTeam + "\"")
	if err != nil {
		homeTeam = m.homeTeam
	}
	message += fmt.Sprintf("<b>Home Team:</b> %s\n", homeTeam)
	awayTeam, err := strconv.Unquote("\"" + m.awayTeam + "\"")
	if err != nil {
		awayTeam = m.awayTeam
	}
	message += fmt.Sprintf("<b>Away Team:</b> %s\n", awayTeam)
	message += fmt.Sprintf("\n")
	message += fmt.Sprintf("<b>%s:</b> %d:%d\n", period, score, score)
	return message
}
