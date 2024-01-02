package prompt

import "fmt"

type QAParams struct {
	Topic          string
	Question       string
	PossibleAnswer string
}

func (qa *QAParams) Prompt() string {
	if qa.PossibleAnswer == "" {
		return qa.questionPrompt()
	}
	return qa.validationPrompt()
}

func (qa *QAParams) questionPrompt() string {
	return fmt.Sprintf(
		`Please review the content from the file, specifically the part related to "%s", and make a question (so I can test myself). Write the question in the same language as the content."`, qa.Topic)
}

func (qa *QAParams) validationPrompt() string {
	return fmt.Sprintf(
		`Please review the content from the file, specifically the part related to "%s" and tell me if the answer "%s" correct to the question "%s". Please answer using the same language as the question.`,
		qa.Topic, qa.PossibleAnswer, qa.Question)
}
