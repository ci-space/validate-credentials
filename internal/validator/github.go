package validator

import (
	"context"
	"net/http"
	"time"

	"github.com/google/go-github/v67/github"
)

const githubTimeout = 30 * time.Second

type GitHub struct {
	client *github.Client
}

func NewGithub(token string) *GitHub {
	return &GitHub{
		client: github.NewClient(nil).WithAuthToken(token),
	}
}

func (gh *GitHub) Validate(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, githubTimeout)
	defer cancel()

	_, resp, err := gh.client.Emojis.List(ctx)
	if err == nil {
		return nil
	}

	if resp != nil && resp.StatusCode == http.StatusUnauthorized {
		return &InvalidCredentialsError{
			message: "provided invalid credentials",
		}
	}

	return err
}
