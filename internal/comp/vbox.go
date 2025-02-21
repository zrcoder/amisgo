package comp

import "github.com/zrcoder/amisgo/schema"

// VBox represents a vertical layout component
type VBox schema.Schema

func NewVBox() VBox {
	return VBox{"type": "vbox"}
}

func (v VBox) set(key string, value any) VBox {
	v[key] = value
	return v
}

// ClassName sets the CSS class name
func (v VBox) ClassName(value string) VBox {
	return v.set("className", value)
}

// Disabled sets whether the component is disabled
func (v VBox) Disabled(value bool) VBox {
	return v.set("disabled", value)
}

// DisabledOn sets the expression to determine if the component is disabled
func (v VBox) DisabledOn(value string) VBox {
	return v.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (v VBox) EditorSetting(value string) VBox {
	return v.set("editorSetting", value)
}

// Hidden sets whether the component is hidden
func (v VBox) Hidden(value bool) VBox {
	return v.set("hidden", value)
}

// HiddenOn sets the expression to determine if the component is hidden
func (v VBox) HiddenOn(value string) VBox {
	return v.set("hiddenOn", value)
}

// ID sets the unique component ID
func (v VBox) ID(value string) VBox {
	return v.set("id", value)
}

// OnEvent sets the event action configuration
func (v VBox) OnEvent(value Event) VBox {
	return v.set("onEvent", value)
}

// Rows sets the rows collection
func (v VBox) Rows(value string) VBox {
	return v.set("rows", value)
}

// Static sets whether the component is static
func (v VBox) Static(value bool) VBox {
	return v.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (v VBox) StaticClassName(value string) VBox {
	return v.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (v VBox) StaticInputClassName(value string) VBox {
	return v.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (v VBox) StaticLabelClassName(value string) VBox {
	return v.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the component is static
func (v VBox) StaticOn(value string) VBox {
	return v.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (v VBox) StaticPlaceholder(value string) VBox {
	return v.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema.Schema
func (v VBox) StaticSchema(value string) VBox {
	return v.set("staticSchema", value)
}

// Style sets the component style
func (v VBox) Style(value any) VBox {
	return v.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (v VBox) TestIdBuilder(value string) VBox {
	return v.set("testIdBuilder", value)
}

// Testid sets the test ID
func (v VBox) Testid(value string) VBox {
	return v.set("testid", value)
}

// UseMobileUI sets whether to use mobile UI
func (v VBox) UseMobileUI(value bool) VBox {
	return v.set("useMobileUI", value)
}

// Visible sets whether the component is visible
func (v VBox) Visible(value bool) VBox {
	return v.set("visible", value)
}

// VisibleOn sets the expression to determine if the component is visible
func (v VBox) VisibleOn(value string) VBox {
	return v.set("visibleOn", value)
}
