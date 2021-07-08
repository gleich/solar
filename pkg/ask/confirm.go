package ask

import (
	"os"

	"github.com/AlecAivazis/survey/v2"
)

// Confirm something with a user and exit with a status of 0 if they say no
func ConfirmOrExit(prompt survey.Prompt) error {
	// Asking the user
	dontExit := false
	err := survey.AskOne(prompt, &dontExit)
	if err != nil {
		return err
	}

	// Exiting if needed
	if !dontExit {
		os.Exit(0)
	}
	return nil
}
