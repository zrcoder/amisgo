package comp

import "github.com/zrcoder/amisgo/schema"

// InputDatetime represents a datetime input control.
type InputDatetime schema.Schema

func NewInputDatetime() InputDatetime {
	return InputDatetime{"type": "input-datetime"}
}

func (d InputDatetime) set(key string, value any) InputDatetime {
	d[key] = value
	return d
}

// AutoFill sets auto-fill value.
func (d InputDatetime) AutoFill(value string) InputDatetime {
	return d.set("autoFill", value)
}

// BorderMode sets border mode: full, half, or none.
func (d InputDatetime) BorderMode(value string) InputDatetime {
	return d.set("borderMode", value)
}

// ClassName sets the CSS class name.
func (d InputDatetime) ClassName(value string) InputDatetime {
	return d.set("className", value)
}

// ClearValueOnHidden removes value when hidden.
func (d InputDatetime) ClearValueOnHidden(value bool) InputDatetime {
	return d.set("clearValueOnHidden", value)
}

// Clearable shows clear button.
func (d InputDatetime) Clearable(value bool) InputDatetime {
	return d.set("clearable", value)
}

// Desc sets description.
func (d InputDatetime) Desc(value string) InputDatetime {
	return d.set("desc", value)
}

// Description sets HTML description.
func (d InputDatetime) Description(value string) InputDatetime {
	return d.set("description", value)
}

// DescriptionClassName sets CSS class for description.
func (d InputDatetime) DescriptionClassName(value string) InputDatetime {
	return d.set("descriptionClassName", value)
}

// Disabled disables the input.
func (d InputDatetime) Disabled(value bool) InputDatetime {
	return d.set("disabled", value)
}

// DisabledDate sets function to disable dates.
func (d InputDatetime) DisabledDate(value string) InputDatetime {
	return d.set("disabledDate", value)
}

// DisabledOn sets expression to disable input.
func (d InputDatetime) DisabledOn(value string) InputDatetime {
	return d.set("disabledOn", value)
}

// DisplayFormat sets display format.
func (d InputDatetime) DisplayFormat(value string) InputDatetime {
	return d.set("displayFormat", value)
}

// EditorSetting sets editor configuration.
func (d InputDatetime) EditorSetting(value string) InputDatetime {
	return d.set("editorSetting", value)
}

// Emebed sets inline mode.
func (d InputDatetime) Emebed(value bool) InputDatetime {
	return d.set("emebed", value)
}

// ExtraName sets extra field name.
func (d InputDatetime) ExtraName(value string) InputDatetime {
	return d.set("extraName", value)
}

// Format sets storage format.
func (d InputDatetime) Format(value string) InputDatetime {
	return d.set("format", value)
}

// Hidden hides the input.
func (d InputDatetime) Hidden(value bool) InputDatetime {
	return d.set("hidden", value)
}

// HiddenOn sets expression to hide input.
func (d InputDatetime) HiddenOn(value string) InputDatetime {
	return d.set("hiddenOn", value)
}

// Hint sets input hint.
func (d InputDatetime) Hint(value string) InputDatetime {
	return d.set("hint", value)
}

// Horizontal sets horizontal layout configuration.
func (d InputDatetime) Horizontal(value string) InputDatetime {
	return d.set("horizontal", value)
}

// ID sets unique component ID.
func (d InputDatetime) ID(value string) InputDatetime {
	return d.set("id", value)
}

// InitAutoFill sets initial auto-fill value.
func (d InputDatetime) InitAutoFill(value string) InputDatetime {
	return d.set("initAutoFill", value)
}

// Inline sets inline mode.
func (d InputDatetime) Inline(value bool) InputDatetime {
	return d.set("inline", value)
}

// InputClassName sets CSS class for input.
func (d InputDatetime) InputClassName(value string) InputDatetime {
	return d.set("inputClassName", value)
}

// InputFormat sets input format.
func (d InputDatetime) InputFormat(value string) InputDatetime {
	return d.set("inputFormat", value)
}

// IsEndDate sets end date flag.
func (d InputDatetime) IsEndDate(value bool) InputDatetime {
	return d.set("isEndDate", value)
}

// Label sets label text.
func (d InputDatetime) Label(value string) InputDatetime {
	return d.set("label", value)
}

// LabelAlign sets label alignment: right, left, top, or inherit.
func (d InputDatetime) LabelAlign(value string) InputDatetime {
	return d.set("labelAlign", value)
}

// LabelClassName sets CSS class for label.
func (d InputDatetime) LabelClassName(value string) InputDatetime {
	return d.set("labelClassName", value)
}

// LabelRemark sets label remark.
func (d InputDatetime) LabelRemark(value string) InputDatetime {
	return d.set("labelRemark", value)
}

// LabelWidth sets label width.
func (d InputDatetime) LabelWidth(value string) InputDatetime {
	return d.set("labelWidth", value)
}

// MaxDate sets maximum date.
func (d InputDatetime) MaxDate(value string) InputDatetime {
	return d.set("maxDate", value)
}

// MinDate sets minimum date.
func (d InputDatetime) MinDate(value string) InputDatetime {
	return d.set("minDate", value)
}

