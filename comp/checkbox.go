package comp

import "github.com/zrcoder/amisgo/model"

// Checkbox represents a checkbox form component for boolean selections
type checkbox model.Schema

// Checkbox creates a new Checkbox instance
func Checkbox() checkbox {
	return checkbox{"type": "checkbox"}
}

func (c checkbox) set(key string, value any) checkbox {
	c[key] = value
	return c
}

// AutoFill enables automatic synchronization of selected values to the form
func (c checkbox) AutoFill(value string) checkbox {
	return c.set("autoFill", value)
}

// Badge sets a badge or indicator for the checkbox
func (c checkbox) Badge(value string) checkbox {
	return c.set("badge", value)
}

// Checked sets the initial checked state of the checkbox
func (c checkbox) Checked(value bool) checkbox {
	return c.set("checked", value)
}

// ClassName sets the CSS class name for the checkbox container
func (c checkbox) ClassName(value string) checkbox {
	return c.set("className", value)
}

// ClearValueOnHidden determines whether to remove the form item's value when hidden
func (c checkbox) ClearValueOnHidden(value bool) checkbox {
	return c.set("clearValueOnHidden", value)
}

// Desc sets a brief description for the checkbox
func (c checkbox) Desc(value string) checkbox {
	return c.set("desc", value)
}

