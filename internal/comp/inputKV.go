package comp

import "github.com/zrcoder/amisgo/model"

type InputKV model.Schema

func NewInputKV() InputKV {
	return InputKV{"type": "input-kv"}
}

// set sets a field value and returns the inputKV instance
func (i InputKV) set(key string, value any) InputKV {
	i[key] = value
	return i
}

// AutoFill sets the autoFill value
func (i InputKV) AutoFill(value string) InputKV {
	return i.set("autoFill", value)
}

// ClassName sets the outer class name
func (i InputKV) ClassName(value string) InputKV {
	return i.set("className", value)
}

// DefaultValue sets the default value
func (i InputKV) DefaultValue(value string) InputKV {
	return i.set("defaultValue", value)
}

// Description sets the description
func (i InputKV) Description(value string) InputKV {
	return i.set("description", value)
}

// Disabled sets whether the input is disabled
func (i InputKV) Disabled(value bool) InputKV {
	return i.set("disabled", value)
}

// DisabledOn sets the condition for disabling the input
func (i InputKV) DisabledOn(value string) InputKV {
	return i.set("disabledOn", value)
}

// Draggable sets whether the input is draggable
func (i InputKV) Draggable(value bool) InputKV {
	return i.set("draggable", value)
}

// Inline sets whether the input is inline
func (i InputKV) Inline(value bool) InputKV {
	return i.set("inline", value)
}

// InputClassName sets the input controller class name
func (i InputKV) InputClassName(value string) InputKV {
	return i.set("inputClassName", value)
}

// KeyPlaceholder sets the key placeholder
func (i InputKV) KeyPlaceholder(value string) InputKV {
	return i.set("keyPlaceholder", value)
}

// Label sets the label
func (i InputKV) Label(value string) InputKV {
	return i.set("label", value)
}

// LabelAlign sets the label alignment
func (i InputKV) LabelAlign(value string) InputKV {
	return i.set("labelAlign", value)
}

// LabelClassName sets the label class name
func (i InputKV) LabelClassName(value string) InputKV {
	return i.set("labelClassName", value)
}

// LabelRemark sets the label remark
func (i InputKV) LabelRemark(value string) InputKV {
	return i.set("labelRemark", value)
}

// Name sets the field name
func (i InputKV) Name(value string) InputKV {
	return i.set("name", value)
}

// Placeholder sets the placeholder
func (i InputKV) Placeholder(value string) InputKV {
	return i.set("placeholder", value)
}

// Required sets whether the input is required
func (i InputKV) Required(value bool) InputKV {
	return i.set("required", value)
}

// RequiredOn sets the condition for requiring the input
func (i InputKV) RequiredOn(value string) InputKV {
	return i.set("requiredOn", value)
}

// SubmitOnChange sets whether to submit the form on change
func (i InputKV) SubmitOnChange(value bool) InputKV {
	return i.set("submitOnChange", value)
}

// ValidateApi sets the validation API
func (i InputKV) ValidateApi(value string) InputKV {
	return i.set("validateApi", value)
}

// Validations sets the validation rules
func (i InputKV) Validations(value string) InputKV {
	return i.set("validations", value)
}

// Value sets the default value
func (i InputKV) Value(value any) InputKV {
	return i.set("value", value)
}

// ValuePlaceholder sets the value placeholder
func (i InputKV) ValuePlaceholder(value string) InputKV {
	return i.set("valuePlaceholder", value)
}

// ValueType sets the value type
func (i InputKV) ValueType(value string) InputKV {
	return i.set("valueType", value)
}

// Visible sets whether the input is visible
func (i InputKV) Visible(value bool) InputKV {
	return i.set("visible", value)
}

// VisibleOn sets the condition for visibility
func (i InputKV) VisibleOn(value string) InputKV {
	return i.set("visibleOn", value)
}
