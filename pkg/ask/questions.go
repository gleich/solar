package ask

import (
	"github.com/AlecAivazis/survey/v2"
)

// Data from the questions ask to the user
type QuestionsData struct {
	PAT   string
	GBCap int
}

// Ask the user a number of questions before running
func Questions() (QuestionsData, error) {
	questions := []*survey.Question{
		{
			Name: "pat",
			Prompt: &survey.Password{
				Message: "What is your GitHub personal access token?",
				Help:    "See more information in the README.md for solar",
			},
			Validate: survey.Required,
		},
		{
			Name: "gbcap",
			Prompt: &survey.Input{
				Message: "What is the max number amount of storage you want to be used locally?",
			},
			Validate: survey.Required,
		},
	}
	answers := QuestionsData{}
	err := survey.Ask(questions, &answers)
	if err != nil {
		return QuestionsData{}, err
	}
	return answers, nil
}
