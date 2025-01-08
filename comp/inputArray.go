package comp

import "github.com/zrcoder/amisgo/model"

// inputArray is an alias for combo. Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/form/input-array
type inputArray model.Schema

// InputArray creates a new InputArray instance
func InputArray() inputArray {
	return make(inputArray).set("type", "input-array")
}

func (ac inputArray) set(key string, value any) inputArray {
	ac[key] = value
	return ac
}

// AddButtonClassName sets the CSS class name for the add button
func (ac inputArray) AddButtonClassName(value string) inputArray {
	return ac.set("addButtonClassName", value)
}

// AddButtonText sets the text for the add button
func (ac inputArray) AddButtonText(value string) inputArray {
	return ac.set("addButtonText", value)
}

// Addable sets whether new items can be added
func (ac inputArray) Addable(value bool) inputArray {
	return ac.set("addable", value)
}

// Addattop sets whether the add button is at the top
func (ac inputArray) Addattop(value bool) inputArray {
	return ac.set("addattop", value)
}

// AutoFill sets the auto-fill value
func (ac inputArray) AutoFill(value string) inputArray {
	return ac.set("autoFill", value)
}

// CanAccessSuperData sets whether parent data can be accessed
func (ac inputArray) CanAccessSuperData(value bool) inputArray {
	return ac.set("canAccessSuperData", value)
}

// ClassName sets the CSS class name for the container
func (ac inputArray) ClassName(value string) inputArray {
	return ac.set("className", value)
}

// ClearValueOnHidden sets whether to clear the value when hidden
func (ac inputArray) ClearValueOnHidden(value bool) inputArray {
	return ac.set("clearValueOnHidden", value)
}

// DeleteApi sets the API to call when deleting
func (ac inputArray) DeleteApi(value string) inputArray {
	return ac.set("deleteApi", value)
}

// DeleteConfirmText sets the confirmation text for deletion
func (ac inputArray) DeleteConfirmText(value string) inputArray {
	return ac.set("deleteConfirmText", value)
}

// Delimiter sets the delimiter for flattening
func (ac inputArray) Delimiter(value string) inputArray {
	return ac.set("delimiter", value)
}

// Desc sets the description
func (ac inputArray) Desc(value string) inputArray {
	return ac.set("desc", value)
}

