package comp

import "github.com/zrcoder/amisgo/model"

// listSelect represents a list control component
type listSelect model.Schema

func ListSelect() listSelect {
	return make(listSelect).set("type", "list-select")
}

// set sets a field value
func (lc listSelect) set(key string, value any) listSelect {
	lc[key] = value
	return lc
}

// ActiveItemSchema sets the active item schema
func (lc listSelect) ActiveItemSchema(value string) listSelect {
	return lc.set("activeItemSchema", value)
}

// AddApi sets the API for adding items
func (lc listSelect) AddApi(value string) listSelect {
	return lc.set("addApi", value)
}

// AddControls sets the form items for adding
func (lc listSelect) AddControls(value string) listSelect {
	return lc.set("addControls", value)
}

// AddDialog sets the add dialog options
func (lc listSelect) AddDialog(value string) listSelect {
	return lc.set("addDialog", value)
}

// AutoFill sets auto-fill options
func (lc listSelect) AutoFill(value string) listSelect {
	return lc.set("autoFill", value)
}

// ClassName sets the container CSS class name
func (lc listSelect) ClassName(value string) listSelect {
	return lc.set("className", value)
}

// ClearValueOnHidden sets whether to clear value when hidden
func (lc listSelect) ClearValueOnHidden(value bool) listSelect {
	return lc.set("clearValueOnHidden", value)
}

// Clearable sets whether the field is clearable
func (lc listSelect) Clearable(value bool) listSelect {
	return lc.set("clearable", value)
}

// Creatable sets whether new items can be created
func (lc listSelect) Creatable(value bool) listSelect {
	return lc.set("creatable", value)
}

// CreateBtnLabel sets the label for the create button
func (lc listSelect) CreateBtnLabel(value string) listSelect {
	return lc.set("createBtnLabel", value)
}

// DeferApi sets the API for deferred loading
func (lc listSelect) DeferApi(value string) listSelect {
	return lc.set("deferApi", value)
}

// DeferField sets the field for lazy loading
func (lc listSelect) DeferField(value string) listSelect {
	return lc.set("deferField", value)
}

// DeleteApi sets the API for deleting items
func (lc listSelect) DeleteApi(value string) listSelect {
	return lc.set("deleteApi", value)
}

// DeleteConfirmText sets the delete confirmation text
func (lc listSelect) DeleteConfirmText(value string) listSelect {
	return lc.set("deleteConfirmText", value)
}

// Delimiter sets the delimiter
func (lc listSelect) Delimiter(value string) listSelect {
	return lc.set("delimiter", value)
}

// Desc sets the description
func (lc listSelect) Desc(value string) listSelect {
	return lc.set("desc", value)
}

