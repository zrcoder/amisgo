package comp

import "github.com/zrcoder/amisgo/model"

// InputKVS represents a key-value object
type InputKVS model.Schema

func NewInputKVS() InputKVS {
	return InputKVS{"type": "input-kvs"}
}

// set sets a field value and returns the instance
func (i InputKVS) set(key string, value any) InputKVS {
	i[key] = value
	return i
}

// AutoFill sets autoFill configuration
func (i InputKVS) AutoFill(value string) InputKVS {
	return i.set("autoFill", value)
}

// ClassName sets the outermost class name of the form
func (i InputKVS) ClassName(value string) InputKVS {
	return i.set("className", value)
}

// Description sets the form item description
func (i InputKVS) Description(value string) InputKVS {
	return i.set("description", value)
}

// Disabled sets whether the form item is disabled
func (i InputKVS) Disabled(value bool) InputKVS {
	return i.set("disabled", value)
}

// DisabledOn sets the condition for disabling the form item
func (i InputKVS) DisabledOn(value string) InputKVS {
	return i.set("disabledOn", value)
}

// Inline sets whether the form item is inline
func (i InputKVS) Inline(value bool) InputKVS {
	return i.set("inline", value)
}

// InputClassName sets the class name of the form controller
func (i InputKVS) InputClassName(value string) InputKVS {
	return i.set("inputClassName", value)
}

// Label sets the form item label
func (i InputKVS) Label(value string) InputKVS {
	return i.set("label", value)
}

// KeyItem sets the key item
func (i InputKVS) KeyItem(value any) InputKVS {
	return i.set("keyItem", value)
}

// ValueItems sets the value items
func (i InputKVS) ValueItems(value ...any) InputKVS {
	return i.set("valueItems", value)
}

// LabelAlign sets the label alignment, effective only when mode is horizontal
func (i InputKVS) LabelAlign(value string) InputKVS {
	return i.set("labelAlign", value)
}

// LabelClassName sets the class name of the label
func (i InputKVS) LabelClassName(value string) InputKVS {
	return i.set("labelClassName", value)
}

// LabelRemark sets the label remark
func (i InputKVS) LabelRemark(value string) InputKVS {
	return i.set("labelRemark", value)
}

// Name sets the field name, specifying the key when the form is submitted
func (i InputKVS) Name(value string) InputKVS {
	return i.set("name", value)
}

// Placeholder sets the form item placeholder
func (i InputKVS) Placeholder(value string) InputKVS {
	return i.set("placeholder", value)
}

// Required sets whether the form item is required
func (i InputKVS) Required(value bool) InputKVS {
	return i.set("required", value)
}

// RequiredOn sets the condition for the form item to be required
func (i InputKVS) RequiredOn(value string) InputKVS {
	return i.set("requiredOn", value)
}

// SubmitOnChange sets whether to submit the form when the form item value changes
func (i InputKVS) SubmitOnChange(value bool) InputKVS {
	return i.set("submitOnChange", value)
}

// ValidateApi sets the form validation API
func (i InputKVS) ValidateApi(value string) InputKVS {
	return i.set("validateApi", value)
}

// Validations sets the form item value format validations
func (i InputKVS) Validations(value string) InputKVS {
	return i.set("validations", value)
}

// Value sets the default value of the form item
func (i InputKVS) Value(value string) InputKVS {
	return i.set("value", value)
}

// Visible sets whether the form item is visible
func (i InputKVS) Visible(value bool) InputKVS {
	return i.set("visible", value)
}

// VisibleOn sets the condition for the form item to be visible
func (i InputKVS) VisibleOn(value string) InputKVS {
	return i.set("visibleOn", value)
}
