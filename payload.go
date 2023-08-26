package rhookie

// Payload is the payload sent to the webhook.
type Payload struct {
	// Content is the content of the message.
	Content string `json:"content"`
	// Username is the username of the webhook.
	Username string `json:"username,omitempty"`
	// AvatarURL is the URL of the avatar of the webhook.
	AvatarURL string `json:"avatar_url,omitempty"`
	// TTS specifies if the message should be sent as a TTS message.
	TTS bool `json:"tts,omitempty"`
	// Embeds are the embeds sent with the message.
	Embeds []Embed `json:"embeds"`
}

// WithUsername adds a username to the payload.
func (p Payload) WithUsername(username string) Payload {
	p.Username = username
	return p
}

// WithTTS sets the TTS state of the payload.
func (p Payload) WithTTS(tts bool) Payload {
	p.TTS = tts
	return p
}

// WithAvatarURL adds an avatar URL to the payload.
func (p Payload) WithAvatarURL(url string) Payload {
	p.AvatarURL = url
	return p
}

// WithContent adds content to the payload.
func (p Payload) WithContent(content string) Payload {
	p.Content = content
	return p
}

// WithEmbeds adds embeds to the payload.
func (p Payload) WithEmbeds(embeds ...Embed) Payload {
	p.Embeds = append(p.Embeds, embeds...)
	return p
}
