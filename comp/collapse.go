package comp

import "github.com/zrcoder/amisgo/model"

// Collapse represents a collapsible container component
type collapse model.Schema

// Collapse creates a new Collapse instance
func Collapse() collapse {
	return make(collapse).set("type", "collapse")
}

func (c collapse) set(key string, value any) collapse {
	c[key] = value
	return c
}

// Body sets the content to be displayed inside the collapsible section
func (c collapse) Body(value ...any) collapse {
	return c.set("body", value)
}

// BodyClassName sets the CSS class name for the body content container
func (c collapse) BodyClassName(value string) collapse {
	return c.set("bodyClassName", value)
}

// ClassName sets the CSS class name for the collapse component
func (c collapse) ClassName(value string) collapse {
	return c.set("className", value)
}

// Collapsable enables or disables the ability to collapse/expand the component
func (c collapse) Collapsable(value bool) collapse {
	return c.set("collapsable", value)
}

// CollapseHeader sets custom header content for the collapse component
func (c collapse) CollapseHeader(value string) collapse {
	return c.set("collapseHeader", value)
}

// Collapsed determines the initial collapsed state of the component
func (c collapse) Collapsed(value bool) collapse {
	return c.set("collapsed", value)
}

// Disabled enables or disables the entire collapse component
func (c collapse) Disabled(value bool) collapse {
	return c.set("disabled", value)
}

// DisabledOn sets a conditional expression to dynamically disable the component
func (c collapse) DisabledOn(value string) collapse {
	return c.set("disabledOn", value)
}

// DivideLine enables or disables a visual dividing line
func (c collapse) DivideLine(value bool) collapse {
	return c.set("divideLine", value)
}

// EditorSetting configures editor-specific settings for the component
func (c collapse) EditorSetting(value string) collapse {
	return c.set("editorSetting", value)
}

// ExpandIcon sets a custom icon for expanding/collapsing
func (c collapse) ExpandIcon(value string) collapse {
	return c.set("expandIcon", value)
}

// Header sets the header content for the collapse component
func (c collapse) Header(value ...any) collapse {
	return c.set("header", value)
}

// HeaderPosition sets the positioning of the header
func (c collapse) HeaderPosition(value string) collapse {
	return c.set("headerPosition", value)
}

// HeadingClassName sets the CSS class name for the header element
func (c collapse) HeadingClassName(value string) collapse {
	return c.set("headingClassName", value)
}

// Hidden controls the overall visibility of the component
func (c collapse) Hidden(value bool) collapse {
	return c.set("hidden", value)
}

// HiddenOn sets a conditional expression to dynamically hide the component
func (c collapse) HiddenOn(value string) collapse {
	return c.set("hiddenOn", value)
}

// ID sets a unique identifier for the component
func (c collapse) ID(value string) collapse {
	return c.set("id", value)
}

// Key sets a unique key for the component (useful in lists)
func (c collapse) Key(value string) collapse {
	return c.set("key", value)
}

// MountOnEnter determines whether to mount the component when it enters the view
func (c collapse) MountOnEnter(value bool) collapse {
	return c.set("mountOnEnter", value)
}

// OnEvent configures event-driven actions for the component
func (c collapse) OnEvent(value any) collapse {
	return c.set("onEvent", value)
}

// ShowArrow enables or disables the expand/collapse arrow
func (c collapse) ShowArrow(value bool) collapse {
	return c.set("showArrow", value)
}

// Size sets the size of the collapse component
func (c collapse) Size(value string) collapse {
	return c.set("size", value)
}

// Static enables or disables static display mode
func (c collapse) Static(value bool) collapse {
	return c.set("static", value)
}

// StaticClassName sets the CSS class name for static form item display
func (c collapse) StaticClassName(value string) collapse {
	return c.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input value display
func (c collapse) StaticInputClassName(value string) collapse {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (c collapse) StaticLabelClassName(value string) collapse {
	return c.set("staticLabelClassName", value)
}

// StaticOn sets a conditional expression for static display
func (c collapse) StaticOn(value string) collapse {
	return c.set("staticOn", value)
}

// StaticPlaceholder defines a placeholder for empty static values
func (c collapse) StaticPlaceholder(value string) collapse {
	return c.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (c collapse) StaticSchema(value string) collapse {
	return c.set("staticSchema", value)
}

// Style sets custom inline styles for the component
func (c collapse) Style(value any) collapse {
	return c.set("style", value)
}

// TestIdBuilder configures test ID generation
func (c collapse) TestIdBuilder(value string) collapse {
	return c.set("testIdBuilder", value)
}

// Testid sets a specific test identifier
func (c collapse) Testid(value string) collapse {
	return c.set("testid", value)
}

// UnmountOnExit determines whether to unmount the component when it exits the view
func (c collapse) UnmountOnExit(value bool) collapse {
	return c.set("unmountOnExit", value)
}

// UseMobileUI enables or disables mobile-specific UI styling
func (c collapse) UseMobileUI(value bool) collapse {
	return c.set("useMobileUI", value)
}

// Visible controls the overall visibility of the component
func (c collapse) Visible(value bool) collapse {
	return c.set("visible", value)
}

// VisibleOn sets a conditional expression for component visibility
func (c collapse) VisibleOn(value string) collapse {
	return c.set("visibleOn", value)
}
