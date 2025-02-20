package comp

import "github.com/zrcoder/amisgo/schema"

// Hidden represents a Hidden field component.
type Hidden schema.Schema

func NewHidden() Hidden {
	return Hidden{"type": "hidden"}
}

func (h Hidden) set(key string, value any) Hidden {
	h[key] = value
	return h
}

// AutoFill sets the autoFill property.
func (h Hidden) AutoFill(value string) Hidden {
	return h.set("autoFill", value)
}

// ClassName sets the container CSS class name.
func (h Hidden) ClassName(value string) Hidden {
	return h.set("className", value)
}

// ClearValueOnHidden sets whether to remove the field value when hidden.
func (h Hidden) ClearValueOnHidden(value bool) Hidden {
	return h.set("clearValueOnHidden", value)
}

// Desc sets the description.
func (h Hidden) Desc(value string) Hidden {
	return h.set("desc", value)
}

// Description sets the description content, supports HTML.
func (h Hidden) Description(value string) Hidden {
	return h.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description.
func (h Hidden) DescriptionClassName(value string) Hidden {
	return h.set("descriptionClassName", value)
}

// Disabled sets whether the field is disabled.
func (h Hidden) Disabled(value bool) Hidden {
	return h.set("disabled", value)
}

// DisabledOn sets the expression to determine if the field is disabled.
func (h Hidden) DisabledOn(value string) Hidden {
	return h.set("disabledOn", value)
}

// EditorSetting sets the editor configuration.
func (h Hidden) EditorSetting(value string) Hidden {
	return h.set("editorSetting", value)
}

// ExtraName sets an additional field name for range components.
func (h Hidden) ExtraName(value string) Hidden {
	return h.set("extraName", value)
}

// Hidden sets whether the field is hidden.
func (h Hidden) Hidden(value bool) Hidden {
	return h.set("hidden", value)
}

// HiddenOn sets the expression to determine if the field is hidden.
func (h Hidden) HiddenOn(value string) Hidden {
	return h.set("hiddenOn", value)
}

// Hint sets the input hint displayed on focus.
func (h Hidden) Hint(value string) Hidden {
	return h.set("hint", value)
}

// Horizontal sets the horizontal layout configuration.
func (h Hidden) Horizontal(value string) Hidden {
	return h.set("horizontal", value)
}

// ID sets the unique component ID.
func (h Hidden) ID(value string) Hidden {
	return h.set("id", value)
}

// InitAutoFill sets the initial autoFill value.
func (h Hidden) InitAutoFill(value string) Hidden {
	return h.set("initAutoFill", value)
}

// Inline sets whether the control is in inline mode.
func (h Hidden) Inline(value bool) Hidden {
	return h.set("inline", value)
}

// InputClassName sets the CSS class name for the input.
func (h Hidden) InputClassName(value string) Hidden {
	return h.set("inputClassName", value)
}

// Label sets the label text.
func (h Hidden) Label(value string) Hidden {
	return h.set("label", value)
}

// LabelAlign sets the label alignment.
func (h Hidden) LabelAlign(value string) Hidden {
	return h.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label.
func (h Hidden) LabelClassName(value string) Hidden {
	return h.set("labelClassName", value)
}

// LabelRemark sets the label remark.
func (h Hidden) LabelRemark(value string) Hidden {
	return h.set("labelRemark", value)
}

// LabelWidth sets the custom width for the label.
func (h Hidden) LabelWidth(value string) Hidden {
	return h.set("labelWidth", value)
}

// Mode sets the display mode of the form item.
func (h Hidden) Mode(value string) Hidden {
	return h.set("mode", value)
}

// Name sets the field name.
func (h Hidden) Name(value string) Hidden {
	return h.set("name", value)
}

// OnEvent sets the event action configuration.
func (h Hidden) OnEvent(value Event) Hidden {
	return h.set("onEvent", value)
}

// Placeholder sets the placeholder text.
func (h Hidden) Placeholder(value string) Hidden {
	return h.set("placeholder", value)
}

// ReadOnly sets whether the field is read-only.
func (h Hidden) ReadOnly(value bool) Hidden {
	return h.set("readOnly", value)
}

// ReadOnlyOn sets the expression to determine if the field is read-only.
func (h Hidden) ReadOnlyOn(value string) Hidden {
	return h.set("readOnlyOn", value)
}

// Remark sets the remark.
func (h Hidden) Remark(value string) Hidden {
	return h.set("remark", value)
}

// Required sets whether the field is required.
func (h Hidden) Required(value bool) Hidden {
	return h.set("required", value)
}

// Row sets the row value.
func (h Hidden) Row(value string) Hidden {
	return h.set("row", value)
}

// SaveImmediately sets whether to save immediately in TableColumn.
func (h Hidden) SaveImmediately(value bool) Hidden {
	return h.set("saveImmediately", value)
}

// Size sets the size of the form item.
func (h Hidden) Size(value string) Hidden {
	return h.set("size", value)
}

// Static sets whether the field is displayed statically.
func (h Hidden) Static(value bool) Hidden {
	return h.set("static", value)
}

// StaticClassName sets the CSS class name for static display.
func (h Hidden) StaticClassName(value string) Hidden {
	return h.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display.
func (h Hidden) StaticInputClassName(value string) Hidden {
	return h.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display.
func (h Hidden) StaticLabelClassName(value string) Hidden {
	return h.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the field is displayed statically.
func (h Hidden) StaticOn(value string) Hidden {
	return h.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display.
func (h Hidden) StaticPlaceholder(value string) Hidden {
	return h.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema.Schema.
func (h Hidden) StaticSchema(value string) Hidden {
	return h.set("staticSchema", value)
}

// Style sets the component style.
func (h Hidden) Style(value any) Hidden {
	return h.set("style", value)
}

// SubmitOnChange sets whether to submit the form on change.
func (h Hidden) SubmitOnChange(value bool) Hidden {
	return h.set("submitOnChange", value)
}

// TestIdBuilder sets the test ID builder.
func (h Hidden) TestIdBuilder(value string) Hidden {
	return h.set("testIdBuilder", value)
}

// UseMobileUI sets whether to use mobile UI styles.
func (h Hidden) UseMobileUI(value bool) Hidden {
	return h.set("useMobileUI", value)
}

// ValidateApi sets the remote validation API.
func (h Hidden) ValidateApi(value string) Hidden {
	return h.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on change.
func (h Hidden) ValidateOnChange(value bool) Hidden {
	return h.set("validateOnChange", value)
}

// ValidationErrors sets the validation error messages.
func (h Hidden) ValidationErrors(value string) Hidden {
	return h.set("validationErrors", value)
}

// Validations sets the validation rules.
func (h Hidden) Validations(value string) Hidden {
	return h.set("validations", value)
}

// Value sets the default value.
func (h Hidden) Value(value string) Hidden {
	return h.set("value", value)
}

// Visible sets whether the field is visible.
func (h Hidden) Visible(value bool) Hidden {
	return h.set("visible", value)
}

// VisibleOn sets the expression to determine if the field is visible.
func (h Hidden) VisibleOn(value string) Hidden {
	return h.set("visibleOn", value)
}

// Width sets the width in a table.
func (h Hidden) Width(value string) Hidden {
	return h.set("width", value)
}
