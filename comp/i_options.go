package comp

type MOptions Schema

// Options creates a new Options instance
func Options() MOptions {
	return MOptions{}
}

// set sets a field value
func (o MOptions) set(key string, value any) MOptions {
	o[key] = value
	return o
}

// AutoFill sets autoFill value
func (o MOptions) AutoFill(value string) MOptions {
	return o.set("autoFill", value)
}

// ClassName sets the outer class name
func (o MOptions) ClassName(value string) MOptions {
	return o.set("className", value)
}

// Description sets the description
func (o MOptions) Description(value string) MOptions {
	return o.set("description", value)
}

// Disabled sets the disabled state
func (o MOptions) Disabled(value bool) MOptions {
	return o.set("disabled", value)
}

// DisabledOn sets the condition for disabling
func (o MOptions) DisabledOn(value string) MOptions {
	return o.set("disabledOn", value)
}

// ExtractValue sets extractValue, effective when joinValues is true
func (o MOptions) ExtractValue(value bool) MOptions {
	return o.set("extractValue", value)
}

// Inline sets the inline state
func (o MOptions) Inline(value bool) MOptions {
	return o.set("inline", value)
}

// InputClassName sets the input class name
func (o MOptions) InputClassName(value string) MOptions {
	return o.set("inputClassName", value)
}

// ItemHeight sets the height of each item for virtual rendering
func (o MOptions) ItemHeight(value string) MOptions {
	return o.set("itemHeight", value)
}

// JoinValues sets joinValues, concatenates selected values with commas
func (o MOptions) JoinValues(value bool) MOptions {
	return o.set("joinValues", value)
}

// Label sets the label
func (o MOptions) Label(value string) MOptions {
	return o.set("label", value)
}

// LabelAlign sets the label alignment, effective when mode is horizontal
func (o MOptions) LabelAlign(value string) MOptions {
	return o.set("labelAlign", value)
}

// LabelClassName sets the label class name
func (o MOptions) LabelClassName(value string) MOptions {
	return o.set("labelClassName", value)
}

// LabelField sets the field used as label
func (o MOptions) LabelField(value string) MOptions {
	return o.set("labelField", value)
}

// LabelRemark sets the label remark
func (o MOptions) LabelRemark(value string) MOptions {
	return o.set("labelRemark", value)
}

// Multiple sets the multiple selection state
func (o MOptions) Multiple(value bool) MOptions {
	return o.set("multiple", value)
}

// Name sets the field name for submission
func (o MOptions) Name(value string) MOptions {
	return o.set("name", value)
}

// Options sets the options group
func (o MOptions) Options(value ...any) MOptions {
	return o.set("options", value)
}

// Required sets the required state
func (o MOptions) Required(value bool) MOptions {
	return o.set("required", value)
}

// RequiredOn sets the condition for required state
func (o MOptions) RequiredOn(value string) MOptions {
	return o.set("requiredOn", value)
}

// Source sets the source for options group
func (o MOptions) Source(value string) MOptions {
	return o.set("source", value)
}

// SubmitOnChange sets the submit on change state
func (o MOptions) SubmitOnChange(value bool) MOptions {
	return o.set("submitOnChange", value)
}

// ValidateApi sets the validation API
func (o MOptions) ValidateApi(value string) MOptions {
	return o.set("validateApi", value)
}

// Validations sets the validation rules
func (o MOptions) Validations(value string) MOptions {
	return o.set("validations", value)
}

// Value sets the default value
func (o MOptions) Value(value string) MOptions {
	return o.set("value", value)
}

// ValueField sets the field used as value
func (o MOptions) ValueField(value string) MOptions {
	return o.set("valueField", value)
}

// ValuesNoWrap sets the no-wrap state for multiple values
func (o MOptions) ValuesNoWrap(value bool) MOptions {
	return o.set("valuesNoWrap", value)
}

// VirtualThreshold sets the threshold for virtual rendering
func (o MOptions) VirtualThreshold(value string) MOptions {
	return o.set("virtualThreshold", value)
}

// Visible sets the visibility state
func (o MOptions) Visible(value bool) MOptions {
	return o.set("visible", value)
}

// VisibleOn sets the condition for visibility
func (o MOptions) VisibleOn(value string) MOptions {
	return o.set("visibleOn", value)
}
