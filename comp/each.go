package comp

import "github.com/zrcoder/amisgo/model"

// each represents a loop renderer.
type each model.Schema

// Each creates a new Each instance with the default type set to 'each'.
func Each() each {
	return make(each).set("type", "each")
}

func (e each) set(key string, value any) each {
	e[key] = value
	return e
}

// ClassName sets the CSS class name.
func (e each) ClassName(value string) each {
	return e.set("className", value)
}

// Disabled sets whether the component is disabled.
func (e each) Disabled(value bool) each {
	return e.set("disabled", value)
}

// DisabledOn sets the expression to determine if the component is disabled.
func (e each) DisabledOn(value string) each {
	return e.set("disabledOn", value)
}

// EditorSetting sets the editor configuration, ignored at runtime.
func (e each) EditorSetting(value string) each {
	return e.set("editorSetting", value)
}

// Hidden sets whether the component is hidden.
func (e each) Hidden(value bool) each {
	return e.set("hidden", value)
}

// HiddenOn sets the expression to determine if the component is hidden.
func (e each) HiddenOn(value string) each {
	return e.set("hiddenOn", value)
}

// ID sets the unique ID for the component, mainly for logging.
func (e each) ID(value string) each {
	return e.set("id", value)
}

// IndexKeyName sets the field name to read the index.
func (e each) IndexKeyName(value string) each {
	return e.set("indexKeyName", value)
}

// ItemKeyName sets the field name to read the item data.
func (e each) ItemKeyName(value string) each {
	return e.set("itemKeyName", value)
}

// Items sets the data items.
func (e each) Items(value ...any) each {
	return e.set("items", value)
}

// Name sets the associated field name.
func (e each) Name(value string) each {
	return e.set("name", value)
}

// OnEvent sets the event action configuration.
func (e each) OnEvent(value any) each {
	return e.set("onEvent", value)
}

// Placeholder sets the placeholder text.
func (e each) Placeholder(value string) each {
	return e.set("placeholder", value)
}

// Source sets the associated field name, supporting data mapping.
func (e each) Source(value string) each {
	return e.set("source", value)
}

// Static sets whether the component is displayed statically.
func (e each) Static(value bool) each {
	return e.set("static", value)
}

// StaticClassName sets the CSS class name for static display.
func (e each) StaticClassName(value string) each {
	return e.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display.
func (e each) StaticInputClassName(value string) each {
	return e.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display.
func (e each) StaticLabelClassName(value string) each {
	return e.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the component is displayed statically.
func (e each) StaticOn(value string) each {
	return e.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display when the value is empty.
func (e each) StaticPlaceholder(value string) each {
	return e.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display.
func (e each) StaticSchema(value string) each {
	return e.set("staticSchema", value)
}

// Style sets the component style.
func (e each) Style(value any) each {
	return e.set("style", value)
}

// TestIdBuilder sets the test ID builder function.
func (e each) TestIdBuilder(value string) each {
	return e.set("testIdBuilder", value)
}

// Testid sets the test ID.
func (e each) Testid(value string) each {
	return e.set("testid", value)
}

// UseMobileUI sets whether to disable mobile UI styles.
func (e each) UseMobileUI(value bool) each {
	return e.set("useMobileUI", value)
}

// Visible sets whether the component is visible.
func (e each) Visible(value bool) each {
	return e.set("visible", value)
}

// VisibleOn sets the expression to determine if the component is visible.
func (e each) VisibleOn(value string) each {
	return e.set("visibleOn", value)
}
