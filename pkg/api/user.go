package api

import (
	"context"

	"github.com/shurcooL/githubv4"
)

// Get the number of stars from a user's account
func StarCount(github *githubv4.Client) (int, error) {
	var query struct {
		Viewer struct {
			StarredRepositories struct {
				TotalCount githubv4.Int
			}
		}
	}
	err := github.Query(context.Background(), &query, nil)
	if err != nil {
		return 0, err
	}
	return int(query.Viewer.StarredRepositories.TotalCount), nil
}
