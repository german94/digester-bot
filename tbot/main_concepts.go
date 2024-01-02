package tbot

import (
	"context"
	"digester-bot/prompt"
	"errors"

	tele "gopkg.in/telebot.v3"
)

func (tb *TBot) handleMainConcepts(c tele.Context) error {
	s := tb.getChatState(c)
	s.Status = StatusMainConcepts
	if s.MainConceptsParams == nil {
		s.MainConceptsParams = &prompt.MainConceptsParams{}

		var langMenu *tele.ReplyMarkup
		langMenu = tb.mainConceptsLangMenu(c, tb.mainConceptsMaxItemsMenu(c, langMenu))
		return c.Edit("Please select a language", langMenu)
	}

	s.MainConceptsParams.Topic = c.Message().Text
	c.Send("Ok, I'm generating the list...")
	resp, err := tb.Ask(context.TODO(), s, s.MainConceptsParams.Prompt())
	if errors.Is(err, ErrRequestsLimitExceeded) {
		return c.Send("You have exceeded the maximum amount of requests per day.")
	}
	if err != nil {
		s.Status = ""
		s.MainConceptsParams = nil
		return c.Send("There was an error, please select an option", tb.afterMainConceptsMenu())
	}

	s.Status = ""
	s.MainConceptsParams = nil
	return tb.Send(c, resp, "main_points", tb.afterMainConceptsMenu())
}

func (tb *TBot) mainConceptsLangMenu(c tele.Context, mainConceptsMaxItemsMenu *tele.ReplyMarkup) *tele.ReplyMarkup {
	return tb.langSelectMenu(func(c tele.Context, lang string) error {
		tb.mainConceptsSelectLang(c, lang)

		return c.Edit("And how many points at most should the list contain?", mainConceptsMaxItemsMenu)
	}, func(c tele.Context) error {
		tb.resetChatState(c)

		return c.Send("Okay, here are the options", tb.mainMenu())
	})
}

func (tb *TBot) mainConceptsSelectLang(c tele.Context, lang string) {
	s := tb.getChatState(c)
	s.MainConceptsParams.Language = lang
}

func (tb *TBot) mainConceptsMaxItemsMenu(c tele.Context, mainConceptsLangMenu *tele.ReplyMarkup) *tele.ReplyMarkup {
	return tb.maxItemsMenu(func(c tele.Context, max int) error {
		tb.mainConceptsSelectMax(c, max)

		return c.Edit("Perfect. One last thing, which topic of the content you provided you want me to make to work with? you can give me just a title for example")
	}, func(c tele.Context) error {
		tb.resetChatState(c)

		return c.Send("Okay, here are the languages", mainConceptsLangMenu)
	})
}

func (tb *TBot) mainConceptsSelectMax(c tele.Context, max int) {
	s := tb.getChatState(c)
	s.MainConceptsParams.Max = max
}

func (tb *TBot) afterMainConceptsMenu() *tele.ReplyMarkup {
	return tb.afterSuccessfulOutput("New list", func(c tele.Context) error {
		tb.resetChatState(c)
		return tb.handleMainConcepts(c)
	})
}
