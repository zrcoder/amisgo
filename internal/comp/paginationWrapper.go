package comp

import "github.com/zrcoder/amisgo/schema"

// PaginationWrapper represents the amis PaginationWrapper renderer
type PaginationWrapper schema.Schema

func NewPaginationWrapper() PaginationWrapper {
	return PaginationWrapper{"type": "pagination-wrapper"}
}

// set sets a field value
func (pw PaginationWrapper) set(key string, value any) PaginationWrapper {
	pw[key] = value
	return pw
}

// Body sets the body content
func (pw PaginationWrapper) Body(value ...any) PaginationWrapper {
	return pw.set("body", value)
}

// ClassName sets the container CSS class name
func (pw PaginationWrapper) ClassName(value string) PaginationWrapper {
	return pw.set("className", value)
}

// Disabled sets the disabled state
func (pw PaginationWrapper) Disabled(value bool) PaginationWrapper {
	return pw.set("disabled", value)
}

// DisabledOn sets the disabled expression
func (pw PaginationWrapper) DisabledOn(value string) PaginationWrapper {
	return pw.set("disabledOn", value)
}

// EditorSetting sets the editor configuration, ignored at runtime
func (pw PaginationWrapper) EditorSetting(value string) PaginationWrapper {
	return pw.set("editorSetting", value)
}

// Hidden sets the hidden state
func (pw PaginationWrapper) Hidden(value bool) PaginationWrapper {
	return pw.set("hidden", value)
}

// HiddenOn sets the hidden expression
func (pw PaginationWrapper) HiddenOn(value string) PaginationWrapper {
	return pw.set("hiddenOn", value)
}

// ID sets the unique component ID
func (pw PaginationWrapper) ID(value string) PaginationWrapper {
	return pw.set("id", value)
}

// InputName sets the input field name
func (pw PaginationWrapper) InputName(value string) PaginationWrapper {
	return pw.set("inputName", value)
}

// MaxButtons sets the maximum number of pagination buttons
func (pw PaginationWrapper) MaxButtons(value string) PaginationWrapper {
	return pw.set("maxButtons", value)
}

// OnEvent sets the event action configuration
func (pw PaginationWrapper) OnEvent(value any) PaginationWrapper {
	return pw.set("onEvent", value)
}

// OutputName sets the output field name
func (pw PaginationWrapper) OutputName(value string) PaginationWrapper {
	return pw.set("outputName", value)
}

// PerPage sets the number of items per page
func (pw PaginationWrapper) PerPage(value string) PaginationWrapper {
	return pw.set("perPage", value)
}

// Position sets the pagination position
func (pw PaginationWrapper) Position(value string) PaginationWrapper {
	return pw.set("position", value)
}

// ShowPageInput sets whether to show the quick jump input box
func (pw PaginationWrapper) ShowPageInput(value bool) PaginationWrapper {
	return pw.set("showPageInput", value)
}

// Static sets the static display state
func (pw PaginationWrapper) Static(value bool) PaginationWrapper {
	return pw.set("static", value)
}

// StaticClassName sets the static display form item class name
func (pw PaginationWrapper) StaticClassName(value string) PaginationWrapper {
	return pw.set("staticClassName", value)
}

// StaticInputClassName sets the static display form item value class name
func (pw PaginationWrapper) StaticInputClassName(value string) PaginationWrapper {
	return pw.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static display form item label class name
func (pw PaginationWrapper) StaticLabelClassName(value string) PaginationWrapper {
	return pw.set("staticLabelClassName", value)
}

// StaticOn sets the static display expression
func (pw PaginationWrapper) StaticOn(value string) PaginationWrapper {
	return pw.set("staticOn", value)
}

// StaticPlaceholder sets the static display placeholder
func (pw PaginationWrapper) StaticPlaceholder(value string) PaginationWrapper {
	return pw.set("staticPlaceholder", value)
}

// StaticSchema sets the static display schema.Schema
func (pw PaginationWrapper) StaticSchema(value string) PaginationWrapper {
	return pw.set("staticSchema", value)
}

// Style sets the component style
func (pw PaginationWrapper) Style(value any) PaginationWrapper {
	return pw.set("style", value)
}

// TestIdBuilder sets the custom test ID builder
func (pw PaginationWrapper) TestIdBuilder(value string) PaginationWrapper {
	return pw.set("testIdBuilder", value)
}

// Testid sets the test ID
func (pw PaginationWrapper) Testid(value string) PaginationWrapper {
	return pw.set("testid", value)
}

// UseMobileUI sets whether to use mobile UI
func (pw PaginationWrapper) UseMobileUI(value bool) PaginationWrapper {
	return pw.set("useMobileUI", value)
}

// Visible sets the visibility state
func (pw PaginationWrapper) Visible(value bool) PaginationWrapper {
	return pw.set("visible", value)
}

// VisibleOn sets the visibility expression
func (pw PaginationWrapper) VisibleOn(value string) PaginationWrapper {
	return pw.set("visibleOn", value)
}
