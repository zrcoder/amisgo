package comp

// treeSelect represents a dropdown tree select component
type treeSelect Schema

func TreeSelect() treeSelect {
	return treeSelect{}.set("type", "tree-select")
}

func (tsc treeSelect) set(key string, value any) treeSelect {
	tsc[key] = value
	return tsc
}

// AddApi sets the API for adding items
func (tsc treeSelect) AddApi(value string) treeSelect {
	return tsc.set("addApi", value)
}

// AddControls sets the form items for adding
func (tsc treeSelect) AddControls(value string) treeSelect {
	return tsc.set("addControls", value)
}

// AddDialog sets the dialog for adding items
func (tsc treeSelect) AddDialog(value string) treeSelect {
	return tsc.set("addDialog", value)
}

// AutoCheckChildren sets whether to auto-check child nodes
func (tsc treeSelect) AutoCheckChildren(value string) treeSelect {
	return tsc.set("autoCheckChildren", value)
}

// AutoFill sets auto-fill
func (tsc treeSelect) AutoFill(value string) treeSelect {
	return tsc.set("autoFill", value)
}

// Cascade sets whether parent and child nodes are independent
func (tsc treeSelect) Cascade(value bool) treeSelect {
	return tsc.set("cascade", value)
}

// ClassName sets the container CSS class name
func (tsc treeSelect) ClassName(value string) treeSelect {
	return tsc.set("className", value)
}

// ClearValueOnHidden sets whether to clear the value when hidden
func (tsc treeSelect) ClearValueOnHidden(value bool) treeSelect {
	return tsc.set("clearValueOnHidden", value)
}

// Clearable sets whether the field is clearable
func (tsc treeSelect) Clearable(value bool) treeSelect {
	return tsc.set("clearable", value)
}

// Creatable sets whether new items can be created
func (tsc treeSelect) Creatable(value bool) treeSelect {
	return tsc.set("creatable", value)
}

// CreateBtnLabel sets the label for the create button
func (tsc treeSelect) CreateBtnLabel(value string) treeSelect {
	return tsc.set("createBtnLabel", value)
}

// DeferApi sets the API for lazy loading
func (tsc treeSelect) DeferApi(value string) treeSelect {
	return tsc.set("deferApi", value)
}

// DeferField sets the field for lazy loading
func (tsc treeSelect) DeferField(value string) treeSelect {
	return tsc.set("deferField", value)
}

// DeleteApi sets the API for deleting items
func (tsc treeSelect) DeleteApi(value string) treeSelect {
	return tsc.set("deleteApi", value)
}

// DeleteConfirmText sets the confirmation text for deletion
func (tsc treeSelect) DeleteConfirmText(value string) treeSelect {
	return tsc.set("deleteConfirmText", value)
}

// Delimiter sets the delimiter
func (tsc treeSelect) Delimiter(value string) treeSelect {
	return tsc.set("delimiter", value)
}

// Desc sets the description
func (tsc treeSelect) Desc(value string) treeSelect {
	return tsc.set("desc", value)
}

