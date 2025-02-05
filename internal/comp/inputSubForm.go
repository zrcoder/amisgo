package comp

import "github.com/zrcoder/amisgo/model"

// InputSubForm represents a sub-form schema
type InputSubForm model.Schema

func NewInputSubForm() InputSubForm {
	return InputSubForm{"type": "input-sub-form"}
}

func (s InputSubForm) set(key string, value any) InputSubForm {
	s[key] = value
	return s
}

// AddButtonClassName sets the CSS class name for the add button
func (s InputSubForm) AddButtonClassName(value string) InputSubForm {
	return s.set("addButtonClassName", value)
}

// AddButtonText sets the text for the add button
func (s InputSubForm) AddButtonText(value string) InputSubForm {
	return s.set("addButtonText", value)
}

// Addable sets whether the form is addable
func (s InputSubForm) Addable(value bool) InputSubForm {
	return s.set("addable", value)
}

// AutoFill sets the auto-fill value
func (s InputSubForm) AutoFill(value string) InputSubForm {
	return s.set("autoFill", value)
}

// BtnLabel sets the default button label
func (s InputSubForm) BtnLabel(value string) InputSubForm {
	return s.set("btnLabel", value)
}

// ClassName sets the CSS class name for the container
func (s InputSubForm) ClassName(value string) InputSubForm {
	return s.set("className", value)
}

// ClearValueOnHidden sets whether to clear the value when hidden
func (s InputSubForm) ClearValueOnHidden(value bool) InputSubForm {
	return s.set("clearValueOnHidden", value)
}

// Desc sets the description
func (s InputSubForm) Desc(value string) InputSubForm {
	return s.set("desc", value)
}

