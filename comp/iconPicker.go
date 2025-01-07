package comp

// iconPicker represents an icon picker component
type iconPicker Schema

// IconPicker creates a new IconPicker instance
func IconPicker() iconPicker {
	return make(iconPicker).set("type", "icon-picker")
}

func (i iconPicker) set(key string, value any) iconPicker {
	i[key] = value
	return i
}

// AutoFill sets autoFill value
func (i iconPicker) AutoFill(value string) iconPicker {
	return i.set("autoFill", value)
}

// ClassName sets container CSS class name
func (i iconPicker) ClassName(value string) iconPicker {
	return i.set("className", value)
}

// ClearValueOnHidden clears value when hidden
func (i iconPicker) ClearValueOnHidden(value bool) iconPicker {
	return i.set("clearValueOnHidden", value)
}

// Desc sets description
func (i iconPicker) Desc(value string) iconPicker {
	return i.set("desc", value)
}

// Description sets HTML description
func (i iconPicker) Description(value string) iconPicker {
	return i.set("description", value)
}

// DescriptionClassName sets description CSS class name
func (i iconPicker) DescriptionClassName(value string) iconPicker {
	return i.set("descriptionClassName", value)
}

// Disabled sets disabled state
func (i iconPicker) Disabled(value bool) iconPicker {
	return i.set("disabled", value)
}

// DisabledOn sets disabled expression
func (i iconPicker) DisabledOn(value string) iconPicker {
	return i.set("disabledOn", value)
}

// EditorSetting sets editor configuration
func (i iconPicker) EditorSetting(value string) iconPicker {
	return i.set("editorSetting", value)
}

// ExtraName sets extra field name
func (i iconPicker) ExtraName(value string) iconPicker {
	return i.set("extraName", value)
}

// Hidden sets hidden state
func (i iconPicker) Hidden(value bool) iconPicker {
	return i.set("hidden", value)
}

// HiddenOn sets hidden expression
func (i iconPicker) HiddenOn(value string) iconPicker {
	return i.set("hiddenOn", value)
}

// Hint sets input hint
func (i iconPicker) Hint(value string) iconPicker {
	return i.set("hint", value)
}

// Horizontal sets horizontal layout configuration
func (i iconPicker) Horizontal(value string) iconPicker {
	return i.set("horizontal", value)
}

// ID sets unique component ID
func (i iconPicker) ID(value string) iconPicker {
	return i.set("id", value)
}

// InitAutoFill sets initial autoFill value
func (i iconPicker) InitAutoFill(value string) iconPicker {
	return i.set("initAutoFill", value)
}

// Inline sets inline mode for form control
func (i iconPicker) Inline(value bool) iconPicker {
	return i.set("inline", value)
}

// InputClassName sets input CSS class name
func (i iconPicker) InputClassName(value string) iconPicker {
	return i.set("inputClassName", value)
}

// Label sets label text
func (i iconPicker) Label(value string) iconPicker {
	return i.set("label", value)
}

// LabelAlign sets label alignment
func (i iconPicker) LabelAlign(value string) iconPicker {
	return i.set("labelAlign", value)
}

// LabelClassName sets label CSS class name
func (i iconPicker) LabelClassName(value string) iconPicker {
	return i.set("labelClassName", value)
}

// LabelRemark sets label remark
func (i iconPicker) LabelRemark(value string) iconPicker {
	return i.set("labelRemark", value)
}

// LabelWidth sets custom label width
func (i iconPicker) LabelWidth(value string) iconPicker {
	return i.set("labelWidth", value)
}

// Mode sets display mode
func (i iconPicker) Mode(value string) iconPicker {
	return i.set("mode", value)
}

// Name sets field name
func (i iconPicker) Name(value string) iconPicker {
	return i.set("name", value)
}

// OnEvent sets event configuration
func (i iconPicker) OnEvent(value any) iconPicker {
	return i.set("onEvent", value)
}

// Placeholder sets placeholder text
func (i iconPicker) Placeholder(value string) iconPicker {
	return i.set("placeholder", value)
}

// ReadOnly sets read-only state
func (i iconPicker) ReadOnly(value bool) iconPicker {
	return i.set("readOnly", value)
}

// ReadOnlyOn sets read-only expression
func (i iconPicker) ReadOnlyOn(value string) iconPicker {
	return i.set("readOnlyOn", value)
}

// Remark sets remark text
func (i iconPicker) Remark(value string) iconPicker {
	return i.set("remark", value)
}

// Required sets required state
func (i iconPicker) Required(value bool) iconPicker {
	return i.set("required", value)
}

// Row sets row value
func (i iconPicker) Row(value string) iconPicker {
	return i.set("row", value)
}

// SaveImmediately sets save immediately state
func (i iconPicker) SaveImmediately(value bool) iconPicker {
	return i.set("saveImmediately", value)
}

// Size sets component size
func (i iconPicker) Size(value string) iconPicker {
	return i.set("size", value)
}

// Static sets static display state
func (i iconPicker) Static(value bool) iconPicker {
	return i.set("static", value)
}

// StaticClassName sets static display CSS class name
func (i iconPicker) StaticClassName(value string) iconPicker {
	return i.set("staticClassName", value)
}

// StaticInputClassName sets static input CSS class name
func (i iconPicker) StaticInputClassName(value string) iconPicker {
	return i.set("staticInputClassName", value)
}

// StaticLabelClassName sets static label CSS class name
func (i iconPicker) StaticLabelClassName(value string) iconPicker {
	return i.set("staticLabelClassName", value)
}

// StaticOn sets static display expression
func (i iconPicker) StaticOn(value string) iconPicker {
	return i.set("staticOn", value)
}

// StaticPlaceholder sets static placeholder text
func (i iconPicker) StaticPlaceholder(value string) iconPicker {
	return i.set("staticPlaceholder", value)
}

// StaticSchema sets static schema
func (i iconPicker) StaticSchema(value string) iconPicker {
	return i.set("staticSchema", value)
}

// Style sets component style
func (i iconPicker) Style(value any) iconPicker {
	return i.set("style", value)
}

// SubmitOnChange sets submit on change state
func (i iconPicker) SubmitOnChange(value bool) iconPicker {
	return i.set("submitOnChange", value)
}

// TestIdBuilder sets test ID builder
func (i iconPicker) TestIdBuilder(value string) iconPicker {
	return i.set("testIdBuilder", value)
}

// UseMobileUI sets mobile UI state
func (i iconPicker) UseMobileUI(value bool) iconPicker {
	return i.set("useMobileUI", value)
}

// ValidateApi sets remote validation API
func (i iconPicker) ValidateApi(value string) iconPicker {
	return i.set("validateApi", value)
}

// ValidateOnChange sets validate on change state
func (i iconPicker) ValidateOnChange(value bool) iconPicker {
	return i.set("validateOnChange", value)
}

// ValidationErrors sets validation error messages
func (i iconPicker) ValidationErrors(value string) iconPicker {
	return i.set("validationErrors", value)
}

// Validations sets validation rules
func (i iconPicker) Validations(value string) iconPicker {
	return i.set("validations", value)
}

// Value sets default value
func (i iconPicker) Value(value string) iconPicker {
	return i.set("value", value)
}

// Visible sets visibility state
func (i iconPicker) Visible(value bool) iconPicker {
	return i.set("visible", value)
}

// VisibleOn sets visibility expression
func (i iconPicker) VisibleOn(value string) iconPicker {
	return i.set("visibleOn", value)
}

// Width sets width in table
func (i iconPicker) Width(value string) iconPicker {
	return i.set("width", value)
}
