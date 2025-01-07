package comp

// inputYearRange represents a year range input component.

type inputYearRange Schema

// InputYearRange creates a new instance of inputYearRange with default type.
func InputYearRange() inputYearRange {
	return make(inputYearRange).set("type", "input-year-range")
}

// set is an internal method to set a field value and return the instance.
func (i inputYearRange) set(key string, value any) inputYearRange {
	i[key] = value
	return i
}

// AutoFill sets the autoFill value.
func (i inputYearRange) AutoFill(value string) inputYearRange {
	return i.set("autoFill", value)
}

// ClassName sets the outermost class name of the form.
func (i inputYearRange) ClassName(value string) inputYearRange {
	return i.set("className", value)
}

// Clearable sets whether the input is clearable.
func (i inputYearRange) Clearable(value bool) inputYearRange {
	return i.set("clearable", value)
}

// Description sets the description of the form item.
func (i inputYearRange) Description(value string) inputYearRange {
	return i.set("description", value)
}

// Disabled sets whether the input is disabled.
func (i inputYearRange) Disabled(value bool) inputYearRange {
	return i.set("disabled", value)
}

// DisabledOn sets the condition to disable the input.
func (i inputYearRange) DisabledOn(value string) inputYearRange {
	return i.set("disabledOn", value)
}

// Embed sets whether the input is in inline mode.
func (i inputYearRange) Embed(value bool) inputYearRange {
	return i.set("embed", value)
}

// Format sets the format of the year range value.
func (i inputYearRange) Format(value string) inputYearRange {
	return i.set("format", value)
}

// Inline sets whether the input is inline.
func (i inputYearRange) Inline(value bool) inputYearRange {
	return i.set("inline", value)
}

// InputClassName sets the class name of the form controller.
func (i inputYearRange) InputClassName(value string) inputYearRange {
	return i.set("inputClassName", value)
}

// InputFormat sets the display format of the year range.
func (i inputYearRange) InputFormat(value string) inputYearRange {
	return i.set("inputFormat", value)
}

// Label sets the label of the form item.
func (i inputYearRange) Label(value string) inputYearRange {
	return i.set("label", value)
}

// LabelAlign sets the alignment of the label, effective only in horizontal mode.
func (i inputYearRange) LabelAlign(value string) inputYearRange {
	return i.set("labelAlign", value)
}

// LabelClassName sets the class name of the label.
func (i inputYearRange) LabelClassName(value string) inputYearRange {
	return i.set("labelClassName", value)
}

// LabelRemark sets the remark of the label.
func (i inputYearRange) LabelRemark(value string) inputYearRange {
	return i.set("labelRemark", value)
}

// MaxDate sets the maximum date limit.
func (i inputYearRange) MaxDate(value string) inputYearRange {
	return i.set("maxDate", value)
}

// MaxDuration sets the maximum duration limit, e.g., 4 years.
func (i inputYearRange) MaxDuration(value string) inputYearRange {
	return i.set("maxDuration", value)
}

// MinDate sets the minimum date limit.
func (i inputYearRange) MinDate(value string) inputYearRange {
	return i.set("minDate", value)
}

// MinDuration sets the minimum duration limit, e.g., 2 years.
func (i inputYearRange) MinDuration(value string) inputYearRange {
	return i.set("minDuration", value)
}

// Name sets the field name for form submission.
func (i inputYearRange) Name(value string) inputYearRange {
	return i.set("name", value)
}

// Placeholder sets the placeholder text.
func (i inputYearRange) Placeholder(value string) inputYearRange {
	return i.set("placeholder", value)
}

// Required sets whether the input is required.
func (i inputYearRange) Required(value bool) inputYearRange {
	return i.set("required", value)
}

// RequiredOn sets the condition to make the input required.
func (i inputYearRange) RequiredOn(value string) inputYearRange {
	return i.set("requiredOn", value)
}

// SubmitOnChange sets whether to submit the form when the input value changes.
func (i inputYearRange) SubmitOnChange(value bool) inputYearRange {
	return i.set("submitOnChange", value)
}

// Utc sets whether to save the value in UTC.
func (i inputYearRange) Utc(value bool) inputYearRange {
	return i.set("utc", value)
}

// ValidateApi sets the validation API.
func (i inputYearRange) ValidateApi(value string) inputYearRange {
	return i.set("validateApi", value)
}

// Validations sets the validation rules, multiple rules separated by commas.
func (i inputYearRange) Validations(value string) inputYearRange {
	return i.set("validations", value)
}

// Value sets the default value of the form item.
func (i inputYearRange) Value(value string) inputYearRange {
	return i.set("value", value)
}

// Visible sets whether the input is visible.
func (i inputYearRange) Visible(value bool) inputYearRange {
	return i.set("visible", value)
}

// VisibleOn sets the condition to make the input visible.
func (i inputYearRange) VisibleOn(value string) inputYearRange {
	return i.set("visibleOn", value)
}
