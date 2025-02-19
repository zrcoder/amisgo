package comp

import "github.com/zrcoder/amisgo/schema"

// InputTag represents an input tag component. Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/form/tag
type InputTag schema.Schema

func NewInputTag() InputTag {
	return InputTag{"type": "input-tag"}
}

// set sets a key-value pair for the input tag component.
func (t InputTag) set(key string, value any) InputTag {
	t[key] = value
	return t
}

// AddApi sets the API to call when adding a tag.
func (t InputTag) AddApi(value string) InputTag {
	return t.set("addApi", value)
}

// AddControls sets the form items for adding a tag.
func (t InputTag) AddControls(value string) InputTag {
	return t.set("addControls", value)
}

// AddDialog sets the configuration for the add dialog.
func (t InputTag) AddDialog(value string) InputTag {
	return t.set("addDialog", value)
}

// AutoFill sets the auto-fill configuration.
func (t InputTag) AutoFill(value string) InputTag {
	return t.set("autoFill", value)
}

// ClassName sets the CSS class name for the container.
func (t InputTag) ClassName(value string) InputTag {
	return t.set("className", value)
}

// ClearValueOnHidden sets whether to clear the value when the form item is hidden.
func (t InputTag) ClearValueOnHidden(value bool) InputTag {
	return t.set("clearValueOnHidden", value)
}

// Clearable sets whether the input tag is clearable.
func (t InputTag) Clearable(value bool) InputTag {
	return t.set("clearable", value)
}

// Creatable sets whether new tags can be created.
func (t InputTag) Creatable(value bool) InputTag {
	return t.set("creatable", value)
}

// CreateBtnLabel sets the label for the create button.
func (t InputTag) CreateBtnLabel(value string) InputTag {
	return t.set("createBtnLabel", value)
}

// DeferApi sets the API for deferred loading.
func (t InputTag) DeferApi(value string) InputTag {
	return t.set("deferApi", value)
}

// DeferField sets the field for lazy loading.
func (t InputTag) DeferField(value string) InputTag {
	return t.set("deferField", value)
}

// DeleteApi sets the API for deleting tags.
func (t InputTag) DeleteApi(value string) InputTag {
	return t.set("deleteApi", value)
}

// DeleteConfirmText sets the confirmation text for deleting tags.
func (t InputTag) DeleteConfirmText(value string) InputTag {
	return t.set("deleteConfirmText", value)
}

// Delimiter sets the delimiter for tags.
func (t InputTag) Delimiter(value string) InputTag {
	return t.set("delimiter", value)
}

// Desc sets the description.
func (t InputTag) Desc(value string) InputTag {
	return t.set("desc", value)
}