// Description sets a detailed description that supports HTML fragments
func (c checkbox) Description(value string) checkbox {
	return c.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description element
func (c checkbox) DescriptionClassName(value string) checkbox {
	return c.set("descriptionClassName", value)
}

// Disabled enables or disables the checkbox component
func (c checkbox) Disabled(value bool) checkbox {
	return c.set("disabled", value)
}

// DisabledOn sets a conditional expression to dynamically disable the checkbox
func (c checkbox) DisabledOn(value string) checkbox {
	return c.set("disabledOn", value)
}

// EditorSetting configures editor-specific settings for the component
func (c checkbox) EditorSetting(value string) checkbox {
	return c.set("editorSetting", value)
}

// ExtraName sets an additional field name for the component
func (c checkbox) ExtraName(value string) checkbox {
	return c.set("extraName", value)
}

// FalseValue sets the value when the checkbox is unchecked
func (c checkbox) FalseValue(value string) checkbox {
	return c.set("falseValue", value)
}

// Hidden controls the visibility of the checkbox component
func (c checkbox) Hidden(value bool) checkbox {
	return c.set("hidden", value)
}

// HiddenOn sets a conditional expression to dynamically hide the checkbox
func (c checkbox) HiddenOn(value string) checkbox {
	return c.set("hiddenOn", value)
}

// Hint provides an input hint or placeholder text
func (c checkbox) Hint(value string) checkbox {
	return c.set("hint", value)
}

// Horizontal configures the specific left and right distribution when using horizontal layout
func (c checkbox) Horizontal(value string) checkbox {
	return c.set("horizontal", value)
}

// ID sets a unique identifier for the component
func (c checkbox) ID(value string) checkbox {
	return c.set("id", value)
}

// InitAutoFill initializes automatic value filling
func (c checkbox) InitAutoFill(value string) checkbox {
	return c.set("initAutoFill", value)
}

// Inline determines if the form control should use inline mode
func (c checkbox) Inline(value bool) checkbox {
	return c.set("inline", value)
}

// InputClassName sets the CSS class name for the input element
func (c checkbox) InputClassName(value string) checkbox {
	return c.set("inputClassName", value)
}

// Label sets the description label for the checkbox
func (c checkbox) Label(value string) checkbox {
	return c.set("label", value)
}

// LabelAlign sets the alignment style for the description label
func (c checkbox) LabelAlign(value string) checkbox {
	return c.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label element
func (c checkbox) LabelClassName(value string) checkbox {
	return c.set("labelClassName", value)
}

// LabelRemark adds a small icon that displays a tooltip when hovered
func (c checkbox) LabelRemark(value string) checkbox {
	return c.set("labelRemark", value)
}

// LabelWidth sets a custom width for the label
func (c checkbox) LabelWidth(value string) checkbox {
	return c.set("labelWidth", value)
}

// Mode configures the current form item display mode
func (c checkbox) Mode(value string) checkbox {
	return c.set("mode", value)
}

// Name sets the field name for the component
func (c checkbox) Name(value string) checkbox {
	return c.set("name", value)
}

// OnEvent configures event-driven actions for the checkbox
func (c checkbox) OnEvent(value any) checkbox {
	return c.set("onEvent", value)
}

// Option sets a specific option for the checkbox
func (c checkbox) Option(value string) checkbox {
	return c.set("option", value)
}

// OptionType configures the type of option for the checkbox
func (c checkbox) OptionType(value string) checkbox {
	return c.set("optionType", value)
}

// Partial enables or disables partial selection mode
func (c checkbox) Partial(value bool) checkbox {
	return c.set("partial", value)
}

// Placeholder sets the input placeholder text
func (c checkbox) Placeholder(value string) checkbox {
	return c.set("placeholder", value)
}

// ReadOnly enables or disables read-only mode
func (c checkbox) ReadOnly(value bool) checkbox {
	return c.set("readOnly", value)
}

// ReadOnlyOn sets a conditional expression for read-only state
func (c checkbox) ReadOnlyOn(value string) checkbox {
	return c.set("readOnlyOn", value)
}

// Remark adds a small icon that displays a tooltip when hovered
func (c checkbox) Remark(value string) checkbox {
	return c.set("remark", value)
}

// Required determines if the checkbox is mandatory
func (c checkbox) Required(value bool) checkbox {
	return c.set("required", value)
}

// Row sets row-related configuration (context-specific)
func (c checkbox) Row(value string) checkbox {
	return c.set("row", value)
}

// SaveImmediately enables immediate saving (used in TableColumn)
func (c checkbox) SaveImmediately(value bool) checkbox {
	return c.set("saveImmediately", value)
}

// Size sets the form item size
func (c checkbox) Size(value string) checkbox {
	return c.set("size", value)
}

// Static enables or disables static display mode
func (c checkbox) Static(value bool) checkbox {
	return c.set("static", value)
}

// StaticClassName sets the CSS class name for static form item display
func (c checkbox) StaticClassName(value string) checkbox {
	return c.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input value display
func (c checkbox) StaticInputClassName(value string) checkbox {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (c checkbox) StaticLabelClassName(value string) checkbox {
	return c.set("staticLabelClassName", value)
}

// StaticOn sets a conditional expression for static display
func (c checkbox) StaticOn(value string) checkbox {
	return c.set("staticOn", value)
}

// StaticPlaceholder defines a placeholder for empty static values
func (c checkbox) StaticPlaceholder(value string) checkbox {
	return c.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (c checkbox) StaticSchema(value string) checkbox {
	return c.set("staticSchema", value)
}

// Style sets custom inline styles for the component
func (c checkbox) Style(value any) checkbox {
	return c.set("style", value)
}

// SubmitOnChange determines if the form should be submitted after value modification
func (c checkbox) SubmitOnChange(value bool) checkbox {
	return c.set("submitOnChange", value)
}

// TestIdBuilder configures test ID generation
func (c checkbox) TestIdBuilder(value string) checkbox {
	return c.set("testIdBuilder", value)
}

// TrueValue sets the value when the checkbox is checked
func (c checkbox) TrueValue(value string) checkbox {
	return c.set("trueValue", value)
}

// UseMobileUI enables or disables mobile-specific UI styling
func (c checkbox) UseMobileUI(value bool) checkbox {
	return c.set("useMobileUI", value)
}

// ValidateApi sets the remote validation endpoint for the form item
func (c checkbox) ValidateApi(value string) checkbox {
	return c.set("validateApi", value)
}

// ValidateOnChange configures validation behavior after form submission
func (c checkbox) ValidateOnChange(value bool) checkbox {
	return c.set("validateOnChange", value)
}

// ValidationErrors sets the error messages for validation failures
func (c checkbox) ValidationErrors(value string) checkbox {
	return c.set("validationErrors", value)
}

// Validations sets validation rules
func (c checkbox) Validations(value string) checkbox {
	return c.set("validations", value)
}

// Value sets the default value for the checkbox
func (c checkbox) Value(value string) checkbox {
	return c.set("value", value)
}

// Visible controls the overall visibility of the component
func (c checkbox) Visible(value bool) checkbox {
	return c.set("visible", value)
}

// VisibleOn sets a conditional expression for component visibility
func (c checkbox) VisibleOn(value string) checkbox {
	return c.set("visibleOn", value)
}

// Width sets the width of the checkbox component
func (c checkbox) Width(value string) checkbox {
	return c.set("width", value)
}
