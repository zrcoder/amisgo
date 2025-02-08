package comp

import "github.com/zrcoder/amisgo/model"

// TooltipWrapper
type TooltipWrapper model.Schema

func NewTooltipWrapper() TooltipWrapper {
	return TooltipWrapper{"type": "tooltip-wrapper"}
}

func (tw TooltipWrapper) set(key string, value any) TooltipWrapper {
	tw[key] = value
	return tw
}

// Body sets the content area
func (tw TooltipWrapper) Body(value ...any) TooltipWrapper {
	return tw.set("body", value)
}

// ClassName sets the CSS class name for the content area
func (tw TooltipWrapper) ClassName(value string) TooltipWrapper {
	return tw.set("className", value)
}

// Content sets the tooltip content
func (tw TooltipWrapper) Content(value string) TooltipWrapper {
	return tw.set("content", value)
}

// Disabled sets whether the tooltip is disabled
func (tw TooltipWrapper) Disabled(value bool) TooltipWrapper {
	return tw.set("disabled", value)
}

// DisabledOn sets the expression to disable the tooltip
func (tw TooltipWrapper) DisabledOn(value string) TooltipWrapper {
	return tw.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (tw TooltipWrapper) EditorSetting(value string) TooltipWrapper {
	return tw.set("editorSetting", value)
}

// Enterable sets whether the tooltip can be entered
func (tw TooltipWrapper) Enterable(value bool) TooltipWrapper {
	return tw.set("enterable", value)
}

// Hidden sets whether the tooltip is hidden
func (tw TooltipWrapper) Hidden(value bool) TooltipWrapper {
	return tw.set("hidden", value)
}

// HiddenOn sets the expression to hide the tooltip
func (tw TooltipWrapper) HiddenOn(value string) TooltipWrapper {
	return tw.set("hiddenOn", value)
}

// ID sets the unique id for the component
func (tw TooltipWrapper) ID(value string) TooltipWrapper {
	return tw.set("id", value)
}

// Inline sets whether the content area is displayed inline
func (tw TooltipWrapper) Inline(value bool) TooltipWrapper {
	return tw.set("inline", value)
}

// MouseEnterDelay sets the delay time for showing the tooltip
func (tw TooltipWrapper) MouseEnterDelay(value string) TooltipWrapper {
	return tw.set("mouseEnterDelay", value)
}

// MouseLeaveDelay sets the delay time for hiding the tooltip
func (tw TooltipWrapper) MouseLeaveDelay(value string) TooltipWrapper {
	return tw.set("mouseLeaveDelay", value)
}

// Offset sets the offset for the tooltip position
func (tw TooltipWrapper) Offset(value string) TooltipWrapper {
	return tw.set("offset", value)
}

// OnEvent sets the event action configuration
func (tw TooltipWrapper) OnEvent(value any) TooltipWrapper {
	return tw.set("onEvent", value)
}

// Placement sets the position of the tooltip
func (tw TooltipWrapper) Placement(value string) TooltipWrapper {
	return tw.set("placement", value)
}

// RootClose sets whether clicking outside closes the tooltip
func (tw TooltipWrapper) RootClose(value bool) TooltipWrapper {
	return tw.set("rootClose", value)
}

// ShowArrow sets whether to show the arrow on the tooltip
func (tw TooltipWrapper) ShowArrow(value bool) TooltipWrapper {
	return tw.set("showArrow", value)
}

// Static sets whether the tooltip is displayed statically
func (tw TooltipWrapper) Static(value bool) TooltipWrapper {
	return tw.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (tw TooltipWrapper) StaticClassName(value string) TooltipWrapper {
	return tw.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (tw TooltipWrapper) StaticInputClassName(value string) TooltipWrapper {
	return tw.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (tw TooltipWrapper) StaticLabelClassName(value string) TooltipWrapper {
	return tw.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display
func (tw TooltipWrapper) StaticOn(value string) TooltipWrapper {
	return tw.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (tw TooltipWrapper) StaticPlaceholder(value string) TooltipWrapper {
	return tw.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (tw TooltipWrapper) StaticSchema(value string) TooltipWrapper {
	return tw.set("staticSchema", value)
}

// Style sets the custom style for the content area
func (tw TooltipWrapper) Style(value any) TooltipWrapper {
	return tw.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (tw TooltipWrapper) TestIdBuilder(value string) TooltipWrapper {
	return tw.set("testIdBuilder", value)
}

// Testid sets the test ID
func (tw TooltipWrapper) Testid(value string) TooltipWrapper {
	return tw.set("testid", value)
}

// Title sets the tooltip title
func (tw TooltipWrapper) Title(value any) TooltipWrapper {
	return tw.set("title", value)
}

// Tooltip sets the tooltip text
func (tw TooltipWrapper) Tooltip(value string) TooltipWrapper {
	return tw.set("tooltip", value)
}

// TooltipClassName sets the CSS class name for the tooltip
func (tw TooltipWrapper) TooltipClassName(value string) TooltipWrapper {
	return tw.set("tooltipClassName", value)
}

// TooltipStyle sets the custom style for the tooltip
func (tw TooltipWrapper) TooltipStyle(value any) TooltipWrapper {
	return tw.set("tooltipStyle", value)
}

// TooltipTheme sets the theme for the tooltip
func (tw TooltipWrapper) TooltipTheme(value string) TooltipWrapper {
	return tw.set("tooltipTheme", value)
}

// Trigger sets the trigger method for the tooltip
func (tw TooltipWrapper) Trigger(value string) TooltipWrapper {
	return tw.set("trigger", value)
}

// UseMobileUI sets whether to use mobile UI
func (tw TooltipWrapper) UseMobileUI(value bool) TooltipWrapper {
	return tw.set("useMobileUI", value)
}

// Visible sets whether the tooltip is visible
func (tw TooltipWrapper) Visible(value bool) TooltipWrapper {
	return tw.set("visible", value)
}

// VisibleOn sets the expression for visibility
func (tw TooltipWrapper) VisibleOn(value string) TooltipWrapper {
	return tw.set("visibleOn", value)
}

// WrapperComponent sets the wrapper component for the content area
func (tw TooltipWrapper) WrapperComponent(value string) TooltipWrapper {
	return tw.set("wrapperComponent", value)
}
