package comp

import "github.com/zrcoder/amisgo/model"

// picker represents the picker control
type picker model.Schema

// Picker creates a new PickerControl instance
func Picker() picker {
	return picker{"type": "picker"}
}

func (pc picker) set(key string, value any) picker {
	pc[key] = value
	return pc
}

// AddApi sets the API for adding items
func (pc picker) AddApi(value string) picker {
	return pc.set("addApi", value)
}

// AddControls sets the form items for adding
func (pc picker) AddControls(value string) picker {
	return pc.set("addControls", value)
}

// AddDialog sets the dialog for adding items
func (pc picker) AddDialog(value string) picker {
	return pc.set("addDialog", value)
}

// AutoFill sets the auto-fill option
func (pc picker) AutoFill(value string) picker {
	return pc.set("autoFill", value)
}

// ClassName sets the CSS class name
func (pc picker) ClassName(value string) picker {
	return pc.set("className", value)
}

// ClearValueOnHidden sets whether to clear value when hidden
func (pc picker) ClearValueOnHidden(value bool) picker {
	return pc.set("clearValueOnHidden", value)
}

// Clearable sets whether the picker is clearable
func (pc picker) Clearable(value bool) picker {
	return pc.set("clearable", value)
}

// Creatable sets whether new items can be created
func (pc picker) Creatable(value bool) picker {
	return pc.set("creatable", value)
}

// CreateBtnLabel sets the label for the create button
func (pc picker) CreateBtnLabel(value string) picker {
	return pc.set("createBtnLabel", value)
}

// DeferApi sets the API for deferred loading
func (pc picker) DeferApi(value string) picker {
	return pc.set("deferApi", value)
}

// DeferField sets the field for lazy loading
func (pc picker) DeferField(value string) picker {
	return pc.set("deferField", value)
}

// DeleteApi sets the API for deleting items
func (pc picker) DeleteApi(value string) picker {
	return pc.set("deleteApi", value)
}

// DeleteConfirmText sets the confirmation text for deletion
func (pc picker) DeleteConfirmText(value string) picker {
	return pc.set("deleteConfirmText", value)
}

// Delimiter sets the delimiter
func (pc picker) Delimiter(value string) picker {
	return pc.set("delimiter", value)
}

// Desc sets the description
func (pc picker) Desc(value string) picker {
	return pc.set("desc", value)
}

