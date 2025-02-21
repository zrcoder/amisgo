package comp

import "github.com/zrcoder/amisgo/schema"

// Transfer control definition
type Transfer schema.Schema

func NewTransfer() Transfer {
	return Transfer{"type": "transfer"}
}

func (tc Transfer) set(key string, value any) Transfer {
	tc[key] = value
	return tc
}

// AddApi sets the API for adding items
func (tc Transfer) AddApi(value string) Transfer {
	return tc.set("addApi", value)
}

// AddControls sets the form items for adding
func (tc Transfer) AddControls(value string) Transfer {
	return tc.set("addControls", value)
}

// AddDialog sets the dialog for adding items
func (tc Transfer) AddDialog(value string) Transfer {
	return tc.set("addDialog", value)
}

// AutoCheckChildren sets whether to auto-check children
func (tc Transfer) AutoCheckChildren(value bool) Transfer {
	return tc.set("autoCheckChildren", value)
}

// AutoFill sets the auto-fill value
func (tc Transfer) AutoFill(value string) Transfer {
	return tc.set("autoFill", value)
}

// ClassName sets the container CSS class name
func (tc Transfer) ClassName(value string) Transfer {
	return tc.set("className", value)
}

// ClearValueOnHidden sets whether to clear value when hidden
func (tc Transfer) ClearValueOnHidden(value bool) Transfer {
	return tc.set("clearValueOnHidden", value)
}

// Clearable sets whether the control is clearable
func (tc Transfer) Clearable(value bool) Transfer {
	return tc.set("clearable", value)
}

// Columns sets the columns for table mode
func (tc Transfer) Columns(value ...any) Transfer {
	return tc.set("columns", value)
}

// Creatable sets whether new items can be created
func (tc Transfer) Creatable(value bool) Transfer {
	return tc.set("creatable", value)
}

// CreateBtnLabel sets the label for the create button
func (tc Transfer) CreateBtnLabel(value string) Transfer {
	return tc.set("createBtnLabel", value)
}

// DeferApi sets the API for deferred loading
func (tc Transfer) DeferApi(value string) Transfer {
	return tc.set("deferApi", value)
}

// DeferField sets the field for deferred loading
func (tc Transfer) DeferField(value string) Transfer {
	return tc.set("deferField", value)
}

// DeleteApi sets the API for deleting items
func (tc Transfer) DeleteApi(value string) Transfer {
	return tc.set("deleteApi", value)
}

// DeleteConfirmText sets the confirmation text for deletion
func (tc Transfer) DeleteConfirmText(value string) Transfer {
	return tc.set("deleteConfirmText", value)
}

// Delimiter sets the delimiter
func (tc Transfer) Delimiter(value string) Transfer {
	return tc.set("delimiter", value)
}

// Desc sets the description
func (tc Transfer) Desc(value string) Transfer {
	return tc.set("desc", value)
}

