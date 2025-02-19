package comp

import "github.com/zrcoder/amisgo/schema"

// IconChecked represents an icon checked renderer
type IconChecked schema.Schema

func NewIconChecked() IconChecked {
	return make(IconChecked)
}

func (i IconChecked) set(key string, value any) IconChecked {
	i[key] = value
	return i
}

// ID sets the unique id of the component
func (i IconChecked) ID(value string) IconChecked {
	return i.set("id", value)
}

// Name sets the name of the component
func (i IconChecked) Name(value string) IconChecked {
	return i.set("name", value)
}

// SVG sets the SVG icon
func (i IconChecked) SVG(value string) IconChecked {
	return i.set("svg", value)
}
