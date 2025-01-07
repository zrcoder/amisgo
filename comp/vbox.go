package comp

// vBox represents a vertical layout component

type vBox Schema

// VBox creates a new VBox instance
func VBox() vBox {
	return vBox{}.set("type", "vbox")
}

func (v vBox) set(key string, value any) vBox {
	v[key] = value
	return v
}

// ClassName sets the CSS class name
func (v vBox) ClassName(value string) vBox {
	return v.set("className", value)
}

// Disabled sets whether the component is disabled
func (v vBox) Disabled(value bool) vBox {
	return v.set("disabled", value)
}

// DisabledOn sets the expression to determine if the component is disabled
func (v vBox) DisabledOn(value string) vBox {
	return v.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (v vBox) EditorSetting(value string) vBox {
	return v.set("editorSetting", value)
}

// Hidden sets whether the component is hidden
func (v vBox) Hidden(value bool) vBox {
	return v.set("hidden", value)
}

// HiddenOn sets the expression to determine if the component is hidden
func (v vBox) HiddenOn(value string) vBox {
	return v.set("hiddenOn", value)
}

// ID sets the unique component ID
func (v vBox) ID(value string) vBox {
	return v.set("id", value)
}

// OnEvent sets the event action configuration
func (v vBox) OnEvent(value any) vBox {
	return v.set("onEvent", value)
}

// Rows sets the rows collection
func (v vBox) Rows(value string) vBox {
	return v.set("rows", value)
}

// Static sets whether the component is static
func (v vBox) Static(value bool) vBox {
	return v.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (v vBox) StaticClassName(value string) vBox {
	return v.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (v vBox) StaticInputClassName(value string) vBox {
	return v.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (v vBox) StaticLabelClassName(value string) vBox {
	return v.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the component is static
func (v vBox) StaticOn(value string) vBox {
	return v.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (v vBox) StaticPlaceholder(value string) vBox {
	return v.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema
func (v vBox) StaticSchema(value string) vBox {
	return v.set("staticSchema", value)
}

// Style sets the component style
func (v vBox) Style(value any) vBox {
	return v.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (v vBox) TestIdBuilder(value string) vBox {
	return v.set("testIdBuilder", value)
}

// Testid sets the test ID
func (v vBox) Testid(value string) vBox {
	return v.set("testid", value)
}

// UseMobileUI sets whether to use mobile UI
func (v vBox) UseMobileUI(value bool) vBox {
	return v.set("useMobileUI", value)
}

// Visible sets whether the component is visible
func (v vBox) Visible(value bool) vBox {
	return v.set("visible", value)
}

// VisibleOn sets the expression to determine if the component is visible
func (v vBox) VisibleOn(value string) vBox {
	return v.set("visibleOn", value)
}
