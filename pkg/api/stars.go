package api

import (
	"context"

	"github.com/shurcooL/githubv4"
)

// A starred github repo
type Star struct {
	Name  string
	Owner struct {
		Login string
	}
	Description string
	DiskUsage   githubv4.Int
	IsEmpty     bool
	IsPrivate   bool
	URL         string
}

// Get star data
func Stars(github *githubv4.Client, starCount int) ([]Star, error) {
	stars := []Star{}
	endCursor := ""
	for len(stars) < starCount {
		var query struct {
			Viewer struct {
				StarredRepositories struct {
					Nodes    []Star
					PageInfo struct {
						EndCursor string
					}
				} `graphql:"starredRepositories(after: $endCursor)"`
			}
		}

		err := github.Query(
			context.Background(),
			&query,
			map[string]interface{}{"endCursor": githubv4.String(endCursor)},
		)
		if err != nil {
			return []Star{}, err
		}

		stars = append(stars, query.Viewer.StarredRepositories.Nodes...)
		endCursor = query.Viewer.StarredRepositories.PageInfo.EndCursor
	}
	return stars, nil
}
