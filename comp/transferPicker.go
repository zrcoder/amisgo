package comp

// transferPicker represents the transfer picker component

type transferPicker Schema

func TransferPicker() transferPicker {
	return transferPicker{}.set("type", "transfer-picker")
}

func (tpc transferPicker) set(key string, value any) transferPicker {
	tpc[key] = value
	return tpc
}

// AddApi sets the API for adding items
func (tpc transferPicker) AddApi(value string) transferPicker {
	return tpc.set("addApi", value)
}

// AddControls sets the form items for adding
func (tpc transferPicker) AddControls(value string) transferPicker {
	return tpc.set("addControls", value)
}

// AddDialog sets the dialog for adding items
func (tpc transferPicker) AddDialog(value string) transferPicker {
	return tpc.set("addDialog", value)
}

// AutoCheckChildren sets whether to auto-check children
func (tpc transferPicker) AutoCheckChildren(value bool) transferPicker {
	return tpc.set("autoCheckChildren", value)
}

// AutoFill sets the auto-fill value
func (tpc transferPicker) AutoFill(value string) transferPicker {
	return tpc.set("autoFill", value)
}

// BorderMode sets the border mode
func (tpc transferPicker) BorderMode(value string) transferPicker {
	return tpc.set("borderMode", value)
}

// ClassName sets the CSS class name
func (tpc transferPicker) ClassName(value string) transferPicker {
	return tpc.set("className", value)
}

// ClearValueOnHidden sets whether to clear value when hidden
func (tpc transferPicker) ClearValueOnHidden(value bool) transferPicker {
	return tpc.set("clearValueOnHidden", value)
}

// Clearable sets whether the component is clearable
func (tpc transferPicker) Clearable(value bool) transferPicker {
	return tpc.set("clearable", value)
}

// Columns sets the columns for table mode
func (tpc transferPicker) Columns(value ...any) transferPicker {
	return tpc.set("columns", value)
}

// Creatable sets whether new items can be created
func (tpc transferPicker) Creatable(value bool) transferPicker {
	return tpc.set("creatable", value)
}

// CreateBtnLabel sets the label for the create button
func (tpc transferPicker) CreateBtnLabel(value string) transferPicker {
	return tpc.set("createBtnLabel", value)
}

// DeferApi sets the API for deferred loading
func (tpc transferPicker) DeferApi(value string) transferPicker {
	return tpc.set("deferApi", value)
}

// DeferField sets the field for deferred loading
func (tpc transferPicker) DeferField(value string) transferPicker {
	return tpc.set("deferField", value)
}

// DeleteApi sets the API for deleting items
func (tpc transferPicker) DeleteApi(value string) transferPicker {
	return tpc.set("deleteApi", value)
}

// DeleteConfirmText sets the confirmation text for deletion
func (tpc transferPicker) DeleteConfirmText(value string) transferPicker {
	return tpc.set("deleteConfirmText", value)
}

// Delimiter sets the delimiter
func (tpc transferPicker) Delimiter(value string) transferPicker {
	return tpc.set("delimiter", value)
}

// Desc sets the description
func (tpc transferPicker) Desc(value string) transferPicker {
	return tpc.set("desc", value)
}

