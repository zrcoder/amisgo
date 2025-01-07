package comp

// locationPicker component
// Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/form/location-picker

type locationPicker Schema

// LocationPicker creates a new LocationPicker instance
func LocationPicker() locationPicker {
	return locationPicker{}.set("type", "location-picker")
}

func (lc locationPicker) set(key string, value any) locationPicker {
	lc[key] = value
	return lc
}

// Ak sets the ak information for some maps
func (lc locationPicker) Ak(value string) locationPicker {
	lc.set("ak", value)
	return lc
}

// AutoFill sets autoFill to sync other values to the form when an option is selected
func (lc locationPicker) AutoFill(value string) locationPicker {
	lc.set("autoFill", value)
	return lc
}

// AutoSelectCurrentLoc sets whether to auto-select the current location
func (lc locationPicker) AutoSelectCurrentLoc(value bool) locationPicker {
	lc.set("autoSelectCurrentLoc", value)
	return lc
}

// ClassName sets the container CSS class name
func (lc locationPicker) ClassName(value string) locationPicker {
	lc.set("className", value)
	return lc
}

// ClearValueOnHidden sets whether to remove the form item value when hidden
func (lc locationPicker) ClearValueOnHidden(value bool) locationPicker {
	lc.set("clearValueOnHidden", value)
	return lc
}

// Desc sets the description content
func (lc locationPicker) Desc(value string) locationPicker {
	lc.set("desc", value)
	return lc
}

// Description sets the description content, supports HTML
func (lc locationPicker) Description(value string) locationPicker {
	lc.set("description", value)
	return lc
}

// DescriptionClassName sets the class name for the description
func (lc locationPicker) DescriptionClassName(value string) locationPicker {
	lc.set("descriptionClassName", value)
	return lc
}

// Disabled sets whether the component is disabled
func (lc locationPicker) Disabled(value bool) locationPicker {
	lc.set("disabled", value)
	return lc
}

// DisabledOn sets the expression to disable the component
func (lc locationPicker) DisabledOn(value string) locationPicker {
	lc.set("disabledOn", value)
	return lc
}

// EditorSetting sets the editor configuration, ignored at runtime
func (lc locationPicker) EditorSetting(value string) locationPicker {
	lc.set("editorSetting", value)
	return lc
}

// ExtraName sets an extra field name for range components
func (lc locationPicker) ExtraName(value string) locationPicker {
	lc.set("extraName", value)
	return lc
}

// GetLocationPlaceholder sets the placeholder when in read-only mode
func (lc locationPicker) GetLocationPlaceholder(value string) locationPicker {
	lc.set("getLocationPlaceholder", value)
	return lc
}

// Hidden sets whether the component is hidden
func (lc locationPicker) Hidden(value bool) locationPicker {
	lc.set("hidden", value)
	return lc
}

// HiddenOn sets the expression to hide the component
func (lc locationPicker) HiddenOn(value string) locationPicker {
	lc.set("hiddenOn", value)
	return lc
}

// Hint sets the input hint displayed when focused
func (lc locationPicker) Hint(value string) locationPicker {
	lc.set("hint", value)
	return lc
}

// Horizontal sets the horizontal layout configuration
func (lc locationPicker) Horizontal(value string) locationPicker {
	lc.set("horizontal", value)
	return lc
}

// ID sets the unique component ID, mainly for logging
func (lc locationPicker) ID(value string) locationPicker {
	lc.set("id", value)
	return lc
}

// InitAutoFill sets the initial auto-fill configuration
func (lc locationPicker) InitAutoFill(value string) locationPicker {
	lc.set("initAutoFill", value)
	return lc
}

// Inline sets whether the form control is in inline mode
func (lc locationPicker) Inline(value bool) locationPicker {
	lc.set("inline", value)
	return lc
}

// InputClassName sets the input class name
func (lc locationPicker) InputClassName(value string) locationPicker {
	lc.set("inputClassName", value)
	return lc
}

// Label sets the label text
func (lc locationPicker) Label(value string) locationPicker {
	lc.set("label", value)
	return lc
}

// LabelAlign sets the label alignment
func (lc locationPicker) LabelAlign(value string) locationPicker {
	lc.set("labelAlign", value)
	return lc
}

// LabelClassName sets the label class name
func (lc locationPicker) LabelClassName(value string) locationPicker {
	lc.set("labelClassName", value)
	return lc
}

// LabelRemark sets a small icon with a tooltip on hover
func (lc locationPicker) LabelRemark(value string) locationPicker {
	lc.set("labelRemark", value)
	return lc
}

// LabelWidth sets the custom label width, default unit is px
func (lc locationPicker) LabelWidth(value string) locationPicker {
	lc.set("labelWidth", value)
	return lc
}

// Mode sets the display mode of the form item
func (lc locationPicker) Mode(value string) locationPicker {
	lc.set("mode", value)
	return lc
}

// Name sets the field name, used as the key when submitting the form
func (lc locationPicker) Name(value string) locationPicker {
	lc.set("name", value)
	return lc
}

