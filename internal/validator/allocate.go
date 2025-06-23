package validator

import (
	"github.com/ci-space/validate-github-credentials/internal/credentials"
)

const (
	envGithubToken = "INPUT_GITHUB_TOKEN"

	envTelegramToken  = "INPUT_TELEGRAM_TOKEN"
	envTelegramChatID = "INPUT_TELEGRAM_CHAT_ID"
)

func Allocate(store credentials.Store) map[string]Validator {
	validators := map[string]Validator{}

	appendGithubValidator(store, validators)
	appendTelegramValidator(store, validators)

	return validators
}

func appendGithubValidator(store credentials.Store, validators map[string]Validator) {
	ghToken, ok := store.Get(envGithubToken)
	if ok {
		validators["github"] = NewGithub(ghToken)
	}
}

func appendTelegramValidator(store credentials.Store, validators map[string]Validator) {
	tgToken, ok := store.Get(envTelegramToken)
	if !ok {
		return
	}

	tgChatID, _ := store.Get(envTelegramChatID)

	validators["telegram"] = NewTelegram(tgToken, tgChatID)
}
