package comp

// operation represents the operation renderer

type operation Schema

// Operation creates a new Operation instance
func Operation() operation {
	return operation{}.set("type", "operation")
}

// set sets a field value
func (o operation) set(key string, value any) operation {
	o[key] = value
	return o
}

// Buttons sets the buttons
func (o operation) Buttons(value string) operation {
	return o.set("buttons", value)
}

// ClassName sets the container CSS class name
func (o operation) ClassName(value string) operation {
	return o.set("className", value)
}

// Disabled sets the disabled state
func (o operation) Disabled(value bool) operation {
	return o.set("disabled", value)
}

// DisabledOn sets the disabled expression
func (o operation) DisabledOn(value string) operation {
	return o.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (o operation) EditorSetting(value string) operation {
	return o.set("editorSetting", value)
}

// Fixed sets the fixed column
func (o operation) Fixed(value string) operation {
	return o.set("fixed", value)
}

// Hidden sets the hidden state
func (o operation) Hidden(value bool) operation {
	return o.set("hidden", value)
}

// HiddenOn sets the hidden expression
func (o operation) HiddenOn(value string) operation {
	return o.set("hiddenOn", value)
}

// ID sets the unique component ID
func (o operation) ID(value string) operation {
	return o.set("id", value)
}

// Label sets the label
func (o operation) Label(value string) operation {
	return o.set("label", value)
}

// OnEvent sets the event action configuration
func (o operation) OnEvent(value any) operation {
	return o.set("onEvent", value)
}

// Placeholder sets the placeholder
func (o operation) Placeholder(value string) operation {
	return o.set("placeholder", value)
}

// Static sets the static display state
func (o operation) Static(value bool) operation {
	return o.set("static", value)
}

// StaticClassName sets the static display form item class name
func (o operation) StaticClassName(value string) operation {
	return o.set("staticClassName", value)
}

// StaticInputClassName sets the static display form item value class name
func (o operation) StaticInputClassName(value string) operation {
	return o.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static display form item label class name
func (o operation) StaticLabelClassName(value string) operation {
	return o.set("staticLabelClassName", value)
}

// StaticOn sets the static display expression
func (o operation) StaticOn(value string) operation {
	return o.set("staticOn", value)
}

// StaticPlaceholder sets the static display placeholder
func (o operation) StaticPlaceholder(value string) operation {
	return o.set("staticPlaceholder", value)
}

// StaticSchema sets the static mode schema
func (o operation) StaticSchema(value string) operation {
	return o.set("staticSchema", value)
}

// Style sets the component style
func (o operation) Style(value any) operation {
	return o.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (o operation) TestIdBuilder(value string) operation {
	return o.set("testIdBuilder", value)
}

// Testid sets the test ID
func (o operation) Testid(value string) operation {
	return o.set("testid", value)
}

// UseMobileUI sets the mobile UI usage state
func (o operation) UseMobileUI(value bool) operation {
	return o.set("useMobileUI", value)
}

// Visible sets the visibility state
func (o operation) Visible(value bool) operation {
	return o.set("visible", value)
}

// VisibleOn sets the visibility expression
func (o operation) VisibleOn(value string) operation {
	return o.set("visibleOn", value)
}
