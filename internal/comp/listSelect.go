package comp

import "github.com/zrcoder/amisgo/model"

// ListSelect represents a list control component
type ListSelect model.Schema

func NewListSelect() ListSelect {
	return ListSelect{"type": "list-select"}
}

// set sets a field value
func (lc ListSelect) set(key string, value any) ListSelect {
	lc[key] = value
	return lc
}

// ActiveItemSchema sets the active item schema
func (lc ListSelect) ActiveItemSchema(value string) ListSelect {
	return lc.set("activeItemSchema", value)
}

// AddApi sets the API for adding items
func (lc ListSelect) AddApi(value string) ListSelect {
	return lc.set("addApi", value)
}

// AddControls sets the form items for adding
func (lc ListSelect) AddControls(value string) ListSelect {
	return lc.set("addControls", value)
}

// AddDialog sets the add dialog options
func (lc ListSelect) AddDialog(value string) ListSelect {
	return lc.set("addDialog", value)
}

// AutoFill sets auto-fill options
func (lc ListSelect) AutoFill(value string) ListSelect {
	return lc.set("autoFill", value)
}

// ClassName sets the container CSS class name
func (lc ListSelect) ClassName(value string) ListSelect {
	return lc.set("className", value)
}

// ClearValueOnHidden sets whether to clear value when hidden
func (lc ListSelect) ClearValueOnHidden(value bool) ListSelect {
	return lc.set("clearValueOnHidden", value)
}

// Clearable sets whether the field is clearable
func (lc ListSelect) Clearable(value bool) ListSelect {
	return lc.set("clearable", value)
}

// Creatable sets whether new items can be created
func (lc ListSelect) Creatable(value bool) ListSelect {
	return lc.set("creatable", value)
}

// CreateBtnLabel sets the label for the create button
func (lc ListSelect) CreateBtnLabel(value string) ListSelect {
	return lc.set("createBtnLabel", value)
}

// DeferApi sets the API for deferred loading
func (lc ListSelect) DeferApi(value string) ListSelect {
	return lc.set("deferApi", value)
}

// DeferField sets the field for lazy loading
func (lc ListSelect) DeferField(value string) ListSelect {
	return lc.set("deferField", value)
}

// DeleteApi sets the API for deleting items
func (lc ListSelect) DeleteApi(value string) ListSelect {
	return lc.set("deleteApi", value)
}

// DeleteConfirmText sets the delete confirmation text
func (lc ListSelect) DeleteConfirmText(value string) ListSelect {
	return lc.set("deleteConfirmText", value)
}

// Delimiter sets the delimiter
func (lc ListSelect) Delimiter(value string) ListSelect {
	return lc.set("delimiter", value)
}

// Desc sets the description
func (lc ListSelect) Desc(value string) ListSelect {
	return lc.set("desc", value)
}

