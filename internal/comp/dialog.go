package comp

import "github.com/zrcoder/amisgo/schema"

type Dialog schema.Schema

func NewDialog() Dialog {
	return Dialog{"type": "dialog"}
}

func (d Dialog) set(key string, value any) Dialog {
	d[key] = value
	return d
}

// Actions defines the actions of the dialog.
func (d Dialog) Actions(value ...Action) Dialog {
	return d.set("actions", value)
}

// Body sets the body of the dialog.
func (d Dialog) Body(value ...any) Dialog {
	return d.set("body", value)
}

// BodyClassName sets the class name of the dialog body.
func (d Dialog) BodyClassName(value string) Dialog {
	return d.set("bodyClassName", value)
}

// ClassName sets the class name of the dialog container.
func (d Dialog) ClassName(value string) Dialog {
	return d.set("className", value)
}

// CloseOnEsc sets whether the dialog can be closed by pressing the ESC key.
func (d Dialog) CloseOnEsc(value bool) Dialog {
	return d.set("closeOnEsc", value)
}

// CloseOnOutside sets whether the dialog can be closed by clicking outside the dialog.
func (d Dialog) CloseOnOutside(value bool) Dialog {
	return d.set("closeOnOutside", value)
}

// Confirm sets whether the dialog needs to be confirmed before it can be closed.
func (d Dialog) Confirm(value bool) Dialog {
	return d.set("confirm", value)
}

// Data sets the data of the dialog.
func (d Dialog) Data(value any) Dialog {
	return d.set("data", value)
}

// DialogType sets the type of the dialog.
func (d Dialog) DialogType(value string) Dialog {
	return d.set("dialogType", value)
}

// Disabled sets whether the dialog is disabled.
func (d Dialog) Disabled(value bool) Dialog {
	return d.set("disabled", value)
}

// DisabledOn sets the disabled expression of the dialog.
func (d Dialog) DisabledOn(value string) Dialog {
	return d.set("disabledOn", value)
}

// Draggable sets whether the dialog can be dragged.
func (d Dialog) Draggable(value bool) Dialog {
	return d.set("draggable", value)
}

// EditorSetting sets the editor setting of the dialog.
func (d Dialog) EditorSetting(value string) Dialog {
	return d.set("editorSetting", value)
}

// Footer sets the footer of the dialog.
func (d Dialog) Footer(value string) Dialog {
	return d.set("footer", value)
}

// Header sets the header of the dialog.
func (d Dialog) Header(value string) Dialog {
	return d.set("header", value)
}

// HeaderClassName sets the class name of the dialog header.
func (d Dialog) HeaderClassName(value string) Dialog {
	return d.set("headerClassName", value)
}

// Height sets the height of the dialog.
func (d Dialog) Height(value string) Dialog {
	return d.set("height", value)
}

// Hidden sets whether the dialog is hidden.
func (d Dialog) Hidden(value bool) Dialog {
	return d.set("hidden", value)
}

// HiddenOn sets the hidden expression of the dialog.
func (d Dialog) HiddenOn(value string) Dialog {
	return d.set("hiddenOn", value)
}

// ID sets the ID of the dialog.
func (d Dialog) ID(value string) Dialog {
	return d.set("id", value)
}

// InputParams sets the input parameters of the dialog.
func (d Dialog) InputParams(value string) Dialog {
	return d.set("inputParams", value)
}

// Name sets the name of the dialog.
func (d Dialog) Name(value string) Dialog {
	return d.set("name", value)
}

// OnEvent sets the event handler of the dialog.
func (d Dialog) OnEvent(value Event) Dialog {
	return d.set("onEvent", value)
}

// Overlay sets whether the dialog has an overlay.
func (d Dialog) Overlay(value bool) Dialog {
	return d.set("overlay", value)
}

// ShowCloseButton sets whether the dialog has a close button.
func (d Dialog) ShowCloseButton(value bool) Dialog {
	return d.set("showCloseButton", value)
}

// ShowErrorMsg sets whether the dialog shows error messages.
func (d Dialog) ShowErrorMsg(value bool) Dialog {
	return d.set("showErrorMsg", value)
}

// ShowLoading sets whether the dialog shows a loading indicator.
func (d Dialog) ShowLoading(value bool) Dialog {
	return d.set("showLoading", value)
}

// Size sets the size of the dialog. supported values: xs、sm、md、lg、xl、full
func (d Dialog) Size(value string) Dialog {
	return d.set("size", value)
}

// Static sets whether the dialog is static.
func (d Dialog) Static(value bool) Dialog {
	return d.set("static", value)
}

// StaticClassName sets the class name of the static dialog.
func (d Dialog) StaticClassName(value string) Dialog {
	return d.set("staticClassName", value)
}

// StaticInputClassName sets the class name of the static dialog input.
func (d Dialog) StaticInputClassName(value string) Dialog {
	return d.set("staticInputClassName", value)
}

// StaticLabelClassName sets the class name of the static dialog label.
func (d Dialog) StaticLabelClassName(value string) Dialog {
	return d.set("staticLabelClassName", value)
}

// StaticOn sets the static expression of the dialog.
func (d Dialog) StaticOn(value string) Dialog {
	return d.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder of the static dialog.
func (d Dialog) StaticPlaceholder(value string) Dialog {
	return d.set("staticPlaceholder", value)
}

// StaticSchema sets the schema.Schema of the static dialog.
func (d Dialog) StaticSchema(value string) Dialog {
	return d.set("staticSchema", value)
}

// Style sets the style of the dialog.
func (d Dialog) Style(value any) Dialog {
	return d.set("style", value)
}

// TestIdBuilder sets the test ID builder of the dialog.
func (d Dialog) TestIdBuilder(value string) Dialog {
	return d.set("testIdBuilder", value)
}

// Testid sets the test ID of the dialog.
func (d Dialog) Testid(value string) Dialog {
	return d.set("testid", value)
}

// Title sets the title of the dialog.
func (d Dialog) Title(value any) Dialog {
	return d.set("title", value)
}

// UseMobileUI sets whether the dialog uses the mobile UI.
func (d Dialog) UseMobileUI(value bool) Dialog {
	return d.set("useMobileUI", value)
}

// Visible sets whether the dialog is visible.
func (d Dialog) Visible(value bool) Dialog {
	return d.set("visible", value)
}

// VisibleOn sets the visible expression of the dialog.
func (d Dialog) VisibleOn(value string) Dialog {
	return d.set("visibleOn", value)
}

// Width sets the width of the dialog.
func (d Dialog) Width(value string) Dialog {
	return d.set("width", value)
}
