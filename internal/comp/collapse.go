package comp

import "github.com/zrcoder/amisgo/schema"

// Collapse represents a collapsible container component
type Collapse schema.Schema

func NewCollapse() Collapse {
	return Collapse{"type": "collapse"}
}

func (c Collapse) set(key string, value any) Collapse {
	c[key] = value
	return c
}

// Body sets the content to be displayed inside the collapsible section
func (c Collapse) Body(value ...any) Collapse {
	return c.set("body", value)
}

// BodyClassName sets the CSS class name for the body content container
func (c Collapse) BodyClassName(value string) Collapse {
	return c.set("bodyClassName", value)
}

// ClassName sets the CSS class name for the collapse component
func (c Collapse) ClassName(value string) Collapse {
	return c.set("className", value)
}

// Collapsable enables or disables the ability to collapse/expand the component
func (c Collapse) Collapsable(value bool) Collapse {
	return c.set("collapsable", value)
}

// CollapseHeader sets custom header content for the collapse component
func (c Collapse) CollapseHeader(value string) Collapse {
	return c.set("collapseHeader", value)
}

// Collapsed determines the initial collapsed state of the component
func (c Collapse) Collapsed(value bool) Collapse {
	return c.set("collapsed", value)
}

// Disabled enables or disables the entire collapse component
func (c Collapse) Disabled(value bool) Collapse {
	return c.set("disabled", value)
}

// DisabledOn sets a conditional expression to dynamically disable the component
func (c Collapse) DisabledOn(value string) Collapse {
	return c.set("disabledOn", value)
}

// DivideLine enables or disables a visual dividing line
func (c Collapse) DivideLine(value bool) Collapse {
	return c.set("divideLine", value)
}

// EditorSetting configures editor-specific settings for the component
func (c Collapse) EditorSetting(value string) Collapse {
	return c.set("editorSetting", value)
}

// ExpandIcon sets a custom icon for expanding/collapsing
func (c Collapse) ExpandIcon(value string) Collapse {
	return c.set("expandIcon", value)
}

// Header sets the header content for the collapse component
func (c Collapse) Header(value ...any) Collapse {
	return c.set("header", value)
}

// HeaderPosition sets the positioning of the header
func (c Collapse) HeaderPosition(value string) Collapse {
	return c.set("headerPosition", value)
}

// HeadingClassName sets the CSS class name for the header element
func (c Collapse) HeadingClassName(value string) Collapse {
	return c.set("headingClassName", value)
}

// Hidden controls the overall visibility of the component
func (c Collapse) Hidden(value bool) Collapse {
	return c.set("hidden", value)
}

// HiddenOn sets a conditional expression to dynamically hide the component
func (c Collapse) HiddenOn(value string) Collapse {
	return c.set("hiddenOn", value)
}

// ID sets a unique identifier for the component
func (c Collapse) ID(value string) Collapse {
	return c.set("id", value)
}

// Key sets a unique key for the component (useful in lists)
func (c Collapse) Key(value string) Collapse {
	return c.set("key", value)
}

// MountOnEnter determines whether to mount the component when it enters the view
func (c Collapse) MountOnEnter(value bool) Collapse {
	return c.set("mountOnEnter", value)
}

// OnEvent configures event-driven actions for the component
func (c Collapse) OnEvent(value any) Collapse {
	return c.set("onEvent", value)
}

// ShowArrow enables or disables the expand/collapse arrow
func (c Collapse) ShowArrow(value bool) Collapse {
	return c.set("showArrow", value)
}

// Size sets the size of the collapse component
func (c Collapse) Size(value string) Collapse {
	return c.set("size", value)
}

// Static enables or disables static display mode
func (c Collapse) Static(value bool) Collapse {
	return c.set("static", value)
}

// StaticClassName sets the CSS class name for static form item display
func (c Collapse) StaticClassName(value string) Collapse {
	return c.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input value display
func (c Collapse) StaticInputClassName(value string) Collapse {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (c Collapse) StaticLabelClassName(value string) Collapse {
	return c.set("staticLabelClassName", value)
}

// StaticOn sets a conditional expression for static display
func (c Collapse) StaticOn(value string) Collapse {
	return c.set("staticOn", value)
}

// StaticPlaceholder defines a placeholder for empty static values
func (c Collapse) StaticPlaceholder(value string) Collapse {
	return c.set("staticPlaceholder", value)
}

// StaticSchema sets the schema.Schema for static display
func (c Collapse) StaticSchema(value string) Collapse {
	return c.set("staticSchema", value)
}

// Style sets custom inline styles for the component
func (c Collapse) Style(value any) Collapse {
	return c.set("style", value)
}

// TestIdBuilder configures test ID generation
func (c Collapse) TestIdBuilder(value string) Collapse {
	return c.set("testIdBuilder", value)
}

// Testid sets a specific test identifier
func (c Collapse) Testid(value string) Collapse {
	return c.set("testid", value)
}

// UnmountOnExit determines whether to unmount the component when it exits the view
func (c Collapse) UnmountOnExit(value bool) Collapse {
	return c.set("unmountOnExit", value)
}

// UseMobileUI enables or disables mobile-specific UI styling
func (c Collapse) UseMobileUI(value bool) Collapse {
	return c.set("useMobileUI", value)
}

// Visible controls the overall visibility of the component
func (c Collapse) Visible(value bool) Collapse {
	return c.set("visible", value)
}

// VisibleOn sets a conditional expression for component visibility
func (c Collapse) VisibleOn(value string) Collapse {
	return c.set("visibleOn", value)
}
