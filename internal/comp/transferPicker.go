package comp

import "github.com/zrcoder/amisgo/schema"

// TransferPicker represents the transfer picker component
type TransferPicker schema.Schema

func NewTransferPicker() TransferPicker {
	return TransferPicker{"type": "transfer-picker"}
}

func (tpc TransferPicker) set(key string, value any) TransferPicker {
	tpc[key] = value
	return tpc
}

// AddApi sets the API for adding items
func (tpc TransferPicker) AddApi(value string) TransferPicker {
	return tpc.set("addApi", value)
}

// AddControls sets the form items for adding
func (tpc TransferPicker) AddControls(value string) TransferPicker {
	return tpc.set("addControls", value)
}

// AddDialog sets the dialog for adding items
func (tpc TransferPicker) AddDialog(value string) TransferPicker {
	return tpc.set("addDialog", value)
}

// AutoCheckChildren sets whether to auto-check children
func (tpc TransferPicker) AutoCheckChildren(value bool) TransferPicker {
	return tpc.set("autoCheckChildren", value)
}

// AutoFill sets the auto-fill value
func (tpc TransferPicker) AutoFill(value string) TransferPicker {
	return tpc.set("autoFill", value)
}

// BorderMode sets the border mode
func (tpc TransferPicker) BorderMode(value string) TransferPicker {
	return tpc.set("borderMode", value)
}

// ClassName sets the CSS class name
func (tpc TransferPicker) ClassName(value string) TransferPicker {
	return tpc.set("className", value)
}

// ClearValueOnHidden sets whether to clear value when hidden
func (tpc TransferPicker) ClearValueOnHidden(value bool) TransferPicker {
	return tpc.set("clearValueOnHidden", value)
}

// Clearable sets whether the component is clearable
func (tpc TransferPicker) Clearable(value bool) TransferPicker {
	return tpc.set("clearable", value)
}

// Columns sets the columns for table mode
func (tpc TransferPicker) Columns(value ...any) TransferPicker {
	return tpc.set("columns", value)
}

// Creatable sets whether new items can be created
func (tpc TransferPicker) Creatable(value bool) TransferPicker {
	return tpc.set("creatable", value)
}

// CreateBtnLabel sets the label for the create button
func (tpc TransferPicker) CreateBtnLabel(value string) TransferPicker {
	return tpc.set("createBtnLabel", value)
}

// DeferApi sets the API for deferred loading
func (tpc TransferPicker) DeferApi(value string) TransferPicker {
	return tpc.set("deferApi", value)
}

// DeferField sets the field for deferred loading
func (tpc TransferPicker) DeferField(value string) TransferPicker {
	return tpc.set("deferField", value)
}

// DeleteApi sets the API for deleting items
func (tpc TransferPicker) DeleteApi(value string) TransferPicker {
	return tpc.set("deleteApi", value)
}

// DeleteConfirmText sets the confirmation text for deletion
func (tpc TransferPicker) DeleteConfirmText(value string) TransferPicker {
	return tpc.set("deleteConfirmText", value)
}

// Delimiter sets the delimiter
func (tpc TransferPicker) Delimiter(value string) TransferPicker {
	return tpc.set("delimiter", value)
}

// Desc sets the description
func (tpc TransferPicker) Desc(value string) TransferPicker {
	return tpc.set("desc", value)
}

