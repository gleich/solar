package ask

import (
	"github.com/AlecAivazis/survey/v2"
)

// Data from the questions ask to the user
type QuestionsData struct {
	PAT    string
	Search bool
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
			Name: "search",
			Prompt: &survey.Confirm{
				Message: "Do you want to be able to search over your cloned stars (you can't change this later)",
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
