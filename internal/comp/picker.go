package comp

import "github.com/zrcoder/amisgo/schema"

// Picker represents the Picker control
type Picker schema.Schema

func NewPicker() Picker {
	return Picker{"type": "picker"}
}

func (pc Picker) set(key string, value any) Picker {
	pc[key] = value
	return pc
}

// AddApi sets the API for adding items
func (pc Picker) AddApi(value string) Picker {
	return pc.set("addApi", value)
}

// AddControls sets the form items for adding
func (pc Picker) AddControls(value string) Picker {
	return pc.set("addControls", value)
}

// AddDialog sets the dialog for adding items
func (pc Picker) AddDialog(value string) Picker {
	return pc.set("addDialog", value)
}

// AutoFill sets the auto-fill option
func (pc Picker) AutoFill(value string) Picker {
	return pc.set("autoFill", value)
}

// ClassName sets the CSS class name
func (pc Picker) ClassName(value string) Picker {
	return pc.set("className", value)
}

// ClearValueOnHidden sets whether to clear value when hidden
func (pc Picker) ClearValueOnHidden(value bool) Picker {
	return pc.set("clearValueOnHidden", value)
}

// Clearable sets whether the picker is clearable
func (pc Picker) Clearable(value bool) Picker {
	return pc.set("clearable", value)
}

// Creatable sets whether new items can be created
func (pc Picker) Creatable(value bool) Picker {
	return pc.set("creatable", value)
}

// CreateBtnLabel sets the label for the create button
func (pc Picker) CreateBtnLabel(value string) Picker {
	return pc.set("createBtnLabel", value)
}

// DeferApi sets the API for deferred loading
func (pc Picker) DeferApi(value string) Picker {
	return pc.set("deferApi", value)
}

// DeferField sets the field for lazy loading
func (pc Picker) DeferField(value string) Picker {
	return pc.set("deferField", value)
}

// DeleteApi sets the API for deleting items
func (pc Picker) DeleteApi(value string) Picker {
	return pc.set("deleteApi", value)
}

// DeleteConfirmText sets the confirmation text for deletion
func (pc Picker) DeleteConfirmText(value string) Picker {
	return pc.set("deleteConfirmText", value)
}

// Delimiter sets the delimiter
func (pc Picker) Delimiter(value string) Picker {
	return pc.set("delimiter", value)
}

// Desc sets the description
func (pc Picker) Desc(value string) Picker {
	return pc.set("desc", value)
}

