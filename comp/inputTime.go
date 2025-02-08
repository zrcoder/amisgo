package comp

import "github.com/zrcoder/amisgo/model"

// InputTime represents a time selection component. Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/form/time
type InputTime model.Schema

func NewInputTime() InputTime {
	return InputTime{"type": "input-time"}
}

// AutoFill sets autoFill value to synchronize other values in the form when an option is selected.
func (tc InputTime) AutoFill(value string) InputTime {
	return tc.set("autoFill", value)
}

// BorderMode sets the border mode. Options: full | half | none.
func (tc InputTime) BorderMode(value string) InputTime {
	return tc.set("borderMode", value)
}

// ClassName sets the container CSS class name.
func (tc InputTime) ClassName(value string) InputTime {
	return tc.set("className", value)
}

// ClearValueOnHidden removes the form item value when hidden.
func (tc InputTime) ClearValueOnHidden(value bool) InputTime {
	return tc.set("clearValueOnHidden", value)
}

// Clearable shows or hides the clear button.
func (tc InputTime) Clearable(value bool) InputTime {
	return tc.set("clearable", value)
}

// Desc sets the description.
func (tc InputTime) Desc(value string) InputTime {
	return tc.set("desc", value)
}

// Description sets the description content, supports HTML.
func (tc InputTime) Description(value string) InputTime {
	return tc.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description.
func (tc InputTime) DescriptionClassName(value string) InputTime {
	return tc.set("descriptionClassName", value)
}

// Disabled enables or disables the component.
func (tc InputTime) Disabled(value bool) InputTime {
	return tc.set("disabled", value)
}

// DisabledDate sets a function to disable specific dates.
func (tc InputTime) DisabledDate(value string) InputTime {
	return tc.set("disabledDate", value)
}

// DisabledOn sets an expression to disable the component.
func (tc InputTime) DisabledOn(value string) InputTime {
	return tc.set("disabledOn", value)
}

// DisplayFormat sets the display format for the date.
func (tc InputTime) DisplayFormat(value string) InputTime {
	return tc.set("displayFormat", value)
}

// EditorSetting sets the editor configuration.
func (tc InputTime) EditorSetting(value string) InputTime {
	return tc.set("editorSetting", value)
}

// Emebed sets whether the component is inline.
func (tc InputTime) Emebed(value bool) InputTime {
	return tc.set("emebed", value)
}

// ExtraName sets an additional field name for range components.
func (tc InputTime) ExtraName(value string) InputTime {
	return tc.set("extraName", value)
}

// Format sets the date storage format.
func (tc InputTime) Format(value string) InputTime {
	return tc.set("format", value)
}

// Hidden hides or shows the component.
func (tc InputTime) Hidden(value bool) InputTime {
	return tc.set("hidden", value)
}

// HiddenOn sets an expression to hide the component.
func (tc InputTime) HiddenOn(value string) InputTime {
	return tc.set("hiddenOn", value)
}

// Hint sets the input hint displayed on focus.
func (tc InputTime) Hint(value string) InputTime {
	return tc.set("hint", value)
}

// Horizontal sets the horizontal layout configuration.
func (tc InputTime) Horizontal(value string) InputTime {
	return tc.set("horizontal", value)
}

// ID sets a unique ID for the component, mainly for logging.
func (tc InputTime) ID(value string) InputTime {
	return tc.set("id", value)
}

// InitAutoFill sets the initial autoFill value.
func (tc InputTime) InitAutoFill(value string) InputTime {
	return tc.set("initAutoFill", value)
}

// Inline sets whether the form control is inline.
func (tc InputTime) Inline(value bool) InputTime {
	return tc.set("inline", value)
}

// InputClassName sets the CSS class name for the input.
func (tc InputTime) InputClassName(value string) InputTime {
	return tc.set("inputClassName", value)
}

// InputFormat sets the input format for the date.
func (tc InputTime) InputFormat(value string) InputTime {
	return tc.set("inputFormat", value)
}

// Label sets the label text.
func (tc InputTime) Label(value string) InputTime {
	return tc.set("label", value)
}

// LabelAlign sets the label alignment. Options: right | left | top | inherit.
func (tc InputTime) LabelAlign(value string) InputTime {
	return tc.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label.
func (tc InputTime) LabelClassName(value string) InputTime {
	return tc.set("labelClassName", value)
}

// LabelRemark sets a small icon with a tooltip next to the label.
func (tc InputTime) LabelRemark(value string) InputTime {
	return tc.set("labelRemark", value)
}

// LabelWidth sets a custom width for the label, default unit is px.
func (tc InputTime) LabelWidth(value string) InputTime {
	return tc.set("labelWidth", value)
}

// Mode sets the display mode for the form item. Options: normal | inline | horizontal.
func (tc InputTime) Mode(value string) InputTime {
	return tc.set("mode", value)
}

// Name sets the field name for form submission.
func (tc InputTime) Name(value string) InputTime {
	return tc.set("name", value)
}

// OnEvent sets the event action configuration.
func (tc InputTime) OnEvent(value any) InputTime {
	return tc.set("onEvent", value)
}

// Placeholder sets the placeholder text.
func (tc InputTime) Placeholder(value string) InputTime {
	return tc.set("placeholder", value)
}

// ReadOnly sets whether the component is read-only.
func (tc InputTime) ReadOnly(value bool) InputTime {
	return tc.set("readOnly", value)
}

// ReadOnlyOn sets an expression to make the component read-only.
func (tc InputTime) ReadOnlyOn(value string) InputTime {
	return tc.set("readOnlyOn", value)
}

// Remark sets a small icon with a tooltip.
func (tc InputTime) Remark(value string) InputTime {
	return tc.set("remark", value)
}

// Required sets whether the component is required.
func (tc InputTime) Required(value bool) InputTime {
	return tc.set("required", value)
}

// Row sets the row value.
func (tc InputTime) Row(value string) InputTime {
	return tc.set("row", value)
}

// SaveImmediately sets whether to save immediately in TableColumn.
func (tc InputTime) SaveImmediately(value bool) InputTime {
	return tc.set("saveImmediately", value)
}

// Shortcuts sets the date shortcuts.
func (tc InputTime) Shortcuts(value string) InputTime {
	return tc.set("shortcuts", value)
}

// Size sets the size of the form item. Options: xs | sm | md | lg | full.
func (tc InputTime) Size(value string) InputTime {
	return tc.set("size", value)
}

// Static sets whether the component is displayed statically.
func (tc InputTime) Static(value bool) InputTime {
	return tc.set("static", value)
}

// StaticClassName sets the CSS class name for static display.
func (tc InputTime) StaticClassName(value string) InputTime {
	return tc.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display.
func (tc InputTime) StaticInputClassName(value string) InputTime {
	return tc.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display.
func (tc InputTime) StaticLabelClassName(value string) InputTime {
	return tc.set("staticLabelClassName", value)
}

// StaticOn sets an expression for static display.
func (tc InputTime) StaticOn(value string) InputTime {
	return tc.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display.
func (tc InputTime) StaticPlaceholder(value string) InputTime {
	return tc.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema.
func (tc InputTime) StaticSchema(value string) InputTime {
	return tc.set("staticSchema", value)
}

// Style sets the component style.
func (tc InputTime) Style(value any) InputTime {
	return tc.set("style", value)
}

// SubmitOnChange sets whether to submit the form on change.
func (tc InputTime) SubmitOnChange(value bool) InputTime {
	return tc.set("submitOnChange", value)
}

// TestIdBuilder sets the test ID builder.
func (tc InputTime) TestIdBuilder(value string) InputTime {
	return tc.set("testIdBuilder", value)
}

// TimeConstraints sets the time input range constraints.
func (tc InputTime) TimeConstraints(value string) InputTime {
	return tc.set("timeConstraints", value)
}

// TimeFormat sets the time format.
func (tc InputTime) TimeFormat(value string) InputTime {
	return tc.set("timeFormat", value)
}

// UseMobileUI sets whether to use mobile UI styles.
func (tc InputTime) UseMobileUI(value bool) InputTime {
	return tc.set("useMobileUI", value)
}

// UTC sets whether to store UTC time.
func (tc InputTime) UTC(value bool) InputTime {
	return tc.set("utc", value)
}

// ValidateApi sets the remote validation API.
func (tc InputTime) ValidateApi(value string) InputTime {
	return tc.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on change.
func (tc InputTime) ValidateOnChange(value bool) InputTime {
	return tc.set("validateOnChange", value)
}

// ValidationErrors sets the validation error messages.
func (tc InputTime) ValidationErrors(value string) InputTime {
	return tc.set("validationErrors", value)
}

// Validations sets the validation rules.
func (tc InputTime) Validations(value string) InputTime {
	return tc.set("validations", value)
}

// Value sets the default value, must be static.
func (tc InputTime) Value(value string) InputTime {
	return tc.set("value", value)
}

// ValueFormat sets the value format, replacing format.
func (tc InputTime) ValueFormat(value string) InputTime {
	return tc.set("valueFormat", value)
}

// Visible sets whether the component is visible.
func (tc InputTime) Visible(value bool) InputTime {
	return tc.set("visible", value)
}

// VisibleOn sets an expression to control visibility.
func (tc InputTime) VisibleOn(value string) InputTime {
	return tc.set("visibleOn", value)
}

// Width sets the width in a table.
func (tc InputTime) Width(value string) InputTime {
	return tc.set("width", value)
}

func (tc InputTime) set(key string, value any) InputTime {
	tc[key] = value
	return tc
}
