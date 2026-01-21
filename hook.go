package rhookie

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	// baseURL is the base URL of the discord webhook API.
	baseURL = "https://discord.com/api/webhooks/%s/%s"
	// noExtraEndpoint is an endpoint that doesn't have any extra path.
	noExtraEndpoint = ""
)

// Hook is a Discord webhook.
type Hook struct {
	id, token string
}

// SendMessage sends a message to the webhook.
func (h Hook) SendMessage(payload Payload) error {
	return h.doRequest(http.MethodPost, noExtraEndpoint, payload)
}

// SendMessageWithResponse sends a message to the webhook and returns the created message.
func (h Hook) SendMessageWithResponse(payload Payload) (Message, error) {
	return h.doRequestWithResponse(http.MethodPost, "?wait=true", payload)
}

// EditMessage edits a message sent by the webhook.
func (h Hook) EditMessage(id string, payload Payload) error {
	return h.doRequest(http.MethodPatch, "/messages/"+id, payload)
}

// doRequest sends a request to the webhook.
func (h Hook) doRequest(method, extraEndpoint string, payload Payload) error {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, h.buildURL(extraEndpoint), &buf)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("webhook request failed: %s: %s", resp.Status, strings.TrimSpace(string(body)))
	}
	return nil
}

func (h Hook) doRequestWithResponse(method, extraEndpoint string, payload Payload) (Message, error) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(payload)
	if err != nil {
		return Message{}, err
	}

	req, err := http.NewRequest(method, h.buildURL(extraEndpoint), &buf)
	if err != nil {
		return Message{}, err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Message{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Message{}, err
	}
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return Message{}, fmt.Errorf("webhook request failed: %s: %s", resp.Status, strings.TrimSpace(string(body)))
	}

	var message Message
	if err := json.Unmarshal(body, &message); err != nil {
		return Message{}, err
	}
	return message, nil
}

func (h Hook) buildURL(extraEndpoint string) string {
	base := fmt.Sprintf(baseURL, h.id, h.token)
	if extraEndpoint == "" {
		return base
	}
	if strings.HasPrefix(extraEndpoint, "?") {
		return base + extraEndpoint
	}
	return base + extraEndpoint
}

// NewHook creates a new webhook.
func NewHook(id, token string) Hook {
	return Hook{
		id:    id,
		token: token,
	}
}
