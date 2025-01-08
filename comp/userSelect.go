package comp

import "github.com/zrcoder/amisgo/model"

// userSelect represents a mobile user selector component.
type userSelect model.Schema

func UsersSelect() userSelect {
	return userSelect{}.set("type", "users-select")
}

func (usc userSelect) set(key string, value any) userSelect {
	usc[key] = value
	return usc
}

// AddApi sets the API for adding users.
func (usc userSelect) AddApi(value string) userSelect {
	return usc.set("addApi", value)
}

// AddControls sets the form controls for adding users.
func (usc userSelect) AddControls(value string) userSelect {
	return usc.set("addControls", value)
}

// AddDialog sets the dialog options for adding users.
func (usc userSelect) AddDialog(value string) userSelect {
	return usc.set("addDialog", value)
}

// AutoFill sets the auto-fill options.
func (usc userSelect) AutoFill(value string) userSelect {
	return usc.set("autoFill", value)
}

// ClassName sets the CSS class name.
func (usc userSelect) ClassName(value string) userSelect {
	return usc.set("className", value)
}

// ClearValueOnHidden sets whether to clear the value when hidden.
func (usc userSelect) ClearValueOnHidden(value bool) userSelect {
	return usc.set("clearValueOnHidden", value)
}

// Clearable sets whether the field is clearable.
func (usc userSelect) Clearable(value bool) userSelect {
	return usc.set("clearable", value)
}

// Creatable sets whether new users can be created.
func (usc userSelect) Creatable(value bool) userSelect {
	return usc.set("creatable", value)
}

// CreateBtnLabel sets the label for the create button.
func (usc userSelect) CreateBtnLabel(value string) userSelect {
	return usc.set("createBtnLabel", value)
}

// DeferApi sets the API for deferred loading.
func (usc userSelect) DeferApi(value string) userSelect {
	return usc.set("deferApi", value)
}

// DeferField sets the field for lazy loading.
func (usc userSelect) DeferField(value string) userSelect {
	return usc.set("deferField", value)
}

// DeleteApi sets the API for deleting users.
func (usc userSelect) DeleteApi(value string) userSelect {
	return usc.set("deleteApi", value)
}

// DeleteConfirmText sets the confirmation text for deletion.
func (usc userSelect) DeleteConfirmText(value string) userSelect {
	return usc.set("deleteConfirmText", value)
}

// Delimiter sets the delimiter.
func (usc userSelect) Delimiter(value string) userSelect {
	return usc.set("delimiter", value)
}

// Desc sets the description.
func (usc userSelect) Desc(value string) userSelect {
	return usc.set("desc", value)
}