// Description sets the description content, supporting HTML.
func (t InputTag) Description(value string) InputTag {
	return t.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description.
func (t InputTag) DescriptionClassName(value string) InputTag {
	return t.set("descriptionClassName", value)
}

// Disabled sets whether the input tag is disabled.
func (t InputTag) Disabled(value bool) InputTag {
	return t.set("disabled", value)
}

// DisabledOn sets the expression to determine if the input tag is disabled.
func (t InputTag) DisabledOn(value string) InputTag {
	return t.set("disabledOn", value)
}

// Dropdown sets whether the input tag is in dropdown mode.
func (t InputTag) Dropdown(value bool) InputTag {
	return t.set("dropdown", value)
}

// EditApi sets the API for editing tags.
func (t InputTag) EditApi(value string) InputTag {
	return t.set("editApi", value)
}

// EditControls sets the form items for editing tags.
func (t InputTag) EditControls(value string) InputTag {
	return t.set("editControls", value)
}

// EditDialog sets the configuration for the edit dialog.
func (t InputTag) EditDialog(value string) InputTag {
	return t.set("editDialog", value)
}

// Editable sets whether the input tag is editable.
func (t InputTag) Editable(value bool) InputTag {
	return t.set("editable", value)
}

// EditorSetting sets the editor configuration.
func (t InputTag) EditorSetting(value string) InputTag {
	return t.set("editorSetting", value)
}

// EnableBatchAdd sets whether batch add mode is enabled.
func (t InputTag) EnableBatchAdd(value bool) InputTag {
	return t.set("enableBatchAdd", value)
}

// ExtraName sets the extra field name for range components.
func (t InputTag) ExtraName(value string) InputTag {
	return t.set("extraName", value)
}

// ExtractValue sets whether to extract the value of selected options as an array.
func (t InputTag) ExtractValue(value bool) InputTag {
	return t.set("extractValue", value)
}

// Hidden sets whether the input tag is hidden.
func (t InputTag) Hidden(value bool) InputTag {
	return t.set("hidden", value)
}

// HiddenOn sets the expression to determine if the input tag is hidden.
func (t InputTag) HiddenOn(value string) InputTag {
	return t.set("hiddenOn", value)
}

// Hint sets the input hint displayed on focus.
func (t InputTag) Hint(value string) InputTag {
	return t.set("hint", value)
}

// Horizontal sets the horizontal layout configuration.
func (t InputTag) Horizontal(value string) InputTag {
	return t.set("horizontal", value)
}

// ID sets the unique ID for the component, mainly for logging.
func (t InputTag) ID(value string) InputTag {
	return t.set("id", value)
}

// InitAutoFill sets the initial auto-fill configuration.
func (t InputTag) InitAutoFill(value string) InputTag {
	return t.set("initAutoFill", value)
}

// InitFetch sets whether to fetch the source API initially.
func (t InputTag) InitFetch(value bool) InputTag {
	return t.set("initFetch", value)
}

// InitFetchOn sets the expression to determine if the source API should be fetched initially.
func (t InputTag) InitFetchOn(value string) InputTag {
	return t.set("initFetchOn", value)
}

// Inline sets whether the form control is in inline mode.
func (t InputTag) Inline(value bool) InputTag {
	return t.set("inline", value)
}

// InputClassName sets the CSS class name for the input.
func (t InputTag) InputClassName(value string) InputTag {
	return t.set("inputClassName", value)
}

// JoinValues sets whether to join selected option values with a delimiter.
func (t InputTag) JoinValues(value bool) InputTag {
	return t.set("joinValues", value)
}

// Label sets the label for the input tag.
func (t InputTag) Label(value string) InputTag {
	return t.set("label", value)
}

// LabelAlign sets the alignment for the label. Options: right, left, top, inherit.
func (t InputTag) LabelAlign(value string) InputTag {
	return t.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label.
func (t InputTag) LabelClassName(value string) InputTag {
	return t.set("labelClassName", value)
}

// LabelRemark sets the remark for the label, displayed as a small icon with a tooltip.
func (t InputTag) LabelRemark(value string) InputTag {
	return t.set("labelRemark", value)
}

// LabelWidth sets the custom width for the label, default unit is px.
func (t InputTag) LabelWidth(value string) InputTag {
	return t.set("labelWidth", value)
}

// Max sets the maximum number of tags that can be added.
func (t InputTag) Max(value string) InputTag {
	return t.set("max", value)
}

// MaxTagCount sets the maximum number of tags to display, with overflow shown in a dropdown.
func (t InputTag) MaxTagCount(value string) InputTag {
	return t.set("maxTagCount", value)
}

// MaxTagText sets the text for the tag overflow indicator, default is "...".
func (t InputTag) MaxTagText(value string) InputTag {
	return t.set("maxTagText", value)
}

// Min sets the minimum number of tags that must be selected.
func (t InputTag) Min(value string) InputTag {
	return t.set("min", value)
}

// Mode sets the display mode for the form item.
func (t InputTag) Mode(value string) InputTag {
	return t.set("mode", value)
}

// Name sets the name attribute for the form item.
func (t InputTag) Name(value string) InputTag {
	return t.set("name", value)
}

// OnEvent sets the event handlers for the input tag.
func (t InputTag) OnEvent(value map[string]string) InputTag {
	return t.set("onEvent", value)
}

// Options sets the options for the input tag.
func (t InputTag) Options(value ...any) InputTag {
	return t.set("options", value)
}

// OptionValuePath sets the path to the value in the options.
func (t InputTag) OptionValuePath(value string) InputTag {
	return t.set("optionValuePath", value)
}

// OptionLabelPath sets the path to the label in the options.
func (t InputTag) OptionLabelPath(value string) InputTag {
	return t.set("optionLabelPath", value)
}

// OptionClassName sets the CSS class name for the options.
func (t InputTag) OptionClassName(value string) InputTag {
	return t.set("optionClassName", value)
}

// Placeholder sets the placeholder text for the input.
func (t InputTag) Placeholder(value string) InputTag {
	return t.set("placeholder", value)
}

// ReadOnly sets whether the input tag is read-only.
func (t InputTag) ReadOnly(value bool) InputTag {
	return t.set("readOnly", value)
}

// Required sets whether the input tag is required.
func (t InputTag) Required(value bool) InputTag {
	return t.set("required", value)
}

// Rules sets the validation rules for the input tag.
func (t InputTag) Rules(value string) InputTag {
	return t.set("rules", value)
}

// Source sets the data source for the input tag.
func (t InputTag) Source(value string) InputTag {
	return t.set("source", value)
}

// TagClassName sets the CSS class name for the tags.
func (t InputTag) TagClassName(value string) InputTag {
	return t.set("tagClassName", value)
}

// TagTpl sets the custom template for the tags.
func (t InputTag) TagTpl(value string) InputTag {
	return t.set("tagTpl", value)
}

// Touchable sets whether the input tag supports touch operations.
func (t InputTag) Touchable(value bool) InputTag {
	return t.set("touchable", value)
}

// Value sets the value for the form item.
func (t InputTag) Value(value any) InputTag {
	return t.set("value", value)
}

// VerticalLayout sets whether the input tag uses vertical layout.
func (t InputTag) VerticalLayout(value bool) InputTag {
	return t.set("verticalLayout", value)
}

// Visible sets whether the input tag is visible.
func (t InputTag) Visible(value bool) InputTag {
	return t.set("visible", value)
}

// ValidationErrors sets custom validation error messages.
func (t InputTag) ValidationErrors(value map[string]string) InputTag {
	return t.set("validationErrors", value)
}

// ValuePath sets the path to the value in the data.
func (t InputTag) ValuePath(value string) InputTag {
	return t.set("valuePath", value)
}

// ValidateOnChange sets whether to validate the input tag on change.
func (t InputTag) ValidateOnChange(value bool) InputTag {
	return t.set("validateOnChange", value)
}

// Validations sets the validation rules for the input tag.
func (t InputTag) Validations(value string) InputTag {
	return t.set("validations", value)
}

// ValuesNoWrap sets whether to avoid line wrapping for multiple values.
func (t InputTag) ValuesNoWrap(value bool) InputTag {
	return t.set("valuesNoWrap", value)
}

// VisibleOn sets the expression to determine if the input tag is visible.
func (t InputTag) VisibleOn(value string) InputTag {
	return t.set("visibleOn", value)
}

// Width sets the width of the input tag in a table.
func (t InputTag) Width(value string) InputTag {
	return t.set("width", value)
}
