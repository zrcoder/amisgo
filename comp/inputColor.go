package comp

import "github.com/zrcoder/amisgo/model"

// inputColor represents a color picker input.
type inputColor model.Schema

func InputColor() inputColor {
	return make(inputColor).set("type", "input-color")
}

func (i inputColor) set(key string, value any) inputColor {
	i[key] = value
	return i
}

// AllowCustomColor allows users to input custom colors.
func (i inputColor) AllowCustomColor(value bool) inputColor {
	return i.set("allowCustomColor", value)
}

// AutoFill auto-fills other form values when an option is selected.
func (i inputColor) AutoFill(value string) inputColor {
	return i.set("autoFill", value)
}

// ClassName sets the CSS class name for the container.
func (i inputColor) ClassName(value string) inputColor {
	return i.set("className", value)
}

// ClearValueOnHidden removes the form item value when hidden.
func (i inputColor) ClearValueOnHidden(value bool) inputColor {
	return i.set("clearValueOnHidden", value)
}

// Clearable shows a clear button.
func (i inputColor) Clearable(value bool) inputColor {
	return i.set("clearable", value)
}

// CloseOnSelect closes the popup after selecting a color.
func (i inputColor) CloseOnSelect(value bool) inputColor {
	return i.set("closeOnSelect", value)
}

// Desc sets the description.
func (i inputColor) Desc(value string) inputColor {
	return i.set("desc", value)
}

