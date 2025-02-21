package comp

import "github.com/zrcoder/amisgo/schema"

// InputCity represents a city selection component.
type InputCity schema.Schema

func NewInputCity() InputCity {
	return InputCity{"type": "input-city"}
}

func (i InputCity) set(key string, value any) InputCity {
	i[key] = value
	return i
}

// AllowCity allows city selection.
func (i InputCity) AllowCity(value bool) InputCity {
	return i.set("allowCity", value)
}

// AllowDistrict allows district selection.
func (i InputCity) AllowDistrict(value bool) InputCity {
	return i.set("allowDistrict", value)
}

// AllowStreet allows street selection.
func (i InputCity) AllowStreet(value bool) InputCity {
	return i.set("allowStreet", value)
}

// AutoFill auto-fills form fields when an option is selected.
func (i InputCity) AutoFill(value string) InputCity {
	return i.set("autoFill", value)
}

// ClassName sets the CSS class name.
func (i InputCity) ClassName(value string) InputCity {
	return i.set("className", value)
}

// ClearValueOnHidden clears the value when the form item is hidden.
func (i InputCity) ClearValueOnHidden(value bool) InputCity {
	return i.set("clearValueOnHidden", value)
}

// Delimiter sets the delimiter for concatenation.
func (i InputCity) Delimiter(value string) InputCity {
	return i.set("delimiter", value)
}

// Desc sets the description.
func (i InputCity) Desc(value string) InputCity {
	return i.set("desc", value)
}

