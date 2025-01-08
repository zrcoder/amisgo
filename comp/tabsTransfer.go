package comp

import "github.com/zrcoder/amisgo/model"

// TabsTransfer documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/form/tabs-transfer

// tabsTransfer represents the schema for tabs transfer component
type tabsTransfer model.Schema

// TabsTransfer initializes a tabsTransfer with type "tabs-transfer"
func TabsTransfer() tabsTransfer {
	return tabsTransfer{}.set("type", "tabs-transfer")
}

// set sets a key-value pair in tabsTransfer
func (t tabsTransfer) set(key string, value any) tabsTransfer {
	t[key] = value
	return t
}

// AddApi sets the API for adding items
func (t tabsTransfer) AddApi(value string) tabsTransfer {
	return t.set("addApi", value)
}

// Adds sets the form items for adding new entries
func (t tabsTransfer) Adds(value string) tabsTransfer {
	return t.set("adds", value)
}

// AddDialog sets the configuration for the add dialog
func (t tabsTransfer) AddDialog(value string) tabsTransfer {
	return t.set("addDialog", value)
}

// AutoCheckChildren sets whether to auto-check children items
func (t tabsTransfer) AutoCheckChildren(value bool) tabsTransfer {
	return t.set("autoCheckChildren", value)
}

// AutoFill sets the auto-fill configuration
func (t tabsTransfer) AutoFill(value string) tabsTransfer {
	return t.set("autoFill", value)
}

// ClassName sets the CSS class name for the container
func (t tabsTransfer) ClassName(value string) tabsTransfer {
	return t.set("className", value)
}

// ClearValueOnHidden sets whether to clear the value when hidden
func (t tabsTransfer) ClearValueOnHidden(value bool) tabsTransfer {
	return t.set("clearValueOnHidden", value)
}

// Clearable sets whether the field is clearable
func (t tabsTransfer) Clearable(value bool) tabsTransfer {
	return t.set("clearable", value)
}

// Columns sets the columns for table mode
func (t tabsTransfer) Columns(value ...any) tabsTransfer {
	return t.set("columns", value)
}

// Creatable sets whether new entries can be created
func (t tabsTransfer) Creatable(value bool) tabsTransfer {
	return t.set("creatable", value)
}

// CreateBtnLabel sets the label for the create button
func (t tabsTransfer) CreateBtnLabel(value string) tabsTransfer {
	return t.set("createBtnLabel", value)
}

// DeferApi sets the API for deferred loading
func (t tabsTransfer) DeferApi(value string) tabsTransfer {
	return t.set("deferApi", value)
}

// DeferField sets the field for deferred loading
func (t tabsTransfer) DeferField(value string) tabsTransfer {
	return t.set("deferField", value)
}

// DeleteApi sets the API for deleting items
func (t tabsTransfer) DeleteApi(value string) tabsTransfer {
	return t.set("deleteApi", value)
}

// DeleteConfirmText sets the confirmation text for deletion
func (t tabsTransfer) DeleteConfirmText(value string) tabsTransfer {
	return t.set("deleteConfirmText", value)
}

// Delimiter sets the delimiter for values
func (t tabsTransfer) Delimiter(value string) tabsTransfer {
	return t.set("delimiter", value)
}

// Desc sets the description
func (t tabsTransfer) Desc(value string) tabsTransfer {
	return t.set("desc", value)
}

