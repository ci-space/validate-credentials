package main

import (
	"context"
	"github.com/ci-space/validate-credentials/internal/credentials"
	"github.com/ci-space/validate-credentials/internal/validator"
	"log/slog"
	"os"
	"strings"
)

func main() {
	ctx := context.Background()

	slog.Info("allocating validators")

	validators := validator.Allocate(credentials.NewEnvStore())
	if len(validators) == 0 {
		fail("no found validators. please check your configuration")
	}

	names := make([]string, len(validators))
	i := 0
	for name := range validators {
		names[i] = name
		i++
	}

	slog.Info("running validators", slog.Any("validators", strings.Join(names, ", ")))

	if !runValidators(ctx, validators) {
		os.Exit(1)
	}
}

func runValidators(ctx context.Context, validators map[string]validator.Validator) bool {
	ok := true

	for name, v := range validators {
		err := v.Validate(ctx)
		if err != nil {
			ok = false

			slog.Error("validation failed", slog.String("validator", name), slog.Any("err", err))

			continue
		}

		slog.Info("validation succeed", slog.String("validator", name))
	}

	return ok
}

func fail(message string) {
	slog.Error(message)
	os.Exit(1)
}
