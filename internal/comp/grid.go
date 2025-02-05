package comp

import "github.com/zrcoder/amisgo/model"

// Grid represents a Grid layout renderer
type Grid model.Schema

func NewGrid() Grid {
	return Grid{"type": "grid"}
}

func (g Grid) set(key string, value any) Grid {
	g[key] = value
	return g
}

// Align sets horizontal alignment: 'left', 'right', 'between', 'center'
func (g Grid) Align(value string) Grid {
	return g.set("align", value)
}

// ClassName sets the CSS class name for the container
func (g Grid) ClassName(value string) Grid {
	return g.set("className", value)
}

// Columns sets the columns collection
func (g Grid) Columns(value ...any) Grid {
	return g.set("columns", value)
}

// Disabled sets whether the grid is disabled
func (g Grid) Disabled(value bool) Grid {
	return g.set("disabled", value)
}

// DisabledOn sets the expression to determine if the grid is disabled
func (g Grid) DisabledOn(value string) Grid {
	return g.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (g Grid) EditorSetting(value string) Grid {
	return g.set("editorSetting", value)
}

// Gap sets the horizontal gap: 'xs', 'sm', 'base', 'none', 'md', 'lg'
func (g Grid) Gap(value string) Grid {
	return g.set("gap", value)
}

// Hidden sets whether the grid is hidden
func (g Grid) Hidden(value bool) Grid {
	return g.set("hidden", value)
}

// HiddenOn sets the expression to determine if the grid is hidden
func (g Grid) HiddenOn(value string) Grid {
	return g.set("hiddenOn", value)
}

// ID sets the unique ID for the component
func (g Grid) ID(value string) Grid {
	return g.set("id", value)
}

// OnEvent sets the event action configuration
func (g Grid) OnEvent(value any) Grid {
	return g.set("onEvent", value)
}

// Static sets whether the grid is displayed statically
func (g Grid) Static(value bool) Grid {
	return g.set("static", value)
}

// StaticClassName sets the CSS class name for static display form items
func (g Grid) StaticClassName(value string) Grid {
	return g.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static display form item values
func (g Grid) StaticInputClassName(value string) Grid {
	return g.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static display form item labels
func (g Grid) StaticLabelClassName(value string) Grid {
	return g.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the grid is displayed statically
func (g Grid) StaticOn(value string) Grid {
	return g.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display when the value is empty
func (g Grid) StaticPlaceholder(value string) Grid {
	return g.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display mode
func (g Grid) StaticSchema(value string) Grid {
	return g.set("staticSchema", value)
}

// Style sets the component style
func (g Grid) Style(value any) Grid {
	return g.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (g Grid) TestIdBuilder(value string) Grid {
	return g.set("testIdBuilder", value)
}

// Testid sets the test ID
func (g Grid) Testid(value string) Grid {
	return g.set("testid", value)
}

// UseMobileUI sets whether to disable mobile UI styles at the component level
func (g Grid) UseMobileUI(value bool) Grid {
	return g.set("useMobileUI", value)
}

// Valign sets vertical alignment: 'top', 'middle', 'bottom', 'between'
func (g Grid) Valign(value string) Grid {
	return g.set("valign", value)
}

// Visible sets whether the grid is visible
func (g Grid) Visible(value bool) Grid {
	return g.set("visible", value)
}

// VisibleOn sets the expression to determine if the grid is visible
func (g Grid) VisibleOn(value string) Grid {
	return g.set("visibleOn", value)
}
