package comp

import "github.com/zrcoder/amisgo/model"

// inputYear represents a year selection component.

type inputYear model.Schema

// InputYear creates a new input-year component.
func InputYear() inputYear {
	return inputYear{"type": "input-year"}
}

// set sets a key-value pair for the inputYear component.
func (yc inputYear) set(key string, value any) inputYear {
	yc[key] = value
	return yc
}

// AutoFill sets the autoFill property.
func (yc inputYear) AutoFill(value string) inputYear {
	return yc.set("autoFill", value)
}

// BorderMode sets the border mode (full, half, none).
func (yc inputYear) BorderMode(value string) inputYear {
	return yc.set("borderMode", value)
}

// ClassName sets the CSS class name.
func (yc inputYear) ClassName(value string) inputYear {
	return yc.set("className", value)
}

// ClearValueOnHidden sets whether to clear the value when hidden.
func (yc inputYear) ClearValueOnHidden(value bool) inputYear {
	return yc.set("clearValueOnHidden", value)
}

// Clearable sets whether the clear button is shown.
func (yc inputYear) Clearable(value bool) inputYear {
	return yc.set("clearable", value)
}

// Desc sets the description.
func (yc inputYear) Desc(value string) inputYear {
	return yc.set("desc", value)
}

// Description sets the HTML description.
func (yc inputYear) Description(value string) inputYear {
	return yc.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description.
func (yc inputYear) DescriptionClassName(value string) inputYear {
	return yc.set("descriptionClassName", value)
}

// Disabled sets whether the component is disabled.
func (yc inputYear) Disabled(value bool) inputYear {
	return yc.set("disabled", value)
}

// DisabledDate sets the function to disable specific dates.
func (yc inputYear) DisabledDate(value string) inputYear {
	return yc.set("disabledDate", value)
}

// DisabledOn sets the expression to disable the component.
func (yc inputYear) DisabledOn(value string) inputYear {
	return yc.set("disabledOn", value)
}

// DisplayFormat sets the display format for the date.
func (yc inputYear) DisplayFormat(value string) inputYear {
	return yc.set("displayFormat", value)
}

// EditorSetting sets the editor configuration.
func (yc inputYear) EditorSetting(value string) inputYear {
	return yc.set("editorSetting", value)
}

// Emebed sets whether the component is inline.
func (yc inputYear) Emebed(value bool) inputYear {
	return yc.set("emebed", value)
}

// ExtraName sets an additional field name for range components.
func (yc inputYear) ExtraName(value string) inputYear {
	return yc.set("extraName", value)
}

// Format sets the storage format for the month.
func (yc inputYear) Format(value string) inputYear {
	return yc.set("format", value)
}

// Hidden sets whether the component is hidden.
func (yc inputYear) Hidden(value bool) inputYear {
	return yc.set("hidden", value)
}

// HiddenOn sets the expression to hide the component.
func (yc inputYear) HiddenOn(value string) inputYear {
	return yc.set("hiddenOn", value)
}

// Hint sets the input hint shown on focus.
func (yc inputYear) Hint(value string) inputYear {
	return yc.set("hint", value)
}

// Horizontal sets the horizontal layout configuration.
func (yc inputYear) Horizontal(value string) inputYear {
	return yc.set("horizontal", value)
}

// ID sets the unique component ID.
func (yc inputYear) ID(value string) inputYear {
	return yc.set("id", value)
}

// InitAutoFill sets the initial autoFill value.
func (yc inputYear) InitAutoFill(value string) inputYear {
	return yc.set("initAutoFill", value)
}

// Inline sets whether the form control is inline.
func (yc inputYear) Inline(value bool) inputYear {
	return yc.set("inline", value)
}

// InputClassName sets the CSS class name for the input.
func (yc inputYear) InputClassName(value string) inputYear {
	return yc.set("inputClassName", value)
}

// InputFormat sets the display format for the month.
func (yc inputYear) InputFormat(value string) inputYear {
	return yc.set("inputFormat", value)
}

// Label sets the label text.
func (yc inputYear) Label(value string) inputYear {
	return yc.set("label", value)
}

// LabelAlign sets the label alignment (right, left, top, inherit).
func (yc inputYear) LabelAlign(value string) inputYear {
	return yc.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label.
func (yc inputYear) LabelClassName(value string) inputYear {
	return yc.set("labelClassName", value)
}

// LabelRemark sets the label remark.
func (yc inputYear) LabelRemark(value string) inputYear {
	return yc.set("labelRemark", value)
}

// LabelWidth sets the custom width for the label.
func (yc inputYear) LabelWidth(value string) inputYear {
	return yc.set("labelWidth", value)
}

// Mode sets the display mode (normal, inline, horizontal).
func (yc inputYear) Mode(value string) inputYear {
	return yc.set("mode", value)
}

// Name sets the field name for form submission.
func (yc inputYear) Name(value string) inputYear {
	return yc.set("name", value)
}

// OnEvent sets the event action configuration.
func (yc inputYear) OnEvent(value any) inputYear {
	return yc.set("onEvent", value)
}

// Placeholder sets the placeholder text.
func (yc inputYear) Placeholder(value string) inputYear {
	return yc.set("placeholder", value)
}

// ReadOnly sets whether the component is read-only.
func (yc inputYear) ReadOnly(value bool) inputYear {
	return yc.set("readOnly", value)
}

// ReadOnlyOn sets the expression for read-only condition.
func (yc inputYear) ReadOnlyOn(value string) inputYear {
	return yc.set("readOnlyOn", value)
}

// Remark sets the remark text.
func (yc inputYear) Remark(value string) inputYear {
	return yc.set("remark", value)
}

// Required sets whether the field is required.
func (yc inputYear) Required(value bool) inputYear {
	return yc.set("required", value)
}

// Row sets the row configuration.
func (yc inputYear) Row(value string) inputYear {
	return yc.set("row", value)
}

// SaveImmediately sets whether to save immediately (used in TableColumn).
func (yc inputYear) SaveImmediately(value bool) inputYear {
	return yc.set("saveImmediately", value)
}

// Shortcuts sets the date shortcuts.
func (yc inputYear) Shortcuts(value string) inputYear {
	return yc.set("shortcuts", value)
}

// Size sets the size of the form item (xs, sm, md, lg, full).
func (yc inputYear) Size(value string) inputYear {
	return yc.set("size", value)
}

// Static sets whether the component is displayed statically.
func (yc inputYear) Static(value bool) inputYear {
	return yc.set("static", value)
}

// StaticClassName sets the CSS class name for static display.
func (yc inputYear) StaticClassName(value string) inputYear {
	return yc.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display.
func (yc inputYear) StaticInputClassName(value string) inputYear {
	return yc.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display.
func (yc inputYear) StaticLabelClassName(value string) inputYear {
	return yc.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display.
func (yc inputYear) StaticOn(value string) inputYear {
	return yc.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display.
func (yc inputYear) StaticPlaceholder(value string) inputYear {
	return yc.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema.
func (yc inputYear) StaticSchema(value string) inputYear {
	return yc.set("staticSchema", value)
}

// Style sets the component style.
func (yc inputYear) Style(value any) inputYear {
	return yc.set("style", value)
}

// SubmitOnChange sets whether to submit the form on change.
func (yc inputYear) SubmitOnChange(value bool) inputYear {
	return yc.set("submitOnChange", value)
}

// TestIdBuilder sets the test ID builder.
func (yc inputYear) TestIdBuilder(value string) inputYear {
	return yc.set("testIdBuilder", value)
}

// UseMobileUI sets whether to use mobile UI.
func (yc inputYear) UseMobileUI(value bool) inputYear {
	return yc.set("useMobileUI", value)
}

// Utc sets whether to store UTC time.
func (yc inputYear) Utc(value bool) inputYear {
	return yc.set("utc", value)
}

// ValidateApi sets the remote validation API.
func (yc inputYear) ValidateApi(value string) inputYear {
	return yc.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on change.
func (yc inputYear) ValidateOnChange(value bool) inputYear {
	return yc.set("validateOnChange", value)
}

// ValidationErrors sets the validation error messages.
func (yc inputYear) ValidationErrors(value string) inputYear {
	return yc.set("validationErrors", value)
}

// Validations sets the validation rules.
func (yc inputYear) Validations(value string) inputYear {
	return yc.set("validations", value)
}

// Value sets the default value.
func (yc inputYear) Value(value string) inputYear {
	return yc.set("value", value)
}

// ValueFormat sets the value format.
func (yc inputYear) ValueFormat(value string) inputYear {
	return yc.set("valueFormat", value)
}

// Visible sets whether the component is visible.
func (yc inputYear) Visible(value bool) inputYear {
	return yc.set("visible", value)
}

// VisibleOn sets the expression for visibility.
func (yc inputYear) VisibleOn(value string) inputYear {
	return yc.set("visibleOn", value)
}

// Width sets the width in a table.
func (yc inputYear) Width(value string) inputYear {
	return yc.set("width", value)
}
