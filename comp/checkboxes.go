package comp

import "github.com/zrcoder/amisgo/model"

type checkboxes model.Schema

func Checkboxes() checkboxes {
	return make(checkboxes).set("type", "checkboxes")
}

func (c checkboxes) set(key string, value any) checkboxes {
	c[key] = value
	return c
}

// AddApi specifies the API endpoint for adding new items
func (c checkboxes) AddApi(value string) checkboxes {
	return c.set("addApi", value)
}

// AddControls defines the form controls for adding new items
func (c checkboxes) AddControls(value string) checkboxes {
	return c.set("addControls", value)
}

// AddDialog configures the dialog settings for adding new items
func (c checkboxes) AddDialog(value string) checkboxes {
	return c.set("addDialog", value)
}

// AutoFill enables automatic synchronization of selected option values to the form
func (c checkboxes) AutoFill(value string) checkboxes {
	return c.set("autoFill", value)
}

// CheckAll enables or disables the select all functionality
func (c checkboxes) CheckAll(value bool) checkboxes {
	return c.set("checkAll", value)
}

// CheckAllText sets the text for select all/none action
func (c checkboxes) CheckAllText(value string) checkboxes {
	return c.set("checkAllText", value)
}

// ClassName sets the CSS class name for the container
func (c checkboxes) ClassName(value string) checkboxes {
	return c.set("className", value)
}

// ClearValueOnHidden determines whether to remove the form item's value when hidden
func (c checkboxes) ClearValueOnHidden(value bool) checkboxes {
	return c.set("clearValueOnHidden", value)
}

// Clearable enables or disables the ability to clear selections
func (c checkboxes) Clearable(value bool) checkboxes {
	return c.set("clearable", value)
}

// ColumnsCount sets the number of columns for displaying checkboxes
func (c checkboxes) ColumnsCount(value string) checkboxes {
	return c.set("columnsCount", value)
}

// Creatable enables or disables the ability to add new items
func (c checkboxes) Creatable(value bool) checkboxes {
	return c.set("creatable", value)
}

// CreateBtnLabel sets the text for the create new item button
func (c checkboxes) CreateBtnLabel(value string) checkboxes {
	return c.set("createBtnLabel", value)
}

// DefaultCheckAll sets whether all options are selected by default
func (c checkboxes) DefaultCheckAll(value bool) checkboxes {
	return c.set("defaultCheckAll", value)
}

// DeferApi specifies the API endpoint for lazy loading
func (c checkboxes) DeferApi(value string) checkboxes {
	return c.set("deferApi", value)
}

// DeferField sets the field for lazy loading
func (c checkboxes) DeferField(value string) checkboxes {
	return c.set("deferField", value)
}

// DeleteApi specifies the API endpoint for deleting options
func (c checkboxes) DeleteApi(value string) checkboxes {
	return c.set("deleteApi", value)
}

// DeleteConfirmText sets the confirmation text for option deletion
func (c checkboxes) DeleteConfirmText(value string) checkboxes {
	return c.set("deleteConfirmText", value)
}

// Delimiter sets the separator for joining multiple selected values
func (c checkboxes) Delimiter(value string) checkboxes {
	return c.set("delimiter", value)
}

// Desc sets a brief description for the checkboxes
func (c checkboxes) Desc(value string) checkboxes {
	return c.set("desc", value)
}

