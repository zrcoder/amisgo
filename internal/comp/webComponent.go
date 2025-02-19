package comp

import "github.com/zrcoder/amisgo/schema"

// WebComponent represents a Web Component
type WebComponent schema.Schema

func NewWebComponent() WebComponent {
	return WebComponent{"type": "web-component"}
}

func (wc WebComponent) set(key string, value any) WebComponent {
	wc[key] = value
	return wc
}

// Body sets child nodes
func (wc WebComponent) Body(value ...any) WebComponent {
	return wc.set("body", value)
}

// Props sets attributes on the tag
func (wc WebComponent) Props(value string) WebComponent {
	return wc.set("props", value)
}

// Tag sets the specific web-component tag
func (wc WebComponent) Tag(value string) WebComponent {
	return wc.set("tag", value)
}
