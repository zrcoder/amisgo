package comp

// grid represents a grid layout renderer
type grid Schema

// Grid creates a new Grid instance
func Grid() grid {
	return make(grid).set("type", "grid")
}

func (g grid) set(key string, value any) grid {
	g[key] = value
	return g
}

// Align sets horizontal alignment: 'left', 'right', 'between', 'center'
func (g grid) Align(value string) grid {
	return g.set("align", value)
}

// ClassName sets the CSS class name for the container
func (g grid) ClassName(value string) grid {
	return g.set("className", value)
}

// Columns sets the columns collection
func (g grid) Columns(value ...any) grid {
	return g.set("columns", value)
}

// Disabled sets whether the grid is disabled
func (g grid) Disabled(value bool) grid {
	return g.set("disabled", value)
}

// DisabledOn sets the expression to determine if the grid is disabled
func (g grid) DisabledOn(value string) grid {
	return g.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (g grid) EditorSetting(value string) grid {
	return g.set("editorSetting", value)
}

// Gap sets the horizontal gap: 'xs', 'sm', 'base', 'none', 'md', 'lg'
func (g grid) Gap(value string) grid {
	return g.set("gap", value)
}

// Hidden sets whether the grid is hidden
func (g grid) Hidden(value bool) grid {
	return g.set("hidden", value)
}

// HiddenOn sets the expression to determine if the grid is hidden
func (g grid) HiddenOn(value string) grid {
	return g.set("hiddenOn", value)
}

// ID sets the unique ID for the component
func (g grid) ID(value string) grid {
	return g.set("id", value)
}

// OnEvent sets the event action configuration
func (g grid) OnEvent(value any) grid {
	return g.set("onEvent", value)
}

// Static sets whether the grid is displayed statically
func (g grid) Static(value bool) grid {
	return g.set("static", value)
}

// StaticClassName sets the CSS class name for static display form items
func (g grid) StaticClassName(value string) grid {
	return g.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static display form item values
func (g grid) StaticInputClassName(value string) grid {
	return g.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static display form item labels
func (g grid) StaticLabelClassName(value string) grid {
	return g.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the grid is displayed statically
func (g grid) StaticOn(value string) grid {
	return g.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display when the value is empty
func (g grid) StaticPlaceholder(value string) grid {
	return g.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display mode
func (g grid) StaticSchema(value string) grid {
	return g.set("staticSchema", value)
}

// Style sets the component style
func (g grid) Style(value any) grid {
	return g.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (g grid) TestIdBuilder(value string) grid {
	return g.set("testIdBuilder", value)
}

// Testid sets the test ID
func (g grid) Testid(value string) grid {
	return g.set("testid", value)
}

// UseMobileUI sets whether to disable mobile UI styles at the component level
func (g grid) UseMobileUI(value bool) grid {
	return g.set("useMobileUI", value)
}

// Valign sets vertical alignment: 'top', 'middle', 'bottom', 'between'
func (g grid) Valign(value string) grid {
	return g.set("valign", value)
}

// Visible sets whether the grid is visible
func (g grid) Visible(value bool) grid {
	return g.set("visible", value)
}

// VisibleOn sets the expression to determine if the grid is visible
func (g grid) VisibleOn(value string) grid {
	return g.set("visibleOn", value)
}