// Description sets a detailed description that supports HTML fragments
func (c checkboxes) Description(value string) checkboxes {
	return c.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description element
func (c checkboxes) DescriptionClassName(value string) checkboxes {
	return c.set("descriptionClassName", value)
}

// Disabled enables or disables the entire checkboxes component
func (c checkboxes) Disabled(value bool) checkboxes {
	return c.set("disabled", value)
}

// DisabledOn sets a conditional expression to dynamically disable the component
func (c checkboxes) DisabledOn(value string) checkboxes {
	return c.set("disabledOn", value)
}

// EditApi specifies the API endpoint for editing existing items
func (c checkboxes) EditApi(value string) checkboxes {
	return c.set("editApi", value)
}

// EditControls defines the form controls for modifying existing items
func (c checkboxes) EditControls(value string) checkboxes {
	return c.set("editControls", value)
}

// EditDialog configures the dialog settings for editing existing items
func (c checkboxes) EditDialog(value string) checkboxes {
	return c.set("editDialog", value)
}

// Editable enables or disables the ability to edit existing items
func (c checkboxes) Editable(value bool) checkboxes {
	return c.set("editable", value)
}

// EditorSetting configures editor-specific settings for the component
func (c checkboxes) EditorSetting(value string) checkboxes {
	return c.set("editorSetting", value)
}

// ExtraName sets an additional field name for the component
func (c checkboxes) ExtraName(value string) checkboxes {
	return c.set("extraName", value)
}

// ExtractValue enables wrapping selected option values as an array
func (c checkboxes) ExtractValue(value bool) checkboxes {
	return c.set("extractValue", value)
}

// Hidden controls the visibility of the entire checkboxes component
func (c checkboxes) Hidden(value bool) checkboxes {
	return c.set("hidden", value)
}

// HiddenOn sets a conditional expression to dynamically hide the component
func (c checkboxes) HiddenOn(value string) checkboxes {
	return c.set("hiddenOn", value)
}

// Hint provides an input hint or placeholder text
func (c checkboxes) Hint(value string) checkboxes {
	return c.set("hint", value)
}

// Horizontal configures the specific left and right distribution when using horizontal layout
func (c checkboxes) Horizontal(value string) checkboxes {
	return c.set("horizontal", value)
}

// ID sets a unique identifier for the component
func (c checkboxes) ID(value string) checkboxes {
	return c.set("id", value)
}

// InitFetch determines whether to initially fetch data from the source API
func (c checkboxes) InitFetch(value bool) checkboxes {
	return c.set("initFetch", value)
}

// InitFetchOn sets a conditional expression to control initial data fetching
func (c checkboxes) InitFetchOn(value string) checkboxes {
	return c.set("initFetchOn", value)
}

// Inline determines if the form control should use inline mode
func (c checkboxes) Inline(value bool) checkboxes {
	return c.set("inline", value)
}

// InputClassName sets the CSS class name for the input element
func (c checkboxes) InputClassName(value string) checkboxes {
	return c.set("inputClassName", value)
}

// JoinValues configures value submission behavior:
// - In single-select mode: submits either the selected value or the entire option object
// - In multi-select mode: joins selected values or submits as an array
func (c checkboxes) JoinValues(value bool) checkboxes {
	return c.set("joinValues", value)
}

// Label sets the description label for the checkboxes
func (c checkboxes) Label(value string) checkboxes {
	return c.set("label", value)
}

// LabelAlign sets the alignment of the description label (right, left, top, or inherit)
func (c checkboxes) LabelAlign(value string) checkboxes {
	return c.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label element
func (c checkboxes) LabelClassName(value string) checkboxes {
	return c.set("labelClassName", value)
}

// LabelRemark adds a small icon that displays a tooltip when hovered
func (c checkboxes) LabelRemark(value string) checkboxes {
	return c.set("labelRemark", value)
}

// LabelWidth sets a custom width for the label (default unit is px)
func (c checkboxes) LabelWidth(value string) checkboxes {
	return c.set("labelWidth", value)
}

// MenuTpl enables custom rendering of checkbox options
func (c checkboxes) MenuTpl(value string) checkboxes {
	return c.set("menuTpl", value)
}

// Mode configures the current form item display mode
func (c checkboxes) Mode(value string) checkboxes {
	return c.set("mode", value)
}

// Multiple enables or disables multi-select mode
func (c checkboxes) Multiple(value bool) checkboxes {
	return c.set("multiple", value)
}

// Name sets the field name for form submission
func (c checkboxes) Name(value string) checkboxes {
	return c.set("name", value)
}

// OnEvent configures event-driven actions for the checkboxes
func (c checkboxes) OnEvent(value any) checkboxes {
	return c.set("onEvent", value)
}

// Options sets the collection of available checkbox options
func (c checkboxes) Options(value ...any) checkboxes {
	return c.set("options", value)
}

// OptionValue sets the field used for option values
func (c checkboxes) OptionValue(value string) checkboxes {
	return c.set("optionValue", value)
}

// Radio enables or disables single-select mode
func (c checkboxes) Radio(value bool) checkboxes {
	return c.set("radio", value)
}

// Reload determines whether to reload options when refreshing
func (c checkboxes) Reload(value bool) checkboxes {
	return c.set("reload", value)
}

// Source sets a custom configuration for the value field name
func (c checkboxes) Source(value string) checkboxes {
	return c.set("source", value)
}

// Value sets the default selected value(s)
func (c checkboxes) Value(value any) checkboxes {
	return c.set("value", value)
}

// DefaultValue sets the initial value before user interaction
func (c checkboxes) DefaultValue(value string) checkboxes {
	return c.set("defaultValue", value)
}