// Description sets the description content
func (tpc TransferPicker) Description(value string) TransferPicker {
	return tpc.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description
func (tpc TransferPicker) DescriptionClassName(value string) TransferPicker {
	return tpc.set("descriptionClassName", value)
}

// Disabled sets whether the component is disabled
func (tpc TransferPicker) Disabled(value bool) TransferPicker {
	return tpc.set("disabled", value)
}

// DisabledOn sets the expression for disabling the component
func (tpc TransferPicker) DisabledOn(value string) TransferPicker {
	return tpc.set("disabledOn", value)
}

// EditApi sets the API for editing items
func (tpc TransferPicker) EditApi(value string) TransferPicker {
	return tpc.set("editApi", value)
}

// EditControls sets the form items for editing
func (tpc TransferPicker) EditControls(value string) TransferPicker {
	return tpc.set("editControls", value)
}

// EditDialog sets the dialog for editing items
func (tpc TransferPicker) EditDialog(value string) TransferPicker {
	return tpc.set("editDialog", value)
}

// Editable sets whether the component is editable
func (tpc TransferPicker) Editable(value bool) TransferPicker {
	return tpc.set("editable", value)
}

// EditorSetting sets the editor configuration
func (tpc TransferPicker) EditorSetting(value string) TransferPicker {
	return tpc.set("editorSetting", value)
}

// ExtraName sets the extra field name
func (tpc TransferPicker) ExtraName(value string) TransferPicker {
	return tpc.set("extraName", value)
}

// ExtractValue sets whether to extract the value as an array
func (tpc TransferPicker) ExtractValue(value bool) TransferPicker {
	return tpc.set("extractValue", value)
}

// Hidden sets whether the component is hidden
func (tpc TransferPicker) Hidden(value bool) TransferPicker {
	return tpc.set("hidden", value)
}

// HiddenOn sets the expression for hiding the component
func (tpc TransferPicker) HiddenOn(value string) TransferPicker {
	return tpc.set("hiddenOn", value)
}

// Hint sets the input hint
func (tpc TransferPicker) Hint(value string) TransferPicker {
	return tpc.set("hint", value)
}

// Horizontal sets the horizontal layout configuration
func (tpc TransferPicker) Horizontal(value string) TransferPicker {
	return tpc.set("horizontal", value)
}

// ID sets the unique component ID
func (tpc TransferPicker) ID(value string) TransferPicker {
	return tpc.set("id", value)
}

// InitAutoFill sets the initial auto-fill value
func (tpc TransferPicker) InitAutoFill(value string) TransferPicker {
	return tpc.set("initAutoFill", value)
}

// InitFetch sets whether to fetch data initially
func (tpc TransferPicker) InitFetch(value bool) TransferPicker {
	return tpc.set("initFetch", value)
}

// InitFetchOn sets the expression for initial data fetch
func (tpc TransferPicker) InitFetchOn(value string) TransferPicker {
	return tpc.set("initFetchOn", value)
}

// InitiallyOpen sets whether the component is initially open
func (tpc TransferPicker) InitiallyOpen(value bool) TransferPicker {
	return tpc.set("initiallyOpen", value)
}

// Inline sets whether the component is inline
func (tpc TransferPicker) Inline(value bool) TransferPicker {
	return tpc.set("inline", value)
}

// InputClassName sets the CSS class name for the input
func (tpc TransferPicker) InputClassName(value string) TransferPicker {
	return tpc.set("inputClassName", value)
}

// ItemHeight sets the height of each item
func (tpc TransferPicker) ItemHeight(value string) TransferPicker {
	return tpc.set("itemHeight", value)
}

// JoinValues sets whether to join values in single select mode
func (tpc TransferPicker) JoinValues(value bool) TransferPicker {
	return tpc.set("joinValues", value)
}

// Label sets the label text
func (tpc TransferPicker) Label(value string) TransferPicker {
	return tpc.set("label", value)
}

// LabelAlign sets the label alignment
func (tpc TransferPicker) LabelAlign(value string) TransferPicker {
	return tpc.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label
func (tpc TransferPicker) LabelClassName(value string) TransferPicker {
	return tpc.set("labelClassName", value)
}

// LabelRemark sets the label remark
func (tpc TransferPicker) LabelRemark(value string) TransferPicker {
	return tpc.set("labelRemark", value)
}

// Max sets the maximum number of options
func (tpc TransferPicker) Max(value int) TransferPicker {
	return tpc.set("max", value)
}

// MergeDataSource sets whether to merge the source data
func (tpc TransferPicker) MergeDataSource(value bool) TransferPicker {
	return tpc.set("mergeDataSource", value)
}

// Min sets the minimum number of options
func (tpc TransferPicker) Min(value int) TransferPicker {
	return tpc.set("min", value)
}

// MultiLine sets whether to display in multi-line mode
func (tpc TransferPicker) MultiLine(value bool) TransferPicker {
	return tpc.set("multiLine", value)
}

// Name sets the component name
func (tpc TransferPicker) Name(value string) TransferPicker {
	return tpc.set("name", value)
}

// Options sets the data source options
func (tpc TransferPicker) Options(value ...any) TransferPicker {
	return tpc.set("options", value)
}

// Placeholder sets the placeholder text
func (tpc TransferPicker) Placeholder(value string) TransferPicker {
	return tpc.set("placeholder", value)
}

// ReadOnly sets whether the component is read-only
func (tpc TransferPicker) ReadOnly(value bool) TransferPicker {
	return tpc.set("readOnly", value)
}

// Remark sets the remark
func (tpc TransferPicker) Remark(value string) TransferPicker {
	return tpc.set("remark", value)
}

// Removable sets whether the component is removable
func (tpc TransferPicker) Removable(value bool) TransferPicker {
	return tpc.set("removable", value)
}

// Renderer sets the component renderer type
func (tpc TransferPicker) Renderer(value string) TransferPicker {
	return tpc.set("renderer", value)
}

// Source sets the data source
func (tpc TransferPicker) Source(value string) TransferPicker {
	return tpc.set("source", value)
}

// SourceEmptyText sets the text when the source is empty
func (tpc TransferPicker) SourceEmptyText(value string) TransferPicker {
	return tpc.set("sourceEmptyText", value)
}

// Static sets whether the component is static
func (tpc TransferPicker) Static(value bool) TransferPicker {
	return tpc.set("static", value)
}

// StaticClassName sets the CSS class name for the static component
func (tpc TransferPicker) StaticClassName(value string) TransferPicker {
	return tpc.set("staticClassName", value)
}

// Unset removes a specific option
func (tpc TransferPicker) Unset(value string) TransferPicker {
	return tpc.set("unset", value)
}

// Value sets the value
func (tpc TransferPicker) Value(value string) TransferPicker {
	return tpc.set("value", value)
}

// ValueField sets the value field for the data source
func (tpc TransferPicker) ValueField(value string) TransferPicker {
	return tpc.set("valueField", value)
}

// ValueJoin sets the field delimiter
func (tpc TransferPicker) ValueJoin(value string) TransferPicker {
	return tpc.set("valueJoin", value)
}

// Vertical sets the vertical layout configuration
func (tpc TransferPicker) Vertical(value string) TransferPicker {
	return tpc.set("vertical", value)
}

// VisibleOn sets the expression for visibility
func (tpc TransferPicker) VisibleOn(value string) TransferPicker {
	return tpc.set("visibleOn", value)
}

// Width sets the width
func (tpc TransferPicker) Width(value string) TransferPicker {
	return tpc.set("width", value)
}

// Wrap sets the minimum width for collapsed display
func (tpc TransferPicker) Wrap(value string) TransferPicker {
	return tpc.set("wrap", value)
}

// ZIndex sets the z-index
func (tpc TransferPicker) ZIndex(value int) TransferPicker {
	return tpc.set("zIndex", value)
}
