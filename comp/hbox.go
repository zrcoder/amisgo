package comp

import "github.com/zrcoder/amisgo/model"

// hBox represents a horizontal layout renderer
type hBox model.Schema

// HBox creates a new hBox instance
func HBox() hBox {
	return hBox{"type": "hbox"}
}

func (h hBox) set(key string, value any) hBox {
	h[key] = value
	return h
}

// Align sets the horizontal alignment: left, right, between, center
func (h hBox) Align(value string) hBox {
	return h.set("align", value)
}

// ClassName sets the container CSS class name
func (h hBox) ClassName(value string) hBox {
	return h.set("className", value)
}

// Columns sets the columns
func (h hBox) Columns(value ...MColumn) hBox {
	return h.set("columns", value)
}

// Disabled sets whether the component is disabled
func (h hBox) Disabled(value bool) hBox {
	return h.set("disabled", value)
}

// DisabledOn sets the expression to determine if the component is disabled
func (h hBox) DisabledOn(value string) hBox {
	return h.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (h hBox) EditorSetting(value string) hBox {
	return h.set("editorSetting", value)
}

// Gap sets the horizontal gap: xs, sm, base, none, md, lg
func (h hBox) Gap(value string) hBox {
	return h.set("gap", value)
}

// Hidden sets whether the component is hidden
func (h hBox) Hidden(value bool) hBox {
	return h.set("hidden", value)
}

// HiddenOn sets the expression to determine if the component is hidden
func (h hBox) HiddenOn(value string) hBox {
	return h.set("hiddenOn", value)
}

// ID sets the unique component ID
func (h hBox) ID(value string) hBox {
	return h.set("id", value)
}

// OnEvent sets the event action configuration
func (h hBox) OnEvent(value any) hBox {
	return h.set("onEvent", value)
}

// Static sets whether the component is statically displayed
func (h hBox) Static(value bool) hBox {
	return h.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (h hBox) StaticClassName(value string) hBox {
	return h.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (h hBox) StaticInputClassName(value string) hBox {
	return h.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (h hBox) StaticLabelClassName(value string) hBox {
	return h.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the component is statically displayed
func (h hBox) StaticOn(value string) hBox {
	return h.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display when value is empty
func (h hBox) StaticPlaceholder(value string) hBox {
	return h.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema
func (h hBox) StaticSchema(value string) hBox {
	return h.set("staticSchema", value)
}

// Style sets the component style
func (h hBox) Style(value any) hBox {
	return h.set("style", value)
}

// SubFormHorizontal sets the horizontal width ratio for sub-forms
func (h hBox) SubFormHorizontal(value string) hBox {
	return h.set("subFormHorizontal", value)
}

// SubFormMode sets the display mode for sub-forms
func (h hBox) SubFormMode(value string) hBox {
	return h.set("subFormMode", value)
}

// TestIdBuilder sets the test ID builder
func (h hBox) TestIdBuilder(value string) hBox {
	return h.set("testIdBuilder", value)
}

// Testid sets the test ID
func (h hBox) Testid(value string) hBox {
	return h.set("testid", value)
}

// UseMobileUI sets whether to disable mobile UI styles
func (h hBox) UseMobileUI(value bool) hBox {
	return h.set("useMobileUI", value)
}

// Valign sets the vertical alignment: top, middle, bottom, between
func (h hBox) Valign(value string) hBox {
	return h.set("valign", value)
}

// Visible sets whether the component is visible
func (h hBox) Visible(value bool) hBox {
	return h.set("visible", value)
}

// VisibleOn sets the expression to determine if the component is visible
func (h hBox) VisibleOn(value string) hBox {
	return h.set("visibleOn", value)
}
