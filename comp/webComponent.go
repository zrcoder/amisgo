package comp

import "github.com/zrcoder/amisgo/model"

// webComponent represents a Web Component

type webComponent model.Schema

// WebComponent creates a new WebComponent instance
func WebComponent() webComponent {
	return webComponent{}.set("type", "web-component")
}

func (wc webComponent) set(key string, value any) webComponent {
	wc[key] = value
	return wc
}

// Body sets child nodes
func (wc webComponent) Body(value ...any) webComponent {
	return wc.set("body", value)
}

// Props sets attributes on the tag
func (wc webComponent) Props(value string) webComponent {
	return wc.set("props", value)
}

// Tag sets the specific web-component tag
func (wc webComponent) Tag(value string) webComponent {
	return wc.set("tag", value)
}
