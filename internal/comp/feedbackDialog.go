package comp

import "github.com/zrcoder/amisgo/schema"

// FeedbackDialog represents a feedback dialog
type FeedbackDialog schema.Schema

func NewFeedbackDialog() FeedbackDialog {
	return make(FeedbackDialog)
}

func (f FeedbackDialog) set(key string, value any) FeedbackDialog {
	f[key] = value
	return f
}

// Actions sets the actions, default buttons are confirm and cancel
func (f FeedbackDialog) Actions(value string) FeedbackDialog {
	return f.set("actions", value)
}

// Body sets the content area
func (f FeedbackDialog) Body(value ...any) FeedbackDialog {
	return f.set("body", value)
}

// BodyClassName sets the class name for the body container
func (f FeedbackDialog) BodyClassName(value string) FeedbackDialog {
	return f.set("bodyClassName", value)
}

// ClassName sets the CSS class name for the container
func (f FeedbackDialog) ClassName(value string) FeedbackDialog {
	return f.set("className", value)
}

// CloseOnEsc sets whether the dialog can be closed with the ESC key
func (f FeedbackDialog) CloseOnEsc(value bool) FeedbackDialog {
	return f.set("closeOnEsc", value)
}

// CloseOnOutside sets whether the dialog can be closed by clicking outside
func (f FeedbackDialog) CloseOnOutside(value bool) FeedbackDialog {
	return f.set("closeOnOutside", value)
}

// Confirm sets whether the confirm button is enabled
func (f FeedbackDialog) Confirm(value bool) FeedbackDialog {
	return f.set("confirm", value)
}

// Data sets the data mapping
func (f FeedbackDialog) Data(value any) FeedbackDialog {
	return f.set("data", value)
}

// DialogType sets the dialog type, e.g., confirm
func (f FeedbackDialog) DialogType(value string) FeedbackDialog {
	return f.set("dialogType", value)
}

// Disabled sets whether the dialog is disabled
func (f FeedbackDialog) Disabled(value bool) FeedbackDialog {
	return f.set("disabled", value)
}

// DisabledOn sets the expression to disable the dialog
func (f FeedbackDialog) DisabledOn(value string) FeedbackDialog {
	return f.set("disabledOn", value)
}

// Draggable sets whether the dialog is draggable
func (f FeedbackDialog) Draggable(value bool) FeedbackDialog {
	return f.set("draggable", value)
}

// EditorSetting sets the editor configuration, ignored at runtime
func (f FeedbackDialog) EditorSetting(value string) FeedbackDialog {
	return f.set("editorSetting", value)
}

// Footer sets the footer content
func (f FeedbackDialog) Footer(value string) FeedbackDialog {
	return f.set("footer", value)
}

// Header sets the header content
func (f FeedbackDialog) Header(value string) FeedbackDialog {
	return f.set("header", value)
}

// HeaderClassName sets the class name for the header
func (f FeedbackDialog) HeaderClassName(value string) FeedbackDialog {
	return f.set("headerClassName", value)
}

// Height sets the dialog height
func (f FeedbackDialog) Height(value string) FeedbackDialog {
	return f.set("height", value)
}

// Hidden sets whether the dialog is hidden
func (f FeedbackDialog) Hidden(value bool) FeedbackDialog {
	return f.set("hidden", value)
}

// HiddenOn sets the expression to hide the dialog
func (f FeedbackDialog) HiddenOn(value string) FeedbackDialog {
	return f.set("hiddenOn", value)
}

// ID sets the unique ID for the component, mainly for logging
func (f FeedbackDialog) ID(value string) FeedbackDialog {
	return f.set("id", value)
}

// InputParams sets the dialog parameters, formatted as JSONSchema
func (f FeedbackDialog) InputParams(value string) FeedbackDialog {
	return f.set("inputParams", value)
}

// Name sets the component name, used for locating and communication
func (f FeedbackDialog) Name(value string) FeedbackDialog {
	return f.set("name", value)
}

// OnEvent sets the event action configuration
func (f FeedbackDialog) OnEvent(value any) FeedbackDialog {
	return f.set("onEvent", value)
}

// Overlay sets whether to show the overlay
func (f FeedbackDialog) Overlay(value bool) FeedbackDialog {
	return f.set("overlay", value)
}

// ShowCloseButton sets whether to show the close button
func (f FeedbackDialog) ShowCloseButton(value bool) FeedbackDialog {
	return f.set("showCloseButton", value)
}

// ShowErrorMsg sets whether to show error messages
func (f FeedbackDialog) ShowErrorMsg(value bool) FeedbackDialog {
	return f.set("showErrorMsg", value)
}

// ShowLoading sets whether to show the loading spinner
func (f FeedbackDialog) ShowLoading(value bool) FeedbackDialog {
	return f.set("showLoading", value)
}

// Size sets the dialog size
func (f FeedbackDialog) Size(value string) FeedbackDialog {
	return f.set("size", value)
}

// SkipRestOnCancel sets whether to interrupt subsequent operations on cancel
func (f FeedbackDialog) SkipRestOnCancel(value bool) FeedbackDialog {
	return f.set("skipRestOnCancel", value)
}

// SkipRestOnConfirm sets whether to interrupt subsequent operations on confirm
func (f FeedbackDialog) SkipRestOnConfirm(value bool) FeedbackDialog {
	return f.set("skipRestOnConfirm", value)
}

// Static sets whether the dialog is displayed statically
func (f FeedbackDialog) Static(value bool) FeedbackDialog {
	return f.set("static", value)
}

// StaticClassName sets the class name for static display form items
func (f FeedbackDialog) StaticClassName(value string) FeedbackDialog {
	return f.set("staticClassName", value)
}

// StaticInputClassName sets the class name for static display form item values
func (f FeedbackDialog) StaticInputClassName(value string) FeedbackDialog {
	return f.set("staticInputClassName", value)
}

// StaticLabelClassName sets the class name for static display form item labels
func (f FeedbackDialog) StaticLabelClassName(value string) FeedbackDialog {
	return f.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display
func (f FeedbackDialog) StaticOn(value string) FeedbackDialog {
	return f.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display empty values
func (f FeedbackDialog) StaticPlaceholder(value string) FeedbackDialog {
	return f.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema.Schema
func (f FeedbackDialog) StaticSchema(value string) FeedbackDialog {
	return f.set("staticSchema", value)
}

// Style sets the component style
func (f FeedbackDialog) Style(value any) FeedbackDialog {
	return f.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (f FeedbackDialog) TestIdBuilder(value string) FeedbackDialog {
	return f.set("testIdBuilder", value)
}

// Testid sets the test ID
func (f FeedbackDialog) Testid(value string) FeedbackDialog {
	return f.set("testid", value)
}

// Title sets the dialog title
func (f FeedbackDialog) Title(value any) FeedbackDialog {
	return f.set("title", value)
}

// UseMobileUI sets whether to disable mobile UI styles at the component level
func (f FeedbackDialog) UseMobileUI(value bool) FeedbackDialog {
	return f.set("useMobileUI", value)
}

// Visible sets whether the dialog is visible
func (f FeedbackDialog) Visible(value bool) FeedbackDialog {
	return f.set("visible", value)
}

// VisibleOn sets the condition for the dialog to appear
func (f FeedbackDialog) VisibleOn(value string) FeedbackDialog {
	return f.set("visibleOn", value)
}

// Width sets the dialog width
func (f FeedbackDialog) Width(value string) FeedbackDialog {
	return f.set("width", value)
}
