package model

type Options Schema

func NewOptions() Options {
	return Options{}
}

// set sets a field value
func (o Options) set(key string, value any) Options {
	o[key] = value
	return o
}

// AutoFill sets autoFill value
func (o Options) AutoFill(value string) Options {
	return o.set("autoFill", value)
}

// ClassName sets the outer class name
func (o Options) ClassName(value string) Options {
	return o.set("className", value)
}

// Description sets the description
func (o Options) Description(value string) Options {
	return o.set("description", value)
}

// Disabled sets the disabled state
func (o Options) Disabled(value bool) Options {
	return o.set("disabled", value)
}

// DisabledOn sets the condition for disabling
func (o Options) DisabledOn(value string) Options {
	return o.set("disabledOn", value)
}

// ExtractValue sets extractValue, effective when joinValues is true
func (o Options) ExtractValue(value bool) Options {
	return o.set("extractValue", value)
}

// Inline sets the inline state
func (o Options) Inline(value bool) Options {
	return o.set("inline", value)
}

// InputClassName sets the input class name
func (o Options) InputClassName(value string) Options {
	return o.set("inputClassName", value)
}

// ItemHeight sets the height of each item for virtual rendering
func (o Options) ItemHeight(value string) Options {
	return o.set("itemHeight", value)
}

// JoinValues sets joinValues, concatenates selected values with commas
func (o Options) JoinValues(value bool) Options {
	return o.set("joinValues", value)
}

// Label sets the label
func (o Options) Label(value string) Options {
	return o.set("label", value)
}

// LabelAlign sets the label alignment, effective when mode is horizontal
func (o Options) LabelAlign(value string) Options {
	return o.set("labelAlign", value)
}

// LabelClassName sets the label class name
func (o Options) LabelClassName(value string) Options {
	return o.set("labelClassName", value)
}

// LabelField sets the field used as label
func (o Options) LabelField(value string) Options {
	return o.set("labelField", value)
}

// LabelRemark sets the label remark
func (o Options) LabelRemark(value string) Options {
	return o.set("labelRemark", value)
}

// Multiple sets the multiple selection state
func (o Options) Multiple(value bool) Options {
	return o.set("multiple", value)
}

// Name sets the field name for submission
func (o Options) Name(value string) Options {
	return o.set("name", value)
}

// Options sets the options group
func (o Options) Options(value ...any) Options {
	return o.set("options", value)
}

// Required sets the required state
func (o Options) Required(value bool) Options {
	return o.set("required", value)
}

// RequiredOn sets the condition for required state
func (o Options) RequiredOn(value string) Options {
	return o.set("requiredOn", value)
}

// Source sets the source for options group
func (o Options) Source(value string) Options {
	return o.set("source", value)
}

// SubmitOnChange sets the submit on change state
func (o Options) SubmitOnChange(value bool) Options {
	return o.set("submitOnChange", value)
}

// ValidateApi sets the validation API
func (o Options) ValidateApi(value string) Options {
	return o.set("validateApi", value)
}

// Validations sets the validation rules
func (o Options) Validations(value string) Options {
	return o.set("validations", value)
}

// Value sets the default value
func (o Options) Value(value string) Options {
	return o.set("value", value)
}

// ValueField sets the field used as value
func (o Options) ValueField(value string) Options {
	return o.set("valueField", value)
}

// ValuesNoWrap sets the no-wrap state for multiple values
func (o Options) ValuesNoWrap(value bool) Options {
	return o.set("valuesNoWrap", value)
}

// VirtualThreshold sets the threshold for virtual rendering
func (o Options) VirtualThreshold(value string) Options {
	return o.set("virtualThreshold", value)
}

// Visible sets the visibility state
func (o Options) Visible(value bool) Options {
	return o.set("visible", value)
}

// VisibleOn sets the condition for visibility
func (o Options) VisibleOn(value string) Options {
	return o.set("visibleOn", value)
}
