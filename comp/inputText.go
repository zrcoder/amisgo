package comp

// inputText represents a text input field. Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/form/text

type inputText Schema

// InputText creates a text input field.
func InputText() inputText {
	return inputText{}.set("type", "input-text")
}

// InputEmail creates an email input field.
func InputEmail() inputText {
	return inputText{}.set("type", "input-email")
}

// InputUrl creates a URL input field.
func InputUrl() inputText {
	return inputText{}.set("type", "input-url")
}

// InputPassword creates a password input field.
func InputPassword() inputText {
	return inputText{}.set("type", "input-password")
}

// AddApi sets the API for adding data.
func (t inputText) AddApi(value string) inputText {
	return t.set("addApi", value)
}

// AddControls sets the form items for adding data.
func (t inputText) AddControls(value string) inputText {
	return t.set("addControls", value)
}

// AddDialog sets the dialog options for adding data.
func (t inputText) AddDialog(value string) inputText {
	return t.set("addDialog", value)
}

// AddOn sets additional options.
func (t inputText) AddOn(value any) inputText {
	return t.set("addOn", value)
}

// AutoComplete sets the API for autocomplete.
func (t inputText) AutoComplete(value string) inputText {
	return t.set("autoComplete", value)
}

// AutoFill sets the autofill options.
func (t inputText) AutoFill(value string) inputText {
	return t.set("autoFill", value)
}

// BorderMode sets the border mode. Options: full, half, none.
func (t inputText) BorderMode(value string) inputText {
	return t.set("borderMode", value)
}

// ClassName sets the CSS class name for the container.
func (t inputText) ClassName(value string) inputText {
	return t.set("className", value)
}

// ColumnClassName sets the CSS class name for the column.
func (t inputText) ColumnClassName(value string) inputText {
	return t.set("columnClassName", value)
}

// ClearValueOnEmpty clears the value when empty.
func (t inputText) ClearValueOnEmpty(value bool) inputText {
	return t.set("clearValueOnEmpty", value)
}

// ClearValueOnHidden clears the value when hidden.
func (t inputText) ClearValueOnHidden(value bool) inputText {
	return t.set("clearValueOnHidden", value)
}

// Clearable sets whether the input is clearable.
func (t inputText) Clearable(value bool) inputText {
	return t.set("clearable", value)
}

// Creatable sets whether new options can be created.
func (t inputText) Creatable(value bool) inputText {
	return t.set("creatable", value)
}

// CreateBtnLabel sets the label for the create button.
func (t inputText) CreateBtnLabel(value string) inputText {
	return t.set("createBtnLabel", value)
}

// DeferApi sets the API for deferred loading.
func (t inputText) DeferApi(value string) inputText {
	return t.set("deferApi", value)
}

// DeferField sets the field for deferred loading.
func (t inputText) DeferField(value string) inputText {
	return t.set("deferField", value)
}

// DeleteApi sets the API for deleting options.
func (t inputText) DeleteApi(value string) inputText {
	return t.set("deleteApi", value)
}

// DeleteConfirmText sets the confirmation text for deletion.
func (t inputText) DeleteConfirmText(value string) inputText {
	return t.set("deleteConfirmText", value)
}

// Delimiter sets the delimiter for multiple values.
func (t inputText) Delimiter(value string) inputText {
	return t.set("delimiter", value)
}

// Desc sets the description.
func (t inputText) Desc(value string) inputText {
	return t.set("desc", value)
}

