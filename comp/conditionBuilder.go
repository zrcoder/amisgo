package comp

import "github.com/zrcoder/amisgo/model"

// ConditionBuilder represents a condition combination control with multiple configuration options
type conditionBuilder model.Schema

// ConditionBuilder creates a new ConditionBuilder instance
func ConditionBuilder() conditionBuilder {
	return make(conditionBuilder).set("type", "condition-builder")
}

func (c conditionBuilder) set(key string, value any) conditionBuilder {
	c[key] = value
	return c
}

// AddBtnVisibleOn sets the visibility condition for the "Add" button
func (c conditionBuilder) AddBtnVisibleOn(value string) conditionBuilder {
	return c.set("addBtnVisibleOn", value)
}

// AddGroupBtnVisibleOn sets the visibility condition for the "Add Condition Group" button
func (c conditionBuilder) AddGroupBtnVisibleOn(value string) conditionBuilder {
	return c.set("addGroupBtnVisibleOn", value)
}

// AutoFill configures automatic filling options
func (c conditionBuilder) AutoFill(value string) conditionBuilder {
	return c.set("autoFill", value)
}

// BuilderMode sets the display mode
func (c conditionBuilder) BuilderMode(value string) conditionBuilder {
	return c.set("builderMode", value)
}

// ClassName sets the CSS class name for the container
func (c conditionBuilder) ClassName(value string) conditionBuilder {
	return c.set("className", value)
}

// ClearValueOnHidden determines whether to remove the value when the form item is hidden
func (c conditionBuilder) ClearValueOnHidden(value bool) conditionBuilder {
	return c.set("clearValueOnHidden", value)
}

// Config sets additional configuration options
func (c conditionBuilder) Config(value string) conditionBuilder {
	return c.set("config", value)
}

// Desc sets a short description
func (c conditionBuilder) Desc(value string) conditionBuilder {
	return c.set("desc", value)
}

