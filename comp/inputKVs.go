package comp

// inputKVS represents a key-value object

type inputKVS Schema

// InputKVS creates a new InputKVS instance with default type
func InputKVS() inputKVS {
	return make(inputKVS).set("type", "input-kvs")
}

// set sets a field value and returns the instance
func (i inputKVS) set(key string, value any) inputKVS {
	i[key] = value
	return i
}

// AutoFill sets autoFill configuration
func (i inputKVS) AutoFill(value string) inputKVS {
	return i.set("autoFill", value)
}

// ClassName sets the outermost class name of the form
func (i inputKVS) ClassName(value string) inputKVS {
	return i.set("className", value)
}

// Description sets the form item description
func (i inputKVS) Description(value string) inputKVS {
	return i.set("description", value)
}

// Disabled sets whether the form item is disabled
func (i inputKVS) Disabled(value bool) inputKVS {
	return i.set("disabled", value)
}

// DisabledOn sets the condition for disabling the form item
func (i inputKVS) DisabledOn(value string) inputKVS {
	return i.set("disabledOn", value)
}

// Inline sets whether the form item is inline
func (i inputKVS) Inline(value bool) inputKVS {
	return i.set("inline", value)
}

// InputClassName sets the class name of the form controller
func (i inputKVS) InputClassName(value string) inputKVS {
	return i.set("inputClassName", value)
}

// Label sets the form item label
func (i inputKVS) Label(value string) inputKVS {
	return i.set("label", value)
}

// KeyItem sets the key item
func (i inputKVS) KeyItem(value any) inputKVS {
	return i.set("keyItem", value)
}

// ValueItems sets the value items
func (i inputKVS) ValueItems(value ...any) inputKVS {
	return i.set("valueItems", value)
}

// LabelAlign sets the label alignment, effective only when mode is horizontal
func (i inputKVS) LabelAlign(value string) inputKVS {
	return i.set("labelAlign", value)
}

// LabelClassName sets the class name of the label
func (i inputKVS) LabelClassName(value string) inputKVS {
	return i.set("labelClassName", value)
}

// LabelRemark sets the label remark
func (i inputKVS) LabelRemark(value string) inputKVS {
	return i.set("labelRemark", value)
}

// Name sets the field name, specifying the key when the form is submitted
func (i inputKVS) Name(value string) inputKVS {
	return i.set("name", value)
}

// Placeholder sets the form item placeholder
func (i inputKVS) Placeholder(value string) inputKVS {
	return i.set("placeholder", value)
}

// Required sets whether the form item is required
func (i inputKVS) Required(value bool) inputKVS {
	return i.set("required", value)
}

// RequiredOn sets the condition for the form item to be required
func (i inputKVS) RequiredOn(value string) inputKVS {
	return i.set("requiredOn", value)
}

// SubmitOnChange sets whether to submit the form when the form item value changes
func (i inputKVS) SubmitOnChange(value bool) inputKVS {
	return i.set("submitOnChange", value)
}

// ValidateApi sets the form validation API
func (i inputKVS) ValidateApi(value string) inputKVS {
	return i.set("validateApi", value)
}

// Validations sets the form item value format validations
func (i inputKVS) Validations(value string) inputKVS {
	return i.set("validations", value)
}

// Value sets the default value of the form item
func (i inputKVS) Value(value string) inputKVS {
	return i.set("value", value)
}

// Visible sets whether the form item is visible
func (i inputKVS) Visible(value bool) inputKVS {
	return i.set("visible", value)
}

// VisibleOn sets the condition for the form item to be visible
func (i inputKVS) VisibleOn(value string) inputKVS {
	return i.set("visibleOn", value)
}