// Description sets the HTML description
func (pc Picker) Description(value string) Picker {
	return pc.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description
func (pc Picker) DescriptionClassName(value string) Picker {
	return pc.set("descriptionClassName", value)
}

// Disabled sets whether the picker is disabled
func (pc Picker) Disabled(value bool) Picker {
	return pc.set("disabled", value)
}

// DisabledOn sets the expression for disabling the picker
func (pc Picker) DisabledOn(value string) Picker {
	return pc.set("disabledOn", value)
}

// EditApi sets the API for editing items
func (pc Picker) EditApi(value string) Picker {
	return pc.set("editApi", value)
}

// EditControls sets the form items for editing
func (pc Picker) EditControls(value string) Picker {
	return pc.set("editControls", value)
}

// EditDialog sets the dialog for editing items
func (pc Picker) EditDialog(value string) Picker {
	return pc.set("editDialog", value)
}

// Editable sets whether the picker is editable
func (pc Picker) Editable(value bool) Picker {
	return pc.set("editable", value)
}

// EditorSetting sets the editor configuration
func (pc Picker) EditorSetting(value string) Picker {
	return pc.set("editorSetting", value)
}

// Embed sets whether the picker is in embedded mode
func (pc Picker) Embed(value bool) Picker {
	return pc.set("embed", value)
}

// ExtraName sets the extra field name
func (pc Picker) ExtraName(value string) Picker {
	return pc.set("extraName", value)
}

// ExtractValue sets whether to extract the value as an array
func (pc Picker) ExtractValue(value bool) Picker {
	return pc.set("extractValue", value)
}

// Hidden sets whether the picker is hidden
func (pc Picker) Hidden(value bool) Picker {
	return pc.set("hidden", value)
}

// HiddenOn sets the expression for hiding the picker
func (pc Picker) HiddenOn(value string) Picker {
	return pc.set("hiddenOn", value)
}

// Hint sets the input hint
func (pc Picker) Hint(value string) Picker {
	return pc.set("hint", value)
}

// Horizontal sets the horizontal layout configuration
func (pc Picker) Horizontal(value string) Picker {
	return pc.set("horizontal", value)
}

// ID sets the unique ID
func (pc Picker) ID(value string) Picker {
	return pc.set("id", value)
}

// InitAutoFill sets the initial auto-fill configuration
func (pc Picker) InitAutoFill(value string) Picker {
	return pc.set("initAutoFill", value)
}

// InitFetch sets whether to fetch data initially
func (pc Picker) InitFetch(value bool) Picker {
	return pc.set("initFetch", value)
}

// InitFetchOn sets the expression for initial data fetch
func (pc Picker) InitFetchOn(value string) Picker {
	return pc.set("initFetchOn", value)
}

// Inline sets whether the control is in inline mode
func (pc Picker) Inline(value bool) Picker {
	return pc.set("inline", value)
}

// InputClassName sets the CSS class name for the input
func (pc Picker) InputClassName(value string) Picker {
	return pc.set("inputClassName", value)
}

// JoinValues sets whether to join values in single selection mode
func (pc Picker) JoinValues(value bool) Picker {
	return pc.set("joinValues", value)
}

// Label sets the label
func (pc Picker) Label(value string) Picker {
	return pc.set("label", value)
}

// LabelAlign sets the label alignment
func (pc Picker) LabelAlign(value string) Picker {
	return pc.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label
func (pc Picker) LabelClassName(value string) Picker {
	return pc.set("labelClassName", value)
}

// LabelField sets the field name for the label
func (pc Picker) LabelField(value string) Picker {
	return pc.set("labelField", value)
}

// LabelRemark sets the label remark
func (pc Picker) LabelRemark(value string) Picker {
	return pc.set("labelRemark", value)
}

// LabelTpl sets the label template
func (pc Picker) LabelTpl(value string) Picker {
	return pc.set("labelTpl", value)
}

// LabelWidth sets the label width
func (pc Picker) LabelWidth(value int) Picker {
	return pc.set("labelWidth", value)
}

// ModalMode sets the modal mode
func (pc Picker) ModalMode(value string) Picker {
	return pc.set("modalMode", value)
}

// ModalTitle sets the modal title
func (pc Picker) ModalTitle(value any) Picker {
	return pc.set("modalTitle", value)
}

// Mode sets the component mode
func (pc Picker) Mode(value string) Picker {
	return pc.set("mode", value)
}

// Multiple sets whether multiple selection is supported
func (pc Picker) Multiple(value bool) Picker {
	return pc.set("multiple", value)
}

// Name sets the component name
func (pc Picker) Name(value string) Picker {
	return pc.set("name", value)
}

// OnEvent sets the event handler
func (pc Picker) OnEvent(value any) Picker {
	return pc.set("onEvent", value)
}

// Options sets the options list
func (pc Picker) Options(value ...any) Picker {
	return pc.set("options", value)
}

// OverflowConfig sets the overflow configuration
func (pc Picker) OverflowConfig(value string) Picker {
	return pc.set("overflowConfig", value)
}

// PickerSchema sets the picker schema.Schema
func (pc Picker) PickerSchema(value string) Picker {
	return pc.set("pickerSchema", value)
}

// Placeholder sets the placeholder text
func (pc Picker) Placeholder(value string) Picker {
	return pc.set("placeholder", value)
}

// ReadOnly sets whether the picker is read-only
func (pc Picker) ReadOnly(value bool) Picker {
	return pc.set("readOnly", value)
}

// ReadOnlyOn sets the expression for read-only mode
func (pc Picker) ReadOnlyOn(value string) Picker {
	return pc.set("readOnlyOn", value)
}

// Removable sets whether the picker is removable
func (pc Picker) Removable(value bool) Picker {
	return pc.set("removable", value)
}

// Required sets whether the picker is required
func (pc Picker) Required(value bool) Picker {
	return pc.set("required", value)
}

// ResetValue sets the reset value
func (pc Picker) ResetValue(value string) Picker {
	return pc.set("resetValue", value)
}

// Row sets the row for the form item
func (pc Picker) Row(value int) Picker {
	return pc.set("row", value)
}

// SaveImmediately sets whether to save immediately
func (pc Picker) SaveImmediately(value bool) Picker {
	return pc.set("saveImmediately", value)
}

// SelectFirst sets whether to select the first item
func (pc Picker) SelectFirst(value bool) Picker {
	return pc.set("selectFirst", value)
}

// Size sets the size
func (pc Picker) Size(value string) Picker {
	return pc.set("size", value)
}

// Source sets the data source
func (pc Picker) Source(value string) Picker {
	return pc.set("source", value)
}

// Static sets whether the picker is in static mode
func (pc Picker) Static(value bool) Picker {
	return pc.set("static", value)
}

// StaticClassName sets the CSS class name for static mode
func (pc Picker) StaticClassName(value string) Picker {
	return pc.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input
func (pc Picker) StaticInputClassName(value string) Picker {
	return pc.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label
func (pc Picker) StaticLabelClassName(value string) Picker {
	return pc.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static mode
func (pc Picker) StaticOn(value string) Picker {
	return pc.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static mode
func (pc Picker) StaticPlaceholder(value string) Picker {
	return pc.set("staticPlaceholder", value)
}

// StaticSchema sets the schema.Schema for static mode
func (pc Picker) StaticSchema(value string) Picker {
	return pc.set("staticSchema", value)
}

// Style sets the custom style
func (pc Picker) Style(value any) Picker {
	return pc.set("style", value)
}

// SubmitOnChange sets whether to submit on change
func (pc Picker) SubmitOnChange(value bool) Picker {
	return pc.set("submitOnChange", value)
}

// TestIdBuilder sets the function to generate test IDs
func (pc Picker) TestIdBuilder(value string) Picker {
	return pc.set("testIdBuilder", value)
}

// UseMobileUI sets whether to use mobile UI
func (pc Picker) UseMobileUI(value bool) Picker {
	return pc.set("useMobileUI", value)
}

// ValidateApi sets the API for validation
func (pc Picker) ValidateApi(value string) Picker {
	return pc.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on change
func (pc Picker) ValidateOnChange(value bool) Picker {
	return pc.set("validateOnChange", value)
}

// ValidationErrors sets the validation error messages
func (pc Picker) ValidationErrors(value string) Picker {
	return pc.set("validationErrors", value)
}

// Validations sets the validation rules
func (pc Picker) Validations(value string) Picker {
	return pc.set("validations", value)
}

// Value sets the value
func (pc Picker) Value(value string) Picker {
	return pc.set("value", value)
}

// ValueField sets the value field
func (pc Picker) ValueField(value string) Picker {
	return pc.set("valueField", value)
}

// ValuesNoWrap sets whether values should not wrap
func (pc Picker) ValuesNoWrap(value bool) Picker {
	return pc.set("valuesNoWrap", value)
}

// Visible sets whether the picker is visible
func (pc Picker) Visible(value bool) Picker {
	return pc.set("visible", value)
}

// VisibleOn sets the expression for visibility
func (pc Picker) VisibleOn(value string) Picker {
	return pc.set("visibleOn", value)
}

// Width sets the width
func (pc Picker) Width(value int) Picker {
	return pc.set("width", value)
}
