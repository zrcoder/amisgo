package comp

import "github.com/zrcoder/amisgo/schema"

// CollapseGroup represents a group of collapsible panels with shared behavior
type CollapseGroup schema.Schema

func NewCollapseGroup() CollapseGroup {
	return CollapseGroup{"type": "collapse-group"}
}

func (c CollapseGroup) set(key string, value any) CollapseGroup {
	c[key] = value
	return c
}

// Accordion enables or disables accordion mode (only one panel open at a time)
func (c CollapseGroup) Accordion(value bool) CollapseGroup {
	return c.set("accordion", value)
}

// ActiveKey sets the currently active panel(s), Array<string | number | never> | string | number
func (c CollapseGroup) ActiveKey(value any) CollapseGroup {
	return c.set("activeKey", value)
}

// Body sets the content for the collapse group
func (c CollapseGroup) Body(value ...any) CollapseGroup {
	return c.set("body", value)
}

// ClassName sets the CSS class name for the container
func (c CollapseGroup) ClassName(value string) CollapseGroup {
	return c.set("className", value)
}

// Disabled enables or disables the entire collapse group
func (c CollapseGroup) Disabled(value bool) CollapseGroup {
	return c.set("disabled", value)
}

// DisabledOn sets a conditional expression to dynamically disable the component
func (c CollapseGroup) DisabledOn(value string) CollapseGroup {
	return c.set("disabledOn", value)
}

// EditorSetting configures editor-specific settings for the component
func (c CollapseGroup) EditorSetting(value string) CollapseGroup {
	return c.set("editorSetting", value)
}

// EnableFieldSetStyle enables or disables the FieldSet styling
func (c CollapseGroup) EnableFieldSetStyle(value bool) CollapseGroup {
	return c.set("enableFieldSetStyle", value)
}

// ExpandIcon sets a custom icon for expanding/collapsing panels
func (c CollapseGroup) ExpandIcon(value any) CollapseGroup {
	return c.set("expandIcon", value)
}

// ExpandIconPosition sets the placement of the expand/collapse icon
func (c CollapseGroup) ExpandIconPosition(value string) CollapseGroup {
	return c.set("expandIconPosition", value)
}

// Hidden controls the overall visibility of the collapse group
func (c CollapseGroup) Hidden(value bool) CollapseGroup {
	return c.set("hidden", value)
}

// HiddenOn sets a conditional expression to dynamically hide the component
func (c CollapseGroup) HiddenOn(value string) CollapseGroup {
	return c.set("hiddenOn", value)
}

// ID sets a unique identifier for the component
func (c CollapseGroup) ID(value string) CollapseGroup {
	return c.set("id", value)
}

// OnEvent configures event-driven actions for the collapse group
func (c CollapseGroup) OnEvent(value Event) CollapseGroup {
	return c.set("onEvent", value)
}

// Static enables or disables static display mode
func (c CollapseGroup) Static(value bool) CollapseGroup {
	return c.set("static", value)
}

// StaticClassName sets the CSS class name for static form item display
func (c CollapseGroup) StaticClassName(value string) CollapseGroup {
	return c.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input value display
func (c CollapseGroup) StaticInputClassName(value string) CollapseGroup {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (c CollapseGroup) StaticLabelClassName(value string) CollapseGroup {
	return c.set("staticLabelClassName", value)
}

// StaticOn sets a conditional expression for static display
func (c CollapseGroup) StaticOn(value string) CollapseGroup {
	return c.set("staticOn", value)
}

// StaticPlaceholder defines a placeholder for empty static values
func (c CollapseGroup) StaticPlaceholder(value string) CollapseGroup {
	return c.set("staticPlaceholder", value)
}

// StaticSchema sets the schema.Schema for static display
func (c CollapseGroup) StaticSchema(value string) CollapseGroup {
	return c.set("staticSchema", value)
}

// Style sets custom inline styles for the component
func (c CollapseGroup) Style(value any) CollapseGroup {
	return c.set("style", value)
}

// TestIdBuilder configures test ID generation
func (c CollapseGroup) TestIdBuilder(value string) CollapseGroup {
	return c.set("testIdBuilder", value)
}

// Testid sets a specific test identifier
func (c CollapseGroup) Testid(value string) CollapseGroup {
	return c.set("testid", value)
}

// UseMobileUI enables or disables mobile-specific UI styling
func (c CollapseGroup) UseMobileUI(value bool) CollapseGroup {
	return c.set("useMobileUI", value)
}

// Visible controls the overall visibility of the component
func (c CollapseGroup) Visible(value bool) CollapseGroup {
	return c.set("visible", value)
}

// VisibleOn sets a conditional expression for component visibility
func (c CollapseGroup) VisibleOn(value string) CollapseGroup {
	return c.set("visibleOn", value)
}