// Description sets detailed description content, supports HTML fragments
func (c conditionBuilder) Description(value string) conditionBuilder {
	return c.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description content
func (c conditionBuilder) DescriptionClassName(value string) conditionBuilder {
	return c.set("descriptionClassName", value)
}

// Disabled enables or disables the component
func (c conditionBuilder) Disabled(value bool) conditionBuilder {
	return c.set("disabled", value)
}

// DisabledOn sets a conditional expression for disabling the component
func (c conditionBuilder) DisabledOn(value string) conditionBuilder {
	return c.set("disabledOn", value)
}

// Draggable enables or disables drag-and-drop functionality
func (c conditionBuilder) Draggable(value bool) conditionBuilder {
	return c.set("draggable", value)
}

// EditorSetting configures editor-specific settings
func (c conditionBuilder) EditorSetting(value string) conditionBuilder {
	return c.set("editorSetting", value)
}

// Embed enables or disables embedded mode
func (c conditionBuilder) Embed(value bool) conditionBuilder {
	return c.set("embed", value)
}

// ExtraName sets an additional field name
func (c conditionBuilder) ExtraName(value string) conditionBuilder {
	return c.set("extraName", value)
}

// Fields sets the collection of available fields
func (c conditionBuilder) Fields(value string) conditionBuilder {
	return c.set("fields", value)
}

// Formula configures the field input control as a formula editor
func (c conditionBuilder) Formula(value string) conditionBuilder {
	return c.set("formula", value)
}

// FormulaForIf configures the formula editor for the 'if' condition
func (c conditionBuilder) FormulaForIf(value string) conditionBuilder {
	return c.set("formulaForIf", value)
}

// Funcs sets the collection of available functions
func (c conditionBuilder) Funcs(value string) conditionBuilder {
	return c.set("funcs", value)
}

// Hidden controls the overall visibility of the component
func (c conditionBuilder) Hidden(value bool) conditionBuilder {
	return c.set("hidden", value)
}

// HiddenOn sets a conditional expression for hiding the component
func (c conditionBuilder) HiddenOn(value string) conditionBuilder {
	return c.set("hiddenOn", value)
}

// Hint sets an input hint or placeholder text
func (c conditionBuilder) Hint(value string) conditionBuilder {
	return c.set("hint", value)
}

// Horizontal configures horizontal layout distribution
func (c conditionBuilder) Horizontal(value string) conditionBuilder {
	return c.set("horizontal", value)
}

// ID sets a unique identifier for the component
func (c conditionBuilder) ID(value string) conditionBuilder {
	return c.set("id", value)
}

// InitAutoFill configures initial automatic filling behavior
func (c conditionBuilder) InitAutoFill(value string) conditionBuilder {
	return c.set("initAutoFill", value)
}

// Inline enables or disables inline form control mode
func (c conditionBuilder) Inline(value bool) conditionBuilder {
	return c.set("inline", value)
}

// InputClassName sets the CSS class name for input elements
func (c conditionBuilder) InputClassName(value string) conditionBuilder {
	return c.set("inputClassName", value)
}

// Label sets the description title
func (c conditionBuilder) Label(value string) conditionBuilder {
	return c.set("label", value)
}

// LabelAlign sets the alignment for the description title
func (c conditionBuilder) LabelAlign(value string) conditionBuilder {
	return c.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the description title
func (c conditionBuilder) LabelClassName(value string) conditionBuilder {
	return c.set("labelClassName", value)
}

// LabelRemark adds a small icon tooltip next to the description title
func (c conditionBuilder) LabelRemark(value string) conditionBuilder {
	return c.set("labelRemark", value)
}

// LabelWidth sets the width of the description title
func (c conditionBuilder) LabelWidth(value string) conditionBuilder {
	return c.set("labelWidth", value)
}

// Mode sets the display mode for the form item
func (c conditionBuilder) Mode(value string) conditionBuilder {
	return c.set("mode", value)
}

// Name sets the field name
func (c conditionBuilder) Name(value string) conditionBuilder {
	return c.set("name", value)
}

// OnEvent configures event-driven actions for the component
func (c conditionBuilder) OnEvent(value any) conditionBuilder {
	return c.set("onEvent", value)
}

// PickerIcon sets the icon for triggering the popup
func (c conditionBuilder) PickerIcon(value string) conditionBuilder {
	return c.set("pickerIcon", value)
}

// Placeholder sets the placeholder text
func (c conditionBuilder) Placeholder(value string) conditionBuilder {
	return c.set("placeholder", value)
}

// ReadOnly enables or disables read-only mode
func (c conditionBuilder) ReadOnly(value bool) conditionBuilder {
	return c.set("readOnly", value)
}

// ReadOnlyOn sets a conditional expression for read-only state
func (c conditionBuilder) ReadOnlyOn(value string) conditionBuilder {
	return c.set("readOnlyOn", value)
}

// Remark sets the tooltip content
func (c conditionBuilder) Remark(value string) conditionBuilder {
	return c.set("remark", value)
}

// Required determines whether the field is mandatory
func (c conditionBuilder) Required(value bool) conditionBuilder {
	return c.set("required", value)
}

// Row sets the number of rows
func (c conditionBuilder) Row(value string) conditionBuilder {
	return c.set("row", value)
}

// SaveImmediately enables or disables immediate saving
func (c conditionBuilder) SaveImmediately(value bool) conditionBuilder {
	return c.set("saveImmediately", value)
}

// ShowANDOR enables or disables the AND/OR toggle button
func (c conditionBuilder) ShowANDOR(value bool) conditionBuilder {
	return c.set("showANDOR", value)
}

// Size sets the form item size
func (c conditionBuilder) Size(value string) conditionBuilder {
	return c.set("size", value)
}

// Source sets the remote configuration options
func (c conditionBuilder) Source(value string) conditionBuilder {
	return c.set("source", value)
}

// Static enables or disables static display mode
func (c conditionBuilder) Static(value bool) conditionBuilder {
	return c.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (c conditionBuilder) StaticClassName(value string) conditionBuilder {
	return c.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (c conditionBuilder) StaticInputClassName(value string) conditionBuilder {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (c conditionBuilder) StaticLabelClassName(value string) conditionBuilder {
	return c.set("staticLabelClassName", value)
}

// StaticOn sets a conditional expression for static display
func (c conditionBuilder) StaticOn(value string) conditionBuilder {
	return c.set("staticOn", value)
}

// StaticPlaceholder defines a placeholder for empty static values
func (c conditionBuilder) StaticPlaceholder(value string) conditionBuilder {
	return c.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (c conditionBuilder) StaticSchema(value string) conditionBuilder {
	return c.set("staticSchema", value)
}

// Style sets custom inline styles for the component
func (c conditionBuilder) Style(value any) conditionBuilder {
	return c.set("style", value)
}

// SubmitOnChange determines whether to submit the form on value change
func (c conditionBuilder) SubmitOnChange(value bool) conditionBuilder {
	return c.set("submitOnChange", value)
}

// TestIdBuilder configures test ID generation
func (c conditionBuilder) TestIdBuilder(value string) conditionBuilder {
	return c.set("testIdBuilder", value)
}

// UseMobileUI enables or disables mobile-specific UI styling
func (c conditionBuilder) UseMobileUI(value bool) conditionBuilder {
	return c.set("useMobileUI", value)
}

// ValidateApi sets the remote validation endpoint for the form item
func (c conditionBuilder) ValidateApi(value string) conditionBuilder {
	return c.set("validateApi", value)
}

// ValidateOnChange determines whether to trigger validation on each modification
func (c conditionBuilder) ValidateOnChange(value bool) conditionBuilder {
	return c.set("validateOnChange", value)
}

// ValidationErrors sets the validation failure error messages
func (c conditionBuilder) ValidationErrors(value string) conditionBuilder {
	return c.set("validationErrors", value)
}

// Validations sets the validation rules
func (c conditionBuilder) Validations(value string) conditionBuilder {
	return c.set("validations", value)
}

// Value sets the default value
func (c conditionBuilder) Value(value string) conditionBuilder {
	return c.set("value", value)
}

// Visible controls the overall visibility of the component
func (c conditionBuilder) Visible(value bool) conditionBuilder {
	return c.set("visible", value)
}

// VisibleOn sets a conditional expression for component visibility
func (c conditionBuilder) VisibleOn(value string) conditionBuilder {
	return c.set("visibleOn", value)
}

// Width sets the component's width
func (c conditionBuilder) Width(value string) conditionBuilder {
	return c.set("width", value)
}
