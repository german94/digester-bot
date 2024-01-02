package prompt

import "fmt"

type QuestionParams struct {
	Question string
}

func Question(params *QuestionParams) string {
	return fmt.Sprintf(
		`Answer me the following question in three double quotes: """%s"""
`, params.Question)
}
