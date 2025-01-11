package comp

import "github.com/zrcoder/amisgo/model"

// AnchorNav represents a navigation component with anchor points for document sections
// Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/anchor-nav
type anchorNav model.Schema

// AnchorNav creates a new AnchorNav instance
func AnchorNav() anchorNav {
	return anchorNav{"type": "anchor-nav"}
}

func (a anchorNav) set(key string, value any) anchorNav {
	a[key] = value
	return a
}

// Active sets the currently activated (located) section
func (a anchorNav) Active(value string) anchorNav {
	return a.set("active", value)
}

// ClassName sets the CSS class name for styling
// Supports string or object configuration with conditional expressions
// Example: className: { "red": "data.progress > 80", "blue": "data.progress > 60" }
func (a anchorNav) ClassName(value string) anchorNav {
	return a.set("className", value)
}

// Direction sets the navigation layout
// Possible values: vertical | horizontal
func (a anchorNav) Direction(value string) anchorNav {
	return a.set("direction", value)
}

// Disabled disables the entire navigation component
func (a anchorNav) Disabled(value bool) anchorNav {
	return a.set("disabled", value)
}

// DisabledOn sets a conditional expression for disabling the component
// Expression syntax: `data.xxx > 5`
func (a anchorNav) DisabledOn(value string) anchorNav {
	return a.set("disabledOn", value)
}

// EditorSetting configures editor-specific settings (ignored during runtime)
func (a anchorNav) EditorSetting(value string) anchorNav {
	return a.set("editorSetting", value)
}

// Hidden controls the visibility of the entire component
func (a anchorNav) Hidden(value bool) anchorNav {
	return a.set("hidden", value)
}

// HiddenOn sets a conditional expression for hiding the component
// Expression syntax: `data.xxx > 5`
func (a anchorNav) HiddenOn(value string) anchorNav {
	return a.set("hiddenOn", value)
}

// ID sets a unique identifier for the component, primarily used for logging
func (a anchorNav) ID(value string) anchorNav {
	return a.set("id", value)
}

// LinkClassName sets the CSS class name for navigation links
// Supports string or object configuration with conditional expressions
func (a anchorNav) LinkClassName(value string) anchorNav {
	return a.set("linkClassName", value)
}

// Links defines the collection of navigation sections
func (a anchorNav) Links(value string) anchorNav {
	return a.set("links", value)
}

// OnEvent configures event-driven actions
func (a anchorNav) OnEvent(value any) anchorNav {
	return a.set("onEvent", value)
}

// SectionClassName sets the CSS class name for individual sections
// Supports string or object configuration with conditional expressions
func (a anchorNav) SectionClassName(value string) anchorNav {
	return a.set("sectionClassName", value)
}

// Static determines if the component is statically displayed
func (a anchorNav) Static(value bool) anchorNav {
	return a.set("static", value)
}

// StaticClassName sets the CSS class for static display
func (a anchorNav) StaticClassName(value string) anchorNav {
	return a.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class for static input display
func (a anchorNav) StaticInputClassName(value string) anchorNav {
	return a.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class for static label display
func (a anchorNav) StaticLabelClassName(value string) anchorNav {
	return a.set("staticLabelClassName", value)
}

// StaticOn sets a conditional expression for static display
// Expression syntax: `data.xxx > 5`
func (a anchorNav) StaticOn(value string) anchorNav {
	return a.set("staticOn", value)
}

// StaticPlaceholder defines a placeholder for empty static values
func (a anchorNav) StaticPlaceholder(value string) anchorNav {
	return a.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (a anchorNav) StaticSchema(value string) anchorNav {
	return a.set("staticSchema", value)
}

// Style sets the component's inline styles
func (a anchorNav) Style(value any) anchorNav {
	return a.set("style", value)
}

// TestIdBuilder configures test ID generation
func (a anchorNav) TestIdBuilder(value string) anchorNav {
	return a.set("testIdBuilder", value)
}

// Testid sets a specific test identifier
func (a anchorNav) Testid(value string) anchorNav {
	return a.set("testid", value)
}

// UseMobileUI enables or disables mobile-specific styling
func (a anchorNav) UseMobileUI(value bool) anchorNav {
	return a.set("useMobileUI", value)
}

// Visible controls the overall visibility of the component
func (a anchorNav) Visible(value bool) anchorNav {
	return a.set("visible", value)
}

// VisibleOn sets a conditional expression for component visibility
// Expression syntax: `data.xxx > 5`
func (a anchorNav) VisibleOn(value string) anchorNav {
	return a.set("visibleOn", value)
}
