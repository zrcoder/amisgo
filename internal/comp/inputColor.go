package comp

import "github.com/zrcoder/amisgo/schema"

// InputColor represents a color picker input.
type InputColor schema.Schema

func NewInputColor() InputColor {
	return InputColor{"type": "input-color"}
}

func (i InputColor) set(key string, value any) InputColor {
	i[key] = value
	return i
}

// AllowCustomColor allows users to input custom colors.
func (i InputColor) AllowCustomColor(value bool) InputColor {
	return i.set("allowCustomColor", value)
}

// AutoFill auto-fills other form values when an option is selected.
func (i InputColor) AutoFill(value string) InputColor {
	return i.set("autoFill", value)
}

// ClassName sets the CSS class name for the container.
func (i InputColor) ClassName(value string) InputColor {
	return i.set("className", value)
}

// ClearValueOnHidden removes the form item value when hidden.
func (i InputColor) ClearValueOnHidden(value bool) InputColor {
	return i.set("clearValueOnHidden", value)
}

// Clearable shows a clear button.
func (i InputColor) Clearable(value bool) InputColor {
	return i.set("clearable", value)
}

// CloseOnSelect closes the popup after selecting a color.
func (i InputColor) CloseOnSelect(value bool) InputColor {
	return i.set("closeOnSelect", value)
}

// Desc sets the description.
func (i InputColor) Desc(value string) InputColor {
	return i.set("desc", value)
}