// Description sets the HTML description
func (lc ListSelect) Description(value string) ListSelect {
	return lc.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description
func (lc ListSelect) DescriptionClassName(value string) ListSelect {
	return lc.set("descriptionClassName", value)
}

// Disabled sets whether the field is disabled
func (lc ListSelect) Disabled(value bool) ListSelect {
	return lc.set("disabled", value)
}

// DisabledOn sets the expression for disabling the field
func (lc ListSelect) DisabledOn(value string) ListSelect {
	return lc.set("disabledOn", value)
}

// EditApi sets the API for editing items
func (lc ListSelect) EditApi(value string) ListSelect {
	return lc.set("editApi", value)
}

// EditControls sets the form items for editing
func (lc ListSelect) EditControls(value string) ListSelect {
	return lc.set("editControls", value)
}

// EditDialog sets the edit dialog options
func (lc ListSelect) EditDialog(value string) ListSelect {
	return lc.set("editDialog", value)
}

// Editable sets whether the field is editable
func (lc ListSelect) Editable(value bool) ListSelect {
	return lc.set("editable", value)
}

// EditorSetting sets the editor configuration
func (lc ListSelect) EditorSetting(value string) ListSelect {
	return lc.set("editorSetting", value)
}

// ExtraName sets the extra field name
func (lc ListSelect) ExtraName(value string) ListSelect {
	return lc.set("extraName", value)
}

// ExtractValue sets whether to extract the value as an array
func (lc ListSelect) ExtractValue(value bool) ListSelect {
	return lc.set("extractValue", value)
}

// Hidden sets whether the field is hidden
func (lc ListSelect) Hidden(value bool) ListSelect {
	return lc.set("hidden", value)
}

// HiddenOn sets the expression for hiding the field
func (lc ListSelect) HiddenOn(value string) ListSelect {
	return lc.set("hiddenOn", value)
}

// Hint sets the input hint
func (lc ListSelect) Hint(value string) ListSelect {
	return lc.set("hint", value)
}

// Horizontal sets the horizontal layout configuration
func (lc ListSelect) Horizontal(value string) ListSelect {
	return lc.set("horizontal", value)
}

// ID sets the unique component ID
func (lc ListSelect) ID(value string) ListSelect {
	return lc.set("id", value)
}

// ImageClassName sets the CSS class name for the image div
func (lc ListSelect) ImageClassName(value string) ListSelect {
	return lc.set("imageClassName", value)
}

// InitAutoFill sets the initial auto-fill configuration
func (lc ListSelect) InitAutoFill(value string) ListSelect {
	return lc.set("initAutoFill", value)
}

// InitFetch sets whether to fetch the source initially
func (lc ListSelect) InitFetch(value bool) ListSelect {
	return lc.set("initFetch", value)
}

// InitFetchOn sets the expression for initial fetch
func (lc ListSelect) InitFetchOn(value string) ListSelect {
	return lc.set("initFetchOn", value)
}

// Inline sets whether the control is inline
func (lc ListSelect) Inline(value bool) ListSelect {
	return lc.set("inline", value)
}

// InputClassName sets the CSS class name for the input
func (lc ListSelect) InputClassName(value string) ListSelect {
	return lc.set("inputClassName", value)
}

// ItemSchema sets the custom display template
func (lc ListSelect) ItemSchema(value string) ListSelect {
	return lc.set("itemSchema", value)
}

// JoinValues sets whether to join values in single select mode
func (lc ListSelect) JoinValues(value bool) ListSelect {
	return lc.set("joinValues", value)
}

// Label sets the label
func (lc ListSelect) Label(value string) ListSelect {
	return lc.set("label", value)
}

// LabelAlign sets the label alignment
func (lc ListSelect) LabelAlign(value string) ListSelect {
	return lc.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label
func (lc ListSelect) LabelClassName(value string) ListSelect {
	return lc.set("labelClassName", value)
}

// LabelRemark sets the label remark
func (lc ListSelect) LabelRemark(value string) ListSelect {
	return lc.set("labelRemark", value)
}

// LabelWidth sets the custom label width
func (lc ListSelect) LabelWidth(value string) ListSelect {
	return lc.set("labelWidth", value)
}

// ListClassName sets the CSS class name for the list div
func (lc ListSelect) ListClassName(value string) ListSelect {
	return lc.set("listClassName", value)
}

// Mode sets the display mode
func (lc ListSelect) Mode(value string) ListSelect {
	return lc.set("mode", value)
}

// Multiple sets whether the field is multi-select
func (lc ListSelect) Multiple(value bool) ListSelect {
	return lc.set("multiple", value)
}

// Name sets the field name for form submission
func (lc ListSelect) Name(value string) ListSelect {
	return lc.set("name", value)
}

// OnEvent sets the event actions
func (lc ListSelect) OnEvent(value any) ListSelect {
	return lc.set("onEvent", value)
}

// Options sets the options
func (lc ListSelect) Options(value ...any) ListSelect {
	return lc.set("options", value)
}

// Placeholder sets the placeholder text
func (lc ListSelect) Placeholder(value string) ListSelect {
	return lc.set("placeholder", value)
}

// ReadOnly sets whether the field is read-only
func (lc ListSelect) ReadOnly(value bool) ListSelect {
	return lc.set("readOnly", value)
}

// ReadOnlyOn sets the expression for read-only
func (lc ListSelect) ReadOnlyOn(value string) ListSelect {
	return lc.set("readOnlyOn", value)
}

// Remark sets the remark
func (lc ListSelect) Remark(value string) ListSelect {
	return lc.set("remark", value)
}

// Removable sets whether the field is removable
func (lc ListSelect) Removable(value bool) ListSelect {
	return lc.set("removable", value)
}

// Required sets whether the field is required
func (lc ListSelect) Required(value bool) ListSelect {
	return lc.set("required", value)
}

// ResetValueOnHidden sets whether to reset value when hidden
func (lc ListSelect) ResetValueOnHidden(value bool) ListSelect {
	return lc.set("resetValueOnHidden", value)
}

// Size sets the size
func (lc ListSelect) Size(value string) ListSelect {
	return lc.set("size", value)
}

// Source sets the data source API
func (lc ListSelect) Source(value string) ListSelect {
	return lc.set("source", value)
}

// Value sets the initial value
func (lc ListSelect) Value(value string) ListSelect {
	return lc.set("value", value)
}

// ValueField sets the value field name
func (lc ListSelect) ValueField(value string) ListSelect {
	return lc.set("valueField", value)
}

// Visible sets whether the field is visible
func (lc ListSelect) Visible(value bool) ListSelect {
	return lc.set("visible", value)
}

// VisibleOn sets the expression for visibility
func (lc ListSelect) VisibleOn(value string) ListSelect {
	return lc.set("visibleOn", value)
}

// Width sets the width
func (lc ListSelect) Width(value string) ListSelect {
	return lc.set("width", value)
}
