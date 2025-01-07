package comp

// inputTimeRange represents a time range input component.

type inputTimeRange Schema

// InputTimeRange creates a new InputTimeRange instance with default type.
func InputTimeRange() inputTimeRange {
	return make(inputTimeRange).set("type", "input-time-range")
}

// set sets a field value and returns the instance.
func (i inputTimeRange) set(key string, value any) inputTimeRange {
	i[key] = value
	return i
}

// AutoFill sets the autoFill value.
func (i inputTimeRange) AutoFill(value string) inputTimeRange {
	return i.set("autoFill", value)
}

// ClassName sets the outermost class name.
func (i inputTimeRange) ClassName(value string) inputTimeRange {
	return i.set("className", value)
}

// Clearable sets whether the input is clearable.
func (i inputTimeRange) Clearable(value bool) inputTimeRange {
	return i.set("clearable", value)
}

// Description sets the description of the form item.
func (i inputTimeRange) Description(value string) inputTimeRange {
	return i.set("description", value)
}

// Disabled sets whether the input is disabled.
func (i inputTimeRange) Disabled(value bool) inputTimeRange {
	return i.set("disabled", value)
}

// DisabledOn sets the condition for disabling the input.
func (i inputTimeRange) DisabledOn(value string) inputTimeRange {
	return i.set("disabledOn", value)
}

// Embed sets whether the input is in inline mode.
func (i inputTimeRange) Embed(value bool) inputTimeRange {
	return i.set("embed", value)
}

// Format sets the value format of the time range picker.
func (i inputTimeRange) Format(value string) inputTimeRange {
	return i.set("format", value)
}

// Inline sets whether the input is inline.
func (i inputTimeRange) Inline(value bool) inputTimeRange {
	return i.set("inline", value)
}

// InputClassName sets the class name of the form controller.
func (i inputTimeRange) InputClassName(value string) inputTimeRange {
	return i.set("inputClassName", value)
}

// InputFormat sets the display format of the time range picker.
func (i inputTimeRange) InputFormat(value string) inputTimeRange {
	return i.set("inputFormat", value)
}

// Label sets the label of the form item.
func (i inputTimeRange) Label(value string) inputTimeRange {
	return i.set("label", value)
}

// LabelAlign sets the alignment of the label.
func (i inputTimeRange) LabelAlign(value string) inputTimeRange {
	return i.set("labelAlign", value)
}

// LabelClassName sets the class name of the label.
func (i inputTimeRange) LabelClassName(value string) inputTimeRange {
	return i.set("labelClassName", value)
}

// LabelRemark sets the remark of the label.
func (i inputTimeRange) LabelRemark(value string) inputTimeRange {
	return i.set("labelRemark", value)
}

// Name sets the name of the field.
func (i inputTimeRange) Name(value string) inputTimeRange {
	return i.set("name", value)
}

// Placeholder sets the placeholder text.
func (i inputTimeRange) Placeholder(value string) inputTimeRange {
	return i.set("placeholder", value)
}

// Required sets whether the input is required.
func (i inputTimeRange) Required(value bool) inputTimeRange {
	return i.set("required", value)
}

// RequiredOn sets the condition for requiring the input.
func (i inputTimeRange) RequiredOn(value string) inputTimeRange {
	return i.set("requiredOn", value)
}

// SubmitOnChange sets whether to submit the form on value change.
func (i inputTimeRange) SubmitOnChange(value bool) inputTimeRange {
	return i.set("submitOnChange", value)
}

// TimeFormat sets the time format of the time range picker.
func (i inputTimeRange) TimeFormat(value string) inputTimeRange {
	return i.set("timeFormat", value)
}

// ValidateApi sets the validation API.
func (i inputTimeRange) ValidateApi(value string) inputTimeRange {
	return i.set("validateApi", value)
}

// Validations sets the validation rules.
func (i inputTimeRange) Validations(value string) inputTimeRange {
	return i.set("validations", value)
}

// Value sets the default value of the form item.
func (i inputTimeRange) Value(value string) inputTimeRange {
	return i.set("value", value)
}

// Visible sets whether the input is visible.
func (i inputTimeRange) Visible(value bool) inputTimeRange {
	return i.set("visible", value)
}

// VisibleOn sets the condition for visibility.
func (i inputTimeRange) VisibleOn(value string) inputTimeRange {
	return i.set("visibleOn", value)
}
