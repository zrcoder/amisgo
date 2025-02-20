package comp

import "github.com/zrcoder/amisgo/schema"

type Remark schema.Schema

func NewRemark() Remark {
	return Remark{"type": "remark"}
}

func (rm Remark) set(key string, value any) Remark {
	rm[key] = value
	return rm
}

// ClassName sets the container's CSS class name.
func (rm Remark) ClassName(value string) Remark {
	return rm.set("className", value)
}

// Content sets the content of the remark.
func (rm Remark) Content(value string) Remark {
	return rm.set("content", value)
}

// Disabled sets whether the remark is disabled.
func (rm Remark) Disabled(value bool) Remark {
	return rm.set("disabled", value)
}

// DisabledOn sets whether the remark is disabled by expression.
func (rm Remark) DisabledOn(value string) Remark {
	return rm.set("disabledOn", value)
}

// EditorSetting sets the editor setting.
func (rm Remark) EditorSetting(value string) Remark {
	return rm.set("editorSetting", value)
}

// Hidden sets whether the remark is hidden.
func (rm Remark) Hidden(value bool) Remark {
	return rm.set("hidden", value)
}

// HiddenOn sets whether the remark is hidden by expression.
func (rm Remark) HiddenOn(value string) Remark {
	return rm.set("hiddenOn", value)
}

// Icon sets the icon of the remark.
func (rm Remark) Icon(value string) Remark {
	return rm.set("icon", value)
}

// ID sets the unique ID of the remark.
func (rm Remark) ID(value string) Remark {
	return rm.set("id", value)
}

// Label sets the label of the remark.
func (rm Remark) Label(value string) Remark {
	return rm.set("label", value)
}

// OnEvent sets the event handler of the remark.
func (rm Remark) OnEvent(value Event) Remark {
	return rm.set("onEvent", value)
}

// Placement sets the placement of the remark.
func (rm Remark) Placement(value string) Remark {
	return rm.set("placement", value)
}

// RootClose sets whether to close the remark when clicking outside.
func (rm Remark) RootClose(value bool) Remark {
	return rm.set("rootClose", value)
}

// Shape sets the shape of the icon.
func (rm Remark) Shape(value string) Remark {
	return rm.set("shape", value)
}

// Static sets whether the remark is static.
func (rm Remark) Static(value bool) Remark {
	return rm.set("static", value)
}

// StaticClassName sets the static CSS class name of the remark.
func (rm Remark) StaticClassName(value string) Remark {
	return rm.set("staticClassName", value)
}

// StaticInputClassName sets the static CSS class name of the input.
func (rm Remark) StaticInputClassName(value string) Remark {
	return rm.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static CSS class name of the label.
func (rm Remark) StaticLabelClassName(value string) Remark {
	return rm.set("staticLabelClassName", value)
}

// StaticOn sets the condition expression for static display.
func (rm Remark) StaticOn(value string) Remark {
	return rm.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder text for static display.
func (rm Remark) StaticPlaceholder(value string) Remark {
	return rm.set("staticPlaceholder", value)
}

// StaticSchema sets the schema.Schema for static display.
func (rm Remark) StaticSchema(value string) Remark {
	return rm.set("staticSchema", value)
}

// Style sets custom inline styles.
func (rm Remark) Style(value any) Remark {
	return rm.set("style", value)
}

// TestIdBuilder configures test ID generation.
func (rm Remark) TestIdBuilder(value string) Remark {
	return rm.set("testIdBuilder", value)
}

// Testid sets a specific test identifier.
func (rm Remark) Testid(value string) Remark {
	return rm.set("testid", value)
}

// Title sets the title of the remark.
func (rm Remark) Title(value string) Remark {
	return rm.set("title", value)
}

// Body sets the body of the remark.
func (rm Remark) Body(value string) Remark {
	return rm.set("body", value)
}

// TooltipClassName sets the tooltip CSS class name.
func (rm Remark) TooltipClassName(value string) Remark {
	return rm.set("tooltipClassName", value)
}

// Trigger sets the trigger rule of the remark.
func (rm Remark) Trigger(value string) Remark {
	return rm.set("trigger", value)
}

// UseMobileUI sets whether to use mobile UI.
func (rm Remark) UseMobileUI(value bool) Remark {
	return rm.set("useMobileUI", value)
}

// Visible sets whether the remark is visible.
func (rm Remark) Visible(value bool) Remark {
	return rm.set("visible", value)
}

// VisibleOn sets whether the remark is visible by expression.
func (rm Remark) VisibleOn(value string) Remark {
	return rm.set("visibleOn", value)
}
