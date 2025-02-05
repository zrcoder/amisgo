package comp

import "github.com/zrcoder/amisgo/model"

// TabsTransferPicker represents the schema for the tabs transfer picker component.
type TabsTransferPicker model.Schema

func NewTabsTransferPicker() TabsTransferPicker {
	return TabsTransferPicker{"type": "tabs-transfer-picker"}
}

// set sets a key-value pair and returns the updated tabsTransferPicker instance.
func (tpc TabsTransferPicker) set(key string, value any) TabsTransferPicker {
	tpc[key] = value
	return tpc
}

// AddApi sets the API for adding items.
func (tpc TabsTransferPicker) AddApi(value string) TabsTransferPicker {
	return tpc.set("addApi", value)
}

// AddControls sets the controls for adding items.
func (tpc TabsTransferPicker) AddControls(value string) TabsTransferPicker {
	return tpc.set("addControls", value)
}

// AddDialog sets the dialog configuration for adding items.
func (tpc TabsTransferPicker) AddDialog(value string) TabsTransferPicker {
	return tpc.set("addDialog", value)
}

// AutoCheckChildren sets whether to auto-check child items.
func (tpc TabsTransferPicker) AutoCheckChildren(value bool) TabsTransferPicker {
	return tpc.set("autoCheckChildren", value)
}

// AutoFill sets the auto-fill configuration.
func (tpc TabsTransferPicker) AutoFill(value string) TabsTransferPicker {
	return tpc.set("autoFill", value)
}

// ClassName sets the CSS class name for the container.
func (tpc TabsTransferPicker) ClassName(value string) TabsTransferPicker {
	return tpc.set("className", value)
}

// ClearValueOnHidden sets whether to clear the value when hidden.
func (tpc TabsTransferPicker) ClearValueOnHidden(value bool) TabsTransferPicker {
	return tpc.set("clearValueOnHidden", value)
}

// Clearable sets whether the field is clearable.
func (tpc TabsTransferPicker) Clearable(value bool) TabsTransferPicker {
	return tpc.set("clearable", value)
}

// Columns sets the columns configuration for table mode.
func (tpc TabsTransferPicker) Columns(value ...any) TabsTransferPicker {
	return tpc.set("columns", value)
}

// Creatable sets whether new items can be created.
func (tpc TabsTransferPicker) Creatable(value bool) TabsTransferPicker {
	return tpc.set("creatable", value)
}

// CreateBtnLabel sets the label for the create button.
func (tpc TabsTransferPicker) CreateBtnLabel(value string) TabsTransferPicker {
	return tpc.set("createBtnLabel", value)
}

// DeferApi sets the API for deferred loading.
func (tpc TabsTransferPicker) DeferApi(value string) TabsTransferPicker {
	return tpc.set("deferApi", value)
}

// DeferField sets the field for deferred loading.
func (tpc TabsTransferPicker) DeferField(value string) TabsTransferPicker {
	return tpc.set("deferField", value)
}

// DeleteApi sets the API for deleting items.
func (tpc TabsTransferPicker) DeleteApi(value string) TabsTransferPicker {
	return tpc.set("deleteApi", value)
}

// DeleteConfirmText sets the confirmation text for deletions.
func (tpc TabsTransferPicker) DeleteConfirmText(value string) TabsTransferPicker {
	return tpc.set("deleteConfirmText", value)
}

// Delimiter sets the delimiter for values.
func (tpc TabsTransferPicker) Delimiter(value string) TabsTransferPicker {
	return tpc.set("delimiter", value)
}

// Desc sets the description.
func (tpc TabsTransferPicker) Desc(value string) TabsTransferPicker {
	return tpc.set("desc", value)
}

