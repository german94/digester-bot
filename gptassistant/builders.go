package gptassistant

type Builder interface {
	NewAssistant() Assistant
}

type GPTBuilder struct {
	apiKey string
}

func NewGPTBuilder(apiKey string) *GPTBuilder {
	return &GPTBuilder{
		apiKey: apiKey,
	}
}

func (gptb *GPTBuilder) NewAssistant() Assistant {
	return New(gptb.apiKey)
}

type MockBuilder struct {
	Response string
}

func (mb *MockBuilder) NewAssistant() Assistant {
	return &MockAssistant{Response: mb.Response}
}
