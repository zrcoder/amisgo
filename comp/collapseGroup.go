package comp

// CollapseGroup represents a group of collapsible panels with shared behavior
type collapseGroup Schema

// CollapseGroup creates a new CollapseGroup instance
func CollapseGroup() collapseGroup {
	return make(collapseGroup).set("type", "collapse-group")
}

func (c collapseGroup) set(key string, value any) collapseGroup {
	c[key] = value
	return c
}

// Accordion enables or disables accordion mode (only one panel open at a time)
func (c collapseGroup) Accordion(value bool) collapseGroup {
	return c.set("accordion", value)
}

// ActiveKey sets the currently active panel(s)
func (c collapseGroup) ActiveKey(value any) collapseGroup {
	return c.set("activeKey", value)
}

// Body sets the content for the collapse group
func (c collapseGroup) Body(value ...any) collapseGroup {
	return c.set("body", value)
}

// ClassName sets the CSS class name for the container
func (c collapseGroup) ClassName(value string) collapseGroup {
	return c.set("className", value)
}

// Disabled enables or disables the entire collapse group
func (c collapseGroup) Disabled(value bool) collapseGroup {
	return c.set("disabled", value)
}

// DisabledOn sets a conditional expression to dynamically disable the component
func (c collapseGroup) DisabledOn(value string) collapseGroup {
	return c.set("disabledOn", value)
}

// EditorSetting configures editor-specific settings for the component
func (c collapseGroup) EditorSetting(value string) collapseGroup {
	return c.set("editorSetting", value)
}

// EnableFieldSetStyle enables or disables the FieldSet styling
func (c collapseGroup) EnableFieldSetStyle(value bool) collapseGroup {
	return c.set("enableFieldSetStyle", value)
}

// ExpandIcon sets a custom icon for expanding/collapsing panels
func (c collapseGroup) ExpandIcon(value any) collapseGroup {
	return c.set("expandIcon", value)
}

// ExpandIconPosition sets the placement of the expand/collapse icon
func (c collapseGroup) ExpandIconPosition(value string) collapseGroup {
	return c.set("expandIconPosition", value)
}

// Hidden controls the overall visibility of the collapse group
func (c collapseGroup) Hidden(value bool) collapseGroup {
	return c.set("hidden", value)
}

// HiddenOn sets a conditional expression to dynamically hide the component
func (c collapseGroup) HiddenOn(value string) collapseGroup {
	return c.set("hiddenOn", value)
}

// ID sets a unique identifier for the component
func (c collapseGroup) ID(value string) collapseGroup {
	return c.set("id", value)
}

// OnEvent configures event-driven actions for the collapse group
func (c collapseGroup) OnEvent(value any) collapseGroup {
	return c.set("onEvent", value)
}

// Static enables or disables static display mode
func (c collapseGroup) Static(value bool) collapseGroup {
	return c.set("static", value)
}

// StaticClassName sets the CSS class name for static form item display
func (c collapseGroup) StaticClassName(value string) collapseGroup {
	return c.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input value display
func (c collapseGroup) StaticInputClassName(value string) collapseGroup {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (c collapseGroup) StaticLabelClassName(value string) collapseGroup {
	return c.set("staticLabelClassName", value)
}

// StaticOn sets a conditional expression for static display
func (c collapseGroup) StaticOn(value string) collapseGroup {
	return c.set("staticOn", value)
}

// StaticPlaceholder defines a placeholder for empty static values
func (c collapseGroup) StaticPlaceholder(value string) collapseGroup {
	return c.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (c collapseGroup) StaticSchema(value string) collapseGroup {
	return c.set("staticSchema", value)
}

// Style sets custom inline styles for the component
func (c collapseGroup) Style(value any) collapseGroup {
	return c.set("style", value)
}

// TestIdBuilder configures test ID generation
func (c collapseGroup) TestIdBuilder(value string) collapseGroup {
	return c.set("testIdBuilder", value)
}

// Testid sets a specific test identifier
func (c collapseGroup) Testid(value string) collapseGroup {
	return c.set("testid", value)
}

// UseMobileUI enables or disables mobile-specific UI styling
func (c collapseGroup) UseMobileUI(value bool) collapseGroup {
	return c.set("useMobileUI", value)
}

// Visible controls the overall visibility of the component
func (c collapseGroup) Visible(value bool) collapseGroup {
	return c.set("visible", value)
}

// VisibleOn sets a conditional expression for component visibility
func (c collapseGroup) VisibleOn(value string) collapseGroup {
	return c.set("visibleOn", value)
}
