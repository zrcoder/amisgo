package comp

import "github.com/zrcoder/amisgo/model"

// ButtonGroupSelect represents a button group selection renderer
type buttonGroupSelect model.Schema

// ButtonGroupSelect creates a new ButtonGroupSelect instance
func ButtonGroupSelect() buttonGroupSelect {
	return buttonGroupSelect{"type": "button-group-select"}
}

func (bgc buttonGroupSelect) set(key string, value any) buttonGroupSelect {
	bgc[key] = value
	return bgc
}

// AddApi sets the API to be called when adding an item
func (bgc buttonGroupSelect) AddApi(value string) buttonGroupSelect {
	return bgc.set("addApi", value)
}

// AddControls sets the form items for adding a new item
func (bgc buttonGroupSelect) AddControls(value string) buttonGroupSelect {
	return bgc.set("addControls", value)
}

// AddDialog configures the settings for the add dialog
func (bgc buttonGroupSelect) AddDialog(value string) buttonGroupSelect {
	return bgc.set("addDialog", value)
}

// AutoFill configures automatic filling
func (bgc buttonGroupSelect) AutoFill(value string) buttonGroupSelect {
	return bgc.set("autoFill", value)
}

// BtnActiveClassName sets the CSS class name for button active state
func (bgc buttonGroupSelect) BtnActiveClassName(value string) buttonGroupSelect {
	return bgc.set("btnActiveClassName", value)
}

// BtnActiveLevel sets the style level for selected buttons
func (bgc buttonGroupSelect) BtnActiveLevel(value string) buttonGroupSelect {
	return bgc.set("btnActiveLevel", value)
}

// BtnClassName sets the CSS class name for buttons
func (bgc buttonGroupSelect) BtnClassName(value string) buttonGroupSelect {
	return bgc.set("btnClassName", value)
}

// BtnLevel sets the style level for buttons
func (bgc buttonGroupSelect) BtnLevel(value string) buttonGroupSelect {
	return bgc.set("btnLevel", value)
}

// Buttons configures the collection of buttons
func (bgc buttonGroupSelect) Buttons(value ...any) buttonGroupSelect {
	return bgc.set("buttons", value)
}

// ClassName sets the CSS class name for the container
func (bgc buttonGroupSelect) ClassName(value string) buttonGroupSelect {
	return bgc.set("className", value)
}

// ClearValueOnHidden determines if the form item value should be deleted when hidden
func (bgc buttonGroupSelect) ClearValueOnHidden(value bool) buttonGroupSelect {
	return bgc.set("clearValueOnHidden", value)
}

// Clearable enables or disables the ability to clear the selection
func (bgc buttonGroupSelect) Clearable(value bool) buttonGroupSelect {
	return bgc.set("clearable", value)
}

// Creatable enables or disables the ability to add new items
func (bgc buttonGroupSelect) Creatable(value bool) buttonGroupSelect {
	return bgc.set("creatable", value)
}

// CreateBtnLabel sets the text for the create button
func (bgc buttonGroupSelect) CreateBtnLabel(value string) buttonGroupSelect {
	return bgc.set("createBtnLabel", value)
}

// DeferApi sets the API for lazy loading
func (bgc buttonGroupSelect) DeferApi(value string) buttonGroupSelect {
	return bgc.set("deferApi", value)
}

// DeferField sets the lazy loading field
func (bgc buttonGroupSelect) DeferField(value string) buttonGroupSelect {
	return bgc.set("deferField", value)
}

// DeleteApi sets the API for deleting options
func (bgc buttonGroupSelect) DeleteApi(value string) buttonGroupSelect {
	return bgc.set("deleteApi", value)
}

// DeleteConfirmText sets the confirmation text for option deletion
func (bgc buttonGroupSelect) DeleteConfirmText(value string) buttonGroupSelect {
	return bgc.set("deleteConfirmText", value)
}

