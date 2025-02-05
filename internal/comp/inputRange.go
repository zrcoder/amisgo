package comp

import "github.com/zrcoder/amisgo/model"

// InputRange controls the range of the component
type InputRange model.Schema

func NewInputRange() InputRange {
	return InputRange{"type": "input-range"}
}

func (rc InputRange) set(key string, value any) InputRange {
	rc[key] = value
	return rc
}

// AutoFill sets autoFill when an option is selected
func (rc InputRange) AutoFill(value string) InputRange {
	return rc.set("autoFill", value)
}

// ClassName sets the container CSS class name
func (rc InputRange) ClassName(value string) InputRange {
	return rc.set("className", value)
}

// ClearValueOnHidden removes the form item value when hidden
func (rc InputRange) ClearValueOnHidden(value bool) InputRange {
	return rc.set("clearValueOnHidden", value)
}

// Clearable sets whether the input box is clearable
func (rc InputRange) Clearable(value bool) InputRange {
	return rc.set("clearable", value)
}

// Delimiter sets the delimiter
func (rc InputRange) Delimiter(value string) InputRange {
	return rc.set("delimiter", value)
}

// Desc sets the description
func (rc InputRange) Desc(value string) InputRange {
	return rc.set("desc", value)
}

// Description sets the description, supports HTML
func (rc InputRange) Description(value string) InputRange {
	return rc.set("description", value)
}

// DescriptionClassName sets the class name for the description
func (rc InputRange) DescriptionClassName(value string) InputRange {
	return rc.set("descriptionClassName", value)
}

// Disabled sets whether the component is disabled
func (rc InputRange) Disabled(value bool) InputRange {
	return rc.set("disabled", value)
}

