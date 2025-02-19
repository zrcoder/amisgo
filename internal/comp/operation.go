package comp

import "github.com/zrcoder/amisgo/schema"

// Operation represents the Operation renderer
type Operation schema.Schema

func NewOperation() Operation {
	return Operation{"type": "operation"}
}

// set sets a field value
func (o Operation) set(key string, value any) Operation {
	o[key] = value
	return o
}

// Buttons sets the buttons
func (o Operation) Buttons(value string) Operation {
	return o.set("buttons", value)
}

// ClassName sets the container CSS class name
func (o Operation) ClassName(value string) Operation {
	return o.set("className", value)
}

// Disabled sets the disabled state
func (o Operation) Disabled(value bool) Operation {
	return o.set("disabled", value)
}

// DisabledOn sets the disabled expression
func (o Operation) DisabledOn(value string) Operation {
	return o.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (o Operation) EditorSetting(value string) Operation {
	return o.set("editorSetting", value)
}

// Fixed sets the fixed column
func (o Operation) Fixed(value string) Operation {
	return o.set("fixed", value)
}

// Hidden sets the hidden state
func (o Operation) Hidden(value bool) Operation {
	return o.set("hidden", value)
}

// HiddenOn sets the hidden expression
func (o Operation) HiddenOn(value string) Operation {
	return o.set("hiddenOn", value)
}

// ID sets the unique component ID
func (o Operation) ID(value string) Operation {
	return o.set("id", value)
}

// Label sets the label
func (o Operation) Label(value string) Operation {
	return o.set("label", value)
}

// OnEvent sets the event action configuration
func (o Operation) OnEvent(value any) Operation {
	return o.set("onEvent", value)
}

// Placeholder sets the placeholder
func (o Operation) Placeholder(value string) Operation {
	return o.set("placeholder", value)
}

// Static sets the static display state
func (o Operation) Static(value bool) Operation {
	return o.set("static", value)
}

// StaticClassName sets the static display form item class name
func (o Operation) StaticClassName(value string) Operation {
	return o.set("staticClassName", value)
}

// StaticInputClassName sets the static display form item value class name
func (o Operation) StaticInputClassName(value string) Operation {
	return o.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static display form item label class name
func (o Operation) StaticLabelClassName(value string) Operation {
	return o.set("staticLabelClassName", value)
}

// StaticOn sets the static display expression
func (o Operation) StaticOn(value string) Operation {
	return o.set("staticOn", value)
}

// StaticPlaceholder sets the static display placeholder
func (o Operation) StaticPlaceholder(value string) Operation {
	return o.set("staticPlaceholder", value)
}

// StaticSchema sets the static mode schema.Schema
func (o Operation) StaticSchema(value string) Operation {
	return o.set("staticSchema", value)
}

// Style sets the component style
func (o Operation) Style(value any) Operation {
	return o.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (o Operation) TestIdBuilder(value string) Operation {
	return o.set("testIdBuilder", value)
}

// Testid sets the test ID
func (o Operation) Testid(value string) Operation {
	return o.set("testid", value)
}

// UseMobileUI sets the mobile UI usage state
func (o Operation) UseMobileUI(value bool) Operation {
	return o.set("useMobileUI", value)
}

// Visible sets the visibility state
func (o Operation) Visible(value bool) Operation {
	return o.set("visible", value)
}

// VisibleOn sets the visibility expression
func (o Operation) VisibleOn(value string) Operation {
	return o.set("visibleOn", value)
}
