package comp

// inputTag represents an input tag component. Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/form/tag

type inputTag Schema

// InputTag initializes an input tag component.
func InputTag() inputTag {
	return inputTag{}.set("type", "input-tag")
}

// set sets a key-value pair for the input tag component.
func (t inputTag) set(key string, value any) inputTag {
	t[key] = value
	return t
}

// AddApi sets the API to call when adding a tag.
func (t inputTag) AddApi(value string) inputTag {
	return t.set("addApi", value)
}

// AddControls sets the form items for adding a tag.
func (t inputTag) AddControls(value string) inputTag {
	return t.set("addControls", value)
}

// AddDialog sets the configuration for the add dialog.
func (t inputTag) AddDialog(value string) inputTag {
	return t.set("addDialog", value)
}

// AutoFill sets the auto-fill configuration.
func (t inputTag) AutoFill(value string) inputTag {
	return t.set("autoFill", value)
}

// ClassName sets the CSS class name for the container.
func (t inputTag) ClassName(value string) inputTag {
	return t.set("className", value)
}

// ClearValueOnHidden sets whether to clear the value when the form item is hidden.
func (t inputTag) ClearValueOnHidden(value bool) inputTag {
	return t.set("clearValueOnHidden", value)
}

// Clearable sets whether the input tag is clearable.
func (t inputTag) Clearable(value bool) inputTag {
	return t.set("clearable", value)
}

// Creatable sets whether new tags can be created.
func (t inputTag) Creatable(value bool) inputTag {
	return t.set("creatable", value)
}

// CreateBtnLabel sets the label for the create button.
func (t inputTag) CreateBtnLabel(value string) inputTag {
	return t.set("createBtnLabel", value)
}

// DeferApi sets the API for deferred loading.
func (t inputTag) DeferApi(value string) inputTag {
	return t.set("deferApi", value)
}

// DeferField sets the field for lazy loading.
func (t inputTag) DeferField(value string) inputTag {
	return t.set("deferField", value)
}

// DeleteApi sets the API for deleting tags.
func (t inputTag) DeleteApi(value string) inputTag {
	return t.set("deleteApi", value)
}

// DeleteConfirmText sets the confirmation text for deleting tags.
func (t inputTag) DeleteConfirmText(value string) inputTag {
	return t.set("deleteConfirmText", value)
}

// Delimiter sets the delimiter for tags.
func (t inputTag) Delimiter(value string) inputTag {
	return t.set("delimiter", value)
}

// Desc sets the description.
func (t inputTag) Desc(value string) inputTag {
	return t.set("desc", value)
}

