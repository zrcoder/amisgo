package comp

import "github.com/zrcoder/amisgo/model"

// AutoFillHeight represents a renderer for automatically filling height
type autoFillHeight model.Schema

// AutoFillHeight creates a new AutoFillHeight instance
func AutoFillHeight() autoFillHeight {
	return make(autoFillHeight)
}

func (a autoFillHeight) set(key string, value any) autoFillHeight {
	a[key] = value
	return a
}

// Height sets the component's height
func (a autoFillHeight) Height(value string) autoFillHeight {
	return a.set("height", value)
}

// MaxHeight sets the maximum height for the component
func (a autoFillHeight) MaxHeight(value string) autoFillHeight {
	return a.set("maxHeight", value)
}
