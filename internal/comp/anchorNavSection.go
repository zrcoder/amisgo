package comp

import "github.com/zrcoder/amisgo/schema"

// AnchorNavSection is an anchor area renderer. Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/anchor-nav
type AnchorNavSection schema.Schema

func NewAnchorNavSection() AnchorNavSection {
	return make(AnchorNavSection)
}

func (a AnchorNavSection) set(key string, value any) AnchorNavSection {
	a[key] = value
	return a
}

// Body sets the content.
func (a AnchorNavSection) Body(value ...any) AnchorNavSection {
	return a.set("body", value)
}

// Children sets the child nodes.
func (a AnchorNavSection) Children(value string) AnchorNavSection {
	return a.set("children", value)
}

// ClassName sets the container CSS class name.
func (a AnchorNavSection) ClassName(value string) AnchorNavSection {
	return a.set("className", value)
}

// Disabled sets whether the component is disabled.
func (a AnchorNavSection) Disabled(value bool) AnchorNavSection {
	return a.set("disabled", value)
}

// DisabledOn sets the expression to determine if the component is disabled.
func (a AnchorNavSection) DisabledOn(value string) AnchorNavSection {
	return a.set("disabledOn", value)
}

// EditorSetting sets the editor configuration, which can be ignored at runtime.
func (a AnchorNavSection) EditorSetting(value string) AnchorNavSection {
	return a.set("editorSetting", value)
}

// Hidden sets whether the component is hidden.
func (a AnchorNavSection) Hidden(value bool) AnchorNavSection {
	return a.set("hidden", value)
}

// HiddenOn sets the expression to determine if the component is hidden.
func (a AnchorNavSection) HiddenOn(value string) AnchorNavSection {
	return a.set("hiddenOn", value)
}

// Href sets the anchor link.
func (a AnchorNavSection) Href(value string) AnchorNavSection {
	return a.set("href", value)
}

// ID sets the unique component ID, mainly used for logging.
func (a AnchorNavSection) ID(value string) AnchorNavSection {
	return a.set("id", value)
}

// OnEvent sets the event action configuration.
func (a AnchorNavSection) OnEvent(value Event) AnchorNavSection {
	return a.set("onEvent", value)
}

// Static sets whether the component is displayed statically.
func (a AnchorNavSection) Static(value bool) AnchorNavSection {
	return a.set("static", value)
}

// StaticClassName sets the CSS class name for static display form items.
func (a AnchorNavSection) StaticClassName(value string) AnchorNavSection {
	return a.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static display form item values.
func (a AnchorNavSection) StaticInputClassName(value string) AnchorNavSection {
	return a.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static display form item labels.
func (a AnchorNavSection) StaticLabelClassName(value string) AnchorNavSection {
	return a.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the component is displayed statically.
func (a AnchorNavSection) StaticOn(value string) AnchorNavSection {
	return a.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display empty values.
func (a AnchorNavSection) StaticPlaceholder(value string) AnchorNavSection {
	return a.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema.Schema.
func (a AnchorNavSection) StaticSchema(value string) AnchorNavSection {
	return a.set("staticSchema", value)
}

// Style sets the component style.
func (a AnchorNavSection) Style(value any) AnchorNavSection {
	return a.set("style", value)
}

// TestIdBuilder sets the test ID builder.
func (a AnchorNavSection) TestIdBuilder(value string) AnchorNavSection {
	return a.set("testIdBuilder", value)
}

// Testid sets the test ID.
func (a AnchorNavSection) Testid(value string) AnchorNavSection {
	return a.set("testid", value)
}

// Title sets the navigation text description.
func (a AnchorNavSection) Title(value any) AnchorNavSection {
	return a.set("title", value)
}

// UseMobileUI sets whether to disable mobile styles at the component level.
func (a AnchorNavSection) UseMobileUI(value bool) AnchorNavSection {
	return a.set("useMobileUI", value)
}

// Visible sets whether the component is visible.
func (a AnchorNavSection) Visible(value bool) AnchorNavSection {
	return a.set("visible", value)
}

// VisibleOn sets the expression to determine if the component is visible.
func (a AnchorNavSection) VisibleOn(value string) AnchorNavSection {
	return a.set("visibleOn", value)
}
