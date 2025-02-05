package comp

import "github.com/zrcoder/amisgo/model"

// InputTimeRange represents a time range input component.
type InputTimeRange model.Schema

func NewInputTimeRange() InputTimeRange {
	return InputTimeRange{"type": "input-time-range"}
}

// set sets a field value and returns the instance.
func (i InputTimeRange) set(key string, value any) InputTimeRange {
	i[key] = value
	return i
}

// AutoFill sets the autoFill value.
func (i InputTimeRange) AutoFill(value string) InputTimeRange {
	return i.set("autoFill", value)
}

// ClassName sets the outermost class name.
func (i InputTimeRange) ClassName(value string) InputTimeRange {
	return i.set("className", value)
}

// Clearable sets whether the input is clearable.
func (i InputTimeRange) Clearable(value bool) InputTimeRange {
	return i.set("clearable", value)
}

// Description sets the description of the form item.
func (i InputTimeRange) Description(value string) InputTimeRange {
	return i.set("description", value)
}

// Disabled sets whether the input is disabled.
func (i InputTimeRange) Disabled(value bool) InputTimeRange {
	return i.set("disabled", value)
}

// DisabledOn sets the condition for disabling the input.
func (i InputTimeRange) DisabledOn(value string) InputTimeRange {
	return i.set("disabledOn", value)
}

// Embed sets whether the input is in inline mode.
func (i InputTimeRange) Embed(value bool) InputTimeRange {
	return i.set("embed", value)
}

// Format sets the value format of the time range picker.
func (i InputTimeRange) Format(value string) InputTimeRange {
	return i.set("format", value)
}

// Inline sets whether the input is inline.
func (i InputTimeRange) Inline(value bool) InputTimeRange {
	return i.set("inline", value)
}

// InputClassName sets the class name of the form controller.
func (i InputTimeRange) InputClassName(value string) InputTimeRange {
	return i.set("inputClassName", value)
}

// InputFormat sets the display format of the time range picker.
func (i InputTimeRange) InputFormat(value string) InputTimeRange {
	return i.set("inputFormat", value)
}

// Label sets the label of the form item.
func (i InputTimeRange) Label(value string) InputTimeRange {
	return i.set("label", value)
}

// LabelAlign sets the alignment of the label.
func (i InputTimeRange) LabelAlign(value string) InputTimeRange {
	return i.set("labelAlign", value)
}

// LabelClassName sets the class name of the label.
func (i InputTimeRange) LabelClassName(value string) InputTimeRange {
	return i.set("labelClassName", value)
}

// LabelRemark sets the remark of the label.
func (i InputTimeRange) LabelRemark(value string) InputTimeRange {
	return i.set("labelRemark", value)
}

// Name sets the name of the field.
func (i InputTimeRange) Name(value string) InputTimeRange {
	return i.set("name", value)
}

// Placeholder sets the placeholder text.
func (i InputTimeRange) Placeholder(value string) InputTimeRange {
	return i.set("placeholder", value)
}

// Required sets whether the input is required.
func (i InputTimeRange) Required(value bool) InputTimeRange {
	return i.set("required", value)
}

// RequiredOn sets the condition for requiring the input.
func (i InputTimeRange) RequiredOn(value string) InputTimeRange {
	return i.set("requiredOn", value)
}

// SubmitOnChange sets whether to submit the form on value change.
func (i InputTimeRange) SubmitOnChange(value bool) InputTimeRange {
	return i.set("submitOnChange", value)
}

// TimeFormat sets the time format of the time range picker.
func (i InputTimeRange) TimeFormat(value string) InputTimeRange {
	return i.set("timeFormat", value)
}

// ValidateApi sets the validation API.
func (i InputTimeRange) ValidateApi(value string) InputTimeRange {
	return i.set("validateApi", value)
}

// Validations sets the validation rules.
func (i InputTimeRange) Validations(value string) InputTimeRange {
	return i.set("validations", value)
}

// Value sets the default value of the form item.
func (i InputTimeRange) Value(value string) InputTimeRange {
	return i.set("value", value)
}

// Visible sets whether the input is visible.
func (i InputTimeRange) Visible(value bool) InputTimeRange {
	return i.set("visible", value)
}

// VisibleOn sets the condition for visibility.
func (i InputTimeRange) VisibleOn(value string) InputTimeRange {
	return i.set("visibleOn", value)
}
