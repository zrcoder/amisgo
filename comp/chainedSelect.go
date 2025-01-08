package comp

import "github.com/zrcoder/amisgo/model"

// ChainedSelect represents a chained select form component that allows hierarchical or cascading selections
// Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/form/chained-select
type chainedSelect model.Schema

// ChainedSelect creates a new ChainedSelect instance
func ChainedSelect() chainedSelect {
	return make(chainedSelect).set("type", "chained-select")
}

func (c chainedSelect) set(key string, value any) chainedSelect {
	c[key] = value
	return c
}

// AddApi sets the API endpoint for adding new items to the select component
func (c chainedSelect) AddApi(value string) chainedSelect {
	return c.set("addApi", value)
}

// AddControls configures the form fields displayed when adding new items
func (c chainedSelect) AddControls(value string) chainedSelect {
	return c.set("addControls", value)
}

// AddDialog configures the settings for the add item dialog modal
func (c chainedSelect) AddDialog(value string) chainedSelect {
	return c.set("addDialog", value)
}

// AutoFill enables automatic synchronization of selected option values to the form
func (c chainedSelect) AutoFill(value string) chainedSelect {
	return c.set("autoFill", value)
}

// ClassName sets the CSS class name for the container element
func (c chainedSelect) ClassName(value string) chainedSelect {
	return c.set("className", value)
}

// ClearValueOnHidden determines whether to remove the form item's value when it becomes hidden
func (c chainedSelect) ClearValueOnHidden(value bool) chainedSelect {
	return c.set("clearValueOnHidden", value)
}

// Clearable enables or disables the ability to clear the current selection
func (c chainedSelect) Clearable(value bool) chainedSelect {
	return c.set("clearable", value)
}

// Creatable enables or disables the ability to add new options
func (c chainedSelect) Creatable(value bool) chainedSelect {
	return c.set("creatable", value)
}

// CreateBtnLabel sets the text for the button used to create new items
func (c chainedSelect) CreateBtnLabel(value string) chainedSelect {
	return c.set("createBtnLabel", value)
}

// DeferApi sets the API endpoint for lazy loading of options
func (c chainedSelect) DeferApi(value string) chainedSelect {
	return c.set("deferApi", value)
}

// DeferField specifies the field to be used for lazy loading
func (c chainedSelect) DeferField(value string) chainedSelect {
	return c.set("deferField", value)
}

// DeleteApi sets the API endpoint for deleting options
func (c chainedSelect) DeleteApi(value string) chainedSelect {
	return c.set("deleteApi", value)
}

// DeleteConfirmText sets the confirmation message displayed when deleting an option
func (c chainedSelect) DeleteConfirmText(value string) chainedSelect {
	return c.set("deleteConfirmText", value)
}

// Delimiter sets the character used to separate multiple selected values
func (c chainedSelect) Delimiter(value string) chainedSelect {
	return c.set("delimiter", value)
}

// Desc sets a brief description for the component
func (c chainedSelect) Desc(value string) chainedSelect {
	return c.set("desc", value)
}

