package comp

// paginationWrapper represents the amis paginationWrapper renderer

type paginationWrapper Schema

// PaginationWrapper creates a new PaginationWrapper instance
func PaginationWrapper() paginationWrapper {
	return paginationWrapper{}.set("type", "pagination-wrapper")
}

// set sets a field value
func (pw paginationWrapper) set(key string, value any) paginationWrapper {
	pw[key] = value
	return pw
}

// Body sets the body content
func (pw paginationWrapper) Body(value ...any) paginationWrapper {
	return pw.set("body", value)
}

// ClassName sets the container CSS class name
func (pw paginationWrapper) ClassName(value string) paginationWrapper {
	return pw.set("className", value)
}

// Disabled sets the disabled state
func (pw paginationWrapper) Disabled(value bool) paginationWrapper {
	return pw.set("disabled", value)
}

// DisabledOn sets the disabled expression
func (pw paginationWrapper) DisabledOn(value string) paginationWrapper {
	return pw.set("disabledOn", value)
}

// EditorSetting sets the editor configuration, ignored at runtime
func (pw paginationWrapper) EditorSetting(value string) paginationWrapper {
	return pw.set("editorSetting", value)
}

// Hidden sets the hidden state
func (pw paginationWrapper) Hidden(value bool) paginationWrapper {
	return pw.set("hidden", value)
}

// HiddenOn sets the hidden expression
func (pw paginationWrapper) HiddenOn(value string) paginationWrapper {
	return pw.set("hiddenOn", value)
}

// ID sets the unique component ID
func (pw paginationWrapper) ID(value string) paginationWrapper {
	return pw.set("id", value)
}

// InputName sets the input field name
func (pw paginationWrapper) InputName(value string) paginationWrapper {
	return pw.set("inputName", value)
}

// MaxButtons sets the maximum number of pagination buttons
func (pw paginationWrapper) MaxButtons(value string) paginationWrapper {
	return pw.set("maxButtons", value)
}

// OnEvent sets the event action configuration
func (pw paginationWrapper) OnEvent(value any) paginationWrapper {
	return pw.set("onEvent", value)
}

// OutputName sets the output field name
func (pw paginationWrapper) OutputName(value string) paginationWrapper {
	return pw.set("outputName", value)
}

// PerPage sets the number of items per page
func (pw paginationWrapper) PerPage(value string) paginationWrapper {
	return pw.set("perPage", value)
}

// Position sets the pagination position
func (pw paginationWrapper) Position(value string) paginationWrapper {
	return pw.set("position", value)
}

// ShowPageInput sets whether to show the quick jump input box
func (pw paginationWrapper) ShowPageInput(value bool) paginationWrapper {
	return pw.set("showPageInput", value)
}

// Static sets the static display state
func (pw paginationWrapper) Static(value bool) paginationWrapper {
	return pw.set("static", value)
}

// StaticClassName sets the static display form item class name
func (pw paginationWrapper) StaticClassName(value string) paginationWrapper {
	return pw.set("staticClassName", value)
}

// StaticInputClassName sets the static display form item value class name
func (pw paginationWrapper) StaticInputClassName(value string) paginationWrapper {
	return pw.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static display form item label class name
func (pw paginationWrapper) StaticLabelClassName(value string) paginationWrapper {
	return pw.set("staticLabelClassName", value)
}

// StaticOn sets the static display expression
func (pw paginationWrapper) StaticOn(value string) paginationWrapper {
	return pw.set("staticOn", value)
}

// StaticPlaceholder sets the static display placeholder
func (pw paginationWrapper) StaticPlaceholder(value string) paginationWrapper {
	return pw.set("staticPlaceholder", value)
}

// StaticSchema sets the static display schema
func (pw paginationWrapper) StaticSchema(value string) paginationWrapper {
	return pw.set("staticSchema", value)
}

// Style sets the component style
func (pw paginationWrapper) Style(value any) paginationWrapper {
	return pw.set("style", value)
}

// TestIdBuilder sets the custom test ID builder
func (pw paginationWrapper) TestIdBuilder(value string) paginationWrapper {
	return pw.set("testIdBuilder", value)
}

// Testid sets the test ID
func (pw paginationWrapper) Testid(value string) paginationWrapper {
	return pw.set("testid", value)
}

// UseMobileUI sets whether to use mobile UI
func (pw paginationWrapper) UseMobileUI(value bool) paginationWrapper {
	return pw.set("useMobileUI", value)
}

// Visible sets the visibility state
func (pw paginationWrapper) Visible(value bool) paginationWrapper {
	return pw.set("visible", value)
}

// VisibleOn sets the visibility expression
func (pw paginationWrapper) VisibleOn(value string) paginationWrapper {
	return pw.set("visibleOn", value)
}
