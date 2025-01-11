package comp

import "github.com/zrcoder/amisgo/model"

// iconChecked represents an icon checked renderer
type iconChecked model.Schema

// IconChecked creates a new IconChecked instance
func IconChecked() iconChecked {
	return make(iconChecked)
}

func (i iconChecked) set(key string, value any) iconChecked {
	i[key] = value
	return i
}

// ID sets the unique id of the component
func (i iconChecked) ID(value string) iconChecked {
	return i.set("id", value)
}

// Name sets the name of the component
func (i iconChecked) Name(value string) iconChecked {
	return i.set("name", value)
}

// SVG sets the SVG icon
func (i iconChecked) SVG(value string) iconChecked {
	return i.set("svg", value)
}
