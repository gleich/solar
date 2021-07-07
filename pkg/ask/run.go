package ask

import (
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

// Ask the user if they are sure they want to run the program.
// Will exit with a status of 0 if the user says they don't want to run.
func ConfirmRun(again bool) error {
	// Asking the user
	run := false
	prompt := &survey.Confirm{
		Message: "Are you sure that you want to run the program?",
		Help:    "This program will use a large number of resources and should only be done once. Please make sure that you really want to run it before running",
		Default: false,
	}
	if again {
		prompt.Message = strings.TrimSuffix(prompt.Message, "?") + " (just double checking)?"
	}
	err := survey.AskOne(prompt, &run)
	if err != nil {
		return err
	}

	// Exiting if needed
	if !run {
		os.Exit(0)
	}
	return nil
}
