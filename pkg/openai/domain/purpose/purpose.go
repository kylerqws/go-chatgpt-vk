package purpose

import "fmt"

type Purpose struct {
	Value       string
	Description string
}

var (
	FineTune = Purpose{
		Value:       "fine-tune",
		Description: "Used to upload files for fine-tuning models.",
	}

	FineTuneResults = Purpose{
		Value:       "fine-tune-results",
		Description: "Used to upload results of fine-tuning jobs.",
	}

	Assistants = Purpose{
		Value:       "assistants",
		Description: "Used to upload files for the Assistants API.",
	}

	AssistantsOutput = Purpose{
		Value:       "assistants_output",
		Description: "Used to upload files generated by assistants.",
	}
)

var AllPurposes = map[string]*Purpose{
	FineTune.Value:         &FineTune,
	FineTuneResults.Value:  &FineTuneResults,
	Assistants.Value:       &Assistants,
	AssistantsOutput.Value: &AssistantsOutput,
}

func Resolve(code string) (*Purpose, error) {
	if code == "" {
		return &FineTune, nil
	}

	if purpose, ok := AllPurposes[code]; ok {
		return purpose, nil
	}

	return nil, fmt.Errorf("invalid purpose value for OpenAI API request: %q", code)
}
