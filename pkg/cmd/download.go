package cmd

import (
	"fmt"

	"code.cloudfoundry.org/bytefmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/gleich/lumber"
	"github.com/gleich/solar/pkg/api"
	"github.com/gleich/solar/pkg/ask"
	"github.com/gleich/solar/pkg/clone"
	"github.com/gleich/solar/pkg/search"
	"github.com/kyokomi/emoji"
	"github.com/spf13/cobra"
)

var downloadCMD = &cobra.Command{
	DisableFlagsInUseLine: true,
	Args:                  cobra.NoArgs,
	Use:                   "download",
	Short:                 "Clone all starred repos",
	Run: func(cmd *cobra.Command, args []string) {
		// Asking the user some questions
		err := ask.ConfirmOrExit(&survey.Confirm{
			Message: "Are you sure that you want to run the program?",
			Help:    "This program will use a large number of resources and should only be done once. Please make sure that you really want to run it before running",
			Default: false,
		})
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
		lumber.Success("Got information for", len(stars), "repos")

		// Getting total number of GB that will be cloned
		var totalKB uint64
		for _, star := range stars {
			if !star.IsEmpty || !star.IsPrivate {
				totalKB += uint64(star.DiskUsage) * bytefmt.KILOBYTE
			}
		}
		err = ask.ConfirmOrExit(&survey.Confirm{
			Message: fmt.Sprintf("Will clone %v. Are you still sure?", bytefmt.ByteSize(totalKB)),
			Default: false,
			Help:    "This amount will actually be less because we are cloning the repo with a depth of 1 instead of the full depth. The size will never be larger than the value provided.",
		})
		if err != nil {
			lumber.Fatal(err, "Failed to confirm size with user")
		}

		if answers.Search {
			err = search.WriteJSON(stars)
			if err != nil {
				lumber.Fatal(err, "Failed to write star data to JSON file")
			}
		}

		clonedSoFar := 0
		for _, star := range stars {
			// Checking if the repo should be skipped
			var (
				ignored       = false
				ignoredReason = ""
			)
			if star.IsEmpty {
				ignored = true
				ignoredReason = "empty"
			}
			if star.IsPrivate {
				ignored = true
				ignoredReason = "private"
			}
			if ignored {
				lumber.Warning(
					star.Name+"/"+star.Owner.Login,
					"won't be cloned because it is",
					ignoredReason,
				)
				continue
			}

			star.Description = emoji.Sprint(star.Description)

			err = clone.Star(star, starCount, clonedSoFar)
			if err != nil {
				lumber.Fatal(err, "Failed to clone", star.Name+"/"+star.Owner.Login)
			}
			clonedSoFar++
		}
		lumber.Success("Cloned", clonedSoFar, "stars. Thank you for using solar!")
	},
}

func init() {
	RootCMD.AddCommand(downloadCMD)
}
