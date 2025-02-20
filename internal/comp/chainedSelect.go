package comp

import "github.com/zrcoder/amisgo/schema"

// ChainedSelect represents a chained select form component that allows hierarchical or cascading selections
// Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/form/chained-select
type ChainedSelect schema.Schema

func NewChainedSelect() ChainedSelect {
	return ChainedSelect{"type": "chained-select"}
}

func (c ChainedSelect) set(key string, value any) ChainedSelect {
	c[key] = value
	return c
}

// AddApi sets the API endpoint for adding new items to the select component
func (c ChainedSelect) AddApi(value string) ChainedSelect {
	return c.set("addApi", value)
}

// AddControls configures the form fields displayed when adding new items
func (c ChainedSelect) AddControls(value string) ChainedSelect {
	return c.set("addControls", value)
}

// AddDialog configures the settings for the add item dialog modal
func (c ChainedSelect) AddDialog(value string) ChainedSelect {
	return c.set("addDialog", value)
}

// AutoFill enables automatic synchronization of selected option values to the form
func (c ChainedSelect) AutoFill(value string) ChainedSelect {
	return c.set("autoFill", value)
}

// ClassName sets the CSS class name for the container element
func (c ChainedSelect) ClassName(value string) ChainedSelect {
	return c.set("className", value)
}

// ClearValueOnHidden determines whether to remove the form item's value when it becomes hidden
func (c ChainedSelect) ClearValueOnHidden(value bool) ChainedSelect {
	return c.set("clearValueOnHidden", value)
}

// Clearable enables or disables the ability to clear the current selection
func (c ChainedSelect) Clearable(value bool) ChainedSelect {
	return c.set("clearable", value)
}

// Creatable enables or disables the ability to add new options
func (c ChainedSelect) Creatable(value bool) ChainedSelect {
	return c.set("creatable", value)
}

// CreateBtnLabel sets the text for the button used to create new items
func (c ChainedSelect) CreateBtnLabel(value string) ChainedSelect {
	return c.set("createBtnLabel", value)
}

// DeferApi sets the API endpoint for lazy loading of options
func (c ChainedSelect) DeferApi(value string) ChainedSelect {
	return c.set("deferApi", value)
}

// DeferField specifies the field to be used for lazy loading
func (c ChainedSelect) DeferField(value string) ChainedSelect {
	return c.set("deferField", value)
}

// DeleteApi sets the API endpoint for deleting options
func (c ChainedSelect) DeleteApi(value string) ChainedSelect {
	return c.set("deleteApi", value)
}

// DeleteConfirmText sets the confirmation message displayed when deleting an option
func (c ChainedSelect) DeleteConfirmText(value string) ChainedSelect {
	return c.set("deleteConfirmText", value)
}

// Delimiter sets the character used to separate multiple selected values
func (c ChainedSelect) Delimiter(value string) ChainedSelect {
	return c.set("delimiter", value)
}

// Desc sets a brief description for the component
func (c ChainedSelect) Desc(value string) ChainedSelect {
	return c.set("desc", value)
}