// Description sets the HTML description.
func (t inputText) Description(value string) inputText {
	return t.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description.
func (t inputText) DescriptionClassName(value string) inputText {
	return t.set("descriptionClassName", value)
}

// Disabled sets whether the input is disabled.
func (t inputText) Disabled(value bool) inputText {
	return t.set("disabled", value)
}

// DisabledOn sets the expression to disable the input.
func (t inputText) DisabledOn(value string) inputText {
	return t.set("disabledOn", value)
}

// EditApi sets the API for editing data.
func (t inputText) EditApi(value string) inputText {
	return t.set("editApi", value)
}

// EditControls sets the form items for editing data.
func (t inputText) EditControls(value string) inputText {
	return t.set("editControls", value)
}

// EditDialog sets the dialog options for editing data.
func (t inputText) EditDialog(value string) inputText {
	return t.set("editDialog", value)
}

// Editable sets whether the input is editable.
func (t inputText) Editable(value bool) inputText {
	return t.set("editable", value)
}

// EditorSetting sets the editor configuration.
func (t inputText) EditorSetting(value string) inputText {
	return t.set("editorSetting", value)
}

// ExtraName sets the extra field name for range components.
func (t inputText) ExtraName(value string) inputText {
	return t.set("extraName", value)
}

// ExtractValue sets whether to extract the value as an array.
func (t inputText) ExtractValue(value bool) inputText {
	return t.set("extractValue", value)
}

// Hidden sets whether the input is hidden.
func (t inputText) Hidden(value bool) inputText {
	return t.set("hidden", value)
}

// HiddenOn sets the expression to hide the input.
func (t inputText) HiddenOn(value string) inputText {
	return t.set("hiddenOn", value)
}

// Hint sets the hint text.
func (t inputText) Hint(value string) inputText {
	return t.set("hint", value)
}

// Horizontal sets the horizontal layout configuration.
func (t inputText) Horizontal(value string) inputText {
	return t.set("horizontal", value)
}

// ID sets the unique ID for the component.
func (t inputText) ID(value string) inputText {
	return t.set("id", value)
}

// InitAutoFill sets the initial autofill value.
func (t inputText) InitAutoFill(value string) inputText {
	return t.set("initAutoFill", value)
}

// InitFetch sets whether to fetch data initially.
func (t inputText) InitFetch(value bool) inputText {
	return t.set("initFetch", value)
}

// InitFetchOn sets the expression to fetch data initially.
func (t inputText) InitFetchOn(value string) inputText {
	return t.set("initFetchOn", value)
}

// Inline sets whether the control is inline.
func (t inputText) Inline(value bool) inputText {
	return t.set("inline", value)
}

// InputClassName sets the CSS class name for the input.
func (t inputText) InputClassName(value string) inputText {
	return t.set("inputClassName", value)
}

// InputControlClassName sets the CSS class name for the control node.
func (t inputText) InputControlClassName(value string) inputText {
	return t.set("inputControlClassName", value)
}

// JoinValues sets whether to join values with a delimiter.
func (t inputText) JoinValues(value bool) inputText {
	return t.set("joinValues", value)
}

// Label sets the label text.
func (t inputText) Label(value string) inputText {
	return t.set("label", value)
}

// Size sets the size of the input. Options: xs, sm, md, lg, full.
func (t inputText) Size(value string) inputText {
	return t.set("size", value)
}

// Remark sets the remark text or object.
func (t inputText) Remark(value any) inputText {
	return t.set("remark", value)
}

// LabelAlign sets the label alignment. Options: right, left, top, inherit.
func (t inputText) LabelAlign(value string) inputText {
	return t.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label.
func (t inputText) LabelClassName(value string) inputText {
	return t.set("labelClassName", value)
}

// LabelRemark sets the remark for the label.
func (t inputText) LabelRemark(value string) inputText {
	return t.set("labelRemark", value)
}

// LabelWidth sets the custom width for the label.
func (t inputText) LabelWidth(value string) inputText {
	return t.set("labelWidth", value)
}

// MaxLength sets the maximum length of the input.
func (t inputText) MaxLength(value string) inputText {
	return t.set("maxLength", value)
}

// MinLength sets the minimum length of the input.
func (t inputText) MinLength(value string) inputText {
	return t.set("minLength", value)
}

// Mode sets the display mode. Options: normal, inline, horizontal.
func (t inputText) Mode(value string) inputText {
	return t.set("mode", value)
}

// Multiple sets whether the input allows multiple values.
func (t inputText) Multiple(value bool) inputText {
	return t.set("multiple", value)
}

// Name sets the name of the form item.
func (t inputText) Name(value string) inputText {
	return t.set("name", value)
}

// Required sets whether the input is required.
func (t inputText) Required(value bool) inputText {
	return t.set("required", value)
}

// Note sets the note text.
func (t inputText) Note(value string) inputText {
	return t.set("note", value)
}

// OnChange sets the event triggered on change.
func (t inputText) OnChange(value string) inputText {
	return t.set("onChange", value)
}

// OnFocus sets the event triggered on focus.
func (t inputText) OnFocus(value string) inputText {
	return t.set("onFocus", value)
}

// OnBlur sets the event triggered on blur.
func (t inputText) OnBlur(value string) inputText {
	return t.set("onBlur", value)
}

// OnInit sets the event triggered on initialization.
func (t inputText) OnInit(value string) inputText {
	return t.set("onInit", value)
}

// OnSearch sets the event triggered on search.
func (t inputText) OnSearch(value string) inputText {
	return t.set("onSearch", value)
}

// Option sets a single option.
func (t inputText) Option(value string) inputText {
	return t.set("option", value)
}

// Options sets multiple options.
func (t inputText) Options(value ...string) inputText {
	return t.set("options", value)
}

// OptionLabel sets the label position for options in multi-select mode. Options: top, left, right.
func (t inputText) OptionLabel(value string) inputText {
	return t.set("optionLabel", value)
}

// Placeholder sets the placeholder text.
func (t inputText) Placeholder(value string) inputText {
	return t.set("placeholder", value)
}

// ReadOnly sets whether the input is read-only.
func (t inputText) ReadOnly(value bool) inputText {
	return t.set("readOnly", value)
}

// RefixApi sets the API for fetching options.
func (t inputText) RefixApi(value string) inputText {
	return t.set("refixApi", value)
}

// ResetValue sets the reset value.
func (t inputText) ResetValue(value string) inputText {
	return t.set("resetValue", value)
}

// Rules sets the validation rules.
func (t inputText) Rules(value string) inputText {
	return t.set("rules", value)
}

// SearchApi sets the API for searching options.
func (t inputText) SearchApi(value string) inputText {
	return t.set("searchApi", value)
}

// Source sets the data source.
func (t inputText) Source(value string) inputText {
	return t.set("source", value)
}

// StaticClassName sets the CSS class name for static display.
func (t inputText) StaticClassName(value string) inputText {
	return t.set("staticClassName", value)
}

// StaticLabel sets the label for static display.
func (t inputText) StaticLabel(value string) inputText {
	return t.set("staticLabel", value)
}

// StaticOn sets the condition for static display.
func (t inputText) StaticOn(value string) inputText {
	return t.set("staticOn", value)
}

// StaticValue sets the static value.
func (t inputText) StaticValue(value string) inputText {
	return t.set("staticValue", value)
}

// Tips sets the tips text.
func (t inputText) Tips(value string) inputText {
	return t.set("tips", value)
}

// ValidationErrors sets the validation error messages.
func (t inputText) ValidationErrors(value string) inputText {
	return t.set("validationErrors", value)
}

// Value sets the value of the form item.
func (t inputText) Value(value any) inputText {
	return t.set("value", value)
}

// Visible sets whether the input is visible.
func (t inputText) Visible(value bool) inputText {
	return t.set("visible", value)
}

// VisibleOn sets the expression to control visibility.
func (t inputText) VisibleOn(value string) inputText {
	return t.set("visibleOn", value)
}

// Width sets the width of the input.
func (t inputText) Width(value string) inputText {
	return t.set("width", value)
}

// OnInput sets the event triggered on input.
func (t inputText) OnInput(value string) inputText {
	return t.set("onInput", value)
}

// ValueType sets the type of the value.
func (t inputText) ValueType(value string) inputText {
	return t.set("valueType", value)
}

// ValueOnEvent sets the value based on an event.
func (t inputText) ValueOnEvent(value any) inputText {
	return t.set("valueOnEvent", value)
}

// set sets a property of the inputText.
func (t inputText) set(key string, value any) inputText {
	t[key] = value
	return t
}
