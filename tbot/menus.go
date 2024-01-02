package tbot

import (
	"strconv"

	tele "gopkg.in/telebot.v3"
)

func (tb *TBot) mainMenu() *tele.ReplyMarkup {
	menu := &tele.ReplyMarkup{}

	btnSummarize := menu.Data("Summarize", "summarize")
	btnPoints := menu.Data("Main Points", "points")
	btnAsk := menu.Data("Ask a question", "ask")
	btnQA := menu.Data("Test", "qa")
	btnStartAgain := menu.Data("Provide new source", "startAgain")

	tb.Handle(&btnAsk, tb.handleQuestion)

	tb.Handle(&btnSummarize, tb.handleSummary)

	tb.Handle(&btnPoints, tb.handleMainConcepts)

	tb.Handle(&btnQA, tb.handleQA)

	tb.Handle(&btnStartAgain, func(c tele.Context) error {
		tb.resetChatState(c)

		return c.Send("Please provide me with either a file or link with the content that you want to learn from.")
	})

	menu.Inline(
		menu.Row(btnSummarize),
		menu.Row(btnPoints),
		menu.Row(btnAsk),
		menu.Row(btnQA),
		menu.Row(btnStartAgain),
	)
	return menu
}

func (tb *TBot) maxWordsMenu(setLangFn func(c tele.Context, maxWords int) error, backHandler func(c tele.Context) error) *tele.ReplyMarkup {
	menu := &tele.ReplyMarkup{}

	btn50 := menu.Data("50", "btn50Words")
	btn100 := menu.Data("100", "btn100Words")
	btn500 := menu.Data("500", "btn500Words")
	btn1000 := menu.Data("1000", "btn1000Words")
	btn2000 := menu.Data("2000", "btn2000Words")
	btn5000 := menu.Data("5000", "btn5000Words")
	for _, btn := range []tele.Btn{btn50, btn100, btn500, btn500, btn1000, btn2000, btn5000} {
		btn := btn
		tb.Handle(&btn, func(c tele.Context) error {
			val, _ := strconv.Atoi(btn.Text)
			return setLangFn(c, val)
		})
	}

	btnBack := menu.Data("Back", "btnBack")
	tb.Handle(&btnBack, backHandler)

	menu.Inline(
		menu.Row(btn50, btn100),
		menu.Row(btn500, btn1000),
		menu.Row(btn2000, btn5000),
		menu.Row(btnBack),
	)
	return menu
}

func (tb *TBot) maxItemsMenu(setLangFn func(c tele.Context, maxWords int) error, backHandler func(c tele.Context) error) *tele.ReplyMarkup {
	menu := &tele.ReplyMarkup{}

	btn3 := menu.Data("3", "btn3Items")
	btn10 := menu.Data("10", "btn10Items")
	btn15 := menu.Data("15", "btn15Items")
	btn20 := menu.Data("20", "btn20Items")
	for _, btn := range []tele.Btn{btn3, btn10, btn15, btn20} {
		btn := btn
		tb.Handle(&btn, func(c tele.Context) error {
			val, _ := strconv.Atoi(btn.Text)
			return setLangFn(c, val)
		})
	}

	btnBack := menu.Data("Back", "btnBack")
	tb.Handle(&btnBack, backHandler)

	menu.Inline(
		menu.Row(btn3, btn10),
		menu.Row(btn15, btn20),
		menu.Row(btnBack),
	)
	return menu
}

func (tb *TBot) langSelectMenu(setLangFn func(c tele.Context, lang string) error, backHandler func(c tele.Context) error) *tele.ReplyMarkup {
	menu := &tele.ReplyMarkup{}

	btnEn := menu.Data("English", "btnEn")
	btnSp := menu.Data("Spanish", "btnSp")
	btnIt := menu.Data("Italian", "btnIt")
	btnCh := menu.Data("Chinese", "btnCh")
	btnJa := menu.Data("Japanese", "btnJp")
	btnKo := menu.Data("Korean", "btnKo")
	btnRu := menu.Data("Russian", "btnRu")
	btnPo := menu.Data("Portugese", "btnPo")
	for _, btn := range []tele.Btn{btnEn, btnSp, btnIt, btnCh, btnJa, btnKo, btnRu, btnPo} {
		btn := btn
		tb.Handle(&btn, func(c tele.Context) error {
			return setLangFn(c, btn.Text)
		})
	}

	btnBack := menu.Data("Back", "btnBack")
	tb.Handle(&btnBack, backHandler)

	menu.Inline(
		menu.Row(btnEn, btnSp),
		menu.Row(btnIt, btnCh),
		menu.Row(btnJa, btnKo),
		menu.Row(btnRu, btnPo),
		menu.Row(btnBack),
	)
	return menu
}

func (tb *TBot) afterSuccessfulOutput(startOverLabel string, startOverHandler func(c tele.Context) error) *tele.ReplyMarkup {
	menu := &tele.ReplyMarkup{}
	startOverBtn := menu.Data(startOverLabel, "btnStartOver")
	backToMenu := menu.Data("Back to menu", "backToMenu")

	tb.Handle(&startOverBtn, startOverHandler)
	tb.Handle(&backToMenu, func(c tele.Context) error {
		tb.resetChatState(c)

		return c.Send("Okay, here are the options", tb.mainMenu())
	})
	menu.Inline(
		menu.Row(startOverBtn),
		menu.Row(backToMenu),
	)

	return menu
}

func (tb *TBot) showMainMenu(c tele.Context) error {
	return c.Send("Alright, now I can help you with different methods, please select one.", tb.mainMenu())
}
