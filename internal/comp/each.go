package comp

import "github.com/zrcoder/amisgo/schema"

// Each represents a loop renderer.
type Each schema.Schema

func NewEach() Each {
	return Each{"type": "each"}
}

func (e Each) set(key string, value any) Each {
	e[key] = value
	return e
}

// ClassName sets the CSS class name.
func (e Each) ClassName(value string) Each {
	return e.set("className", value)
}

// Disabled sets whether the component is disabled.
func (e Each) Disabled(value bool) Each {
	return e.set("disabled", value)
}

// DisabledOn sets the expression to determine if the component is disabled.
func (e Each) DisabledOn(value string) Each {
	return e.set("disabledOn", value)
}

// EditorSetting sets the editor configuration, ignored at runtime.
func (e Each) EditorSetting(value string) Each {
	return e.set("editorSetting", value)
}

// Hidden sets whether the component is hidden.
func (e Each) Hidden(value bool) Each {
	return e.set("hidden", value)
}

// HiddenOn sets the expression to determine if the component is hidden.
func (e Each) HiddenOn(value string) Each {
	return e.set("hiddenOn", value)
}

// ID sets the unique ID for the component, mainly for logging.
func (e Each) ID(value string) Each {
	return e.set("id", value)
}

// IndexKeyName sets the field name to read the index.
func (e Each) IndexKeyName(value string) Each {
	return e.set("indexKeyName", value)
}

// ItemKeyName sets the field name to read the item data.
func (e Each) ItemKeyName(value string) Each {
	return e.set("itemKeyName", value)
}

// Items sets the data items.
func (e Each) Items(value ...any) Each {
	return e.set("items", value)
}

// Name sets the associated field name.
func (e Each) Name(value string) Each {
	return e.set("name", value)
}

// OnEvent sets the event action configuration.
func (e Each) OnEvent(value Event) Each {
	return e.set("onEvent", value)
}

// Placeholder sets the placeholder text.
func (e Each) Placeholder(value string) Each {
	return e.set("placeholder", value)
}

// Source sets the associated field name, supporting data mapping.
func (e Each) Source(value string) Each {
	return e.set("source", value)
}

// Static sets whether the component is displayed statically.
func (e Each) Static(value bool) Each {
	return e.set("static", value)
}

// StaticClassName sets the CSS class name for static display.
func (e Each) StaticClassName(value string) Each {
	return e.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display.
func (e Each) StaticInputClassName(value string) Each {
	return e.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display.
func (e Each) StaticLabelClassName(value string) Each {
	return e.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the component is displayed statically.
func (e Each) StaticOn(value string) Each {
	return e.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display when the value is empty.
func (e Each) StaticPlaceholder(value string) Each {
	return e.set("staticPlaceholder", value)
}

// StaticSchema sets the schema.Schema for static display.
func (e Each) StaticSchema(value string) Each {
	return e.set("staticSchema", value)
}

// Style sets the component style.
func (e Each) Style(value any) Each {
	return e.set("style", value)
}

// TestIdBuilder sets the test ID builder function.
func (e Each) TestIdBuilder(value string) Each {
	return e.set("testIdBuilder", value)
}

// Testid sets the test ID.
func (e Each) Testid(value string) Each {
	return e.set("testid", value)
}

// UseMobileUI sets whether to disable mobile UI styles.
func (e Each) UseMobileUI(value bool) Each {
	return e.set("useMobileUI", value)
}

// Visible sets whether the component is visible.
func (e Each) Visible(value bool) Each {
	return e.set("visible", value)
}

// VisibleOn sets the expression to determine if the component is visible.
func (e Each) VisibleOn(value string) Each {
	return e.set("visibleOn", value)
}
