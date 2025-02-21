package comp

import "github.com/zrcoder/amisgo/schema"

// InputYear represents a year selection component.
type InputYear schema.Schema

func NewInputYear() InputYear {
	return InputYear{"type": "input-year"}
}

// set sets a key-value pair for the inputYear component.
func (yc InputYear) set(key string, value any) InputYear {
	yc[key] = value
	return yc
}

// AutoFill sets the autoFill property.
func (yc InputYear) AutoFill(value string) InputYear {
	return yc.set("autoFill", value)
}

// BorderMode sets the border mode (full, half, none).
func (yc InputYear) BorderMode(value string) InputYear {
	return yc.set("borderMode", value)
}

// ClassName sets the CSS class name.
func (yc InputYear) ClassName(value string) InputYear {
	return yc.set("className", value)
}

// ClearValueOnHidden sets whether to clear the value when hidden.
func (yc InputYear) ClearValueOnHidden(value bool) InputYear {
	return yc.set("clearValueOnHidden", value)
}

// Clearable sets whether the clear button is shown.
func (yc InputYear) Clearable(value bool) InputYear {
	return yc.set("clearable", value)
}

// Desc sets the description.
func (yc InputYear) Desc(value string) InputYear {
	return yc.set("desc", value)
}

// Description sets the HTML description.
func (yc InputYear) Description(value string) InputYear {
	return yc.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description.
func (yc InputYear) DescriptionClassName(value string) InputYear {
	return yc.set("descriptionClassName", value)
}

// Disabled sets whether the component is disabled.
func (yc InputYear) Disabled(value bool) InputYear {
	return yc.set("disabled", value)
}

// DisabledDate sets the function to disable specific dates.
func (yc InputYear) DisabledDate(value string) InputYear {
	return yc.set("disabledDate", value)
}

// DisabledOn sets the expression to disable the component.
func (yc InputYear) DisabledOn(value string) InputYear {
	return yc.set("disabledOn", value)
}

// DisplayFormat sets the display format for the date.
func (yc InputYear) DisplayFormat(value string) InputYear {
	return yc.set("displayFormat", value)
}

// EditorSetting sets the editor configuration.
func (yc InputYear) EditorSetting(value string) InputYear {
	return yc.set("editorSetting", value)
}

// Emebed sets whether the component is inline.
func (yc InputYear) Emebed(value bool) InputYear {
	return yc.set("emebed", value)
}

// ExtraName sets an additional field name for range components.
func (yc InputYear) ExtraName(value string) InputYear {
	return yc.set("extraName", value)
}

// Format sets the storage format for the month.
func (yc InputYear) Format(value string) InputYear {
	return yc.set("format", value)
}

// Hidden sets whether the component is hidden.
func (yc InputYear) Hidden(value bool) InputYear {
	return yc.set("hidden", value)
}

// HiddenOn sets the expression to hide the component.
func (yc InputYear) HiddenOn(value string) InputYear {
	return yc.set("hiddenOn", value)
}

// Hint sets the input hint shown on focus.
func (yc InputYear) Hint(value string) InputYear {
	return yc.set("hint", value)
}

// Horizontal sets the horizontal layout configuration.
func (yc InputYear) Horizontal(value string) InputYear {
	return yc.set("horizontal", value)
}

// ID sets the unique component ID.
func (yc InputYear) ID(value string) InputYear {
	return yc.set("id", value)
}

// InitAutoFill sets the initial autoFill value.
func (yc InputYear) InitAutoFill(value string) InputYear {
	return yc.set("initAutoFill", value)
}

// Inline sets whether the form control is inline.
func (yc InputYear) Inline(value bool) InputYear {
	return yc.set("inline", value)
}

// InputClassName sets the CSS class name for the input.
func (yc InputYear) InputClassName(value string) InputYear {
	return yc.set("inputClassName", value)
}

// InputFormat sets the display format for the month.
func (yc InputYear) InputFormat(value string) InputYear {
	return yc.set("inputFormat", value)
}

// Label sets the label text.
func (yc InputYear) Label(value string) InputYear {
	return yc.set("label", value)
}

// LabelAlign sets the label alignment (right, left, top, inherit).
func (yc InputYear) LabelAlign(value string) InputYear {
	return yc.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label.
func (yc InputYear) LabelClassName(value string) InputYear {
	return yc.set("labelClassName", value)
}

// LabelRemark sets the label remark.
func (yc InputYear) LabelRemark(value string) InputYear {
	return yc.set("labelRemark", value)
}

// LabelWidth sets the custom width for the label.
func (yc InputYear) LabelWidth(value string) InputYear {
	return yc.set("labelWidth", value)
}

// Mode sets the display mode (normal, inline, horizontal).
func (yc InputYear) Mode(value string) InputYear {
	return yc.set("mode", value)
}

// Name sets the field name for form submission.
func (yc InputYear) Name(value string) InputYear {
	return yc.set("name", value)
}

// OnEvent sets the event action configuration.
func (yc InputYear) OnEvent(value Event) InputYear {
	return yc.set("onEvent", value)
}

// Placeholder sets the placeholder text.
func (yc InputYear) Placeholder(value string) InputYear {
	return yc.set("placeholder", value)
}

// ReadOnly sets whether the component is read-only.
func (yc InputYear) ReadOnly(value bool) InputYear {
	return yc.set("readOnly", value)
}

// ReadOnlyOn sets the expression for read-only condition.
func (yc InputYear) ReadOnlyOn(value string) InputYear {
	return yc.set("readOnlyOn", value)
}

// Remark sets the remark text.
func (yc InputYear) Remark(value string) InputYear {
	return yc.set("remark", value)
}

// Required sets whether the field is required.
func (yc InputYear) Required(value bool) InputYear {
	return yc.set("required", value)
}

// Row sets the row configuration.
func (yc InputYear) Row(value string) InputYear {
	return yc.set("row", value)
}

// SaveImmediately sets whether to save immediately (used in TableColumn).
func (yc InputYear) SaveImmediately(value bool) InputYear {
	return yc.set("saveImmediately", value)
}

// Shortcuts sets the date shortcuts.
func (yc InputYear) Shortcuts(value string) InputYear {
	return yc.set("shortcuts", value)
}

// Size sets the size of the form item (xs, sm, md, lg, full).
func (yc InputYear) Size(value string) InputYear {
	return yc.set("size", value)
}

// Static sets whether the component is displayed statically.
func (yc InputYear) Static(value bool) InputYear {
	return yc.set("static", value)
}

// StaticClassName sets the CSS class name for static display.
func (yc InputYear) StaticClassName(value string) InputYear {
	return yc.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display.
func (yc InputYear) StaticInputClassName(value string) InputYear {
	return yc.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display.
func (yc InputYear) StaticLabelClassName(value string) InputYear {
	return yc.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display.
func (yc InputYear) StaticOn(value string) InputYear {
	return yc.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display.
func (yc InputYear) StaticPlaceholder(value string) InputYear {
	return yc.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema.Schema.
func (yc InputYear) StaticSchema(value string) InputYear {
	return yc.set("staticSchema", value)
}

// Style sets the component style.
func (yc InputYear) Style(value any) InputYear {
	return yc.set("style", value)
}

// SubmitOnChange sets whether to submit the form on change.
func (yc InputYear) SubmitOnChange(value bool) InputYear {
	return yc.set("submitOnChange", value)
}

// TestIdBuilder sets the test ID builder.
func (yc InputYear) TestIdBuilder(value string) InputYear {
	return yc.set("testIdBuilder", value)
}

// UseMobileUI sets whether to use mobile UI.
func (yc InputYear) UseMobileUI(value bool) InputYear {
	return yc.set("useMobileUI", value)
}

// Utc sets whether to store UTC time.
func (yc InputYear) Utc(value bool) InputYear {
	return yc.set("utc", value)
}

// ValidateApi sets the remote validation API.
func (yc InputYear) ValidateApi(value string) InputYear {
	return yc.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on change.
func (yc InputYear) ValidateOnChange(value bool) InputYear {
	return yc.set("validateOnChange", value)
}

// ValidationErrors sets the validation error messages.
func (yc InputYear) ValidationErrors(value string) InputYear {
	return yc.set("validationErrors", value)
}

// Validations sets the validation rules.
func (yc InputYear) Validations(value string) InputYear {
	return yc.set("validations", value)
}

// Value sets the default value.
func (yc InputYear) Value(value string) InputYear {
	return yc.set("value", value)
}

// ValueFormat sets the value format.
func (yc InputYear) ValueFormat(value string) InputYear {
	return yc.set("valueFormat", value)
}

// Visible sets whether the component is visible.
func (yc InputYear) Visible(value bool) InputYear {
	return yc.set("visible", value)
}

// VisibleOn sets the expression for visibility.
func (yc InputYear) VisibleOn(value string) InputYear {
	return yc.set("visibleOn", value)
}

// Width sets the width in a table.
func (yc InputYear) Width(value string) InputYear {
	return yc.set("width", value)
}
