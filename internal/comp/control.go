package comp

import "github.com/zrcoder/amisgo/schema"

// Control represents a form item wrapper. Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/form/Control
type Control schema.Schema

func NewControl() Control {
	return Control{"type": "control"}
}

func (fc Control) set(key string, value any) Control {
	fc[key] = value
	return fc
}

// AutoFill configures automatic filling, synchronizing other values from the selected option to the form
func (fc Control) AutoFill(value string) Control {
	return fc.set("autoFill", value)
}

// Body sets the FormItem content
func (fc Control) Body(value ...any) Control {
	return fc.set("body", value)
}

// ClassName sets the container CSS class name (supports string or object configuration with conditional expressions)
func (fc Control) ClassName(value string) Control {
	return fc.set("className", value)
}

// ClearValueOnHidden determines whether to remove the form item's value when hidden, including values of unhidden items with the same name
func (fc Control) ClearValueOnHidden(value bool) Control {
	return fc.set("clearValueOnHidden", value)
}

// Desc sets a short description
func (fc Control) Desc(value string) Control {
	return fc.set("desc", value)
}

// Description sets detailed description content, supporting HTML fragments
func (fc Control) Description(value string) Control {
	return fc.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description
func (fc Control) DescriptionClassName(value string) Control {
	return fc.set("descriptionClassName", value)
}

// Disabled enables or disables the component
func (fc Control) Disabled(value bool) Control {
	return fc.set("disabled", value)
}

// DisabledOn sets a conditional expression for disabling the component (e.g., `data.xxx > 5`)
func (fc Control) DisabledOn(value string) Control {
	return fc.set("disabledOn", value)
}

// EditorSetting configures editor-specific settings (can be ignored at runtime)
func (fc Control) EditorSetting(value string) Control {
	return fc.set("editorSetting", value)
}

// ExtraName sets an additional field name for flattening values in range components
func (fc Control) ExtraName(value string) Control {
	return fc.set("extraName", value)
}

// Hidden controls the overall visibility of the component
func (fc Control) Hidden(value bool) Control {
	return fc.set("hidden", value)
}

// HiddenOn sets a conditional expression for hiding the component (e.g., `data.xxx > 5`)
func (fc Control) HiddenOn(value string) Control {
	return fc.set("hiddenOn", value)
}

// Hint sets an input hint displayed on focus
func (fc Control) Hint(value string) Control {
	return fc.set("hint", value)
}

// Horizontal configures specific left-right distribution for horizontal layout
func (fc Control) Horizontal(value string) Control {
	return fc.set("horizontal", value)
}

// ID sets a unique identifier, primarily used for log collection
func (fc Control) ID(value string) Control {
	return fc.set("id", value)
}

// InitAutoFill configures initial automatic filling behavior
func (fc Control) InitAutoFill(value string) Control {
	return fc.set("initAutoFill", value)
}

// Inline enables or disables inline form control mode
func (fc Control) Inline(value bool) Control {
	return fc.set("inline", value)
}

// InputClassName sets the CSS class name for input elements
func (fc Control) InputClassName(value string) Control {
	return fc.set("inputClassName", value)
}

// Label sets the description title
func (fc Control) Label(value string) Control {
	return fc.set("label", value)
}

// LabelAlign sets the alignment for the description title (options: right | left | top | inherit)
func (fc Control) LabelAlign(value string) Control {
	return fc.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label
func (fc Control) LabelClassName(value string) Control {
	return fc.set("labelClassName", value)
}

// LabelRemark adds a small tooltip icon next to the label
func (fc Control) LabelRemark(value string) Control {
	return fc.set("labelRemark", value)
}

// LabelWidth sets a custom width for the label (default unit is px)
func (fc Control) LabelWidth(value string) Control {
	return fc.set("labelWidth", value)
}

// Mode sets the display mode for the form item (options: normal | inline | horizontal)
func (fc Control) Mode(value string) Control {
	return fc.set("mode", value)
}

// Name sets the field name for form submission, supporting multi-level naming with dot notation (e.g., a.b.c)
func (fc Control) Name(value string) Control {
	return fc.set("name", value)
}

// OnEvent configures event-driven actions for the component
func (fc Control) OnEvent(value Event) Control {
	return fc.set("onEvent", value)
}

// Placeholder sets the input placeholder text
func (fc Control) Placeholder(value string) Control {
	return fc.set("placeholder", value)
}

// ReadOnly enables or disables read-only mode for the component
func (fc Control) ReadOnly(value bool) Control {
	return fc.set("readOnly", value)
}

// ReadOnlyOn sets a conditional expression for read-only state
func (fc Control) ReadOnlyOn(value string) Control {
	return fc.set("readOnlyOn", value)
}

// Remark adds a small tooltip icon with additional information on hover
func (fc Control) Remark(value string) Control {
	return fc.set("remark", value)
}

// Required determines whether the field is mandatory
func (fc Control) Required(value bool) Control {
	return fc.set("required", value)
}

// Row sets the number of rows (for multi-line inputs)
func (fc Control) Row(value string) Control {
	return fc.set("row", value)
}

// SaveImmediately enables immediate saving (used in TableColumn)
func (fc Control) SaveImmediately(value bool) Control {
	return fc.set("saveImmediately", value)
}

// Size sets the form item size (options: xs | sm | md | lg | full)
func (fc Control) Size(value string) Control {
	return fc.set("size", value)
}

// Static enables or disables static display mode
func (fc Control) Static(value bool) Control {
	return fc.set("static", value)
}

// StaticClassName sets the CSS class name for static form item display (supports string or object configuration with conditional expressions)
func (fc Control) StaticClassName(value string) Control {
	return fc.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input value display (supports string or object configuration with conditional expressions)
func (fc Control) StaticInputClassName(value string) Control {
	return fc.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display (supports string or object configuration with conditional expressions)
func (fc Control) StaticLabelClassName(value string) Control {
	return fc.set("staticLabelClassName", value)
}

// StaticOn sets a conditional expression for static display (e.g., `data.xxx > 5`)
func (fc Control) StaticOn(value string) Control {
	return fc.set("staticOn", value)
}

// StaticPlaceholder defines a placeholder for empty static values
func (fc Control) StaticPlaceholder(value string) Control {
	return fc.set("staticPlaceholder", value)
}

// StaticSchema sets the schema.Schema for static display
func (fc Control) StaticSchema(value string) Control {
	return fc.set("staticSchema", value)
}

// Style sets custom inline styles for the component
func (fc Control) Style(value any) Control {
	return fc.set("style", value)
}

// SubmitOnChange determines whether to submit the form when the value is modified
func (fc Control) SubmitOnChange(value bool) Control {
	return fc.set("submitOnChange", value)
}

// TestIdBuilder configures test ID generation
func (fc Control) TestIdBuilder(value string) Control {
	return fc.set("testIdBuilder", value)
}

// UseMobileUI enables or disables mobile-specific UI styling at the component level
func (fc Control) UseMobileUI(value bool) Control {
	return fc.set("useMobileUI", value)
}

// ValidateApi sets the remote validation endpoint for the form item
func (fc Control) ValidateApi(value string) Control {
	return fc.set("validateApi", value)
}

// ValidateOnChange configures validation behavior on modification (by default, validation occurs after form submission)
func (fc Control) ValidateOnChange(value bool) Control {
	return fc.set("validateOnChange", value)
}

// ValidationErrors sets the error messages for validation failures
func (fc Control) ValidationErrors(value string) Control {
	return fc.set("validationErrors", value)
}

// Validations sets the validation rules
func (fc Control) Validations(value string) Control {
	return fc.set("validations", value)
}

// Value sets the default value (static only, not supporting variable retrieval; data association is done via the name attribute)
func (fc Control) Value(value string) Control {
	return fc.set("value", value)
}

// Visible controls the overall visibility of the component
func (fc Control) Visible(value bool) Control {
	return fc.set("visible", value)
}

// VisibleOn sets a conditional expression for component visibility (e.g., `data.xxx > 5`)
func (fc Control) VisibleOn(value string) Control {
	return fc.set("visibleOn", value)
}

// Width adjusts the component width (used in Table)
func (fc Control) Width(value string) Control {
	return fc.set("width", value)
}
