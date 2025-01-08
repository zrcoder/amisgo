package comp

import "github.com/zrcoder/amisgo/model"

// inputRange controls the range of the component

type inputRange model.Schema

// InputRange creates a new RangeControl instance
func InputRange() inputRange {
	return inputRange{}.set("type", "input-range")
}

func (rc inputRange) set(key string, value any) inputRange {
	rc[key] = value
	return rc
}

// AutoFill sets autoFill when an option is selected
func (rc inputRange) AutoFill(value string) inputRange {
	return rc.set("autoFill", value)
}

// ClassName sets the container CSS class name
func (rc inputRange) ClassName(value string) inputRange {
	return rc.set("className", value)
}

// ClearValueOnHidden removes the form item value when hidden
func (rc inputRange) ClearValueOnHidden(value bool) inputRange {
	return rc.set("clearValueOnHidden", value)
}

// Clearable sets whether the input box is clearable
func (rc inputRange) Clearable(value bool) inputRange {
	return rc.set("clearable", value)
}

// Delimiter sets the delimiter
func (rc inputRange) Delimiter(value string) inputRange {
	return rc.set("delimiter", value)
}

// Desc sets the description
func (rc inputRange) Desc(value string) inputRange {
	return rc.set("desc", value)
}

// Description sets the description, supports HTML
func (rc inputRange) Description(value string) inputRange {
	return rc.set("description", value)
}

// DescriptionClassName sets the class name for the description
func (rc inputRange) DescriptionClassName(value string) inputRange {
	return rc.set("descriptionClassName", value)
}

// Disabled sets whether the component is disabled
func (rc inputRange) Disabled(value bool) inputRange {
	return rc.set("disabled", value)
}

