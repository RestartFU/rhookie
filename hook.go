package rhookie

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
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

// EditMessage edits a message sent by the webhook.
func (h Hook) EditMessage(id string, payload Payload) error {
	return h.doRequest(http.MethodPatch, "message/"+id, payload)
}

// doRequest sends a request to the webhook.
func (h Hook) doRequest(method, extraEndpoint string, payload Payload) error {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, fmt.Sprintf(baseURL, h.id, h.token)+extraEndpoint, &buf)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	_, err = http.DefaultClient.Do(req)
	return err
}

// NewHook creates a new webhook.
func NewHook(id, token string) Hook {
	return Hook{
		id:    id,
		token: token,
	}
}
