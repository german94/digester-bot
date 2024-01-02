package tbot

import (
	"digester-bot/gptassistant"
	"digester-bot/prompt"

	tele "gopkg.in/telebot.v3"
)

const (
	StatusSummary      = "summary"
	StatusMainConcepts = "main_concepts"
	StatusQuestion     = "question"
	StatusQA           = "qa"
)

type Status string

type State struct {
	Status             Status
	Assistant          gptassistant.Assistant
	SummaryParams      *prompt.SummaryParams
	MainConceptsParams *prompt.MainConceptsParams
	QuestionParams     *prompt.QuestionParams
	QAParams           *prompt.QAParams
	User               *tele.User
}

func (tb *TBot) getChatState(c tele.Context) *State {
	s, ok := tb.chatState.Load(c.Chat().ID)
	if !ok {
		s = &State{
			User: c.Sender(),
		}
		tb.chatState.Store(c.Chat().ID, s)
	}
	return s
}

func (tb *TBot) resetChatState(c tele.Context) {
	s, ok := tb.chatState.Load(c.Chat().ID)
	if ok {
		tb.chatState.Store(c.Chat().ID, &State{
			User:      s.User,
			Assistant: s.Assistant,
		})
	}
}
