package comp

import "github.com/zrcoder/amisgo/model"

// inputSubForm represents a sub-form schema
type inputSubForm model.Schema

func InputSubForm() inputSubForm {
	return inputSubForm{}.set("type", "input-sub-form")
}

func (s inputSubForm) set(key string, value any) inputSubForm {
	s[key] = value
	return s
}

// AddButtonClassName sets the CSS class name for the add button
func (s inputSubForm) AddButtonClassName(value string) inputSubForm {
	return s.set("addButtonClassName", value)
}

// AddButtonText sets the text for the add button
func (s inputSubForm) AddButtonText(value string) inputSubForm {
	return s.set("addButtonText", value)
}

// Addable sets whether the form is addable
func (s inputSubForm) Addable(value bool) inputSubForm {
	return s.set("addable", value)
}

// AutoFill sets the auto-fill value
func (s inputSubForm) AutoFill(value string) inputSubForm {
	return s.set("autoFill", value)
}

// BtnLabel sets the default button label
func (s inputSubForm) BtnLabel(value string) inputSubForm {
	return s.set("btnLabel", value)
}

// ClassName sets the CSS class name for the container
func (s inputSubForm) ClassName(value string) inputSubForm {
	return s.set("className", value)
}

// ClearValueOnHidden sets whether to clear the value when hidden
func (s inputSubForm) ClearValueOnHidden(value bool) inputSubForm {
	return s.set("clearValueOnHidden", value)
}

// Desc sets the description
func (s inputSubForm) Desc(value string) inputSubForm {
	return s.set("desc", value)
}

