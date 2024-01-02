package tbot

import (
	"digester-bot/gptassistant"
	"testing"

	"github.com/stretchr/testify/require"
	tele "gopkg.in/telebot.v3"
)

func TestSummary(t *testing.T) {
	inputContent := "this is the input content that is supposed to be summarized"

	expectedSummary := "this is the summary"
	tb := new(&MockBot{}, &gptassistant.MockBuilder{
		Response: expectedSummary,
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

	require.NoError(t, tb.handleSummary(c))

	require.NotNil(t, tb.getChatState(c))
	require.NotNil(t, tb.getChatState(c).SummaryParams)

	tb.summarySelectLang(c, "english")
	require.Equal(t, "english", tb.getChatState(c).SummaryParams.Language)

	tb.summarySelectMaxWords(c, 100)
	require.Equal(t, 100, tb.getChatState(c).SummaryParams.MaxWords)

	c.Msg.Text = "chapter 1"
	require.NoError(t, tb.handleMessage(c))
	require.Equal(t, expectedSummary, c.ReceivedMessage)
}
