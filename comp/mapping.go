package comp

import "github.com/zrcoder/amisgo/model"

// mapping represents a display component

type mapping model.Schema

// Mapping creates a new Mapping instance
func Mapping() mapping {
	return mapping{}.set("type", "map")
}

func (m mapping) set(key string, value any) mapping {
	m[key] = value
	return m
}

// ClassName sets the CSS class name
func (m mapping) ClassName(value string) mapping {
	m.set("className", value)
	return m
}

// Disabled sets the disabled state
func (m mapping) Disabled(value bool) mapping {
	m.set("disabled", value)
	return m
}

// DisabledOn sets the disabled expression
func (m mapping) DisabledOn(value string) mapping {
	m.set("disabledOn", value)
	return m
}

// EditorSetting sets the editor configuration
func (m mapping) EditorSetting(value string) mapping {
	m.set("editorSetting", value)
	return m
}

// Hidden sets the hidden state
func (m mapping) Hidden(value bool) mapping {
	m.set("hidden", value)
	return m
}

// HiddenOn sets the hidden expression
func (m mapping) HiddenOn(value string) mapping {
	m.set("hiddenOn", value)
	return m
}

// ID sets the unique component ID
func (m mapping) ID(value string) mapping {
	m.set("id", value)
	return m
}

// ItemSchema sets the custom render mapping value
func (m mapping) ItemSchema(value string) mapping {
	m.set("itemSchema", value)
	return m
}

// LabelField sets the label field name
func (m mapping) LabelField(value string) mapping {
	m.set("labelField", value)
	return m
}

// Map sets the mapping rules
func (m mapping) Map(value string) mapping {
	m.set("map", value)
	return m
}

// Name sets the associated field name
func (m mapping) Name(value string) mapping {
	m.set("name", value)
	return m
}

// OnEvent sets the event action configuration
func (m mapping) OnEvent(value any) mapping {
	m.set("onEvent", value)
	return m
}

// Placeholder sets the placeholder text
func (m mapping) Placeholder(value string) mapping {
	m.set("placeholder", value)
	return m
}

// Source sets the source for remote dictionary
func (m mapping) Source(value string) mapping {
	m.set("source", value)
	return m
}

// Static sets the static display state
func (m mapping) Static(value bool) mapping {
	m.set("static", value)
	return m
}

// StaticClassName sets the static display class name
func (m mapping) StaticClassName(value string) mapping {
	m.set("staticClassName", value)
	return m
}

// StaticInputClassName sets the static input class name
func (m mapping) StaticInputClassName(value string) mapping {
	m.set("staticInputClassName", value)
	return m
}

// StaticLabelClassName sets the static label class name
func (m mapping) StaticLabelClassName(value string) mapping {
	m.set("staticLabelClassName", value)
	return m
}

// StaticOn sets the static display expression
func (m mapping) StaticOn(value string) mapping {
	m.set("staticOn", value)
	return m
}

// StaticPlaceholder sets the static placeholder text
func (m mapping) StaticPlaceholder(value string) mapping {
	m.set("staticPlaceholder", value)
	return m
}

// StaticSchema sets the static schema
func (m mapping) StaticSchema(value string) mapping {
	m.set("staticSchema", value)
	return m
}

// Style sets the component style
func (m mapping) Style(value any) mapping {
	m.set("style", value)
	return m
}

// TestIdBuilder sets the test ID builder
func (m mapping) TestIdBuilder(value string) mapping {
	m.set("testIdBuilder", value)
	return m
}

// Testid sets the test ID
func (m mapping) Testid(value string) mapping {
	m.set("testid", value)
	return m
}

// UseMobileUI sets the mobile UI usage
func (m mapping) UseMobileUI(value bool) mapping {
	m.set("useMobileUI", value)
	return m
}

// ValueField sets the value field name
func (m mapping) ValueField(value string) mapping {
	m.set("valueField", value)
	return m
}

// Visible sets the visibility state
func (m mapping) Visible(value bool) mapping {
	m.set("visible", value)
	return m
}

// VisibleOn sets the visibility expression
func (m mapping) VisibleOn(value string) mapping {
	m.set("visibleOn", value)
	return m
}