// Description sets the HTML description
func (pc picker) Description(value string) picker {
	return pc.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description
func (pc picker) DescriptionClassName(value string) picker {
	return pc.set("descriptionClassName", value)
}

// Disabled sets whether the picker is disabled
func (pc picker) Disabled(value bool) picker {
	return pc.set("disabled", value)
}

// DisabledOn sets the expression for disabling the picker
func (pc picker) DisabledOn(value string) picker {
	return pc.set("disabledOn", value)
}

// EditApi sets the API for editing items
func (pc picker) EditApi(value string) picker {
	return pc.set("editApi", value)
}

// EditControls sets the form items for editing
func (pc picker) EditControls(value string) picker {
	return pc.set("editControls", value)
}

// EditDialog sets the dialog for editing items
func (pc picker) EditDialog(value string) picker {
	return pc.set("editDialog", value)
}

// Editable sets whether the picker is editable
func (pc picker) Editable(value bool) picker {
	return pc.set("editable", value)
}

// EditorSetting sets the editor configuration
func (pc picker) EditorSetting(value string) picker {
	return pc.set("editorSetting", value)
}

// Embed sets whether the picker is in embedded mode
func (pc picker) Embed(value bool) picker {
	return pc.set("embed", value)
}

// ExtraName sets the extra field name
func (pc picker) ExtraName(value string) picker {
	return pc.set("extraName", value)
}

// ExtractValue sets whether to extract the value as an array
func (pc picker) ExtractValue(value bool) picker {
	return pc.set("extractValue", value)
}

// Hidden sets whether the picker is hidden
func (pc picker) Hidden(value bool) picker {
	return pc.set("hidden", value)
}

// HiddenOn sets the expression for hiding the picker
func (pc picker) HiddenOn(value string) picker {
	return pc.set("hiddenOn", value)
}

// Hint sets the input hint
func (pc picker) Hint(value string) picker {
	return pc.set("hint", value)
}

// Horizontal sets the horizontal layout configuration
func (pc picker) Horizontal(value string) picker {
	return pc.set("horizontal", value)
}

// ID sets the unique ID
func (pc picker) ID(value string) picker {
	return pc.set("id", value)
}

// InitAutoFill sets the initial auto-fill configuration
func (pc picker) InitAutoFill(value string) picker {
	return pc.set("initAutoFill", value)
}

// InitFetch sets whether to fetch data initially
func (pc picker) InitFetch(value bool) picker {
	return pc.set("initFetch", value)
}

// InitFetchOn sets the expression for initial data fetch
func (pc picker) InitFetchOn(value string) picker {
	return pc.set("initFetchOn", value)
}

// Inline sets whether the control is in inline mode
func (pc picker) Inline(value bool) picker {
	return pc.set("inline", value)
}

// InputClassName sets the CSS class name for the input
func (pc picker) InputClassName(value string) picker {
	return pc.set("inputClassName", value)
}

// JoinValues sets whether to join values in single selection mode
func (pc picker) JoinValues(value bool) picker {
	return pc.set("joinValues", value)
}

// Label sets the label
func (pc picker) Label(value string) picker {
	return pc.set("label", value)
}

// LabelAlign sets the label alignment
func (pc picker) LabelAlign(value string) picker {
	return pc.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label
func (pc picker) LabelClassName(value string) picker {
	return pc.set("labelClassName", value)
}

// LabelField sets the field name for the label
func (pc picker) LabelField(value string) picker {
	return pc.set("labelField", value)
}

// LabelRemark sets the label remark
func (pc picker) LabelRemark(value string) picker {
	return pc.set("labelRemark", value)
}

// LabelTpl sets the label template
func (pc picker) LabelTpl(value string) picker {
	return pc.set("labelTpl", value)
}

// LabelWidth sets the label width
func (pc picker) LabelWidth(value int) picker {
	return pc.set("labelWidth", value)
}

// ModalMode sets the modal mode
func (pc picker) ModalMode(value string) picker {
	return pc.set("modalMode", value)
}

// ModalTitle sets the modal title
func (pc picker) ModalTitle(value any) picker {
	return pc.set("modalTitle", value)
}

// Mode sets the component mode
func (pc picker) Mode(value string) picker {
	return pc.set("mode", value)
}

// Multiple sets whether multiple selection is supported
func (pc picker) Multiple(value bool) picker {
	return pc.set("multiple", value)
}

// Name sets the component name
func (pc picker) Name(value string) picker {
	return pc.set("name", value)
}

// OnEvent sets the event handler
func (pc picker) OnEvent(value any) picker {
	return pc.set("onEvent", value)
}

// Options sets the options list
func (pc picker) Options(value ...any) picker {
	return pc.set("options", value)
}

// OverflowConfig sets the overflow configuration
func (pc picker) OverflowConfig(value string) picker {
	return pc.set("overflowConfig", value)
}

// PickerSchema sets the picker schema
func (pc picker) PickerSchema(value string) picker {
	return pc.set("pickerSchema", value)
}

// Placeholder sets the placeholder text
func (pc picker) Placeholder(value string) picker {
	return pc.set("placeholder", value)
}

// ReadOnly sets whether the picker is read-only
func (pc picker) ReadOnly(value bool) picker {
	return pc.set("readOnly", value)
}

// ReadOnlyOn sets the expression for read-only mode
func (pc picker) ReadOnlyOn(value string) picker {
	return pc.set("readOnlyOn", value)
}

// Removable sets whether the picker is removable
func (pc picker) Removable(value bool) picker {
	return pc.set("removable", value)
}

// Required sets whether the picker is required
func (pc picker) Required(value bool) picker {
	return pc.set("required", value)
}

// ResetValue sets the reset value
func (pc picker) ResetValue(value string) picker {
	return pc.set("resetValue", value)
}

// Row sets the row for the form item
func (pc picker) Row(value int) picker {
	return pc.set("row", value)
}

// SaveImmediately sets whether to save immediately
func (pc picker) SaveImmediately(value bool) picker {
	return pc.set("saveImmediately", value)
}

// SelectFirst sets whether to select the first item
func (pc picker) SelectFirst(value bool) picker {
	return pc.set("selectFirst", value)
}

// Size sets the size
func (pc picker) Size(value string) picker {
	return pc.set("size", value)
}

// Source sets the data source
func (pc picker) Source(value string) picker {
	return pc.set("source", value)
}

// Static sets whether the picker is in static mode
func (pc picker) Static(value bool) picker {
	return pc.set("static", value)
}

// StaticClassName sets the CSS class name for static mode
func (pc picker) StaticClassName(value string) picker {
	return pc.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input
func (pc picker) StaticInputClassName(value string) picker {
	return pc.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label
func (pc picker) StaticLabelClassName(value string) picker {
	return pc.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static mode
func (pc picker) StaticOn(value string) picker {
	return pc.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static mode
func (pc picker) StaticPlaceholder(value string) picker {
	return pc.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static mode
func (pc picker) StaticSchema(value string) picker {
	return pc.set("staticSchema", value)
}

// Style sets the custom style
func (pc picker) Style(value any) picker {
	return pc.set("style", value)
}

// SubmitOnChange sets whether to submit on change
func (pc picker) SubmitOnChange(value bool) picker {
	return pc.set("submitOnChange", value)
}

// TestIdBuilder sets the function to generate test IDs
func (pc picker) TestIdBuilder(value string) picker {
	return pc.set("testIdBuilder", value)
}

// UseMobileUI sets whether to use mobile UI
func (pc picker) UseMobileUI(value bool) picker {
	return pc.set("useMobileUI", value)
}

// ValidateApi sets the API for validation
func (pc picker) ValidateApi(value string) picker {
	return pc.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on change
func (pc picker) ValidateOnChange(value bool) picker {
	return pc.set("validateOnChange", value)
}

// ValidationErrors sets the validation error messages
func (pc picker) ValidationErrors(value string) picker {
	return pc.set("validationErrors", value)
}

// Validations sets the validation rules
func (pc picker) Validations(value string) picker {
	return pc.set("validations", value)
}

// Value sets the value
func (pc picker) Value(value string) picker {
	return pc.set("value", value)
}

// ValueField sets the value field
func (pc picker) ValueField(value string) picker {
	return pc.set("valueField", value)
}

// ValuesNoWrap sets whether values should not wrap
func (pc picker) ValuesNoWrap(value bool) picker {
	return pc.set("valuesNoWrap", value)
}

// Visible sets whether the picker is visible
func (pc picker) Visible(value bool) picker {
	return pc.set("visible", value)
}

// VisibleOn sets the expression for visibility
func (pc picker) VisibleOn(value string) picker {
	return pc.set("visibleOn", value)
}

// Width sets the width
func (pc picker) Width(value int) picker {
	return pc.set("width", value)
}
