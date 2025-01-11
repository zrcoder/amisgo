package comp

import "github.com/zrcoder/amisgo/model"

// inputCity represents a city selection component.
type inputCity model.Schema

func InputCity() inputCity {
	return inputCity{"type": "input-city"}
}

func (i inputCity) set(key string, value any) inputCity {
	i[key] = value
	return i
}

// AllowCity allows city selection.
func (i inputCity) AllowCity(value bool) inputCity {
	return i.set("allowCity", value)
}

// AllowDistrict allows district selection.
func (i inputCity) AllowDistrict(value bool) inputCity {
	return i.set("allowDistrict", value)
}

// AllowStreet allows street selection.
func (i inputCity) AllowStreet(value bool) inputCity {
	return i.set("allowStreet", value)
}

// AutoFill auto-fills form fields when an option is selected.
func (i inputCity) AutoFill(value string) inputCity {
	return i.set("autoFill", value)
}

// ClassName sets the CSS class name.
func (i inputCity) ClassName(value string) inputCity {
	return i.set("className", value)
}

// ClearValueOnHidden clears the value when the form item is hidden.
func (i inputCity) ClearValueOnHidden(value bool) inputCity {
	return i.set("clearValueOnHidden", value)
}

// Delimiter sets the delimiter for concatenation.
func (i inputCity) Delimiter(value string) inputCity {
	return i.set("delimiter", value)
}

// Desc sets the description.
func (i inputCity) Desc(value string) inputCity {
	return i.set("desc", value)
}

