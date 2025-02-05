package comp

import "github.com/zrcoder/amisgo/model"

// LocationPicker component
// Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/form/location-picker
type LocationPicker model.Schema

func NewLocationPicker() LocationPicker {
	return LocationPicker{"type": "location-picker"}
}

func (lc LocationPicker) set(key string, value any) LocationPicker {
	lc[key] = value
	return lc
}

// Ak sets the ak information for some maps
func (lc LocationPicker) Ak(value string) LocationPicker {
	lc.set("ak", value)
	return lc
}

// AutoFill sets autoFill to sync other values to the form when an option is selected
func (lc LocationPicker) AutoFill(value string) LocationPicker {
	lc.set("autoFill", value)
	return lc
}

// AutoSelectCurrentLoc sets whether to auto-select the current location
func (lc LocationPicker) AutoSelectCurrentLoc(value bool) LocationPicker {
	lc.set("autoSelectCurrentLoc", value)
	return lc
}

// ClassName sets the container CSS class name
func (lc LocationPicker) ClassName(value string) LocationPicker {
	lc.set("className", value)
	return lc
}

// ClearValueOnHidden sets whether to remove the form item value when hidden
func (lc LocationPicker) ClearValueOnHidden(value bool) LocationPicker {
	lc.set("clearValueOnHidden", value)
	return lc
}

// Desc sets the description content
func (lc LocationPicker) Desc(value string) LocationPicker {
	lc.set("desc", value)
	return lc
}

// Description sets the description content, supports HTML
func (lc LocationPicker) Description(value string) LocationPicker {
	lc.set("description", value)
	return lc
}

// DescriptionClassName sets the class name for the description
func (lc LocationPicker) DescriptionClassName(value string) LocationPicker {
	lc.set("descriptionClassName", value)
	return lc
}

// Disabled sets whether the component is disabled
func (lc LocationPicker) Disabled(value bool) LocationPicker {
	lc.set("disabled", value)
	return lc
}

// DisabledOn sets the expression to disable the component
func (lc LocationPicker) DisabledOn(value string) LocationPicker {
	lc.set("disabledOn", value)
	return lc
}

// EditorSetting sets the editor configuration, ignored at runtime
func (lc LocationPicker) EditorSetting(value string) LocationPicker {
	lc.set("editorSetting", value)
	return lc
}

// ExtraName sets an extra field name for range components
func (lc LocationPicker) ExtraName(value string) LocationPicker {
	lc.set("extraName", value)
	return lc
}

// GetLocationPlaceholder sets the placeholder when in read-only mode
func (lc LocationPicker) GetLocationPlaceholder(value string) LocationPicker {
	lc.set("getLocationPlaceholder", value)
	return lc
}

// Hidden sets whether the component is hidden
func (lc LocationPicker) Hidden(value bool) LocationPicker {
	lc.set("hidden", value)
	return lc
}

// HiddenOn sets the expression to hide the component
func (lc LocationPicker) HiddenOn(value string) LocationPicker {
	lc.set("hiddenOn", value)
	return lc
}

// Hint sets the input hint displayed when focused
func (lc LocationPicker) Hint(value string) LocationPicker {
	lc.set("hint", value)
	return lc
}

// Horizontal sets the horizontal layout configuration
func (lc LocationPicker) Horizontal(value string) LocationPicker {
	lc.set("horizontal", value)
	return lc
}

// ID sets the unique component ID, mainly for logging
func (lc LocationPicker) ID(value string) LocationPicker {
	lc.set("id", value)
	return lc
}

// InitAutoFill sets the initial auto-fill configuration
func (lc LocationPicker) InitAutoFill(value string) LocationPicker {
	lc.set("initAutoFill", value)
	return lc
}

// Inline sets whether the form control is in inline mode
func (lc LocationPicker) Inline(value bool) LocationPicker {
	lc.set("inline", value)
	return lc
}

// InputClassName sets the input class name
func (lc LocationPicker) InputClassName(value string) LocationPicker {
	lc.set("inputClassName", value)
	return lc
}

// Label sets the label text
func (lc LocationPicker) Label(value string) LocationPicker {
	lc.set("label", value)
	return lc
}

// LabelAlign sets the label alignment
func (lc LocationPicker) LabelAlign(value string) LocationPicker {
	lc.set("labelAlign", value)
	return lc
}

// LabelClassName sets the label class name
func (lc LocationPicker) LabelClassName(value string) LocationPicker {
	lc.set("labelClassName", value)
	return lc
}

// LabelRemark sets a small icon with a tooltip on hover
func (lc LocationPicker) LabelRemark(value string) LocationPicker {
	lc.set("labelRemark", value)
	return lc
}

// LabelWidth sets the custom label width, default unit is px
func (lc LocationPicker) LabelWidth(value string) LocationPicker {
	lc.set("labelWidth", value)
	return lc
}

// Mode sets the display mode of the form item
func (lc LocationPicker) Mode(value string) LocationPicker {
	lc.set("mode", value)
	return lc
}

// Name sets the field name, used as the key when submitting the form
func (lc LocationPicker) Name(value string) LocationPicker {
	lc.set("name", value)
	return lc
}

