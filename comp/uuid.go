package comp

import "github.com/zrcoder/amisgo/model"

// uuid UUID functional component

type uuid model.Schema

func UUID() uuid {
	return uuid{"type": "uuid"}
}

func (uc uuid) set(key string, value any) uuid {
	uc[key] = value
	return uc
}

// AutoFill sets autoFill value
func (uc uuid) AutoFill(value string) uuid {
	return uc.set("autoFill", value)
}

// ClassName sets container CSS class name
func (uc uuid) ClassName(value string) uuid {
	return uc.set("className", value)
}

// ClearValueOnHidden removes value when hidden
func (uc uuid) ClearValueOnHidden(value bool) uuid {
	return uc.set("clearValueOnHidden", value)
}

// Desc sets description
func (uc uuid) Desc(value string) uuid {
	return uc.set("desc", value)
}

// Description sets HTML description
func (uc uuid) Description(value string) uuid {
	return uc.set("description", value)
}

// DescriptionClassName sets description CSS class name
func (uc uuid) DescriptionClassName(value string) uuid {
	return uc.set("descriptionClassName", value)
}

// Disabled sets disabled state
func (uc uuid) Disabled(value bool) uuid {
	return uc.set("disabled", value)
}

// DisabledOn sets disabled expression
func (uc uuid) DisabledOn(value string) uuid {
	return uc.set("disabledOn", value)
}

// EditorSetting sets editor configuration
func (uc uuid) EditorSetting(value string) uuid {
	return uc.set("editorSetting", value)
}

// ExtraName sets extra field name
func (uc uuid) ExtraName(value string) uuid {
	return uc.set("extraName", value)
}

// Hidden sets hidden state
func (uc uuid) Hidden(value bool) uuid {
	return uc.set("hidden", value)
}

// HiddenOn sets hidden expression
func (uc uuid) HiddenOn(value string) uuid {
	return uc.set("hiddenOn", value)
}

// Hint sets input hint
func (uc uuid) Hint(value string) uuid {
	return uc.set("hint", value)
}

// Horizontal sets horizontal layout configuration
func (uc uuid) Horizontal(value string) uuid {
	return uc.set("horizontal", value)
}

// ID sets unique component ID
func (uc uuid) ID(value string) uuid {
	return uc.set("id", value)
}

// InitAutoFill sets initial autoFill value
func (uc uuid) InitAutoFill(value string) uuid {
	return uc.set("initAutoFill", value)
}

// Inline sets inline mode
func (uc uuid) Inline(value bool) uuid {
	return uc.set("inline", value)
}

// InputClassName sets input CSS class name
func (uc uuid) InputClassName(value string) uuid {
	return uc.set("inputClassName", value)
}

// Label sets label text
func (uc uuid) Label(value string) uuid {
	return uc.set("label", value)
}

// LabelAlign sets label alignment
func (uc uuid) LabelAlign(value string) uuid {
	return uc.set("labelAlign", value)
}

// LabelClassName sets label CSS class name
func (uc uuid) LabelClassName(value string) uuid {
	return uc.set("labelClassName", value)
}

// LabelRemark sets label remark
func (uc uuid) LabelRemark(value string) uuid {
	return uc.set("labelRemark", value)
}

// LabelWidth sets label width
func (uc uuid) LabelWidth(value string) uuid {
	return uc.set("labelWidth", value)
}

// Length sets UUID length
func (uc uuid) Length(value string) uuid {
	return uc.set("length", value)
}

// Mode sets display mode
func (uc uuid) Mode(value string) uuid {
	return uc.set("mode", value)
}

// Name sets field name
func (uc uuid) Name(value string) uuid {
	return uc.set("name", value)
}

// OnEvent sets event actions
func (uc uuid) OnEvent(value any) uuid {
	return uc.set("onEvent", value)
}

// Placeholder sets placeholder text
func (uc uuid) Placeholder(value string) uuid {
	return uc.set("placeholder", value)
}

// ReadOnly sets read-only state
func (uc uuid) ReadOnly(value bool) uuid {
	return uc.set("readOnly", value)
}

// ReadOnlyOn sets read-only expression
func (uc uuid) ReadOnlyOn(value string) uuid {
	return uc.set("readOnlyOn", value)
}

// Remark sets remark
func (uc uuid) Remark(value string) uuid {
	return uc.set("remark", value)
}

// Required sets required state
func (uc uuid) Required(value bool) uuid {
	return uc.set("required", value)
}

// Row sets row value
func (uc uuid) Row(value string) uuid {
	return uc.set("row", value)
}

// SaveImmediately sets immediate save state
func (uc uuid) SaveImmediately(value bool) uuid {
	return uc.set("saveImmediately", value)
}

// Size sets form item size
func (uc uuid) Size(value string) uuid {
	return uc.set("size", value)
}

// Static sets static display state
func (uc uuid) Static(value bool) uuid {
	return uc.set("static", value)
}

// StaticClassName sets static display CSS class name
func (uc uuid) StaticClassName(value string) uuid {
	return uc.set("staticClassName", value)
}

// StaticInputClassName sets static input CSS class name
func (uc uuid) StaticInputClassName(value string) uuid {
	return uc.set("staticInputClassName", value)
}

// StaticLabelClassName sets static label CSS class name
func (uc uuid) StaticLabelClassName(value string) uuid {
	return uc.set("staticLabelClassName", value)
}

// StaticOn sets static display expression
func (uc uuid) StaticOn(value string) uuid {
	return uc.set("staticOn", value)
}

// StaticPlaceholder sets static placeholder
func (uc uuid) StaticPlaceholder(value string) uuid {
	return uc.set("staticPlaceholder", value)
}

// StaticSchema sets static schema
func (uc uuid) StaticSchema(value string) uuid {
	return uc.set("staticSchema", value)
}

// Style sets component style
func (uc uuid) Style(value any) uuid {
	return uc.set("style", value)
}

// SubmitOnChange sets submit on change state
func (uc uuid) SubmitOnChange(value bool) uuid {
	return uc.set("submitOnChange", value)
}

// TestIdBuilder sets test ID builder
func (uc uuid) TestIdBuilder(value string) uuid {
	return uc.set("testIdBuilder", value)
}

// UseMobileUI sets mobile UI usage
func (uc uuid) UseMobileUI(value bool) uuid {
	return uc.set("useMobileUI", value)
}

// ValidateApi sets remote validation API
func (uc uuid) ValidateApi(value string) uuid {
	return uc.set("validateApi", value)
}

// ValidateOnChange sets validation on change state
func (uc uuid) ValidateOnChange(value bool) uuid {
	return uc.set("validateOnChange", value)
}

// ValidationErrors sets validation error messages
func (uc uuid) ValidationErrors(value string) uuid {
	return uc.set("validationErrors", value)
}

// Validations sets validations
func (uc uuid) Validations(value string) uuid {
	return uc.set("validations", value)
}

// Value sets default value
func (uc uuid) Value(value string) uuid {
	return uc.set("value", value)
}

// Visible sets visibility state
func (uc uuid) Visible(value bool) uuid {
	return uc.set("visible", value)
}

// VisibleOn sets visibility expression
func (uc uuid) VisibleOn(value string) uuid {
	return uc.set("visibleOn", value)
}

// Width sets width in table
func (uc uuid) Width(value string) uuid {
	return uc.set("width", value)
}
