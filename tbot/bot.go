package tbot

import tele "gopkg.in/telebot.v3"

type Bot interface {
	Handle(endpoint interface{}, h tele.HandlerFunc, m ...tele.MiddlewareFunc)
	Token() string
	Start()
	FileByID(fileID string) (tele.File, error)
}

type TeleBot struct {
	tele.Bot
}

func (t *TeleBot) Token() string {
	return t.Bot.Token
}
