package comp

import "github.com/zrcoder/amisgo/schema"

// HBox represents a horizontal layout renderer
type HBox schema.Schema

func NewHBox() HBox {
	return HBox{"type": "hbox"}
}

func (h HBox) set(key string, value any) HBox {
	h[key] = value
	return h
}

// Align sets the horizontal alignment: left, right, between, center
func (h HBox) Align(value string) HBox {
	return h.set("align", value)
}

// ClassName sets the container CSS class name
func (h HBox) ClassName(value string) HBox {
	return h.set("className", value)
}

// Columns sets the columns
func (h HBox) Columns(value ...Column) HBox {
	return h.set("columns", value)
}

// Disabled sets whether the component is disabled
func (h HBox) Disabled(value bool) HBox {
	return h.set("disabled", value)
}

// DisabledOn sets the expression to determine if the component is disabled
func (h HBox) DisabledOn(value string) HBox {
	return h.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (h HBox) EditorSetting(value string) HBox {
	return h.set("editorSetting", value)
}

// Gap sets the horizontal gap: xs, sm, base, none, md, lg
func (h HBox) Gap(value string) HBox {
	return h.set("gap", value)
}

// Hidden sets whether the component is hidden
func (h HBox) Hidden(value bool) HBox {
	return h.set("hidden", value)
}

// HiddenOn sets the expression to determine if the component is hidden
func (h HBox) HiddenOn(value string) HBox {
	return h.set("hiddenOn", value)
}

// ID sets the unique component ID
func (h HBox) ID(value string) HBox {
	return h.set("id", value)
}

// OnEvent sets the event action configuration
func (h HBox) OnEvent(value Event) HBox {
	return h.set("onEvent", value)
}

// Static sets whether the component is statically displayed
func (h HBox) Static(value bool) HBox {
	return h.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (h HBox) StaticClassName(value string) HBox {
	return h.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (h HBox) StaticInputClassName(value string) HBox {
	return h.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (h HBox) StaticLabelClassName(value string) HBox {
	return h.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the component is statically displayed
func (h HBox) StaticOn(value string) HBox {
	return h.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display when value is empty
func (h HBox) StaticPlaceholder(value string) HBox {
	return h.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema.Schema
func (h HBox) StaticSchema(value string) HBox {
	return h.set("staticSchema", value)
}

// Style sets the component style
func (h HBox) Style(value any) HBox {
	return h.set("style", value)
}

// SubFormHorizontal sets the horizontal width ratio for sub-forms
func (h HBox) SubFormHorizontal(value string) HBox {
	return h.set("subFormHorizontal", value)
}

// SubFormMode sets the display mode for sub-forms
func (h HBox) SubFormMode(value string) HBox {
	return h.set("subFormMode", value)
}

// TestIdBuilder sets the test ID builder
func (h HBox) TestIdBuilder(value string) HBox {
	return h.set("testIdBuilder", value)
}

// Testid sets the test ID
func (h HBox) Testid(value string) HBox {
	return h.set("testid", value)
}

// UseMobileUI sets whether to disable mobile UI styles
func (h HBox) UseMobileUI(value bool) HBox {
	return h.set("useMobileUI", value)
}

// Valign sets the vertical alignment: top, middle, bottom, between
func (h HBox) Valign(value string) HBox {
	return h.set("valign", value)
}

// Visible sets whether the component is visible
func (h HBox) Visible(value bool) HBox {
	return h.set("visible", value)
}

// VisibleOn sets the expression to determine if the component is visible
func (h HBox) VisibleOn(value string) HBox {
	return h.set("visibleOn", value)
}
