package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func SendToTelegram(m *VolleyBall, period string) {
	if !CheckIfExist(fmt.Sprintf("%d", m.id), period) {
		return
	}
	bot, err := tgbotapi.NewBotAPI(BotToken)
	if err != nil {
		Logging(err)
	}
	msg := tgbotapi.NewMessage(ChannelId, CreateMessage(m, period))
	msg.ParseMode = "html"
	_, err = bot.Send(msg)
	if err != nil {
		Logging(err)
	}
	Logging("send message")
}

func CheckIfExist(id_game, period string) bool {
	db, err := DbConnection()
	if err != nil {
		Logging(err)
		return true
	}
	defer db.Close()
	rows, err := db.Query("SELECT id FROM sofa WHERE id_game=$1 AND period=$2", id_game, period)
	if err != nil {
		Logging(err)
		return true
	}
	if rows.Next() {
		rows.Close()
		return false
	}
	rows.Close()
	_, err = db.Exec("INSERT INTO sofa (id, id_game, period) VALUES (NULL, $1, $2)", id_game, period)
	if err != nil {
		Logging(err)
		return true
	}
	return true
}

func CreateMessage(m *VolleyBall, period string) string {
	message := ""
	message += fmt.Sprintf("<b>Season name:</b> %s\n", m.seasonName)
	//message += fmt.Sprintf("<b>Date Change:</b> %s\n", m.changeDate)
	//message += fmt.Sprintf("<b>Status game:</b> %s\n", m.statusType)
	message += fmt.Sprintf("\n")
	message += fmt.Sprintf("<b>Home Team:</b> %s\n", m.homeTeam)
	message += fmt.Sprintf("<b>Away Team:</b> %s\n", m.awayTeam)
	message += fmt.Sprintf("\n")
	message += fmt.Sprintf("<b>%s:</b> 21:21\n", period)
	return message
}