// Description sets the description content
func (tc Transfer) Description(value string) Transfer {
	return tc.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description
func (tc Transfer) DescriptionClassName(value string) Transfer {
	return tc.set("descriptionClassName", value)
}

// Disabled sets whether the control is disabled
func (tc Transfer) Disabled(value bool) Transfer {
	return tc.set("disabled", value)
}

// DisabledOn sets the expression for disabling the control
func (tc Transfer) DisabledOn(value string) Transfer {
	return tc.set("disabledOn", value)
}

// EditApi sets the API for editing items
func (tc Transfer) EditApi(value string) Transfer {
	return tc.set("editApi", value)
}

// EditControls sets the form items for editing
func (tc Transfer) EditControls(value string) Transfer {
	return tc.set("editControls", value)
}

// EditDialog sets the dialog for editing items
func (tc Transfer) EditDialog(value string) Transfer {
	return tc.set("editDialog", value)
}

// Editable sets whether the control is editable
func (tc Transfer) Editable(value bool) Transfer {
	return tc.set("editable", value)
}

// EditorSetting sets the editor configuration
func (tc Transfer) EditorSetting(value string) Transfer {
	return tc.set("editorSetting", value)
}

// ExtraName sets the extra field name
func (tc Transfer) ExtraName(value string) Transfer {
	return tc.set("extraName", value)
}

// ExtractValue sets whether to extract the value as an array
func (tc Transfer) ExtractValue(value bool) Transfer {
	return tc.set("extractValue", value)
}

// Hidden sets whether the control is hidden
func (tc Transfer) Hidden(value bool) Transfer {
	return tc.set("hidden", value)
}

// HiddenOn sets the expression for hiding the control
func (tc Transfer) HiddenOn(value string) Transfer {
	return tc.set("hiddenOn", value)
}

// Hint sets the input hint
func (tc Transfer) Hint(value string) Transfer {
	return tc.set("hint", value)
}

// Horizontal sets the horizontal layout configuration
func (tc Transfer) Horizontal(value string) Transfer {
	return tc.set("horizontal", value)
}

// ID sets the unique ID for the component
func (tc Transfer) ID(value string) Transfer {
	return tc.set("id", value)
}

// InitAutoFill sets the initial auto-fill value
func (tc Transfer) InitAutoFill(value string) Transfer {
	return tc.set("initAutoFill", value)
}

// InitFetch sets whether to fetch data initially
func (tc Transfer) InitFetch(value bool) Transfer {
	return tc.set("initFetch", value)
}

// InitFetchOn sets the expression for initial data fetch
func (tc Transfer) InitFetchOn(value string) Transfer {
	return tc.set("initFetchOn", value)
}

// InitiallyOpen sets whether the control is initially open
func (tc Transfer) InitiallyOpen(value bool) Transfer {
	return tc.set("initiallyOpen", value)
}

// Inline sets whether the control is in inline mode
func (tc Transfer) Inline(value bool) Transfer {
	return tc.set("inline", value)
}

// InputClassName sets the CSS class name for the input
func (tc Transfer) InputClassName(value string) Transfer {
	return tc.set("inputClassName", value)
}

// ItemHeight sets the height of each item
func (tc Transfer) ItemHeight(value string) Transfer {
	return tc.set("itemHeight", value)
}

// JoinValues sets whether to join values with a delimiter
func (tc Transfer) JoinValues(value bool) Transfer {
	return tc.set("joinValues", value)
}

// Label sets the label
func (tc Transfer) Label(value string) Transfer {
	return tc.set("label", value)
}

// LabelAlign sets the label alignment
func (tc Transfer) LabelAlign(value string) Transfer {
	return tc.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label
func (tc Transfer) LabelClassName(value string) Transfer {
	return tc.set("labelClassName", value)
}

// LabelRemark sets the remark for the label
func (tc Transfer) LabelRemark(value string) Transfer {
	return tc.set("labelRemark", value)
}

// LabelWidth sets the custom width for the label
func (tc Transfer) LabelWidth(value string) Transfer {
	return tc.set("labelWidth", value)
}

// LeftMode sets the selection mode for the left side
func (tc Transfer) LeftMode(value string) Transfer {
	return tc.set("leftMode", value)
}

// Max sets the maximum number of selectable items
func (tc Transfer) Max(value int) Transfer {
	return tc.set("max", value)
}

// Min sets the minimum number of selectable items
func (tc Transfer) Min(value int) Transfer {
	return tc.set("min", value)
}

// Name sets the form item name
func (tc Transfer) Name(value string) Transfer {
	return tc.set("name", value)
}

// OnChange sets the callback function for changes
func (tc Transfer) OnChange(value string) Transfer {
	return tc.set("onChange", value)
}

// OnEvent sets the event handler
func (tc Transfer) OnEvent(value Event) Transfer {
	return tc.set("onEvent", value)
}

// Option sets a single option
func (tc Transfer) Option(value string) Transfer {
	return tc.set("option", value)
}

// Options sets the options data
func (tc Transfer) Options(value ...any) Transfer {
	return tc.set("options", value)
}

// Placeholder sets the input placeholder
func (tc Transfer) Placeholder(value string) Transfer {
	return tc.set("placeholder", value)
}

// ReadOnly sets whether the control is read-only
func (tc Transfer) ReadOnly(value bool) Transfer {
	return tc.set("readOnly", value)
}

// RemoteOptions sets the remote options
func (tc Transfer) RemoteOptions(value ...any) Transfer {
	return tc.set("remoteOptions", value)
}

// Render sets the form render mode
func (tc Transfer) Render(value string) Transfer {
	return tc.set("render", value)
}

// ResetValue sets the default value when the form is reset
func (tc Transfer) ResetValue(value string) Transfer {
	return tc.set("resetValue", value)
}

// SaveImmediately sets whether to save immediately on change
func (tc Transfer) SaveImmediately(value bool) Transfer {
	return tc.set("saveImmediately", value)
}

// Source sets the data source
func (tc Transfer) Source(value string) Transfer {
	return tc.set("source", value)
}

// Searchable sets whether the control is searchable
func (tc Transfer) Searchable(value bool) Transfer {
	return tc.set("searchable", value)
}

// SearchApi sets the API for searching
func (tc Transfer) SearchApi(value string) Transfer {
	return tc.set("searchApi", value)
}

// SelectMode sets the selection mode
func (tc Transfer) SelectMode(value string) Transfer {
	return tc.set("selectMode", value)
}

// Sortable sets whether the control is sortable
func (tc Transfer) Sortable(value bool) Transfer {
	return tc.set("sortable", value)
}

// StaticMode sets whether the control is in static mode
func (tc Transfer) StaticMode(value bool) Transfer {
	return tc.set("staticMode", value)
}

// StaticClassName sets the CSS class name for static mode
func (tc Transfer) StaticClassName(value string) Transfer {
	return tc.set("staticClassName", value)
}

// StaticLabel sets the label for static mode
func (tc Transfer) StaticLabel(value string) Transfer {
	return tc.set("staticLabel", value)
}

// StaticPlaceholder sets the placeholder for static mode
func (tc Transfer) StaticPlaceholder(value string) Transfer {
	return tc.set("staticPlaceholder", value)
}

// StrictMode sets whether the control is in strict mode
func (tc Transfer) StrictMode(value bool) Transfer {
	return tc.set("strictMode", value)
}

// SubmitOnChange sets whether to submit the form on change
func (tc Transfer) SubmitOnChange(value bool) Transfer {
	return tc.set("submitOnChange", value)
}

// Sync sets whether to sync the value
func (tc Transfer) Sync(value bool) Transfer {
	return tc.set("sync", value)
}

// TextField sets the display field for options
func (tc Transfer) TextField(value string) Transfer {
	return tc.set("textField", value)
}

// Value sets the value
func (tc Transfer) Value(value any) Transfer {
	return tc.set("value", value)
}

// Visible sets whether the control is visible
func (tc Transfer) Visible(value bool) Transfer {
	return tc.set("visible", value)
}

// VisibleOn sets the expression for visibility
func (tc Transfer) VisibleOn(value string) Transfer {
	return tc.set("visibleOn", value)
}

// Width sets the control width
func (tc Transfer) Width(value string) Transfer {
	return tc.set("width", value)
}

// Wrappers sets the wrapper components
func (tc Transfer) Wrappers(value string) Transfer {
	return tc.set("wrappers", value)
}

// AllowMultiple sets whether multiple selection is allowed
func (tc Transfer) AllowMultiple(value bool) Transfer {
	return tc.set("allowMultiple", value)
}