// Description sets the description content, supports HTML.
func (i inputColor) Description(value string) inputColor {
	return i.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description.
func (i inputColor) DescriptionClassName(value string) inputColor {
	return i.set("descriptionClassName", value)
}

// Disabled disables the input.
func (i inputColor) Disabled(value bool) inputColor {
	return i.set("disabled", value)
}

// DisabledOn sets the condition to disable the input.
func (i inputColor) DisabledOn(value string) inputColor {
	return i.set("disabledOn", value)
}

// EditorSetting sets the editor configuration.
func (i inputColor) EditorSetting(value string) inputColor {
	return i.set("editorSetting", value)
}

// ExtraName sets an extra field name for range components.
func (i inputColor) ExtraName(value string) inputColor {
	return i.set("extraName", value)
}

// Format sets the color format (hex, rgb, rgba, hsl).
func (i inputColor) Format(value string) inputColor {
	return i.set("format", value)
}

// Hidden hides the input.
func (i inputColor) Hidden(value bool) inputColor {
	return i.set("hidden", value)
}

// HiddenOn sets the condition to hide the input.
func (i inputColor) HiddenOn(value string) inputColor {
	return i.set("hiddenOn", value)
}

// Hint sets the input hint.
func (i inputColor) Hint(value string) inputColor {
	return i.set("hint", value)
}

// Horizontal sets the horizontal layout configuration.
func (i inputColor) Horizontal(value string) inputColor {
	return i.set("horizontal", value)
}

// ID sets the unique component ID.
func (i inputColor) ID(value string) inputColor {
	return i.set("id", value)
}

// InitAutoFill sets the initial auto-fill value.
func (i inputColor) InitAutoFill(value string) inputColor {
	return i.set("initAutoFill", value)
}

// Inline sets the input to inline mode.
func (i inputColor) Inline(value bool) inputColor {
	return i.set("inline", value)
}

// InputClassName sets the CSS class name for the input.
func (i inputColor) InputClassName(value string) inputColor {
	return i.set("inputClassName", value)
}

// Label sets the label text.
func (i inputColor) Label(value string) inputColor {
	return i.set("label", value)
}

// LabelAlign sets the label alignment (right, left, top, inherit).
func (i inputColor) LabelAlign(value string) inputColor {
	return i.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label.
func (i inputColor) LabelClassName(value string) inputColor {
	return i.set("labelClassName", value)
}

// LabelRemark sets a remark icon with tooltip for the label.
func (i inputColor) LabelRemark(value string) inputColor {
	return i.set("labelRemark", value)
}

// LabelWidth sets the custom width for the label.
func (i inputColor) LabelWidth(value string) inputColor {
	return i.set("labelWidth", value)
}

// Mode sets the display mode (normal, inline, horizontal).
func (i inputColor) Mode(value string) inputColor {
	return i.set("mode", value)
}

// Name sets the field name for form submission.
func (i inputColor) Name(value string) inputColor {
	return i.set("name", value)
}

// OnEvent sets the event action configuration.
func (i inputColor) OnEvent(value any) inputColor {
	return i.set("onEvent", value)
}

// Placeholder sets the placeholder text.
func (i inputColor) Placeholder(value string) inputColor {
	return i.set("placeholder", value)
}

// PopOverContainerSelector sets the popup container selector.
func (i inputColor) PopOverContainerSelector(value string) inputColor {
	return i.set("popOverContainerSelector", value)
}

// PresetColors sets the preset colors for selection.
func (i inputColor) PresetColors(value string) inputColor {
	return i.set("presetColors", value)
}

// ReadOnly sets the input to read-only mode.
func (i inputColor) ReadOnly(value bool) inputColor {
	return i.set("readOnly", value)
}

// ReadOnlyOn sets the condition for read-only mode.
func (i inputColor) ReadOnlyOn(value string) inputColor {
	return i.set("readOnlyOn", value)
}

// Remark sets a remark icon with tooltip.
func (i inputColor) Remark(value string) inputColor {
	return i.set("remark", value)
}

// Required sets the input as required.
func (i inputColor) Required(value bool) inputColor {
	return i.set("required", value)
}

// Row sets the row configuration.
func (i inputColor) Row(value string) inputColor {
	return i.set("row", value)
}

// SaveImmediately saves the value immediately in TableColumn.
func (i inputColor) SaveImmediately(value bool) inputColor {
	return i.set("saveImmediately", value)
}

// Size sets the input size (xs, sm, md, lg, full).
func (i inputColor) Size(value string) inputColor {
	return i.set("size", value)
}

// Static sets the input to static display mode.
func (i inputColor) Static(value bool) inputColor {
	return i.set("static", value)
}

// StaticClassName sets the CSS class name for static display.
func (i inputColor) StaticClassName(value string) inputColor {
	return i.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input value.
func (i inputColor) StaticInputClassName(value string) inputColor {
	return i.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label.
func (i inputColor) StaticLabelClassName(value string) inputColor {
	return i.set("staticLabelClassName", value)
}

// StaticOn sets the condition for static display.
func (i inputColor) StaticOn(value string) inputColor {
	return i.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display.
func (i inputColor) StaticPlaceholder(value string) inputColor {
	return i.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema.
func (i inputColor) StaticSchema(value string) inputColor {
	return i.set("staticSchema", value)
}

// Style sets the component style.
func (i inputColor) Style(value any) inputColor {
	return i.set("style", value)
}

// SubmitOnChange submits the form when the value changes.
func (i inputColor) SubmitOnChange(value bool) inputColor {
	return i.set("submitOnChange", value)
}

// TestIdBuilder sets the test ID builder.
func (i inputColor) TestIdBuilder(value string) inputColor {
	return i.set("testIdBuilder", value)
}

// UseMobileUI disables mobile UI styles.
func (i inputColor) UseMobileUI(value bool) inputColor {
	return i.set("useMobileUI", value)
}

// ValidateApi sets the remote validation API.
func (i inputColor) ValidateApi(value string) inputColor {
	return i.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on change.
func (i inputColor) ValidateOnChange(value bool) inputColor {
	return i.set("validateOnChange", value)
}

// ValidationErrors sets the validation error messages.
func (i inputColor) ValidationErrors(value string) inputColor {
	return i.set("validationErrors", value)
}

// Validations sets the validation rules.
func (i inputColor) Validations(value string) inputColor {
	return i.set("validations", value)
}

// Value sets the default value.
func (i inputColor) Value(value string) inputColor {
	return i.set("value", value)
}

// Visible sets the input visibility.
func (i inputColor) Visible(value bool) inputColor {
	return i.set("visible", value)
}

// VisibleOn sets the condition for visibility.
func (i inputColor) VisibleOn(value string) inputColor {
	return i.set("visibleOn", value)
}

// Width sets the width in a table.
func (i inputColor) Width(value string) inputColor {
	return i.set("width", value)
}
