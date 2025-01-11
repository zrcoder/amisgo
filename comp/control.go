package comp

import "github.com/zrcoder/amisgo/model"

// Control represents a form item wrapper. Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/form/control
type control model.Schema

// Control creates a new Control instance with default type
func Control() control {
	return control{"type": "control"}
}

func (fc control) set(key string, value any) control {
	fc[key] = value
	return fc
}

// AutoFill configures automatic filling, synchronizing other values from the selected option to the form
func (fc control) AutoFill(value string) control {
	return fc.set("autoFill", value)
}

// Body sets the FormItem content
func (fc control) Body(value ...any) control {
	return fc.set("body", value)
}

// ClassName sets the container CSS class name (supports string or object configuration with conditional expressions)
func (fc control) ClassName(value string) control {
	return fc.set("className", value)
}

// ClearValueOnHidden determines whether to remove the form item's value when hidden, including values of unhidden items with the same name
func (fc control) ClearValueOnHidden(value bool) control {
	return fc.set("clearValueOnHidden", value)
}

// Desc sets a short description
func (fc control) Desc(value string) control {
	return fc.set("desc", value)
}

// Description sets detailed description content, supporting HTML fragments
func (fc control) Description(value string) control {
	return fc.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description
func (fc control) DescriptionClassName(value string) control {
	return fc.set("descriptionClassName", value)
}

// Disabled enables or disables the component
func (fc control) Disabled(value bool) control {
	return fc.set("disabled", value)
}

// DisabledOn sets a conditional expression for disabling the component (e.g., `data.xxx > 5`)
func (fc control) DisabledOn(value string) control {
	return fc.set("disabledOn", value)
}

// EditorSetting configures editor-specific settings (can be ignored at runtime)
func (fc control) EditorSetting(value string) control {
	return fc.set("editorSetting", value)
}

// ExtraName sets an additional field name for flattening values in range components
func (fc control) ExtraName(value string) control {
	return fc.set("extraName", value)
}

// Hidden controls the overall visibility of the component
func (fc control) Hidden(value bool) control {
	return fc.set("hidden", value)
}

// HiddenOn sets a conditional expression for hiding the component (e.g., `data.xxx > 5`)
func (fc control) HiddenOn(value string) control {
	return fc.set("hiddenOn", value)
}

// Hint sets an input hint displayed on focus
func (fc control) Hint(value string) control {
	return fc.set("hint", value)
}

// Horizontal configures specific left-right distribution for horizontal layout
func (fc control) Horizontal(value string) control {
	return fc.set("horizontal", value)
}

// ID sets a unique identifier, primarily used for log collection
func (fc control) ID(value string) control {
	return fc.set("id", value)
}

// InitAutoFill configures initial automatic filling behavior
func (fc control) InitAutoFill(value string) control {
	return fc.set("initAutoFill", value)
}

// Inline enables or disables inline form control mode
func (fc control) Inline(value bool) control {
	return fc.set("inline", value)
}

// InputClassName sets the CSS class name for input elements
func (fc control) InputClassName(value string) control {
	return fc.set("inputClassName", value)
}

// Label sets the description title
func (fc control) Label(value string) control {
	return fc.set("label", value)
}

// LabelAlign sets the alignment for the description title (options: right | left | top | inherit)
func (fc control) LabelAlign(value string) control {
	return fc.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label
func (fc control) LabelClassName(value string) control {
	return fc.set("labelClassName", value)
}

// LabelRemark adds a small tooltip icon next to the label
func (fc control) LabelRemark(value string) control {
	return fc.set("labelRemark", value)
}

// LabelWidth sets a custom width for the label (default unit is px)
func (fc control) LabelWidth(value string) control {
	return fc.set("labelWidth", value)
}

// Mode sets the display mode for the form item (options: normal | inline | horizontal)
func (fc control) Mode(value string) control {
	return fc.set("mode", value)
}

// Name sets the field name for form submission, supporting multi-level naming with dot notation (e.g., a.b.c)
func (fc control) Name(value string) control {
	return fc.set("name", value)
}

// OnEvent configures event-driven actions for the component
func (fc control) OnEvent(value any) control {
	return fc.set("onEvent", value)
}

// Placeholder sets the input placeholder text
func (fc control) Placeholder(value string) control {
	return fc.set("placeholder", value)
}

// ReadOnly enables or disables read-only mode for the component
func (fc control) ReadOnly(value bool) control {
	return fc.set("readOnly", value)
}

// ReadOnlyOn sets a conditional expression for read-only state
func (fc control) ReadOnlyOn(value string) control {
	return fc.set("readOnlyOn", value)
}

// Remark adds a small tooltip icon with additional information on hover
func (fc control) Remark(value string) control {
	return fc.set("remark", value)
}

// Required determines whether the field is mandatory
func (fc control) Required(value bool) control {
	return fc.set("required", value)
}

// Row sets the number of rows (for multi-line inputs)
func (fc control) Row(value string) control {
	return fc.set("row", value)
}

// SaveImmediately enables immediate saving (used in TableColumn)
func (fc control) SaveImmediately(value bool) control {
	return fc.set("saveImmediately", value)
}

// Size sets the form item size (options: xs | sm | md | lg | full)
func (fc control) Size(value string) control {
	return fc.set("size", value)
}

// Static enables or disables static display mode
func (fc control) Static(value bool) control {
	return fc.set("static", value)
}

// StaticClassName sets the CSS class name for static form item display (supports string or object configuration with conditional expressions)
func (fc control) StaticClassName(value string) control {
	return fc.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input value display (supports string or object configuration with conditional expressions)
func (fc control) StaticInputClassName(value string) control {
	return fc.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display (supports string or object configuration with conditional expressions)
func (fc control) StaticLabelClassName(value string) control {
	return fc.set("staticLabelClassName", value)
}

// StaticOn sets a conditional expression for static display (e.g., `data.xxx > 5`)
func (fc control) StaticOn(value string) control {
	return fc.set("staticOn", value)
}

// StaticPlaceholder defines a placeholder for empty static values
func (fc control) StaticPlaceholder(value string) control {
	return fc.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (fc control) StaticSchema(value string) control {
	return fc.set("staticSchema", value)
}

// Style sets custom inline styles for the component
func (fc control) Style(value any) control {
	return fc.set("style", value)
}

// SubmitOnChange determines whether to submit the form when the value is modified
func (fc control) SubmitOnChange(value bool) control {
	return fc.set("submitOnChange", value)
}

// TestIdBuilder configures test ID generation
func (fc control) TestIdBuilder(value string) control {
	return fc.set("testIdBuilder", value)
}

// UseMobileUI enables or disables mobile-specific UI styling at the component level
func (fc control) UseMobileUI(value bool) control {
	return fc.set("useMobileUI", value)
}

// ValidateApi sets the remote validation endpoint for the form item
func (fc control) ValidateApi(value string) control {
	return fc.set("validateApi", value)
}

// ValidateOnChange configures validation behavior on modification (by default, validation occurs after form submission)
func (fc control) ValidateOnChange(value bool) control {
	return fc.set("validateOnChange", value)
}

// ValidationErrors sets the error messages for validation failures
func (fc control) ValidationErrors(value string) control {
	return fc.set("validationErrors", value)
}

// Validations sets the validation rules
func (fc control) Validations(value string) control {
	return fc.set("validations", value)
}

// Value sets the default value (static only, not supporting variable retrieval; data association is done via the name attribute)
func (fc control) Value(value string) control {
	return fc.set("value", value)
}

// Visible controls the overall visibility of the component
func (fc control) Visible(value bool) control {
	return fc.set("visible", value)
}

// VisibleOn sets a conditional expression for component visibility (e.g., `data.xxx > 5`)
func (fc control) VisibleOn(value string) control {
	return fc.set("visibleOn", value)
}

// Width adjusts the component width (used in Table)
func (fc control) Width(value string) control {
	return fc.set("width", value)
}
