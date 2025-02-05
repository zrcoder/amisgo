package comp

import "github.com/zrcoder/amisgo/model"

// Grid2d represents a 2D grid layout renderer.
type Grid2d model.Schema

func NewGrid2D() Grid2d {
	return Grid2d{"type": "grid-2d"}
}

func (g Grid2d) set(key string, value any) Grid2d {
	g[key] = value
	return g
}

// ClassName sets the CSS class name for the container.
func (g Grid2d) ClassName(value string) Grid2d {
	return g.set("gridClassName", value)
}

// Cols sets the number of columns, default is 12.
func (g Grid2d) Cols(value string) Grid2d {
	return g.set("cols", value)
}

// Disabled sets whether the grid is disabled.
func (g Grid2d) Disabled(value bool) Grid2d {
	return g.set("disabled", value)
}

// DisabledOn sets the expression to determine if the grid is disabled.
func (g Grid2d) DisabledOn(value string) Grid2d {
	return g.set("disabledOn", value)
}

// EditorSetting sets the editor configuration, ignored at runtime.
func (g Grid2d) EditorSetting(value string) Grid2d {
	return g.set("editorSetting", value)
}

// Gap sets the gap between grid items, default is 0.
func (g Grid2d) Gap(value string) Grid2d {
	return g.set("gap", value)
}

// RowGap sets the vertical gap between grid items.
func (g Grid2d) RowGap(value string) Grid2d {
	return g.set("rowGap", value)
}

// Grids sets the configuration for each grid item.
func (g Grid2d) Grids(value ...model.GridItem) Grid2d {
	return g.set("grids", value)
}

// Hidden sets whether the grid is hidden.
func (g Grid2d) Hidden(value bool) Grid2d {
	return g.set("hidden", value)
}

// HiddenOn sets the expression to determine if the grid is hidden.
func (g Grid2d) HiddenOn(value string) Grid2d {
	return g.set("hiddenOn", value)
}

// ID sets the unique ID for the component, mainly for logging.
func (g Grid2d) ID(value string) Grid2d {
	return g.set("id", value)
}

// OnEvent sets the event action configuration.
func (g Grid2d) OnEvent(value any) Grid2d {
	return g.set("onEvent", value)
}

// RowHeight sets the height of each row, default is 50px.
func (g Grid2d) RowHeight(value string) Grid2d {
	return g.set("rowHeight", value)
}

// Static sets whether the grid is displayed statically.
func (g Grid2d) Static(value bool) Grid2d {
	return g.set("static", value)
}

// StaticClassName sets the CSS class name for static display.
func (g Grid2d) StaticClassName(value string) Grid2d {
	return g.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display.
func (g Grid2d) StaticInputClassName(value string) Grid2d {
	return g.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display.
func (g Grid2d) StaticLabelClassName(value string) Grid2d {
	return g.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the grid is displayed statically.
func (g Grid2d) StaticOn(value string) Grid2d {
	return g.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display.
func (g Grid2d) StaticPlaceholder(value string) Grid2d {
	return g.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display mode.
func (g Grid2d) StaticSchema(value string) Grid2d {
	return g.set("staticSchema", value)
}

// Style sets the style for the component.
func (g Grid2d) Style(value any) Grid2d {
	return g.set("style", value)
}

// TestIdBuilder sets the test ID builder.
func (g Grid2d) TestIdBuilder(value string) Grid2d {
	return g.set("testIdBuilder", value)
}

// Testid sets the test ID.
func (g Grid2d) Testid(value string) Grid2d {
	return g.set("testid", value)
}

// UseMobileUI sets whether to use mobile UI styles.
func (g Grid2d) UseMobileUI(value bool) Grid2d {
	return g.set("useMobileUI", value)
}

// Visible sets whether the grid is visible.
func (g Grid2d) Visible(value bool) Grid2d {
	return g.set("visible", value)
}

// VisibleOn sets the expression to determine if the grid is visible.
func (g Grid2d) VisibleOn(value string) Grid2d {
	return g.set("visibleOn", value)
}

// Width sets the width of the grid container, default is auto.
func (g Grid2d) Width(value string) Grid2d {
	return g.set("width", value)
}
