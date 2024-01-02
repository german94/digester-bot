package tbot

import (
	"context"
	"digester-bot/prompt"
	"errors"

	tele "gopkg.in/telebot.v3"
)

func (tb *TBot) handleQuestion(c tele.Context) error {
	s := tb.getChatState(c)
	s.Status = StatusQuestion
	if s.QuestionParams == nil {
		s.QuestionParams = &prompt.QuestionParams{}

		menu := &tele.ReplyMarkup{}
		backToMenu := menu.Data("Back to menu", "btnBackToMenu")
		tb.Handle(&backToMenu, func(c tele.Context) error {
			s := tb.getChatState(c)
			s.Status = ""
			s.QuestionParams = nil
			return c.Send("Okay, here are the options", tb.mainMenu())
		})
		menu.Inline(
			menu.Row(backToMenu),
		)

		return c.Send("Perfect, so what is your question?", menu)
	}

	c.Send("Ok, I'm processing your question...")
	s.QuestionParams.Question = c.Message().Text
	resp, err := tb.Ask(context.TODO(), s, s.QuestionParams.Question)
	if errors.Is(err, ErrRequestsLimitExceeded) {
		return c.Send("You have exceeded the maximum amount of requests per day.")
	}
	if err != nil {
		s.Status = ""
		s.QuestionParams = nil
		return c.Send("There was an error, please select an option", tb.afterQuestionMenu())
	}

	s.Status = ""
	s.QuestionParams = nil
	return tb.Send(c, resp, "answer", tb.afterQuestionMenu())
}

func (tb *TBot) afterQuestionMenu() *tele.ReplyMarkup {
	return tb.afterSuccessfulOutput("Ask new question", func(c tele.Context) error {
		tb.resetChatState(c)

		return tb.handleQuestion(c)
	})
}
