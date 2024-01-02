package tbot

import (
	"context"

	tele "gopkg.in/telebot.v3"
)

func (tb *TBot) initAssistant(c tele.Context, fileContent []byte) error {
	s := tb.getChatState(c)
	s.Assistant = tb.assitantBuilder.NewAssistant()
	if err := s.Assistant.Init(context.TODO(), fileContent, "file"); err != nil {
		return err
	}
	return nil
}