// DisabledOn sets the expression to disable the component
func (rc InputRange) DisabledOn(value string) InputRange {
	return rc.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (rc InputRange) EditorSetting(value string) InputRange {
	return rc.set("editorSetting", value)
}

// ExtraName sets an extra field name for range components
func (rc InputRange) ExtraName(value string) InputRange {
	return rc.set("extraName", value)
}

// Hidden sets whether the component is hidden
func (rc InputRange) Hidden(value bool) InputRange {
	return rc.set("hidden", value)
}

// HiddenOn sets the expression to hide the component
func (rc InputRange) HiddenOn(value string) InputRange {
	return rc.set("hiddenOn", value)
}

// Hint sets the input hint
func (rc InputRange) Hint(value string) InputRange {
	return rc.set("hint", value)
}

// Horizontal sets the horizontal layout configuration
func (rc InputRange) Horizontal(value string) InputRange {
	return rc.set("horizontal", value)
}

// ID sets the unique component ID
func (rc InputRange) ID(value string) InputRange {
	return rc.set("id", value)
}

// InitAutoFill sets the initial autoFill value
func (rc InputRange) InitAutoFill(value string) InputRange {
	return rc.set("initAutoFill", value)
}

// Inline sets whether the form control is inline
func (rc InputRange) Inline(value bool) InputRange {
	return rc.set("inline", value)
}

// InputClassName sets the input class name
func (rc InputRange) InputClassName(value string) InputRange {
	return rc.set("inputClassName", value)
}

// JoinValues sets whether to join values with a delimiter
func (rc InputRange) JoinValues(value bool) InputRange {
	return rc.set("joinValues", value)
}

// Label sets the label
func (rc InputRange) Label(value string) InputRange {
	return rc.set("label", value)
}

// LabelAlign sets the label alignment
func (rc InputRange) LabelAlign(value string) InputRange {
	return rc.set("labelAlign", value)
}

// LabelClassName sets the label class name
func (rc InputRange) LabelClassName(value string) InputRange {
	return rc.set("labelClassName", value)
}

// LabelRemark sets the label remark
func (rc InputRange) LabelRemark(value string) InputRange {
	return rc.set("labelRemark", value)
}

// LabelWidth sets the custom label width
func (rc InputRange) LabelWidth(value string) InputRange {
	return rc.set("labelWidth", value)
}

// Marks sets the marks
func (rc InputRange) Marks(value string) InputRange {
	return rc.set("marks", value)
}

// Max sets the maximum value, number | string
func (rc InputRange) Max(value any) InputRange {
	return rc.set("max", value)
}

// Min sets the minimum value, number | string
func (rc InputRange) Min(value any) InputRange {
	return rc.set("min", value)
}

// Mode sets the display mode
func (rc InputRange) Mode(value string) InputRange {
	return rc.set("mode", value)
}

// Multiple sets whether to use dual sliders
func (rc InputRange) Multiple(value bool) InputRange {
	return rc.set("multiple", value)
}

// Name sets the field name
func (rc InputRange) Name(value string) InputRange {
	return rc.set("name", value)
}

// OnEvent sets the event action configuration
func (rc InputRange) OnEvent(value any) InputRange {
	return rc.set("onEvent", value)
}

// Parts sets the number of segments
func (rc InputRange) Parts(value string) InputRange {
	return rc.set("parts", value)
}

// Placeholder sets the placeholder
func (rc InputRange) Placeholder(value string) InputRange {
	return rc.set("placeholder", value)
}

// ReadOnly sets whether the component is read-only
func (rc InputRange) ReadOnly(value bool) InputRange {
	return rc.set("readOnly", value)
}

// ReadOnlyOn sets the expression for read-only
func (rc InputRange) ReadOnlyOn(value string) InputRange {
	return rc.set("readOnlyOn", value)
}

// Remark sets the remark
func (rc InputRange) Remark(value string) InputRange {
	return rc.set("remark", value)
}

// Required sets whether the component is required
func (rc InputRange) Required(value bool) InputRange {
	return rc.set("required", value)
}

// Row sets the row
func (rc InputRange) Row(value string) InputRange {
	return rc.set("row", value)
}

// SaveImmediately sets whether to save immediately
func (rc InputRange) SaveImmediately(value bool) InputRange {
	return rc.set("saveImmediately", value)
}

// ShowInput sets whether to show the input box
func (rc InputRange) ShowInput(value bool) InputRange {
	return rc.set("showInput", value)
}

// ShowSteps sets whether to show steps
func (rc InputRange) ShowSteps(value bool) InputRange {
	return rc.set("showSteps", value)
}

// Size sets the size of the form item
func (rc InputRange) Size(value string) InputRange {
	return rc.set("size", value)
}

// Static sets whether the component is static
func (rc InputRange) Static(value bool) InputRange {
	return rc.set("static", value)
}

// StaticClassName sets the class name for static display
func (rc InputRange) StaticClassName(value string) InputRange {
	return rc.set("staticClassName", value)
}

// StaticInputClassName sets the class name for static input value
func (rc InputRange) StaticInputClassName(value string) InputRange {
	return rc.set("staticInputClassName", value)
}

// StaticLabelClassName sets the class name for static label
func (rc InputRange) StaticLabelClassName(value string) InputRange {
	return rc.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display
func (rc InputRange) StaticOn(value string) InputRange {
	return rc.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (rc InputRange) StaticPlaceholder(value string) InputRange {
	return rc.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema
func (rc InputRange) StaticSchema(value string) InputRange {
	return rc.set("staticSchema", value)
}

// Step sets the step value number | string
func (rc InputRange) Step(value any) InputRange {
	return rc.set("step", value)
}

// Style sets the component style
func (rc InputRange) Style(value any) InputRange {
	return rc.set("style", value)
}

// SubmitOnChange sets whether to submit the form on change
func (rc InputRange) SubmitOnChange(value bool) InputRange {
	return rc.set("submitOnChange", value)
}

// TestIdBuilder sets the test ID builder
func (rc InputRange) TestIdBuilder(value string) InputRange {
	return rc.set("testIdBuilder", value)
}

// TooltipPlacement sets the tooltip placement
func (rc InputRange) TooltipPlacement(value string) InputRange {
	return rc.set("tooltipPlacement", value)
}

// TooltipVisible sets whether to show the tooltip
func (rc InputRange) TooltipVisible(value bool) InputRange {
	return rc.set("tooltipVisible", value)
}

// Unit sets the unit
func (rc InputRange) Unit(value string) InputRange {
	return rc.set("unit", value)
}

// UseMobileUI sets whether to use mobile UI
func (rc InputRange) UseMobileUI(value bool) InputRange {
	return rc.set("useMobileUI", value)
}

// ValidateApi sets the remote validation API
func (rc InputRange) ValidateApi(value string) InputRange {
	return rc.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on change
func (rc InputRange) ValidateOnChange(value bool) InputRange {
	return rc.set("validateOnChange", value)
}

// ValidationErrors sets the validation error messages
func (rc InputRange) ValidationErrors(value string) InputRange {
	return rc.set("validationErrors", value)
}

// Validations sets the validations
func (rc InputRange) Validations(value string) InputRange {
	return rc.set("validations", value)
}

// Value sets the slider value, number or string or {min: number, max: number} or [number, number]
func (rc InputRange) Value(value any) InputRange {
	return rc.set("value", value)
}

// Visible sets whether the component is visible
func (rc InputRange) Visible(value bool) InputRange {
	return rc.set("visible", value)
}

// VisibleOn sets the expression for visibility
func (rc InputRange) VisibleOn(value string) InputRange {
	return rc.set("visibleOn", value)
}

// Width sets the width in a table
func (rc InputRange) Width(value string) InputRange {
	return rc.set("width", value)
}
