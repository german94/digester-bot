package gptassistant

import (
	"context"
	"log"
)

type MockAssistant struct {
	Question string
	Response string
}

func (*MockAssistant) Init(ctx context.Context, fileContent []byte, fileName string) error {
	return nil
}

func (am *MockAssistant) Ask(ctx context.Context, question string) (string, error) {
	am.Question = question
	log.Printf("question: %v\nresponse: %v\n", am.Question, am.Response)
	return am.Response, nil
}
