package comp

import "github.com/zrcoder/amisgo/model"

// UserSelect represents a mobile user selector component.
type UserSelect model.Schema

func NewUsersSelect() UserSelect {
	return UserSelect{"type": "users-select"}
}

func (usc UserSelect) set(key string, value any) UserSelect {
	usc[key] = value
	return usc
}

// AddApi sets the API for adding users.
func (usc UserSelect) AddApi(value string) UserSelect {
	return usc.set("addApi", value)
}

// AddControls sets the form controls for adding users.
func (usc UserSelect) AddControls(value string) UserSelect {
	return usc.set("addControls", value)
}

// AddDialog sets the dialog options for adding users.
func (usc UserSelect) AddDialog(value string) UserSelect {
	return usc.set("addDialog", value)
}

// AutoFill sets the auto-fill options.
func (usc UserSelect) AutoFill(value string) UserSelect {
	return usc.set("autoFill", value)
}

// ClassName sets the CSS class name.
func (usc UserSelect) ClassName(value string) UserSelect {
	return usc.set("className", value)
}

// ClearValueOnHidden sets whether to clear the value when hidden.
func (usc UserSelect) ClearValueOnHidden(value bool) UserSelect {
	return usc.set("clearValueOnHidden", value)
}

// Clearable sets whether the field is clearable.
func (usc UserSelect) Clearable(value bool) UserSelect {
	return usc.set("clearable", value)
}

// Creatable sets whether new users can be created.
func (usc UserSelect) Creatable(value bool) UserSelect {
	return usc.set("creatable", value)
}

// CreateBtnLabel sets the label for the create button.
func (usc UserSelect) CreateBtnLabel(value string) UserSelect {
	return usc.set("createBtnLabel", value)
}

// DeferApi sets the API for deferred loading.
func (usc UserSelect) DeferApi(value string) UserSelect {
	return usc.set("deferApi", value)
}

// DeferField sets the field for lazy loading.
func (usc UserSelect) DeferField(value string) UserSelect {
	return usc.set("deferField", value)
}

// DeleteApi sets the API for deleting users.
func (usc UserSelect) DeleteApi(value string) UserSelect {
	return usc.set("deleteApi", value)
}

// DeleteConfirmText sets the confirmation text for deletion.
func (usc UserSelect) DeleteConfirmText(value string) UserSelect {
	return usc.set("deleteConfirmText", value)
}

// Delimiter sets the delimiter.
func (usc UserSelect) Delimiter(value string) UserSelect {
	return usc.set("delimiter", value)
}

// Desc sets the description.
func (usc UserSelect) Desc(value string) UserSelect {
	return usc.set("desc", value)
}