// Description sets the HTML description.
func (usc userSelect) Description(value string) userSelect {
	return usc.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description.
func (usc userSelect) DescriptionClassName(value string) userSelect {
	return usc.set("descriptionClassName", value)
}

// Disabled sets whether the field is disabled.
func (usc userSelect) Disabled(value bool) userSelect {
	return usc.set("disabled", value)
}

// DisabledOn sets the expression for disabling the field.
func (usc userSelect) DisabledOn(value string) userSelect {
	return usc.set("disabledOn", value)
}

// EditApi sets the API for editing users.
func (usc userSelect) EditApi(value string) userSelect {
	return usc.set("editApi", value)
}

// EditControls sets the form controls for editing users.
func (usc userSelect) EditControls(value string) userSelect {
	return usc.set("editControls", value)
}

// EditDialog sets the dialog options for editing users.
func (usc userSelect) EditDialog(value string) userSelect {
	return usc.set("editDialog", value)
}

// Editable sets whether the field is editable.
func (usc userSelect) Editable(value bool) userSelect {
	return usc.set("editable", value)
}

// EditorSetting sets the editor configuration.
func (usc userSelect) EditorSetting(value string) userSelect {
	return usc.set("editorSetting", value)
}

// ExtraName sets the extra field name.
func (usc userSelect) ExtraName(value string) userSelect {
	return usc.set("extraName", value)
}

// ExtractValue sets whether to extract the value as an array.
func (usc userSelect) ExtractValue(value bool) userSelect {
	return usc.set("extractValue", value)
}

// Hidden sets whether the field is hidden.
func (usc userSelect) Hidden(value bool) userSelect {
	return usc.set("hidden", value)
}

// HiddenOn sets the expression for hiding the field.
func (usc userSelect) HiddenOn(value string) userSelect {
	return usc.set("hiddenOn", value)
}

// Hint sets the input hint.
func (usc userSelect) Hint(value string) userSelect {
	return usc.set("hint", value)
}

// Horizontal sets the horizontal layout configuration.
func (usc userSelect) Horizontal(value string) userSelect {
	return usc.set("horizontal", value)
}

// ID sets the unique component ID.
func (usc userSelect) ID(value string) userSelect {
	return usc.set("id", value)
}

// InitAutoFill sets the initial auto-fill options.
func (usc userSelect) InitAutoFill(value string) userSelect {
	return usc.set("initAutoFill", value)
}

// InitFetch sets whether to fetch data initially.
func (usc userSelect) InitFetch(value bool) userSelect {
	return usc.set("initFetch", value)
}

// InitFetchOn sets the expression for initial data fetch.
func (usc userSelect) InitFetchOn(value string) userSelect {
	return usc.set("initFetchOn", value)
}

// Inline sets whether the control is inline.
func (usc userSelect) Inline(value bool) userSelect {
	return usc.set("inline", value)
}

// InputClassName sets the CSS class name for the input.
func (usc userSelect) InputClassName(value string) userSelect {
	return usc.set("inputClassName", value)
}

// JoinValues sets whether to join values in single or multiple mode.
func (usc userSelect) JoinValues(value bool) userSelect {
	return usc.set("joinValues", value)
}

// Label sets the label.
func (usc userSelect) Label(value string) userSelect {
	return usc.set("label", value)
}

// LabelAlign sets the label alignment.
func (usc userSelect) LabelAlign(value string) userSelect {
	return usc.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label.
func (usc userSelect) LabelClassName(value string) userSelect {
	return usc.set("labelClassName", value)
}

// LabelRemark sets the label remark.
func (usc userSelect) LabelRemark(value string) userSelect {
	return usc.set("labelRemark", value)
}

// LabelWidth sets the custom label width.
func (usc userSelect) LabelWidth(value string) userSelect {
	return usc.set("labelWidth", value)
}

// Mode sets the display mode.
func (usc userSelect) Mode(value string) userSelect {
	return usc.set("mode", value)
}

// Multiple sets whether the field is in multiple selection mode.
func (usc userSelect) Multiple(value bool) userSelect {
	return usc.set("multiple", value)
}

// Name sets the field name.
func (usc userSelect) Name(value string) userSelect {
	return usc.set("name", value)
}

// OnEvent sets the event actions.
func (usc userSelect) OnEvent(value any) userSelect {
	return usc.set("onEvent", value)
}

// Options sets the options.
func (usc userSelect) Options(value ...any) userSelect {
	return usc.set("options", value)
}

// Placeholder sets the placeholder text.
func (usc userSelect) Placeholder(value string) userSelect {
	return usc.set("placeholder", value)
}

// ReadOnly sets whether the field is read-only.
func (usc userSelect) ReadOnly(value bool) userSelect {
	return usc.set("readOnly", value)
}

// ReadOnlyOn sets the expression for read-only condition.
func (usc userSelect) ReadOnlyOn(value string) userSelect {
	return usc.set("readOnlyOn", value)
}

// Remark sets the remark.
func (usc userSelect) Remark(value string) userSelect {
	return usc.set("remark", value)
}

// Removable sets whether the field is removable.
func (usc userSelect) Removable(value bool) userSelect {
	return usc.set("removable", value)
}

// Required sets whether the field is required.
func (usc userSelect) Required(value bool) userSelect {
	return usc.set("required", value)
}

// ResetValue sets the reset value.
func (usc userSelect) ResetValue(value string) userSelect {
	return usc.set("resetValue", value)
}

// ShowErrorMsg sets whether to show error messages.
func (usc userSelect) ShowErrorMsg(value bool) userSelect {
	return usc.set("showErrorMsg", value)
}

// Size sets the input size.
func (usc userSelect) Size(value string) userSelect {
	return usc.set("size", value)
}

// StaticClassName sets the CSS class name for static content.
func (usc userSelect) StaticClassName(value string) userSelect {
	return usc.set("staticClassName", value)
}

// StaticLabel sets the label for static display.
func (usc userSelect) StaticLabel(value string) userSelect {
	return usc.set("staticLabel", value)
}

// StaticPlaceholder sets the placeholder for static display.
func (usc userSelect) StaticPlaceholder(value string) userSelect {
	return usc.set("staticPlaceholder", value)
}

// StaticOptions sets the static options.
func (usc userSelect) StaticOptions(value ...any) userSelect {
	return usc.set("staticOptions", value)
}

// Width sets the input width.
func (usc userSelect) Width(value string) userSelect {
	return usc.set("width", value)
}