// Description sets the HTML description
func (s inputSubForm) Description(value string) inputSubForm {
	return s.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description
func (s inputSubForm) DescriptionClassName(value string) inputSubForm {
	return s.set("descriptionClassName", value)
}

// Disabled sets whether the form is disabled
func (s inputSubForm) Disabled(value bool) inputSubForm {
	return s.set("disabled", value)
}

// DisabledOn sets the expression to disable the form
func (s inputSubForm) DisabledOn(value string) inputSubForm {
	return s.set("disabledOn", value)
}

// Draggable sets whether the form is draggable
func (s inputSubForm) Draggable(value bool) inputSubForm {
	return s.set("draggable", value)
}

// DraggableTip sets the draggable tip
func (s inputSubForm) DraggableTip(value string) inputSubForm {
	return s.set("draggableTip", value)
}

// EditorSetting sets the editor configuration
func (s inputSubForm) EditorSetting(value string) inputSubForm {
	return s.set("editorSetting", value)
}

// ExtraName sets the extra field name
func (s inputSubForm) ExtraName(value string) inputSubForm {
	return s.set("extraName", value)
}

// Form sets the sub-form details
func (s inputSubForm) Form(value string) inputSubForm {
	return s.set("form", value)
}

// Hidden sets whether the form is hidden
func (s inputSubForm) Hidden(value bool) inputSubForm {
	return s.set("hidden", value)
}

// HiddenOn sets the expression to hide the form
func (s inputSubForm) HiddenOn(value string) inputSubForm {
	return s.set("hiddenOn", value)
}

// Hint sets the input hint
func (s inputSubForm) Hint(value string) inputSubForm {
	return s.set("hint", value)
}

// Horizontal sets the horizontal layout configuration
func (s inputSubForm) Horizontal(value string) inputSubForm {
	return s.set("horizontal", value)
}

// ID sets the unique component ID
func (s inputSubForm) ID(value string) inputSubForm {
	return s.set("id", value)
}

// InitAutoFill sets the initial auto-fill value
func (s inputSubForm) InitAutoFill(value string) inputSubForm {
	return s.set("initAutoFill", value)
}

// Inline sets whether the form control is inline
func (s inputSubForm) Inline(value bool) inputSubForm {
	return s.set("inline", value)
}

// InputClassName sets the CSS class name for the input
func (s inputSubForm) InputClassName(value string) inputSubForm {
	return s.set("inputClassName", value)
}

// ItemClassName sets the CSS class name for the item
func (s inputSubForm) ItemClassName(value string) inputSubForm {
	return s.set("itemClassName", value)
}

// ItemsClassName sets the CSS class name for the items
func (s inputSubForm) ItemsClassName(value string) inputSubForm {
	return s.set("itemsClassName", value)
}

// Label sets the label
func (s inputSubForm) Label(value string) inputSubForm {
	return s.set("label", value)
}

// LabelAlign sets the label alignment
func (s inputSubForm) LabelAlign(value string) inputSubForm {
	return s.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label
func (s inputSubForm) LabelClassName(value string) inputSubForm {
	return s.set("labelClassName", value)
}

// LabelField sets the label field
func (s inputSubForm) LabelField(value string) inputSubForm {
	return s.set("labelField", value)
}

// LabelRemark sets the label remark
func (s inputSubForm) LabelRemark(value string) inputSubForm {
	return s.set("labelRemark", value)
}

// LabelWidth sets the label width
func (s inputSubForm) LabelWidth(value string) inputSubForm {
	return s.set("labelWidth", value)
}

// MaxLength sets the maximum length
func (s inputSubForm) MaxLength(value string) inputSubForm {
	return s.set("maxLength", value)
}

// MinLength sets the minimum length
func (s inputSubForm) MinLength(value string) inputSubForm {
	return s.set("minLength", value)
}

// Mode sets the display mode
func (s inputSubForm) Mode(value string) inputSubForm {
	return s.set("mode", value)
}

// Multiple sets whether multiple selection is allowed
func (s inputSubForm) Multiple(value bool) inputSubForm {
	return s.set("multiple", value)
}

// Name sets the field name
func (s inputSubForm) Name(value string) inputSubForm {
	return s.set("name", value)
}

// OnEvent sets the event configuration
func (s inputSubForm) OnEvent(value any) inputSubForm {
	return s.set("onEvent", value)
}

// Placeholder sets the placeholder text
func (s inputSubForm) Placeholder(value string) inputSubForm {
	return s.set("placeholder", value)
}

// ReadOnly sets whether the form is read-only
func (s inputSubForm) ReadOnly(value bool) inputSubForm {
	return s.set("readOnly", value)
}

// ReadOnlyOn sets the read-only condition
func (s inputSubForm) ReadOnlyOn(value string) inputSubForm {
	return s.set("readOnlyOn", value)
}

// Remark sets the remark
func (s inputSubForm) Remark(value string) inputSubForm {
	return s.set("remark", value)
}

// Removable sets whether the form is removable
func (s inputSubForm) Removable(value bool) inputSubForm {
	return s.set("removable", value)
}

// Required sets whether the field is required
func (s inputSubForm) Required(value bool) inputSubForm {
	return s.set("required", value)
}

// Row sets the row value
func (s inputSubForm) Row(value string) inputSubForm {
	return s.set("row", value)
}

// SaveImmediately sets whether to save immediately
func (s inputSubForm) SaveImmediately(value bool) inputSubForm {
	return s.set("saveImmediately", value)
}

// Scaffold sets the scaffold value
func (s inputSubForm) Scaffold(value string) inputSubForm {
	return s.set("scaffold", value)
}

// ShowErrorMsg sets whether to show error messages
func (s inputSubForm) ShowErrorMsg(value bool) inputSubForm {
	return s.set("showErrorMsg", value)
}

// Size sets the size of the form item
func (s inputSubForm) Size(value string) inputSubForm {
	return s.set("size", value)
}

// Static sets whether the form is static
func (s inputSubForm) Static(value bool) inputSubForm {
	return s.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (s inputSubForm) StaticClassName(value string) inputSubForm {
	return s.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input
func (s inputSubForm) StaticInputClassName(value string) inputSubForm {
	return s.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label
func (s inputSubForm) StaticLabelClassName(value string) inputSubForm {
	return s.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display
func (s inputSubForm) StaticOn(value string) inputSubForm {
	return s.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (s inputSubForm) StaticPlaceholder(value string) inputSubForm {
	return s.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (s inputSubForm) StaticSchema(value string) inputSubForm {
	return s.set("staticSchema", value)
}

// Style sets the component style
func (s inputSubForm) Style(value any) inputSubForm {
	return s.set("style", value)
}

// SubmitOnChange sets whether to submit the form on change
func (s inputSubForm) SubmitOnChange(value bool) inputSubForm {
	return s.set("submitOnChange", value)
}

// TestIdBuilder sets the test ID builder
func (s inputSubForm) TestIdBuilder(value string) inputSubForm {
	return s.set("testIdBuilder", value)
}

// UseMobileUI sets whether to use mobile UI
func (s inputSubForm) UseMobileUI(value bool) inputSubForm {
	return s.set("useMobileUI", value)
}

// ValidateApi sets the API for remote validation
func (s inputSubForm) ValidateApi(value string) inputSubForm {
	return s.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on change
func (s inputSubForm) ValidateOnChange(value bool) inputSubForm {
	return s.set("validateOnChange", value)
}

// ValidationErrors sets the validation error messages
func (s inputSubForm) ValidationErrors(value string) inputSubForm {
	return s.set("validationErrors", value)
}

// Validations sets the validation rules
func (s inputSubForm) Validations(value string) inputSubForm {
	return s.set("validations", value)
}

// Value sets the default value
func (s inputSubForm) Value(value string) inputSubForm {
	return s.set("value", value)
}

// Visible sets whether the form is visible
func (s inputSubForm) Visible(value bool) inputSubForm {
	return s.set("visible", value)
}

// VisibleOn sets the expression for visibility
func (s inputSubForm) VisibleOn(value string) inputSubForm {
	return s.set("visibleOn", value)
}

// Width sets the width in a table
func (s inputSubForm) Width(value string) inputSubForm {
	return s.set("width", value)
}