// Description sets the description content (supports HTML)
func (tsc treeSelect) Description(value string) treeSelect {
	return tsc.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description
func (tsc treeSelect) DescriptionClassName(value string) treeSelect {
	return tsc.set("descriptionClassName", value)
}

// Disabled sets whether the field is disabled
func (tsc treeSelect) Disabled(value bool) treeSelect {
	return tsc.set("disabled", value)
}

// DisabledOn sets the expression for disabling the field
func (tsc treeSelect) DisabledOn(value string) treeSelect {
	return tsc.set("disabledOn", value)
}

// EditApi sets the API for editing items
func (tsc treeSelect) EditApi(value string) treeSelect {
	return tsc.set("editApi", value)
}

// EditControls sets the form items for editing
func (tsc treeSelect) EditControls(value string) treeSelect {
	return tsc.set("editControls", value)
}

// EditDialog sets the dialog for editing items
func (tsc treeSelect) EditDialog(value string) treeSelect {
	return tsc.set("editDialog", value)
}

// Editable sets whether the field is editable
func (tsc treeSelect) Editable(value bool) treeSelect {
	return tsc.set("editable", value)
}

// EditorSetting sets the editor configuration
func (tsc treeSelect) EditorSetting(value string) treeSelect {
	return tsc.set("editorSetting", value)
}

// EnableDefaultIcon sets whether to add a default icon to options
func (tsc treeSelect) EnableDefaultIcon(value bool) treeSelect {
	return tsc.set("enableDefaultIcon", value)
}

// EnableNodePath sets whether to enable node path mode
func (tsc treeSelect) EnableNodePath(value bool) treeSelect {
	return tsc.set("enableNodePath", value)
}

// ExtraName sets the extra field name
func (tsc treeSelect) ExtraName(value string) treeSelect {
	return tsc.set("extraName", value)
}

// ExtractValue sets whether to wrap selected option values in an array
func (tsc treeSelect) ExtractValue(value bool) treeSelect {
	return tsc.set("extractValue", value)
}

// Hidden sets whether the field is hidden
func (tsc treeSelect) Hidden(value bool) treeSelect {
	return tsc.set("hidden", value)
}

// HiddenOn sets the expression for hiding the field
func (tsc treeSelect) HiddenOn(value string) treeSelect {
	return tsc.set("hiddenOn", value)
}

// HideNodePathLabel sets whether to hide ancestor node text in the select box
func (tsc treeSelect) HideNodePathLabel(value bool) treeSelect {
	return tsc.set("hideNodePathLabel", value)
}

// HideRoot sets whether to hide the root node
func (tsc treeSelect) HideRoot(value bool) treeSelect {
	return tsc.set("hideRoot", value)
}

// Hint sets the input hint
func (tsc treeSelect) Hint(value string) treeSelect {
	return tsc.set("hint", value)
}

// Horizontal sets the horizontal layout configuration
func (tsc treeSelect) Horizontal(value string) treeSelect {
	return tsc.set("horizontal", value)
}

// ID sets the unique component ID
func (tsc treeSelect) ID(value string) treeSelect {
	return tsc.set("id", value)
}

// InitAutoFill sets the initial auto-fill value
func (tsc treeSelect) InitAutoFill(value string) treeSelect {
	return tsc.set("initAutoFill", value)
}

// InitFetch sets whether to fetch the source initially
func (tsc treeSelect) InitFetch(value bool) treeSelect {
	return tsc.set("initFetch", value)
}

// InitFetchOn sets the expression for initial fetch
func (tsc treeSelect) InitFetchOn(value string) treeSelect {
	return tsc.set("initFetchOn", value)
}

// Inline sets whether the form control is inline
func (tsc treeSelect) Inline(value bool) treeSelect {
	return tsc.set("inline", value)
}

// InputClassName sets the input CSS class name
func (tsc treeSelect) InputClassName(value string) treeSelect {
	return tsc.set("inputClassName", value)
}

// JoinValues sets whether to submit the value of the selected option in single select mode
func (tsc treeSelect) JoinValues(value bool) treeSelect {
	return tsc.set("joinValues", value)
}

// Label sets the label
func (tsc treeSelect) Label(value string) treeSelect {
	return tsc.set("label", value)
}

// LabelAlign sets the label alignment
func (tsc treeSelect) LabelAlign(value string) treeSelect {
	return tsc.set("labelAlign", value)
}

// LabelClassName sets the label CSS class name
func (tsc treeSelect) LabelClassName(value string) treeSelect {
	return tsc.set("labelClassName", value)
}

// LabelField sets the label field
func (tsc treeSelect) LabelField(value string) treeSelect {
	return tsc.set("labelField", value)
}

// LabelRemark sets the label remark
func (tsc treeSelect) LabelRemark(value string) treeSelect {
	return tsc.set("labelRemark", value)
}

// LabelWidth sets the custom label width
func (tsc treeSelect) LabelWidth(value string) treeSelect {
	return tsc.set("labelWidth", value)
}

// MaxTagCount sets the maximum number of tags to display
func (tsc treeSelect) MaxTagCount(value string) treeSelect {
	return tsc.set("maxTagCount", value)
}

// MenuTpl sets the custom option template
func (tsc treeSelect) MenuTpl(value string) treeSelect {
	return tsc.set("menuTpl", value)
}

// Multiple sets whether multiple selection is allowed
func (tsc treeSelect) Multiple(value bool) treeSelect {
	return tsc.set("multiple", value)
}

// Name sets the form field name
func (tsc treeSelect) Name(value string) treeSelect {
	return tsc.set("name", value)
}

// Options sets the options
func (tsc treeSelect) Options(value any) treeSelect {
	return tsc.set("options", value)
}

// OptionValues sets the option values
func (tsc treeSelect) OptionValues(value string) treeSelect {
	return tsc.set("optionValues", value)
}

// OptionLabel sets the option label
func (tsc treeSelect) OptionLabel(value string) treeSelect {
	return tsc.set("optionLabel", value)
}

// OptionName sets the option name
func (tsc treeSelect) OptionName(value string) treeSelect {
	return tsc.set("optionName", value)
}

// OptionValue sets the option value
func (tsc treeSelect) OptionValue(value string) treeSelect {
	return tsc.set("optionValue", value)
}

// OptionIcon sets the option icon
func (tsc treeSelect) OptionIcon(value string) treeSelect {
	return tsc.set("optionIcon", value)
}

// Placeholder sets the placeholder text
func (tsc treeSelect) Placeholder(value string) treeSelect {
	return tsc.set("placeholder", value)
}

// ReadOnly sets whether the field is read-only
func (tsc treeSelect) ReadOnly(value bool) treeSelect {
	return tsc.set("readOnly", value)
}

// ReBuildApi sets the API for rebuilding the dropdown
func (tsc treeSelect) ReBuildApi(value string) treeSelect {
	return tsc.set("reBuildApi", value)
}

// Reload sets whether the options can be refreshed via reloadApi
func (tsc treeSelect) Reload(value bool) treeSelect {
	return tsc.set("reload", value)
}

// ReloadApi sets the API for refreshing options
func (tsc treeSelect) ReloadApi(value string) treeSelect {
	return tsc.set("reloadApi", value)
}

// ResetValue sets the reset value for the form
func (tsc treeSelect) ResetValue(value any) treeSelect {
	return tsc.set("resetValue", value)
}

// RootLabel sets the root node label
func (tsc treeSelect) RootLabel(value string) treeSelect {
	return tsc.set("rootLabel", value)
}

// Selectable sets the expression for controlling whether items are selectable
func (tsc treeSelect) Selectable(value string) treeSelect {
	return tsc.set("selectable", value)
}

// ShowIcon sets whether to show icons
func (tsc treeSelect) ShowIcon(value bool) treeSelect {
	return tsc.set("showIcon", value)
}

// ShowRadio sets whether to show radio buttons
func (tsc treeSelect) ShowRadio(value bool) treeSelect {
	return tsc.set("showRadio", value)
}

// ShowSearch sets whether to show the search box
func (tsc treeSelect) ShowSearch(value bool) treeSelect {
	return tsc.set("showSearch", value)
}

// Source sets the data source
func (tsc treeSelect) Source(value string) treeSelect {
	return tsc.set("source", value)
}

// StaticClassName sets the static CSS class name
func (tsc treeSelect) StaticClassName(value string) treeSelect {
	return tsc.set("staticClassName", value)
}

// Static sets whether the field is statically displayed
func (tsc treeSelect) Static(value bool) treeSelect {
	return tsc.set("static", value)
}

// StaticLabel sets the static display label
func (tsc treeSelect) StaticLabel(value string) treeSelect {
	return tsc.set("staticLabel", value)
}

// Taggable sets whether tags can be added
func (tsc treeSelect) Taggable(value bool) treeSelect {
	return tsc.set("taggable", value)
}

// Tags sets whether tags are enabled
func (tsc treeSelect) Tags(value bool) treeSelect {
	return tsc.set("tags", value)
}

// Tip sets the tip text
func (tsc treeSelect) Tip(value string) treeSelect {
	return tsc.set("tip", value)
}

// Value sets the value of the form field
func (tsc treeSelect) Value(value any) treeSelect {
	return tsc.set("value", value)
}

// ValueField sets the value field for options
func (tsc treeSelect) ValueField(value string) treeSelect {
	return tsc.set("valueField", value)
}

// Vertical sets the vertical layout configuration
func (tsc treeSelect) Vertical(value string) treeSelect {
	return tsc.set("vertical", value)
}

// Visible sets whether the field is visible
func (tsc treeSelect) Visible(value bool) treeSelect {
	return tsc.set("visible", value)
}

// VisibleOn sets the expression for visibility
func (tsc treeSelect) VisibleOn(value string) treeSelect {
	return tsc.set("visibleOn", value)
}

// Width sets the width of the component
func (tsc treeSelect) Width(value string) treeSelect {
	return tsc.set("width", value)
}

// Wrapping sets whether to wrap lines
func (tsc treeSelect) Wrapping(value bool) treeSelect {
	return tsc.set("wrapping", value)
}

// Other sets custom fields
func (tsc treeSelect) Other(key string, value any) treeSelect {
	return tsc.set(key, value)
}

// RootCreatable sets whether the root node can create child nodes
func (tsc treeSelect) RootCreatable(value bool) treeSelect {
	return tsc.set("rootCreatable", value)
}

// RootValue sets the value of the root option
func (tsc treeSelect) RootValue(value string) treeSelect {
	return tsc.set("rootValue", value)
}

// Row sets the row value
func (tsc treeSelect) Row(value string) treeSelect {
	return tsc.set("row", value)
}

// SaveImmediately sets whether to save immediately (used in TableColumn)
func (tsc treeSelect) SaveImmediately(value bool) treeSelect {
	return tsc.set("saveImmediately", value)
}

// Searchable sets whether the field is searchable
func (tsc treeSelect) Searchable(value bool) treeSelect {
	return tsc.set("searchable", value)
}

// SelectFirst sets whether to select the first option by default
func (tsc treeSelect) SelectFirst(value bool) treeSelect {
	return tsc.set("selectFirst", value)
}

// ShowOutline sets whether to show the outline
func (tsc treeSelect) ShowOutline(value bool) treeSelect {
	return tsc.set("showOutline", value)
}

// Size sets the size of the form field
func (tsc treeSelect) Size(value string) treeSelect {
	return tsc.set("size", value)
}

// StaticInputClassName sets the CSS class name for the static input
func (tsc treeSelect) StaticInputClassName(value string) treeSelect {
	return tsc.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for the static label
func (tsc treeSelect) StaticLabelClassName(value string) treeSelect {
	return tsc.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display
func (tsc treeSelect) StaticOn(value string) treeSelect {
	return tsc.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (tsc treeSelect) StaticPlaceholder(value string) treeSelect {
	return tsc.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema
func (tsc treeSelect) StaticSchema(value string) treeSelect {
	return tsc.set("staticSchema", value)
}

// Style sets the component style
func (tsc treeSelect) Style(value any) treeSelect {
	return tsc.set("style", value)
}

// SubmitOnChange sets whether to submit the form on change
func (tsc treeSelect) SubmitOnChange(value bool) treeSelect {
	return tsc.set("submitOnChange", value)
}

// TestIdBuilder sets the test ID builder
func (tsc treeSelect) TestIdBuilder(value string) treeSelect {
	return tsc.set("testIdBuilder", value)
}

// UseMobileUI sets whether to use mobile UI
func (tsc treeSelect) UseMobileUI(value bool) treeSelect {
	return tsc.set("useMobileUI", value)
}

// ValidateApi sets the API for remote validation
func (tsc treeSelect) ValidateApi(value string) treeSelect {
	return tsc.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on change
func (tsc treeSelect) ValidateOnChange(value bool) treeSelect {
	return tsc.set("validateOnChange", value)
}

// ValidationErrors sets the validation error messages
func (tsc treeSelect) ValidationErrors(value string) treeSelect {
	return tsc.set("validationErrors", value)
}

// Validations sets the validations
func (tsc treeSelect) Validations(value string) treeSelect {
	return tsc.set("validations", value)
}

// ValuesNoWrap sets whether to avoid wrapping in multi-select mode
func (tsc treeSelect) ValuesNoWrap(value bool) treeSelect {
	return tsc.set("valuesNoWrap", value)
}

// WithChildren sets whether to include child node values when selecting a parent
func (tsc treeSelect) WithChildren(value bool) treeSelect {
	return tsc.set("withChildren", value)
}
