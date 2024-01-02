package tbot

import (
	"digester-bot/futils"
	"digester-bot/gptassistant"
	"digester-bot/user_usage"
	"time"

	xsync "github.com/puzpuzpuz/xsync/v3"
	tele "gopkg.in/telebot.v3"
)

type TBot struct {
	Bot
	futils.FileHelper

	chatState       *xsync.MapOf[int64, *State]
	userRequests    *xsync.MapOf[int64, *user_usage.UserUsage]
	assitantBuilder gptassistant.Builder
}

func New(token string, assitantBuilder gptassistant.Builder) (*TBot, error) {
	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
	b, err := tele.NewBot(pref)
	if err != nil {
		return nil, err
	}

	result := new(&TeleBot{Bot: *b}, assitantBuilder, &futils.Service{})
	result.Handle(tele.OnText, result.handleMessage)
	result.Handle(tele.OnDocument, result.handleFile)

	return result, nil
}

func new(b Bot, assitantBuilder gptassistant.Builder, fileHelper futils.FileHelper) *TBot {
	return &TBot{
		Bot:             b,
		chatState:       xsync.NewMapOf[int64, *State](),
		userRequests:    xsync.NewMapOf[int64, *user_usage.UserUsage](),
		assitantBuilder: assitantBuilder,
		FileHelper:      fileHelper,
	}
}
