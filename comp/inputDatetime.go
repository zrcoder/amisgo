package comp

// inputDatetime represents a datetime input control.
type inputDatetime Schema

func InputDatetime() inputDatetime {
	return make(inputDatetime).set("type", "input-datetime")
}

func (d inputDatetime) set(key string, value any) inputDatetime {
	d[key] = value
	return d
}

// AutoFill sets auto-fill value.
func (d inputDatetime) AutoFill(value string) inputDatetime {
	return d.set("autoFill", value)
}

// BorderMode sets border mode: full, half, or none.
func (d inputDatetime) BorderMode(value string) inputDatetime {
	return d.set("borderMode", value)
}

// ClassName sets the CSS class name.
func (d inputDatetime) ClassName(value string) inputDatetime {
	return d.set("className", value)
}

// ClearValueOnHidden removes value when hidden.
func (d inputDatetime) ClearValueOnHidden(value bool) inputDatetime {
	return d.set("clearValueOnHidden", value)
}

// Clearable shows clear button.
func (d inputDatetime) Clearable(value bool) inputDatetime {
	return d.set("clearable", value)
}

// Desc sets description.
func (d inputDatetime) Desc(value string) inputDatetime {
	return d.set("desc", value)
}

// Description sets HTML description.
func (d inputDatetime) Description(value string) inputDatetime {
	return d.set("description", value)
}

// DescriptionClassName sets CSS class for description.
func (d inputDatetime) DescriptionClassName(value string) inputDatetime {
	return d.set("descriptionClassName", value)
}

// Disabled disables the input.
func (d inputDatetime) Disabled(value bool) inputDatetime {
	return d.set("disabled", value)
}

// DisabledDate sets function to disable dates.
func (d inputDatetime) DisabledDate(value string) inputDatetime {
	return d.set("disabledDate", value)
}

// DisabledOn sets expression to disable input.
func (d inputDatetime) DisabledOn(value string) inputDatetime {
	return d.set("disabledOn", value)
}

// DisplayFormat sets display format.
func (d inputDatetime) DisplayFormat(value string) inputDatetime {
	return d.set("displayFormat", value)
}

// EditorSetting sets editor configuration.
func (d inputDatetime) EditorSetting(value string) inputDatetime {
	return d.set("editorSetting", value)
}

// Emebed sets inline mode.
func (d inputDatetime) Emebed(value bool) inputDatetime {
	return d.set("emebed", value)
}

// ExtraName sets extra field name.
func (d inputDatetime) ExtraName(value string) inputDatetime {
	return d.set("extraName", value)
}

// Format sets storage format.
func (d inputDatetime) Format(value string) inputDatetime {
	return d.set("format", value)
}

// Hidden hides the input.
func (d inputDatetime) Hidden(value bool) inputDatetime {
	return d.set("hidden", value)
}

// HiddenOn sets expression to hide input.
func (d inputDatetime) HiddenOn(value string) inputDatetime {
	return d.set("hiddenOn", value)
}

// Hint sets input hint.
func (d inputDatetime) Hint(value string) inputDatetime {
	return d.set("hint", value)
}

// Horizontal sets horizontal layout configuration.
func (d inputDatetime) Horizontal(value string) inputDatetime {
	return d.set("horizontal", value)
}

// ID sets unique component ID.
func (d inputDatetime) ID(value string) inputDatetime {
	return d.set("id", value)
}

// InitAutoFill sets initial auto-fill value.
func (d inputDatetime) InitAutoFill(value string) inputDatetime {
	return d.set("initAutoFill", value)
}

// Inline sets inline mode.
func (d inputDatetime) Inline(value bool) inputDatetime {
	return d.set("inline", value)
}

// InputClassName sets CSS class for input.
func (d inputDatetime) InputClassName(value string) inputDatetime {
	return d.set("inputClassName", value)
}

// InputFormat sets input format.
func (d inputDatetime) InputFormat(value string) inputDatetime {
	return d.set("inputFormat", value)
}

// IsEndDate sets end date flag.
func (d inputDatetime) IsEndDate(value bool) inputDatetime {
	return d.set("isEndDate", value)
}

// Label sets label text.
func (d inputDatetime) Label(value string) inputDatetime {
	return d.set("label", value)
}

// LabelAlign sets label alignment: right, left, top, or inherit.
func (d inputDatetime) LabelAlign(value string) inputDatetime {
	return d.set("labelAlign", value)
}

// LabelClassName sets CSS class for label.
func (d inputDatetime) LabelClassName(value string) inputDatetime {
	return d.set("labelClassName", value)
}

// LabelRemark sets label remark.
func (d inputDatetime) LabelRemark(value string) inputDatetime {
	return d.set("labelRemark", value)
}

// LabelWidth sets label width.
func (d inputDatetime) LabelWidth(value string) inputDatetime {
	return d.set("labelWidth", value)
}

// MaxDate sets maximum date.
func (d inputDatetime) MaxDate(value string) inputDatetime {
	return d.set("maxDate", value)
}

// MinDate sets minimum date.
func (d inputDatetime) MinDate(value string) inputDatetime {
	return d.set("minDate", value)
}