// Description sets the description content
func (ac inputArray) Description(value string) inputArray {
	return ac.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description
func (ac inputArray) DescriptionClassName(value string) inputArray {
	return ac.set("descriptionClassName", value)
}

// Disabled sets whether the input is disabled
func (ac inputArray) Disabled(value bool) inputArray {
	return ac.set("disabled", value)
}

// DisabledOn sets the expression for disabling the input
func (ac inputArray) DisabledOn(value string) inputArray {
	return ac.set("disabledOn", value)
}

// Draggable sets whether the items are draggable
func (ac inputArray) Draggable(value bool) inputArray {
	return ac.set("draggable", value)
}

// DraggableTip sets the tip for dragging
func (ac inputArray) DraggableTip(value string) inputArray {
	return ac.set("draggableTip", value)
}

// EditorSetting sets the editor configuration
func (ac inputArray) EditorSetting(value string) inputArray {
	return ac.set("editorSetting", value)
}

// ExtraName sets the extra field name
func (ac inputArray) ExtraName(value string) inputArray {
	return ac.set("extraName", value)
}

// Flat sets whether to flatten the result
func (ac inputArray) Flat(value bool) inputArray {
	return ac.set("flat", value)
}

// FormClassName sets the CSS class name for the form items
func (ac inputArray) FormClassName(value string) inputArray {
	return ac.set("formClassName", value)
}

// Hidden sets whether the input is hidden
func (ac inputArray) Hidden(value bool) inputArray {
	return ac.set("hidden", value)
}

// HiddenOn sets the expression for hiding the input
func (ac inputArray) HiddenOn(value string) inputArray {
	return ac.set("hiddenOn", value)
}

// Hint sets the input hint
func (ac inputArray) Hint(value string) inputArray {
	return ac.set("hint", value)
}

// Horizontal sets the horizontal layout
func (ac inputArray) Horizontal(value string) inputArray {
	return ac.set("horizontal", value)
}

// ID sets the unique ID for the component
func (ac inputArray) ID(value string) inputArray {
	return ac.set("id", value)
}

// InitAutoFill sets the initial auto-fill value
func (ac inputArray) InitAutoFill(value string) inputArray {
	return ac.set("initAutoFill", value)
}

// Inline sets whether the input is inline
func (ac inputArray) Inline(value bool) inputArray {
	return ac.set("inline", value)
}

// InputClassName sets the CSS class name for the input
func (ac inputArray) InputClassName(value string) inputArray {
	return ac.set("inputClassName", value)
}

// Items sets the configuration for the items
func (ac inputArray) Items(value ...any) inputArray {
	return ac.set("items", value)
}

// JoinValues sets whether to send values with a delimiter to the backend
func (ac inputArray) JoinValues(value bool) inputArray {
	return ac.set("joinValues", value)
}

// Label sets the label
func (ac inputArray) Label(value string) inputArray {
	return ac.set("label", value)
}

// LabelAlign sets the label alignment
func (ac inputArray) LabelAlign(value string) inputArray {
	return ac.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label
func (ac inputArray) LabelClassName(value string) inputArray {
	return ac.set("labelClassName", value)
}

// LabelRemark sets the label remark
func (ac inputArray) LabelRemark(value string) inputArray {
	return ac.set("labelRemark", value)
}

// LabelWidth sets the label width
func (ac inputArray) LabelWidth(value string) inputArray {
	return ac.set("labelWidth", value)
}

// LazyLoad sets whether to enable lazy loading
func (ac inputArray) LazyLoad(value bool) inputArray {
	return ac.set("lazyLoad", value)
}

// MaxLength sets the maximum number of items
func (ac inputArray) MaxLength(value string) inputArray {
	return ac.set("maxLength", value)
}

// Messages sets the messages
func (ac inputArray) Messages(value string) inputArray {
	return ac.set("messages", value)
}

// MinLength sets the minimum number of items
func (ac inputArray) MinLength(value string) inputArray {
	return ac.set("minLength", value)
}

// Mode sets the display mode
func (ac inputArray) Mode(value string) inputArray {
	return ac.set("mode", value)
}

// MultiLine sets whether the input is multi-line
func (ac inputArray) MultiLine(value bool) inputArray {
	return ac.set("multiLine", value)
}

// Multiple sets whether multiple items can be selected
func (ac inputArray) Multiple(value bool) inputArray {
	return ac.set("multiple", value)
}

// Name sets the field name
func (ac inputArray) Name(value string) inputArray {
	return ac.set("name", value)
}

// NoBorder sets whether the input has no border
func (ac inputArray) NoBorder(value bool) inputArray {
	return ac.set("noBorder", value)
}

// Nullable sets whether the input can be null
func (ac inputArray) Nullable(value bool) inputArray {
	return ac.set("nullable", value)
}

// OnEvent sets the event configuration
func (ac inputArray) OnEvent(value any) inputArray {
	return ac.set("onEvent", value)
}

// Placeholder sets the placeholder text
func (ac inputArray) Placeholder(value string) inputArray {
	return ac.set("placeholder", value)
}

// ReadOnly sets whether the input is read-only
func (ac inputArray) ReadOnly(value bool) inputArray {
	return ac.set("readOnly", value)
}

// ReadOnlyOn sets the expression for read-only
func (ac inputArray) ReadOnlyOn(value string) inputArray {
	return ac.set("readOnlyOn", value)
}

// Remark sets the remark
func (ac inputArray) Remark(value string) inputArray {
	return ac.set("remark", value)
}

// Removable sets whether items can be removed
func (ac inputArray) Removable(value bool) inputArray {
	return ac.set("removable", value)
}

// Required sets whether the input is required
func (ac inputArray) Required(value bool) inputArray {
	return ac.set("required", value)
}

// Row sets the number of rows
func (ac inputArray) Row(value string) inputArray {
	return ac.set("row", value)
}

// SaveImmediately sets whether to save immediately
func (ac inputArray) SaveImmediately(value bool) inputArray {
	return ac.set("saveImmediately", value)
}

// Scaffold sets the default value for new items
func (ac inputArray) Scaffold(value string) inputArray {
	return ac.set("scaffold", value)
}

// Size sets the size of the form item
func (ac inputArray) Size(value string) inputArray {
	return ac.set("size", value)
}

// Static sets whether the input is static
func (ac inputArray) Static(value bool) inputArray {
	return ac.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (ac inputArray) StaticClassName(value string) inputArray {
	return ac.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input value
func (ac inputArray) StaticInputClassName(value string) inputArray {
	return ac.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label
func (ac inputArray) StaticLabelClassName(value string) inputArray {
	return ac.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display
func (ac inputArray) StaticOn(value string) inputArray {
	return ac.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (ac inputArray) StaticPlaceholder(value string) inputArray {
	return ac.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (ac inputArray) StaticSchema(value string) inputArray {
	return ac.set("staticSchema", value)
}

// StrictMode sets whether to enable strict mode
func (ac inputArray) StrictMode(value bool) inputArray {
	return ac.set("strictMode", value)
}

// Style sets the component style
func (ac inputArray) Style(value any) inputArray {
	return ac.set("style", value)
}

// SubFormHorizontal sets the horizontal width ratio for sub-forms
func (ac inputArray) SubFormHorizontal(value string) inputArray {
	return ac.set("subFormHorizontal", value)
}

// SubFormMode sets the mode for sub-forms
func (ac inputArray) SubFormMode(value string) inputArray {
	return ac.set("subFormMode", value)
}

// SubmitOnChange sets whether to submit the form on change
func (ac inputArray) SubmitOnChange(value bool) inputArray {
	return ac.set("submitOnChange", value)
}

// SyncFields sets the fields to synchronize
func (ac inputArray) SyncFields(value string) inputArray {
	return ac.set("syncFields", value)
}

// TabsLabelTpl sets the template for tab labels
func (ac inputArray) TabsLabelTpl(value string) inputArray {
	return ac.set("tabsLabelTpl", value)
}

// TabsMode sets whether to use tabs for display
func (ac inputArray) TabsMode(value bool) inputArray {
	return ac.set("tabsMode", value)
}

// TabsStyle sets the style for tabs
func (ac inputArray) TabsStyle(value any) inputArray {
	return ac.set("tabsStyle", value)
}

// TestIdBuilder sets the test ID builder
func (ac inputArray) TestIdBuilder(value string) inputArray {
	return ac.set("testIdBuilder", value)
}

// TypeSwitchable sets whether the type is switchable
func (ac inputArray) TypeSwitchable(value bool) inputArray {
	return ac.set("typeSwitchable", value)
}

// UpdatePristineAfterStoreDataReInit sets whether to update the pristine value after data re-initialization
func (ac inputArray) UpdatePristineAfterStoreDataReInit(value bool) inputArray {
	return ac.set("updatePristineAfterStoreDataReInit", value)
}

// UseMobileUI sets whether to use mobile UI
func (ac inputArray) UseMobileUI(value bool) inputArray {
	return ac.set("useMobileUI", value)
}

// ValidateApi sets the API for remote validation
func (ac inputArray) ValidateApi(value string) inputArray {
	return ac.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on each change
func (ac inputArray) ValidateOnChange(value bool) inputArray {
	return ac.set("validateOnChange", value)
}

// ValidationErrors sets the validation error messages
func (ac inputArray) ValidationErrors(value string) inputArray {
	return ac.set("validationErrors", value)
}

// Validations sets the validation rules
func (ac inputArray) Validations(value string) inputArray {
	return ac.set("validations", value)
}

// Value sets the default value
func (ac inputArray) Value(value string) inputArray {
	return ac.set("value", value)
}

// Visible sets whether the input is visible
func (ac inputArray) Visible(value bool) inputArray {
	return ac.set("visible", value)
}

// VisibleOn sets the expression for visibility
func (ac inputArray) VisibleOn(value string) inputArray {
	return ac.set("visibleOn", value)
}

// Width sets the width in a table
func (ac inputArray) Width(value string) inputArray {
	return ac.set("width", value)
}
