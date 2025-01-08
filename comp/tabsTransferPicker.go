package comp

import "github.com/zrcoder/amisgo/model"

// tabsTransferPicker represents the schema for the tabs transfer picker component.

type tabsTransferPicker model.Schema

// TabsTransferPicker initializes a new tabsTransferPicker instance.
func TabsTransferPicker() tabsTransferPicker {
	return tabsTransferPicker{}.set("type", "tabs-transfer-picker")
}

// set sets a key-value pair and returns the updated tabsTransferPicker instance.
func (tpc tabsTransferPicker) set(key string, value any) tabsTransferPicker {
	tpc[key] = value
	return tpc
}

// AddApi sets the API for adding items.
func (tpc tabsTransferPicker) AddApi(value string) tabsTransferPicker {
	return tpc.set("addApi", value)
}

// AddControls sets the controls for adding items.
func (tpc tabsTransferPicker) AddControls(value string) tabsTransferPicker {
	return tpc.set("addControls", value)
}

// AddDialog sets the dialog configuration for adding items.
func (tpc tabsTransferPicker) AddDialog(value string) tabsTransferPicker {
	return tpc.set("addDialog", value)
}

// AutoCheckChildren sets whether to auto-check child items.
func (tpc tabsTransferPicker) AutoCheckChildren(value bool) tabsTransferPicker {
	return tpc.set("autoCheckChildren", value)
}

// AutoFill sets the auto-fill configuration.
func (tpc tabsTransferPicker) AutoFill(value string) tabsTransferPicker {
	return tpc.set("autoFill", value)
}

// ClassName sets the CSS class name for the container.
func (tpc tabsTransferPicker) ClassName(value string) tabsTransferPicker {
	return tpc.set("className", value)
}

// ClearValueOnHidden sets whether to clear the value when hidden.
func (tpc tabsTransferPicker) ClearValueOnHidden(value bool) tabsTransferPicker {
	return tpc.set("clearValueOnHidden", value)
}

// Clearable sets whether the field is clearable.
func (tpc tabsTransferPicker) Clearable(value bool) tabsTransferPicker {
	return tpc.set("clearable", value)
}

// Columns sets the columns configuration for table mode.
func (tpc tabsTransferPicker) Columns(value ...any) tabsTransferPicker {
	return tpc.set("columns", value)
}

// Creatable sets whether new items can be created.
func (tpc tabsTransferPicker) Creatable(value bool) tabsTransferPicker {
	return tpc.set("creatable", value)
}

// CreateBtnLabel sets the label for the create button.
func (tpc tabsTransferPicker) CreateBtnLabel(value string) tabsTransferPicker {
	return tpc.set("createBtnLabel", value)
}

// DeferApi sets the API for deferred loading.
func (tpc tabsTransferPicker) DeferApi(value string) tabsTransferPicker {
	return tpc.set("deferApi", value)
}

// DeferField sets the field for deferred loading.
func (tpc tabsTransferPicker) DeferField(value string) tabsTransferPicker {
	return tpc.set("deferField", value)
}

// DeleteApi sets the API for deleting items.
func (tpc tabsTransferPicker) DeleteApi(value string) tabsTransferPicker {
	return tpc.set("deleteApi", value)
}

// DeleteConfirmText sets the confirmation text for deletions.
func (tpc tabsTransferPicker) DeleteConfirmText(value string) tabsTransferPicker {
	return tpc.set("deleteConfirmText", value)
}

// Delimiter sets the delimiter for values.
func (tpc tabsTransferPicker) Delimiter(value string) tabsTransferPicker {
	return tpc.set("delimiter", value)
}

// Desc sets the description.
func (tpc tabsTransferPicker) Desc(value string) tabsTransferPicker {
	return tpc.set("desc", value)
}

