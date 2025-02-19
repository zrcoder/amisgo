package comp

import "github.com/zrcoder/amisgo/schema"

// Radios single selection.
// Doc: https://aisuda.bce.baidu.com/amis/zh-CN/components/form/Radios
type Radios schema.Schema

func NewRadios() Radios {
	return Radios{"type": "radios"}
}

func (rc Radios) set(key string, value any) Radios {
	rc[key] = value
	return rc
}

// AddApi API to call when adding
func (rc Radios) AddApi(value string) Radios {
	return rc.set("addApi", value)
}

// AddControls form items when adding
func (rc Radios) AddControls(value string) Radios {
	return rc.set("addControls", value)
}

// AddDialog settings for adding dialog
func (rc Radios) AddDialog(value string) Radios {
	return rc.set("addDialog", value)
}

// AutoFill auto fill when selected
func (rc Radios) AutoFill(value string) Radios {
	return rc.set("autoFill", value)
}

// ClassName container CSS class name
func (rc Radios) ClassName(value string) Radios {
	return rc.set("className", value)
}

// ClearValueOnHidden whether to clear value when hidden
func (rc Radios) ClearValueOnHidden(value bool) Radios {
	return rc.set("clearValueOnHidden", value)
}

// Clearable whether to clear
func (rc Radios) Clearable(value bool) Radios {
	return rc.set("clearable", value)
}

// ColumnsCount show how many per row
func (rc Radios) ColumnsCount(value string) Radios {
	return rc.set("columnsCount", value)
}

// Creatable whether to allow create
func (rc Radios) Creatable(value bool) Radios {
	return rc.set("creatable", value)
}

// CreateBtnLabel label for create button
func (rc Radios) CreateBtnLabel(value string) Radios {
	return rc.set("createBtnLabel", value)
}

// DeferApi API for lazy loading
func (rc Radios) DeferApi(value string) Radios {
	return rc.set("deferApi", value)
}

// DeferField field for lazy loading
func (rc Radios) DeferField(value string) Radios {
	return rc.set("deferField", value)
}

// DeleteApi API for deleting
func (rc Radios) DeleteApi(value string) Radios {
	return rc.set("deleteApi", value)
}

// DeleteConfirmText confirm text for deleting
func (rc Radios) DeleteConfirmText(value string) Radios {
	return rc.set("deleteConfirmText", value)
}

// Delimiter delimiter
func (rc Radios) Delimiter(value string) Radios {
	return rc.set("delimiter", value)
}

// Desc
func (rc Radios) Desc(value string) Radios {
	return rc.set("desc", value)
}

// Description description
func (rc Radios) Description(value string) Radios {
	return rc.set("description", value)
}

// DescriptionClassName className for description
func (rc Radios) DescriptionClassName(value string) Radios {
	return rc.set("descriptionClassName", value)
}

// Disabled whether to disable
func (rc Radios) Disabled(value bool) Radios {
	return rc.set("disabled", value)
}

// DisabledOn whether to disable by expression
func (rc Radios) DisabledOn(value string) Radios {
	return rc.set("disabledOn", value)
}

// EditApi API for editing
func (rc Radios) EditApi(value string) Radios {
	return rc.set("editApi", value)
}

// EditControls form items for editing
func (rc Radios) EditControls(value string) Radios {
	return rc.set("editControls", value)
}

// EditDialog settings for editing dialog
func (rc Radios) EditDialog(value string) Radios {
	return rc.set("editDialog", value)
}

// Editable whether to allow edit
func (rc Radios) Editable(value bool) Radios {
	return rc.set("editable", value)
}

// EditorSetting editor settings
func (rc Radios) EditorSetting(value string) Radios {
	return rc.set("editorSetting", value)
}

// ExtraName extra field name
func (rc Radios) ExtraName(value string) Radios {
	return rc.set("extraName", value)
}

// ExtractValue whether to extract value
func (rc Radios) ExtractValue(value bool) Radios {
	return rc.set("extractValue", value)
}

// Hidden whether to hide
func (rc Radios) Hidden(value bool) Radios {
	return rc.set("hidden", value)
}

// HiddenOn whether to hide by expression
func (rc Radios) HiddenOn(value string) Radios {
	return rc.set("hiddenOn", value)
}

// Hint hint
func (rc Radios) Hint(value string) Radios {
	return rc.set("hint", value)
}

// Horizontal horizontal layout
func (rc Radios) Horizontal(value string) Radios {
	return rc.set("horizontal", value)
}

// ID component unique id
func (rc Radios) ID(value string) Radios {
	return rc.set("id", value)
}

// InitAutoFill
func (rc Radios) InitAutoFill(value string) Radios {
	return rc.set("initAutoFill", value)
}

