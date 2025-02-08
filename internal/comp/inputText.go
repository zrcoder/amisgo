package comp

import "github.com/zrcoder/amisgo/model"

// InputText represents a text input field. Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/form/text
type InputText model.Schema

func NewInputText() InputText {
	return InputText{"type": "input-text"}
}

func NewInputEmail() InputText {
	return InputText{"type": "input-email"}
}

func NewInputUrl() InputText {
	return InputText{"type": "input-url"}
}

func NewInputPassword() InputText {
	return InputText{"type": "input-password"}
}

// AddApi sets the API for adding data.
func (t InputText) AddApi(value string) InputText {
	return t.set("addApi", value)
}

// AddControls sets the form items for adding data.
func (t InputText) AddControls(value string) InputText {
	return t.set("addControls", value)
}

// AddDialog sets the dialog options for adding data.
func (t InputText) AddDialog(value string) InputText {
	return t.set("addDialog", value)
}

// AddOn sets additional options.
func (t InputText) AddOn(value any) InputText {
	return t.set("addOn", value)
}

// AutoComplete sets the API for autocomplete.
func (t InputText) AutoComplete(value string) InputText {
	return t.set("autoComplete", value)
}

// AutoFill sets the autofill options.
func (t InputText) AutoFill(value string) InputText {
	return t.set("autoFill", value)
}

// BorderMode sets the border mode. Options: full, half, none.
func (t InputText) BorderMode(value string) InputText {
	return t.set("borderMode", value)
}

// ClassName sets the CSS class name for the container.
func (t InputText) ClassName(value string) InputText {
	return t.set("className", value)
}

// ColumnClassName sets the CSS class name for the column.
func (t InputText) ColumnClassName(value string) InputText {
	return t.set("columnClassName", value)
}

// ClearValueOnEmpty clears the value when empty.
func (t InputText) ClearValueOnEmpty(value bool) InputText {
	return t.set("clearValueOnEmpty", value)
}

// ClearValueOnHidden clears the value when hidden.
func (t InputText) ClearValueOnHidden(value bool) InputText {
	return t.set("clearValueOnHidden", value)
}

// Clearable sets whether the input is clearable.
func (t InputText) Clearable(value bool) InputText {
	return t.set("clearable", value)
}

// Creatable sets whether new options can be created.
func (t InputText) Creatable(value bool) InputText {
	return t.set("creatable", value)
}

// CreateBtnLabel sets the label for the create button.
func (t InputText) CreateBtnLabel(value string) InputText {
	return t.set("createBtnLabel", value)
}

// DeferApi sets the API for deferred loading.
func (t InputText) DeferApi(value string) InputText {
	return t.set("deferApi", value)
}

// DeferField sets the field for deferred loading.
func (t InputText) DeferField(value string) InputText {
	return t.set("deferField", value)
}

// DeleteApi sets the API for deleting options.
func (t InputText) DeleteApi(value string) InputText {
	return t.set("deleteApi", value)
}

// DeleteConfirmText sets the confirmation text for deletion.
func (t InputText) DeleteConfirmText(value string) InputText {
	return t.set("deleteConfirmText", value)
}

// Delimiter sets the delimiter for multiple values.
func (t InputText) Delimiter(value string) InputText {
	return t.set("delimiter", value)
}

// Desc sets the description.
func (t InputText) Desc(value string) InputText {
	return t.set("desc", value)
}