// Description sets the description content, supporting HTML.
func (t inputTag) Description(value string) inputTag {
	return t.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description.
func (t inputTag) DescriptionClassName(value string) inputTag {
	return t.set("descriptionClassName", value)
}

// Disabled sets whether the input tag is disabled.
func (t inputTag) Disabled(value bool) inputTag {
	return t.set("disabled", value)
}

// DisabledOn sets the expression to determine if the input tag is disabled.
func (t inputTag) DisabledOn(value string) inputTag {
	return t.set("disabledOn", value)
}

// Dropdown sets whether the input tag is in dropdown mode.
func (t inputTag) Dropdown(value bool) inputTag {
	return t.set("dropdown", value)
}

// EditApi sets the API for editing tags.
func (t inputTag) EditApi(value string) inputTag {
	return t.set("editApi", value)
}

// EditControls sets the form items for editing tags.
func (t inputTag) EditControls(value string) inputTag {
	return t.set("editControls", value)
}

// EditDialog sets the configuration for the edit dialog.
func (t inputTag) EditDialog(value string) inputTag {
	return t.set("editDialog", value)
}

// Editable sets whether the input tag is editable.
func (t inputTag) Editable(value bool) inputTag {
	return t.set("editable", value)
}

// EditorSetting sets the editor configuration.
func (t inputTag) EditorSetting(value string) inputTag {
	return t.set("editorSetting", value)
}

// EnableBatchAdd sets whether batch add mode is enabled.
func (t inputTag) EnableBatchAdd(value bool) inputTag {
	return t.set("enableBatchAdd", value)
}

// ExtraName sets the extra field name for range components.
func (t inputTag) ExtraName(value string) inputTag {
	return t.set("extraName", value)
}

// ExtractValue sets whether to extract the value of selected options as an array.
func (t inputTag) ExtractValue(value bool) inputTag {
	return t.set("extractValue", value)
}

// Hidden sets whether the input tag is hidden.
func (t inputTag) Hidden(value bool) inputTag {
	return t.set("hidden", value)
}

// HiddenOn sets the expression to determine if the input tag is hidden.
func (t inputTag) HiddenOn(value string) inputTag {
	return t.set("hiddenOn", value)
}

// Hint sets the input hint displayed on focus.
func (t inputTag) Hint(value string) inputTag {
	return t.set("hint", value)
}

// Horizontal sets the horizontal layout configuration.
func (t inputTag) Horizontal(value string) inputTag {
	return t.set("horizontal", value)
}

// ID sets the unique ID for the component, mainly for logging.
func (t inputTag) ID(value string) inputTag {
	return t.set("id", value)
}

// InitAutoFill sets the initial auto-fill configuration.
func (t inputTag) InitAutoFill(value string) inputTag {
	return t.set("initAutoFill", value)
}

// InitFetch sets whether to fetch the source API initially.
func (t inputTag) InitFetch(value bool) inputTag {
	return t.set("initFetch", value)
}

// InitFetchOn sets the expression to determine if the source API should be fetched initially.
func (t inputTag) InitFetchOn(value string) inputTag {
	return t.set("initFetchOn", value)
}

// Inline sets whether the form control is in inline mode.
func (t inputTag) Inline(value bool) inputTag {
	return t.set("inline", value)
}

// InputClassName sets the CSS class name for the input.
func (t inputTag) InputClassName(value string) inputTag {
	return t.set("inputClassName", value)
}

// JoinValues sets whether to join selected option values with a delimiter.
func (t inputTag) JoinValues(value bool) inputTag {
	return t.set("joinValues", value)
}

// Label sets the label for the input tag.
func (t inputTag) Label(value string) inputTag {
	return t.set("label", value)
}

// LabelAlign sets the alignment for the label. Options: right, left, top, inherit.
func (t inputTag) LabelAlign(value string) inputTag {
	return t.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label.
func (t inputTag) LabelClassName(value string) inputTag {
	return t.set("labelClassName", value)
}

// LabelRemark sets the remark for the label, displayed as a small icon with a tooltip.
func (t inputTag) LabelRemark(value string) inputTag {
	return t.set("labelRemark", value)
}

// LabelWidth sets the custom width for the label, default unit is px.
func (t inputTag) LabelWidth(value string) inputTag {
	return t.set("labelWidth", value)
}

// Max sets the maximum number of tags that can be added.
func (t inputTag) Max(value string) inputTag {
	return t.set("max", value)
}

// MaxTagCount sets the maximum number of tags to display, with overflow shown in a dropdown.
func (t inputTag) MaxTagCount(value string) inputTag {
	return t.set("maxTagCount", value)
}

// MaxTagText sets the text for the tag overflow indicator, default is "...".
func (t inputTag) MaxTagText(value string) inputTag {
	return t.set("maxTagText", value)
}

// Min sets the minimum number of tags that must be selected.
func (t inputTag) Min(value string) inputTag {
	return t.set("min", value)
}

// Mode sets the display mode for the form item.
func (t inputTag) Mode(value string) inputTag {
	return t.set("mode", value)
}

// Name sets the name attribute for the form item.
func (t inputTag) Name(value string) inputTag {
	return t.set("name", value)
}

// OnEvent sets the event handlers for the input tag.
func (t inputTag) OnEvent(value map[string]string) inputTag {
	return t.set("onEvent", value)
}

// Options sets the options for the input tag.
func (t inputTag) Options(value ...any) inputTag {
	return t.set("options", value)
}

// OptionValuePath sets the path to the value in the options.
func (t inputTag) OptionValuePath(value string) inputTag {
	return t.set("optionValuePath", value)
}

// OptionLabelPath sets the path to the label in the options.
func (t inputTag) OptionLabelPath(value string) inputTag {
	return t.set("optionLabelPath", value)
}

// OptionClassName sets the CSS class name for the options.
func (t inputTag) OptionClassName(value string) inputTag {
	return t.set("optionClassName", value)
}

// Placeholder sets the placeholder text for the input.
func (t inputTag) Placeholder(value string) inputTag {
	return t.set("placeholder", value)
}

// ReadOnly sets whether the input tag is read-only.
func (t inputTag) ReadOnly(value bool) inputTag {
	return t.set("readOnly", value)
}

// Required sets whether the input tag is required.
func (t inputTag) Required(value bool) inputTag {
	return t.set("required", value)
}

// Rules sets the validation rules for the input tag.
func (t inputTag) Rules(value string) inputTag {
	return t.set("rules", value)
}

// Source sets the data source for the input tag.
func (t inputTag) Source(value string) inputTag {
	return t.set("source", value)
}

// TagClassName sets the CSS class name for the tags.
func (t inputTag) TagClassName(value string) inputTag {
	return t.set("tagClassName", value)
}

// TagTpl sets the custom template for the tags.
func (t inputTag) TagTpl(value string) inputTag {
	return t.set("tagTpl", value)
}

// Touchable sets whether the input tag supports touch operations.
func (t inputTag) Touchable(value bool) inputTag {
	return t.set("touchable", value)
}

// Value sets the value for the form item.
func (t inputTag) Value(value any) inputTag {
	return t.set("value", value)
}

// VerticalLayout sets whether the input tag uses vertical layout.
func (t inputTag) VerticalLayout(value bool) inputTag {
	return t.set("verticalLayout", value)
}

// Visible sets whether the input tag is visible.
func (t inputTag) Visible(value bool) inputTag {
	return t.set("visible", value)
}

// ValidationErrors sets custom validation error messages.
func (t inputTag) ValidationErrors(value map[string]string) inputTag {
	return t.set("validationErrors", value)
}

// ValuePath sets the path to the value in the data.
func (t inputTag) ValuePath(value string) inputTag {
	return t.set("valuePath", value)
}

// ValidateOnChange sets whether to validate the input tag on change.
func (t inputTag) ValidateOnChange(value bool) inputTag {
	return t.set("validateOnChange", value)
}

// Validations sets the validation rules for the input tag.
func (t inputTag) Validations(value string) inputTag {
	return t.set("validations", value)
}

// ValuesNoWrap sets whether to avoid line wrapping for multiple values.
func (t inputTag) ValuesNoWrap(value bool) inputTag {
	return t.set("valuesNoWrap", value)
}

// VisibleOn sets the expression to determine if the input tag is visible.
func (t inputTag) VisibleOn(value string) inputTag {
	return t.set("visibleOn", value)
}

// Width sets the width of the input tag in a table.
func (t inputTag) Width(value string) inputTag {
	return t.set("width", value)
}
