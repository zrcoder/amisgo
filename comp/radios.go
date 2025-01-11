package comp

import "github.com/zrcoder/amisgo/model"

// radios single selection.
// Doc: https://aisuda.bce.baidu.com/amis/zh-CN/components/form/radios

type radios model.Schema

// Radios create a new RadiosControl instance
func Radios() radios {
	return radios{"type": "radios"}
}

func (rc radios) set(key string, value any) radios {
	rc[key] = value
	return rc
}

// AddApi API to call when adding
func (rc radios) AddApi(value string) radios {
	return rc.set("addApi", value)
}

// AddControls form items when adding
func (rc radios) AddControls(value string) radios {
	return rc.set("addControls", value)
}

// AddDialog settings for adding dialog
func (rc radios) AddDialog(value string) radios {
	return rc.set("addDialog", value)
}

// AutoFill auto fill when selected
func (rc radios) AutoFill(value string) radios {
	return rc.set("autoFill", value)
}

// ClassName container CSS class name
func (rc radios) ClassName(value string) radios {
	return rc.set("className", value)
}

// ClearValueOnHidden whether to clear value when hidden
func (rc radios) ClearValueOnHidden(value bool) radios {
	return rc.set("clearValueOnHidden", value)
}

// Clearable whether to clear
func (rc radios) Clearable(value bool) radios {
	return rc.set("clearable", value)
}

// ColumnsCount show how many per row
func (rc radios) ColumnsCount(value string) radios {
	return rc.set("columnsCount", value)
}

// Creatable whether to allow create
func (rc radios) Creatable(value bool) radios {
	return rc.set("creatable", value)
}

// CreateBtnLabel label for create button
func (rc radios) CreateBtnLabel(value string) radios {
	return rc.set("createBtnLabel", value)
}

// DeferApi API for lazy loading
func (rc radios) DeferApi(value string) radios {
	return rc.set("deferApi", value)
}

// DeferField field for lazy loading
func (rc radios) DeferField(value string) radios {
	return rc.set("deferField", value)
}

// DeleteApi API for deleting
func (rc radios) DeleteApi(value string) radios {
	return rc.set("deleteApi", value)
}

// DeleteConfirmText confirm text for deleting
func (rc radios) DeleteConfirmText(value string) radios {
	return rc.set("deleteConfirmText", value)
}

// Delimiter delimiter
func (rc radios) Delimiter(value string) radios {
	return rc.set("delimiter", value)
}

// Desc
func (rc radios) Desc(value string) radios {
	return rc.set("desc", value)
}

// Description description
func (rc radios) Description(value string) radios {
	return rc.set("description", value)
}

// DescriptionClassName className for description
func (rc radios) DescriptionClassName(value string) radios {
	return rc.set("descriptionClassName", value)
}

// Disabled whether to disable
func (rc radios) Disabled(value bool) radios {
	return rc.set("disabled", value)
}

// DisabledOn whether to disable by expression
func (rc radios) DisabledOn(value string) radios {
	return rc.set("disabledOn", value)
}

// EditApi API for editing
func (rc radios) EditApi(value string) radios {
	return rc.set("editApi", value)
}

// EditControls form items for editing
func (rc radios) EditControls(value string) radios {
	return rc.set("editControls", value)
}

// EditDialog settings for editing dialog
func (rc radios) EditDialog(value string) radios {
	return rc.set("editDialog", value)
}

// Editable whether to allow edit
func (rc radios) Editable(value bool) radios {
	return rc.set("editable", value)
}

// EditorSetting editor settings
func (rc radios) EditorSetting(value string) radios {
	return rc.set("editorSetting", value)
}

// ExtraName extra field name
func (rc radios) ExtraName(value string) radios {
	return rc.set("extraName", value)
}

// ExtractValue whether to extract value
func (rc radios) ExtractValue(value bool) radios {
	return rc.set("extractValue", value)
}

// Hidden whether to hide
func (rc radios) Hidden(value bool) radios {
	return rc.set("hidden", value)
}

// HiddenOn whether to hide by expression
func (rc radios) HiddenOn(value string) radios {
	return rc.set("hiddenOn", value)
}

// Hint hint
func (rc radios) Hint(value string) radios {
	return rc.set("hint", value)
}

// Horizontal horizontal layout
func (rc radios) Horizontal(value string) radios {
	return rc.set("horizontal", value)
}

// ID component unique id
func (rc radios) ID(value string) radios {
	return rc.set("id", value)
}

// InitAutoFill
func (rc radios) InitAutoFill(value string) radios {
	return rc.set("initAutoFill", value)
}