// Delimiter sets the delimiter character
func (bgc buttonGroupSelect) Delimiter(value string) buttonGroupSelect {
	return bgc.set("delimiter", value)
}

// Desc sets the description
func (bgc buttonGroupSelect) Desc(value string) buttonGroupSelect {
	return bgc.set("desc", value)
}

// Description sets the description content
func (bgc buttonGroupSelect) Description(value string) buttonGroupSelect {
	return bgc.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description
func (bgc buttonGroupSelect) DescriptionClassName(value string) buttonGroupSelect {
	return bgc.set("descriptionClassName", value)
}

// Disabled enables or disables the component
func (bgc buttonGroupSelect) Disabled(value bool) buttonGroupSelect {
	return bgc.set("disabled", value)
}

// DisabledOn sets a conditional expression for disabling the component
func (bgc buttonGroupSelect) DisabledOn(value string) buttonGroupSelect {
	return bgc.set("disabledOn", value)
}

// EditApi sets the API for editing
func (bgc buttonGroupSelect) EditApi(value string) buttonGroupSelect {
	return bgc.set("editApi", value)
}

// EditControls sets the form items for editing an option
func (bgc buttonGroupSelect) EditControls(value string) buttonGroupSelect {
	return bgc.set("editControls", value)
}

// EditDialog configures the settings for the edit dialog
func (bgc buttonGroupSelect) EditDialog(value string) buttonGroupSelect {
	return bgc.set("editDialog", value)
}

// Editable enables or disables editing
func (bgc buttonGroupSelect) Editable(value bool) buttonGroupSelect {
	return bgc.set("editable", value)
}

// EditorSetting configures editor-specific settings
func (bgc buttonGroupSelect) EditorSetting(value string) buttonGroupSelect {
	return bgc.set("editorSetting", value)
}

// ExtraName sets an additional field name
func (bgc buttonGroupSelect) ExtraName(value string) buttonGroupSelect {
	return bgc.set("extraName", value)
}

// ExtractValue enables wrapping selected option values as an array
func (bgc buttonGroupSelect) ExtractValue(value bool) buttonGroupSelect {
	return bgc.set("extractValue", value)
}

// Hidden controls the visibility of the component
func (bgc buttonGroupSelect) Hidden(value bool) buttonGroupSelect {
	return bgc.set("hidden", value)
}

// HiddenOn sets a conditional expression for hiding the component
func (bgc buttonGroupSelect) HiddenOn(value string) buttonGroupSelect {
	return bgc.set("hiddenOn", value)
}

// Hint sets the input hint
func (bgc buttonGroupSelect) Hint(value string) buttonGroupSelect {
	return bgc.set("hint", value)
}

// Horizontal sets the specific left-right allocation for horizontal layout
func (bgc buttonGroupSelect) Horizontal(value string) buttonGroupSelect {
	return bgc.set("horizontal", value)
}

// ID sets a unique identifier for the component
func (bgc buttonGroupSelect) ID(value string) buttonGroupSelect {
	return bgc.set("id", value)
}

// InitAutoFill sets initial automatic filling
func (bgc buttonGroupSelect) InitAutoFill(value string) buttonGroupSelect {
	return bgc.set("initAutoFill", value)
}

// InitFetch configures initial source API fetching
func (bgc buttonGroupSelect) InitFetch(value bool) buttonGroupSelect {
	return bgc.set("initFetch", value)
}

// InitFetchOn sets a conditional expression for initial source API fetching
func (bgc buttonGroupSelect) InitFetchOn(value string) buttonGroupSelect {
	return bgc.set("initFetchOn", value)
}

// Inline enables or disables inline mode for form controls
func (bgc buttonGroupSelect) Inline(value bool) buttonGroupSelect {
	return bgc.set("inline", value)
}

// InputClassName sets the CSS class name for the input
func (bgc buttonGroupSelect) InputClassName(value string) buttonGroupSelect {
	return bgc.set("inputClassName", value)
}

// JoinValues configures handling for single and multiple selection modes
func (bgc buttonGroupSelect) JoinValues(value bool) buttonGroupSelect {
	return bgc.set("joinValues", value)
}

// Label sets the description title
func (bgc buttonGroupSelect) Label(value string) buttonGroupSelect {
	return bgc.set("label", value)
}

// LabelAlign sets the alignment for the description title
func (bgc buttonGroupSelect) LabelAlign(value string) buttonGroupSelect {
	return bgc.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label
func (bgc buttonGroupSelect) LabelClassName(value string) buttonGroupSelect {
	return bgc.set("labelClassName", value)
}

// LabelRemark sets the small icon next to the description title
func (bgc buttonGroupSelect) LabelRemark(value string) buttonGroupSelect {
	return bgc.set("labelRemark", value)
}

// LabelWidth sets the width of the label
func (bgc buttonGroupSelect) LabelWidth(value string) buttonGroupSelect {
	return bgc.set("labelWidth", value)
}

// Mode sets the display mode for the current form item
func (bgc buttonGroupSelect) Mode(value string) buttonGroupSelect {
	return bgc.set("mode", value)
}

// Multiple enables or disables multiple selection mode
func (bgc buttonGroupSelect) Multiple(value bool) buttonGroupSelect {
	return bgc.set("multiple", value)
}

// Name sets the form field name
func (bgc buttonGroupSelect) Name(value string) buttonGroupSelect {
	return bgc.set("name", value)
}

// Value sets the component value
func (bgc buttonGroupSelect) Value(value any) buttonGroupSelect {
	return bgc.set("value", value)
}

// OnChange sets the callback function for value changes
func (bgc buttonGroupSelect) OnChange(value string) buttonGroupSelect {
	return bgc.set("onChange", value)
}

// OnEvent sets the event callback function
func (bgc buttonGroupSelect) OnEvent(value any) buttonGroupSelect {
	return bgc.set("onEvent", value)
}

// OptionIcon sets the option icon
func (bgc buttonGroupSelect) OptionIcon(value string) buttonGroupSelect {
	return bgc.set("optionIcon", value)
}

// OptionLabel sets the option label
func (bgc buttonGroupSelect) OptionLabel(value string) buttonGroupSelect {
	return bgc.set("optionLabel", value)
}

// OptionValue sets the option value
func (bgc buttonGroupSelect) OptionValue(value string) buttonGroupSelect {
	return bgc.set("optionValue", value)
}

// Options sets the array of options
func (bgc buttonGroupSelect) Options(value ...any) buttonGroupSelect {
	return bgc.set("options", value)
}

// Outline enables or disables outline button style
func (bgc buttonGroupSelect) Outline(value bool) buttonGroupSelect {
	return bgc.set("outline", value)
}

// Overlay enables or disables Overlay
func (bgc buttonGroupSelect) Overlay(value bool) buttonGroupSelect {
	return bgc.set("overlay", value)
}

// Source sets the source API
func (bgc buttonGroupSelect) Source(value string) buttonGroupSelect {
	return bgc.set("source", value)
}

// Size sets the control size
func (bgc buttonGroupSelect) Size(value string) buttonGroupSelect {
	return bgc.set("size", value)
}

// ValueField sets the value field name
func (bgc buttonGroupSelect) ValueField(value string) buttonGroupSelect {
	return bgc.set("valueField", value)
}

// Validation sets the validation rules
func (bgc buttonGroupSelect) Validation(value string) buttonGroupSelect {
	return bgc.set("validation", value)
}

// Visible controls the visibility of the component
func (bgc buttonGroupSelect) Visible(value bool) buttonGroupSelect {
	return bgc.set("visible", value)
}

// VisibleOn sets a conditional expression for component visibility
func (bgc buttonGroupSelect) VisibleOn(value string) buttonGroupSelect {
	return bgc.set("visibleOn", value)
}
