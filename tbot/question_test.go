package tbot

import (
	"digester-bot/gptassistant"
	"testing"

	"github.com/stretchr/testify/require"
	tele "gopkg.in/telebot.v3"
)

func TestQuestion(t *testing.T) {
	inputContent := "Christopher Columbus is credited with discovering the Americas in 1492. Americans get a day off work on October 10 to celebrate Columbus Day"
	expectedAnswer := "In 1492"

	tb := new(&MockBot{}, &gptassistant.MockBuilder{
		Response: expectedAnswer,
	}, &MockFileHelper{
		Content: []byte(inputContent),
	})

	c := &MockContext{
		UserID: int64(1),
		ChatID: int64(200),
		Msg: &tele.Message{
			Text:     "/start",
			Document: &tele.Document{},
		},
	}

	require.NoError(t, tb.handleFile(c))

	require.NoError(t, tb.handleQuestion(c))

	require.NotNil(t, tb.getChatState(c))
	require.NotNil(t, tb.getChatState(c).QuestionParams)

	c.Msg.Text = "when was America discovered?"
	require.NoError(t, tb.handleMessage(c))

	require.Equal(t, expectedAnswer, c.ReceivedMessage)
}
