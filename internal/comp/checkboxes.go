package comp

import "github.com/zrcoder/amisgo/schema"

type Checkboxes schema.Schema

func NewCheckboxes() Checkboxes {
	return Checkboxes{"type": "checkboxes"}
}

func (c Checkboxes) set(key string, value any) Checkboxes {
	c[key] = value
	return c
}

// AddApi specifies the API endpoint for adding new items
func (c Checkboxes) AddApi(value string) Checkboxes {
	return c.set("addApi", value)
}

// AddControls defines the form controls for adding new items
func (c Checkboxes) AddControls(value string) Checkboxes {
	return c.set("addControls", value)
}

// AddDialog configures the dialog settings for adding new items
func (c Checkboxes) AddDialog(value string) Checkboxes {
	return c.set("addDialog", value)
}

// AutoFill enables automatic synchronization of selected option values to the form
func (c Checkboxes) AutoFill(value string) Checkboxes {
	return c.set("autoFill", value)
}

// CheckAll enables or disables the select all functionality
func (c Checkboxes) CheckAll(value bool) Checkboxes {
	return c.set("checkAll", value)
}

// CheckAllText sets the text for select all/none action
func (c Checkboxes) CheckAllText(value string) Checkboxes {
	return c.set("checkAllText", value)
}

// ClassName sets the CSS class name for the container
func (c Checkboxes) ClassName(value string) Checkboxes {
	return c.set("className", value)
}

// ClearValueOnHidden determines whether to remove the form item's value when hidden
func (c Checkboxes) ClearValueOnHidden(value bool) Checkboxes {
	return c.set("clearValueOnHidden", value)
}

// Clearable enables or disables the ability to clear selections
func (c Checkboxes) Clearable(value bool) Checkboxes {
	return c.set("clearable", value)
}

// ColumnsCount sets the number of columns for displaying checkboxes
func (c Checkboxes) ColumnsCount(value string) Checkboxes {
	return c.set("columnsCount", value)
}

// Creatable enables or disables the ability to add new items
func (c Checkboxes) Creatable(value bool) Checkboxes {
	return c.set("creatable", value)
}

// CreateBtnLabel sets the text for the create new item button
func (c Checkboxes) CreateBtnLabel(value string) Checkboxes {
	return c.set("createBtnLabel", value)
}

// DefaultCheckAll sets whether all options are selected by default
func (c Checkboxes) DefaultCheckAll(value bool) Checkboxes {
	return c.set("defaultCheckAll", value)
}

// DeferApi specifies the API endpoint for lazy loading
func (c Checkboxes) DeferApi(value string) Checkboxes {
	return c.set("deferApi", value)
}

// DeferField sets the field for lazy loading
func (c Checkboxes) DeferField(value string) Checkboxes {
	return c.set("deferField", value)
}

// DeleteApi specifies the API endpoint for deleting options
func (c Checkboxes) DeleteApi(value string) Checkboxes {
	return c.set("deleteApi", value)
}

// DeleteConfirmText sets the confirmation text for option deletion
func (c Checkboxes) DeleteConfirmText(value string) Checkboxes {
	return c.set("deleteConfirmText", value)
}

// Delimiter sets the separator for joining multiple selected values
func (c Checkboxes) Delimiter(value string) Checkboxes {
	return c.set("delimiter", value)
}

// Desc sets a brief description for the checkboxes
func (c Checkboxes) Desc(value string) Checkboxes {
	return c.set("desc", value)
}

