package comp

import "github.com/zrcoder/amisgo/model"

// inputDatetimeRange represents a date-time range input component.
type inputDatetimeRange model.Schema

// InputDatetimeRange creates a new instance with default type.
func InputDatetimeRange() inputDatetimeRange {
	return make(inputDatetimeRange).set("type", "input-datetime-range")
}

// AutoFill sets auto-fill configuration.
func (i inputDatetimeRange) AutoFill(value string) inputDatetimeRange {
	return i.set("autoFill", value)
}

// ClassName sets the outermost class name.
func (i inputDatetimeRange) ClassName(value string) inputDatetimeRange {
	return i.set("className", value)
}

// Clearable sets whether the input is clearable.
func (i inputDatetimeRange) Clearable(value bool) inputDatetimeRange {
	return i.set("clearable", value)
}

// Description sets the description of the form item.
func (i inputDatetimeRange) Description(value string) inputDatetimeRange {
	return i.set("description", value)
}

// Disabled sets whether the input is disabled.
func (i inputDatetimeRange) Disabled(value bool) inputDatetimeRange {
	return i.set("disabled", value)
}

// DisabledOn sets the condition to disable the input.
func (i inputDatetimeRange) DisabledOn(value string) inputDatetimeRange {
	return i.set("disabledOn", value)
}

// Format sets the value format of the date-time picker.
func (i inputDatetimeRange) Format(value string) inputDatetimeRange {
	return i.set("format", value)
}

// Inline sets whether the input is inline.
func (i inputDatetimeRange) Inline(value bool) inputDatetimeRange {
	return i.set("inline", value)
}

// InputClassName sets the class name of the input control.
func (i inputDatetimeRange) InputClassName(value string) inputDatetimeRange {
	return i.set("inputClassName", value)
}

// InputFormat sets the display format of the date-time picker.
func (i inputDatetimeRange) InputFormat(value string) inputDatetimeRange {
	return i.set("inputFormat", value)
}

// Label sets the label of the form item.
func (i inputDatetimeRange) Label(value string) inputDatetimeRange {
	return i.set("label", value)
}

// LabelAlign sets the alignment of the label.
func (i inputDatetimeRange) LabelAlign(value string) inputDatetimeRange {
	return i.set("labelAlign", value)
}

// LabelClassName sets the class name of the label.
func (i inputDatetimeRange) LabelClassName(value string) inputDatetimeRange {
	return i.set("labelClassName", value)
}

// LabelRemark sets the remark for the label.
func (i inputDatetimeRange) LabelRemark(value string) inputDatetimeRange {
	return i.set("labelRemark", value)
}

// MaxDate sets the maximum date-time limit.
func (i inputDatetimeRange) MaxDate(value string) inputDatetimeRange {
	return i.set("maxDate", value)
}

// MinDate sets the minimum date-time limit.
func (i inputDatetimeRange) MinDate(value string) inputDatetimeRange {
	return i.set("minDate", value)
}

// Name sets the field name for form submission.
func (i inputDatetimeRange) Name(value string) inputDatetimeRange {
	return i.set("name", value)
}

// Placeholder sets the placeholder text.
func (i inputDatetimeRange) Placeholder(value string) inputDatetimeRange {
	return i.set("placeholder", value)
}

// Ranges sets the date range shortcuts.
func (i inputDatetimeRange) Ranges(value string) inputDatetimeRange {
	return i.set("ranges", value)
}

// Required sets whether the input is required.
func (i inputDatetimeRange) Required(value bool) inputDatetimeRange {
	return i.set("required", value)
}

// RequiredOn sets the condition to make the input required.
func (i inputDatetimeRange) RequiredOn(value string) inputDatetimeRange {
	return i.set("requiredOn", value)
}

// SubmitOnChange sets whether to submit the form on value change.
func (i inputDatetimeRange) SubmitOnChange(value bool) inputDatetimeRange {
	return i.set("submitOnChange", value)
}

// Utc sets whether to save the value in UTC.
func (i inputDatetimeRange) Utc(value bool) inputDatetimeRange {
	return i.set("utc", value)
}

// ValidateApi sets the validation API.
func (i inputDatetimeRange) ValidateApi(value string) inputDatetimeRange {
	return i.set("validateApi", value)
}

// Validations sets the validation rules.
func (i inputDatetimeRange) Validations(value string) inputDatetimeRange {
	return i.set("validations", value)
}

// Value sets the default value.
func (i inputDatetimeRange) Value(value string) inputDatetimeRange {
	return i.set("value", value)
}

// Visible sets whether the input is visible.
func (i inputDatetimeRange) Visible(value bool) inputDatetimeRange {
	return i.set("visible", value)
}

// VisibleOn sets the condition to make the input visible.
func (i inputDatetimeRange) VisibleOn(value string) inputDatetimeRange {
	return i.set("visibleOn", value)
}

func (i inputDatetimeRange) set(key string, value any) inputDatetimeRange {
	i[key] = value
	return i
}
