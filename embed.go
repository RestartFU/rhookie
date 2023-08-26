package rhookie

// Field is a field in an embed.
type Field struct {
	// Name is the name of the field.
	Name string `json:"name"`
	// Value is the value of the field.
	Value string `json:"value"`
	// Inline specifies if the field should be displayed inline.
	Inline bool `json:"inline,omitempty"`
}

// Embed is an embed sent with a webhook.
type Embed struct {
	// Title is the title of the embed.
	Title string `json:"title,omitempty"`
	// Type is the type of the embed.
	Type string `json:"type,omitempty"`
	// Description is the description of the embed.
	Description string `json:"description,omitempty"`
	// URL is the URL of the embed.
	URL string `json:"url,omitempty"`
	// Color is the color of the embed.
	Color int `json:"color,omitempty"`
	// Footer is the footer of the embed.
	Footer string `json:"footer,omitempty"`
	// Fields are the fields of the embed.
	Fields []Field `json:"fields,omitempty"`
}