// Description sets the HTML description
func (s InputSubForm) Description(value string) InputSubForm {
	return s.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description
func (s InputSubForm) DescriptionClassName(value string) InputSubForm {
	return s.set("descriptionClassName", value)
}

// Disabled sets whether the form is disabled
func (s InputSubForm) Disabled(value bool) InputSubForm {
	return s.set("disabled", value)
}

// DisabledOn sets the expression to disable the form
func (s InputSubForm) DisabledOn(value string) InputSubForm {
	return s.set("disabledOn", value)
}

// Draggable sets whether the form is draggable
func (s InputSubForm) Draggable(value bool) InputSubForm {
	return s.set("draggable", value)
}

// DraggableTip sets the draggable tip
func (s InputSubForm) DraggableTip(value string) InputSubForm {
	return s.set("draggableTip", value)
}

// EditorSetting sets the editor configuration
func (s InputSubForm) EditorSetting(value string) InputSubForm {
	return s.set("editorSetting", value)
}

// ExtraName sets the extra field name
func (s InputSubForm) ExtraName(value string) InputSubForm {
	return s.set("extraName", value)
}

// Form sets the sub-form details
func (s InputSubForm) Form(value string) InputSubForm {
	return s.set("form", value)
}

// Hidden sets whether the form is hidden
func (s InputSubForm) Hidden(value bool) InputSubForm {
	return s.set("hidden", value)
}

// HiddenOn sets the expression to hide the form
func (s InputSubForm) HiddenOn(value string) InputSubForm {
	return s.set("hiddenOn", value)
}

// Hint sets the input hint
func (s InputSubForm) Hint(value string) InputSubForm {
	return s.set("hint", value)
}

// Horizontal sets the horizontal layout configuration
func (s InputSubForm) Horizontal(value string) InputSubForm {
	return s.set("horizontal", value)
}

// ID sets the unique component ID
func (s InputSubForm) ID(value string) InputSubForm {
	return s.set("id", value)
}

// InitAutoFill sets the initial auto-fill value
func (s InputSubForm) InitAutoFill(value string) InputSubForm {
	return s.set("initAutoFill", value)
}

// Inline sets whether the form control is inline
func (s InputSubForm) Inline(value bool) InputSubForm {
	return s.set("inline", value)
}

// InputClassName sets the CSS class name for the input
func (s InputSubForm) InputClassName(value string) InputSubForm {
	return s.set("inputClassName", value)
}

// ItemClassName sets the CSS class name for the item
func (s InputSubForm) ItemClassName(value string) InputSubForm {
	return s.set("itemClassName", value)
}

// ItemsClassName sets the CSS class name for the items
func (s InputSubForm) ItemsClassName(value string) InputSubForm {
	return s.set("itemsClassName", value)
}

// Label sets the label
func (s InputSubForm) Label(value string) InputSubForm {
	return s.set("label", value)
}

// LabelAlign sets the label alignment
func (s InputSubForm) LabelAlign(value string) InputSubForm {
	return s.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label
func (s InputSubForm) LabelClassName(value string) InputSubForm {
	return s.set("labelClassName", value)
}

// LabelField sets the label field
func (s InputSubForm) LabelField(value string) InputSubForm {
	return s.set("labelField", value)
}

// LabelRemark sets the label remark
func (s InputSubForm) LabelRemark(value string) InputSubForm {
	return s.set("labelRemark", value)
}

// LabelWidth sets the label width
func (s InputSubForm) LabelWidth(value string) InputSubForm {
	return s.set("labelWidth", value)
}

// MaxLength sets the maximum length
func (s InputSubForm) MaxLength(value string) InputSubForm {
	return s.set("maxLength", value)
}

// MinLength sets the minimum length
func (s InputSubForm) MinLength(value string) InputSubForm {
	return s.set("minLength", value)
}

// Mode sets the display mode
func (s InputSubForm) Mode(value string) InputSubForm {
	return s.set("mode", value)
}

// Multiple sets whether multiple selection is allowed
func (s InputSubForm) Multiple(value bool) InputSubForm {
	return s.set("multiple", value)
}

// Name sets the field name
func (s InputSubForm) Name(value string) InputSubForm {
	return s.set("name", value)
}

// OnEvent sets the event configuration
func (s InputSubForm) OnEvent(value any) InputSubForm {
	return s.set("onEvent", value)
}

// Placeholder sets the placeholder text
func (s InputSubForm) Placeholder(value string) InputSubForm {
	return s.set("placeholder", value)
}

// ReadOnly sets whether the form is read-only
func (s InputSubForm) ReadOnly(value bool) InputSubForm {
	return s.set("readOnly", value)
}

// ReadOnlyOn sets the read-only condition
func (s InputSubForm) ReadOnlyOn(value string) InputSubForm {
	return s.set("readOnlyOn", value)
}

// Remark sets the remark
func (s InputSubForm) Remark(value string) InputSubForm {
	return s.set("remark", value)
}

// Removable sets whether the form is removable
func (s InputSubForm) Removable(value bool) InputSubForm {
	return s.set("removable", value)
}

// Required sets whether the field is required
func (s InputSubForm) Required(value bool) InputSubForm {
	return s.set("required", value)
}

// Row sets the row value
func (s InputSubForm) Row(value string) InputSubForm {
	return s.set("row", value)
}

// SaveImmediately sets whether to save immediately
func (s InputSubForm) SaveImmediately(value bool) InputSubForm {
	return s.set("saveImmediately", value)
}

// Scaffold sets the scaffold value
func (s InputSubForm) Scaffold(value string) InputSubForm {
	return s.set("scaffold", value)
}

// ShowErrorMsg sets whether to show error messages
func (s InputSubForm) ShowErrorMsg(value bool) InputSubForm {
	return s.set("showErrorMsg", value)
}

// Size sets the size of the form item
func (s InputSubForm) Size(value string) InputSubForm {
	return s.set("size", value)
}

// Static sets whether the form is static
func (s InputSubForm) Static(value bool) InputSubForm {
	return s.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (s InputSubForm) StaticClassName(value string) InputSubForm {
	return s.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input
func (s InputSubForm) StaticInputClassName(value string) InputSubForm {
	return s.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label
func (s InputSubForm) StaticLabelClassName(value string) InputSubForm {
	return s.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display
func (s InputSubForm) StaticOn(value string) InputSubForm {
	return s.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (s InputSubForm) StaticPlaceholder(value string) InputSubForm {
	return s.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (s InputSubForm) StaticSchema(value string) InputSubForm {
	return s.set("staticSchema", value)
}

// Style sets the component style
func (s InputSubForm) Style(value any) InputSubForm {
	return s.set("style", value)
}

// SubmitOnChange sets whether to submit the form on change
func (s InputSubForm) SubmitOnChange(value bool) InputSubForm {
	return s.set("submitOnChange", value)
}

// TestIdBuilder sets the test ID builder
func (s InputSubForm) TestIdBuilder(value string) InputSubForm {
	return s.set("testIdBuilder", value)
}

// UseMobileUI sets whether to use mobile UI
func (s InputSubForm) UseMobileUI(value bool) InputSubForm {
	return s.set("useMobileUI", value)
}

// ValidateApi sets the API for remote validation
func (s InputSubForm) ValidateApi(value string) InputSubForm {
	return s.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on change
func (s InputSubForm) ValidateOnChange(value bool) InputSubForm {
	return s.set("validateOnChange", value)
}

// ValidationErrors sets the validation error messages
func (s InputSubForm) ValidationErrors(value string) InputSubForm {
	return s.set("validationErrors", value)
}

// Validations sets the validation rules
func (s InputSubForm) Validations(value string) InputSubForm {
	return s.set("validations", value)
}

// Value sets the default value
func (s InputSubForm) Value(value string) InputSubForm {
	return s.set("value", value)
}

// Visible sets whether the form is visible
func (s InputSubForm) Visible(value bool) InputSubForm {
	return s.set("visible", value)
}

// VisibleOn sets the expression for visibility
func (s InputSubForm) VisibleOn(value string) InputSubForm {
	return s.set("visibleOn", value)
}

// Width sets the width in a table
func (s InputSubForm) Width(value string) InputSubForm {
	return s.set("width", value)
}