// Description sets the HTML description content.
func (tpc TabsTransferPicker) Description(value string) TabsTransferPicker {
	return tpc.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description.
func (tpc TabsTransferPicker) DescriptionClassName(value string) TabsTransferPicker {
	return tpc.set("descriptionClassName", value)
}

// Disabled sets whether the field is disabled.
func (tpc TabsTransferPicker) Disabled(value bool) TabsTransferPicker {
	return tpc.set("disabled", value)
}

// DisabledOn sets the expression to determine if the field is disabled.
func (tpc TabsTransferPicker) DisabledOn(value string) TabsTransferPicker {
	return tpc.set("disabledOn", value)
}

// EditApi sets the API for editing items.
func (tpc TabsTransferPicker) EditApi(value string) TabsTransferPicker {
	return tpc.set("editApi", value)
}

// EditControls sets the controls for editing items.
func (tpc TabsTransferPicker) EditControls(value string) TabsTransferPicker {
	return tpc.set("editControls", value)
}

// EditDialog sets the dialog configuration for editing items.
func (tpc TabsTransferPicker) EditDialog(value string) TabsTransferPicker {
	return tpc.set("editDialog", value)
}

// Editable sets whether the field is editable.
func (tpc TabsTransferPicker) Editable(value bool) TabsTransferPicker {
	return tpc.set("editable", value)
}

// EditorSetting sets the editor configuration.
func (tpc TabsTransferPicker) EditorSetting(value string) TabsTransferPicker {
	return tpc.set("editorSetting", value)
}

// ExtraName sets the extra field name for range components.
func (tpc TabsTransferPicker) ExtraName(value string) TabsTransferPicker {
	return tpc.set("extraName", value)
}

// ExtractValue sets whether to extract the value as an array.
func (tpc TabsTransferPicker) ExtractValue(value bool) TabsTransferPicker {
	return tpc.set("extractValue", value)
}

// Hidden sets whether the field is hidden.
func (tpc TabsTransferPicker) Hidden(value bool) TabsTransferPicker {
	return tpc.set("hidden", value)
}

// HiddenOn sets the expression to determine if the field is hidden.
func (tpc TabsTransferPicker) HiddenOn(value string) TabsTransferPicker {
	return tpc.set("hiddenOn", value)
}

// Hint sets the input hint.
func (tpc TabsTransferPicker) Hint(value string) TabsTransferPicker {
	return tpc.set("hint", value)
}

// Horizontal sets the horizontal layout configuration.
func (tpc TabsTransferPicker) Horizontal(value string) TabsTransferPicker {
	return tpc.set("horizontal", value)
}

// ID sets the unique component ID.
func (tpc TabsTransferPicker) ID(value string) TabsTransferPicker {
	return tpc.set("id", value)
}

// InitAutoFill sets the initial auto-fill configuration.
func (tpc TabsTransferPicker) InitAutoFill(value string) TabsTransferPicker {
	return tpc.set("initAutoFill", value)
}

// InitFetch sets whether to fetch data initially.
func (tpc TabsTransferPicker) InitFetch(value bool) TabsTransferPicker {
	return tpc.set("initFetch", value)
}

// InitFetchOn sets the expression to determine if data should be fetched initially.
func (tpc TabsTransferPicker) InitFetchOn(value string) TabsTransferPicker {
	return tpc.set("initFetchOn", value)
}

// InitiallyOpen sets whether the tabs are initially open.
func (tpc TabsTransferPicker) InitiallyOpen(value bool) TabsTransferPicker {
	return tpc.set("initiallyOpen", value)
}

// Inline sets whether the control is in inline mode.
func (tpc TabsTransferPicker) Inline(value bool) TabsTransferPicker {
	return tpc.set("inline", value)
}

// InputClassName sets the CSS class name for the input.
func (tpc TabsTransferPicker) InputClassName(value string) TabsTransferPicker {
	return tpc.set("inputClassName", value)
}

// ItemHeight sets the height of individual items for virtual rendering.
func (tpc TabsTransferPicker) ItemHeight(value string) TabsTransferPicker {
	return tpc.set("itemHeight", value)
}

// JoinValues sets whether to link the value with the selected items.
func (tpc TabsTransferPicker) JoinValues(value bool) TabsTransferPicker {
	return tpc.set("joinValues", value)
}

// Label sets the label.
func (tpc TabsTransferPicker) Label(value string) TabsTransferPicker {
	return tpc.set("label", value)
}

// LabelClassName sets the CSS class name for the label.
func (tpc TabsTransferPicker) LabelClassName(value string) TabsTransferPicker {
	return tpc.set("labelClassName", value)
}

// LabelRemark sets the remark for the label.
func (tpc TabsTransferPicker) LabelRemark(value string) TabsTransferPicker {
	return tpc.set("labelRemark", value)
}

// MaxLength sets the maximum number of selectable items.
func (tpc TabsTransferPicker) MaxLength(value int) TabsTransferPicker {
	return tpc.set("maxLength", value)
}

// Mode sets the selection mode.
func (tpc TabsTransferPicker) Mode(value string) TabsTransferPicker {
	return tpc.set("mode", value)
}

// Multiple sets whether multiple items can be selected.
func (tpc TabsTransferPicker) Multiple(value bool) TabsTransferPicker {
	return tpc.set("multiple", value)
}

// Name sets the form item name.
func (tpc TabsTransferPicker) Name(value string) TabsTransferPicker {
	return tpc.set("name", value)
}

// OnChange sets the callback for value changes.
func (tpc TabsTransferPicker) OnChange(value string) TabsTransferPicker {
	return tpc.set("onChange", value)
}

// Options sets the options.
func (tpc TabsTransferPicker) Options(value any) TabsTransferPicker {
	return tpc.set("options", value)
}

// OptionsSrc sets the API for fetching options.
func (tpc TabsTransferPicker) OptionsSrc(value string) TabsTransferPicker {
	return tpc.set("optionsSrc", value)
}

// ReadOnly sets whether the field is read-only.
func (tpc TabsTransferPicker) ReadOnly(value bool) TabsTransferPicker {
	return tpc.set("readOnly", value)
}

// Required sets whether the field is required.
func (tpc TabsTransferPicker) Required(value bool) TabsTransferPicker {
	return tpc.set("required", value)
}

// RequiredApi sets the API for required validation.
func (tpc TabsTransferPicker) RequiredApi(value string) TabsTransferPicker {
	return tpc.set("requiredApi", value)
}

// Source sets the API for fetching data.
func (tpc TabsTransferPicker) Source(value string) TabsTransferPicker {
	return tpc.set("source", value)
}

// Tabs sets the tabs configuration.
func (tpc TabsTransferPicker) Tabs(value any) TabsTransferPicker {
	return tpc.set("tabs", value)
}

// Tooltip sets the tooltip text.
func (tpc TabsTransferPicker) Tooltip(value string) TabsTransferPicker {
	return tpc.set("tooltip", value)
}

// Value sets the form value.
func (tpc TabsTransferPicker) Value(value any) TabsTransferPicker {
	return tpc.set("value", value)
}

// Vertical sets whether the list is vertical.
func (tpc TabsTransferPicker) Vertical(value bool) TabsTransferPicker {
	return tpc.set("vertical", value)
}

// Visible sets whether the field is visible.
func (tpc TabsTransferPicker) Visible(value bool) TabsTransferPicker {
	return tpc.set("visible", value)
}

// VisibleOn sets the expression to determine if the field is visible.
func (tpc TabsTransferPicker) VisibleOn(value string) TabsTransferPicker {
	return tpc.set("visibleOn", value)
}

// Width sets the width.
func (tpc TabsTransferPicker) Width(value string) TabsTransferPicker {
	return tpc.set("width", value)
}

// WidthConfig sets the width configuration.
func (tpc TabsTransferPicker) WidthConfig(value string) TabsTransferPicker {
	return tpc.set("widthConfig", value)
}

// Xname sets the key name.
func (tpc TabsTransferPicker) Xname(value string) TabsTransferPicker {
	return tpc.set("xname", value)
}

// Xtype sets the type.
func (tpc TabsTransferPicker) Xtype(value string) TabsTransferPicker {
	return tpc.set("xtype", value)
}

// Component sets the custom component.
func (tpc TabsTransferPicker) Component(value string) TabsTransferPicker {
	return tpc.set("component", value)
}
