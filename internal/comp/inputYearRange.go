package comp

import "github.com/zrcoder/amisgo/schema"

// InputYearRange represents a year range input component.
type InputYearRange schema.Schema

func NewInputYearRange() InputYearRange {
	return InputYearRange{"type": "input-year-range"}
}

// set is an internal method to set a field value and return the instance.
func (i InputYearRange) set(key string, value any) InputYearRange {
	i[key] = value
	return i
}

// AutoFill sets the autoFill value.
func (i InputYearRange) AutoFill(value string) InputYearRange {
	return i.set("autoFill", value)
}

// ClassName sets the outermost class name of the form.
func (i InputYearRange) ClassName(value string) InputYearRange {
	return i.set("className", value)
}

// Clearable sets whether the input is clearable.
func (i InputYearRange) Clearable(value bool) InputYearRange {
	return i.set("clearable", value)
}

// Description sets the description of the form item.
func (i InputYearRange) Description(value string) InputYearRange {
	return i.set("description", value)
}

// Disabled sets whether the input is disabled.
func (i InputYearRange) Disabled(value bool) InputYearRange {
	return i.set("disabled", value)
}

// DisabledOn sets the condition to disable the input.
func (i InputYearRange) DisabledOn(value string) InputYearRange {
	return i.set("disabledOn", value)
}

// Embed sets whether the input is in inline mode.
func (i InputYearRange) Embed(value bool) InputYearRange {
	return i.set("embed", value)
}

// Format sets the format of the year range value.
func (i InputYearRange) Format(value string) InputYearRange {
	return i.set("format", value)
}

// Inline sets whether the input is inline.
func (i InputYearRange) Inline(value bool) InputYearRange {
	return i.set("inline", value)
}

// InputClassName sets the class name of the form controller.
func (i InputYearRange) InputClassName(value string) InputYearRange {
	return i.set("inputClassName", value)
}

// InputFormat sets the display format of the year range.
func (i InputYearRange) InputFormat(value string) InputYearRange {
	return i.set("inputFormat", value)
}

// Label sets the label of the form item.
func (i InputYearRange) Label(value string) InputYearRange {
	return i.set("label", value)
}

// LabelAlign sets the alignment of the label, effective only in horizontal mode.
func (i InputYearRange) LabelAlign(value string) InputYearRange {
	return i.set("labelAlign", value)
}

// LabelClassName sets the class name of the label.
func (i InputYearRange) LabelClassName(value string) InputYearRange {
	return i.set("labelClassName", value)
}

// LabelRemark sets the remark of the label.
func (i InputYearRange) LabelRemark(value string) InputYearRange {
	return i.set("labelRemark", value)
}

// MaxDate sets the maximum date limit.
func (i InputYearRange) MaxDate(value string) InputYearRange {
	return i.set("maxDate", value)
}

// MaxDuration sets the maximum duration limit, e.g., 4 years.
func (i InputYearRange) MaxDuration(value string) InputYearRange {
	return i.set("maxDuration", value)
}

// MinDate sets the minimum date limit.
func (i InputYearRange) MinDate(value string) InputYearRange {
	return i.set("minDate", value)
}

// MinDuration sets the minimum duration limit, e.g., 2 years.
func (i InputYearRange) MinDuration(value string) InputYearRange {
	return i.set("minDuration", value)
}

// Name sets the field name for form submission.
func (i InputYearRange) Name(value string) InputYearRange {
	return i.set("name", value)
}

// Placeholder sets the placeholder text.
func (i InputYearRange) Placeholder(value string) InputYearRange {
	return i.set("placeholder", value)
}

// Required sets whether the input is required.
func (i InputYearRange) Required(value bool) InputYearRange {
	return i.set("required", value)
}

// RequiredOn sets the condition to make the input required.
func (i InputYearRange) RequiredOn(value string) InputYearRange {
	return i.set("requiredOn", value)
}

// SubmitOnChange sets whether to submit the form when the input value changes.
func (i InputYearRange) SubmitOnChange(value bool) InputYearRange {
	return i.set("submitOnChange", value)
}

// Utc sets whether to save the value in UTC.
func (i InputYearRange) Utc(value bool) InputYearRange {
	return i.set("utc", value)
}

// ValidateApi sets the validation API.
func (i InputYearRange) ValidateApi(value string) InputYearRange {
	return i.set("validateApi", value)
}

// Validations sets the validation rules, multiple rules separated by commas.
func (i InputYearRange) Validations(value string) InputYearRange {
	return i.set("validations", value)
}

// Value sets the default value of the form item.
func (i InputYearRange) Value(value string) InputYearRange {
	return i.set("value", value)
}

// Visible sets whether the input is visible.
func (i InputYearRange) Visible(value bool) InputYearRange {
	return i.set("visible", value)
}

// VisibleOn sets the condition to make the input visible.
func (i InputYearRange) VisibleOn(value string) InputYearRange {
	return i.set("visibleOn", value)
}

// UTC sets if use utc time
func (y InputYearRange) UTC(value bool) InputYearRange {
	return y.set("utc", value)
}

// Animation enables or disables cursor animation, default is enabled
func (y InputYearRange) Animation(value bool) InputYearRange {
	return y.set("animation", value)
}
