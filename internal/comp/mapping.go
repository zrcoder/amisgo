package comp

import "github.com/zrcoder/amisgo/model"

// Mapping represents a display component
type Mapping model.Schema

func NewMapping() Mapping {
	return Mapping{"type": "map"}
}

func (m Mapping) set(key string, value any) Mapping {
	m[key] = value
	return m
}

// ClassName sets the CSS class name
func (m Mapping) ClassName(value string) Mapping {
	m.set("className", value)
	return m
}

// Disabled sets the disabled state
func (m Mapping) Disabled(value bool) Mapping {
	m.set("disabled", value)
	return m
}

// DisabledOn sets the disabled expression
func (m Mapping) DisabledOn(value string) Mapping {
	m.set("disabledOn", value)
	return m
}

// EditorSetting sets the editor configuration
func (m Mapping) EditorSetting(value string) Mapping {
	m.set("editorSetting", value)
	return m
}

// Hidden sets the hidden state
func (m Mapping) Hidden(value bool) Mapping {
	m.set("hidden", value)
	return m
}

// HiddenOn sets the hidden expression
func (m Mapping) HiddenOn(value string) Mapping {
	m.set("hiddenOn", value)
	return m
}

// ID sets the unique component ID
func (m Mapping) ID(value string) Mapping {
	m.set("id", value)
	return m
}

// ItemSchema sets the custom render mapping value
func (m Mapping) ItemSchema(value string) Mapping {
	m.set("itemSchema", value)
	return m
}

// LabelField sets the label field name
func (m Mapping) LabelField(value string) Mapping {
	m.set("labelField", value)
	return m
}

// Map sets the mapping rules
func (m Mapping) Map(value string) Mapping {
	m.set("map", value)
	return m
}

// Name sets the associated field name
func (m Mapping) Name(value string) Mapping {
	m.set("name", value)
	return m
}

// OnEvent sets the event action configuration
func (m Mapping) OnEvent(value any) Mapping {
	m.set("onEvent", value)
	return m
}

// Placeholder sets the placeholder text
func (m Mapping) Placeholder(value string) Mapping {
	m.set("placeholder", value)
	return m
}

// Source sets the source for remote dictionary
func (m Mapping) Source(value string) Mapping {
	m.set("source", value)
	return m
}

// Static sets the static display state
func (m Mapping) Static(value bool) Mapping {
	m.set("static", value)
	return m
}

// StaticClassName sets the static display class name
func (m Mapping) StaticClassName(value string) Mapping {
	m.set("staticClassName", value)
	return m
}

// StaticInputClassName sets the static input class name
func (m Mapping) StaticInputClassName(value string) Mapping {
	m.set("staticInputClassName", value)
	return m
}

// StaticLabelClassName sets the static label class name
func (m Mapping) StaticLabelClassName(value string) Mapping {
	m.set("staticLabelClassName", value)
	return m
}

// StaticOn sets the static display expression
func (m Mapping) StaticOn(value string) Mapping {
	m.set("staticOn", value)
	return m
}

// StaticPlaceholder sets the static placeholder text
func (m Mapping) StaticPlaceholder(value string) Mapping {
	m.set("staticPlaceholder", value)
	return m
}

// StaticSchema sets the static schema
func (m Mapping) StaticSchema(value string) Mapping {
	m.set("staticSchema", value)
	return m
}

// Style sets the component style
func (m Mapping) Style(value any) Mapping {
	m.set("style", value)
	return m
}

// TestIdBuilder sets the test ID builder
func (m Mapping) TestIdBuilder(value string) Mapping {
	m.set("testIdBuilder", value)
	return m
}

// Testid sets the test ID
func (m Mapping) Testid(value string) Mapping {
	m.set("testid", value)
	return m
}

// UseMobileUI sets the mobile UI usage
func (m Mapping) UseMobileUI(value bool) Mapping {
	m.set("useMobileUI", value)
	return m
}

// ValueField sets the value field name
func (m Mapping) ValueField(value string) Mapping {
	m.set("valueField", value)
	return m
}

// Visible sets the visibility state
func (m Mapping) Visible(value bool) Mapping {
	m.set("visible", value)
	return m
}

// VisibleOn sets the visibility expression
func (m Mapping) VisibleOn(value string) Mapping {
	m.set("visibleOn", value)
	return m
}
