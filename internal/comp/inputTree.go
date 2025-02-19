package comp

import "github.com/zrcoder/amisgo/schema"

// InputTree represents a dropdown selection box. Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/form/tree
type InputTree schema.Schema

func NewInputTree() InputTree {
	return InputTree{"type": "input-tree"}
}

// set sets a key-value pair in the inputTree.
func (tc InputTree) set(key string, value any) InputTree {
	tc[key] = value
	return tc
}

// AddApi sets the API for adding items.
func (tc InputTree) AddApi(value string) InputTree {
	return tc.set("addApi", value)
}

// AddControls sets the form items for adding new entries.
func (tc InputTree) AddControls(value string) InputTree {
	return tc.set("addControls", value)
}

// AddDialog sets the configuration for the add dialog.
func (tc InputTree) AddDialog(value string) InputTree {
	return tc.set("addDialog", value)
}

// AutoCheckChildren sets whether to auto-check child nodes.
func (tc InputTree) AutoCheckChildren(value bool) InputTree {
	return tc.set("autoCheckChildren", value)
}

// AutoFill sets the auto-fill value when an option is selected.
func (tc InputTree) AutoFill(value string) InputTree {
	return tc.set("autoFill", value)
}

// Cascade sets the data cascade relationship.
func (tc InputTree) Cascade(value bool) InputTree {
	return tc.set("cascade", value)
}

// ClassName sets the CSS class name for the container.
func (tc InputTree) ClassName(value string) InputTree {
	return tc.set("className", value)
}

// ClearValueOnHidden sets whether to clear the value when the form item is hidden.
func (tc InputTree) ClearValueOnHidden(value bool) InputTree {
	return tc.set("clearValueOnHidden", value)
}

// Clearable sets whether the input can be cleared.
func (tc InputTree) Clearable(value bool) InputTree {
	return tc.set("clearable", value)
}

// Creatable sets whether new items can be created.
func (tc InputTree) Creatable(value bool) InputTree {
	return tc.set("creatable", value)
}

// CreateBtnLabel sets the label for the create button.
func (tc InputTree) CreateBtnLabel(value string) InputTree {
	return tc.set("createBtnLabel", value)
}

// DeferApi sets the API for lazy loading.
func (tc InputTree) DeferApi(value string) InputTree {
	return tc.set("deferApi", value)
}

// DeferField sets the field for lazy loading.
func (tc InputTree) DeferField(value string) InputTree {
	return tc.set("deferField", value)
}

// DeleteApi sets the API for deleting items.
func (tc InputTree) DeleteApi(value string) InputTree {
	return tc.set("deleteApi", value)
}

// DeleteConfirmText sets the confirmation text for deleting items.
func (tc InputTree) DeleteConfirmText(value string) InputTree {
	return tc.set("deleteConfirmText", value)
}

// Delimiter sets the delimiter for values.
func (tc InputTree) Delimiter(value string) InputTree {
	return tc.set("delimiter", value)
}

// Desc sets the description.
func (tc InputTree) Desc(value string) InputTree {
	return tc.set("desc", value)
}

