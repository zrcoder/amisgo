package comp

import "github.com/zrcoder/amisgo/schema"

// Formula represents a Formula control component
type Formula schema.Schema

func NewFormula() Formula {
	return Formula{"type": "formula"}
}

func (fc Formula) set(key string, value any) Formula {
	fc[key] = value
	return fc
}

// AutoFill sets the autoFill property
func (fc Formula) AutoFill(value string) Formula {
	return fc.set("autoFill", value)
}

// AutoSet sets the autoSet property
func (fc Formula) AutoSet(value bool) Formula {
	return fc.set("autoSet", value)
}

// ClassName sets the CSS class name
func (fc Formula) ClassName(value string) Formula {
	return fc.set("className", value)
}

// ClearValueOnHidden sets whether to clear value when hidden
func (fc Formula) ClearValueOnHidden(value bool) Formula {
	return fc.set("clearValueOnHidden", value)
}

// Condition sets the condition to trigger the formula
func (fc Formula) Condition(value string) Formula {
	return fc.set("condition", value)
}

// Desc sets the description
func (fc Formula) Desc(value string) Formula {
	return fc.set("desc", value)
}

// Description sets the HTML description
func (fc Formula) Description(value string) Formula {
	return fc.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description
func (fc Formula) DescriptionClassName(value string) Formula {
	return fc.set("descriptionClassName", value)
}

// Disabled sets whether the component is disabled
func (fc Formula) Disabled(value bool) Formula {
	return fc.set("disabled", value)
}

// DisabledOn sets the expression to disable the component
func (fc Formula) DisabledOn(value string) Formula {
	return fc.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (fc Formula) EditorSetting(value string) Formula {
	return fc.set("editorSetting", value)
}

// ExtraName sets an extra field name
func (fc Formula) ExtraName(value string) Formula {
	return fc.set("extraName", value)
}

// Formula sets the formula
func (fc Formula) Formula(value string) Formula {
	return fc.set("formula", value)
}

// Hidden sets whether the component is hidden
func (fc Formula) Hidden(value bool) Formula {
	return fc.set("hidden", value)
}

// HiddenOn sets the expression to hide the component
func (fc Formula) HiddenOn(value string) Formula {
	return fc.set("hiddenOn", value)
}

// Hint sets the input hint
func (fc Formula) Hint(value string) Formula {
	return fc.set("hint", value)
}

// Horizontal sets the horizontal layout configuration
func (fc Formula) Horizontal(value string) Formula {
	return fc.set("horizontal", value)
}

// ID sets the ID of the button to trigger the formula
func (fc Formula) ID(value string) Formula {
	return fc.set("id", value)
}

// InitAutoFill sets the initial auto-fill value
func (fc Formula) InitAutoFill(value string) Formula {
	return fc.set("initAutoFill", value)
}

// InitSet sets whether to apply initially
func (fc Formula) InitSet(value bool) Formula {
	return fc.set("initSet", value)
}

// Inline sets whether the control is inline
func (fc Formula) Inline(value bool) Formula {
	return fc.set("inline", value)
}

// InputClassName sets the CSS class name for the input
func (fc Formula) InputClassName(value string) Formula {
	return fc.set("inputClassName", value)
}

// Label sets the label
func (fc Formula) Label(value string) Formula {
	return fc.set("label", value)
}

// LabelAlign sets the label alignment
func (fc Formula) LabelAlign(value string) Formula {
	return fc.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label
func (fc Formula) LabelClassName(value string) Formula {
	return fc.set("labelClassName", value)
}

// LabelRemark sets the label remark
func (fc Formula) LabelRemark(value string) Formula {
	return fc.set("labelRemark", value)
}

// LabelWidth sets the custom width for the label
func (fc Formula) LabelWidth(value string) Formula {
	return fc.set("labelWidth", value)
}

// Mode sets the display mode
func (fc Formula) Mode(value string) Formula {
	return fc.set("mode", value)
}

// Name sets the field name
func (fc Formula) Name(value string) Formula {
	return fc.set("name", value)
}

// OnEvent sets the event configuration
func (fc Formula) OnEvent(value Event) Formula {
	return fc.set("onEvent", value)
}

// Placeholder sets the placeholder
func (fc Formula) Placeholder(value string) Formula {
	return fc.set("placeholder", value)
}

// ReadOnly sets whether the component is read-only
func (fc Formula) ReadOnly(value bool) Formula {
	return fc.set("readOnly", value)
}

// ReadOnlyOn sets the expression to make the component read-only
func (fc Formula) ReadOnlyOn(value string) Formula {
	return fc.set("readOnlyOn", value)
}

// Remark sets the remark
func (fc Formula) Remark(value string) Formula {
	return fc.set("remark", value)
}

// Required sets whether the field is required
func (fc Formula) Required(value bool) Formula {
	return fc.set("required", value)
}

// Row sets the row configuration
func (fc Formula) Row(value string) Formula {
	return fc.set("row", value)
}

// SaveImmediately sets whether to save immediately
func (fc Formula) SaveImmediately(value bool) Formula {
	return fc.set("saveImmediately", value)
}

// Size sets the size of the form item
func (fc Formula) Size(value string) Formula {
	return fc.set("size", value)
}

// Static sets whether the component is static
func (fc Formula) Static(value bool) Formula {
	return fc.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (fc Formula) StaticClassName(value string) Formula {
	return fc.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input
func (fc Formula) StaticInputClassName(value string) Formula {
	return fc.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label
func (fc Formula) StaticLabelClassName(value string) Formula {
	return fc.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display
func (fc Formula) StaticOn(value string) Formula {
	return fc.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (fc Formula) StaticPlaceholder(value string) Formula {
	return fc.set("staticPlaceholder", value)
}

// StaticSchema sets the schema.Schema for static display
func (fc Formula) StaticSchema(value string) Formula {
	return fc.set("staticSchema", value)
}

// Style sets the component style
func (fc Formula) Style(value any) Formula {
	return fc.set("style", value)
}

// SubmitOnChange sets whether to submit the form on change
func (fc Formula) SubmitOnChange(value bool) Formula {
	return fc.set("submitOnChange", value)
}

// TestIdBuilder sets the test ID builder
func (fc Formula) TestIdBuilder(value string) Formula {
	return fc.set("testIdBuilder", value)
}

// UseMobileUI sets whether to use mobile UI
func (fc Formula) UseMobileUI(value bool) Formula {
	return fc.set("useMobileUI", value)
}

// ValidateApi sets the API for remote validation
func (fc Formula) ValidateApi(value string) Formula {
	return fc.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on change
func (fc Formula) ValidateOnChange(value bool) Formula {
	return fc.set("validateOnChange", value)
}

// ValidationErrors sets the validation error messages
func (fc Formula) ValidationErrors(value string) Formula {
	return fc.set("validationErrors", value)
}

// Validations sets the validations
func (fc Formula) Validations(value string) Formula {
	return fc.set("validations", value)
}

// Value sets the default value
func (fc Formula) Value(value string) Formula {
	return fc.set("value", value)
}

// Visible sets whether the component is visible
func (fc Formula) Visible(value bool) Formula {
	return fc.set("visible", value)
}

// VisibleOn sets the expression to make the component visible
func (fc Formula) VisibleOn(value string) Formula {
	return fc.set("visibleOn", value)
}

// Width sets the width in a table
func (fc Formula) Width(value string) Formula {
	return fc.set("width", value)
}