// Mode sets display mode: normal, inline, or horizontal.
func (d InputDatetime) Mode(value string) InputDatetime {
	return d.set("mode", value)
}

// Name sets field name.
func (d InputDatetime) Name(value string) InputDatetime {
	return d.set("name", value)
}

// OnEvent sets event actions.
func (d InputDatetime) OnEvent(value Event) InputDatetime {
	return d.set("onEvent", value)
}

// Placeholder sets placeholder text.
func (d InputDatetime) Placeholder(value string) InputDatetime {
	return d.set("placeholder", value)
}

// ReadOnly sets read-only mode.
func (d InputDatetime) ReadOnly(value bool) InputDatetime {
	return d.set("readOnly", value)
}

// ReadOnlyOn sets expression for read-only mode.
func (d InputDatetime) ReadOnlyOn(value string) InputDatetime {
	return d.set("readOnlyOn", value)
}

// Remark sets remark.
func (d InputDatetime) Remark(value string) InputDatetime {
	return d.set("remark", value)
}

// Required sets required flag.
func (d InputDatetime) Required(value bool) InputDatetime {
	return d.set("required", value)
}

// Row sets row value.
func (d InputDatetime) Row(value string) InputDatetime {
	return d.set("row", value)
}

// SaveImmediately sets immediate save flag.
func (d InputDatetime) SaveImmediately(value bool) InputDatetime {
	return d.set("saveImmediately", value)
}

// Shortcuts sets date shortcuts.
func (d InputDatetime) Shortcuts(value string) InputDatetime {
	return d.set("shortcuts", value)
}

// Size sets input size: xs, sm, md, lg, or full.
func (d InputDatetime) Size(value string) InputDatetime {
	return d.set("size", value)
}

// Static sets static display mode.
func (d InputDatetime) Static(value bool) InputDatetime {
	return d.set("static", value)
}

// StaticClassName sets CSS class for static display.
func (d InputDatetime) StaticClassName(value string) InputDatetime {
	return d.set("staticClassName", value)
}

// StaticInputClassName sets CSS class for static input value.
func (d InputDatetime) StaticInputClassName(value string) InputDatetime {
	return d.set("staticInputClassName", value)
}

// StaticLabelClassName sets CSS class for static label.
func (d InputDatetime) StaticLabelClassName(value string) InputDatetime {
	return d.set("staticLabelClassName", value)
}

// StaticOn sets expression for static display.
func (d InputDatetime) StaticOn(value string) InputDatetime {
	return d.set("staticOn", value)
}

// StaticPlaceholder sets placeholder for static display.
func (d InputDatetime) StaticPlaceholder(value string) InputDatetime {
	return d.set("staticPlaceholder", value)
}

// StaticSchema sets static schema.Schema.
func (d InputDatetime) StaticSchema(value string) InputDatetime {
	return d.set("staticSchema", value)
}

// Style sets component style.
func (d InputDatetime) Style(value any) InputDatetime {
	return d.set("style", value)
}

// SubmitOnChange sets flag to submit on change.
func (d InputDatetime) SubmitOnChange(value bool) InputDatetime {
	return d.set("submitOnChange", value)
}

// TestIdBuilder sets test ID builder.
func (d InputDatetime) TestIdBuilder(value string) InputDatetime {
	return d.set("testIdBuilder", value)
}

// TimeConstraints sets time constraints.
func (d InputDatetime) TimeConstraints(value string) InputDatetime {
	return d.set("timeConstraints", value)
}

// TimeFormat sets time format.
func (d InputDatetime) TimeFormat(value string) InputDatetime {
	return d.set("timeFormat", value)
}

// UseMobileUI sets flag to use mobile UI.
func (d InputDatetime) UseMobileUI(value bool) InputDatetime {
	return d.set("useMobileUI", value)
}

// Utc sets flag to store UTC time.
func (d InputDatetime) Utc(value bool) InputDatetime {
	return d.set("utc", value)
}

// ValidateApi sets remote validation API.
func (d InputDatetime) ValidateApi(value string) InputDatetime {
	return d.set("validateApi", value)
}

// ValidateOnChange sets flag to validate on change.
func (d InputDatetime) ValidateOnChange(value bool) InputDatetime {
	return d.set("validateOnChange", value)
}

// ValidationErrors sets validation error messages.
func (d InputDatetime) ValidationErrors(value string) InputDatetime {
	return d.set("validationErrors", value)
}

// Validations sets validation rules.
func (d InputDatetime) Validations(value string) InputDatetime {
	return d.set("validations", value)
}

// Value sets default value.
func (d InputDatetime) Value(value string) InputDatetime {
	return d.set("value", value)
}

// ValueFormat sets value format.
func (d InputDatetime) ValueFormat(value string) InputDatetime {
	return d.set("valueFormat", value)
}

// Visible sets visibility.
func (d InputDatetime) Visible(value bool) InputDatetime {
	return d.set("visible", value)
}

// VisibleOn sets expression for visibility.
func (d InputDatetime) VisibleOn(value string) InputDatetime {
	return d.set("visibleOn", value)
}

// Width sets width in table.
func (d InputDatetime) Width(value string) InputDatetime {
	return d.set("width", value)
}
