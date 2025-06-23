package validator

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Telegram struct {
	token string

	chatID string
}

func NewTelegram(
	token string,
	chatID string,
) *Telegram {
	return &Telegram{
		token:  token,
		chatID: chatID,
	}
}

func (t *Telegram) Validate(ctx context.Context) error {
	err := t.getMe(ctx)
	if err != nil {
		return err
	}

	if t.chatID != "" {
		return t.getChat(ctx)
	}

	return nil
}

func (t *Telegram) getChat(ctx context.Context) error {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("https://api.telegram.org/bot%s/getChat?chat_id=%s", t.token, t.chatID),
		nil,
	)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("exec request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return &InvalidCredentialsError{
			message: "provided invalid credentials",
		}
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read reesponse body: %w", err)
	}

	respBodyString := string(respBody)
	if strings.Contains(respBodyString, "chat not found") {
		return &InvalidCredentialsError{
			message: "chat not found",
		}
	}

	return fmt.Errorf("telegram returns: %s", respBody)
}

func (t *Telegram) getMe(ctx context.Context) error {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("https://api.telegram.org/bot%s/getMe", t.token),
		nil,
	)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("exec request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return &InvalidCredentialsError{
			message: "provided invalid credentials",
		}
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read reesponse body: %w", err)
	}

	return fmt.Errorf("telegram returns: %s", respBody)
}