// Description sets the HTML description content.
func (tpc tabsTransferPicker) Description(value string) tabsTransferPicker {
	return tpc.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description.
func (tpc tabsTransferPicker) DescriptionClassName(value string) tabsTransferPicker {
	return tpc.set("descriptionClassName", value)
}

// Disabled sets whether the field is disabled.
func (tpc tabsTransferPicker) Disabled(value bool) tabsTransferPicker {
	return tpc.set("disabled", value)
}

// DisabledOn sets the expression to determine if the field is disabled.
func (tpc tabsTransferPicker) DisabledOn(value string) tabsTransferPicker {
	return tpc.set("disabledOn", value)
}

// EditApi sets the API for editing items.
func (tpc tabsTransferPicker) EditApi(value string) tabsTransferPicker {
	return tpc.set("editApi", value)
}

// EditControls sets the controls for editing items.
func (tpc tabsTransferPicker) EditControls(value string) tabsTransferPicker {
	return tpc.set("editControls", value)
}

// EditDialog sets the dialog configuration for editing items.
func (tpc tabsTransferPicker) EditDialog(value string) tabsTransferPicker {
	return tpc.set("editDialog", value)
}

// Editable sets whether the field is editable.
func (tpc tabsTransferPicker) Editable(value bool) tabsTransferPicker {
	return tpc.set("editable", value)
}

// EditorSetting sets the editor configuration.
func (tpc tabsTransferPicker) EditorSetting(value string) tabsTransferPicker {
	return tpc.set("editorSetting", value)
}

// ExtraName sets the extra field name for range components.
func (tpc tabsTransferPicker) ExtraName(value string) tabsTransferPicker {
	return tpc.set("extraName", value)
}

// ExtractValue sets whether to extract the value as an array.
func (tpc tabsTransferPicker) ExtractValue(value bool) tabsTransferPicker {
	return tpc.set("extractValue", value)
}

// Hidden sets whether the field is hidden.
func (tpc tabsTransferPicker) Hidden(value bool) tabsTransferPicker {
	return tpc.set("hidden", value)
}

// HiddenOn sets the expression to determine if the field is hidden.
func (tpc tabsTransferPicker) HiddenOn(value string) tabsTransferPicker {
	return tpc.set("hiddenOn", value)
}

// Hint sets the input hint.
func (tpc tabsTransferPicker) Hint(value string) tabsTransferPicker {
	return tpc.set("hint", value)
}

// Horizontal sets the horizontal layout configuration.
func (tpc tabsTransferPicker) Horizontal(value string) tabsTransferPicker {
	return tpc.set("horizontal", value)
}

// ID sets the unique component ID.
func (tpc tabsTransferPicker) ID(value string) tabsTransferPicker {
	return tpc.set("id", value)
}

// InitAutoFill sets the initial auto-fill configuration.
func (tpc tabsTransferPicker) InitAutoFill(value string) tabsTransferPicker {
	return tpc.set("initAutoFill", value)
}

// InitFetch sets whether to fetch data initially.
func (tpc tabsTransferPicker) InitFetch(value bool) tabsTransferPicker {
	return tpc.set("initFetch", value)
}

// InitFetchOn sets the expression to determine if data should be fetched initially.
func (tpc tabsTransferPicker) InitFetchOn(value string) tabsTransferPicker {
	return tpc.set("initFetchOn", value)
}

// InitiallyOpen sets whether the tabs are initially open.
func (tpc tabsTransferPicker) InitiallyOpen(value bool) tabsTransferPicker {
	return tpc.set("initiallyOpen", value)
}

// Inline sets whether the control is in inline mode.
func (tpc tabsTransferPicker) Inline(value bool) tabsTransferPicker {
	return tpc.set("inline", value)
}

// InputClassName sets the CSS class name for the input.
func (tpc tabsTransferPicker) InputClassName(value string) tabsTransferPicker {
	return tpc.set("inputClassName", value)
}

// ItemHeight sets the height of individual items for virtual rendering.
func (tpc tabsTransferPicker) ItemHeight(value string) tabsTransferPicker {
	return tpc.set("itemHeight", value)
}

// JoinValues sets whether to link the value with the selected items.
func (tpc tabsTransferPicker) JoinValues(value bool) tabsTransferPicker {
	return tpc.set("joinValues", value)
}

// Label sets the label.
func (tpc tabsTransferPicker) Label(value string) tabsTransferPicker {
	return tpc.set("label", value)
}

// LabelClassName sets the CSS class name for the label.
func (tpc tabsTransferPicker) LabelClassName(value string) tabsTransferPicker {
	return tpc.set("labelClassName", value)
}

// LabelRemark sets the remark for the label.
func (tpc tabsTransferPicker) LabelRemark(value string) tabsTransferPicker {
	return tpc.set("labelRemark", value)
}

// MaxLength sets the maximum number of selectable items.
func (tpc tabsTransferPicker) MaxLength(value int) tabsTransferPicker {
	return tpc.set("maxLength", value)
}

// Mode sets the selection mode.
func (tpc tabsTransferPicker) Mode(value string) tabsTransferPicker {
	return tpc.set("mode", value)
}

// Multiple sets whether multiple items can be selected.
func (tpc tabsTransferPicker) Multiple(value bool) tabsTransferPicker {
	return tpc.set("multiple", value)
}

// Name sets the form item name.
func (tpc tabsTransferPicker) Name(value string) tabsTransferPicker {
	return tpc.set("name", value)
}

// OnChange sets the callback for value changes.
func (tpc tabsTransferPicker) OnChange(value string) tabsTransferPicker {
	return tpc.set("onChange", value)
}

// Options sets the options.
func (tpc tabsTransferPicker) Options(value any) tabsTransferPicker {
	return tpc.set("options", value)
}

// OptionsSrc sets the API for fetching options.
func (tpc tabsTransferPicker) OptionsSrc(value string) tabsTransferPicker {
	return tpc.set("optionsSrc", value)
}

// ReadOnly sets whether the field is read-only.
func (tpc tabsTransferPicker) ReadOnly(value bool) tabsTransferPicker {
	return tpc.set("readOnly", value)
}

// Required sets whether the field is required.
func (tpc tabsTransferPicker) Required(value bool) tabsTransferPicker {
	return tpc.set("required", value)
}

// RequiredApi sets the API for required validation.
func (tpc tabsTransferPicker) RequiredApi(value string) tabsTransferPicker {
	return tpc.set("requiredApi", value)
}

// Source sets the API for fetching data.
func (tpc tabsTransferPicker) Source(value string) tabsTransferPicker {
	return tpc.set("source", value)
}

// Tabs sets the tabs configuration.
func (tpc tabsTransferPicker) Tabs(value any) tabsTransferPicker {
	return tpc.set("tabs", value)
}

// Tooltip sets the tooltip text.
func (tpc tabsTransferPicker) Tooltip(value string) tabsTransferPicker {
	return tpc.set("tooltip", value)
}

// Value sets the form value.
func (tpc tabsTransferPicker) Value(value any) tabsTransferPicker {
	return tpc.set("value", value)
}

// Vertical sets whether the list is vertical.
func (tpc tabsTransferPicker) Vertical(value bool) tabsTransferPicker {
	return tpc.set("vertical", value)
}

// Visible sets whether the field is visible.
func (tpc tabsTransferPicker) Visible(value bool) tabsTransferPicker {
	return tpc.set("visible", value)
}

// VisibleOn sets the expression to determine if the field is visible.
func (tpc tabsTransferPicker) VisibleOn(value string) tabsTransferPicker {
	return tpc.set("visibleOn", value)
}

// Width sets the width.
func (tpc tabsTransferPicker) Width(value string) tabsTransferPicker {
	return tpc.set("width", value)
}

// WidthConfig sets the width configuration.
func (tpc tabsTransferPicker) WidthConfig(value string) tabsTransferPicker {
	return tpc.set("widthConfig", value)
}

// Xname sets the key name.
func (tpc tabsTransferPicker) Xname(value string) tabsTransferPicker {
	return tpc.set("xname", value)
}

// Xtype sets the type.
func (tpc tabsTransferPicker) Xtype(value string) tabsTransferPicker {
	return tpc.set("xtype", value)
}

// Component sets the custom component.
func (tpc tabsTransferPicker) Component(value string) tabsTransferPicker {
	return tpc.set("component", value)
}