// DisabledOn sets the expression to disable the component
func (rc inputRange) DisabledOn(value string) inputRange {
	return rc.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (rc inputRange) EditorSetting(value string) inputRange {
	return rc.set("editorSetting", value)
}

// ExtraName sets an extra field name for range components
func (rc inputRange) ExtraName(value string) inputRange {
	return rc.set("extraName", value)
}

// Hidden sets whether the component is hidden
func (rc inputRange) Hidden(value bool) inputRange {
	return rc.set("hidden", value)
}

// HiddenOn sets the expression to hide the component
func (rc inputRange) HiddenOn(value string) inputRange {
	return rc.set("hiddenOn", value)
}

// Hint sets the input hint
func (rc inputRange) Hint(value string) inputRange {
	return rc.set("hint", value)
}

// Horizontal sets the horizontal layout configuration
func (rc inputRange) Horizontal(value string) inputRange {
	return rc.set("horizontal", value)
}

// ID sets the unique component ID
func (rc inputRange) ID(value string) inputRange {
	return rc.set("id", value)
}

// InitAutoFill sets the initial autoFill value
func (rc inputRange) InitAutoFill(value string) inputRange {
	return rc.set("initAutoFill", value)
}

// Inline sets whether the form control is inline
func (rc inputRange) Inline(value bool) inputRange {
	return rc.set("inline", value)
}

// InputClassName sets the input class name
func (rc inputRange) InputClassName(value string) inputRange {
	return rc.set("inputClassName", value)
}

// JoinValues sets whether to join values with a delimiter
func (rc inputRange) JoinValues(value bool) inputRange {
	return rc.set("joinValues", value)
}

// Label sets the label
func (rc inputRange) Label(value string) inputRange {
	return rc.set("label", value)
}

// LabelAlign sets the label alignment
func (rc inputRange) LabelAlign(value string) inputRange {
	return rc.set("labelAlign", value)
}

// LabelClassName sets the label class name
func (rc inputRange) LabelClassName(value string) inputRange {
	return rc.set("labelClassName", value)
}

// LabelRemark sets the label remark
func (rc inputRange) LabelRemark(value string) inputRange {
	return rc.set("labelRemark", value)
}

// LabelWidth sets the custom label width
func (rc inputRange) LabelWidth(value string) inputRange {
	return rc.set("labelWidth", value)
}

// Marks sets the marks
func (rc inputRange) Marks(value string) inputRange {
	return rc.set("marks", value)
}

// Max sets the maximum value
func (rc inputRange) Max(value string) inputRange {
	return rc.set("max", value)
}

// Min sets the minimum value
func (rc inputRange) Min(value string) inputRange {
	return rc.set("min", value)
}

// Mode sets the display mode
func (rc inputRange) Mode(value string) inputRange {
	return rc.set("mode", value)
}

// Multiple sets whether to use dual sliders
func (rc inputRange) Multiple(value bool) inputRange {
	return rc.set("multiple", value)
}

// Name sets the field name
func (rc inputRange) Name(value string) inputRange {
	return rc.set("name", value)
}

// OnEvent sets the event action configuration
func (rc inputRange) OnEvent(value any) inputRange {
	return rc.set("onEvent", value)
}

// Parts sets the number of segments
func (rc inputRange) Parts(value string) inputRange {
	return rc.set("parts", value)
}

// Placeholder sets the placeholder
func (rc inputRange) Placeholder(value string) inputRange {
	return rc.set("placeholder", value)
}

// ReadOnly sets whether the component is read-only
func (rc inputRange) ReadOnly(value bool) inputRange {
	return rc.set("readOnly", value)
}

// ReadOnlyOn sets the expression for read-only
func (rc inputRange) ReadOnlyOn(value string) inputRange {
	return rc.set("readOnlyOn", value)
}

// Remark sets the remark
func (rc inputRange) Remark(value string) inputRange {
	return rc.set("remark", value)
}

// Required sets whether the component is required
func (rc inputRange) Required(value bool) inputRange {
	return rc.set("required", value)
}

// Row sets the row
func (rc inputRange) Row(value string) inputRange {
	return rc.set("row", value)
}

// SaveImmediately sets whether to save immediately
func (rc inputRange) SaveImmediately(value bool) inputRange {
	return rc.set("saveImmediately", value)
}

// ShowInput sets whether to show the input box
func (rc inputRange) ShowInput(value bool) inputRange {
	return rc.set("showInput", value)
}

// ShowSteps sets whether to show steps
func (rc inputRange) ShowSteps(value bool) inputRange {
	return rc.set("showSteps", value)
}

// Size sets the size of the form item
func (rc inputRange) Size(value string) inputRange {
	return rc.set("size", value)
}

// Static sets whether the component is static
func (rc inputRange) Static(value bool) inputRange {
	return rc.set("static", value)
}

// StaticClassName sets the class name for static display
func (rc inputRange) StaticClassName(value string) inputRange {
	return rc.set("staticClassName", value)
}

// StaticInputClassName sets the class name for static input value
func (rc inputRange) StaticInputClassName(value string) inputRange {
	return rc.set("staticInputClassName", value)
}

// StaticLabelClassName sets the class name for static label
func (rc inputRange) StaticLabelClassName(value string) inputRange {
	return rc.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display
func (rc inputRange) StaticOn(value string) inputRange {
	return rc.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (rc inputRange) StaticPlaceholder(value string) inputRange {
	return rc.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema
func (rc inputRange) StaticSchema(value string) inputRange {
	return rc.set("staticSchema", value)
}

// Step sets the step value
func (rc inputRange) Step(value string) inputRange {
	return rc.set("step", value)
}

// Style sets the component style
func (rc inputRange) Style(value any) inputRange {
	return rc.set("style", value)
}

// SubmitOnChange sets whether to submit the form on change
func (rc inputRange) SubmitOnChange(value bool) inputRange {
	return rc.set("submitOnChange", value)
}

// TestIdBuilder sets the test ID builder
func (rc inputRange) TestIdBuilder(value string) inputRange {
	return rc.set("testIdBuilder", value)
}

// TooltipPlacement sets the tooltip placement
func (rc inputRange) TooltipPlacement(value string) inputRange {
	return rc.set("tooltipPlacement", value)
}

// TooltipVisible sets whether to show the tooltip
func (rc inputRange) TooltipVisible(value bool) inputRange {
	return rc.set("tooltipVisible", value)
}

// Unit sets the unit
func (rc inputRange) Unit(value string) inputRange {
	return rc.set("unit", value)
}

// UseMobileUI sets whether to use mobile UI
func (rc inputRange) UseMobileUI(value bool) inputRange {
	return rc.set("useMobileUI", value)
}

// ValidateApi sets the remote validation API
func (rc inputRange) ValidateApi(value string) inputRange {
	return rc.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on change
func (rc inputRange) ValidateOnChange(value bool) inputRange {
	return rc.set("validateOnChange", value)
}

// ValidationErrors sets the validation error messages
func (rc inputRange) ValidationErrors(value string) inputRange {
	return rc.set("validationErrors", value)
}

// Validations sets the validations
func (rc inputRange) Validations(value string) inputRange {
	return rc.set("validations", value)
}

// Value sets the slider value
func (rc inputRange) Value(value string) inputRange {
	return rc.set("value", value)
}

// Visible sets whether the component is visible
func (rc inputRange) Visible(value bool) inputRange {
	return rc.set("visible", value)
}

// VisibleOn sets the expression for visibility
func (rc inputRange) VisibleOn(value string) inputRange {
	return rc.set("visibleOn", value)
}

// Width sets the width in a table
func (rc inputRange) Width(value string) inputRange {
	return rc.set("width", value)
}
