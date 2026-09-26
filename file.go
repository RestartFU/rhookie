package rhookie

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

// SendFileWithResponse sends a message with one file attached and returns the created message.
func (h Hook) SendFileWithResponse(ctx context.Context, payload Payload, filename string, data []byte) (Message, error) {
	var body bytes.Buffer
	form := multipart.NewWriter(&body)
	encoded, err := json.Marshal(payload)
	if err != nil {
		return Message{}, err
	}
	if err := form.WriteField("payload_json", string(encoded)); err != nil {
		return Message{}, err
	}
	file, err := form.CreateFormFile("files[0]", filename)
	if err != nil {
		return Message{}, err
	}
	if _, err := file.Write(data); err != nil {
		return Message{}, err
	}
	if err := form.Close(); err != nil {
		return Message{}, err
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, h.buildURL("?wait=true"), &body)
	if err != nil {
		return Message{}, err
	}
	request.Header.Set("Content-Type", form.FormDataContentType())
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return Message{}, err
	}
	defer response.Body.Close()
	responseBody, err := io.ReadAll(io.LimitReader(response.Body, 1<<20))
	if err != nil {
		return Message{}, err
	}
	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
		return Message{}, fmt.Errorf("webhook request failed: %s: %s", response.Status, strings.TrimSpace(string(responseBody)))
	}
	var message Message
	if err := json.Unmarshal(responseBody, &message); err != nil {
		return Message{}, err
	}
	return message, nil
}
