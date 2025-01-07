package comp

// inputTree represents a dropdown selection box. Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/form/tree

type inputTree Schema

// InputTree initializes an inputTree with type "input-tree".
func InputTree() inputTree {
	return inputTree{}.set("type", "input-tree")
}

// set sets a key-value pair in the inputTree.
func (tc inputTree) set(key string, value any) inputTree {
	tc[key] = value
	return tc
}

// AddApi sets the API for adding items.
func (tc inputTree) AddApi(value string) inputTree {
	return tc.set("addApi", value)
}

// AddControls sets the form items for adding new entries.
func (tc inputTree) AddControls(value string) inputTree {
	return tc.set("addControls", value)
}

// AddDialog sets the configuration for the add dialog.
func (tc inputTree) AddDialog(value string) inputTree {
	return tc.set("addDialog", value)
}

// AutoCheckChildren sets whether to auto-check child nodes.
func (tc inputTree) AutoCheckChildren(value bool) inputTree {
	return tc.set("autoCheckChildren", value)
}

// AutoFill sets the auto-fill value when an option is selected.
func (tc inputTree) AutoFill(value string) inputTree {
	return tc.set("autoFill", value)
}

// Cascade sets the data cascade relationship.
func (tc inputTree) Cascade(value bool) inputTree {
	return tc.set("cascade", value)
}

// ClassName sets the CSS class name for the container.
func (tc inputTree) ClassName(value string) inputTree {
	return tc.set("className", value)
}

// ClearValueOnHidden sets whether to clear the value when the form item is hidden.
func (tc inputTree) ClearValueOnHidden(value bool) inputTree {
	return tc.set("clearValueOnHidden", value)
}

// Clearable sets whether the input can be cleared.
func (tc inputTree) Clearable(value bool) inputTree {
	return tc.set("clearable", value)
}

// Creatable sets whether new items can be created.
func (tc inputTree) Creatable(value bool) inputTree {
	return tc.set("creatable", value)
}

// CreateBtnLabel sets the label for the create button.
func (tc inputTree) CreateBtnLabel(value string) inputTree {
	return tc.set("createBtnLabel", value)
}

// DeferApi sets the API for lazy loading.
func (tc inputTree) DeferApi(value string) inputTree {
	return tc.set("deferApi", value)
}

// DeferField sets the field for lazy loading.
func (tc inputTree) DeferField(value string) inputTree {
	return tc.set("deferField", value)
}

// DeleteApi sets the API for deleting items.
func (tc inputTree) DeleteApi(value string) inputTree {
	return tc.set("deleteApi", value)
}

// DeleteConfirmText sets the confirmation text for deleting items.
func (tc inputTree) DeleteConfirmText(value string) inputTree {
	return tc.set("deleteConfirmText", value)
}

// Delimiter sets the delimiter for values.
func (tc inputTree) Delimiter(value string) inputTree {
	return tc.set("delimiter", value)
}

// Desc sets the description.
func (tc inputTree) Desc(value string) inputTree {
	return tc.set("desc", value)
}

