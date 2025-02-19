package comp

import "github.com/zrcoder/amisgo/schema"

// InputDatetimeRange represents a date-time range input component.
type InputDatetimeRange schema.Schema

func NewInputDatetimeRange() InputDatetimeRange {
	return InputDatetimeRange{"type": "input-datetime-range"}
}

// AutoFill sets auto-fill configuration.
func (i InputDatetimeRange) AutoFill(value string) InputDatetimeRange {
	return i.set("autoFill", value)
}

// ClassName sets the outermost class name.
func (i InputDatetimeRange) ClassName(value string) InputDatetimeRange {
	return i.set("className", value)
}

// Clearable sets whether the input is clearable.
func (i InputDatetimeRange) Clearable(value bool) InputDatetimeRange {
	return i.set("clearable", value)
}

// Description sets the description of the form item.
func (i InputDatetimeRange) Description(value string) InputDatetimeRange {
	return i.set("description", value)
}

// Disabled sets whether the input is disabled.
func (i InputDatetimeRange) Disabled(value bool) InputDatetimeRange {
	return i.set("disabled", value)
}

// DisabledOn sets the condition to disable the input.
func (i InputDatetimeRange) DisabledOn(value string) InputDatetimeRange {
	return i.set("disabledOn", value)
}

// Format sets the value format of the date-time picker.
func (i InputDatetimeRange) Format(value string) InputDatetimeRange {
	return i.set("format", value)
}

// Inline sets whether the input is inline.
func (i InputDatetimeRange) Inline(value bool) InputDatetimeRange {
	return i.set("inline", value)
}

// InputClassName sets the class name of the input control.
func (i InputDatetimeRange) InputClassName(value string) InputDatetimeRange {
	return i.set("inputClassName", value)
}

// InputFormat sets the display format of the date-time picker.
func (i InputDatetimeRange) InputFormat(value string) InputDatetimeRange {
	return i.set("inputFormat", value)
}

// Label sets the label of the form item.
func (i InputDatetimeRange) Label(value string) InputDatetimeRange {
	return i.set("label", value)
}

// LabelAlign sets the alignment of the label.
func (i InputDatetimeRange) LabelAlign(value string) InputDatetimeRange {
	return i.set("labelAlign", value)
}

// LabelClassName sets the class name of the label.
func (i InputDatetimeRange) LabelClassName(value string) InputDatetimeRange {
	return i.set("labelClassName", value)
}

// LabelRemark sets the remark for the label.
func (i InputDatetimeRange) LabelRemark(value string) InputDatetimeRange {
	return i.set("labelRemark", value)
}

// MaxDate sets the maximum date-time limit.
func (i InputDatetimeRange) MaxDate(value string) InputDatetimeRange {
	return i.set("maxDate", value)
}

// MinDate sets the minimum date-time limit.
func (i InputDatetimeRange) MinDate(value string) InputDatetimeRange {
	return i.set("minDate", value)
}

// Name sets the field name for form submission.
func (i InputDatetimeRange) Name(value string) InputDatetimeRange {
	return i.set("name", value)
}

// Placeholder sets the placeholder text.
func (i InputDatetimeRange) Placeholder(value string) InputDatetimeRange {
	return i.set("placeholder", value)
}

// Ranges sets the date range shortcuts.
func (i InputDatetimeRange) Ranges(value string) InputDatetimeRange {
	return i.set("ranges", value)
}

// Required sets whether the input is required.
func (i InputDatetimeRange) Required(value bool) InputDatetimeRange {
	return i.set("required", value)
}

// RequiredOn sets the condition to make the input required.
func (i InputDatetimeRange) RequiredOn(value string) InputDatetimeRange {
	return i.set("requiredOn", value)
}

// SubmitOnChange sets whether to submit the form on value change.
func (i InputDatetimeRange) SubmitOnChange(value bool) InputDatetimeRange {
	return i.set("submitOnChange", value)
}

// Utc sets whether to save the value in UTC.
func (i InputDatetimeRange) Utc(value bool) InputDatetimeRange {
	return i.set("utc", value)
}

// ValidateApi sets the validation API.
func (i InputDatetimeRange) ValidateApi(value string) InputDatetimeRange {
	return i.set("validateApi", value)
}

// Validations sets the validation rules.
func (i InputDatetimeRange) Validations(value string) InputDatetimeRange {
	return i.set("validations", value)
}

// Value sets the default value.
func (i InputDatetimeRange) Value(value string) InputDatetimeRange {
	return i.set("value", value)
}

// Visible sets whether the input is visible.
func (i InputDatetimeRange) Visible(value bool) InputDatetimeRange {
	return i.set("visible", value)
}

// VisibleOn sets the condition to make the input visible.
func (i InputDatetimeRange) VisibleOn(value string) InputDatetimeRange {
	return i.set("visibleOn", value)
}

func (i InputDatetimeRange) set(key string, value any) InputDatetimeRange {
	i[key] = value
	return i
}
