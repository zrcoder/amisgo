package comp

import "github.com/zrcoder/amisgo/model"

// InputArray is an alias for combo. Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/form/input-array
type InputArray model.Schema

func NewInputArray() InputArray {
	return InputArray{"type": "input-array"}
}

func (ac InputArray) set(key string, value any) InputArray {
	ac[key] = value
	return ac
}

// AddButtonClassName sets the CSS class name for the add button
func (ac InputArray) AddButtonClassName(value string) InputArray {
	return ac.set("addButtonClassName", value)
}

// AddButtonText sets the text for the add button
func (ac InputArray) AddButtonText(value string) InputArray {
	return ac.set("addButtonText", value)
}

// Addable sets whether new items can be added
func (ac InputArray) Addable(value bool) InputArray {
	return ac.set("addable", value)
}

// Addattop sets whether the add button is at the top
func (ac InputArray) Addattop(value bool) InputArray {
	return ac.set("addattop", value)
}

// AutoFill sets the auto-fill value
func (ac InputArray) AutoFill(value string) InputArray {
	return ac.set("autoFill", value)
}

// CanAccessSuperData sets whether parent data can be accessed
func (ac InputArray) CanAccessSuperData(value bool) InputArray {
	return ac.set("canAccessSuperData", value)
}

// ClassName sets the CSS class name for the container
func (ac InputArray) ClassName(value string) InputArray {
	return ac.set("className", value)
}

// ClearValueOnHidden sets whether to clear the value when hidden
func (ac InputArray) ClearValueOnHidden(value bool) InputArray {
	return ac.set("clearValueOnHidden", value)
}

// DeleteApi sets the API to call when deleting
func (ac InputArray) DeleteApi(value string) InputArray {
	return ac.set("deleteApi", value)
}

// DeleteConfirmText sets the confirmation text for deletion
func (ac InputArray) DeleteConfirmText(value string) InputArray {
	return ac.set("deleteConfirmText", value)
}

// Delimiter sets the delimiter for flattening
func (ac InputArray) Delimiter(value string) InputArray {
	return ac.set("delimiter", value)
}

// Desc sets the description
func (ac InputArray) Desc(value string) InputArray {
	return ac.set("desc", value)
}

