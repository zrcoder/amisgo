package comp

import "github.com/zrcoder/amisgo/model"

// TreeSelect represents a dropdown tree select component
type TreeSelect model.Schema

func NewTreeSelect() TreeSelect {
	return TreeSelect{"type": "tree-select"}
}

func (tsc TreeSelect) set(key string, value any) TreeSelect {
	tsc[key] = value
	return tsc
}

// AddApi sets the API for adding items
func (tsc TreeSelect) AddApi(value string) TreeSelect {
	return tsc.set("addApi", value)
}

// AddControls sets the form items for adding
func (tsc TreeSelect) AddControls(value string) TreeSelect {
	return tsc.set("addControls", value)
}

// AddDialog sets the dialog for adding items
func (tsc TreeSelect) AddDialog(value string) TreeSelect {
	return tsc.set("addDialog", value)
}

// AutoCheckChildren sets whether to auto-check child nodes
func (tsc TreeSelect) AutoCheckChildren(value string) TreeSelect {
	return tsc.set("autoCheckChildren", value)
}

// AutoFill sets auto-fill
func (tsc TreeSelect) AutoFill(value string) TreeSelect {
	return tsc.set("autoFill", value)
}

// Cascade sets whether parent and child nodes are independent
func (tsc TreeSelect) Cascade(value bool) TreeSelect {
	return tsc.set("cascade", value)
}

// ClassName sets the container CSS class name
func (tsc TreeSelect) ClassName(value string) TreeSelect {
	return tsc.set("className", value)
}

// ClearValueOnHidden sets whether to clear the value when hidden
func (tsc TreeSelect) ClearValueOnHidden(value bool) TreeSelect {
	return tsc.set("clearValueOnHidden", value)
}

// Clearable sets whether the field is clearable
func (tsc TreeSelect) Clearable(value bool) TreeSelect {
	return tsc.set("clearable", value)
}

// Creatable sets whether new items can be created
func (tsc TreeSelect) Creatable(value bool) TreeSelect {
	return tsc.set("creatable", value)
}

// CreateBtnLabel sets the label for the create button
func (tsc TreeSelect) CreateBtnLabel(value string) TreeSelect {
	return tsc.set("createBtnLabel", value)
}

// DeferApi sets the API for lazy loading
func (tsc TreeSelect) DeferApi(value string) TreeSelect {
	return tsc.set("deferApi", value)
}

// DeferField sets the field for lazy loading
func (tsc TreeSelect) DeferField(value string) TreeSelect {
	return tsc.set("deferField", value)
}

// DeleteApi sets the API for deleting items
func (tsc TreeSelect) DeleteApi(value string) TreeSelect {
	return tsc.set("deleteApi", value)
}

// DeleteConfirmText sets the confirmation text for deletion
func (tsc TreeSelect) DeleteConfirmText(value string) TreeSelect {
	return tsc.set("deleteConfirmText", value)
}

// Delimiter sets the delimiter
func (tsc TreeSelect) Delimiter(value string) TreeSelect {
	return tsc.set("delimiter", value)
}

// Desc sets the description
func (tsc TreeSelect) Desc(value string) TreeSelect {
	return tsc.set("desc", value)
}