// Description sets the description content, supports HTML.
func (i inputCity) Description(value string) inputCity {
	return i.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description.
func (i inputCity) DescriptionClassName(value string) inputCity {
	return i.set("descriptionClassName", value)
}

// Disabled disables the component.
func (i inputCity) Disabled(value bool) inputCity {
	return i.set("disabled", value)
}

// DisabledOn sets the expression to disable the component.
func (i inputCity) DisabledOn(value string) inputCity {
	return i.set("disabledOn", value)
}

// EditorSetting sets the editor configuration.
func (i inputCity) EditorSetting(value string) inputCity {
	return i.set("editorSetting", value)
}

// ExtraName sets an extra field name for range components.
func (i inputCity) ExtraName(value string) inputCity {
	return i.set("extraName", value)
}

// ExtractValue stores only the city code.
func (i inputCity) ExtractValue(value bool) inputCity {
	return i.set("extractValue", value)
}

// Hidden hides the component.
func (i inputCity) Hidden(value bool) inputCity {
	return i.set("hidden", value)
}

// HiddenOn sets the expression to hide the component.
func (i inputCity) HiddenOn(value string) inputCity {
	return i.set("hiddenOn", value)
}

// Hint sets the input hint.
func (i inputCity) Hint(value string) inputCity {
	return i.set("hint", value)
}

// Horizontal sets the horizontal layout configuration.
func (i inputCity) Horizontal(value string) inputCity {
	return i.set("horizontal", value)
}

// ID sets the unique component ID.
func (i inputCity) ID(value string) inputCity {
	return i.set("id", value)
}

// InitAutoFill sets the initial auto-fill value.
func (i inputCity) InitAutoFill(value string) inputCity {
	return i.set("initAutoFill", value)
}

// Inline sets the component to inline mode.
func (i inputCity) Inline(value bool) inputCity {
	return i.set("inline", value)
}

// InputClassName sets the CSS class name for the input.
func (i inputCity) InputClassName(value string) inputCity {
	return i.set("inputClassName", value)
}

// ItemClassName sets the CSS class name for the dropdown.
func (i inputCity) ItemClassName(value string) inputCity {
	return i.set("itemClassName", value)
}

// JoinValues concatenates values into a string.
func (i inputCity) JoinValues(value bool) inputCity {
	return i.set("joinValues", value)
}

// Label sets the label text.
func (i inputCity) Label(value string) inputCity {
	return i.set("label", value)
}

// LabelAlign sets the label alignment.
func (i inputCity) LabelAlign(value string) inputCity {
	return i.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label.
func (i inputCity) LabelClassName(value string) inputCity {
	return i.set("labelClassName", value)
}

// LabelRemark sets the label remark.
func (i inputCity) LabelRemark(value string) inputCity {
	return i.set("labelRemark", value)
}

// LabelWidth sets the custom label width.
func (i inputCity) LabelWidth(value string) inputCity {
	return i.set("labelWidth", value)
}

// LoadingConfig sets the loading configuration.
func (i inputCity) LoadingConfig(value string) inputCity {
	return i.set("loadingConfig", value)
}

// Mode sets the display mode.
func (i inputCity) Mode(value string) inputCity {
	return i.set("mode", value)
}

// Name sets the field name for form submission.
func (i inputCity) Name(value string) inputCity {
	return i.set("name", value)
}

// OnEvent sets the event action configuration.
func (i inputCity) OnEvent(value any) inputCity {
	return i.set("onEvent", value)
}

// Placeholder sets the placeholder text.
func (i inputCity) Placeholder(value string) inputCity {
	return i.set("placeholder", value)
}

// ReadOnly sets the component to read-only.
func (i inputCity) ReadOnly(value bool) inputCity {
	return i.set("readOnly", value)
}

// ReadOnlyOn sets the expression for read-only mode.
func (i inputCity) ReadOnlyOn(value string) inputCity {
	return i.set("readOnlyOn", value)
}

// Remark sets the remark text.
func (i inputCity) Remark(value string) inputCity {
	return i.set("remark", value)
}

// Required sets the component as required.
func (i inputCity) Required(value bool) inputCity {
	return i.set("required", value)
}

// Row sets the row configuration.
func (i inputCity) Row(value string) inputCity {
	return i.set("row", value)
}

// SaveImmediately saves the value immediately in TableColumn.
func (i inputCity) SaveImmediately(value bool) inputCity {
	return i.set("saveImmediately", value)
}

// Searchable enables the search box.
func (i inputCity) Searchable(value bool) inputCity {
	return i.set("searchable", value)
}

// Size sets the component size.
func (i inputCity) Size(value string) inputCity {
	return i.set("size", value)
}

// Static sets the component to static display.
func (i inputCity) Static(value bool) inputCity {
	return i.set("static", value)
}

// StaticClassName sets the CSS class name for static display.
func (i inputCity) StaticClassName(value string) inputCity {
	return i.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display.
func (i inputCity) StaticInputClassName(value string) inputCity {
	return i.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display.
func (i inputCity) StaticLabelClassName(value string) inputCity {
	return i.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display.
func (i inputCity) StaticOn(value string) inputCity {
	return i.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display.
func (i inputCity) StaticPlaceholder(value string) inputCity {
	return i.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema.
func (i inputCity) StaticSchema(value string) inputCity {
	return i.set("staticSchema", value)
}

// Style sets the component style.
func (i inputCity) Style(value any) inputCity {
	return i.set("style", value)
}

// SubmitOnChange submits the form on value change.
func (i inputCity) SubmitOnChange(value bool) inputCity {
	return i.set("submitOnChange", value)
}

// TestIdBuilder sets the test ID builder.
func (i inputCity) TestIdBuilder(value string) inputCity {
	return i.set("testIdBuilder", value)
}

// UseMobile enables mobile adaptation.
func (i inputCity) UseMobile(value bool) inputCity {
	return i.set("useMobile", value)
}

// ValidateApi sets the remote validation API.
func (i inputCity) ValidateApi(value string) inputCity {
	return i.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on change.
func (i inputCity) ValidateOnChange(value bool) inputCity {
	return i.set("validateOnChange", value)
}

// ValidationErrors sets the validation error messages.
func (i inputCity) ValidationErrors(value string) inputCity {
	return i.set("validationErrors", value)
}

// Validations sets the validation rules.
func (i inputCity) Validations(value string) inputCity {
	return i.set("validations", value)
}

// Value sets the default value.
func (i inputCity) Value(value any) inputCity {
	return i.set("value", value)
}

// Visible sets the component visibility.
func (i inputCity) Visible(value bool) inputCity {
	return i.set("visible", value)
}

// VisibleOn sets the expression for visibility.
func (i inputCity) VisibleOn(value string) inputCity {
	return i.set("visibleOn", value)
}

// Width sets the width in Table.
func (i inputCity) Width(value string) inputCity {
	return i.set("width", value)
}
