package tbot

import (
	"context"
	"digester-bot/prompt"
	"errors"

	tele "gopkg.in/telebot.v3"
)

func (tb *TBot) handleSummary(c tele.Context) error {
	s := tb.getChatState(c)
	s.Status = StatusSummary
	if s.SummaryParams == nil {
		s.SummaryParams = &prompt.SummaryParams{}

		var langMenu *tele.ReplyMarkup
		langMenu = tb.summaryLangMenu(c, tb.summaryMaxWordsMenu(c, langMenu))
		return c.Edit("Please select a language", langMenu)
	}

	s.SummaryParams.Topic = c.Message().Text
	c.Send("Ok, I'm generating the summary...")
	resp, err := tb.Ask(context.TODO(), s, s.SummaryParams.Prompt())
	if errors.Is(err, ErrRequestsLimitExceeded) {
		return c.Send("You have exceeded the maximum amount of requests per day.")
	}
	if err != nil {
		s.Status = ""
		s.SummaryParams = nil
		return c.Send("There was an error, please select an option", tb.afterSummaryMenu())
	}

	s.Status = ""
	s.SummaryParams = nil

	return tb.Send(c, resp, "summary", tb.afterSummaryMenu())
}

func (tb *TBot) summaryLangMenu(c tele.Context, summaryMaxWordsMenu *tele.ReplyMarkup) *tele.ReplyMarkup {
	return tb.langSelectMenu(func(c tele.Context, lang string) error {
		tb.summarySelectLang(c, lang)

		return c.Edit("And how many words at most should the summary contain?", summaryMaxWordsMenu)
	}, func(c tele.Context) error {
		tb.resetChatState(c)

		return c.Send("Okay, here are the options", tb.mainMenu())
	})
}

func (tb *TBot) summarySelectLang(c tele.Context, lang string) {
	s := tb.getChatState(c)
	s.SummaryParams.Language = lang
}

func (tb *TBot) summaryMaxWordsMenu(c tele.Context, summaryLangMenu *tele.ReplyMarkup) *tele.ReplyMarkup {
	return tb.maxWordsMenu(func(c tele.Context, maxWords int) error {
		tb.summarySelectMaxWords(c, maxWords)

		return c.Edit("Perfect. One last thing, which topic of the content you provided you want me to summarize? you can give me just a title for example")
	}, func(c tele.Context) error {
		tb.resetChatState(c)

		return c.Send("Okay, here are the languages", summaryLangMenu)
	})
}

func (tb *TBot) summarySelectMaxWords(c tele.Context, maxWords int) {
	s := tb.getChatState(c)
	s.SummaryParams.MaxWords = maxWords
}

func (tb *TBot) afterSummaryMenu() *tele.ReplyMarkup {
	return tb.afterSuccessfulOutput("New summary", func(c tele.Context) error {
		tb.resetChatState(c)

		return tb.handleSummary(c)
	})
}