// OnEvent sets the event action configuration
func (lc LocationPicker) OnEvent(value any) LocationPicker {
	lc.set("onEvent", value)
	return lc
}

// OnlySelectCurrentLoc sets whether to restrict selection to the current location
func (lc LocationPicker) OnlySelectCurrentLoc(value bool) LocationPicker {
	lc.set("onlySelectCurrentLoc", value)
	return lc
}

// Placeholder sets the placeholder text
func (lc LocationPicker) Placeholder(value string) LocationPicker {
	lc.set("placeholder", value)
	return lc
}

// ReadOnly sets whether the component is read-only
func (lc LocationPicker) ReadOnly(value bool) LocationPicker {
	lc.set("readOnly", value)
	return lc
}

// ReadOnlyOn sets the expression to make the component read-only
func (lc LocationPicker) ReadOnlyOn(value string) LocationPicker {
	lc.set("readOnlyOn", value)
	return lc
}

// Remark sets a small icon with a tooltip on hover
func (lc LocationPicker) Remark(value string) LocationPicker {
	lc.set("remark", value)
	return lc
}

// Required sets whether the component is required
func (lc LocationPicker) Required(value bool) LocationPicker {
	lc.set("required", value)
	return lc
}

// Row sets the number of rows
func (lc LocationPicker) Row(value string) LocationPicker {
	lc.set("row", value)
	return lc
}

// SaveImmediately sets whether to save immediately
func (lc LocationPicker) SaveImmediately(value bool) LocationPicker {
	lc.set("saveImmediately", value)
	return lc
}

// Size sets the size of the form item
func (lc LocationPicker) Size(value string) LocationPicker {
	lc.set("size", value)
	return lc
}

// Static sets whether the component is displayed statically
func (lc LocationPicker) Static(value bool) LocationPicker {
	lc.set("static", value)
	return lc
}

// StaticClassName sets the class name for static display
func (lc LocationPicker) StaticClassName(value string) LocationPicker {
	lc.set("staticClassName", value)
	return lc
}

// StaticInputClassName sets the class name for static display value
func (lc LocationPicker) StaticInputClassName(value string) LocationPicker {
	lc.set("staticInputClassName", value)
	return lc
}

// StaticLabelClassName sets the class name for static display label
func (lc LocationPicker) StaticLabelClassName(value string) LocationPicker {
	lc.set("staticLabelClassName", value)
	return lc
}

// StaticOn sets the expression for static display
func (lc LocationPicker) StaticOn(value string) LocationPicker {
	lc.set("staticOn", value)
	return lc
}

// StaticPlaceholder sets the placeholder for static display
func (lc LocationPicker) StaticPlaceholder(value string) LocationPicker {
	lc.set("staticPlaceholder", value)
	return lc
}

// StaticSchema sets the schema for static display
func (lc LocationPicker) StaticSchema(value string) LocationPicker {
	lc.set("staticSchema", value)
	return lc
}

// Style sets the component style
func (lc LocationPicker) Style(value any) LocationPicker {
	lc.set("style", value)
	return lc
}

// SubmitOnChange sets whether to submit the form when changed
func (lc LocationPicker) SubmitOnChange(value bool) LocationPicker {
	lc.set("submitOnChange", value)
	return lc
}

// TestIdBuilder sets the test ID builder
func (lc LocationPicker) TestIdBuilder(value string) LocationPicker {
	lc.set("testIdBuilder", value)
	return lc
}

// UseMobileUI sets whether to disable mobile styles at the component level
func (lc LocationPicker) UseMobileUI(value bool) LocationPicker {
	lc.set("useMobileUI", value)
	return lc
}

// ValidateApi sets the remote validation API
func (lc LocationPicker) ValidateApi(value string) LocationPicker {
	lc.set("validateApi", value)
	return lc
}

// ValidateOnChange sets whether to revalidate on change after form submission
func (lc LocationPicker) ValidateOnChange(value bool) LocationPicker {
	lc.set("validateOnChange", value)
	return lc
}

// ValidationErrors sets the validation error messages
func (lc LocationPicker) ValidationErrors(value string) LocationPicker {
	lc.set("validationErrors", value)
	return lc
}

// Validations sets the validation configuration
func (lc LocationPicker) Validations(value string) LocationPicker {
	lc.set("validations", value)
	return lc
}

// Value sets the default value
func (lc LocationPicker) Value(value string) LocationPicker {
	lc.set("value", value)
	return lc
}

// Vendor sets the map type
func (lc LocationPicker) Vendor(value string) LocationPicker {
	lc.set("vendor", value)
	return lc
}

// Visible sets whether the component is visible
func (lc LocationPicker) Visible(value bool) LocationPicker {
	lc.set("visible", value)
	return lc
}

// VisibleOn sets the expression to show the component
func (lc LocationPicker) VisibleOn(value string) LocationPicker {
	lc.set("visibleOn", value)
	return lc
}

// Width sets the width in a table
func (lc LocationPicker) Width(value string) LocationPicker {
	lc.set("width", value)
	return lc
}