// Description sets a detailed description that supports HTML fragments
func (c chainedSelect) Description(value string) chainedSelect {
	return c.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description element
func (c chainedSelect) DescriptionClassName(value string) chainedSelect {
	return c.set("descriptionClassName", value)
}

// Disabled enables or disables the entire component
func (c chainedSelect) Disabled(value bool) chainedSelect {
	return c.set("disabled", value)
}

// DisabledOn sets a conditional expression to dynamically disable the component
func (c chainedSelect) DisabledOn(value string) chainedSelect {
	return c.set("disabledOn", value)
}

// EditApi sets the API endpoint for editing existing options
func (c chainedSelect) EditApi(value string) chainedSelect {
	return c.set("editApi", value)
}

// EditControls configures the form fields used for modifying existing options
func (c chainedSelect) EditControls(value string) chainedSelect {
	return c.set("editControls", value)
}

// EditDialog configures the settings for the edit option dialog modal
func (c chainedSelect) EditDialog(value string) chainedSelect {
	return c.set("editDialog", value)
}

// Editable enables or disables the ability to edit existing options
func (c chainedSelect) Editable(value bool) chainedSelect {
	return c.set("editable", value)
}

// EditorSetting configures editor-specific settings for the component
func (c chainedSelect) EditorSetting(value string) chainedSelect {
	return c.set("editorSetting", value)
}

// ExtraName sets an additional field name for the component
func (c chainedSelect) ExtraName(value string) chainedSelect {
	return c.set("extraName", value)
}

// ExtractValue wraps the selected option's value into an array when enabled
func (c chainedSelect) ExtractValue(value bool) chainedSelect {
	return c.set("extractValue", value)
}

// Hidden controls the visibility of the component
func (c chainedSelect) Hidden(value bool) chainedSelect {
	return c.set("hidden", value)
}

// HiddenOn sets a conditional expression to dynamically hide the component
func (c chainedSelect) HiddenOn(value string) chainedSelect {
	return c.set("hiddenOn", value)
}

// Hint provides an input hint or placeholder text
func (c chainedSelect) Hint(value string) chainedSelect {
	return c.set("hint", value)
}

// Horizontal configures the specific left and right distribution when using horizontal layout
func (c chainedSelect) Horizontal(value string) chainedSelect {
	return c.set("horizontal", value)
}

// ID sets a unique identifier for the component
func (c chainedSelect) ID(value string) chainedSelect {
	return c.set("id", value)
}

// InitAutoFill initializes automatic value filling
func (c chainedSelect) InitAutoFill(value string) chainedSelect {
	return c.set("initAutoFill", value)
}

// InitFetch configures whether to initially fetch data from the source API
func (c chainedSelect) InitFetch(value bool) chainedSelect {
	return c.set("initFetch", value)
}

// InitFetchOn sets a conditional expression to determine initial API data fetching
func (c chainedSelect) InitFetchOn(value string) chainedSelect {
	return c.set("initFetchOn", value)
}

// Inline determines if the form control should use inline mode
func (c chainedSelect) Inline(value bool) chainedSelect {
	return c.set("inline", value)
}

// InputClassName sets the CSS class name for the input element
func (c chainedSelect) InputClassName(value string) chainedSelect {
	return c.set("inputClassName", value)
}

// JoinValues configures value handling for single or multiple selection modes
func (c chainedSelect) JoinValues(value bool) chainedSelect {
	return c.set("joinValues", value)
}

// Label sets the description title for the component
func (c chainedSelect) Label(value string) chainedSelect {
	return c.set("label", value)
}

// LabelAlign sets the alignment style for the description title
func (c chainedSelect) LabelAlign(value string) chainedSelect {
	return c.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label element
func (c chainedSelect) LabelClassName(value string) chainedSelect {
	return c.set("labelClassName", value)
}

// LabelRemark adds a small icon that displays a tooltip when hovered
func (c chainedSelect) LabelRemark(value string) chainedSelect {
	return c.set("labelRemark", value)
}

// LabelWidth sets a custom width for the label
func (c chainedSelect) LabelWidth(value string) chainedSelect {
	return c.set("labelWidth", value)
}

// Mode configures the current form item display mode
func (c chainedSelect) Mode(value string) chainedSelect {
	return c.set("mode", value)
}

// Multiple enables or disables multiple selection mode
func (c chainedSelect) Multiple(value bool) chainedSelect {
	return c.set("multiple", value)
}

// Name sets the field name for the component
func (c chainedSelect) Name(value string) chainedSelect {
	return c.set("name", value)
}

// OnEvent configures event-driven actions for the component
func (c chainedSelect) OnEvent(value any) chainedSelect {
	return c.set("onEvent", value)
}

// Options sets the collection of selectable options
func (c chainedSelect) Options(value ...any) chainedSelect {
	return c.set("options", value)
}

// Placeholder sets the input placeholder text
func (c chainedSelect) Placeholder(value string) chainedSelect {
	return c.set("placeholder", value)
}

// ReadOnly enables or disables read-only mode
func (c chainedSelect) ReadOnly(value bool) chainedSelect {
	return c.set("readOnly", value)
}

// ReadOnlyOn sets a conditional expression for read-only state
func (c chainedSelect) ReadOnlyOn(value string) chainedSelect {
	return c.set("readOnlyOn", value)
}

// Remark adds a small icon that displays a tooltip when hovered
func (c chainedSelect) Remark(value string) chainedSelect {
	return c.set("remark", value)
}

// Removable enables or disables the ability to remove options
func (c chainedSelect) Removable(value bool) chainedSelect {
	return c.set("removable", value)
}

// Required determines if the field is mandatory
func (c chainedSelect) Required(value bool) chainedSelect {
	return c.set("required", value)
}

// ResetValue sets the form item's value when the clear button is clicked
func (c chainedSelect) ResetValue(value string) chainedSelect {
	return c.set("resetValue", value)
}

// Row sets row-related configuration (context-specific)
func (c chainedSelect) Row(value string) chainedSelect {
	return c.set("row", value)
}

// SaveImmediately enables immediate saving (used in TableColumn)
func (c chainedSelect) SaveImmediately(value bool) chainedSelect {
	return c.set("saveImmediately", value)
}

// SelectFirst automatically selects the first option by default
func (c chainedSelect) SelectFirst(value bool) chainedSelect {
	return c.set("selectFirst", value)
}

// Size sets the form item size (available values: xs | sm | md | lg | full)
func (c chainedSelect) Size(value string) chainedSelect {
	return c.set("size", value)
}

// Source sets the API endpoint for fetching options
func (c chainedSelect) Source(value string) chainedSelect {
	return c.set("source", value)
}

// Static enables or disables static display mode
func (c chainedSelect) Static(value bool) chainedSelect {
	return c.set("static", value)
}

// StaticClassName sets the CSS class name for static form item display
func (c chainedSelect) StaticClassName(value string) chainedSelect {
	return c.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input value display
func (c chainedSelect) StaticInputClassName(value string) chainedSelect {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (c chainedSelect) StaticLabelClassName(value string) chainedSelect {
	return c.set("staticLabelClassName", value)
}

// StaticOn sets a conditional expression for static display
func (c chainedSelect) StaticOn(value string) chainedSelect {
	return c.set("staticOn", value)
}

// StaticPlaceholder defines a placeholder for empty static values
func (c chainedSelect) StaticPlaceholder(value string) chainedSelect {
	return c.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (c chainedSelect) StaticSchema(value string) chainedSelect {
	return c.set("staticSchema", value)
}

// Style sets custom inline styles for the component
func (c chainedSelect) Style(value any) chainedSelect {
	return c.set("style", value)
}

// SubmitOnChange determines if the form should be submitted after value modification
func (c chainedSelect) SubmitOnChange(value bool) chainedSelect {
	return c.set("submitOnChange", value)
}

// TestIdBuilder configures test ID generation
func (c chainedSelect) TestIdBuilder(value string) chainedSelect {
	return c.set("testIdBuilder", value)
}

// UseMobileUI enables or disables mobile-specific UI styling
func (c chainedSelect) UseMobileUI(value bool) chainedSelect {
	return c.set("useMobileUI", value)
}

// ValidateApi sets the remote validation endpoint for the form item
func (c chainedSelect) ValidateApi(value string) chainedSelect {
	return c.set("validateApi", value)
}

// ValidateOnChange configures validation behavior after form submission
func (c chainedSelect) ValidateOnChange(value bool) chainedSelect {
	return c.set("validateOnChange", value)
}

// ValidationErrors sets the error messages for validation failures
func (c chainedSelect) ValidationErrors(value string) chainedSelect {
	return c.set("validationErrors", value)
}

// Validations sets validation rules
func (c chainedSelect) Validations(value string) chainedSelect {
	return c.set("validations", value)
}

// Value sets the default value (must be a static value, variables not supported)
func (c chainedSelect) Value(value string) chainedSelect {
	return c.set("value", value)
}

// ValuesNoWrap prevents line wrapping for multiple selected values
func (c chainedSelect) ValuesNoWrap(value bool) chainedSelect {
	return c.set("valuesNoWrap", value)
}

// Visible controls the overall visibility of the component
func (c chainedSelect) Visible(value bool) chainedSelect {
	return c.set("visible", value)
}

// VisibleOn sets a conditional expression for component visibility
func (c chainedSelect) VisibleOn(value string) chainedSelect {
	return c.set("visibleOn", value)
}

// Width adjusts the component width (typically used in Table contexts)
func (c chainedSelect) Width(value string) chainedSelect {
	return c.set("width", value)
}