// Description sets the HTML description content
func (t tabsTransfer) Description(value string) tabsTransfer {
	return t.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description
func (t tabsTransfer) DescriptionClassName(value string) tabsTransfer {
	return t.set("descriptionClassName", value)
}

// Disabled sets whether the field is disabled
func (t tabsTransfer) Disabled(value bool) tabsTransfer {
	return t.set("disabled", value)
}

// DisabledOn sets the expression to determine if the field is disabled
func (t tabsTransfer) DisabledOn(value string) tabsTransfer {
	return t.set("disabledOn", value)
}

// EditApi sets the API for editing items
func (t tabsTransfer) EditApi(value string) tabsTransfer {
	return t.set("editApi", value)
}

// Edits sets the form items for editing entries
func (t tabsTransfer) Edits(value string) tabsTransfer {
	return t.set("edits", value)
}

// EditDialog sets the configuration for the edit dialog
func (t tabsTransfer) EditDialog(value string) tabsTransfer {
	return t.set("editDialog", value)
}

// Editable sets whether the field is editable
func (t tabsTransfer) Editable(value bool) tabsTransfer {
	return t.set("editable", value)
}

// EditorSetting sets the editor configuration
func (t tabsTransfer) EditorSetting(value string) tabsTransfer {
	return t.set("editorSetting", value)
}

// ExtraName sets the extra field name for range components
func (t tabsTransfer) ExtraName(value string) tabsTransfer {
	return t.set("extraName", value)
}

// ExtractValue sets whether to extract the value as an array
func (t tabsTransfer) ExtractValue(value bool) tabsTransfer {
	return t.set("extractValue", value)
}

// Hidden sets whether the field is hidden
func (t tabsTransfer) Hidden(value bool) tabsTransfer {
	return t.set("hidden", value)
}

// HiddenOn sets the expression to determine if the field is hidden
func (t tabsTransfer) HiddenOn(value string) tabsTransfer {
	return t.set("hiddenOn", value)
}

// Hint sets the input hint
func (t tabsTransfer) Hint(value string) tabsTransfer {
	return t.set("hint", value)
}

// Horizontal sets the horizontal layout configuration
func (t tabsTransfer) Horizontal(value string) tabsTransfer {
	return t.set("horizontal", value)
}

// ID sets the unique ID for the component
func (t tabsTransfer) ID(value string) tabsTransfer {
	return t.set("id", value)
}

// InitAutoFill sets the initial auto-fill configuration
func (t tabsTransfer) InitAutoFill(value string) tabsTransfer {
	return t.set("initAutoFill", value)
}

// InitFetch sets whether to fetch data initially
func (t tabsTransfer) InitFetch(value bool) tabsTransfer {
	return t.set("initFetch", value)
}

// InitFetchOn sets the expression to determine if data should be fetched initially
func (t tabsTransfer) InitFetchOn(value string) tabsTransfer {
	return t.set("initFetchOn", value)
}

// InitiallyOpen sets whether the tabs are initially open
func (t tabsTransfer) InitiallyOpen(value bool) tabsTransfer {
	return t.set("initiallyOpen", value)
}

// Inline sets whether the form is in inline mode
func (t tabsTransfer) Inline(value bool) tabsTransfer {
	return t.set("inline", value)
}

// InputClassName sets the CSS class name for the input
func (t tabsTransfer) InputClassName(value string) tabsTransfer {
	return t.set("inputClassName", value)
}

// ItemHeight sets the height of each item for virtual rendering
func (t tabsTransfer) ItemHeight(value string) tabsTransfer {
	return t.set("itemHeight", value)
}

// JoinValues sets whether to join values with a delimiter
func (t tabsTransfer) JoinValues(value bool) tabsTransfer {
	return t.set("joinValues", value)
}

// Label sets the label text
func (t tabsTransfer) Label(value string) tabsTransfer {
	return t.set("label", value)
}

// LabelAlign sets the label alignment
func (t tabsTransfer) LabelAlign(value string) tabsTransfer {
	return t.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label
func (t tabsTransfer) LabelClassName(value string) tabsTransfer {
	return t.set("labelClassName", value)
}

// LabelRemark sets the remark for the label
func (t tabsTransfer) LabelRemark(value string) tabsTransfer {
	return t.set("labelRemark", value)
}

// LabelWidth sets the width of the label
func (t tabsTransfer) LabelWidth(value string) tabsTransfer {
	return t.set("labelWidth", value)
}

// Level sets the level
func (t tabsTransfer) Level(value string) tabsTransfer {
	return t.set("level", value)
}

// Max sets the maximum value
func (t tabsTransfer) Max(value int) tabsTransfer {
	return t.set("max", value)
}

// Min sets the minimum value
func (t tabsTransfer) Min(value int) tabsTransfer {
	return t.set("min", value)
}

// Name sets the unique name for the component
func (t tabsTransfer) Name(value string) tabsTransfer {
	return t.set("name", value)
}

// Nested sets whether nested options are supported
func (t tabsTransfer) Nested(value bool) tabsTransfer {
	return t.set("nested", value)
}

// Options sets the options
func (t tabsTransfer) Options(value ...any) tabsTransfer {
	return t.set("options", value)
}

// OptionTextField sets the text field for options
func (t tabsTransfer) OptionTextField(value string) tabsTransfer {
	return t.set("optionTextField", value)
}

// OptionValueField sets the value field for options
func (t tabsTransfer) OptionValueField(value string) tabsTransfer {
	return t.set("optionValueField", value)
}

// PageField sets the pagination field
func (t tabsTransfer) PageField(value string) tabsTransfer {
	return t.set("pageField", value)
}

// Paginate sets whether pagination is enabled
func (t tabsTransfer) Paginate(value bool) tabsTransfer {
	return t.set("paginate", value)
}

// PageSizeField sets the field for page size
func (t tabsTransfer) PageSizeField(value string) tabsTransfer {
	return t.set("pageSizeField", value)
}

// Placeholder sets the placeholder text
func (t tabsTransfer) Placeholder(value string) tabsTransfer {
	return t.set("placeholder", value)
}

// ReadOnly sets whether the field is read-only
func (t tabsTransfer) ReadOnly(value bool) tabsTransfer {
	return t.set("readOnly", value)
}

// ReadOnlyOn sets the expression to determine if the field is read-only
func (t tabsTransfer) ReadOnlyOn(value string) tabsTransfer {
	return t.set("readOnlyOn", value)
}

// RemoveBtnLabel sets the label for the remove button
func (t tabsTransfer) RemoveBtnLabel(value string) tabsTransfer {
	return t.set("removeBtnLabel", value)
}

// Remote sets the remote data source
func (t tabsTransfer) Remote(value string) tabsTransfer {
	return t.set("remote", value)
}

// ResetValue sets the default value for reset
func (t tabsTransfer) ResetValue(value string) tabsTransfer {
	return t.set("resetValue", value)
}

// Searchable sets whether the field is searchable
func (t tabsTransfer) Searchable(value bool) tabsTransfer {
	return t.set("searchable", value)
}

// ShowBorder sets whether to show the border
func (t tabsTransfer) ShowBorder(value bool) tabsTransfer {
	return t.set("showBorder", value)
}

// ShowInputPlaceholder sets whether to show the input placeholder
func (t tabsTransfer) ShowInputPlaceholder(value bool) tabsTransfer {
	return t.set("showInputPlaceholder", value)
}

// ShowLabel sets whether to show the label
func (t tabsTransfer) ShowLabel(value bool) tabsTransfer {
	return t.set("showLabel", value)
}

// ShowRemovable sets whether options can be removed
func (t tabsTransfer) ShowRemovable(value bool) tabsTransfer {
	return t.set("showRemovable", value)
}

// ShowSearch sets whether options can be searched
func (t tabsTransfer) ShowSearch(value bool) tabsTransfer {
	return t.set("showSearch", value)
}

// Size sets the size
func (t tabsTransfer) Size(value string) tabsTransfer {
	return t.set("size", value)
}

// Source sets the data source API
func (t tabsTransfer) Source(value string) tabsTransfer {
	return t.set("source", value)
}

// StaticClassName sets the CSS class name for the static part
func (t tabsTransfer) StaticClassName(value string) tabsTransfer {
	return t.set("staticClassName", value)
}

// Step sets the step value
func (t tabsTransfer) Step(value int) tabsTransfer {
	return t.set("step", value)
}

// TabField sets the tab field for pagination
func (t tabsTransfer) TabField(value string) tabsTransfer {
	return t.set("tabField", value)
}

// Tabs sets the content for tabs
func (t tabsTransfer) Tabs(value string) tabsTransfer {
	return t.set("tabs", value)
}

// Toggable sets whether the field is toggable
func (t tabsTransfer) Toggable(value bool) tabsTransfer {
	return t.set("toggable", value)
}

// TreeMode sets whether tree mode is enabled
func (t tabsTransfer) TreeMode(value bool) tabsTransfer {
	return t.set("treeMode", value)
}

// Value sets the selected value
func (t tabsTransfer) Value(value string) tabsTransfer {
	return t.set("value", value)
}

// Visible sets whether the field is visible
func (t tabsTransfer) Visible(value bool) tabsTransfer {
	return t.set("visible", value)
}

// VisibleOn sets the expression to determine if the field is visible
func (t tabsTransfer) VisibleOn(value string) tabsTransfer {
	return t.set("visibleOn", value)
}

// Width sets the width
func (t tabsTransfer) Width(value string) tabsTransfer {
	return t.set("width", value)
}

// WrapperClassName sets the CSS class name for the wrapper
func (t tabsTransfer) WrapperClassName(value string) tabsTransfer {
	return t.set("wrapperClassName", value)
}

// WrapperStyle sets the style for the wrapper
func (t tabsTransfer) WrapperStyle(value any) tabsTransfer {
	return t.set("wrapperStyle", value)
}
