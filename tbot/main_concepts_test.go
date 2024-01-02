package tbot

import (
	"digester-bot/gptassistant"
	"testing"

	"github.com/stretchr/testify/require"
	tele "gopkg.in/telebot.v3"
)

func TestMainConcepts(t *testing.T) {
	inputContent := "Christopher Columbus is credited with discovering the Americas in 1492. Americans get a day off work on October 10 to celebrate Columbus Day"

	expectedMainConcepts := "-Colombus discovered America\nThis happened in 1492\nAmericans celebrate this on October 10"
	tb := new(&MockBot{}, &gptassistant.MockBuilder{
		Response: expectedMainConcepts,
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

	require.NoError(t, tb.handleMainConcepts(c))

	require.NotNil(t, tb.getChatState(c))
	require.NotNil(t, tb.getChatState(c).MainConceptsParams)

	tb.mainConceptsSelectLang(c, "english")
	require.Equal(t, "english", tb.getChatState(c).MainConceptsParams.Language)

	tb.mainConceptsSelectMax(c, 100)
	require.Equal(t, 100, tb.getChatState(c).MainConceptsParams.Max)

	c.Msg.Text = "chapter 1"
	require.NoError(t, tb.handleMessage(c))
	require.Equal(t, expectedMainConcepts, c.ReceivedMessage)
}