// Description sets the HTML description.
func (t InputText) Description(value string) InputText {
	return t.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description.
func (t InputText) DescriptionClassName(value string) InputText {
	return t.set("descriptionClassName", value)
}

// Disabled sets whether the input is disabled.
func (t InputText) Disabled(value bool) InputText {
	return t.set("disabled", value)
}

// DisabledOn sets the expression to disable the input.
func (t InputText) DisabledOn(value string) InputText {
	return t.set("disabledOn", value)
}

// EditApi sets the API for editing data.
func (t InputText) EditApi(value string) InputText {
	return t.set("editApi", value)
}

// EditControls sets the form items for editing data.
func (t InputText) EditControls(value string) InputText {
	return t.set("editControls", value)
}

// EditDialog sets the dialog options for editing data.
func (t InputText) EditDialog(value string) InputText {
	return t.set("editDialog", value)
}

// Editable sets whether the input is editable.
func (t InputText) Editable(value bool) InputText {
	return t.set("editable", value)
}

// EditorSetting sets the editor configuration.
func (t InputText) EditorSetting(value string) InputText {
	return t.set("editorSetting", value)
}

// ExtraName sets the extra field name for range components.
func (t InputText) ExtraName(value string) InputText {
	return t.set("extraName", value)
}

// ExtractValue sets whether to extract the value as an array.
func (t InputText) ExtractValue(value bool) InputText {
	return t.set("extractValue", value)
}

// Hidden sets whether the input is hidden.
func (t InputText) Hidden(value bool) InputText {
	return t.set("hidden", value)
}

// HiddenOn sets the expression to hide the input.
func (t InputText) HiddenOn(value string) InputText {
	return t.set("hiddenOn", value)
}

// Hint sets the hint text.
func (t InputText) Hint(value string) InputText {
	return t.set("hint", value)
}

// Horizontal sets the horizontal layout configuration.
func (t InputText) Horizontal(value string) InputText {
	return t.set("horizontal", value)
}

// ID sets the unique ID for the component.
func (t InputText) ID(value string) InputText {
	return t.set("id", value)
}

// InitAutoFill sets the initial autofill value.
func (t InputText) InitAutoFill(value string) InputText {
	return t.set("initAutoFill", value)
}

// InitFetch sets whether to fetch data initially.
func (t InputText) InitFetch(value bool) InputText {
	return t.set("initFetch", value)
}

// InitFetchOn sets the expression to fetch data initially.
func (t InputText) InitFetchOn(value string) InputText {
	return t.set("initFetchOn", value)
}

// Inline sets whether the control is inline.
func (t InputText) Inline(value bool) InputText {
	return t.set("inline", value)
}

// InputClassName sets the CSS class name for the input.
func (t InputText) InputClassName(value string) InputText {
	return t.set("inputClassName", value)
}

// InputControlClassName sets the CSS class name for the control node.
func (t InputText) InputControlClassName(value string) InputText {
	return t.set("inputControlClassName", value)
}

// JoinValues sets whether to join values with a delimiter.
func (t InputText) JoinValues(value bool) InputText {
	return t.set("joinValues", value)
}

// Label sets the label text.
func (t InputText) Label(value string) InputText {
	return t.set("label", value)
}

// Size sets the size of the input. Options: xs, sm, md, lg, full.
func (t InputText) Size(value string) InputText {
	return t.set("size", value)
}

// Remark sets the remark text or object.
func (t InputText) Remark(value any) InputText {
	return t.set("remark", value)
}

// LabelAlign sets the label alignment. Options: right, left, top, inherit.
func (t InputText) LabelAlign(value string) InputText {
	return t.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label.
func (t InputText) LabelClassName(value string) InputText {
	return t.set("labelClassName", value)
}

// LabelRemark sets the remark for the label.
func (t InputText) LabelRemark(value string) InputText {
	return t.set("labelRemark", value)
}

// LabelWidth sets the custom width for the label.
func (t InputText) LabelWidth(value string) InputText {
	return t.set("labelWidth", value)
}

// MaxLength sets the maximum length of the input.
func (t InputText) MaxLength(value string) InputText {
	return t.set("maxLength", value)
}

// MinLength sets the minimum length of the input.
func (t InputText) MinLength(value string) InputText {
	return t.set("minLength", value)
}

// Mode sets the display mode. Options: normal, inline, horizontal.
func (t InputText) Mode(value string) InputText {
	return t.set("mode", value)
}

// Multiple sets whether the input allows multiple values.
func (t InputText) Multiple(value bool) InputText {
	return t.set("multiple", value)
}

// Name sets the name of the form item.
func (t InputText) Name(value string) InputText {
	return t.set("name", value)
}

// Required sets whether the input is required.
func (t InputText) Required(value bool) InputText {
	return t.set("required", value)
}

// Note sets the note text.
func (t InputText) Note(value string) InputText {
	return t.set("note", value)
}

// OnChange sets the event triggered on change.
func (t InputText) OnChange(value string) InputText {
	return t.set("onChange", value)
}

// OnFocus sets the event triggered on focus.
func (t InputText) OnFocus(value string) InputText {
	return t.set("onFocus", value)
}

// OnBlur sets the event triggered on blur.
func (t InputText) OnBlur(value string) InputText {
	return t.set("onBlur", value)
}

// OnInit sets the event triggered on initialization.
func (t InputText) OnInit(value string) InputText {
	return t.set("onInit", value)
}

// OnSearch sets the event triggered on search.
func (t InputText) OnSearch(value string) InputText {
	return t.set("onSearch", value)
}

// Option sets a single option.
func (t InputText) Option(value string) InputText {
	return t.set("option", value)
}

// Options sets multiple options.
func (t InputText) Options(value ...string) InputText {
	return t.set("options", value)
}

// OptionLabel sets the label position for options in multi-select mode. Options: top, left, right.
func (t InputText) OptionLabel(value string) InputText {
	return t.set("optionLabel", value)
}

// Placeholder sets the placeholder text.
func (t InputText) Placeholder(value string) InputText {
	return t.set("placeholder", value)
}

// ReadOnly sets whether the input is read-only.
func (t InputText) ReadOnly(value bool) InputText {
	return t.set("readOnly", value)
}

// RefixApi sets the API for fetching options.
func (t InputText) RefixApi(value string) InputText {
	return t.set("refixApi", value)
}

// ResetValue sets the reset value.
func (t InputText) ResetValue(value string) InputText {
	return t.set("resetValue", value)
}

// Rules sets the validation rules.
func (t InputText) Rules(value string) InputText {
	return t.set("rules", value)
}

// SearchApi sets the API for searching options.
func (t InputText) SearchApi(value string) InputText {
	return t.set("searchApi", value)
}

// Source sets the data source.
func (t InputText) Source(value string) InputText {
	return t.set("source", value)
}

// StaticClassName sets the CSS class name for static display.
func (t InputText) StaticClassName(value string) InputText {
	return t.set("staticClassName", value)
}

// StaticLabel sets the label for static display.
func (t InputText) StaticLabel(value string) InputText {
	return t.set("staticLabel", value)
}

// StaticOn sets the condition for static display.
func (t InputText) StaticOn(value string) InputText {
	return t.set("staticOn", value)
}

// StaticValue sets the static value.
func (t InputText) StaticValue(value string) InputText {
	return t.set("staticValue", value)
}

// Tips sets the tips text.
func (t InputText) Tips(value string) InputText {
	return t.set("tips", value)
}

// ValidationErrors sets the validation error messages.
func (t InputText) ValidationErrors(value string) InputText {
	return t.set("validationErrors", value)
}

// Value sets the value of the form item.
func (t InputText) Value(value any) InputText {
	return t.set("value", value)
}

// Visible sets whether the input is visible.
func (t InputText) Visible(value bool) InputText {
	return t.set("visible", value)
}

// VisibleOn sets the expression to control visibility.
func (t InputText) VisibleOn(value string) InputText {
	return t.set("visibleOn", value)
}

// Width sets the width of the input.
func (t InputText) Width(value string) InputText {
	return t.set("width", value)
}

// OnInput sets the event triggered on input.
func (t InputText) OnInput(value string) InputText {
	return t.set("onInput", value)
}

// ValueType sets the type of the value.
func (t InputText) ValueType(value string) InputText {
	return t.set("valueType", value)
}

// ValueOnEvent sets the value based on an event.
func (t InputText) ValueOnEvent(value any) InputText {
	return t.set("valueOnEvent", value)
}

// set sets a property of the inputText.
func (t InputText) set(key string, value any) InputText {
	t[key] = value
	return t
}