// OnEvent sets the event action configuration
func (lc locationPicker) OnEvent(value any) locationPicker {
	lc.set("onEvent", value)
	return lc
}

// OnlySelectCurrentLoc sets whether to restrict selection to the current location
func (lc locationPicker) OnlySelectCurrentLoc(value bool) locationPicker {
	lc.set("onlySelectCurrentLoc", value)
	return lc
}

// Placeholder sets the placeholder text
func (lc locationPicker) Placeholder(value string) locationPicker {
	lc.set("placeholder", value)
	return lc
}

// ReadOnly sets whether the component is read-only
func (lc locationPicker) ReadOnly(value bool) locationPicker {
	lc.set("readOnly", value)
	return lc
}

// ReadOnlyOn sets the expression to make the component read-only
func (lc locationPicker) ReadOnlyOn(value string) locationPicker {
	lc.set("readOnlyOn", value)
	return lc
}

// Remark sets a small icon with a tooltip on hover
func (lc locationPicker) Remark(value string) locationPicker {
	lc.set("remark", value)
	return lc
}

// Required sets whether the component is required
func (lc locationPicker) Required(value bool) locationPicker {
	lc.set("required", value)
	return lc
}

// Row sets the number of rows
func (lc locationPicker) Row(value string) locationPicker {
	lc.set("row", value)
	return lc
}

// SaveImmediately sets whether to save immediately
func (lc locationPicker) SaveImmediately(value bool) locationPicker {
	lc.set("saveImmediately", value)
	return lc
}

// Size sets the size of the form item
func (lc locationPicker) Size(value string) locationPicker {
	lc.set("size", value)
	return lc
}

// Static sets whether the component is displayed statically
func (lc locationPicker) Static(value bool) locationPicker {
	lc.set("static", value)
	return lc
}

// StaticClassName sets the class name for static display
func (lc locationPicker) StaticClassName(value string) locationPicker {
	lc.set("staticClassName", value)
	return lc
}

// StaticInputClassName sets the class name for static display value
func (lc locationPicker) StaticInputClassName(value string) locationPicker {
	lc.set("staticInputClassName", value)
	return lc
}

// StaticLabelClassName sets the class name for static display label
func (lc locationPicker) StaticLabelClassName(value string) locationPicker {
	lc.set("staticLabelClassName", value)
	return lc
}

// StaticOn sets the expression for static display
func (lc locationPicker) StaticOn(value string) locationPicker {
	lc.set("staticOn", value)
	return lc
}

// StaticPlaceholder sets the placeholder for static display
func (lc locationPicker) StaticPlaceholder(value string) locationPicker {
	lc.set("staticPlaceholder", value)
	return lc
}

// StaticSchema sets the schema for static display
func (lc locationPicker) StaticSchema(value string) locationPicker {
	lc.set("staticSchema", value)
	return lc
}

// Style sets the component style
func (lc locationPicker) Style(value any) locationPicker {
	lc.set("style", value)
	return lc
}

// SubmitOnChange sets whether to submit the form when changed
func (lc locationPicker) SubmitOnChange(value bool) locationPicker {
	lc.set("submitOnChange", value)
	return lc
}

// TestIdBuilder sets the test ID builder
func (lc locationPicker) TestIdBuilder(value string) locationPicker {
	lc.set("testIdBuilder", value)
	return lc
}

// UseMobileUI sets whether to disable mobile styles at the component level
func (lc locationPicker) UseMobileUI(value bool) locationPicker {
	lc.set("useMobileUI", value)
	return lc
}

// ValidateApi sets the remote validation API
func (lc locationPicker) ValidateApi(value string) locationPicker {
	lc.set("validateApi", value)
	return lc
}

// ValidateOnChange sets whether to revalidate on change after form submission
func (lc locationPicker) ValidateOnChange(value bool) locationPicker {
	lc.set("validateOnChange", value)
	return lc
}

// ValidationErrors sets the validation error messages
func (lc locationPicker) ValidationErrors(value string) locationPicker {
	lc.set("validationErrors", value)
	return lc
}

// Validations sets the validation configuration
func (lc locationPicker) Validations(value string) locationPicker {
	lc.set("validations", value)
	return lc
}

// Value sets the default value
func (lc locationPicker) Value(value string) locationPicker {
	lc.set("value", value)
	return lc
}

// Vendor sets the map type
func (lc locationPicker) Vendor(value string) locationPicker {
	lc.set("vendor", value)
	return lc
}

// Visible sets whether the component is visible
func (lc locationPicker) Visible(value bool) locationPicker {
	lc.set("visible", value)
	return lc
}

// VisibleOn sets the expression to show the component
func (lc locationPicker) VisibleOn(value string) locationPicker {
	lc.set("visibleOn", value)
	return lc
}

// Width sets the width in a table
func (lc locationPicker) Width(value string) locationPicker {
	lc.set("width", value)
	return lc
}
