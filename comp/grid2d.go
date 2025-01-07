package comp

// grid2d represents a 2D grid layout renderer.
type grid2d Schema

// Grid2D creates a new Grid2D instance.
func Grid2D() grid2d {
	return make(grid2d).set("type", "grid-2d")
}

func (g grid2d) set(key string, value any) grid2d {
	g[key] = value
	return g
}

// ClassName sets the CSS class name for the container.
func (g grid2d) ClassName(value string) grid2d {
	return g.set("gridClassName", value)
}

// Cols sets the number of columns, default is 12.
func (g grid2d) Cols(value string) grid2d {
	return g.set("cols", value)
}

// Disabled sets whether the grid is disabled.
func (g grid2d) Disabled(value bool) grid2d {
	return g.set("disabled", value)
}

// DisabledOn sets the expression to determine if the grid is disabled.
func (g grid2d) DisabledOn(value string) grid2d {
	return g.set("disabledOn", value)
}

// EditorSetting sets the editor configuration, ignored at runtime.
func (g grid2d) EditorSetting(value string) grid2d {
	return g.set("editorSetting", value)
}

// Gap sets the gap between grid items, default is 0.
func (g grid2d) Gap(value string) grid2d {
	return g.set("gap", value)
}

// RowGap sets the vertical gap between grid items.
func (g grid2d) RowGap(value string) grid2d {
	return g.set("rowGap", value)
}

// Grids sets the configuration for each grid item.
func (g grid2d) Grids(value ...MGridItem) grid2d {
	return g.set("grids", value)
}

// Hidden sets whether the grid is hidden.
func (g grid2d) Hidden(value bool) grid2d {
	return g.set("hidden", value)
}

// HiddenOn sets the expression to determine if the grid is hidden.
func (g grid2d) HiddenOn(value string) grid2d {
	return g.set("hiddenOn", value)
}

// ID sets the unique ID for the component, mainly for logging.
func (g grid2d) ID(value string) grid2d {
	return g.set("id", value)
}

// OnEvent sets the event action configuration.
func (g grid2d) OnEvent(value any) grid2d {
	return g.set("onEvent", value)
}

// RowHeight sets the height of each row, default is 50px.
func (g grid2d) RowHeight(value string) grid2d {
	return g.set("rowHeight", value)
}

// Static sets whether the grid is displayed statically.
func (g grid2d) Static(value bool) grid2d {
	return g.set("static", value)
}

// StaticClassName sets the CSS class name for static display.
func (g grid2d) StaticClassName(value string) grid2d {
	return g.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display.
func (g grid2d) StaticInputClassName(value string) grid2d {
	return g.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display.
func (g grid2d) StaticLabelClassName(value string) grid2d {
	return g.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the grid is displayed statically.
func (g grid2d) StaticOn(value string) grid2d {
	return g.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display.
func (g grid2d) StaticPlaceholder(value string) grid2d {
	return g.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display mode.
func (g grid2d) StaticSchema(value string) grid2d {
	return g.set("staticSchema", value)
}

// Style sets the style for the component.
func (g grid2d) Style(value any) grid2d {
	return g.set("style", value)
}

// TestIdBuilder sets the test ID builder.
func (g grid2d) TestIdBuilder(value string) grid2d {
	return g.set("testIdBuilder", value)
}

// Testid sets the test ID.
func (g grid2d) Testid(value string) grid2d {
	return g.set("testid", value)
}

// UseMobileUI sets whether to use mobile UI styles.
func (g grid2d) UseMobileUI(value bool) grid2d {
	return g.set("useMobileUI", value)
}

// Visible sets whether the grid is visible.
func (g grid2d) Visible(value bool) grid2d {
	return g.set("visible", value)
}

// VisibleOn sets the expression to determine if the grid is visible.
func (g grid2d) VisibleOn(value string) grid2d {
	return g.set("visibleOn", value)
}

// Width sets the width of the grid container, default is auto.
func (g grid2d) Width(value string) grid2d {
	return g.set("width", value)
}
