package validator

import (
	"context"
	"github.com/google/go-github/v67/github"
	"net/http"
)

type GitHub struct {
	client *github.Client
}

func NewGithub(token string) *GitHub {
	return &GitHub{
		client: github.NewClient(nil).WithAuthToken(token),
	}
}

func (gh *GitHub) Validate(ctx context.Context) error {
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
