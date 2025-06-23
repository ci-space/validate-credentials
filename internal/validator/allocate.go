package validator

import "github.com/ci-space/validate-github-credentials/internal/credentials"

const (
	envGithubToken = "INPUT_GITHUB_TOKEN"
)

func Allocate(store credentials.Store) map[string]Validator {
	validators := map[string]Validator{}

	appendGithubValidator(store, validators)

	return validators
}

func appendGithubValidator(store credentials.Store, validators map[string]Validator) {
	ghToken, ok := store.Get(envGithubToken)
	if ok {
		validators["github"] = NewGithub(ghToken)
	}
}
