package comp

import "github.com/zrcoder/amisgo/model"

// ConditionBuilder represents a condition combination control with multiple configuration options
type ConditionBuilder model.Schema

func NewConditionBuilder() ConditionBuilder {
	return ConditionBuilder{"type": "condition-builder"}
}

func (c ConditionBuilder) set(key string, value any) ConditionBuilder {
	c[key] = value
	return c
}

// AddBtnVisibleOn sets the visibility condition for the "Add" button
func (c ConditionBuilder) AddBtnVisibleOn(value string) ConditionBuilder {
	return c.set("addBtnVisibleOn", value)
}

// AddGroupBtnVisibleOn sets the visibility condition for the "Add Condition Group" button
func (c ConditionBuilder) AddGroupBtnVisibleOn(value string) ConditionBuilder {
	return c.set("addGroupBtnVisibleOn", value)
}

// AutoFill configures automatic filling options
func (c ConditionBuilder) AutoFill(value string) ConditionBuilder {
	return c.set("autoFill", value)
}

// BuilderMode sets the display mode
func (c ConditionBuilder) BuilderMode(value string) ConditionBuilder {
	return c.set("builderMode", value)
}

// ClassName sets the CSS class name for the container
func (c ConditionBuilder) ClassName(value string) ConditionBuilder {
	return c.set("className", value)
}

// ClearValueOnHidden determines whether to remove the value when the form item is hidden
func (c ConditionBuilder) ClearValueOnHidden(value bool) ConditionBuilder {
	return c.set("clearValueOnHidden", value)
}

// Config sets additional configuration options
func (c ConditionBuilder) Config(value string) ConditionBuilder {
	return c.set("config", value)
}

// Desc sets a short description
func (c ConditionBuilder) Desc(value string) ConditionBuilder {
	return c.set("desc", value)
}

