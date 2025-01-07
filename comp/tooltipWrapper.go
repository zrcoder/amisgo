package comp

// tooltipWrapper

type tooltipWrapper Schema

// TooltipWrapper creates a new TooltipWrapper instance
func TooltipWrapper() tooltipWrapper {
	return tooltipWrapper{}.set("type", "tooltip-wrapper")
}

func (tw tooltipWrapper) set(key string, value any) tooltipWrapper {
	tw[key] = value
	return tw
}

// Body sets the content area
func (tw tooltipWrapper) Body(value ...any) tooltipWrapper {
	return tw.set("body", value)
}

// ClassName sets the CSS class name for the content area
func (tw tooltipWrapper) ClassName(value string) tooltipWrapper {
	return tw.set("className", value)
}

// Content sets the tooltip content
func (tw tooltipWrapper) Content(value string) tooltipWrapper {
	return tw.set("content", value)
}

// Disabled sets whether the tooltip is disabled
func (tw tooltipWrapper) Disabled(value bool) tooltipWrapper {
	return tw.set("disabled", value)
}

// DisabledOn sets the expression to disable the tooltip
func (tw tooltipWrapper) DisabledOn(value string) tooltipWrapper {
	return tw.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (tw tooltipWrapper) EditorSetting(value string) tooltipWrapper {
	return tw.set("editorSetting", value)
}

// Enterable sets whether the tooltip can be entered
func (tw tooltipWrapper) Enterable(value bool) tooltipWrapper {
	return tw.set("enterable", value)
}

// Hidden sets whether the tooltip is hidden
func (tw tooltipWrapper) Hidden(value bool) tooltipWrapper {
	return tw.set("hidden", value)
}

// HiddenOn sets the expression to hide the tooltip
func (tw tooltipWrapper) HiddenOn(value string) tooltipWrapper {
	return tw.set("hiddenOn", value)
}

// ID sets the unique id for the component
func (tw tooltipWrapper) ID(value string) tooltipWrapper {
	return tw.set("id", value)
}

// Inline sets whether the content area is displayed inline
func (tw tooltipWrapper) Inline(value bool) tooltipWrapper {
	return tw.set("inline", value)
}

// MouseEnterDelay sets the delay time for showing the tooltip
func (tw tooltipWrapper) MouseEnterDelay(value string) tooltipWrapper {
	return tw.set("mouseEnterDelay", value)
}

// MouseLeaveDelay sets the delay time for hiding the tooltip
func (tw tooltipWrapper) MouseLeaveDelay(value string) tooltipWrapper {
	return tw.set("mouseLeaveDelay", value)
}

// Offset sets the offset for the tooltip position
func (tw tooltipWrapper) Offset(value string) tooltipWrapper {
	return tw.set("offset", value)
}

// OnEvent sets the event action configuration
func (tw tooltipWrapper) OnEvent(value any) tooltipWrapper {
	return tw.set("onEvent", value)
}

// Placement sets the position of the tooltip
func (tw tooltipWrapper) Placement(value string) tooltipWrapper {
	return tw.set("placement", value)
}

// RootClose sets whether clicking outside closes the tooltip
func (tw tooltipWrapper) RootClose(value bool) tooltipWrapper {
	return tw.set("rootClose", value)
}

// ShowArrow sets whether to show the arrow on the tooltip
func (tw tooltipWrapper) ShowArrow(value bool) tooltipWrapper {
	return tw.set("showArrow", value)
}

// Static sets whether the tooltip is displayed statically
func (tw tooltipWrapper) Static(value bool) tooltipWrapper {
	return tw.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (tw tooltipWrapper) StaticClassName(value string) tooltipWrapper {
	return tw.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (tw tooltipWrapper) StaticInputClassName(value string) tooltipWrapper {
	return tw.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (tw tooltipWrapper) StaticLabelClassName(value string) tooltipWrapper {
	return tw.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display
func (tw tooltipWrapper) StaticOn(value string) tooltipWrapper {
	return tw.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (tw tooltipWrapper) StaticPlaceholder(value string) tooltipWrapper {
	return tw.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (tw tooltipWrapper) StaticSchema(value string) tooltipWrapper {
	return tw.set("staticSchema", value)
}

// Style sets the custom style for the content area
func (tw tooltipWrapper) Style(value any) tooltipWrapper {
	return tw.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (tw tooltipWrapper) TestIdBuilder(value string) tooltipWrapper {
	return tw.set("testIdBuilder", value)
}

// Testid sets the test ID
func (tw tooltipWrapper) Testid(value string) tooltipWrapper {
	return tw.set("testid", value)
}

// Title sets the tooltip title
func (tw tooltipWrapper) Title(value any) tooltipWrapper {
	return tw.set("title", value)
}

// Tooltip sets the tooltip text
func (tw tooltipWrapper) Tooltip(value string) tooltipWrapper {
	return tw.set("tooltip", value)
}

// TooltipClassName sets the CSS class name for the tooltip
func (tw tooltipWrapper) TooltipClassName(value string) tooltipWrapper {
	return tw.set("tooltipClassName", value)
}

// TooltipStyle sets the custom style for the tooltip
func (tw tooltipWrapper) TooltipStyle(value any) tooltipWrapper {
	return tw.set("tooltipStyle", value)
}

// TooltipTheme sets the theme for the tooltip
func (tw tooltipWrapper) TooltipTheme(value string) tooltipWrapper {
	return tw.set("tooltipTheme", value)
}

// Trigger sets the trigger method for the tooltip
func (tw tooltipWrapper) Trigger(value string) tooltipWrapper {
	return tw.set("trigger", value)
}

// UseMobileUI sets whether to use mobile UI
func (tw tooltipWrapper) UseMobileUI(value bool) tooltipWrapper {
	return tw.set("useMobileUI", value)
}

// Visible sets whether the tooltip is visible
func (tw tooltipWrapper) Visible(value bool) tooltipWrapper {
	return tw.set("visible", value)
}

// VisibleOn sets the expression for visibility
func (tw tooltipWrapper) VisibleOn(value string) tooltipWrapper {
	return tw.set("visibleOn", value)
}

// WrapperComponent sets the wrapper component for the content area
func (tw tooltipWrapper) WrapperComponent(value string) tooltipWrapper {
	return tw.set("wrapperComponent", value)
}
