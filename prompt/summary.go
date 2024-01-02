package prompt

import "fmt"

type SummaryParams struct {
	Topic    string
	MaxWords int
	Language string
}

func (sp *SummaryParams) Prompt() string {
	if sp.Language == "" {
		sp.Language = "english"
	}
	if sp.MaxWords <= 0 {
		sp.MaxWords = 100
	}
	return fmt.Sprintf(
		`Please provide me a summary that explains the "%s", written in %s language and with no more than %v words.`,
		sp.Topic, sp.Language, sp.MaxWords)
}
