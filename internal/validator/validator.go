package validator

import "context"

type Validator interface {
	Validate(ctx context.Context) error
}

type InvalidCredentialsError struct {
	message string
}

func (e *InvalidCredentialsError) Error() string {
	return e.message
}