// InitFetch whether to fetch data when init
func (rc Radios) InitFetch(value bool) Radios {
	return rc.set("initFetch", value)
}

// InitFetchOn whether to fetch data when init by expression
func (rc Radios) InitFetchOn(value string) Radios {
	return rc.set("initFetchOn", value)
}

// Inline whether to inline
func (rc Radios) Inline(value bool) Radios {
	return rc.set("inline", value)
}

// InputClassName className for input
func (rc Radios) InputClassName(value string) Radios {
	return rc.set("inputClassName", value)
}

// JoinValues whether to join values
func (rc Radios) JoinValues(value bool) Radios {
	return rc.set("joinValues", value)
}

// Label label
func (rc Radios) Label(value string) Radios {
	return rc.set("label", value)
}

// LabelAlign label align
func (rc Radios) LabelAlign(value string) Radios {
	return rc.set("labelAlign", value)
}

// LabelClassName className for label
func (rc Radios) LabelClassName(value string) Radios {
	return rc.set("labelClassName", value)
}

// LabelRemark label remark
func (rc Radios) LabelRemark(value string) Radios {
	return rc.set("labelRemark", value)
}

// LabelWidth label width
func (rc Radios) LabelWidth(value string) Radios {
	return rc.set("labelWidth", value)
}

// Mode mode
func (rc Radios) Mode(value string) Radios {
	return rc.set("mode", value)
}

// Multiple whether to multiple
func (rc Radios) Multiple(value bool) Radios {
	return rc.set("multiple", value)
}

// Name field name
func (rc Radios) Name(value string) Radios {
	return rc.set("name", value)
}

// OnEvent event action configuration
func (rc Radios) OnEvent(value any) Radios {
	return rc.set("onEvent", value)
}

// Option option
func (rc Radios) Option(value any) Radios {
	return rc.set("option", value)
}

// Options options
func (rc Radios) Options(value ...any) Radios {
	return rc.set("options", value)
}

// Placeholder placeholder
func (rc Radios) Placeholder(value string) Radios {
	return rc.set("placeholder", value)
}

// ReadOnly whether to read only
func (rc Radios) ReadOnly(value bool) Radios {
	return rc.set("readOnly", value)
}

// Removable whether to removable
func (rc Radios) Removable(value bool) Radios {
	return rc.set("removable", value)
}

// RenderMode render mode
func (rc Radios) RenderMode(value string) Radios {
	return rc.set("renderMode", value)
}

// Required whether to required
func (rc Radios) Required(value bool) Radios {
	return rc.set("required", value)
}

// ShowIcon whether to show icon
func (rc Radios) ShowIcon(value bool) Radios {
	return rc.set("showIcon", value)
}

// ShowRadio whether to show radio
func (rc Radios) ShowRadio(value bool) Radios {
	return rc.set("showRadio", value)
}

// ShowText whether to show text
func (rc Radios) ShowText(value bool) Radios {
	return rc.set("showText", value)
}

// Size size
func (rc Radios) Size(value string) Radios {
	return rc.set("size", value)
}

// Source data source
func (rc Radios) Source(value string) Radios {
	return rc.set("source", value)
}

// SourceApi custom source API
func (rc Radios) SourceApi(value string) Radios {
	return rc.set("sourceApi", value)
}

// Toggle whether to toggle
func (rc Radios) Toggle(value bool) Radios {
	return rc.set("toggle", value)
}

// Value value
func (rc Radios) Value(value any) Radios {
	return rc.set("value", value)
}

// Visible whether to visible
func (rc Radios) Visible(value bool) Radios {
	return rc.set("visible", value)
}

// VisibleOn whether to visible by expression
func (rc Radios) VisibleOn(value string) Radios {
	return rc.set("visibleOn", value)
}

// Width width
func (rc Radios) Width(value string) Radios {
	return rc.set("width", value)
}

// WrapperClassName className for wrapper
func (rc Radios) WrapperClassName(value string) Radios {
	return rc.set("wrapperClassName", value)
}

// ValueField value field
func (rc Radios) ValueField(value string) Radios {
	return rc.set("valueField", value)
}

// LabelField label field
func (rc Radios) LabelField(value string) Radios {
	return rc.set("labelField", value)
}

// DescriptionField description field
func (rc Radios) DescriptionField(value string) Radios {
	return rc.set("descriptionField", value)
}

// OptionValueField option value field
func (rc Radios) OptionValueField(value string) Radios {
	return rc.set("optionValueField", value)
}

// Option label field
func (rc Radios) OptionLabelField(value string) Radios {
	return rc.set("optionLabelField", value)
}

// Option description field
func (rc Radios) OptionDescriptionField(value string) Radios {
	return rc.set("optionDescriptionField", value)
}
