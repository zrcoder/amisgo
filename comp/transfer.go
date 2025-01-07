package comp

// transfer control definition

type transfer Schema

// Transfer creates a new TransferControl instance
func Transfer() transfer {
	return transfer{}.set("type", "transfer")
}

func (tc transfer) set(key string, value any) transfer {
	tc[key] = value
	return tc
}

// AddApi sets the API for adding items
func (tc transfer) AddApi(value string) transfer {
	return tc.set("addApi", value)
}

// AddControls sets the form items for adding
func (tc transfer) AddControls(value string) transfer {
	return tc.set("addControls", value)
}

// AddDialog sets the dialog for adding items
func (tc transfer) AddDialog(value string) transfer {
	return tc.set("addDialog", value)
}

// AutoCheckChildren sets whether to auto-check children
func (tc transfer) AutoCheckChildren(value bool) transfer {
	return tc.set("autoCheckChildren", value)
}

// AutoFill sets the auto-fill value
func (tc transfer) AutoFill(value string) transfer {
	return tc.set("autoFill", value)
}

// ClassName sets the container CSS class name
func (tc transfer) ClassName(value string) transfer {
	return tc.set("className", value)
}

// ClearValueOnHidden sets whether to clear value when hidden
func (tc transfer) ClearValueOnHidden(value bool) transfer {
	return tc.set("clearValueOnHidden", value)
}

// Clearable sets whether the control is clearable
func (tc transfer) Clearable(value bool) transfer {
	return tc.set("clearable", value)
}

// Columns sets the columns for table mode
func (tc transfer) Columns(value ...any) transfer {
	return tc.set("columns", value)
}

// Creatable sets whether new items can be created
func (tc transfer) Creatable(value bool) transfer {
	return tc.set("creatable", value)
}

// CreateBtnLabel sets the label for the create button
func (tc transfer) CreateBtnLabel(value string) transfer {
	return tc.set("createBtnLabel", value)
}

// DeferApi sets the API for deferred loading
func (tc transfer) DeferApi(value string) transfer {
	return tc.set("deferApi", value)
}

// DeferField sets the field for deferred loading
func (tc transfer) DeferField(value string) transfer {
	return tc.set("deferField", value)
}

// DeleteApi sets the API for deleting items
func (tc transfer) DeleteApi(value string) transfer {
	return tc.set("deleteApi", value)
}

// DeleteConfirmText sets the confirmation text for deletion
func (tc transfer) DeleteConfirmText(value string) transfer {
	return tc.set("deleteConfirmText", value)
}

// Delimiter sets the delimiter
func (tc transfer) Delimiter(value string) transfer {
	return tc.set("delimiter", value)
}

// Desc sets the description
func (tc transfer) Desc(value string) transfer {
	return tc.set("desc", value)
}