// Description sets the description content, supports HTML.
func (tc inputTree) Description(value string) inputTree {
	return tc.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description.
func (tc inputTree) DescriptionClassName(value string) inputTree {
	return tc.set("descriptionClassName", value)
}

// Disabled sets whether the input is disabled.
func (tc inputTree) Disabled(value bool) inputTree {
	return tc.set("disabled", value)
}

// DisabledOn sets the expression to disable the input.
func (tc inputTree) DisabledOn(value string) inputTree {
	return tc.set("disabledOn", value)
}

// EditApi sets the API for editing items.
func (tc inputTree) EditApi(value string) inputTree {
	return tc.set("editApi", value)
}

// EditControls sets the form items for editing entries.
func (tc inputTree) EditControls(value string) inputTree {
	return tc.set("editControls", value)
}

// EditDialog sets the configuration for the edit dialog.
func (tc inputTree) EditDialog(value string) inputTree {
	return tc.set("editDialog", value)
}

// Editable sets whether items can be edited.
func (tc inputTree) Editable(value bool) inputTree {
	return tc.set("editable", value)
}

// EditorSetting sets the editor configuration.
func (tc inputTree) EditorSetting(value string) inputTree {
	return tc.set("editorSetting", value)
}

// EnableDefaultIcon sets whether to add a default icon to options.
func (tc inputTree) EnableDefaultIcon(value bool) inputTree {
	return tc.set("enableDefaultIcon", value)
}

// EnableNodePath sets whether to enable node path mode.
func (tc inputTree) EnableNodePath(value bool) inputTree {
	return tc.set("enableNodePath", value)
}

// ExtraName sets the extra field name for range components.
func (tc inputTree) ExtraName(value string) inputTree {
	return tc.set("extraName", value)
}

// ExtractValue sets whether to wrap selected option values as an array.
func (tc inputTree) ExtractValue(value bool) inputTree {
	return tc.set("extractValue", value)
}

// HeightAuto sets whether the height auto-grows.
func (tc inputTree) HeightAuto(value bool) inputTree {
	return tc.set("heightAuto", value)
}

// Hidden sets whether the input is hidden.
func (tc inputTree) Hidden(value bool) inputTree {
	return tc.set("hidden", value)
}

// HiddenOn sets the expression to hide the input.
func (tc inputTree) HiddenOn(value string) inputTree {
	return tc.set("hiddenOn", value)
}

// HideRoot sets whether to hide the root node.
func (tc inputTree) HideRoot(value bool) inputTree {
	return tc.set("hideRoot", value)
}

// HighlightTxt sets the text to highlight.
func (tc inputTree) HighlightTxt(value string) inputTree {
	return tc.set("highlightTxt", value)
}

// Hint sets the input hint shown on focus.
func (tc inputTree) Hint(value string) inputTree {
	return tc.set("hint", value)
}

// Horizontal sets the horizontal layout configuration.
func (tc inputTree) Horizontal(value string) inputTree {
	return tc.set("horizontal", value)
}

// ID sets the unique component ID for logging.
func (tc inputTree) ID(value string) inputTree {
	return tc.set("id", value)
}

// InitAutoFill sets the initial auto-fill value.
func (tc inputTree) InitAutoFill(value string) inputTree {
	return tc.set("initAutoFill", value)
}

// InitFetch sets whether to fetch the source API initially.
func (tc inputTree) InitFetch(value bool) inputTree {
	return tc.set("initFetch", value)
}

// InitFetchOn sets the expression to fetch the source API initially.
func (tc inputTree) InitFetchOn(value string) inputTree {
	return tc.set("initFetchOn", value)
}

// Inline sets whether the form control is in inline mode.
func (tc inputTree) Inline(value bool) inputTree {
	return tc.set("inline", value)
}

// InputClassName sets the CSS class name for the input.
func (tc inputTree) InputClassName(value string) inputTree {
	return tc.set("inputClassName", value)
}

// JoinValues sets whether to join values with a delimiter in single/multiple mode.
func (tc inputTree) JoinValues(value bool) inputTree {
	return tc.set("joinValues", value)
}

// Label sets the label text.
func (tc inputTree) Label(value string) inputTree {
	return tc.set("label", value)
}

// LabelClassName sets the CSS class name for the label.
func (tc inputTree) LabelClassName(value string) inputTree {
	return tc.set("labelClassName", value)
}

// LabelRemark sets the label remark.
func (tc inputTree) LabelRemark(value string) inputTree {
	return tc.set("labelRemark", value)
}

// LabelWidth sets the label width.
func (tc inputTree) LabelWidth(value string) inputTree {
	return tc.set("labelWidth", value)
}

// MaxLength sets the maximum length.
func (tc inputTree) MaxLength(value int) inputTree {
	return tc.set("maxLength", value)
}

// MinLength sets the minimum length.
func (tc inputTree) MinLength(value int) inputTree {
	return tc.set("minLength", value)
}

// Name sets the form item name.
func (tc inputTree) Name(value string) inputTree {
	return tc.set("name", value)
}

// NodeValue sets the node value for the selected option.
func (tc inputTree) NodeValue(value string) inputTree {
	return tc.set("nodeValue", value)
}

// Options sets the options.
func (tc inputTree) Options(value ...any) inputTree {
	return tc.set("options", value)
}

// Placeholder sets the placeholder text.
func (tc inputTree) Placeholder(value string) inputTree {
	return tc.set("placeholder", value)
}

// ReadOnly sets whether the input is read-only.
func (tc inputTree) ReadOnly(value bool) inputTree {
	return tc.set("readOnly", value)
}

// RootLabel sets the root label.
func (tc inputTree) RootLabel(value string) inputTree {
	return tc.set("rootLabel", value)
}

// SelectMode sets the selection mode (multiple, tags, single).
func (tc inputTree) SelectMode(value string) inputTree {
	return tc.set("selectMode", value)
}

// ShowIcon sets whether to show icons.
func (tc inputTree) ShowIcon(value bool) inputTree {
	return tc.set("showIcon", value)
}

// ShowSearch sets whether to show the search box.
func (tc inputTree) ShowSearch(value bool) inputTree {
	return tc.set("showSearch", value)
}

// ShowTreeIcon sets whether to show tree icons.
func (tc inputTree) ShowTreeIcon(value bool) inputTree {
	return tc.set("showTreeIcon", value)
}

// Size sets the form item size.
func (tc inputTree) Size(value string) inputTree {
	return tc.set("size", value)
}

// Source sets the source for options.
func (tc inputTree) Source(value string) inputTree {
	return tc.set("source", value)
}

// StrictMode sets whether to enable strict mode.
func (tc inputTree) StrictMode(value bool) inputTree {
	return tc.set("strictMode", value)
}

// SwitcherIcon sets the switcher icon.
func (tc inputTree) SwitcherIcon(value string) inputTree {
	return tc.set("switcherIcon", value)
}

// Tips sets the tips text.
func (tc inputTree) Tips(value string) inputTree {
	return tc.set("tips", value)
}

// Value sets the form value.
func (tc inputTree) Value(value any) inputTree {
	return tc.set("value", value)
}

// ValueField sets the value field for options.
func (tc inputTree) ValueField(value string) inputTree {
	return tc.set("valueField", value)
}

// Virtualized sets whether to enable virtual lists.
func (tc inputTree) Virtualized(value bool) inputTree {
	return tc.set("virtualized", value)
}

// Width sets the component width.
func (tc inputTree) Width(value string) inputTree {
	return tc.set("width", value)
}

// WidthUnit sets the component width unit.
func (tc inputTree) WidthUnit(value string) inputTree {
	return tc.set("widthUnit", value)
}

// StaticClassName sets the static display class name.
func (tc inputTree) StaticClassName(value string) inputTree {
	return tc.set("staticClassName", value)
}

// StaticInputClassName sets the static display input class name.
func (tc inputTree) StaticInputClassName(value string) inputTree {
	return tc.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static display label class name.
func (tc inputTree) StaticLabelClassName(value string) inputTree {
	return tc.set("staticLabelClassName", value)
}

// StaticOn sets the static display expression.
func (tc inputTree) StaticOn(value string) inputTree {
	return tc.set("staticOn", value)
}

// StaticPlaceholder sets the static display placeholder.
func (tc inputTree) StaticPlaceholder(value string) inputTree {
	return tc.set("staticPlaceholder", value)
}

// StaticSchema sets the static display schema.
func (tc inputTree) StaticSchema(value string) inputTree {
	return tc.set("staticSchema", value)
}

// Style sets the component style.
func (tc inputTree) Style(value any) inputTree {
	return tc.set("style", value)
}

// SubmitOnChange sets whether to submit the form on change.
func (tc inputTree) SubmitOnChange(value bool) inputTree {
	return tc.set("submitOnChange", value)
}

// TestIdBuilder sets the test ID builder.
func (tc inputTree) TestIdBuilder(value string) inputTree {
	return tc.set("testIdBuilder", value)
}

// UseMobileUI sets whether to disable mobile UI styles.
func (tc inputTree) UseMobileUI(value bool) inputTree {
	return tc.set("useMobileUI", value)
}

// ValidateApi sets the remote validation API.
func (tc inputTree) ValidateApi(value string) inputTree {
	return tc.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on change after form submission.
func (tc inputTree) ValidateOnChange(value bool) inputTree {
	return tc.set("validateOnChange", value)
}

// ValidationErrors sets the validation error messages.
func (tc inputTree) ValidationErrors(value string) inputTree {
	return tc.set("validationErrors", value)
}

// Validations sets the validation rules.
func (tc inputTree) Validations(value string) inputTree {
	return tc.set("validations", value)
}

// ValuesNoWrap sets whether to avoid wrapping values in multi-select mode.
func (tc inputTree) ValuesNoWrap(value bool) inputTree {
	return tc.set("valuesNoWrap", value)
}

// Visible sets whether the input is visible.
func (tc inputTree) Visible(value bool) inputTree {
	return tc.set("visible", value)
}

// VisibleOn sets the expression to show the input.
func (tc inputTree) VisibleOn(value string) inputTree {
	return tc.set("visibleOn", value)
}

// WithChildren sets whether to include child node values when selecting a parent.
func (tc inputTree) WithChildren(value bool) inputTree {
	return tc.set("withChildren", value)
}
