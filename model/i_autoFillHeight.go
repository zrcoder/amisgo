package model

// AutoFillHeight represents a renderer for automatically filling height
type AutoFillHeight Schema

func NewAutoFillHeight() AutoFillHeight {
	return make(AutoFillHeight)
}

func (a AutoFillHeight) set(key string, value any) AutoFillHeight {
	a[key] = value
	return a
}

// Height sets the component's height
func (a AutoFillHeight) Height(value string) AutoFillHeight {
	return a.set("height", value)
}

// MaxHeight sets the maximum height for the component
func (a AutoFillHeight) MaxHeight(value string) AutoFillHeight {
	return a.set("maxHeight", value)
}