// Description sets a detailed description that supports HTML fragments
func (c Checkboxes) Description(value string) Checkboxes {
	return c.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description element
func (c Checkboxes) DescriptionClassName(value string) Checkboxes {
	return c.set("descriptionClassName", value)
}

// Disabled enables or disables the entire checkboxes component
func (c Checkboxes) Disabled(value bool) Checkboxes {
	return c.set("disabled", value)
}

// DisabledOn sets a conditional expression to dynamically disable the component
func (c Checkboxes) DisabledOn(value string) Checkboxes {
	return c.set("disabledOn", value)
}

// EditApi specifies the API endpoint for editing existing items
func (c Checkboxes) EditApi(value string) Checkboxes {
	return c.set("editApi", value)
}

// EditControls defines the form controls for modifying existing items
func (c Checkboxes) EditControls(value string) Checkboxes {
	return c.set("editControls", value)
}

// EditDialog configures the dialog settings for editing existing items
func (c Checkboxes) EditDialog(value string) Checkboxes {
	return c.set("editDialog", value)
}

// Editable enables or disables the ability to edit existing items
func (c Checkboxes) Editable(value bool) Checkboxes {
	return c.set("editable", value)
}

// EditorSetting configures editor-specific settings for the component
func (c Checkboxes) EditorSetting(value string) Checkboxes {
	return c.set("editorSetting", value)
}

// ExtraName sets an additional field name for the component
func (c Checkboxes) ExtraName(value string) Checkboxes {
	return c.set("extraName", value)
}

// ExtractValue enables wrapping selected option values as an array
func (c Checkboxes) ExtractValue(value bool) Checkboxes {
	return c.set("extractValue", value)
}

// Hidden controls the visibility of the entire checkboxes component
func (c Checkboxes) Hidden(value bool) Checkboxes {
	return c.set("hidden", value)
}

// HiddenOn sets a conditional expression to dynamically hide the component
func (c Checkboxes) HiddenOn(value string) Checkboxes {
	return c.set("hiddenOn", value)
}

// Hint provides an input hint or placeholder text
func (c Checkboxes) Hint(value string) Checkboxes {
	return c.set("hint", value)
}

// Horizontal configures the specific left and right distribution when using horizontal layout
func (c Checkboxes) Horizontal(value string) Checkboxes {
	return c.set("horizontal", value)
}

// ID sets a unique identifier for the component
func (c Checkboxes) ID(value string) Checkboxes {
	return c.set("id", value)
}

// InitFetch determines whether to initially fetch data from the source API
func (c Checkboxes) InitFetch(value bool) Checkboxes {
	return c.set("initFetch", value)
}

// InitFetchOn sets a conditional expression to control initial data fetching
func (c Checkboxes) InitFetchOn(value string) Checkboxes {
	return c.set("initFetchOn", value)
}

// Inline determines if the form control should use inline mode
func (c Checkboxes) Inline(value bool) Checkboxes {
	return c.set("inline", value)
}

// InputClassName sets the CSS class name for the input element
func (c Checkboxes) InputClassName(value string) Checkboxes {
	return c.set("inputClassName", value)
}

// JoinValues configures value submission behavior:
// - In single-select mode: submits either the selected value or the entire option object
// - In multi-select mode: joins selected values or submits as an array
func (c Checkboxes) JoinValues(value bool) Checkboxes {
	return c.set("joinValues", value)
}

// Label sets the description label for the checkboxes
func (c Checkboxes) Label(value string) Checkboxes {
	return c.set("label", value)
}

// LabelAlign sets the alignment of the description label (right, left, top, or inherit)
func (c Checkboxes) LabelAlign(value string) Checkboxes {
	return c.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label element
func (c Checkboxes) LabelClassName(value string) Checkboxes {
	return c.set("labelClassName", value)
}

// LabelRemark adds a small icon that displays a tooltip when hovered
func (c Checkboxes) LabelRemark(value string) Checkboxes {
	return c.set("labelRemark", value)
}

// LabelWidth sets a custom width for the label (default unit is px)
func (c Checkboxes) LabelWidth(value string) Checkboxes {
	return c.set("labelWidth", value)
}

// MenuTpl enables custom rendering of checkbox options
func (c Checkboxes) MenuTpl(value string) Checkboxes {
	return c.set("menuTpl", value)
}

// Mode configures the current form item display mode
func (c Checkboxes) Mode(value string) Checkboxes {
	return c.set("mode", value)
}

// Multiple enables or disables multi-select mode
func (c Checkboxes) Multiple(value bool) Checkboxes {
	return c.set("multiple", value)
}

// Name sets the field name for form submission
func (c Checkboxes) Name(value string) Checkboxes {
	return c.set("name", value)
}

// OnEvent configures event-driven actions for the checkboxes
func (c Checkboxes) OnEvent(value any) Checkboxes {
	return c.set("onEvent", value)
}

// Options sets the collection of available checkbox options
func (c Checkboxes) Options(value ...any) Checkboxes {
	return c.set("options", value)
}

// OptionValue sets the field used for option values
func (c Checkboxes) OptionValue(value string) Checkboxes {
	return c.set("optionValue", value)
}

// Radio enables or disables single-select mode
func (c Checkboxes) Radio(value bool) Checkboxes {
	return c.set("radio", value)
}

// Reload determines whether to reload options when refreshing
func (c Checkboxes) Reload(value bool) Checkboxes {
	return c.set("reload", value)
}

// Source sets a custom configuration for the value field name
func (c Checkboxes) Source(value string) Checkboxes {
	return c.set("source", value)
}

// Value sets the default selected value(s)
func (c Checkboxes) Value(value any) Checkboxes {
	return c.set("value", value)
}

// DefaultValue sets the initial value before user interaction
func (c Checkboxes) DefaultValue(value string) Checkboxes {
	return c.set("defaultValue", value)
}