// Description sets detailed description content, supports HTML fragments
func (c ConditionBuilder) Description(value string) ConditionBuilder {
	return c.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description content
func (c ConditionBuilder) DescriptionClassName(value string) ConditionBuilder {
	return c.set("descriptionClassName", value)
}

// Disabled enables or disables the component
func (c ConditionBuilder) Disabled(value bool) ConditionBuilder {
	return c.set("disabled", value)
}

// DisabledOn sets a conditional expression for disabling the component
func (c ConditionBuilder) DisabledOn(value string) ConditionBuilder {
	return c.set("disabledOn", value)
}

// Draggable enables or disables drag-and-drop functionality
func (c ConditionBuilder) Draggable(value bool) ConditionBuilder {
	return c.set("draggable", value)
}

// EditorSetting configures editor-specific settings
func (c ConditionBuilder) EditorSetting(value string) ConditionBuilder {
	return c.set("editorSetting", value)
}

// Embed enables or disables embedded mode
func (c ConditionBuilder) Embed(value bool) ConditionBuilder {
	return c.set("embed", value)
}

// ExtraName sets an additional field name
func (c ConditionBuilder) ExtraName(value string) ConditionBuilder {
	return c.set("extraName", value)
}

// Fields sets the collection of available fields
func (c ConditionBuilder) Fields(value string) ConditionBuilder {
	return c.set("fields", value)
}

// Formula configures the field input control as a formula editor
func (c ConditionBuilder) Formula(value string) ConditionBuilder {
	return c.set("formula", value)
}

// FormulaForIf configures the formula editor for the 'if' condition
func (c ConditionBuilder) FormulaForIf(value string) ConditionBuilder {
	return c.set("formulaForIf", value)
}

// Funcs sets the collection of available functions
func (c ConditionBuilder) Funcs(value string) ConditionBuilder {
	return c.set("funcs", value)
}

// Hidden controls the overall visibility of the component
func (c ConditionBuilder) Hidden(value bool) ConditionBuilder {
	return c.set("hidden", value)
}

// HiddenOn sets a conditional expression for hiding the component
func (c ConditionBuilder) HiddenOn(value string) ConditionBuilder {
	return c.set("hiddenOn", value)
}

// Hint sets an input hint or placeholder text
func (c ConditionBuilder) Hint(value string) ConditionBuilder {
	return c.set("hint", value)
}

// Horizontal configures horizontal layout distribution
func (c ConditionBuilder) Horizontal(value string) ConditionBuilder {
	return c.set("horizontal", value)
}

// ID sets a unique identifier for the component
func (c ConditionBuilder) ID(value string) ConditionBuilder {
	return c.set("id", value)
}

// InitAutoFill configures initial automatic filling behavior
func (c ConditionBuilder) InitAutoFill(value string) ConditionBuilder {
	return c.set("initAutoFill", value)
}

// Inline enables or disables inline form control mode
func (c ConditionBuilder) Inline(value bool) ConditionBuilder {
	return c.set("inline", value)
}

// InputClassName sets the CSS class name for input elements
func (c ConditionBuilder) InputClassName(value string) ConditionBuilder {
	return c.set("inputClassName", value)
}

// Label sets the description title
func (c ConditionBuilder) Label(value string) ConditionBuilder {
	return c.set("label", value)
}

// LabelAlign sets the alignment for the description title
func (c ConditionBuilder) LabelAlign(value string) ConditionBuilder {
	return c.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the description title
func (c ConditionBuilder) LabelClassName(value string) ConditionBuilder {
	return c.set("labelClassName", value)
}

// LabelRemark adds a small icon tooltip next to the description title
func (c ConditionBuilder) LabelRemark(value string) ConditionBuilder {
	return c.set("labelRemark", value)
}

// LabelWidth sets the width of the description title
func (c ConditionBuilder) LabelWidth(value string) ConditionBuilder {
	return c.set("labelWidth", value)
}

// Mode sets the display mode for the form item
func (c ConditionBuilder) Mode(value string) ConditionBuilder {
	return c.set("mode", value)
}

// Name sets the field name
func (c ConditionBuilder) Name(value string) ConditionBuilder {
	return c.set("name", value)
}

// OnEvent configures event-driven actions for the component
func (c ConditionBuilder) OnEvent(value any) ConditionBuilder {
	return c.set("onEvent", value)
}

// PickerIcon sets the icon for triggering the popup
func (c ConditionBuilder) PickerIcon(value string) ConditionBuilder {
	return c.set("pickerIcon", value)
}

// Placeholder sets the placeholder text
func (c ConditionBuilder) Placeholder(value string) ConditionBuilder {
	return c.set("placeholder", value)
}

// ReadOnly enables or disables read-only mode
func (c ConditionBuilder) ReadOnly(value bool) ConditionBuilder {
	return c.set("readOnly", value)
}

// ReadOnlyOn sets a conditional expression for read-only state
func (c ConditionBuilder) ReadOnlyOn(value string) ConditionBuilder {
	return c.set("readOnlyOn", value)
}

// Remark sets the tooltip content
func (c ConditionBuilder) Remark(value string) ConditionBuilder {
	return c.set("remark", value)
}

// Required determines whether the field is mandatory
func (c ConditionBuilder) Required(value bool) ConditionBuilder {
	return c.set("required", value)
}

// Row sets the number of rows
func (c ConditionBuilder) Row(value string) ConditionBuilder {
	return c.set("row", value)
}

// SaveImmediately enables or disables immediate saving
func (c ConditionBuilder) SaveImmediately(value bool) ConditionBuilder {
	return c.set("saveImmediately", value)
}

// ShowANDOR enables or disables the AND/OR toggle button
func (c ConditionBuilder) ShowANDOR(value bool) ConditionBuilder {
	return c.set("showANDOR", value)
}

// Size sets the form item size
func (c ConditionBuilder) Size(value string) ConditionBuilder {
	return c.set("size", value)
}

// Source sets the remote configuration options
func (c ConditionBuilder) Source(value string) ConditionBuilder {
	return c.set("source", value)
}

// Static enables or disables static display mode
func (c ConditionBuilder) Static(value bool) ConditionBuilder {
	return c.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (c ConditionBuilder) StaticClassName(value string) ConditionBuilder {
	return c.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (c ConditionBuilder) StaticInputClassName(value string) ConditionBuilder {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (c ConditionBuilder) StaticLabelClassName(value string) ConditionBuilder {
	return c.set("staticLabelClassName", value)
}

// StaticOn sets a conditional expression for static display
func (c ConditionBuilder) StaticOn(value string) ConditionBuilder {
	return c.set("staticOn", value)
}

// StaticPlaceholder defines a placeholder for empty static values
func (c ConditionBuilder) StaticPlaceholder(value string) ConditionBuilder {
	return c.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (c ConditionBuilder) StaticSchema(value string) ConditionBuilder {
	return c.set("staticSchema", value)
}

// Style sets custom inline styles for the component
func (c ConditionBuilder) Style(value any) ConditionBuilder {
	return c.set("style", value)
}

// SubmitOnChange determines whether to submit the form on value change
func (c ConditionBuilder) SubmitOnChange(value bool) ConditionBuilder {
	return c.set("submitOnChange", value)
}

// TestIdBuilder configures test ID generation
func (c ConditionBuilder) TestIdBuilder(value string) ConditionBuilder {
	return c.set("testIdBuilder", value)
}

// UseMobileUI enables or disables mobile-specific UI styling
func (c ConditionBuilder) UseMobileUI(value bool) ConditionBuilder {
	return c.set("useMobileUI", value)
}

// ValidateApi sets the remote validation endpoint for the form item
func (c ConditionBuilder) ValidateApi(value string) ConditionBuilder {
	return c.set("validateApi", value)
}

// ValidateOnChange determines whether to trigger validation on each modification
func (c ConditionBuilder) ValidateOnChange(value bool) ConditionBuilder {
	return c.set("validateOnChange", value)
}

// ValidationErrors sets the validation failure error messages
func (c ConditionBuilder) ValidationErrors(value string) ConditionBuilder {
	return c.set("validationErrors", value)
}

// Validations sets the validation rules
func (c ConditionBuilder) Validations(value string) ConditionBuilder {
	return c.set("validations", value)
}

// Value sets the default value
func (c ConditionBuilder) Value(value string) ConditionBuilder {
	return c.set("value", value)
}

// Visible controls the overall visibility of the component
func (c ConditionBuilder) Visible(value bool) ConditionBuilder {
	return c.set("visible", value)
}

// VisibleOn sets a conditional expression for component visibility
func (c ConditionBuilder) VisibleOn(value string) ConditionBuilder {
	return c.set("visibleOn", value)
}

// Width sets the component's width
func (c ConditionBuilder) Width(value string) ConditionBuilder {
	return c.set("width", value)
}