// Description sets the HTML description
func (lc listSelect) Description(value string) listSelect {
	return lc.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description
func (lc listSelect) DescriptionClassName(value string) listSelect {
	return lc.set("descriptionClassName", value)
}

// Disabled sets whether the field is disabled
func (lc listSelect) Disabled(value bool) listSelect {
	return lc.set("disabled", value)
}

// DisabledOn sets the expression for disabling the field
func (lc listSelect) DisabledOn(value string) listSelect {
	return lc.set("disabledOn", value)
}

// EditApi sets the API for editing items
func (lc listSelect) EditApi(value string) listSelect {
	return lc.set("editApi", value)
}

// EditControls sets the form items for editing
func (lc listSelect) EditControls(value string) listSelect {
	return lc.set("editControls", value)
}

// EditDialog sets the edit dialog options
func (lc listSelect) EditDialog(value string) listSelect {
	return lc.set("editDialog", value)
}

// Editable sets whether the field is editable
func (lc listSelect) Editable(value bool) listSelect {
	return lc.set("editable", value)
}

// EditorSetting sets the editor configuration
func (lc listSelect) EditorSetting(value string) listSelect {
	return lc.set("editorSetting", value)
}

// ExtraName sets the extra field name
func (lc listSelect) ExtraName(value string) listSelect {
	return lc.set("extraName", value)
}

// ExtractValue sets whether to extract the value as an array
func (lc listSelect) ExtractValue(value bool) listSelect {
	return lc.set("extractValue", value)
}

// Hidden sets whether the field is hidden
func (lc listSelect) Hidden(value bool) listSelect {
	return lc.set("hidden", value)
}

// HiddenOn sets the expression for hiding the field
func (lc listSelect) HiddenOn(value string) listSelect {
	return lc.set("hiddenOn", value)
}

// Hint sets the input hint
func (lc listSelect) Hint(value string) listSelect {
	return lc.set("hint", value)
}

// Horizontal sets the horizontal layout configuration
func (lc listSelect) Horizontal(value string) listSelect {
	return lc.set("horizontal", value)
}

// ID sets the unique component ID
func (lc listSelect) ID(value string) listSelect {
	return lc.set("id", value)
}

// ImageClassName sets the CSS class name for the image div
func (lc listSelect) ImageClassName(value string) listSelect {
	return lc.set("imageClassName", value)
}

// InitAutoFill sets the initial auto-fill configuration
func (lc listSelect) InitAutoFill(value string) listSelect {
	return lc.set("initAutoFill", value)
}

// InitFetch sets whether to fetch the source initially
func (lc listSelect) InitFetch(value bool) listSelect {
	return lc.set("initFetch", value)
}

// InitFetchOn sets the expression for initial fetch
func (lc listSelect) InitFetchOn(value string) listSelect {
	return lc.set("initFetchOn", value)
}

// Inline sets whether the control is inline
func (lc listSelect) Inline(value bool) listSelect {
	return lc.set("inline", value)
}

// InputClassName sets the CSS class name for the input
func (lc listSelect) InputClassName(value string) listSelect {
	return lc.set("inputClassName", value)
}

// ItemSchema sets the custom display template
func (lc listSelect) ItemSchema(value string) listSelect {
	return lc.set("itemSchema", value)
}

// JoinValues sets whether to join values in single select mode
func (lc listSelect) JoinValues(value bool) listSelect {
	return lc.set("joinValues", value)
}

// Label sets the label
func (lc listSelect) Label(value string) listSelect {
	return lc.set("label", value)
}

// LabelAlign sets the label alignment
func (lc listSelect) LabelAlign(value string) listSelect {
	return lc.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label
func (lc listSelect) LabelClassName(value string) listSelect {
	return lc.set("labelClassName", value)
}

// LabelRemark sets the label remark
func (lc listSelect) LabelRemark(value string) listSelect {
	return lc.set("labelRemark", value)
}

// LabelWidth sets the custom label width
func (lc listSelect) LabelWidth(value string) listSelect {
	return lc.set("labelWidth", value)
}

// ListClassName sets the CSS class name for the list div
func (lc listSelect) ListClassName(value string) listSelect {
	return lc.set("listClassName", value)
}

// Mode sets the display mode
func (lc listSelect) Mode(value string) listSelect {
	return lc.set("mode", value)
}

// Multiple sets whether the field is multi-select
func (lc listSelect) Multiple(value bool) listSelect {
	return lc.set("multiple", value)
}

// Name sets the field name for form submission
func (lc listSelect) Name(value string) listSelect {
	return lc.set("name", value)
}

// OnEvent sets the event actions
func (lc listSelect) OnEvent(value any) listSelect {
	return lc.set("onEvent", value)
}

// Options sets the options
func (lc listSelect) Options(value ...any) listSelect {
	return lc.set("options", value)
}

// Placeholder sets the placeholder text
func (lc listSelect) Placeholder(value string) listSelect {
	return lc.set("placeholder", value)
}

// ReadOnly sets whether the field is read-only
func (lc listSelect) ReadOnly(value bool) listSelect {
	return lc.set("readOnly", value)
}

// ReadOnlyOn sets the expression for read-only
func (lc listSelect) ReadOnlyOn(value string) listSelect {
	return lc.set("readOnlyOn", value)
}

// Remark sets the remark
func (lc listSelect) Remark(value string) listSelect {
	return lc.set("remark", value)
}

// Removable sets whether the field is removable
func (lc listSelect) Removable(value bool) listSelect {
	return lc.set("removable", value)
}

// Required sets whether the field is required
func (lc listSelect) Required(value bool) listSelect {
	return lc.set("required", value)
}

// ResetValueOnHidden sets whether to reset value when hidden
func (lc listSelect) ResetValueOnHidden(value bool) listSelect {
	return lc.set("resetValueOnHidden", value)
}

// Size sets the size
func (lc listSelect) Size(value string) listSelect {
	return lc.set("size", value)
}

// Source sets the data source API
func (lc listSelect) Source(value string) listSelect {
	return lc.set("source", value)
}

// Value sets the initial value
func (lc listSelect) Value(value string) listSelect {
	return lc.set("value", value)
}

// ValueField sets the value field name
func (lc listSelect) ValueField(value string) listSelect {
	return lc.set("valueField", value)
}

// Visible sets whether the field is visible
func (lc listSelect) Visible(value bool) listSelect {
	return lc.set("visible", value)
}

// VisibleOn sets the expression for visibility
func (lc listSelect) VisibleOn(value string) listSelect {
	return lc.set("visibleOn", value)
}

// Width sets the width
func (lc listSelect) Width(value string) listSelect {
	return lc.set("width", value)
}
