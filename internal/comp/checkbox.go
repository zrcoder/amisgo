package comp

import "github.com/zrcoder/amisgo/schema"

// Checkbox represents a Checkbox form component for boolean selections
type Checkbox schema.Schema

func NewCheckbox() Checkbox {
	return Checkbox{"type": "checkbox"}
}

func (c Checkbox) set(key string, value any) Checkbox {
	c[key] = value
	return c
}

// AutoFill enables automatic synchronization of selected values to the form
func (c Checkbox) AutoFill(value string) Checkbox {
	return c.set("autoFill", value)
}

// Badge sets a badge or indicator for the checkbox
func (c Checkbox) Badge(value string) Checkbox {
	return c.set("badge", value)
}

// Checked sets the initial checked state of the checkbox
func (c Checkbox) Checked(value bool) Checkbox {
	return c.set("checked", value)
}

// ClassName sets the CSS class name for the checkbox container
func (c Checkbox) ClassName(value string) Checkbox {
	return c.set("className", value)
}

// ClearValueOnHidden determines whether to remove the form item's value when hidden
func (c Checkbox) ClearValueOnHidden(value bool) Checkbox {
	return c.set("clearValueOnHidden", value)
}

// Desc sets a brief description for the checkbox
func (c Checkbox) Desc(value string) Checkbox {
	return c.set("desc", value)
}

