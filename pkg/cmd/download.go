package cmd

import (
	"fmt"

	"github.com/gleich/lumber"
	"github.com/gleich/solar/pkg/api"
	"github.com/gleich/solar/pkg/ask"
	"github.com/spf13/cobra"
)

var downloadCMD = &cobra.Command{
	DisableFlagsInUseLine: true,
	Args:                  cobra.NoArgs,
	Use:                   "download",
	Short:                 "Clone all starred repos",
	Run: func(cmd *cobra.Command, args []string) {
		// Asking the user some questions
		err := ask.ConfirmRun(false)
		if err != nil {
			lumber.Fatal(err, "Failed to confirm run")
		}
		answers, err := ask.Questions()
		if err != nil {
			lumber.Fatal(err, "Failed to ask questions")
		}

		// Getting some information
		ghClient := api.GenClient(answers.PAT)
		starCount, err := api.StarCount(ghClient)
		if err != nil {
			lumber.Fatal(err, "Failed to get number of stars for your account")
		}

		fmt.Println()
		lumber.Info("Getting information for", starCount, "repos")

		// Getting star information
		stars, err := api.Stars(ghClient, starCount)
		if err != nil {
			lumber.Fatal(err, "Failed to get stars from GitHub API")
		}
		fmt.Printf("%#v\n", len(stars))
	},
}

func init() {
	RootCMD.AddCommand(downloadCMD)
}
