package tbot

import (
	"digester-bot/gptassistant"
	"testing"

	"github.com/stretchr/testify/require"
	tele "gopkg.in/telebot.v3"
)

func TestQA(t *testing.T) {
	inputContent := "Christopher Columbus is credited with discovering the Americas in 1492. Americans get a day off work on October 10 to celebrate Columbus Day"
	expectedQuestion := "When was America discovered?"

	tb := new(&MockBot{}, &gptassistant.MockBuilder{
		Response: expectedQuestion,
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

	require.NoError(t, tb.handleQA(c))
	require.NotNil(t, tb.getChatState(c))
	require.NotNil(t, tb.getChatState(c).QAParams)

	c.Msg.Text = "chapter 1"
	require.NoError(t, tb.handleMessage(c))
	require.Equal(t, "chapter 1", tb.getChatState(c).QAParams.Topic)
	require.Equal(t, expectedQuestion, tb.getChatState(c).QAParams.Question)

	c.Msg.Text = "In 1900"
	expectedCorrection := "America was not discovered in 1900. America was discoverd in 1492 by Christopher Columbus."
	tb.getChatState(c).Assistant = &gptassistant.MockAssistant{
		Response: expectedCorrection,
	}
	require.NoError(t, tb.handleMessage(c))
	require.Equal(t, expectedCorrection, c.ReceivedMessage)
}
