package comp

import "github.com/zrcoder/amisgo/model"

// anchorNavSection is an anchor area renderer. Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/anchor-nav
type anchorNavSection model.Schema

// AnchorNavSection creates a new instance of AnchorNavSection.
func AnchorNavSection() anchorNavSection {
	return make(anchorNavSection)
}

func (a anchorNavSection) set(key string, value any) anchorNavSection {
	a[key] = value
	return a
}

// Body sets the content.
func (a anchorNavSection) Body(value ...any) anchorNavSection {
	return a.set("body", value)
}

// Children sets the child nodes.
func (a anchorNavSection) Children(value string) anchorNavSection {
	return a.set("children", value)
}

// ClassName sets the container CSS class name.
func (a anchorNavSection) ClassName(value string) anchorNavSection {
	return a.set("className", value)
}

// Disabled sets whether the component is disabled.
func (a anchorNavSection) Disabled(value bool) anchorNavSection {
	return a.set("disabled", value)
}

// DisabledOn sets the expression to determine if the component is disabled.
func (a anchorNavSection) DisabledOn(value string) anchorNavSection {
	return a.set("disabledOn", value)
}

// EditorSetting sets the editor configuration, which can be ignored at runtime.
func (a anchorNavSection) EditorSetting(value string) anchorNavSection {
	return a.set("editorSetting", value)
}

// Hidden sets whether the component is hidden.
func (a anchorNavSection) Hidden(value bool) anchorNavSection {
	return a.set("hidden", value)
}

// HiddenOn sets the expression to determine if the component is hidden.
func (a anchorNavSection) HiddenOn(value string) anchorNavSection {
	return a.set("hiddenOn", value)
}

// Href sets the anchor link.
func (a anchorNavSection) Href(value string) anchorNavSection {
	return a.set("href", value)
}

// ID sets the unique component ID, mainly used for logging.
func (a anchorNavSection) ID(value string) anchorNavSection {
	return a.set("id", value)
}

// OnEvent sets the event action configuration.
func (a anchorNavSection) OnEvent(value any) anchorNavSection {
	return a.set("onEvent", value)
}

// Static sets whether the component is displayed statically.
func (a anchorNavSection) Static(value bool) anchorNavSection {
	return a.set("static", value)
}

// StaticClassName sets the CSS class name for static display form items.
func (a anchorNavSection) StaticClassName(value string) anchorNavSection {
	return a.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static display form item values.
func (a anchorNavSection) StaticInputClassName(value string) anchorNavSection {
	return a.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static display form item labels.
func (a anchorNavSection) StaticLabelClassName(value string) anchorNavSection {
	return a.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the component is displayed statically.
func (a anchorNavSection) StaticOn(value string) anchorNavSection {
	return a.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display empty values.
func (a anchorNavSection) StaticPlaceholder(value string) anchorNavSection {
	return a.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema.
func (a anchorNavSection) StaticSchema(value string) anchorNavSection {
	return a.set("staticSchema", value)
}

// Style sets the component style.
func (a anchorNavSection) Style(value any) anchorNavSection {
	return a.set("style", value)
}

// TestIdBuilder sets the test ID builder.
func (a anchorNavSection) TestIdBuilder(value string) anchorNavSection {
	return a.set("testIdBuilder", value)
}

// Testid sets the test ID.
func (a anchorNavSection) Testid(value string) anchorNavSection {
	return a.set("testid", value)
}

// Title sets the navigation text description.
func (a anchorNavSection) Title(value any) anchorNavSection {
	return a.set("title", value)
}

// UseMobileUI sets whether to disable mobile styles at the component level.
func (a anchorNavSection) UseMobileUI(value bool) anchorNavSection {
	return a.set("useMobileUI", value)
}

// Visible sets whether the component is visible.
func (a anchorNavSection) Visible(value bool) anchorNavSection {
	return a.set("visible", value)
}

// VisibleOn sets the expression to determine if the component is visible.
func (a anchorNavSection) VisibleOn(value string) anchorNavSection {
	return a.set("visibleOn", value)
}