// Description sets the description content, supports HTML.
func (i InputCity) Description(value string) InputCity {
	return i.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description.
func (i InputCity) DescriptionClassName(value string) InputCity {
	return i.set("descriptionClassName", value)
}

// Disabled disables the component.
func (i InputCity) Disabled(value bool) InputCity {
	return i.set("disabled", value)
}

// DisabledOn sets the expression to disable the component.
func (i InputCity) DisabledOn(value string) InputCity {
	return i.set("disabledOn", value)
}

// EditorSetting sets the editor configuration.
func (i InputCity) EditorSetting(value string) InputCity {
	return i.set("editorSetting", value)
}

// ExtraName sets an extra field name for range components.
func (i InputCity) ExtraName(value string) InputCity {
	return i.set("extraName", value)
}

// ExtractValue stores only the city code.
func (i InputCity) ExtractValue(value bool) InputCity {
	return i.set("extractValue", value)
}

// Hidden hides the component.
func (i InputCity) Hidden(value bool) InputCity {
	return i.set("hidden", value)
}

// HiddenOn sets the expression to hide the component.
func (i InputCity) HiddenOn(value string) InputCity {
	return i.set("hiddenOn", value)
}

// Hint sets the input hint.
func (i InputCity) Hint(value string) InputCity {
	return i.set("hint", value)
}

// Horizontal sets the horizontal layout configuration.
func (i InputCity) Horizontal(value string) InputCity {
	return i.set("horizontal", value)
}

// ID sets the unique component ID.
func (i InputCity) ID(value string) InputCity {
	return i.set("id", value)
}

// InitAutoFill sets the initial auto-fill value.
func (i InputCity) InitAutoFill(value string) InputCity {
	return i.set("initAutoFill", value)
}

// Inline sets the component to inline mode.
func (i InputCity) Inline(value bool) InputCity {
	return i.set("inline", value)
}

// InputClassName sets the CSS class name for the input.
func (i InputCity) InputClassName(value string) InputCity {
	return i.set("inputClassName", value)
}

// ItemClassName sets the CSS class name for the dropdown.
func (i InputCity) ItemClassName(value string) InputCity {
	return i.set("itemClassName", value)
}

// JoinValues concatenates values into a string.
func (i InputCity) JoinValues(value bool) InputCity {
	return i.set("joinValues", value)
}

// Label sets the label text.
func (i InputCity) Label(value string) InputCity {
	return i.set("label", value)
}

// LabelAlign sets the label alignment.
func (i InputCity) LabelAlign(value string) InputCity {
	return i.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label.
func (i InputCity) LabelClassName(value string) InputCity {
	return i.set("labelClassName", value)
}

// LabelRemark sets the label remark.
func (i InputCity) LabelRemark(value string) InputCity {
	return i.set("labelRemark", value)
}

// LabelWidth sets the custom label width.
func (i InputCity) LabelWidth(value string) InputCity {
	return i.set("labelWidth", value)
}

// LoadingConfig sets the loading configuration.
func (i InputCity) LoadingConfig(value string) InputCity {
	return i.set("loadingConfig", value)
}

// Mode sets the display mode.
func (i InputCity) Mode(value string) InputCity {
	return i.set("mode", value)
}

// Name sets the field name for form submission.
func (i InputCity) Name(value string) InputCity {
	return i.set("name", value)
}

// OnEvent sets the event action configuration.
func (i InputCity) OnEvent(value Event) InputCity {
	return i.set("onEvent", value)
}

// Placeholder sets the placeholder text.
func (i InputCity) Placeholder(value string) InputCity {
	return i.set("placeholder", value)
}

// ReadOnly sets the component to read-only.
func (i InputCity) ReadOnly(value bool) InputCity {
	return i.set("readOnly", value)
}

// ReadOnlyOn sets the expression for read-only mode.
func (i InputCity) ReadOnlyOn(value string) InputCity {
	return i.set("readOnlyOn", value)
}

// Remark sets the remark text.
func (i InputCity) Remark(value string) InputCity {
	return i.set("remark", value)
}

// Required sets the component as required.
func (i InputCity) Required(value bool) InputCity {
	return i.set("required", value)
}

// Row sets the row configuration.
func (i InputCity) Row(value string) InputCity {
	return i.set("row", value)
}

// SaveImmediately saves the value immediately in TableColumn.
func (i InputCity) SaveImmediately(value bool) InputCity {
	return i.set("saveImmediately", value)
}

// Searchable enables the search box.
func (i InputCity) Searchable(value bool) InputCity {
	return i.set("searchable", value)
}

// Size sets the component size.
func (i InputCity) Size(value string) InputCity {
	return i.set("size", value)
}

// Static sets the component to static display.
func (i InputCity) Static(value bool) InputCity {
	return i.set("static", value)
}

// StaticClassName sets the CSS class name for static display.
func (i InputCity) StaticClassName(value string) InputCity {
	return i.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display.
func (i InputCity) StaticInputClassName(value string) InputCity {
	return i.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display.
func (i InputCity) StaticLabelClassName(value string) InputCity {
	return i.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display.
func (i InputCity) StaticOn(value string) InputCity {
	return i.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display.
func (i InputCity) StaticPlaceholder(value string) InputCity {
	return i.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema.Schema.
func (i InputCity) StaticSchema(value string) InputCity {
	return i.set("staticSchema", value)
}

// Style sets the component style.
func (i InputCity) Style(value any) InputCity {
	return i.set("style", value)
}

// SubmitOnChange submits the form on value change.
func (i InputCity) SubmitOnChange(value bool) InputCity {
	return i.set("submitOnChange", value)
}

// TestIdBuilder sets the test ID builder.
func (i InputCity) TestIdBuilder(value string) InputCity {
	return i.set("testIdBuilder", value)
}

// UseMobile enables mobile adaptation.
func (i InputCity) UseMobile(value bool) InputCity {
	return i.set("useMobile", value)
}

// ValidateApi sets the remote validation API.
func (i InputCity) ValidateApi(value string) InputCity {
	return i.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on change.
func (i InputCity) ValidateOnChange(value bool) InputCity {
	return i.set("validateOnChange", value)
}

// ValidationErrors sets the validation error messages.
func (i InputCity) ValidationErrors(value string) InputCity {
	return i.set("validationErrors", value)
}

// Validations sets the validation rules.
func (i InputCity) Validations(value string) InputCity {
	return i.set("validations", value)
}

// Value sets the default value.
func (i InputCity) Value(value any) InputCity {
	return i.set("value", value)
}

// Visible sets the component visibility.
func (i InputCity) Visible(value bool) InputCity {
	return i.set("visible", value)
}

// VisibleOn sets the expression for visibility.
func (i InputCity) VisibleOn(value string) InputCity {
	return i.set("visibleOn", value)
}

// Width sets the width in Table.
func (i InputCity) Width(value string) InputCity {
	return i.set("width", value)
}
