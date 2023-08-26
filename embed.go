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

// WithName adds a name to the field.
func (f Field) WithName(name string) Field {
	f.Name = name
	return f
}

// WithValue adds a value to the field.
func (f Field) WithValue(value string) Field {
	f.Value = value
	return f
}

// WithInline sets the inline state of the field.
func (f Field) WithInline(inline bool) Field {
	f.Inline = inline
	return f
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

// WithTitle adds a title to the embed.
func (e Embed) WithTitle(title string) Embed {
	e.Title = title
	return e
}

// WithType adds a type to the embed.
func (e Embed) WithType(t string) Embed {
	e.Type = t
	return e
}

// WithDescription adds a description to the embed.
func (e Embed) WithDescription(description string) Embed {
	e.Description = description
	return e
}

// WithURL adds a URL to the embed.
func (e Embed) WithURL(url string) Embed {
	e.URL = url
	return e
}

// WithColor adds a color to the embed.
func (e Embed) WithColor(color int) Embed {
	e.Color = color
	return e
}

// WithFooter adds a footer to the embed.
func (e Embed) WithFooter(footer string) Embed {
	e.Footer = footer
	return e
}

// WithFields adds fields to the embed.
func (e Embed) WithFields(fields ...Field) Embed {
	e.Fields = append(e.Fields, fields...)
	return e
}
