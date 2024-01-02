package tbot

import (
	"context"
	"digester-bot/prompt"
	"errors"

	tele "gopkg.in/telebot.v3"
)

func (tb *TBot) handleQA(c tele.Context) error {
	s := tb.getChatState(c)
	s.Status = StatusQA
	if s.QAParams == nil {
		s.QAParams = &prompt.QAParams{}
		return c.Send("Ok, please tell me about which topic (you can give me just a title for example) you want to get a question")
	}

	if s.QAParams.Question == "" {
		if c.Message().Text == "" {
			return c.Send("Please tell me about which topic (you can give me just a title for example) you want to get a question")
		}
		s.QAParams.Topic = c.Message().Text
		c.Send("Ok, I'm generating the question...")
		resp, err := tb.Ask(context.TODO(), s, s.QAParams.Prompt())
		if errors.Is(err, ErrRequestsLimitExceeded) {
			return c.Send("You have exceeded the maximum amount of requests per day.")
		}
		if err != nil {
			s.Status = ""
			s.QAParams = nil
			return c.Edit("There was an error, please select an option", tb.afterQAMenu())
		}
		s.QAParams.Question = resp
		resp += "\nYou can answer this question and I will tell you if you're right or not"
		return tb.Send(c, resp, "qa", tb.afterQAMenu())
	}

	c.Send("Let me see...")
	s.QAParams.PossibleAnswer = c.Message().Text
	resp, err := tb.Ask(context.TODO(), s, s.QAParams.Prompt())
	if errors.Is(err, ErrRequestsLimitExceeded) {
		return c.Send("You have exceeded the maximum amount of requests per day.")
	}
	if err != nil {
		s.Status = ""
		s.QAParams = nil
		return c.Send("There was an error, please select an option", tb.afterQAMenu())
	}

	s.Status = ""
	s.QAParams = nil

	return tb.Send(c, resp, "validation", tb.afterQAMenu())
}

func (tb *TBot) afterQAMenu() *tele.ReplyMarkup {
	return tb.afterSuccessfulOutput("New question", func(c tele.Context) error {
		tb.resetChatState(c)

		return tb.handleQA(c)
	})
}
