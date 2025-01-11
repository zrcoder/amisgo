package comp

import "github.com/zrcoder/amisgo/model"

// inputRepeat documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/form/repeat

type inputRepeat model.Schema

// InputRepeat creates a new RepeatControl instance
func InputRepeat() inputRepeat {
	return inputRepeat{"type": "input-repeat"}
}

func (rc inputRepeat) set(key string, value any) inputRepeat {
	rc[key] = value
	return rc
}

// AutoFill sets autoFill when an option is selected
func (rc inputRepeat) AutoFill(value string) inputRepeat {
	return rc.set("autoFill", value)
}

// ClassName sets the container CSS class name
func (rc inputRepeat) ClassName(value string) inputRepeat {
	return rc.set("className", value)
}

// ClearValueOnHidden removes the form item value when hidden
func (rc inputRepeat) ClearValueOnHidden(value bool) inputRepeat {
	return rc.set("clearValueOnHidden", value)
}

// Desc sets the description
func (rc inputRepeat) Desc(value string) inputRepeat {
	return rc.set("desc", value)
}

// Description sets the description content, supports HTML
func (rc inputRepeat) Description(value string) inputRepeat {
	return rc.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description
func (rc inputRepeat) DescriptionClassName(value string) inputRepeat {
	return rc.set("descriptionClassName", value)
}

// Disabled sets whether the input is disabled
func (rc inputRepeat) Disabled(value bool) inputRepeat {
	return rc.set("disabled", value)
}

// DisabledOn sets the condition to disable the input
func (rc inputRepeat) DisabledOn(value string) inputRepeat {
	return rc.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (rc inputRepeat) EditorSetting(value string) inputRepeat {
	return rc.set("editorSetting", value)
}

// ExtraName sets an extra field name for range components
func (rc inputRepeat) ExtraName(value string) inputRepeat {
	return rc.set("extraName", value)
}

// Hidden sets whether the input is hidden
func (rc inputRepeat) Hidden(value bool) inputRepeat {
	return rc.set("hidden", value)
}

// HiddenOn sets the condition to hide the input
func (rc inputRepeat) HiddenOn(value string) inputRepeat {
	return rc.set("hiddenOn", value)
}

// Hint sets the input hint displayed on focus
func (rc inputRepeat) Hint(value string) inputRepeat {
	return rc.set("hint", value)
}

// Horizontal sets the horizontal layout configuration
func (rc inputRepeat) Horizontal(value string) inputRepeat {
	return rc.set("horizontal", value)
}

// ID sets the unique component ID
func (rc inputRepeat) ID(value string) inputRepeat {
	return rc.set("id", value)
}

// InitAutoFill sets the initial autoFill value
func (rc inputRepeat) InitAutoFill(value string) inputRepeat {
	return rc.set("initAutoFill", value)
}

// Inline sets whether the form control is inline
func (rc inputRepeat) Inline(value bool) inputRepeat {
	return rc.set("inline", value)
}

// InputClassName sets the CSS class name for the input
func (rc inputRepeat) InputClassName(value string) inputRepeat {
	return rc.set("inputClassName", value)
}

// Label sets the label text
func (rc inputRepeat) Label(value string) inputRepeat {
	return rc.set("label", value)
}

// LabelAlign sets the label alignment
func (rc inputRepeat) LabelAlign(value string) inputRepeat {
	return rc.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label
func (rc inputRepeat) LabelClassName(value string) inputRepeat {
	return rc.set("labelClassName", value)
}

// LabelRemark sets the label remark
func (rc inputRepeat) LabelRemark(value string) inputRepeat {
	return rc.set("labelRemark", value)
}

// LabelWidth sets the custom width for the label
func (rc inputRepeat) LabelWidth(value string) inputRepeat {
	return rc.set("labelWidth", value)
}

// Mode sets the display mode of the form item
func (rc inputRepeat) Mode(value string) inputRepeat {
	return rc.set("mode", value)
}

// Name sets the field name for form submission
func (rc inputRepeat) Name(value string) inputRepeat {
	return rc.set("name", value)
}

// OnEvent sets the event action configuration
func (rc inputRepeat) OnEvent(value any) inputRepeat {
	return rc.set("onEvent", value)
}

// Options sets the options
func (rc inputRepeat) Options(value ...any) inputRepeat {
	return rc.set("options", value)
}

// Placeholder sets the placeholder text
func (rc inputRepeat) Placeholder(value string) inputRepeat {
	return rc.set("placeholder", value)
}

// ReadOnly sets whether the input is read-only
func (rc inputRepeat) ReadOnly(value bool) inputRepeat {
	return rc.set("readOnly", value)
}

// ReadOnlyOn sets the condition for read-only
func (rc inputRepeat) ReadOnlyOn(value string) inputRepeat {
	return rc.set("readOnlyOn", value)
}

// Remark sets the remark
func (rc inputRepeat) Remark(value string) inputRepeat {
	return rc.set("remark", value)
}

// Required sets whether the input is required
func (rc inputRepeat) Required(value bool) inputRepeat {
	return rc.set("required", value)
}

// Row sets the row value
func (rc inputRepeat) Row(value string) inputRepeat {
	return rc.set("row", value)
}

// SaveImmediately sets whether to save immediately in TableColumn
func (rc inputRepeat) SaveImmediately(value bool) inputRepeat {
	return rc.set("saveImmediately", value)
}

// Size sets the size of the form item
func (rc inputRepeat) Size(value string) inputRepeat {
	return rc.set("size", value)
}

// Static sets whether the input is displayed statically
func (rc inputRepeat) Static(value bool) inputRepeat {
	return rc.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (rc inputRepeat) StaticClassName(value string) inputRepeat {
	return rc.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (rc inputRepeat) StaticInputClassName(value string) inputRepeat {
	return rc.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (rc inputRepeat) StaticLabelClassName(value string) inputRepeat {
	return rc.set("staticLabelClassName", value)
}

// StaticOn sets the condition for static display
func (rc inputRepeat) StaticOn(value string) inputRepeat {
	return rc.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (rc inputRepeat) StaticPlaceholder(value string) inputRepeat {
	return rc.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema
func (rc inputRepeat) StaticSchema(value string) inputRepeat {
	return rc.set("staticSchema", value)
}

// Style sets the component style
func (rc inputRepeat) Style(value any) inputRepeat {
	return rc.set("style", value)
}

// SubmitOnChange sets whether to submit the form on change
func (rc inputRepeat) SubmitOnChange(value bool) inputRepeat {
	return rc.set("submitOnChange", value)
}

// TestIdBuilder sets the test ID builder
func (rc inputRepeat) TestIdBuilder(value string) inputRepeat {
	return rc.set("testIdBuilder", value)
}

// UseMobileUI sets whether to use mobile UI
func (rc inputRepeat) UseMobileUI(value bool) inputRepeat {
	return rc.set("useMobileUI", value)
}

// ValidateApi sets the remote validation API
func (rc inputRepeat) ValidateApi(value string) inputRepeat {
	return rc.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on change
func (rc inputRepeat) ValidateOnChange(value bool) inputRepeat {
	return rc.set("validateOnChange", value)
}

// ValidationErrors sets the validation error messages
func (rc inputRepeat) ValidationErrors(value string) inputRepeat {
	return rc.set("validationErrors", value)
}

// Validations sets the validations
func (rc inputRepeat) Validations(value string) inputRepeat {
	return rc.set("validations", value)
}

// Value sets the default value
func (rc inputRepeat) Value(value string) inputRepeat {
	return rc.set("value", value)
}

// Visible sets whether the input is visible
func (rc inputRepeat) Visible(value bool) inputRepeat {
	return rc.set("visible", value)
}

// VisibleOn sets the condition for visibility
func (rc inputRepeat) VisibleOn(value string) inputRepeat {
	return rc.set("visibleOn", value)
}

// Width sets the width in Table
func (rc inputRepeat) Width(value string) inputRepeat {
	return rc.set("width", value)
}