// Description sets the description content
func (ac InputArray) Description(value string) InputArray {
	return ac.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description
func (ac InputArray) DescriptionClassName(value string) InputArray {
	return ac.set("descriptionClassName", value)
}

// Disabled sets whether the input is disabled
func (ac InputArray) Disabled(value bool) InputArray {
	return ac.set("disabled", value)
}

// DisabledOn sets the expression for disabling the input
func (ac InputArray) DisabledOn(value string) InputArray {
	return ac.set("disabledOn", value)
}

// Draggable sets whether the items are draggable
func (ac InputArray) Draggable(value bool) InputArray {
	return ac.set("draggable", value)
}

// DraggableTip sets the tip for dragging
func (ac InputArray) DraggableTip(value string) InputArray {
	return ac.set("draggableTip", value)
}

// EditorSetting sets the editor configuration
func (ac InputArray) EditorSetting(value string) InputArray {
	return ac.set("editorSetting", value)
}

// ExtraName sets the extra field name
func (ac InputArray) ExtraName(value string) InputArray {
	return ac.set("extraName", value)
}

// Flat sets whether to flatten the result
func (ac InputArray) Flat(value bool) InputArray {
	return ac.set("flat", value)
}

// FormClassName sets the CSS class name for the form items
func (ac InputArray) FormClassName(value string) InputArray {
	return ac.set("formClassName", value)
}

// Hidden sets whether the input is hidden
func (ac InputArray) Hidden(value bool) InputArray {
	return ac.set("hidden", value)
}

// HiddenOn sets the expression for hiding the input
func (ac InputArray) HiddenOn(value string) InputArray {
	return ac.set("hiddenOn", value)
}

// Hint sets the input hint
func (ac InputArray) Hint(value string) InputArray {
	return ac.set("hint", value)
}

// Horizontal sets the horizontal layout
func (ac InputArray) Horizontal(value string) InputArray {
	return ac.set("horizontal", value)
}

// ID sets the unique ID for the component
func (ac InputArray) ID(value string) InputArray {
	return ac.set("id", value)
}

// InitAutoFill sets the initial auto-fill value
func (ac InputArray) InitAutoFill(value string) InputArray {
	return ac.set("initAutoFill", value)
}

// Inline sets whether the input is inline
func (ac InputArray) Inline(value bool) InputArray {
	return ac.set("inline", value)
}

// InputClassName sets the CSS class name for the input
func (ac InputArray) InputClassName(value string) InputArray {
	return ac.set("inputClassName", value)
}

// Items sets the configuration for the items
func (ac InputArray) Items(value ...any) InputArray {
	return ac.set("items", value)
}

// JoinValues sets whether to send values with a delimiter to the backend
func (ac InputArray) JoinValues(value bool) InputArray {
	return ac.set("joinValues", value)
}

// Label sets the label
func (ac InputArray) Label(value string) InputArray {
	return ac.set("label", value)
}

// LabelAlign sets the label alignment
func (ac InputArray) LabelAlign(value string) InputArray {
	return ac.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label
func (ac InputArray) LabelClassName(value string) InputArray {
	return ac.set("labelClassName", value)
}

// LabelRemark sets the label remark
func (ac InputArray) LabelRemark(value string) InputArray {
	return ac.set("labelRemark", value)
}

// LabelWidth sets the label width
func (ac InputArray) LabelWidth(value string) InputArray {
	return ac.set("labelWidth", value)
}

// LazyLoad sets whether to enable lazy loading
func (ac InputArray) LazyLoad(value bool) InputArray {
	return ac.set("lazyLoad", value)
}

// MaxLength sets the maximum number of items
func (ac InputArray) MaxLength(value string) InputArray {
	return ac.set("maxLength", value)
}

// Messages sets the messages
func (ac InputArray) Messages(value string) InputArray {
	return ac.set("messages", value)
}

// MinLength sets the minimum number of items
func (ac InputArray) MinLength(value string) InputArray {
	return ac.set("minLength", value)
}

// Mode sets the display mode
func (ac InputArray) Mode(value string) InputArray {
	return ac.set("mode", value)
}

// MultiLine sets whether the input is multi-line
func (ac InputArray) MultiLine(value bool) InputArray {
	return ac.set("multiLine", value)
}

// Multiple sets whether multiple items can be selected
func (ac InputArray) Multiple(value bool) InputArray {
	return ac.set("multiple", value)
}

// Name sets the field name
func (ac InputArray) Name(value string) InputArray {
	return ac.set("name", value)
}

// NoBorder sets whether the input has no border
func (ac InputArray) NoBorder(value bool) InputArray {
	return ac.set("noBorder", value)
}

// Nullable sets whether the input can be null
func (ac InputArray) Nullable(value bool) InputArray {
	return ac.set("nullable", value)
}

// OnEvent sets the event configuration
func (ac InputArray) OnEvent(value any) InputArray {
	return ac.set("onEvent", value)
}

// Placeholder sets the placeholder text
func (ac InputArray) Placeholder(value string) InputArray {
	return ac.set("placeholder", value)
}

// ReadOnly sets whether the input is read-only
func (ac InputArray) ReadOnly(value bool) InputArray {
	return ac.set("readOnly", value)
}

// ReadOnlyOn sets the expression for read-only
func (ac InputArray) ReadOnlyOn(value string) InputArray {
	return ac.set("readOnlyOn", value)
}

// Remark sets the remark
func (ac InputArray) Remark(value string) InputArray {
	return ac.set("remark", value)
}

// Removable sets whether items can be removed
func (ac InputArray) Removable(value bool) InputArray {
	return ac.set("removable", value)
}

// Required sets whether the input is required
func (ac InputArray) Required(value bool) InputArray {
	return ac.set("required", value)
}

// Row sets the number of rows
func (ac InputArray) Row(value string) InputArray {
	return ac.set("row", value)
}

// SaveImmediately sets whether to save immediately
func (ac InputArray) SaveImmediately(value bool) InputArray {
	return ac.set("saveImmediately", value)
}

// Scaffold sets the default value for new items
func (ac InputArray) Scaffold(value string) InputArray {
	return ac.set("scaffold", value)
}

// Size sets the size of the form item
func (ac InputArray) Size(value string) InputArray {
	return ac.set("size", value)
}

// Static sets whether the input is static
func (ac InputArray) Static(value bool) InputArray {
	return ac.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (ac InputArray) StaticClassName(value string) InputArray {
	return ac.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input value
func (ac InputArray) StaticInputClassName(value string) InputArray {
	return ac.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label
func (ac InputArray) StaticLabelClassName(value string) InputArray {
	return ac.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display
func (ac InputArray) StaticOn(value string) InputArray {
	return ac.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (ac InputArray) StaticPlaceholder(value string) InputArray {
	return ac.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (ac InputArray) StaticSchema(value string) InputArray {
	return ac.set("staticSchema", value)
}

// StrictMode sets whether to enable strict mode
func (ac InputArray) StrictMode(value bool) InputArray {
	return ac.set("strictMode", value)
}

// Style sets the component style
func (ac InputArray) Style(value any) InputArray {
	return ac.set("style", value)
}

// SubFormHorizontal sets the horizontal width ratio for sub-forms
func (ac InputArray) SubFormHorizontal(value string) InputArray {
	return ac.set("subFormHorizontal", value)
}

// SubFormMode sets the mode for sub-forms
func (ac InputArray) SubFormMode(value string) InputArray {
	return ac.set("subFormMode", value)
}

// SubmitOnChange sets whether to submit the form on change
func (ac InputArray) SubmitOnChange(value bool) InputArray {
	return ac.set("submitOnChange", value)
}

// SyncFields sets the fields to synchronize
func (ac InputArray) SyncFields(value string) InputArray {
	return ac.set("syncFields", value)
}

// TabsLabelTpl sets the template for tab labels
func (ac InputArray) TabsLabelTpl(value string) InputArray {
	return ac.set("tabsLabelTpl", value)
}

// TabsMode sets whether to use tabs for display
func (ac InputArray) TabsMode(value bool) InputArray {
	return ac.set("tabsMode", value)
}

// TabsStyle sets the style for tabs
func (ac InputArray) TabsStyle(value any) InputArray {
	return ac.set("tabsStyle", value)
}

// TestIdBuilder sets the test ID builder
func (ac InputArray) TestIdBuilder(value string) InputArray {
	return ac.set("testIdBuilder", value)
}

// TypeSwitchable sets whether the type is switchable
func (ac InputArray) TypeSwitchable(value bool) InputArray {
	return ac.set("typeSwitchable", value)
}

// UpdatePristineAfterStoreDataReInit sets whether to update the pristine value after data re-initialization
func (ac InputArray) UpdatePristineAfterStoreDataReInit(value bool) InputArray {
	return ac.set("updatePristineAfterStoreDataReInit", value)
}

// UseMobileUI sets whether to use mobile UI
func (ac InputArray) UseMobileUI(value bool) InputArray {
	return ac.set("useMobileUI", value)
}

// ValidateApi sets the API for remote validation
func (ac InputArray) ValidateApi(value string) InputArray {
	return ac.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on each change
func (ac InputArray) ValidateOnChange(value bool) InputArray {
	return ac.set("validateOnChange", value)
}

// ValidationErrors sets the validation error messages
func (ac InputArray) ValidationErrors(value string) InputArray {
	return ac.set("validationErrors", value)
}

// Validations sets the validation rules
func (ac InputArray) Validations(value string) InputArray {
	return ac.set("validations", value)
}

// Value sets the default value
func (ac InputArray) Value(value string) InputArray {
	return ac.set("value", value)
}

// Visible sets whether the input is visible
func (ac InputArray) Visible(value bool) InputArray {
	return ac.set("visible", value)
}

// VisibleOn sets the expression for visibility
func (ac InputArray) VisibleOn(value string) InputArray {
	return ac.set("visibleOn", value)
}

// Width sets the width in a table
func (ac InputArray) Width(value string) InputArray {
	return ac.set("width", value)
}