// InitFetch whether to fetch data when init
func (rc radios) InitFetch(value bool) radios {
	return rc.set("initFetch", value)
}

// InitFetchOn whether to fetch data when init by expression
func (rc radios) InitFetchOn(value string) radios {
	return rc.set("initFetchOn", value)
}

// Inline whether to inline
func (rc radios) Inline(value bool) radios {
	return rc.set("inline", value)
}

// InputClassName className for input
func (rc radios) InputClassName(value string) radios {
	return rc.set("inputClassName", value)
}

// JoinValues whether to join values
func (rc radios) JoinValues(value bool) radios {
	return rc.set("joinValues", value)
}

// Label label
func (rc radios) Label(value string) radios {
	return rc.set("label", value)
}

// LabelAlign label align
func (rc radios) LabelAlign(value string) radios {
	return rc.set("labelAlign", value)
}

// LabelClassName className for label
func (rc radios) LabelClassName(value string) radios {
	return rc.set("labelClassName", value)
}

// LabelRemark label remark
func (rc radios) LabelRemark(value string) radios {
	return rc.set("labelRemark", value)
}

// LabelWidth label width
func (rc radios) LabelWidth(value string) radios {
	return rc.set("labelWidth", value)
}

// Mode mode
func (rc radios) Mode(value string) radios {
	return rc.set("mode", value)
}

// Multiple whether to multiple
func (rc radios) Multiple(value bool) radios {
	return rc.set("multiple", value)
}

// Name field name
func (rc radios) Name(value string) radios {
	return rc.set("name", value)
}

// OnEvent event action configuration
func (rc radios) OnEvent(value any) radios {
	return rc.set("onEvent", value)
}

// Option option
func (rc radios) Option(value any) radios {
	return rc.set("option", value)
}

// Options options
func (rc radios) Options(value ...any) radios {
	return rc.set("options", value)
}

// Placeholder placeholder
func (rc radios) Placeholder(value string) radios {
	return rc.set("placeholder", value)
}

// ReadOnly whether to read only
func (rc radios) ReadOnly(value bool) radios {
	return rc.set("readOnly", value)
}

// Removable whether to removable
func (rc radios) Removable(value bool) radios {
	return rc.set("removable", value)
}

// RenderMode render mode
func (rc radios) RenderMode(value string) radios {
	return rc.set("renderMode", value)
}

// Required whether to required
func (rc radios) Required(value bool) radios {
	return rc.set("required", value)
}

// ShowIcon whether to show icon
func (rc radios) ShowIcon(value bool) radios {
	return rc.set("showIcon", value)
}

// ShowRadio whether to show radio
func (rc radios) ShowRadio(value bool) radios {
	return rc.set("showRadio", value)
}

// ShowText whether to show text
func (rc radios) ShowText(value bool) radios {
	return rc.set("showText", value)
}

// Size size
func (rc radios) Size(value string) radios {
	return rc.set("size", value)
}

// Source data source
func (rc radios) Source(value string) radios {
	return rc.set("source", value)
}

// SourceApi custom source API
func (rc radios) SourceApi(value string) radios {
	return rc.set("sourceApi", value)
}

// Toggle whether to toggle
func (rc radios) Toggle(value bool) radios {
	return rc.set("toggle", value)
}

// Value value
func (rc radios) Value(value any) radios {
	return rc.set("value", value)
}

// Visible whether to visible
func (rc radios) Visible(value bool) radios {
	return rc.set("visible", value)
}

// VisibleOn whether to visible by expression
func (rc radios) VisibleOn(value string) radios {
	return rc.set("visibleOn", value)
}

// Width width
func (rc radios) Width(value string) radios {
	return rc.set("width", value)
}

// WrapperClassName className for wrapper
func (rc radios) WrapperClassName(value string) radios {
	return rc.set("wrapperClassName", value)
}

// ValueField value field
func (rc radios) ValueField(value string) radios {
	return rc.set("valueField", value)
}

// LabelField label field
func (rc radios) LabelField(value string) radios {
	return rc.set("labelField", value)
}

// DescriptionField description field
func (rc radios) DescriptionField(value string) radios {
	return rc.set("descriptionField", value)
}

// OptionValueField option value field
func (rc radios) OptionValueField(value string) radios {
	return rc.set("optionValueField", value)
}

// Option label field
func (rc radios) OptionLabelField(value string) radios {
	return rc.set("optionLabelField", value)
}

// Option description field
func (rc radios) OptionDescriptionField(value string) radios {
	return rc.set("optionDescriptionField", value)
}