// Description sets the description content, supports HTML.
func (tc InputTree) Description(value string) InputTree {
	return tc.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description.
func (tc InputTree) DescriptionClassName(value string) InputTree {
	return tc.set("descriptionClassName", value)
}

// Disabled sets whether the input is disabled.
func (tc InputTree) Disabled(value bool) InputTree {
	return tc.set("disabled", value)
}

// DisabledOn sets the expression to disable the input.
func (tc InputTree) DisabledOn(value string) InputTree {
	return tc.set("disabledOn", value)
}

// EditApi sets the API for editing items.
func (tc InputTree) EditApi(value string) InputTree {
	return tc.set("editApi", value)
}

// EditControls sets the form items for editing entries.
func (tc InputTree) EditControls(value string) InputTree {
	return tc.set("editControls", value)
}

// EditDialog sets the configuration for the edit dialog.
func (tc InputTree) EditDialog(value string) InputTree {
	return tc.set("editDialog", value)
}

// Editable sets whether items can be edited.
func (tc InputTree) Editable(value bool) InputTree {
	return tc.set("editable", value)
}

// EditorSetting sets the editor configuration.
func (tc InputTree) EditorSetting(value string) InputTree {
	return tc.set("editorSetting", value)
}

// EnableDefaultIcon sets whether to add a default icon to options.
func (tc InputTree) EnableDefaultIcon(value bool) InputTree {
	return tc.set("enableDefaultIcon", value)
}

// EnableNodePath sets whether to enable node path mode.
func (tc InputTree) EnableNodePath(value bool) InputTree {
	return tc.set("enableNodePath", value)
}

// ExtraName sets the extra field name for range components.
func (tc InputTree) ExtraName(value string) InputTree {
	return tc.set("extraName", value)
}

// ExtractValue sets whether to wrap selected option values as an array.
func (tc InputTree) ExtractValue(value bool) InputTree {
	return tc.set("extractValue", value)
}

// HeightAuto sets whether the height auto-grows.
func (tc InputTree) HeightAuto(value bool) InputTree {
	return tc.set("heightAuto", value)
}

// Hidden sets whether the input is hidden.
func (tc InputTree) Hidden(value bool) InputTree {
	return tc.set("hidden", value)
}

// HiddenOn sets the expression to hide the input.
func (tc InputTree) HiddenOn(value string) InputTree {
	return tc.set("hiddenOn", value)
}

// HideRoot sets whether to hide the root node.
func (tc InputTree) HideRoot(value bool) InputTree {
	return tc.set("hideRoot", value)
}

// HighlightTxt sets the text to highlight.
func (tc InputTree) HighlightTxt(value string) InputTree {
	return tc.set("highlightTxt", value)
}

// Hint sets the input hint shown on focus.
func (tc InputTree) Hint(value string) InputTree {
	return tc.set("hint", value)
}

// Horizontal sets the horizontal layout configuration.
func (tc InputTree) Horizontal(value string) InputTree {
	return tc.set("horizontal", value)
}

// ID sets the unique component ID for logging.
func (tc InputTree) ID(value string) InputTree {
	return tc.set("id", value)
}

// InitAutoFill sets the initial auto-fill value.
func (tc InputTree) InitAutoFill(value string) InputTree {
	return tc.set("initAutoFill", value)
}

// InitFetch sets whether to fetch the source API initially.
func (tc InputTree) InitFetch(value bool) InputTree {
	return tc.set("initFetch", value)
}

// InitFetchOn sets the expression to fetch the source API initially.
func (tc InputTree) InitFetchOn(value string) InputTree {
	return tc.set("initFetchOn", value)
}

// Inline sets whether the form control is in inline mode.
func (tc InputTree) Inline(value bool) InputTree {
	return tc.set("inline", value)
}

// InputClassName sets the CSS class name for the input.
func (tc InputTree) InputClassName(value string) InputTree {
	return tc.set("inputClassName", value)
}

// JoinValues sets whether to join values with a delimiter in single/multiple mode.
func (tc InputTree) JoinValues(value bool) InputTree {
	return tc.set("joinValues", value)
}

// Label sets the label text.
func (tc InputTree) Label(value string) InputTree {
	return tc.set("label", value)
}

// LabelClassName sets the CSS class name for the label.
func (tc InputTree) LabelClassName(value string) InputTree {
	return tc.set("labelClassName", value)
}

// LabelRemark sets the label remark.
func (tc InputTree) LabelRemark(value string) InputTree {
	return tc.set("labelRemark", value)
}

// LabelWidth sets the label width.
func (tc InputTree) LabelWidth(value string) InputTree {
	return tc.set("labelWidth", value)
}

// MaxLength sets the maximum length.
func (tc InputTree) MaxLength(value int) InputTree {
	return tc.set("maxLength", value)
}

// MinLength sets the minimum length.
func (tc InputTree) MinLength(value int) InputTree {
	return tc.set("minLength", value)
}

// Name sets the form item name.
func (tc InputTree) Name(value string) InputTree {
	return tc.set("name", value)
}

// NodeValue sets the node value for the selected option.
func (tc InputTree) NodeValue(value string) InputTree {
	return tc.set("nodeValue", value)
}

// Options sets the options.
func (tc InputTree) Options(value ...any) InputTree {
	return tc.set("options", value)
}

// Placeholder sets the placeholder text.
func (tc InputTree) Placeholder(value string) InputTree {
	return tc.set("placeholder", value)
}

// ReadOnly sets whether the input is read-only.
func (tc InputTree) ReadOnly(value bool) InputTree {
	return tc.set("readOnly", value)
}

// RootLabel sets the root label.
func (tc InputTree) RootLabel(value string) InputTree {
	return tc.set("rootLabel", value)
}

// SelectMode sets the selection mode (multiple, tags, single).
func (tc InputTree) SelectMode(value string) InputTree {
	return tc.set("selectMode", value)
}

// ShowIcon sets whether to show icons.
func (tc InputTree) ShowIcon(value bool) InputTree {
	return tc.set("showIcon", value)
}

// ShowSearch sets whether to show the search box.
func (tc InputTree) ShowSearch(value bool) InputTree {
	return tc.set("showSearch", value)
}

// ShowTreeIcon sets whether to show tree icons.
func (tc InputTree) ShowTreeIcon(value bool) InputTree {
	return tc.set("showTreeIcon", value)
}

// Size sets the form item size.
func (tc InputTree) Size(value string) InputTree {
	return tc.set("size", value)
}

// Source sets the source for options.
func (tc InputTree) Source(value string) InputTree {
	return tc.set("source", value)
}

// StrictMode sets whether to enable strict mode.
func (tc InputTree) StrictMode(value bool) InputTree {
	return tc.set("strictMode", value)
}

// SwitcherIcon sets the switcher icon.
func (tc InputTree) SwitcherIcon(value string) InputTree {
	return tc.set("switcherIcon", value)
}

// Tips sets the tips text.
func (tc InputTree) Tips(value string) InputTree {
	return tc.set("tips", value)
}

// Value sets the form value.
func (tc InputTree) Value(value any) InputTree {
	return tc.set("value", value)
}

// ValueField sets the value field for options.
func (tc InputTree) ValueField(value string) InputTree {
	return tc.set("valueField", value)
}

// Virtualized sets whether to enable virtual lists.
func (tc InputTree) Virtualized(value bool) InputTree {
	return tc.set("virtualized", value)
}

// Width sets the component width.
func (tc InputTree) Width(value string) InputTree {
	return tc.set("width", value)
}

// WidthUnit sets the component width unit.
func (tc InputTree) WidthUnit(value string) InputTree {
	return tc.set("widthUnit", value)
}

// StaticClassName sets the static display class name.
func (tc InputTree) StaticClassName(value string) InputTree {
	return tc.set("staticClassName", value)
}

// StaticInputClassName sets the static display input class name.
func (tc InputTree) StaticInputClassName(value string) InputTree {
	return tc.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static display label class name.
func (tc InputTree) StaticLabelClassName(value string) InputTree {
	return tc.set("staticLabelClassName", value)
}

// StaticOn sets the static display expression.
func (tc InputTree) StaticOn(value string) InputTree {
	return tc.set("staticOn", value)
}

// StaticPlaceholder sets the static display placeholder.
func (tc InputTree) StaticPlaceholder(value string) InputTree {
	return tc.set("staticPlaceholder", value)
}

// StaticSchema sets the static display schema.Schema.
func (tc InputTree) StaticSchema(value string) InputTree {
	return tc.set("staticSchema", value)
}

// Style sets the component style.
func (tc InputTree) Style(value any) InputTree {
	return tc.set("style", value)
}

// SubmitOnChange sets whether to submit the form on change.
func (tc InputTree) SubmitOnChange(value bool) InputTree {
	return tc.set("submitOnChange", value)
}

// TestIdBuilder sets the test ID builder.
func (tc InputTree) TestIdBuilder(value string) InputTree {
	return tc.set("testIdBuilder", value)
}

// UseMobileUI sets whether to disable mobile UI styles.
func (tc InputTree) UseMobileUI(value bool) InputTree {
	return tc.set("useMobileUI", value)
}

// ValidateApi sets the remote validation API.
func (tc InputTree) ValidateApi(value string) InputTree {
	return tc.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on change after form submission.
func (tc InputTree) ValidateOnChange(value bool) InputTree {
	return tc.set("validateOnChange", value)
}

// ValidationErrors sets the validation error messages.
func (tc InputTree) ValidationErrors(value string) InputTree {
	return tc.set("validationErrors", value)
}

// Validations sets the validation rules.
func (tc InputTree) Validations(value string) InputTree {
	return tc.set("validations", value)
}

// ValuesNoWrap sets whether to avoid wrapping values in multi-select mode.
func (tc InputTree) ValuesNoWrap(value bool) InputTree {
	return tc.set("valuesNoWrap", value)
}

// Visible sets whether the input is visible.
func (tc InputTree) Visible(value bool) InputTree {
	return tc.set("visible", value)
}

// VisibleOn sets the expression to show the input.
func (tc InputTree) VisibleOn(value string) InputTree {
	return tc.set("visibleOn", value)
}

// WithChildren sets whether to include child node values when selecting a parent.
func (tc InputTree) WithChildren(value bool) InputTree {
	return tc.set("withChildren", value)
}
