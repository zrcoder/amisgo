package comp

// formula represents a formula control component
type formula Schema

// Formula creates a new FormulaControl instance
func Formula() formula {
	return make(formula).set("type", "formula")
}

func (fc formula) set(key string, value any) formula {
	fc[key] = value
	return fc
}

// AutoFill sets the autoFill property
func (fc formula) AutoFill(value string) formula {
	return fc.set("autoFill", value)
}

// AutoSet sets the autoSet property
func (fc formula) AutoSet(value bool) formula {
	return fc.set("autoSet", value)
}

// ClassName sets the CSS class name
func (fc formula) ClassName(value string) formula {
	return fc.set("className", value)
}

// ClearValueOnHidden sets whether to clear value when hidden
func (fc formula) ClearValueOnHidden(value bool) formula {
	return fc.set("clearValueOnHidden", value)
}

// Condition sets the condition to trigger the formula
func (fc formula) Condition(value string) formula {
	return fc.set("condition", value)
}

// Desc sets the description
func (fc formula) Desc(value string) formula {
	return fc.set("desc", value)
}

// Description sets the HTML description
func (fc formula) Description(value string) formula {
	return fc.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description
func (fc formula) DescriptionClassName(value string) formula {
	return fc.set("descriptionClassName", value)
}

// Disabled sets whether the component is disabled
func (fc formula) Disabled(value bool) formula {
	return fc.set("disabled", value)
}

// DisabledOn sets the expression to disable the component
func (fc formula) DisabledOn(value string) formula {
	return fc.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (fc formula) EditorSetting(value string) formula {
	return fc.set("editorSetting", value)
}

// ExtraName sets an extra field name
func (fc formula) ExtraName(value string) formula {
	return fc.set("extraName", value)
}

// Formula sets the formula
func (fc formula) Formula(value string) formula {
	return fc.set("formula", value)
}

// Hidden sets whether the component is hidden
func (fc formula) Hidden(value bool) formula {
	return fc.set("hidden", value)
}

// HiddenOn sets the expression to hide the component
func (fc formula) HiddenOn(value string) formula {
	return fc.set("hiddenOn", value)
}

// Hint sets the input hint
func (fc formula) Hint(value string) formula {
	return fc.set("hint", value)
}

// Horizontal sets the horizontal layout configuration
func (fc formula) Horizontal(value string) formula {
	return fc.set("horizontal", value)
}

// ID sets the ID of the button to trigger the formula
func (fc formula) ID(value string) formula {
	return fc.set("id", value)
}

// InitAutoFill sets the initial auto-fill value
func (fc formula) InitAutoFill(value string) formula {
	return fc.set("initAutoFill", value)
}

// InitSet sets whether to apply initially
func (fc formula) InitSet(value bool) formula {
	return fc.set("initSet", value)
}

// Inline sets whether the control is inline
func (fc formula) Inline(value bool) formula {
	return fc.set("inline", value)
}

// InputClassName sets the CSS class name for the input
func (fc formula) InputClassName(value string) formula {
	return fc.set("inputClassName", value)
}

// Label sets the label
func (fc formula) Label(value string) formula {
	return fc.set("label", value)
}

// LabelAlign sets the label alignment
func (fc formula) LabelAlign(value string) formula {
	return fc.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label
func (fc formula) LabelClassName(value string) formula {
	return fc.set("labelClassName", value)
}

// LabelRemark sets the label remark
func (fc formula) LabelRemark(value string) formula {
	return fc.set("labelRemark", value)
}

// LabelWidth sets the custom width for the label
func (fc formula) LabelWidth(value string) formula {
	return fc.set("labelWidth", value)
}

// Mode sets the display mode
func (fc formula) Mode(value string) formula {
	return fc.set("mode", value)
}

// Name sets the field name
func (fc formula) Name(value string) formula {
	return fc.set("name", value)
}

// OnEvent sets the event configuration
func (fc formula) OnEvent(value any) formula {
	return fc.set("onEvent", value)
}

// Placeholder sets the placeholder
func (fc formula) Placeholder(value string) formula {
	return fc.set("placeholder", value)
}

// ReadOnly sets whether the component is read-only
func (fc formula) ReadOnly(value bool) formula {
	return fc.set("readOnly", value)
}

// ReadOnlyOn sets the expression to make the component read-only
func (fc formula) ReadOnlyOn(value string) formula {
	return fc.set("readOnlyOn", value)
}

// Remark sets the remark
func (fc formula) Remark(value string) formula {
	return fc.set("remark", value)
}

// Required sets whether the field is required
func (fc formula) Required(value bool) formula {
	return fc.set("required", value)
}

// Row sets the row configuration
func (fc formula) Row(value string) formula {
	return fc.set("row", value)
}

// SaveImmediately sets whether to save immediately
func (fc formula) SaveImmediately(value bool) formula {
	return fc.set("saveImmediately", value)
}

// Size sets the size of the form item
func (fc formula) Size(value string) formula {
	return fc.set("size", value)
}

// Static sets whether the component is static
func (fc formula) Static(value bool) formula {
	return fc.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (fc formula) StaticClassName(value string) formula {
	return fc.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input
func (fc formula) StaticInputClassName(value string) formula {
	return fc.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label
func (fc formula) StaticLabelClassName(value string) formula {
	return fc.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display
func (fc formula) StaticOn(value string) formula {
	return fc.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (fc formula) StaticPlaceholder(value string) formula {
	return fc.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (fc formula) StaticSchema(value string) formula {
	return fc.set("staticSchema", value)
}

// Style sets the component style
func (fc formula) Style(value any) formula {
	return fc.set("style", value)
}

// SubmitOnChange sets whether to submit the form on change
func (fc formula) SubmitOnChange(value bool) formula {
	return fc.set("submitOnChange", value)
}

// TestIdBuilder sets the test ID builder
func (fc formula) TestIdBuilder(value string) formula {
	return fc.set("testIdBuilder", value)
}

// UseMobileUI sets whether to use mobile UI
func (fc formula) UseMobileUI(value bool) formula {
	return fc.set("useMobileUI", value)
}

// ValidateApi sets the API for remote validation
func (fc formula) ValidateApi(value string) formula {
	return fc.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on change
func (fc formula) ValidateOnChange(value bool) formula {
	return fc.set("validateOnChange", value)
}

// ValidationErrors sets the validation error messages
func (fc formula) ValidationErrors(value string) formula {
	return fc.set("validationErrors", value)
}

// Validations sets the validations
func (fc formula) Validations(value string) formula {
	return fc.set("validations", value)
}

// Value sets the default value
func (fc formula) Value(value string) formula {
	return fc.set("value", value)
}

// Visible sets whether the component is visible
func (fc formula) Visible(value bool) formula {
	return fc.set("visible", value)
}

// VisibleOn sets the expression to make the component visible
func (fc formula) VisibleOn(value string) formula {
	return fc.set("visibleOn", value)
}

// Width sets the width in a table
func (fc formula) Width(value string) formula {
	return fc.set("width", value)
}
