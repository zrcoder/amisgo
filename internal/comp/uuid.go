package comp

import "github.com/zrcoder/amisgo/schema"

// UUID UUID functional component
type UUID schema.Schema

func NewUUID() UUID {
	return UUID{"type": "uuid"}
}

func (uc UUID) set(key string, value any) UUID {
	uc[key] = value
	return uc
}

// AutoFill sets autoFill value
func (uc UUID) AutoFill(value string) UUID {
	return uc.set("autoFill", value)
}

// ClassName sets container CSS class name
func (uc UUID) ClassName(value string) UUID {
	return uc.set("className", value)
}

// ClearValueOnHidden removes value when hidden
func (uc UUID) ClearValueOnHidden(value bool) UUID {
	return uc.set("clearValueOnHidden", value)
}

// Desc sets description
func (uc UUID) Desc(value string) UUID {
	return uc.set("desc", value)
}

// Description sets HTML description
func (uc UUID) Description(value string) UUID {
	return uc.set("description", value)
}

// DescriptionClassName sets description CSS class name
func (uc UUID) DescriptionClassName(value string) UUID {
	return uc.set("descriptionClassName", value)
}

// Disabled sets disabled state
func (uc UUID) Disabled(value bool) UUID {
	return uc.set("disabled", value)
}

// DisabledOn sets disabled expression
func (uc UUID) DisabledOn(value string) UUID {
	return uc.set("disabledOn", value)
}

// EditorSetting sets editor configuration
func (uc UUID) EditorSetting(value string) UUID {
	return uc.set("editorSetting", value)
}

// ExtraName sets extra field name
func (uc UUID) ExtraName(value string) UUID {
	return uc.set("extraName", value)
}

// Hidden sets hidden state
func (uc UUID) Hidden(value bool) UUID {
	return uc.set("hidden", value)
}

// HiddenOn sets hidden expression
func (uc UUID) HiddenOn(value string) UUID {
	return uc.set("hiddenOn", value)
}

// Hint sets input hint
func (uc UUID) Hint(value string) UUID {
	return uc.set("hint", value)
}

// Horizontal sets horizontal layout configuration
func (uc UUID) Horizontal(value string) UUID {
	return uc.set("horizontal", value)
}

// ID sets unique component ID
func (uc UUID) ID(value string) UUID {
	return uc.set("id", value)
}

// InitAutoFill sets initial autoFill value
func (uc UUID) InitAutoFill(value string) UUID {
	return uc.set("initAutoFill", value)
}

// Inline sets inline mode
func (uc UUID) Inline(value bool) UUID {
	return uc.set("inline", value)
}

// InputClassName sets input CSS class name
func (uc UUID) InputClassName(value string) UUID {
	return uc.set("inputClassName", value)
}

// Label sets label text
func (uc UUID) Label(value string) UUID {
	return uc.set("label", value)
}

// LabelAlign sets label alignment
func (uc UUID) LabelAlign(value string) UUID {
	return uc.set("labelAlign", value)
}

// LabelClassName sets label CSS class name
func (uc UUID) LabelClassName(value string) UUID {
	return uc.set("labelClassName", value)
}

// LabelRemark sets label remark
func (uc UUID) LabelRemark(value string) UUID {
	return uc.set("labelRemark", value)
}

// LabelWidth sets label width
func (uc UUID) LabelWidth(value string) UUID {
	return uc.set("labelWidth", value)
}

// Length sets UUID length
func (uc UUID) Length(value string) UUID {
	return uc.set("length", value)
}

// Mode sets display mode
func (uc UUID) Mode(value string) UUID {
	return uc.set("mode", value)
}

// Name sets field name
func (uc UUID) Name(value string) UUID {
	return uc.set("name", value)
}

// OnEvent sets event actions
func (uc UUID) OnEvent(value Event) UUID {
	return uc.set("onEvent", value)
}

// Placeholder sets placeholder text
func (uc UUID) Placeholder(value string) UUID {
	return uc.set("placeholder", value)
}

// ReadOnly sets read-only state
func (uc UUID) ReadOnly(value bool) UUID {
	return uc.set("readOnly", value)
}

// ReadOnlyOn sets read-only expression
func (uc UUID) ReadOnlyOn(value string) UUID {
	return uc.set("readOnlyOn", value)
}

// Remark sets remark
func (uc UUID) Remark(value string) UUID {
	return uc.set("remark", value)
}

// Required sets required state
func (uc UUID) Required(value bool) UUID {
	return uc.set("required", value)
}

// Row sets row value
func (uc UUID) Row(value string) UUID {
	return uc.set("row", value)
}

// SaveImmediately sets immediate save state
func (uc UUID) SaveImmediately(value bool) UUID {
	return uc.set("saveImmediately", value)
}

// Size sets form item size
func (uc UUID) Size(value string) UUID {
	return uc.set("size", value)
}

// Static sets static display state
func (uc UUID) Static(value bool) UUID {
	return uc.set("static", value)
}

// StaticClassName sets static display CSS class name
func (uc UUID) StaticClassName(value string) UUID {
	return uc.set("staticClassName", value)
}

// StaticInputClassName sets static input CSS class name
func (uc UUID) StaticInputClassName(value string) UUID {
	return uc.set("staticInputClassName", value)
}

// StaticLabelClassName sets static label CSS class name
func (uc UUID) StaticLabelClassName(value string) UUID {
	return uc.set("staticLabelClassName", value)
}

// StaticOn sets static display expression
func (uc UUID) StaticOn(value string) UUID {
	return uc.set("staticOn", value)
}

// StaticPlaceholder sets static placeholder
func (uc UUID) StaticPlaceholder(value string) UUID {
	return uc.set("staticPlaceholder", value)
}

// StaticSchema sets static schema.Schema
func (uc UUID) StaticSchema(value string) UUID {
	return uc.set("staticSchema", value)
}

// Style sets component style
func (uc UUID) Style(value any) UUID {
	return uc.set("style", value)
}

// SubmitOnChange sets submit on change state
func (uc UUID) SubmitOnChange(value bool) UUID {
	return uc.set("submitOnChange", value)
}

// TestIdBuilder sets test ID builder
func (uc UUID) TestIdBuilder(value string) UUID {
	return uc.set("testIdBuilder", value)
}

// UseMobileUI sets mobile UI usage
func (uc UUID) UseMobileUI(value bool) UUID {
	return uc.set("useMobileUI", value)
}

// ValidateApi sets remote validation API
func (uc UUID) ValidateApi(value string) UUID {
	return uc.set("validateApi", value)
}

// ValidateOnChange sets validation on change state
func (uc UUID) ValidateOnChange(value bool) UUID {
	return uc.set("validateOnChange", value)
}

// ValidationErrors sets validation error messages
func (uc UUID) ValidationErrors(value string) UUID {
	return uc.set("validationErrors", value)
}

// Validations sets validations
func (uc UUID) Validations(value string) UUID {
	return uc.set("validations", value)
}

// Value sets default value
func (uc UUID) Value(value string) UUID {
	return uc.set("value", value)
}

// Visible sets visibility state
func (uc UUID) Visible(value bool) UUID {
	return uc.set("visible", value)
}

// VisibleOn sets visibility expression
func (uc UUID) VisibleOn(value string) UUID {
	return uc.set("visibleOn", value)
}

// Width sets width in table
func (uc UUID) Width(value string) UUID {
	return uc.set("width", value)
}
