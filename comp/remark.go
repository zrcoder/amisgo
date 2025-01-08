package comp

import "github.com/zrcoder/amisgo/model"

// remark 提示渲染器.

type remark model.Schema

// Remark creates a new remark.
func Remark() remark {
	return remark{}.set("type", "remark")
}

func (rm remark) set(key string, value any) remark {
	rm[key] = value
	return rm
}

// ClassName sets the container's CSS class name.
func (rm remark) ClassName(value string) remark {
	return rm.set("className", value)
}

// Content sets the content of the remark.
func (rm remark) Content(value string) remark {
	return rm.set("content", value)
}

// Disabled sets whether the remark is disabled.
func (rm remark) Disabled(value bool) remark {
	return rm.set("disabled", value)
}

// DisabledOn sets whether the remark is disabled by expression.
func (rm remark) DisabledOn(value string) remark {
	return rm.set("disabledOn", value)
}

// EditorSetting sets the editor setting.
func (rm remark) EditorSetting(value string) remark {
	return rm.set("editorSetting", value)
}

// Hidden sets whether the remark is hidden.
func (rm remark) Hidden(value bool) remark {
	return rm.set("hidden", value)
}

// HiddenOn sets whether the remark is hidden by expression.
func (rm remark) HiddenOn(value string) remark {
	return rm.set("hiddenOn", value)
}

// Icon sets the icon of the remark.
func (rm remark) Icon(value string) remark {
	return rm.set("icon", value)
}

// ID sets the unique ID of the remark.
func (rm remark) ID(value string) remark {
	return rm.set("id", value)
}

// Label sets the label of the remark.
func (rm remark) Label(value string) remark {
	return rm.set("label", value)
}

// OnEvent sets the event handler of the remark.
func (rm remark) OnEvent(value any) remark {
	return rm.set("onEvent", value)
}

// Placement sets the placement of the remark.
func (rm remark) Placement(value string) remark {
	return rm.set("placement", value)
}

// RootClose sets whether to close the remark when clicking outside.
func (rm remark) RootClose(value bool) remark {
	return rm.set("rootClose", value)
}

// Shape sets the shape of the icon.
func (rm remark) Shape(value string) remark {
	return rm.set("shape", value)
}

// Static sets whether the remark is static.
func (rm remark) Static(value bool) remark {
	return rm.set("static", value)
}

// StaticClassName sets the static CSS class name of the remark.
func (rm remark) StaticClassName(value string) remark {
	return rm.set("staticClassName", value)
}

// StaticInputClassName sets the static CSS class name of the input.
func (rm remark) StaticInputClassName(value string) remark {
	return rm.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static CSS class name of the label.
func (rm remark) StaticLabelClassName(value string) remark {
	return rm.set("staticLabelClassName", value)
}

// StaticOn sets the condition expression for static display.
func (rm remark) StaticOn(value string) remark {
	return rm.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder text for static display.
func (rm remark) StaticPlaceholder(value string) remark {
	return rm.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display.
func (rm remark) StaticSchema(value string) remark {
	return rm.set("staticSchema", value)
}

// Style sets custom inline styles.
func (rm remark) Style(value any) remark {
	return rm.set("style", value)
}

// TestIdBuilder configures test ID generation.
func (rm remark) TestIdBuilder(value string) remark {
	return rm.set("testIdBuilder", value)
}

// Testid sets a specific test identifier.
func (rm remark) Testid(value string) remark {
	return rm.set("testid", value)
}

// Title sets the title of the remark.
func (rm remark) Title(value string) remark {
	return rm.set("title", value)
}

// Body sets the body of the remark.
func (rm remark) Body(value string) remark {
	return rm.set("body", value)
}

// TooltipClassName sets the tooltip CSS class name.
func (rm remark) TooltipClassName(value string) remark {
	return rm.set("tooltipClassName", value)
}

// Trigger sets the trigger rule of the remark.
func (rm remark) Trigger(value string) remark {
	return rm.set("trigger", value)
}

// UseMobileUI sets whether to use mobile UI.
func (rm remark) UseMobileUI(value bool) remark {
	return rm.set("useMobileUI", value)
}

// Visible sets whether the remark is visible.
func (rm remark) Visible(value bool) remark {
	return rm.set("visible", value)
}

// VisibleOn sets whether the remark is visible by expression.
func (rm remark) VisibleOn(value string) remark {
	return rm.set("visibleOn", value)
}
