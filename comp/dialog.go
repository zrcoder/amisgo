package comp

import "github.com/zrcoder/amisgo/model"

type dialog model.Schema

func Dialog() dialog {
	return make(dialog).set("type", "dialog")
}

func (d dialog) set(key string, value any) dialog {
	d[key] = value
	return d
}

// Actions defines the actions of the dialog.
func (d dialog) Actions(value ...action) dialog {
	return d.set("actions", value)
}

// Body sets the body of the dialog.
func (d dialog) Body(value ...any) dialog {
	return d.set("body", value)
}

// BodyClassName sets the class name of the dialog body.
func (d dialog) BodyClassName(value string) dialog {
	return d.set("bodyClassName", value)
}

// ClassName sets the class name of the dialog container.
func (d dialog) ClassName(value string) dialog {
	return d.set("className", value)
}

// CloseOnEsc sets whether the dialog can be closed by pressing the ESC key.
func (d dialog) CloseOnEsc(value bool) dialog {
	return d.set("closeOnEsc", value)
}

// CloseOnOutside sets whether the dialog can be closed by clicking outside the dialog.
func (d dialog) CloseOnOutside(value bool) dialog {
	return d.set("closeOnOutside", value)
}

// Confirm sets whether the dialog needs to be confirmed before it can be closed.
func (d dialog) Confirm(value bool) dialog {
	return d.set("confirm", value)
}

// Data sets the data of the dialog.
func (d dialog) Data(value model.Data) dialog {
	return d.set("data", value)
}

// DialogType sets the type of the dialog.
func (d dialog) DialogType(value string) dialog {
	return d.set("dialogType", value)
}

// Disabled sets whether the dialog is disabled.
func (d dialog) Disabled(value bool) dialog {
	return d.set("disabled", value)
}

// DisabledOn sets the disabled expression of the dialog.
func (d dialog) DisabledOn(value string) dialog {
	return d.set("disabledOn", value)
}

// Draggable sets whether the dialog can be dragged.
func (d dialog) Draggable(value bool) dialog {
	return d.set("draggable", value)
}

// EditorSetting sets the editor setting of the dialog.
func (d dialog) EditorSetting(value string) dialog {
	return d.set("editorSetting", value)
}

// Footer sets the footer of the dialog.
func (d dialog) Footer(value string) dialog {
	return d.set("footer", value)
}

// Header sets the header of the dialog.
func (d dialog) Header(value string) dialog {
	return d.set("header", value)
}

// HeaderClassName sets the class name of the dialog header.
func (d dialog) HeaderClassName(value string) dialog {
	return d.set("headerClassName", value)
}

// Height sets the height of the dialog.
func (d dialog) Height(value string) dialog {
	return d.set("height", value)
}

// Hidden sets whether the dialog is hidden.
func (d dialog) Hidden(value bool) dialog {
	return d.set("hidden", value)
}

// HiddenOn sets the hidden expression of the dialog.
func (d dialog) HiddenOn(value string) dialog {
	return d.set("hiddenOn", value)
}

// ID sets the ID of the dialog.
func (d dialog) ID(value string) dialog {
	return d.set("id", value)
}

// InputParams sets the input parameters of the dialog.
func (d dialog) InputParams(value string) dialog {
	return d.set("inputParams", value)
}

// Name sets the name of the dialog.
func (d dialog) Name(value string) dialog {
	return d.set("name", value)
}

// OnEvent sets the event handler of the dialog.
func (d dialog) OnEvent(value any) dialog {
	return d.set("onEvent", value)
}

// Overlay sets whether the dialog has an overlay.
func (d dialog) Overlay(value bool) dialog {
	return d.set("overlay", value)
}

// ShowCloseButton sets whether the dialog has a close button.
func (d dialog) ShowCloseButton(value bool) dialog {
	return d.set("showCloseButton", value)
}

// ShowErrorMsg sets whether the dialog shows error messages.
func (d dialog) ShowErrorMsg(value bool) dialog {
	return d.set("showErrorMsg", value)
}

// ShowLoading sets whether the dialog shows a loading indicator.
func (d dialog) ShowLoading(value bool) dialog {
	return d.set("showLoading", value)
}

// Size sets the size of the dialog.
func (d dialog) Size(value string) dialog {
	return d.set("size", value)
}

// Static sets whether the dialog is static.
func (d dialog) Static(value bool) dialog {
	return d.set("static", value)
}

// StaticClassName sets the class name of the static dialog.
func (d dialog) StaticClassName(value string) dialog {
	return d.set("staticClassName", value)
}

// StaticInputClassName sets the class name of the static dialog input.
func (d dialog) StaticInputClassName(value string) dialog {
	return d.set("staticInputClassName", value)
}

// StaticLabelClassName sets the class name of the static dialog label.
func (d dialog) StaticLabelClassName(value string) dialog {
	return d.set("staticLabelClassName", value)
}

// StaticOn sets the static expression of the dialog.
func (d dialog) StaticOn(value string) dialog {
	return d.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder of the static dialog.
func (d dialog) StaticPlaceholder(value string) dialog {
	return d.set("staticPlaceholder", value)
}

// StaticSchema sets the schema of the static dialog.
func (d dialog) StaticSchema(value string) dialog {
	return d.set("staticSchema", value)
}

// Style sets the style of the dialog.
func (d dialog) Style(value any) dialog {
	return d.set("style", value)
}

// TestIdBuilder sets the test ID builder of the dialog.
func (d dialog) TestIdBuilder(value string) dialog {
	return d.set("testIdBuilder", value)
}

// Testid sets the test ID of the dialog.
func (d dialog) Testid(value string) dialog {
	return d.set("testid", value)
}

// Title sets the title of the dialog.
func (d dialog) Title(value any) dialog {
	return d.set("title", value)
}

// UseMobileUI sets whether the dialog uses the mobile UI.
func (d dialog) UseMobileUI(value bool) dialog {
	return d.set("useMobileUI", value)
}

// Visible sets whether the dialog is visible.
func (d dialog) Visible(value bool) dialog {
	return d.set("visible", value)
}

// VisibleOn sets the visible expression of the dialog.
func (d dialog) VisibleOn(value string) dialog {
	return d.set("visibleOn", value)
}

// Width sets the width of the dialog.
func (d dialog) Width(value string) dialog {
	return d.set("width", value)
}
