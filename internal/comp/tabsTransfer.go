package comp

import "github.com/zrcoder/amisgo/schema"

// TabsTransfer represents the schema.Schema for tabs transfer component
// documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/form/tabs-transfer
type TabsTransfer schema.Schema

func NewTabsTransfer() TabsTransfer {
	return TabsTransfer{"type": "tabs-transfer"}
}

// set sets a key-value pair in tabsTransfer
func (t TabsTransfer) set(key string, value any) TabsTransfer {
	t[key] = value
	return t
}

// AddApi sets the API for adding items
func (t TabsTransfer) AddApi(value string) TabsTransfer {
	return t.set("addApi", value)
}

// Adds sets the form items for adding new entries
func (t TabsTransfer) Adds(value string) TabsTransfer {
	return t.set("adds", value)
}

// AddDialog sets the configuration for the add dialog
func (t TabsTransfer) AddDialog(value string) TabsTransfer {
	return t.set("addDialog", value)
}

// AutoCheckChildren sets whether to auto-check children items
func (t TabsTransfer) AutoCheckChildren(value bool) TabsTransfer {
	return t.set("autoCheckChildren", value)
}

// AutoFill sets the auto-fill configuration
func (t TabsTransfer) AutoFill(value string) TabsTransfer {
	return t.set("autoFill", value)
}

// ClassName sets the CSS class name for the container
func (t TabsTransfer) ClassName(value string) TabsTransfer {
	return t.set("className", value)
}

// ClearValueOnHidden sets whether to clear the value when hidden
func (t TabsTransfer) ClearValueOnHidden(value bool) TabsTransfer {
	return t.set("clearValueOnHidden", value)
}

// Clearable sets whether the field is clearable
func (t TabsTransfer) Clearable(value bool) TabsTransfer {
	return t.set("clearable", value)
}

// Columns sets the columns for table mode
func (t TabsTransfer) Columns(value ...any) TabsTransfer {
	return t.set("columns", value)
}

// Creatable sets whether new entries can be created
func (t TabsTransfer) Creatable(value bool) TabsTransfer {
	return t.set("creatable", value)
}

// CreateBtnLabel sets the label for the create button
func (t TabsTransfer) CreateBtnLabel(value string) TabsTransfer {
	return t.set("createBtnLabel", value)
}

// DeferApi sets the API for deferred loading
func (t TabsTransfer) DeferApi(value string) TabsTransfer {
	return t.set("deferApi", value)
}

// DeferField sets the field for deferred loading
func (t TabsTransfer) DeferField(value string) TabsTransfer {
	return t.set("deferField", value)
}

// DeleteApi sets the API for deleting items
func (t TabsTransfer) DeleteApi(value string) TabsTransfer {
	return t.set("deleteApi", value)
}

// DeleteConfirmText sets the confirmation text for deletion
func (t TabsTransfer) DeleteConfirmText(value string) TabsTransfer {
	return t.set("deleteConfirmText", value)
}

// Delimiter sets the delimiter for values
func (t TabsTransfer) Delimiter(value string) TabsTransfer {
	return t.set("delimiter", value)
}

// Desc sets the description
func (t TabsTransfer) Desc(value string) TabsTransfer {
	return t.set("desc", value)
}