// Description sets a detailed description that supports HTML fragments
func (c ChainedSelect) Description(value string) ChainedSelect {
	return c.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description element
func (c ChainedSelect) DescriptionClassName(value string) ChainedSelect {
	return c.set("descriptionClassName", value)
}

// Disabled enables or disables the entire component
func (c ChainedSelect) Disabled(value bool) ChainedSelect {
	return c.set("disabled", value)
}

// DisabledOn sets a conditional expression to dynamically disable the component
func (c ChainedSelect) DisabledOn(value string) ChainedSelect {
	return c.set("disabledOn", value)
}

// EditApi sets the API endpoint for editing existing options
func (c ChainedSelect) EditApi(value string) ChainedSelect {
	return c.set("editApi", value)
}

// EditControls configures the form fields used for modifying existing options
func (c ChainedSelect) EditControls(value string) ChainedSelect {
	return c.set("editControls", value)
}

// EditDialog configures the settings for the edit option dialog modal
func (c ChainedSelect) EditDialog(value string) ChainedSelect {
	return c.set("editDialog", value)
}

// Editable enables or disables the ability to edit existing options
func (c ChainedSelect) Editable(value bool) ChainedSelect {
	return c.set("editable", value)
}

// EditorSetting configures editor-specific settings for the component
func (c ChainedSelect) EditorSetting(value string) ChainedSelect {
	return c.set("editorSetting", value)
}

// ExtraName sets an additional field name for the component
func (c ChainedSelect) ExtraName(value string) ChainedSelect {
	return c.set("extraName", value)
}

// ExtractValue wraps the selected option's value into an array when enabled
func (c ChainedSelect) ExtractValue(value bool) ChainedSelect {
	return c.set("extractValue", value)
}

// Hidden controls the visibility of the component
func (c ChainedSelect) Hidden(value bool) ChainedSelect {
	return c.set("hidden", value)
}

// HiddenOn sets a conditional expression to dynamically hide the component
func (c ChainedSelect) HiddenOn(value string) ChainedSelect {
	return c.set("hiddenOn", value)
}

// Hint provides an input hint or placeholder text
func (c ChainedSelect) Hint(value string) ChainedSelect {
	return c.set("hint", value)
}

// Horizontal configures the specific left and right distribution when using horizontal layout
func (c ChainedSelect) Horizontal(value string) ChainedSelect {
	return c.set("horizontal", value)
}

// ID sets a unique identifier for the component
func (c ChainedSelect) ID(value string) ChainedSelect {
	return c.set("id", value)
}

// InitAutoFill initializes automatic value filling
func (c ChainedSelect) InitAutoFill(value string) ChainedSelect {
	return c.set("initAutoFill", value)
}

// InitFetch configures whether to initially fetch data from the source API
func (c ChainedSelect) InitFetch(value bool) ChainedSelect {
	return c.set("initFetch", value)
}

// InitFetchOn sets a conditional expression to determine initial API data fetching
func (c ChainedSelect) InitFetchOn(value string) ChainedSelect {
	return c.set("initFetchOn", value)
}

// Inline determines if the form control should use inline mode
func (c ChainedSelect) Inline(value bool) ChainedSelect {
	return c.set("inline", value)
}

// InputClassName sets the CSS class name for the input element
func (c ChainedSelect) InputClassName(value string) ChainedSelect {
	return c.set("inputClassName", value)
}

// JoinValues configures value handling for single or multiple selection modes
func (c ChainedSelect) JoinValues(value bool) ChainedSelect {
	return c.set("joinValues", value)
}

// Label sets the description title for the component
func (c ChainedSelect) Label(value string) ChainedSelect {
	return c.set("label", value)
}

// LabelAlign sets the alignment style for the description title
func (c ChainedSelect) LabelAlign(value string) ChainedSelect {
	return c.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label element
func (c ChainedSelect) LabelClassName(value string) ChainedSelect {
	return c.set("labelClassName", value)
}

// LabelRemark adds a small icon that displays a tooltip when hovered
func (c ChainedSelect) LabelRemark(value string) ChainedSelect {
	return c.set("labelRemark", value)
}

// LabelWidth sets a custom width for the label
func (c ChainedSelect) LabelWidth(value string) ChainedSelect {
	return c.set("labelWidth", value)
}

// Mode configures the current form item display mode
func (c ChainedSelect) Mode(value string) ChainedSelect {
	return c.set("mode", value)
}

// Multiple enables or disables multiple selection mode
func (c ChainedSelect) Multiple(value bool) ChainedSelect {
	return c.set("multiple", value)
}

// Name sets the field name for the component
func (c ChainedSelect) Name(value string) ChainedSelect {
	return c.set("name", value)
}

// OnEvent configures event-driven actions for the component
func (c ChainedSelect) OnEvent(value Event) ChainedSelect {
	return c.set("onEvent", value)
}

// Options sets the collection of selectable options
func (c ChainedSelect) Options(value ...any) ChainedSelect {
	return c.set("options", value)
}

// Placeholder sets the input placeholder text
func (c ChainedSelect) Placeholder(value string) ChainedSelect {
	return c.set("placeholder", value)
}

// ReadOnly enables or disables read-only mode
func (c ChainedSelect) ReadOnly(value bool) ChainedSelect {
	return c.set("readOnly", value)
}

// ReadOnlyOn sets a conditional expression for read-only state
func (c ChainedSelect) ReadOnlyOn(value string) ChainedSelect {
	return c.set("readOnlyOn", value)
}

// Remark adds a small icon that displays a tooltip when hovered
func (c ChainedSelect) Remark(value string) ChainedSelect {
	return c.set("remark", value)
}

// Removable enables or disables the ability to remove options
func (c ChainedSelect) Removable(value bool) ChainedSelect {
	return c.set("removable", value)
}

// Required determines if the field is mandatory
func (c ChainedSelect) Required(value bool) ChainedSelect {
	return c.set("required", value)
}

// ResetValue sets the form item's value when the clear button is clicked
func (c ChainedSelect) ResetValue(value string) ChainedSelect {
	return c.set("resetValue", value)
}

// Row sets row-related configuration (context-specific)
func (c ChainedSelect) Row(value string) ChainedSelect {
	return c.set("row", value)
}

// SaveImmediately enables immediate saving (used in TableColumn)
func (c ChainedSelect) SaveImmediately(value bool) ChainedSelect {
	return c.set("saveImmediately", value)
}

// SelectFirst automatically selects the first option by default
func (c ChainedSelect) SelectFirst(value bool) ChainedSelect {
	return c.set("selectFirst", value)
}

// Size sets the form item size (available values: xs | sm | md | lg | full)
func (c ChainedSelect) Size(value string) ChainedSelect {
	return c.set("size", value)
}

// Source sets the API endpoint for fetching options
func (c ChainedSelect) Source(value string) ChainedSelect {
	return c.set("source", value)
}

// Static enables or disables static display mode
func (c ChainedSelect) Static(value bool) ChainedSelect {
	return c.set("static", value)
}

// StaticClassName sets the CSS class name for static form item display
func (c ChainedSelect) StaticClassName(value string) ChainedSelect {
	return c.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input value display
func (c ChainedSelect) StaticInputClassName(value string) ChainedSelect {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (c ChainedSelect) StaticLabelClassName(value string) ChainedSelect {
	return c.set("staticLabelClassName", value)
}

// StaticOn sets a conditional expression for static display
func (c ChainedSelect) StaticOn(value string) ChainedSelect {
	return c.set("staticOn", value)
}

// StaticPlaceholder defines a placeholder for empty static values
func (c ChainedSelect) StaticPlaceholder(value string) ChainedSelect {
	return c.set("staticPlaceholder", value)
}

// StaticSchema sets the schema.Schema for static display
func (c ChainedSelect) StaticSchema(value string) ChainedSelect {
	return c.set("staticSchema", value)
}

// Style sets custom inline styles for the component
func (c ChainedSelect) Style(value any) ChainedSelect {
	return c.set("style", value)
}

// SubmitOnChange determines if the form should be submitted after value modification
func (c ChainedSelect) SubmitOnChange(value bool) ChainedSelect {
	return c.set("submitOnChange", value)
}

// TestIdBuilder configures test ID generation
func (c ChainedSelect) TestIdBuilder(value string) ChainedSelect {
	return c.set("testIdBuilder", value)
}

// UseMobileUI enables or disables mobile-specific UI styling
func (c ChainedSelect) UseMobileUI(value bool) ChainedSelect {
	return c.set("useMobileUI", value)
}

// ValidateApi sets the remote validation endpoint for the form item
func (c ChainedSelect) ValidateApi(value string) ChainedSelect {
	return c.set("validateApi", value)
}

// ValidateOnChange configures validation behavior after form submission
func (c ChainedSelect) ValidateOnChange(value bool) ChainedSelect {
	return c.set("validateOnChange", value)
}

// ValidationErrors sets the error messages for validation failures
func (c ChainedSelect) ValidationErrors(value string) ChainedSelect {
	return c.set("validationErrors", value)
}

// Validations sets validation rules
func (c ChainedSelect) Validations(value string) ChainedSelect {
	return c.set("validations", value)
}

// Value sets the default value (must be a static value, variables not supported)
func (c ChainedSelect) Value(value string) ChainedSelect {
	return c.set("value", value)
}

// ValuesNoWrap prevents line wrapping for multiple selected values
func (c ChainedSelect) ValuesNoWrap(value bool) ChainedSelect {
	return c.set("valuesNoWrap", value)
}

// Visible controls the overall visibility of the component
func (c ChainedSelect) Visible(value bool) ChainedSelect {
	return c.set("visible", value)
}

// VisibleOn sets a conditional expression for component visibility
func (c ChainedSelect) VisibleOn(value string) ChainedSelect {
	return c.set("visibleOn", value)
}

// Width adjusts the component width (typically used in Table contexts)
func (c ChainedSelect) Width(value string) ChainedSelect {
	return c.set("width", value)
}