// Description sets the description content
func (tpc transferPicker) Description(value string) transferPicker {
	return tpc.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description
func (tpc transferPicker) DescriptionClassName(value string) transferPicker {
	return tpc.set("descriptionClassName", value)
}

// Disabled sets whether the component is disabled
func (tpc transferPicker) Disabled(value bool) transferPicker {
	return tpc.set("disabled", value)
}

// DisabledOn sets the expression for disabling the component
func (tpc transferPicker) DisabledOn(value string) transferPicker {
	return tpc.set("disabledOn", value)
}

// EditApi sets the API for editing items
func (tpc transferPicker) EditApi(value string) transferPicker {
	return tpc.set("editApi", value)
}

// EditControls sets the form items for editing
func (tpc transferPicker) EditControls(value string) transferPicker {
	return tpc.set("editControls", value)
}

// EditDialog sets the dialog for editing items
func (tpc transferPicker) EditDialog(value string) transferPicker {
	return tpc.set("editDialog", value)
}

// Editable sets whether the component is editable
func (tpc transferPicker) Editable(value bool) transferPicker {
	return tpc.set("editable", value)
}

// EditorSetting sets the editor configuration
func (tpc transferPicker) EditorSetting(value string) transferPicker {
	return tpc.set("editorSetting", value)
}

// ExtraName sets the extra field name
func (tpc transferPicker) ExtraName(value string) transferPicker {
	return tpc.set("extraName", value)
}

// ExtractValue sets whether to extract the value as an array
func (tpc transferPicker) ExtractValue(value bool) transferPicker {
	return tpc.set("extractValue", value)
}

// Hidden sets whether the component is hidden
func (tpc transferPicker) Hidden(value bool) transferPicker {
	return tpc.set("hidden", value)
}

// HiddenOn sets the expression for hiding the component
func (tpc transferPicker) HiddenOn(value string) transferPicker {
	return tpc.set("hiddenOn", value)
}

// Hint sets the input hint
func (tpc transferPicker) Hint(value string) transferPicker {
	return tpc.set("hint", value)
}

// Horizontal sets the horizontal layout configuration
func (tpc transferPicker) Horizontal(value string) transferPicker {
	return tpc.set("horizontal", value)
}

// ID sets the unique component ID
func (tpc transferPicker) ID(value string) transferPicker {
	return tpc.set("id", value)
}

// InitAutoFill sets the initial auto-fill value
func (tpc transferPicker) InitAutoFill(value string) transferPicker {
	return tpc.set("initAutoFill", value)
}

// InitFetch sets whether to fetch data initially
func (tpc transferPicker) InitFetch(value bool) transferPicker {
	return tpc.set("initFetch", value)
}

// InitFetchOn sets the expression for initial data fetch
func (tpc transferPicker) InitFetchOn(value string) transferPicker {
	return tpc.set("initFetchOn", value)
}

// InitiallyOpen sets whether the component is initially open
func (tpc transferPicker) InitiallyOpen(value bool) transferPicker {
	return tpc.set("initiallyOpen", value)
}

// Inline sets whether the component is inline
func (tpc transferPicker) Inline(value bool) transferPicker {
	return tpc.set("inline", value)
}

// InputClassName sets the CSS class name for the input
func (tpc transferPicker) InputClassName(value string) transferPicker {
	return tpc.set("inputClassName", value)
}

// ItemHeight sets the height of each item
func (tpc transferPicker) ItemHeight(value string) transferPicker {
	return tpc.set("itemHeight", value)
}

// JoinValues sets whether to join values in single select mode
func (tpc transferPicker) JoinValues(value bool) transferPicker {
	return tpc.set("joinValues", value)
}

// Label sets the label text
func (tpc transferPicker) Label(value string) transferPicker {
	return tpc.set("label", value)
}

// LabelAlign sets the label alignment
func (tpc transferPicker) LabelAlign(value string) transferPicker {
	return tpc.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label
func (tpc transferPicker) LabelClassName(value string) transferPicker {
	return tpc.set("labelClassName", value)
}

// LabelRemark sets the label remark
func (tpc transferPicker) LabelRemark(value string) transferPicker {
	return tpc.set("labelRemark", value)
}

// Max sets the maximum number of options
func (tpc transferPicker) Max(value int) transferPicker {
	return tpc.set("max", value)
}

// MergeDataSource sets whether to merge the source data
func (tpc transferPicker) MergeDataSource(value bool) transferPicker {
	return tpc.set("mergeDataSource", value)
}

// Min sets the minimum number of options
func (tpc transferPicker) Min(value int) transferPicker {
	return tpc.set("min", value)
}

// MultiLine sets whether to display in multi-line mode
func (tpc transferPicker) MultiLine(value bool) transferPicker {
	return tpc.set("multiLine", value)
}

// Name sets the component name
func (tpc transferPicker) Name(value string) transferPicker {
	return tpc.set("name", value)
}

// Options sets the data source options
func (tpc transferPicker) Options(value ...any) transferPicker {
	return tpc.set("options", value)
}

// Placeholder sets the placeholder text
func (tpc transferPicker) Placeholder(value string) transferPicker {
	return tpc.set("placeholder", value)
}

// ReadOnly sets whether the component is read-only
func (tpc transferPicker) ReadOnly(value bool) transferPicker {
	return tpc.set("readOnly", value)
}

// Remark sets the remark
func (tpc transferPicker) Remark(value string) transferPicker {
	return tpc.set("remark", value)
}

// Removable sets whether the component is removable
func (tpc transferPicker) Removable(value bool) transferPicker {
	return tpc.set("removable", value)
}

// Renderer sets the component renderer type
func (tpc transferPicker) Renderer(value string) transferPicker {
	return tpc.set("renderer", value)
}

// Source sets the data source
func (tpc transferPicker) Source(value string) transferPicker {
	return tpc.set("source", value)
}

// SourceEmptyText sets the text when the source is empty
func (tpc transferPicker) SourceEmptyText(value string) transferPicker {
	return tpc.set("sourceEmptyText", value)
}

// Static sets whether the component is static
func (tpc transferPicker) Static(value bool) transferPicker {
	return tpc.set("static", value)
}

// StaticClassName sets the CSS class name for the static component
func (tpc transferPicker) StaticClassName(value string) transferPicker {
	return tpc.set("staticClassName", value)
}

// Unset removes a specific option
func (tpc transferPicker) Unset(value string) transferPicker {
	return tpc.set("unset", value)
}

// Value sets the value
func (tpc transferPicker) Value(value string) transferPicker {
	return tpc.set("value", value)
}

// ValueField sets the value field for the data source
func (tpc transferPicker) ValueField(value string) transferPicker {
	return tpc.set("valueField", value)
}

// ValueJoin sets the field delimiter
func (tpc transferPicker) ValueJoin(value string) transferPicker {
	return tpc.set("valueJoin", value)
}

// Vertical sets the vertical layout configuration
func (tpc transferPicker) Vertical(value string) transferPicker {
	return tpc.set("vertical", value)
}

// VisibleOn sets the expression for visibility
func (tpc transferPicker) VisibleOn(value string) transferPicker {
	return tpc.set("visibleOn", value)
}

// Width sets the width
func (tpc transferPicker) Width(value string) transferPicker {
	return tpc.set("width", value)
}

// Wrap sets the minimum width for collapsed display
func (tpc transferPicker) Wrap(value string) transferPicker {
	return tpc.set("wrap", value)
}

// ZIndex sets the z-index
func (tpc transferPicker) ZIndex(value int) transferPicker {
	return tpc.set("zIndex", value)
}
