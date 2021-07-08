package search

import (
	"encoding/json"
	"os"

	"github.com/gleich/solar/pkg/api"
)

// Convert list of stars to JSON and write that JSON to a file called repos.json
func WriteJSON(stars []api.Star) error {
	starsJSON, err := json.Marshal(stars)
	if err != nil {
		return err
	}
	err = os.WriteFile("repos.json", starsJSON, 0655)
	if err != nil {
		return err
	}
	return nil
}