// Description sets a detailed description that supports HTML fragments
func (c Checkbox) Description(value string) Checkbox {
	return c.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description element
func (c Checkbox) DescriptionClassName(value string) Checkbox {
	return c.set("descriptionClassName", value)
}

// Disabled enables or disables the checkbox component
func (c Checkbox) Disabled(value bool) Checkbox {
	return c.set("disabled", value)
}

// DisabledOn sets a conditional expression to dynamically disable the checkbox
func (c Checkbox) DisabledOn(value string) Checkbox {
	return c.set("disabledOn", value)
}

// EditorSetting configures editor-specific settings for the component
func (c Checkbox) EditorSetting(value string) Checkbox {
	return c.set("editorSetting", value)
}

// ExtraName sets an additional field name for the component
func (c Checkbox) ExtraName(value string) Checkbox {
	return c.set("extraName", value)
}

// FalseValue sets the value when the checkbox is unchecked
func (c Checkbox) FalseValue(value string) Checkbox {
	return c.set("falseValue", value)
}

// Hidden controls the visibility of the checkbox component
func (c Checkbox) Hidden(value bool) Checkbox {
	return c.set("hidden", value)
}

// HiddenOn sets a conditional expression to dynamically hide the checkbox
func (c Checkbox) HiddenOn(value string) Checkbox {
	return c.set("hiddenOn", value)
}

// Hint provides an input hint or placeholder text
func (c Checkbox) Hint(value string) Checkbox {
	return c.set("hint", value)
}

// Horizontal configures the specific left and right distribution when using horizontal layout
func (c Checkbox) Horizontal(value string) Checkbox {
	return c.set("horizontal", value)
}

// ID sets a unique identifier for the component
func (c Checkbox) ID(value string) Checkbox {
	return c.set("id", value)
}

// InitAutoFill initializes automatic value filling
func (c Checkbox) InitAutoFill(value string) Checkbox {
	return c.set("initAutoFill", value)
}

// Inline determines if the form control should use inline mode
func (c Checkbox) Inline(value bool) Checkbox {
	return c.set("inline", value)
}

// InputClassName sets the CSS class name for the input element
func (c Checkbox) InputClassName(value string) Checkbox {
	return c.set("inputClassName", value)
}

// Label sets the description label for the checkbox
func (c Checkbox) Label(value string) Checkbox {
	return c.set("label", value)
}

// LabelAlign sets the alignment style for the description label
func (c Checkbox) LabelAlign(value string) Checkbox {
	return c.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label element
func (c Checkbox) LabelClassName(value string) Checkbox {
	return c.set("labelClassName", value)
}

// LabelRemark adds a small icon that displays a tooltip when hovered
func (c Checkbox) LabelRemark(value string) Checkbox {
	return c.set("labelRemark", value)
}

// LabelWidth sets a custom width for the label
func (c Checkbox) LabelWidth(value string) Checkbox {
	return c.set("labelWidth", value)
}

// Mode configures the current form item display mode
func (c Checkbox) Mode(value string) Checkbox {
	return c.set("mode", value)
}

// Name sets the field name for the component
func (c Checkbox) Name(value string) Checkbox {
	return c.set("name", value)
}

// OnEvent configures event-driven actions for the checkbox
func (c Checkbox) OnEvent(value any) Checkbox {
	return c.set("onEvent", value)
}

// Option sets a specific option for the checkbox
func (c Checkbox) Option(value string) Checkbox {
	return c.set("option", value)
}

// OptionType configures the type of option for the checkbox
func (c Checkbox) OptionType(value string) Checkbox {
	return c.set("optionType", value)
}

// Partial enables or disables partial selection mode
func (c Checkbox) Partial(value bool) Checkbox {
	return c.set("partial", value)
}

// Placeholder sets the input placeholder text
func (c Checkbox) Placeholder(value string) Checkbox {
	return c.set("placeholder", value)
}

// ReadOnly enables or disables read-only mode
func (c Checkbox) ReadOnly(value bool) Checkbox {
	return c.set("readOnly", value)
}

// ReadOnlyOn sets a conditional expression for read-only state
func (c Checkbox) ReadOnlyOn(value string) Checkbox {
	return c.set("readOnlyOn", value)
}

// Remark adds a small icon that displays a tooltip when hovered
func (c Checkbox) Remark(value string) Checkbox {
	return c.set("remark", value)
}

// Required determines if the checkbox is mandatory
func (c Checkbox) Required(value bool) Checkbox {
	return c.set("required", value)
}

// Row sets row-related configuration (context-specific)
func (c Checkbox) Row(value string) Checkbox {
	return c.set("row", value)
}

// SaveImmediately enables immediate saving (used in TableColumn)
func (c Checkbox) SaveImmediately(value bool) Checkbox {
	return c.set("saveImmediately", value)
}

// Size sets the form item size
func (c Checkbox) Size(value string) Checkbox {
	return c.set("size", value)
}

// Static enables or disables static display mode
func (c Checkbox) Static(value bool) Checkbox {
	return c.set("static", value)
}

// StaticClassName sets the CSS class name for static form item display
func (c Checkbox) StaticClassName(value string) Checkbox {
	return c.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input value display
func (c Checkbox) StaticInputClassName(value string) Checkbox {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (c Checkbox) StaticLabelClassName(value string) Checkbox {
	return c.set("staticLabelClassName", value)
}

// StaticOn sets a conditional expression for static display
func (c Checkbox) StaticOn(value string) Checkbox {
	return c.set("staticOn", value)
}

// StaticPlaceholder defines a placeholder for empty static values
func (c Checkbox) StaticPlaceholder(value string) Checkbox {
	return c.set("staticPlaceholder", value)
}

// StaticSchema sets the schema.Schema for static display
func (c Checkbox) StaticSchema(value string) Checkbox {
	return c.set("staticSchema", value)
}

// Style sets custom inline styles for the component
func (c Checkbox) Style(value any) Checkbox {
	return c.set("style", value)
}

// SubmitOnChange determines if the form should be submitted after value modification
func (c Checkbox) SubmitOnChange(value bool) Checkbox {
	return c.set("submitOnChange", value)
}

// TestIdBuilder configures test ID generation
func (c Checkbox) TestIdBuilder(value string) Checkbox {
	return c.set("testIdBuilder", value)
}

// TrueValue sets the value when the checkbox is checked
func (c Checkbox) TrueValue(value string) Checkbox {
	return c.set("trueValue", value)
}

// UseMobileUI enables or disables mobile-specific UI styling
func (c Checkbox) UseMobileUI(value bool) Checkbox {
	return c.set("useMobileUI", value)
}

// ValidateApi sets the remote validation endpoint for the form item
func (c Checkbox) ValidateApi(value string) Checkbox {
	return c.set("validateApi", value)
}

// ValidateOnChange configures validation behavior after form submission
func (c Checkbox) ValidateOnChange(value bool) Checkbox {
	return c.set("validateOnChange", value)
}

// ValidationErrors sets the error messages for validation failures
func (c Checkbox) ValidationErrors(value string) Checkbox {
	return c.set("validationErrors", value)
}

// Validations sets validation rules
func (c Checkbox) Validations(value string) Checkbox {
	return c.set("validations", value)
}

// Value sets the default value for the checkbox
func (c Checkbox) Value(value string) Checkbox {
	return c.set("value", value)
}

// Visible controls the overall visibility of the component
func (c Checkbox) Visible(value bool) Checkbox {
	return c.set("visible", value)
}

// VisibleOn sets a conditional expression for component visibility
func (c Checkbox) VisibleOn(value string) Checkbox {
	return c.set("visibleOn", value)
}

// Width sets the width of the checkbox component
func (c Checkbox) Width(value string) Checkbox {
	return c.set("width", value)
}