// Mode sets display mode: normal, inline, or horizontal.
func (d inputDatetime) Mode(value string) inputDatetime {
	return d.set("mode", value)
}

// Name sets field name.
func (d inputDatetime) Name(value string) inputDatetime {
	return d.set("name", value)
}

// OnEvent sets event actions.
func (d inputDatetime) OnEvent(value any) inputDatetime {
	return d.set("onEvent", value)
}

// Placeholder sets placeholder text.
func (d inputDatetime) Placeholder(value string) inputDatetime {
	return d.set("placeholder", value)
}

// ReadOnly sets read-only mode.
func (d inputDatetime) ReadOnly(value bool) inputDatetime {
	return d.set("readOnly", value)
}

// ReadOnlyOn sets expression for read-only mode.
func (d inputDatetime) ReadOnlyOn(value string) inputDatetime {
	return d.set("readOnlyOn", value)
}

// Remark sets remark.
func (d inputDatetime) Remark(value string) inputDatetime {
	return d.set("remark", value)
}

// Required sets required flag.
func (d inputDatetime) Required(value bool) inputDatetime {
	return d.set("required", value)
}

// Row sets row value.
func (d inputDatetime) Row(value string) inputDatetime {
	return d.set("row", value)
}

// SaveImmediately sets immediate save flag.
func (d inputDatetime) SaveImmediately(value bool) inputDatetime {
	return d.set("saveImmediately", value)
}

// Shortcuts sets date shortcuts.
func (d inputDatetime) Shortcuts(value string) inputDatetime {
	return d.set("shortcuts", value)
}

// Size sets input size: xs, sm, md, lg, or full.
func (d inputDatetime) Size(value string) inputDatetime {
	return d.set("size", value)
}

// Static sets static display mode.
func (d inputDatetime) Static(value bool) inputDatetime {
	return d.set("static", value)
}

// StaticClassName sets CSS class for static display.
func (d inputDatetime) StaticClassName(value string) inputDatetime {
	return d.set("staticClassName", value)
}

// StaticInputClassName sets CSS class for static input value.
func (d inputDatetime) StaticInputClassName(value string) inputDatetime {
	return d.set("staticInputClassName", value)
}

// StaticLabelClassName sets CSS class for static label.
func (d inputDatetime) StaticLabelClassName(value string) inputDatetime {
	return d.set("staticLabelClassName", value)
}

// StaticOn sets expression for static display.
func (d inputDatetime) StaticOn(value string) inputDatetime {
	return d.set("staticOn", value)
}

// StaticPlaceholder sets placeholder for static display.
func (d inputDatetime) StaticPlaceholder(value string) inputDatetime {
	return d.set("staticPlaceholder", value)
}

// StaticSchema sets static schema.
func (d inputDatetime) StaticSchema(value string) inputDatetime {
	return d.set("staticSchema", value)
}

// Style sets component style.
func (d inputDatetime) Style(value any) inputDatetime {
	return d.set("style", value)
}

// SubmitOnChange sets flag to submit on change.
func (d inputDatetime) SubmitOnChange(value bool) inputDatetime {
	return d.set("submitOnChange", value)
}

// TestIdBuilder sets test ID builder.
func (d inputDatetime) TestIdBuilder(value string) inputDatetime {
	return d.set("testIdBuilder", value)
}

// TimeConstraints sets time constraints.
func (d inputDatetime) TimeConstraints(value string) inputDatetime {
	return d.set("timeConstraints", value)
}

// TimeFormat sets time format.
func (d inputDatetime) TimeFormat(value string) inputDatetime {
	return d.set("timeFormat", value)
}

// UseMobileUI sets flag to use mobile UI.
func (d inputDatetime) UseMobileUI(value bool) inputDatetime {
	return d.set("useMobileUI", value)
}

// Utc sets flag to store UTC time.
func (d inputDatetime) Utc(value bool) inputDatetime {
	return d.set("utc", value)
}

// ValidateApi sets remote validation API.
func (d inputDatetime) ValidateApi(value string) inputDatetime {
	return d.set("validateApi", value)
}

// ValidateOnChange sets flag to validate on change.
func (d inputDatetime) ValidateOnChange(value bool) inputDatetime {
	return d.set("validateOnChange", value)
}

// ValidationErrors sets validation error messages.
func (d inputDatetime) ValidationErrors(value string) inputDatetime {
	return d.set("validationErrors", value)
}

// Validations sets validation rules.
func (d inputDatetime) Validations(value string) inputDatetime {
	return d.set("validations", value)
}

// Value sets default value.
func (d inputDatetime) Value(value string) inputDatetime {
	return d.set("value", value)
}

// ValueFormat sets value format.
func (d inputDatetime) ValueFormat(value string) inputDatetime {
	return d.set("valueFormat", value)
}

// Visible sets visibility.
func (d inputDatetime) Visible(value bool) inputDatetime {
	return d.set("visible", value)
}

// VisibleOn sets expression for visibility.
func (d inputDatetime) VisibleOn(value string) inputDatetime {
	return d.set("visibleOn", value)
}

// Width sets width in table.
func (d inputDatetime) Width(value string) inputDatetime {
	return d.set("width", value)
}