// Description sets the description content
func (tc transfer) Description(value string) transfer {
	return tc.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description
func (tc transfer) DescriptionClassName(value string) transfer {
	return tc.set("descriptionClassName", value)
}

// Disabled sets whether the control is disabled
func (tc transfer) Disabled(value bool) transfer {
	return tc.set("disabled", value)
}

// DisabledOn sets the expression for disabling the control
func (tc transfer) DisabledOn(value string) transfer {
	return tc.set("disabledOn", value)
}

// EditApi sets the API for editing items
func (tc transfer) EditApi(value string) transfer {
	return tc.set("editApi", value)
}

// EditControls sets the form items for editing
func (tc transfer) EditControls(value string) transfer {
	return tc.set("editControls", value)
}

// EditDialog sets the dialog for editing items
func (tc transfer) EditDialog(value string) transfer {
	return tc.set("editDialog", value)
}

// Editable sets whether the control is editable
func (tc transfer) Editable(value bool) transfer {
	return tc.set("editable", value)
}

// EditorSetting sets the editor configuration
func (tc transfer) EditorSetting(value string) transfer {
	return tc.set("editorSetting", value)
}

// ExtraName sets the extra field name
func (tc transfer) ExtraName(value string) transfer {
	return tc.set("extraName", value)
}

// ExtractValue sets whether to extract the value as an array
func (tc transfer) ExtractValue(value bool) transfer {
	return tc.set("extractValue", value)
}

// Hidden sets whether the control is hidden
func (tc transfer) Hidden(value bool) transfer {
	return tc.set("hidden", value)
}

// HiddenOn sets the expression for hiding the control
func (tc transfer) HiddenOn(value string) transfer {
	return tc.set("hiddenOn", value)
}

// Hint sets the input hint
func (tc transfer) Hint(value string) transfer {
	return tc.set("hint", value)
}

// Horizontal sets the horizontal layout configuration
func (tc transfer) Horizontal(value string) transfer {
	return tc.set("horizontal", value)
}

// ID sets the unique ID for the component
func (tc transfer) ID(value string) transfer {
	return tc.set("id", value)
}

// InitAutoFill sets the initial auto-fill value
func (tc transfer) InitAutoFill(value string) transfer {
	return tc.set("initAutoFill", value)
}

// InitFetch sets whether to fetch data initially
func (tc transfer) InitFetch(value bool) transfer {
	return tc.set("initFetch", value)
}

// InitFetchOn sets the expression for initial data fetch
func (tc transfer) InitFetchOn(value string) transfer {
	return tc.set("initFetchOn", value)
}

// InitiallyOpen sets whether the control is initially open
func (tc transfer) InitiallyOpen(value bool) transfer {
	return tc.set("initiallyOpen", value)
}

// Inline sets whether the control is in inline mode
func (tc transfer) Inline(value bool) transfer {
	return tc.set("inline", value)
}

// InputClassName sets the CSS class name for the input
func (tc transfer) InputClassName(value string) transfer {
	return tc.set("inputClassName", value)
}

// ItemHeight sets the height of each item
func (tc transfer) ItemHeight(value string) transfer {
	return tc.set("itemHeight", value)
}

// JoinValues sets whether to join values with a delimiter
func (tc transfer) JoinValues(value bool) transfer {
	return tc.set("joinValues", value)
}

// Label sets the label
func (tc transfer) Label(value string) transfer {
	return tc.set("label", value)
}

// LabelAlign sets the label alignment
func (tc transfer) LabelAlign(value string) transfer {
	return tc.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label
func (tc transfer) LabelClassName(value string) transfer {
	return tc.set("labelClassName", value)
}

// LabelRemark sets the remark for the label
func (tc transfer) LabelRemark(value string) transfer {
	return tc.set("labelRemark", value)
}

// LabelWidth sets the custom width for the label
func (tc transfer) LabelWidth(value string) transfer {
	return tc.set("labelWidth", value)
}

// LeftMode sets the selection mode for the left side
func (tc transfer) LeftMode(value string) transfer {
	return tc.set("leftMode", value)
}

// Max sets the maximum number of selectable items
func (tc transfer) Max(value int) transfer {
	return tc.set("max", value)
}

// Min sets the minimum number of selectable items
func (tc transfer) Min(value int) transfer {
	return tc.set("min", value)
}

// Name sets the form item name
func (tc transfer) Name(value string) transfer {
	return tc.set("name", value)
}

// OnChange sets the callback function for changes
func (tc transfer) OnChange(value string) transfer {
	return tc.set("onChange", value)
}

// OnEvent sets the event handler
func (tc transfer) OnEvent(value any) transfer {
	return tc.set("onEvent", value)
}

// Option sets a single option
func (tc transfer) Option(value string) transfer {
	return tc.set("option", value)
}

// Options sets the options data
func (tc transfer) Options(value ...any) transfer {
	return tc.set("options", value)
}

// Placeholder sets the input placeholder
func (tc transfer) Placeholder(value string) transfer {
	return tc.set("placeholder", value)
}

// ReadOnly sets whether the control is read-only
func (tc transfer) ReadOnly(value bool) transfer {
	return tc.set("readOnly", value)
}

// RemoteOptions sets the remote options
func (tc transfer) RemoteOptions(value ...any) transfer {
	return tc.set("remoteOptions", value)
}

// Render sets the form render mode
func (tc transfer) Render(value string) transfer {
	return tc.set("render", value)
}

// ResetValue sets the default value when the form is reset
func (tc transfer) ResetValue(value string) transfer {
	return tc.set("resetValue", value)
}

// SaveImmediately sets whether to save immediately on change
func (tc transfer) SaveImmediately(value bool) transfer {
	return tc.set("saveImmediately", value)
}

// Source sets the data source
func (tc transfer) Source(value string) transfer {
	return tc.set("source", value)
}

// Searchable sets whether the control is searchable
func (tc transfer) Searchable(value bool) transfer {
	return tc.set("searchable", value)
}

// SearchApi sets the API for searching
func (tc transfer) SearchApi(value string) transfer {
	return tc.set("searchApi", value)
}

// SelectMode sets the selection mode
func (tc transfer) SelectMode(value string) transfer {
	return tc.set("selectMode", value)
}

// Sortable sets whether the control is sortable
func (tc transfer) Sortable(value bool) transfer {
	return tc.set("sortable", value)
}

// StaticMode sets whether the control is in static mode
func (tc transfer) StaticMode(value bool) transfer {
	return tc.set("staticMode", value)
}

// StaticClassName sets the CSS class name for static mode
func (tc transfer) StaticClassName(value string) transfer {
	return tc.set("staticClassName", value)
}

// StaticLabel sets the label for static mode
func (tc transfer) StaticLabel(value string) transfer {
	return tc.set("staticLabel", value)
}

// StaticPlaceholder sets the placeholder for static mode
func (tc transfer) StaticPlaceholder(value string) transfer {
	return tc.set("staticPlaceholder", value)
}

// StrictMode sets whether the control is in strict mode
func (tc transfer) StrictMode(value bool) transfer {
	return tc.set("strictMode", value)
}

// SubmitOnChange sets whether to submit the form on change
func (tc transfer) SubmitOnChange(value bool) transfer {
	return tc.set("submitOnChange", value)
}

// Sync sets whether to sync the value
func (tc transfer) Sync(value bool) transfer {
	return tc.set("sync", value)
}

// TextField sets the display field for options
func (tc transfer) TextField(value string) transfer {
	return tc.set("textField", value)
}

// Value sets the value
func (tc transfer) Value(value any) transfer {
	return tc.set("value", value)
}

// Visible sets whether the control is visible
func (tc transfer) Visible(value bool) transfer {
	return tc.set("visible", value)
}

// VisibleOn sets the expression for visibility
func (tc transfer) VisibleOn(value string) transfer {
	return tc.set("visibleOn", value)
}

// Width sets the control width
func (tc transfer) Width(value string) transfer {
	return tc.set("width", value)
}

// Wrappers sets the wrapper components
func (tc transfer) Wrappers(value string) transfer {
	return tc.set("wrappers", value)
}

// AllowMultiple sets whether multiple selection is allowed
func (tc transfer) AllowMultiple(value bool) transfer {
	return tc.set("allowMultiple", value)
}