// Description sets the HTML description.
func (usc UserSelect) Description(value string) UserSelect {
	return usc.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description.
func (usc UserSelect) DescriptionClassName(value string) UserSelect {
	return usc.set("descriptionClassName", value)
}

// Disabled sets whether the field is disabled.
func (usc UserSelect) Disabled(value bool) UserSelect {
	return usc.set("disabled", value)
}

// DisabledOn sets the expression for disabling the field.
func (usc UserSelect) DisabledOn(value string) UserSelect {
	return usc.set("disabledOn", value)
}

// EditApi sets the API for editing users.
func (usc UserSelect) EditApi(value string) UserSelect {
	return usc.set("editApi", value)
}

// EditControls sets the form controls for editing users.
func (usc UserSelect) EditControls(value string) UserSelect {
	return usc.set("editControls", value)
}

// EditDialog sets the dialog options for editing users.
func (usc UserSelect) EditDialog(value string) UserSelect {
	return usc.set("editDialog", value)
}

// Editable sets whether the field is editable.
func (usc UserSelect) Editable(value bool) UserSelect {
	return usc.set("editable", value)
}

// EditorSetting sets the editor configuration.
func (usc UserSelect) EditorSetting(value string) UserSelect {
	return usc.set("editorSetting", value)
}

// ExtraName sets the extra field name.
func (usc UserSelect) ExtraName(value string) UserSelect {
	return usc.set("extraName", value)
}

// ExtractValue sets whether to extract the value as an array.
func (usc UserSelect) ExtractValue(value bool) UserSelect {
	return usc.set("extractValue", value)
}

// Hidden sets whether the field is hidden.
func (usc UserSelect) Hidden(value bool) UserSelect {
	return usc.set("hidden", value)
}

// HiddenOn sets the expression for hiding the field.
func (usc UserSelect) HiddenOn(value string) UserSelect {
	return usc.set("hiddenOn", value)
}

// Hint sets the input hint.
func (usc UserSelect) Hint(value string) UserSelect {
	return usc.set("hint", value)
}

// Horizontal sets the horizontal layout configuration.
func (usc UserSelect) Horizontal(value string) UserSelect {
	return usc.set("horizontal", value)
}

// ID sets the unique component ID.
func (usc UserSelect) ID(value string) UserSelect {
	return usc.set("id", value)
}

// InitAutoFill sets the initial auto-fill options.
func (usc UserSelect) InitAutoFill(value string) UserSelect {
	return usc.set("initAutoFill", value)
}

// InitFetch sets whether to fetch data initially.
func (usc UserSelect) InitFetch(value bool) UserSelect {
	return usc.set("initFetch", value)
}

// InitFetchOn sets the expression for initial data fetch.
func (usc UserSelect) InitFetchOn(value string) UserSelect {
	return usc.set("initFetchOn", value)
}

// Inline sets whether the control is inline.
func (usc UserSelect) Inline(value bool) UserSelect {
	return usc.set("inline", value)
}

// InputClassName sets the CSS class name for the input.
func (usc UserSelect) InputClassName(value string) UserSelect {
	return usc.set("inputClassName", value)
}

// JoinValues sets whether to join values in single or multiple mode.
func (usc UserSelect) JoinValues(value bool) UserSelect {
	return usc.set("joinValues", value)
}

// Label sets the label.
func (usc UserSelect) Label(value string) UserSelect {
	return usc.set("label", value)
}

// LabelAlign sets the label alignment.
func (usc UserSelect) LabelAlign(value string) UserSelect {
	return usc.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label.
func (usc UserSelect) LabelClassName(value string) UserSelect {
	return usc.set("labelClassName", value)
}

// LabelRemark sets the label remark.
func (usc UserSelect) LabelRemark(value string) UserSelect {
	return usc.set("labelRemark", value)
}

// LabelWidth sets the custom label width.
func (usc UserSelect) LabelWidth(value string) UserSelect {
	return usc.set("labelWidth", value)
}

// Mode sets the display mode.
func (usc UserSelect) Mode(value string) UserSelect {
	return usc.set("mode", value)
}

// Multiple sets whether the field is in multiple selection mode.
func (usc UserSelect) Multiple(value bool) UserSelect {
	return usc.set("multiple", value)
}

// Name sets the field name.
func (usc UserSelect) Name(value string) UserSelect {
	return usc.set("name", value)
}

// OnEvent sets the event actions.
func (usc UserSelect) OnEvent(value any) UserSelect {
	return usc.set("onEvent", value)
}

// Options sets the options.
func (usc UserSelect) Options(value ...any) UserSelect {
	return usc.set("options", value)
}

// Placeholder sets the placeholder text.
func (usc UserSelect) Placeholder(value string) UserSelect {
	return usc.set("placeholder", value)
}

// ReadOnly sets whether the field is read-only.
func (usc UserSelect) ReadOnly(value bool) UserSelect {
	return usc.set("readOnly", value)
}

// ReadOnlyOn sets the expression for read-only condition.
func (usc UserSelect) ReadOnlyOn(value string) UserSelect {
	return usc.set("readOnlyOn", value)
}

// Remark sets the remark.
func (usc UserSelect) Remark(value string) UserSelect {
	return usc.set("remark", value)
}

// Removable sets whether the field is removable.
func (usc UserSelect) Removable(value bool) UserSelect {
	return usc.set("removable", value)
}

// Required sets whether the field is required.
func (usc UserSelect) Required(value bool) UserSelect {
	return usc.set("required", value)
}

// ResetValue sets the reset value.
func (usc UserSelect) ResetValue(value string) UserSelect {
	return usc.set("resetValue", value)
}

// ShowErrorMsg sets whether to show error messages.
func (usc UserSelect) ShowErrorMsg(value bool) UserSelect {
	return usc.set("showErrorMsg", value)
}

// Size sets the input size.
func (usc UserSelect) Size(value string) UserSelect {
	return usc.set("size", value)
}

// StaticClassName sets the CSS class name for static content.
func (usc UserSelect) StaticClassName(value string) UserSelect {
	return usc.set("staticClassName", value)
}

// StaticLabel sets the label for static display.
func (usc UserSelect) StaticLabel(value string) UserSelect {
	return usc.set("staticLabel", value)
}

// StaticPlaceholder sets the placeholder for static display.
func (usc UserSelect) StaticPlaceholder(value string) UserSelect {
	return usc.set("staticPlaceholder", value)
}

// StaticOptions sets the static options.
func (usc UserSelect) StaticOptions(value ...any) UserSelect {
	return usc.set("staticOptions", value)
}

// Width sets the input width.
func (usc UserSelect) Width(value string) UserSelect {
	return usc.set("width", value)
}
