package comp

type inputKV Schema

// InputKV creates a new inputKV instance with default type
func InputKV() inputKV {
	return make(inputKV).set("type", "input-kv")
}

// set sets a field value and returns the inputKV instance
func (i inputKV) set(key string, value any) inputKV {
	i[key] = value
	return i
}

// AutoFill sets the autoFill value
func (i inputKV) AutoFill(value string) inputKV {
	return i.set("autoFill", value)
}

// ClassName sets the outer class name
func (i inputKV) ClassName(value string) inputKV {
	return i.set("className", value)
}

// DefaultValue sets the default value
func (i inputKV) DefaultValue(value string) inputKV {
	return i.set("defaultValue", value)
}

// Description sets the description
func (i inputKV) Description(value string) inputKV {
	return i.set("description", value)
}

// Disabled sets whether the input is disabled
func (i inputKV) Disabled(value bool) inputKV {
	return i.set("disabled", value)
}

// DisabledOn sets the condition for disabling the input
func (i inputKV) DisabledOn(value string) inputKV {
	return i.set("disabledOn", value)
}

// Draggable sets whether the input is draggable
func (i inputKV) Draggable(value bool) inputKV {
	return i.set("draggable", value)
}

// Inline sets whether the input is inline
func (i inputKV) Inline(value bool) inputKV {
	return i.set("inline", value)
}

// InputClassName sets the input controller class name
func (i inputKV) InputClassName(value string) inputKV {
	return i.set("inputClassName", value)
}

// KeyPlaceholder sets the key placeholder
func (i inputKV) KeyPlaceholder(value string) inputKV {
	return i.set("keyPlaceholder", value)
}

// Label sets the label
func (i inputKV) Label(value string) inputKV {
	return i.set("label", value)
}

// LabelAlign sets the label alignment
func (i inputKV) LabelAlign(value string) inputKV {
	return i.set("labelAlign", value)
}

// LabelClassName sets the label class name
func (i inputKV) LabelClassName(value string) inputKV {
	return i.set("labelClassName", value)
}

// LabelRemark sets the label remark
func (i inputKV) LabelRemark(value string) inputKV {
	return i.set("labelRemark", value)
}

// Name sets the field name
func (i inputKV) Name(value string) inputKV {
	return i.set("name", value)
}

// Placeholder sets the placeholder
func (i inputKV) Placeholder(value string) inputKV {
	return i.set("placeholder", value)
}

// Required sets whether the input is required
func (i inputKV) Required(value bool) inputKV {
	return i.set("required", value)
}

// RequiredOn sets the condition for requiring the input
func (i inputKV) RequiredOn(value string) inputKV {
	return i.set("requiredOn", value)
}

// SubmitOnChange sets whether to submit the form on change
func (i inputKV) SubmitOnChange(value bool) inputKV {
	return i.set("submitOnChange", value)
}

// ValidateApi sets the validation API
func (i inputKV) ValidateApi(value string) inputKV {
	return i.set("validateApi", value)
}

// Validations sets the validation rules
func (i inputKV) Validations(value string) inputKV {
	return i.set("validations", value)
}

// Value sets the default value
func (i inputKV) Value(value any) inputKV {
	return i.set("value", value)
}

// ValuePlaceholder sets the value placeholder
func (i inputKV) ValuePlaceholder(value string) inputKV {
	return i.set("valuePlaceholder", value)
}

// ValueType sets the value type
func (i inputKV) ValueType(value string) inputKV {
	return i.set("valueType", value)
}

// Visible sets whether the input is visible
func (i inputKV) Visible(value bool) inputKV {
	return i.set("visible", value)
}

// VisibleOn sets the condition for visibility
func (i inputKV) VisibleOn(value string) inputKV {
	return i.set("visibleOn", value)
}
