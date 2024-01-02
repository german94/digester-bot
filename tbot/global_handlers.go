package tbot

import (
	"digester-bot/urlutils"
	"errors"
	"fmt"
	"log"

	tele "gopkg.in/telebot.v3"
)

func (tb *TBot) handleMessage(c tele.Context) error {
	s := tb.getChatState(c)

	if c.Message().Text == "/start" {
		return c.Send("Please provide me with either a file or link with the content that you want to learn from.")
	}
	switch s.Status {
	case StatusQuestion:
		return tb.handleQuestion(c)
	case StatusSummary:
		return tb.handleSummary(c)
	case StatusMainConcepts:
		return tb.handleMainConcepts(c)
	case StatusQA:
		return tb.handleQA(c)
	}
	msgText := c.Message().Text
	if urlutils.IsURL(msgText) {
		c.Send("Okay, give me a few seconds, I'm processing what you just sent me...")
		content, err := urlutils.ExtractContent(msgText)
		if err != nil {
			log.Printf("Error extracting content from website %v: %v\n", msgText, err)
			return c.Send("There was an error, please try again.")
		}
		if err := tb.initAssistant(c, []byte(content)); err != nil {
			log.Printf("Error initializing assistant from file: %v\n", err)
			return c.Send("There was an error, please try again.")
		}
		return c.Send("Alright, now I can help you with different methods, please select one.", tb.mainMenu())
	}
	return errors.New("no action")
}

func (tb *TBot) handleFile(c tele.Context) error {
	c.Send("Okay, give me a few seconds, I'm processing your file...")
	fileID := c.Message().Document.FileID
	fileContent, err := tb.getTelegramFile(fileID, c.Chat().ID)
	if err != nil {
		log.Printf("error getting telegram file %v: %v\n", fileID, err)
		return c.Send("there was an error, please try again")
	}
	if err := tb.initAssistant(c, fileContent); err != nil {
		log.Printf("error initializing assistant: %v\n", err)
		return c.Send("there was an error, please try again")
	}

	return tb.showMainMenu(c)
}

func (tb *TBot) getTelegramFile(fileID string, chatID int64) ([]byte, error) {
	f, err := tb.FileByID(fileID)
	if err != nil {
		return nil, err
	}

	fileURL := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", tb.Token(), f.FilePath)
	return tb.GetFileFromURL(fileURL)
}
