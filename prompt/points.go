package prompt

import "fmt"

type MainConceptsParams struct {
	Topic    string
	Max      int
	Language string
}

func (mc *MainConceptsParams) Prompt() string {
	if mc.Language == "" {
		mc.Language = "english"
	}
	if mc.Max <= 0 {
		mc.Max = 10
	}
	return fmt.Sprintf(
		`Please provide me a list of (at most) %v sentences (written in %s language) briefly describing the main concepts related to "%v". Use - to start every new point, like:
-explanation of one thing
-explanation of another thing
...`, mc.Max, mc.Language, mc.Topic)
}
