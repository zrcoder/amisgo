package comp

import "github.com/zrcoder/amisgo/schema"

// AnchorNav represents a navigation component with anchor points for document sections
// Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/anchor-nav
type AnchorNav schema.Schema

func NewAnchorNav() AnchorNav {
	return AnchorNav{"type": "anchor-nav"}
}

func (a AnchorNav) set(key string, value any) AnchorNav {
	a[key] = value
	return a
}

// Active sets the currently activated (located) section
func (a AnchorNav) Active(value string) AnchorNav {
	return a.set("active", value)
}

// ClassName sets the CSS class name for styling
// Supports string or object configuration with conditional expressions
// Example: className: { "red": "data.progress > 80", "blue": "data.progress > 60" }
func (a AnchorNav) ClassName(value string) AnchorNav {
	return a.set("className", value)
}

// Direction sets the navigation layout
// Possible values: vertical | horizontal
func (a AnchorNav) Direction(value string) AnchorNav {
	return a.set("direction", value)
}

// Disabled disables the entire navigation component
func (a AnchorNav) Disabled(value bool) AnchorNav {
	return a.set("disabled", value)
}

// DisabledOn sets a conditional expression for disabling the component
// Expression syntax: `data.xxx > 5`
func (a AnchorNav) DisabledOn(value string) AnchorNav {
	return a.set("disabledOn", value)
}

// EditorSetting configures editor-specific settings (ignored during runtime)
func (a AnchorNav) EditorSetting(value string) AnchorNav {
	return a.set("editorSetting", value)
}

// Hidden controls the visibility of the entire component
func (a AnchorNav) Hidden(value bool) AnchorNav {
	return a.set("hidden", value)
}

// HiddenOn sets a conditional expression for hiding the component
// Expression syntax: `data.xxx > 5`
func (a AnchorNav) HiddenOn(value string) AnchorNav {
	return a.set("hiddenOn", value)
}

// ID sets a unique identifier for the component, primarily used for logging
func (a AnchorNav) ID(value string) AnchorNav {
	return a.set("id", value)
}

// LinkClassName sets the CSS class name for navigation links
// Supports string or object configuration with conditional expressions
func (a AnchorNav) LinkClassName(value string) AnchorNav {
	return a.set("linkClassName", value)
}

// Links defines the collection of navigation sections
func (a AnchorNav) Links(value string) AnchorNav {
	return a.set("links", value)
}

// OnEvent configures event-driven actions
func (a AnchorNav) OnEvent(value Event) AnchorNav {
	return a.set("onEvent", value)
}

// SectionClassName sets the CSS class name for individual sections
// Supports string or object configuration with conditional expressions
func (a AnchorNav) SectionClassName(value string) AnchorNav {
	return a.set("sectionClassName", value)
}

// Static determines if the component is statically displayed
func (a AnchorNav) Static(value bool) AnchorNav {
	return a.set("static", value)
}

// StaticClassName sets the CSS class for static display
func (a AnchorNav) StaticClassName(value string) AnchorNav {
	return a.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class for static input display
func (a AnchorNav) StaticInputClassName(value string) AnchorNav {
	return a.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class for static label display
func (a AnchorNav) StaticLabelClassName(value string) AnchorNav {
	return a.set("staticLabelClassName", value)
}

// StaticOn sets a conditional expression for static display
// Expression syntax: `data.xxx > 5`
func (a AnchorNav) StaticOn(value string) AnchorNav {
	return a.set("staticOn", value)
}

// StaticPlaceholder defines a placeholder for empty static values
func (a AnchorNav) StaticPlaceholder(value string) AnchorNav {
	return a.set("staticPlaceholder", value)
}

// StaticSchema sets the schema.Schema for static display
func (a AnchorNav) StaticSchema(value string) AnchorNav {
	return a.set("staticSchema", value)
}

// Style sets the component's inline styles
func (a AnchorNav) Style(value any) AnchorNav {
	return a.set("style", value)
}

// TestIdBuilder configures test ID generation
func (a AnchorNav) TestIdBuilder(value string) AnchorNav {
	return a.set("testIdBuilder", value)
}

// Testid sets a specific test identifier
func (a AnchorNav) Testid(value string) AnchorNav {
	return a.set("testid", value)
}

// UseMobileUI enables or disables mobile-specific styling
func (a AnchorNav) UseMobileUI(value bool) AnchorNav {
	return a.set("useMobileUI", value)
}

// Visible controls the overall visibility of the component
func (a AnchorNav) Visible(value bool) AnchorNav {
	return a.set("visible", value)
}

// VisibleOn sets a conditional expression for component visibility
// Expression syntax: `data.xxx > 5`
func (a AnchorNav) VisibleOn(value string) AnchorNav {
	return a.set("visibleOn", value)
}