// Description sets the description content, supports HTML.
func (i InputColor) Description(value string) InputColor {
	return i.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description.
func (i InputColor) DescriptionClassName(value string) InputColor {
	return i.set("descriptionClassName", value)
}

// Disabled disables the input.
func (i InputColor) Disabled(value bool) InputColor {
	return i.set("disabled", value)
}

// DisabledOn sets the condition to disable the input.
func (i InputColor) DisabledOn(value string) InputColor {
	return i.set("disabledOn", value)
}

// EditorSetting sets the editor configuration.
func (i InputColor) EditorSetting(value string) InputColor {
	return i.set("editorSetting", value)
}

// ExtraName sets an extra field name for range components.
func (i InputColor) ExtraName(value string) InputColor {
	return i.set("extraName", value)
}

// Format sets the color format (hex, rgb, rgba, hsl).
func (i InputColor) Format(value string) InputColor {
	return i.set("format", value)
}

// Hidden hides the input.
func (i InputColor) Hidden(value bool) InputColor {
	return i.set("hidden", value)
}

// HiddenOn sets the condition to hide the input.
func (i InputColor) HiddenOn(value string) InputColor {
	return i.set("hiddenOn", value)
}

// Hint sets the input hint.
func (i InputColor) Hint(value string) InputColor {
	return i.set("hint", value)
}

// Horizontal sets the horizontal layout configuration.
func (i InputColor) Horizontal(value string) InputColor {
	return i.set("horizontal", value)
}

// ID sets the unique component ID.
func (i InputColor) ID(value string) InputColor {
	return i.set("id", value)
}

// InitAutoFill sets the initial auto-fill value.
func (i InputColor) InitAutoFill(value string) InputColor {
	return i.set("initAutoFill", value)
}

// Inline sets the input to inline mode.
func (i InputColor) Inline(value bool) InputColor {
	return i.set("inline", value)
}

// InputClassName sets the CSS class name for the input.
func (i InputColor) InputClassName(value string) InputColor {
	return i.set("inputClassName", value)
}

// Label sets the label text.
func (i InputColor) Label(value string) InputColor {
	return i.set("label", value)
}

// LabelAlign sets the label alignment (right, left, top, inherit).
func (i InputColor) LabelAlign(value string) InputColor {
	return i.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label.
func (i InputColor) LabelClassName(value string) InputColor {
	return i.set("labelClassName", value)
}

// LabelRemark sets a remark icon with tooltip for the label.
func (i InputColor) LabelRemark(value string) InputColor {
	return i.set("labelRemark", value)
}

// LabelWidth sets the custom width for the label.
func (i InputColor) LabelWidth(value string) InputColor {
	return i.set("labelWidth", value)
}

// Mode sets the display mode (normal, inline, horizontal).
func (i InputColor) Mode(value string) InputColor {
	return i.set("mode", value)
}

// Name sets the field name for form submission.
func (i InputColor) Name(value string) InputColor {
	return i.set("name", value)
}

// OnEvent sets the event action configuration.
func (i InputColor) OnEvent(value Event) InputColor {
	return i.set("onEvent", value)
}

// Placeholder sets the placeholder text.
func (i InputColor) Placeholder(value string) InputColor {
	return i.set("placeholder", value)
}

// PopOverContainerSelector sets the popup container selector.
func (i InputColor) PopOverContainerSelector(value string) InputColor {
	return i.set("popOverContainerSelector", value)
}

// PresetColors sets the preset colors for selection.
func (i InputColor) PresetColors(value string) InputColor {
	return i.set("presetColors", value)
}

// ReadOnly sets the input to read-only mode.
func (i InputColor) ReadOnly(value bool) InputColor {
	return i.set("readOnly", value)
}

// ReadOnlyOn sets the condition for read-only mode.
func (i InputColor) ReadOnlyOn(value string) InputColor {
	return i.set("readOnlyOn", value)
}

// Remark sets a remark icon with tooltip.
func (i InputColor) Remark(value string) InputColor {
	return i.set("remark", value)
}

// Required sets the input as required.
func (i InputColor) Required(value bool) InputColor {
	return i.set("required", value)
}

// Row sets the row configuration.
func (i InputColor) Row(value string) InputColor {
	return i.set("row", value)
}

// SaveImmediately saves the value immediately in TableColumn.
func (i InputColor) SaveImmediately(value bool) InputColor {
	return i.set("saveImmediately", value)
}

// Size sets the input size (xs, sm, md, lg, full).
func (i InputColor) Size(value string) InputColor {
	return i.set("size", value)
}

// Static sets the input to static display mode.
func (i InputColor) Static(value bool) InputColor {
	return i.set("static", value)
}

// StaticClassName sets the CSS class name for static display.
func (i InputColor) StaticClassName(value string) InputColor {
	return i.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input value.
func (i InputColor) StaticInputClassName(value string) InputColor {
	return i.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label.
func (i InputColor) StaticLabelClassName(value string) InputColor {
	return i.set("staticLabelClassName", value)
}

// StaticOn sets the condition for static display.
func (i InputColor) StaticOn(value string) InputColor {
	return i.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display.
func (i InputColor) StaticPlaceholder(value string) InputColor {
	return i.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema.Schema.
func (i InputColor) StaticSchema(value string) InputColor {
	return i.set("staticSchema", value)
}

// Style sets the component style.
func (i InputColor) Style(value any) InputColor {
	return i.set("style", value)
}

// SubmitOnChange submits the form when the value changes.
func (i InputColor) SubmitOnChange(value bool) InputColor {
	return i.set("submitOnChange", value)
}

// TestIdBuilder sets the test ID builder.
func (i InputColor) TestIdBuilder(value string) InputColor {
	return i.set("testIdBuilder", value)
}

// UseMobileUI disables mobile UI styles.
func (i InputColor) UseMobileUI(value bool) InputColor {
	return i.set("useMobileUI", value)
}

// ValidateApi sets the remote validation API.
func (i InputColor) ValidateApi(value string) InputColor {
	return i.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on change.
func (i InputColor) ValidateOnChange(value bool) InputColor {
	return i.set("validateOnChange", value)
}

// ValidationErrors sets the validation error messages.
func (i InputColor) ValidationErrors(value string) InputColor {
	return i.set("validationErrors", value)
}

// Validations sets the validation rules.
func (i InputColor) Validations(value string) InputColor {
	return i.set("validations", value)
}

// Value sets the default value.
func (i InputColor) Value(value string) InputColor {
	return i.set("value", value)
}

// Visible sets the input visibility.
func (i InputColor) Visible(value bool) InputColor {
	return i.set("visible", value)
}

// VisibleOn sets the condition for visibility.
func (i InputColor) VisibleOn(value string) InputColor {
	return i.set("visibleOn", value)
}

// Width sets the width in a table.
func (i InputColor) Width(value string) InputColor {
	return i.set("width", value)
}
