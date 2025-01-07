package comp

// inputTime represents a time selection component. Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/form/time

type inputTime Schema

// InputTime creates a new input-time component.
func InputTime() inputTime {
	return inputTime{}.set("type", "input-time")
}

// AutoFill sets autoFill value to synchronize other values in the form when an option is selected.
func (tc inputTime) AutoFill(value string) inputTime {
	return tc.set("autoFill", value)
}

// BorderMode sets the border mode. Options: full | half | none.
func (tc inputTime) BorderMode(value string) inputTime {
	return tc.set("borderMode", value)
}

// ClassName sets the container CSS class name.
func (tc inputTime) ClassName(value string) inputTime {
	return tc.set("className", value)
}

// ClearValueOnHidden removes the form item value when hidden.
func (tc inputTime) ClearValueOnHidden(value bool) inputTime {
	return tc.set("clearValueOnHidden", value)
}

// Clearable shows or hides the clear button.
func (tc inputTime) Clearable(value bool) inputTime {
	return tc.set("clearable", value)
}

// Desc sets the description.
func (tc inputTime) Desc(value string) inputTime {
	return tc.set("desc", value)
}

// Description sets the description content, supports HTML.
func (tc inputTime) Description(value string) inputTime {
	return tc.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description.
func (tc inputTime) DescriptionClassName(value string) inputTime {
	return tc.set("descriptionClassName", value)
}

// Disabled enables or disables the component.
func (tc inputTime) Disabled(value bool) inputTime {
	return tc.set("disabled", value)
}

// DisabledDate sets a function to disable specific dates.
func (tc inputTime) DisabledDate(value string) inputTime {
	return tc.set("disabledDate", value)
}

// DisabledOn sets an expression to disable the component.
func (tc inputTime) DisabledOn(value string) inputTime {
	return tc.set("disabledOn", value)
}

// DisplayFormat sets the display format for the date.
func (tc inputTime) DisplayFormat(value string) inputTime {
	return tc.set("displayFormat", value)
}

// EditorSetting sets the editor configuration.
func (tc inputTime) EditorSetting(value string) inputTime {
	return tc.set("editorSetting", value)
}

// Emebed sets whether the component is inline.
func (tc inputTime) Emebed(value bool) inputTime {
	return tc.set("emebed", value)
}

// ExtraName sets an additional field name for range components.
func (tc inputTime) ExtraName(value string) inputTime {
	return tc.set("extraName", value)
}

// Format sets the date storage format.
func (tc inputTime) Format(value string) inputTime {
	return tc.set("format", value)
}

// Hidden hides or shows the component.
func (tc inputTime) Hidden(value bool) inputTime {
	return tc.set("hidden", value)
}

// HiddenOn sets an expression to hide the component.
func (tc inputTime) HiddenOn(value string) inputTime {
	return tc.set("hiddenOn", value)
}

// Hint sets the input hint displayed on focus.
func (tc inputTime) Hint(value string) inputTime {
	return tc.set("hint", value)
}

// Horizontal sets the horizontal layout configuration.
func (tc inputTime) Horizontal(value string) inputTime {
	return tc.set("horizontal", value)
}

// ID sets a unique ID for the component, mainly for logging.
func (tc inputTime) ID(value string) inputTime {
	return tc.set("id", value)
}

// InitAutoFill sets the initial autoFill value.
func (tc inputTime) InitAutoFill(value string) inputTime {
	return tc.set("initAutoFill", value)
}

// Inline sets whether the form control is inline.
func (tc inputTime) Inline(value bool) inputTime {
	return tc.set("inline", value)
}

// InputClassName sets the CSS class name for the input.
func (tc inputTime) InputClassName(value string) inputTime {
	return tc.set("inputClassName", value)
}

// InputFormat sets the input format for the date.
func (tc inputTime) InputFormat(value string) inputTime {
	return tc.set("inputFormat", value)
}

// Label sets the label text.
func (tc inputTime) Label(value string) inputTime {
	return tc.set("label", value)
}

// LabelAlign sets the label alignment. Options: right | left | top | inherit.
func (tc inputTime) LabelAlign(value string) inputTime {
	return tc.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label.
func (tc inputTime) LabelClassName(value string) inputTime {
	return tc.set("labelClassName", value)
}

// LabelRemark sets a small icon with a tooltip next to the label.
func (tc inputTime) LabelRemark(value string) inputTime {
	return tc.set("labelRemark", value)
}

// LabelWidth sets a custom width for the label, default unit is px.
func (tc inputTime) LabelWidth(value string) inputTime {
	return tc.set("labelWidth", value)
}

// Mode sets the display mode for the form item. Options: normal | inline | horizontal.
func (tc inputTime) Mode(value string) inputTime {
	return tc.set("mode", value)
}

// Name sets the field name for form submission.
func (tc inputTime) Name(value string) inputTime {
	return tc.set("name", value)
}

// OnEvent sets the event action configuration.
func (tc inputTime) OnEvent(value any) inputTime {
	return tc.set("onEvent", value)
}

// Placeholder sets the placeholder text.
func (tc inputTime) Placeholder(value string) inputTime {
	return tc.set("placeholder", value)
}

// ReadOnly sets whether the component is read-only.
func (tc inputTime) ReadOnly(value bool) inputTime {
	return tc.set("readOnly", value)
}

// ReadOnlyOn sets an expression to make the component read-only.
func (tc inputTime) ReadOnlyOn(value string) inputTime {
	return tc.set("readOnlyOn", value)
}

// Remark sets a small icon with a tooltip.
func (tc inputTime) Remark(value string) inputTime {
	return tc.set("remark", value)
}

// Required sets whether the component is required.
func (tc inputTime) Required(value bool) inputTime {
	return tc.set("required", value)
}

// Row sets the row value.
func (tc inputTime) Row(value string) inputTime {
	return tc.set("row", value)
}

// SaveImmediately sets whether to save immediately in TableColumn.
func (tc inputTime) SaveImmediately(value bool) inputTime {
	return tc.set("saveImmediately", value)
}

// Shortcuts sets the date shortcuts.
func (tc inputTime) Shortcuts(value string) inputTime {
	return tc.set("shortcuts", value)
}

// Size sets the size of the form item. Options: xs | sm | md | lg | full.
func (tc inputTime) Size(value string) inputTime {
	return tc.set("size", value)
}

// Static sets whether the component is displayed statically.
func (tc inputTime) Static(value bool) inputTime {
	return tc.set("static", value)
}

// StaticClassName sets the CSS class name for static display.
func (tc inputTime) StaticClassName(value string) inputTime {
	return tc.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display.
func (tc inputTime) StaticInputClassName(value string) inputTime {
	return tc.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display.
func (tc inputTime) StaticLabelClassName(value string) inputTime {
	return tc.set("staticLabelClassName", value)
}

// StaticOn sets an expression for static display.
func (tc inputTime) StaticOn(value string) inputTime {
	return tc.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display.
func (tc inputTime) StaticPlaceholder(value string) inputTime {
	return tc.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema.
func (tc inputTime) StaticSchema(value string) inputTime {
	return tc.set("staticSchema", value)
}

// Style sets the component style.
func (tc inputTime) Style(value any) inputTime {
	return tc.set("style", value)
}

// SubmitOnChange sets whether to submit the form on change.
func (tc inputTime) SubmitOnChange(value bool) inputTime {
	return tc.set("submitOnChange", value)
}

// TestIdBuilder sets the test ID builder.
func (tc inputTime) TestIdBuilder(value string) inputTime {
	return tc.set("testIdBuilder", value)
}

// TimeConstraints sets the time input range constraints.
func (tc inputTime) TimeConstraints(value string) inputTime {
	return tc.set("timeConstraints", value)
}

// TimeFormat sets the time format.
func (tc inputTime) TimeFormat(value string) inputTime {
	return tc.set("timeFormat", value)
}

// UseMobileUI sets whether to use mobile UI styles.
func (tc inputTime) UseMobileUI(value bool) inputTime {
	return tc.set("useMobileUI", value)
}

// UTC sets whether to store UTC time.
func (tc inputTime) UTC(value bool) inputTime {
	return tc.set("utc", value)
}

// ValidateApi sets the remote validation API.
func (tc inputTime) ValidateApi(value string) inputTime {
	return tc.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on change.
func (tc inputTime) ValidateOnChange(value bool) inputTime {
	return tc.set("validateOnChange", value)
}

// ValidationErrors sets the validation error messages.
func (tc inputTime) ValidationErrors(value string) inputTime {
	return tc.set("validationErrors", value)
}

// Validations sets the validation rules.
func (tc inputTime) Validations(value string) inputTime {
	return tc.set("validations", value)
}

// Value sets the default value, must be static.
func (tc inputTime) Value(value string) inputTime {
	return tc.set("value", value)
}

// ValueFormat sets the value format, replacing format.
func (tc inputTime) ValueFormat(value string) inputTime {
	return tc.set("valueFormat", value)
}

// Visible sets whether the component is visible.
func (tc inputTime) Visible(value bool) inputTime {
	return tc.set("visible", value)
}

// VisibleOn sets an expression to control visibility.
func (tc inputTime) VisibleOn(value string) inputTime {
	return tc.set("visibleOn", value)
}

// Width sets the width in a table.
func (tc inputTime) Width(value string) inputTime {
	return tc.set("width", value)
}

func (tc inputTime) set(key string, value any) inputTime {
	tc[key] = value
	return tc
}