// Description sets the description content (supports HTML)
func (tsc TreeSelect) Description(value string) TreeSelect {
	return tsc.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description
func (tsc TreeSelect) DescriptionClassName(value string) TreeSelect {
	return tsc.set("descriptionClassName", value)
}

// Disabled sets whether the field is disabled
func (tsc TreeSelect) Disabled(value bool) TreeSelect {
	return tsc.set("disabled", value)
}

// DisabledOn sets the expression for disabling the field
func (tsc TreeSelect) DisabledOn(value string) TreeSelect {
	return tsc.set("disabledOn", value)
}

// EditApi sets the API for editing items
func (tsc TreeSelect) EditApi(value string) TreeSelect {
	return tsc.set("editApi", value)
}

// EditControls sets the form items for editing
func (tsc TreeSelect) EditControls(value string) TreeSelect {
	return tsc.set("editControls", value)
}

// EditDialog sets the dialog for editing items
func (tsc TreeSelect) EditDialog(value string) TreeSelect {
	return tsc.set("editDialog", value)
}

// Editable sets whether the field is editable
func (tsc TreeSelect) Editable(value bool) TreeSelect {
	return tsc.set("editable", value)
}

// EditorSetting sets the editor configuration
func (tsc TreeSelect) EditorSetting(value string) TreeSelect {
	return tsc.set("editorSetting", value)
}

// EnableDefaultIcon sets whether to add a default icon to options
func (tsc TreeSelect) EnableDefaultIcon(value bool) TreeSelect {
	return tsc.set("enableDefaultIcon", value)
}

// EnableNodePath sets whether to enable node path mode
func (tsc TreeSelect) EnableNodePath(value bool) TreeSelect {
	return tsc.set("enableNodePath", value)
}

// ExtraName sets the extra field name
func (tsc TreeSelect) ExtraName(value string) TreeSelect {
	return tsc.set("extraName", value)
}

// ExtractValue sets whether to wrap selected option values in an array
func (tsc TreeSelect) ExtractValue(value bool) TreeSelect {
	return tsc.set("extractValue", value)
}

// Hidden sets whether the field is hidden
func (tsc TreeSelect) Hidden(value bool) TreeSelect {
	return tsc.set("hidden", value)
}

// HiddenOn sets the expression for hiding the field
func (tsc TreeSelect) HiddenOn(value string) TreeSelect {
	return tsc.set("hiddenOn", value)
}

// HideNodePathLabel sets whether to hide ancestor node text in the select box
func (tsc TreeSelect) HideNodePathLabel(value bool) TreeSelect {
	return tsc.set("hideNodePathLabel", value)
}

// HideRoot sets whether to hide the root node
func (tsc TreeSelect) HideRoot(value bool) TreeSelect {
	return tsc.set("hideRoot", value)
}

// Hint sets the input hint
func (tsc TreeSelect) Hint(value string) TreeSelect {
	return tsc.set("hint", value)
}

// Horizontal sets the horizontal layout configuration
func (tsc TreeSelect) Horizontal(value string) TreeSelect {
	return tsc.set("horizontal", value)
}

// ID sets the unique component ID
func (tsc TreeSelect) ID(value string) TreeSelect {
	return tsc.set("id", value)
}

// InitAutoFill sets the initial auto-fill value
func (tsc TreeSelect) InitAutoFill(value string) TreeSelect {
	return tsc.set("initAutoFill", value)
}

// InitFetch sets whether to fetch the source initially
func (tsc TreeSelect) InitFetch(value bool) TreeSelect {
	return tsc.set("initFetch", value)
}

// InitFetchOn sets the expression for initial fetch
func (tsc TreeSelect) InitFetchOn(value string) TreeSelect {
	return tsc.set("initFetchOn", value)
}

// Inline sets whether the form control is inline
func (tsc TreeSelect) Inline(value bool) TreeSelect {
	return tsc.set("inline", value)
}

// InputClassName sets the input CSS class name
func (tsc TreeSelect) InputClassName(value string) TreeSelect {
	return tsc.set("inputClassName", value)
}

// JoinValues sets whether to submit the value of the selected option in single select mode
func (tsc TreeSelect) JoinValues(value bool) TreeSelect {
	return tsc.set("joinValues", value)
}

// Label sets the label
func (tsc TreeSelect) Label(value string) TreeSelect {
	return tsc.set("label", value)
}

// LabelAlign sets the label alignment
func (tsc TreeSelect) LabelAlign(value string) TreeSelect {
	return tsc.set("labelAlign", value)
}

// LabelClassName sets the label CSS class name
func (tsc TreeSelect) LabelClassName(value string) TreeSelect {
	return tsc.set("labelClassName", value)
}

// LabelField sets the label field
func (tsc TreeSelect) LabelField(value string) TreeSelect {
	return tsc.set("labelField", value)
}

// LabelRemark sets the label remark
func (tsc TreeSelect) LabelRemark(value string) TreeSelect {
	return tsc.set("labelRemark", value)
}

// LabelWidth sets the custom label width
func (tsc TreeSelect) LabelWidth(value string) TreeSelect {
	return tsc.set("labelWidth", value)
}

// MaxTagCount sets the maximum number of tags to display
func (tsc TreeSelect) MaxTagCount(value string) TreeSelect {
	return tsc.set("maxTagCount", value)
}

// MenuTpl sets the custom option template
func (tsc TreeSelect) MenuTpl(value string) TreeSelect {
	return tsc.set("menuTpl", value)
}

// Multiple sets whether multiple selection is allowed
func (tsc TreeSelect) Multiple(value bool) TreeSelect {
	return tsc.set("multiple", value)
}

// Name sets the form field name
func (tsc TreeSelect) Name(value string) TreeSelect {
	return tsc.set("name", value)
}

// Options sets the options
func (tsc TreeSelect) Options(value any) TreeSelect {
	return tsc.set("options", value)
}

// OptionValues sets the option values
func (tsc TreeSelect) OptionValues(value string) TreeSelect {
	return tsc.set("optionValues", value)
}

// OptionLabel sets the option label
func (tsc TreeSelect) OptionLabel(value string) TreeSelect {
	return tsc.set("optionLabel", value)
}

// OptionName sets the option name
func (tsc TreeSelect) OptionName(value string) TreeSelect {
	return tsc.set("optionName", value)
}

// OptionValue sets the option value
func (tsc TreeSelect) OptionValue(value string) TreeSelect {
	return tsc.set("optionValue", value)
}

// OptionIcon sets the option icon
func (tsc TreeSelect) OptionIcon(value string) TreeSelect {
	return tsc.set("optionIcon", value)
}

// Placeholder sets the placeholder text
func (tsc TreeSelect) Placeholder(value string) TreeSelect {
	return tsc.set("placeholder", value)
}

// ReadOnly sets whether the field is read-only
func (tsc TreeSelect) ReadOnly(value bool) TreeSelect {
	return tsc.set("readOnly", value)
}

// ReBuildApi sets the API for rebuilding the dropdown
func (tsc TreeSelect) ReBuildApi(value string) TreeSelect {
	return tsc.set("reBuildApi", value)
}

// Reload sets whether the options can be refreshed via reloadApi
func (tsc TreeSelect) Reload(value bool) TreeSelect {
	return tsc.set("reload", value)
}

// ReloadApi sets the API for refreshing options
func (tsc TreeSelect) ReloadApi(value string) TreeSelect {
	return tsc.set("reloadApi", value)
}

// ResetValue sets the reset value for the form
func (tsc TreeSelect) ResetValue(value any) TreeSelect {
	return tsc.set("resetValue", value)
}

// RootLabel sets the root node label
func (tsc TreeSelect) RootLabel(value string) TreeSelect {
	return tsc.set("rootLabel", value)
}

// Selectable sets the expression for controlling whether items are selectable
func (tsc TreeSelect) Selectable(value string) TreeSelect {
	return tsc.set("selectable", value)
}

// ShowIcon sets whether to show icons
func (tsc TreeSelect) ShowIcon(value bool) TreeSelect {
	return tsc.set("showIcon", value)
}

// ShowRadio sets whether to show radio buttons
func (tsc TreeSelect) ShowRadio(value bool) TreeSelect {
	return tsc.set("showRadio", value)
}

// ShowSearch sets whether to show the search box
func (tsc TreeSelect) ShowSearch(value bool) TreeSelect {
	return tsc.set("showSearch", value)
}

// Source sets the data source
func (tsc TreeSelect) Source(value string) TreeSelect {
	return tsc.set("source", value)
}

// StaticClassName sets the static CSS class name
func (tsc TreeSelect) StaticClassName(value string) TreeSelect {
	return tsc.set("staticClassName", value)
}

// Static sets whether the field is statically displayed
func (tsc TreeSelect) Static(value bool) TreeSelect {
	return tsc.set("static", value)
}

// StaticLabel sets the static display label
func (tsc TreeSelect) StaticLabel(value string) TreeSelect {
	return tsc.set("staticLabel", value)
}

// Taggable sets whether tags can be added
func (tsc TreeSelect) Taggable(value bool) TreeSelect {
	return tsc.set("taggable", value)
}

// Tags sets whether tags are enabled
func (tsc TreeSelect) Tags(value bool) TreeSelect {
	return tsc.set("tags", value)
}

// Tip sets the tip text
func (tsc TreeSelect) Tip(value string) TreeSelect {
	return tsc.set("tip", value)
}

// Value sets the value of the form field
func (tsc TreeSelect) Value(value any) TreeSelect {
	return tsc.set("value", value)
}

// ValueField sets the value field for options
func (tsc TreeSelect) ValueField(value string) TreeSelect {
	return tsc.set("valueField", value)
}

// Vertical sets the vertical layout configuration
func (tsc TreeSelect) Vertical(value string) TreeSelect {
	return tsc.set("vertical", value)
}

// Visible sets whether the field is visible
func (tsc TreeSelect) Visible(value bool) TreeSelect {
	return tsc.set("visible", value)
}

// VisibleOn sets the expression for visibility
func (tsc TreeSelect) VisibleOn(value string) TreeSelect {
	return tsc.set("visibleOn", value)
}

// Width sets the width of the component
func (tsc TreeSelect) Width(value string) TreeSelect {
	return tsc.set("width", value)
}

// Wrapping sets whether to wrap lines
func (tsc TreeSelect) Wrapping(value bool) TreeSelect {
	return tsc.set("wrapping", value)
}

// Other sets custom fields
func (tsc TreeSelect) Other(key string, value any) TreeSelect {
	return tsc.set(key, value)
}

// RootCreatable sets whether the root node can create child nodes
func (tsc TreeSelect) RootCreatable(value bool) TreeSelect {
	return tsc.set("rootCreatable", value)
}

// RootValue sets the value of the root option
func (tsc TreeSelect) RootValue(value string) TreeSelect {
	return tsc.set("rootValue", value)
}

// Row sets the row value
func (tsc TreeSelect) Row(value string) TreeSelect {
	return tsc.set("row", value)
}

// SaveImmediately sets whether to save immediately (used in TableColumn)
func (tsc TreeSelect) SaveImmediately(value bool) TreeSelect {
	return tsc.set("saveImmediately", value)
}

// Searchable sets whether the field is searchable
func (tsc TreeSelect) Searchable(value bool) TreeSelect {
	return tsc.set("searchable", value)
}

// SelectFirst sets whether to select the first option by default
func (tsc TreeSelect) SelectFirst(value bool) TreeSelect {
	return tsc.set("selectFirst", value)
}

// ShowOutline sets whether to show the outline
func (tsc TreeSelect) ShowOutline(value bool) TreeSelect {
	return tsc.set("showOutline", value)
}

// Size sets the size of the form field
func (tsc TreeSelect) Size(value string) TreeSelect {
	return tsc.set("size", value)
}

// StaticInputClassName sets the CSS class name for the static input
func (tsc TreeSelect) StaticInputClassName(value string) TreeSelect {
	return tsc.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for the static label
func (tsc TreeSelect) StaticLabelClassName(value string) TreeSelect {
	return tsc.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display
func (tsc TreeSelect) StaticOn(value string) TreeSelect {
	return tsc.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (tsc TreeSelect) StaticPlaceholder(value string) TreeSelect {
	return tsc.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema
func (tsc TreeSelect) StaticSchema(value string) TreeSelect {
	return tsc.set("staticSchema", value)
}

// Style sets the component style
func (tsc TreeSelect) Style(value any) TreeSelect {
	return tsc.set("style", value)
}

// SubmitOnChange sets whether to submit the form on change
func (tsc TreeSelect) SubmitOnChange(value bool) TreeSelect {
	return tsc.set("submitOnChange", value)
}

// TestIdBuilder sets the test ID builder
func (tsc TreeSelect) TestIdBuilder(value string) TreeSelect {
	return tsc.set("testIdBuilder", value)
}

// UseMobileUI sets whether to use mobile UI
func (tsc TreeSelect) UseMobileUI(value bool) TreeSelect {
	return tsc.set("useMobileUI", value)
}

// ValidateApi sets the API for remote validation
func (tsc TreeSelect) ValidateApi(value string) TreeSelect {
	return tsc.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on change
func (tsc TreeSelect) ValidateOnChange(value bool) TreeSelect {
	return tsc.set("validateOnChange", value)
}

// ValidationErrors sets the validation error messages
func (tsc TreeSelect) ValidationErrors(value string) TreeSelect {
	return tsc.set("validationErrors", value)
}

// Validations sets the validations
func (tsc TreeSelect) Validations(value string) TreeSelect {
	return tsc.set("validations", value)
}

// ValuesNoWrap sets whether to avoid wrapping in multi-select mode
func (tsc TreeSelect) ValuesNoWrap(value bool) TreeSelect {
	return tsc.set("valuesNoWrap", value)
}

// WithChildren sets whether to include child node values when selecting a parent
func (tsc TreeSelect) WithChildren(value bool) TreeSelect {
	return tsc.set("withChildren", value)
}