// Description sets the HTML description content
func (t TabsTransfer) Description(value string) TabsTransfer {
	return t.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description
func (t TabsTransfer) DescriptionClassName(value string) TabsTransfer {
	return t.set("descriptionClassName", value)
}

// Disabled sets whether the field is disabled
func (t TabsTransfer) Disabled(value bool) TabsTransfer {
	return t.set("disabled", value)
}

// DisabledOn sets the expression to determine if the field is disabled
func (t TabsTransfer) DisabledOn(value string) TabsTransfer {
	return t.set("disabledOn", value)
}

// EditApi sets the API for editing items
func (t TabsTransfer) EditApi(value string) TabsTransfer {
	return t.set("editApi", value)
}

// Edits sets the form items for editing entries
func (t TabsTransfer) Edits(value string) TabsTransfer {
	return t.set("edits", value)
}

// EditDialog sets the configuration for the edit dialog
func (t TabsTransfer) EditDialog(value string) TabsTransfer {
	return t.set("editDialog", value)
}

// Editable sets whether the field is editable
func (t TabsTransfer) Editable(value bool) TabsTransfer {
	return t.set("editable", value)
}

// EditorSetting sets the editor configuration
func (t TabsTransfer) EditorSetting(value string) TabsTransfer {
	return t.set("editorSetting", value)
}

// ExtraName sets the extra field name for range components
func (t TabsTransfer) ExtraName(value string) TabsTransfer {
	return t.set("extraName", value)
}

// ExtractValue sets whether to extract the value as an array
func (t TabsTransfer) ExtractValue(value bool) TabsTransfer {
	return t.set("extractValue", value)
}

// Hidden sets whether the field is hidden
func (t TabsTransfer) Hidden(value bool) TabsTransfer {
	return t.set("hidden", value)
}

// HiddenOn sets the expression to determine if the field is hidden
func (t TabsTransfer) HiddenOn(value string) TabsTransfer {
	return t.set("hiddenOn", value)
}

// Hint sets the input hint
func (t TabsTransfer) Hint(value string) TabsTransfer {
	return t.set("hint", value)
}

// Horizontal sets the horizontal layout configuration
func (t TabsTransfer) Horizontal(value string) TabsTransfer {
	return t.set("horizontal", value)
}

// ID sets the unique ID for the component
func (t TabsTransfer) ID(value string) TabsTransfer {
	return t.set("id", value)
}

// InitAutoFill sets the initial auto-fill configuration
func (t TabsTransfer) InitAutoFill(value string) TabsTransfer {
	return t.set("initAutoFill", value)
}

// InitFetch sets whether to fetch data initially
func (t TabsTransfer) InitFetch(value bool) TabsTransfer {
	return t.set("initFetch", value)
}

// InitFetchOn sets the expression to determine if data should be fetched initially
func (t TabsTransfer) InitFetchOn(value string) TabsTransfer {
	return t.set("initFetchOn", value)
}

// InitiallyOpen sets whether the tabs are initially open
func (t TabsTransfer) InitiallyOpen(value bool) TabsTransfer {
	return t.set("initiallyOpen", value)
}

// Inline sets whether the form is in inline mode
func (t TabsTransfer) Inline(value bool) TabsTransfer {
	return t.set("inline", value)
}

// InputClassName sets the CSS class name for the input
func (t TabsTransfer) InputClassName(value string) TabsTransfer {
	return t.set("inputClassName", value)
}

// ItemHeight sets the height of each item for virtual rendering
func (t TabsTransfer) ItemHeight(value string) TabsTransfer {
	return t.set("itemHeight", value)
}

// JoinValues sets whether to join values with a delimiter
func (t TabsTransfer) JoinValues(value bool) TabsTransfer {
	return t.set("joinValues", value)
}

// Label sets the label text
func (t TabsTransfer) Label(value string) TabsTransfer {
	return t.set("label", value)
}

// LabelAlign sets the label alignment
func (t TabsTransfer) LabelAlign(value string) TabsTransfer {
	return t.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label
func (t TabsTransfer) LabelClassName(value string) TabsTransfer {
	return t.set("labelClassName", value)
}

// LabelRemark sets the remark for the label
func (t TabsTransfer) LabelRemark(value string) TabsTransfer {
	return t.set("labelRemark", value)
}

// LabelWidth sets the width of the label
func (t TabsTransfer) LabelWidth(value string) TabsTransfer {
	return t.set("labelWidth", value)
}

// Level sets the level
func (t TabsTransfer) Level(value string) TabsTransfer {
	return t.set("level", value)
}

// Max sets the maximum value
func (t TabsTransfer) Max(value int) TabsTransfer {
	return t.set("max", value)
}

// Min sets the minimum value
func (t TabsTransfer) Min(value int) TabsTransfer {
	return t.set("min", value)
}

// Name sets the unique name for the component
func (t TabsTransfer) Name(value string) TabsTransfer {
	return t.set("name", value)
}

// Nested sets whether nested options are supported
func (t TabsTransfer) Nested(value bool) TabsTransfer {
	return t.set("nested", value)
}

// Options sets the options
func (t TabsTransfer) Options(value ...any) TabsTransfer {
	return t.set("options", value)
}

// OptionTextField sets the text field for options
func (t TabsTransfer) OptionTextField(value string) TabsTransfer {
	return t.set("optionTextField", value)
}

// OptionValueField sets the value field for options
func (t TabsTransfer) OptionValueField(value string) TabsTransfer {
	return t.set("optionValueField", value)
}

// PageField sets the pagination field
func (t TabsTransfer) PageField(value string) TabsTransfer {
	return t.set("pageField", value)
}

// Paginate sets whether pagination is enabled
func (t TabsTransfer) Paginate(value bool) TabsTransfer {
	return t.set("paginate", value)
}

// PageSizeField sets the field for page size
func (t TabsTransfer) PageSizeField(value string) TabsTransfer {
	return t.set("pageSizeField", value)
}

// Placeholder sets the placeholder text
func (t TabsTransfer) Placeholder(value string) TabsTransfer {
	return t.set("placeholder", value)
}

// ReadOnly sets whether the field is read-only
func (t TabsTransfer) ReadOnly(value bool) TabsTransfer {
	return t.set("readOnly", value)
}

// ReadOnlyOn sets the expression to determine if the field is read-only
func (t TabsTransfer) ReadOnlyOn(value string) TabsTransfer {
	return t.set("readOnlyOn", value)
}

// RemoveBtnLabel sets the label for the remove button
func (t TabsTransfer) RemoveBtnLabel(value string) TabsTransfer {
	return t.set("removeBtnLabel", value)
}

// Remote sets the remote data source
func (t TabsTransfer) Remote(value string) TabsTransfer {
	return t.set("remote", value)
}

// ResetValue sets the default value for reset
func (t TabsTransfer) ResetValue(value string) TabsTransfer {
	return t.set("resetValue", value)
}

// Searchable sets whether the field is searchable
func (t TabsTransfer) Searchable(value bool) TabsTransfer {
	return t.set("searchable", value)
}

// ShowBorder sets whether to show the border
func (t TabsTransfer) ShowBorder(value bool) TabsTransfer {
	return t.set("showBorder", value)
}

// ShowInputPlaceholder sets whether to show the input placeholder
func (t TabsTransfer) ShowInputPlaceholder(value bool) TabsTransfer {
	return t.set("showInputPlaceholder", value)
}

// ShowLabel sets whether to show the label
func (t TabsTransfer) ShowLabel(value bool) TabsTransfer {
	return t.set("showLabel", value)
}

// ShowRemovable sets whether options can be removed
func (t TabsTransfer) ShowRemovable(value bool) TabsTransfer {
	return t.set("showRemovable", value)
}

// ShowSearch sets whether options can be searched
func (t TabsTransfer) ShowSearch(value bool) TabsTransfer {
	return t.set("showSearch", value)
}

// Size sets the size
func (t TabsTransfer) Size(value string) TabsTransfer {
	return t.set("size", value)
}

// Source sets the data source API
func (t TabsTransfer) Source(value string) TabsTransfer {
	return t.set("source", value)
}

// StaticClassName sets the CSS class name for the static part
func (t TabsTransfer) StaticClassName(value string) TabsTransfer {
	return t.set("staticClassName", value)
}

// Step sets the step value
func (t TabsTransfer) Step(value int) TabsTransfer {
	return t.set("step", value)
}

// TabField sets the tab field for pagination
func (t TabsTransfer) TabField(value string) TabsTransfer {
	return t.set("tabField", value)
}

// Tabs sets the content for tabs
func (t TabsTransfer) Tabs(value string) TabsTransfer {
	return t.set("tabs", value)
}

// Toggable sets whether the field is toggable
func (t TabsTransfer) Toggable(value bool) TabsTransfer {
	return t.set("toggable", value)
}

// TreeMode sets whether tree mode is enabled
func (t TabsTransfer) TreeMode(value bool) TabsTransfer {
	return t.set("treeMode", value)
}

// Value sets the selected value
func (t TabsTransfer) Value(value string) TabsTransfer {
	return t.set("value", value)
}

// Visible sets whether the field is visible
func (t TabsTransfer) Visible(value bool) TabsTransfer {
	return t.set("visible", value)
}

// VisibleOn sets the expression to determine if the field is visible
func (t TabsTransfer) VisibleOn(value string) TabsTransfer {
	return t.set("visibleOn", value)
}

// Width sets the width
func (t TabsTransfer) Width(value string) TabsTransfer {
	return t.set("width", value)
}

// WrapperClassName sets the CSS class name for the wrapper
func (t TabsTransfer) WrapperClassName(value string) TabsTransfer {
	return t.set("wrapperClassName", value)
}

// WrapperStyle sets the style for the wrapper
func (t TabsTransfer) WrapperStyle(value any) TabsTransfer {
	return t.set("wrapperStyle", value)
}
