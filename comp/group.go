package comp

// group represents a form group renderer that allows multiple forms to be displayed in a row.
type group Schema

// Group creates a new GroupControl instance.
func Group() group {
	return make(group).set("type", "group")
}

func (g group) set(key string, value any) group {
	g[key] = value
	return g
}

// AutoFill sets the autoFill property to automatically fill values when an option is selected.
func (g group) AutoFill(value string) group {
	return g.set("autoFill", value)
}

// Body sets the body property with form items.
func (g group) Body(value ...any) group {
	return g.set("body", value)
}

// ClassName sets the CSS class name for the container.
func (g group) ClassName(value string) group {
	return g.set("className", value)
}

// ClearValueOnHidden sets whether to remove the form item value when hidden.
func (g group) ClearValueOnHidden(value bool) group {
	return g.set("clearValueOnHidden", value)
}

// Desc sets the description.
func (g group) Desc(value string) group {
	return g.set("desc", value)
}

// Description sets the description content, supporting HTML fragments.
func (g group) Description(value string) group {
	return g.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description.
func (g group) DescriptionClassName(value string) group {
	return g.set("descriptionClassName", value)
}

// Direction sets the layout direction, either horizontal or vertical.
func (g group) Direction(value string) group {
	return g.set("direction", value)
}

// Disabled sets whether the group is disabled.
func (g group) Disabled(value bool) group {
	return g.set("disabled", value)
}

// DisabledOn sets the expression to determine if the group is disabled.
func (g group) DisabledOn(value string) group {
	return g.set("disabledOn", value)
}

// EditorSetting sets the editor configuration, ignored at runtime.
func (g group) EditorSetting(value string) group {
	return g.set("editorSetting", value)
}

// ExtraName sets an additional field name for range components.
func (g group) ExtraName(value string) group {
	return g.set("extraName", value)
}

// Gap sets the gap size, options are: xs, sm, normal.
func (g group) Gap(value string) group {
	return g.set("gap", value)
}

// Hidden sets whether the group is hidden.
func (g group) Hidden(value bool) group {
	return g.set("hidden", value)
}

// HiddenOn sets the expression to determine if the group is hidden.
func (g group) HiddenOn(value string) group {
	return g.set("hiddenOn", value)
}

// Hint sets the input hint, displayed when focused.
func (g group) Hint(value string) group {
	return g.set("hint", value)
}

// Horizontal sets the horizontal layout configuration.
func (g group) Horizontal(value string) group {
	return g.set("horizontal", value)
}

// ID sets the unique component ID, mainly for logging.
func (g group) ID(value string) group {
	return g.set("id", value)
}

// InitAutoFill sets the initial autoFill value.
func (g group) InitAutoFill(value string) group {
	return g.set("initAutoFill", value)
}

// Inline sets whether the form control is in inline mode.
func (g group) Inline(value bool) group {
	return g.set("inline", value)
}

// InputClassName sets the CSS class name for the input.
func (g group) InputClassName(value string) group {
	return g.set("inputClassName", value)
}

// Label sets the label text.
func (g group) Label(value string) group {
	return g.set("label", value)
}

// LabelAlign sets the label alignment, options are: right, left, top, inherit.
func (g group) LabelAlign(value string) group {
	return g.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label.
func (g group) LabelClassName(value string) group {
	return g.set("labelClassName", value)
}

// LabelRemark sets a small icon with a tooltip next to the label.
func (g group) LabelRemark(value string) group {
	return g.set("labelRemark", value)
}

// LabelWidth sets the custom width for the label, default unit is px.
func (g group) LabelWidth(value string) group {
	return g.set("labelWidth", value)
}

// Mode sets the display mode for the form item, options are: normal, inline, horizontal.
func (g group) Mode(value string) group {
	return g.set("mode", value)
}

// Name sets the field name for form submission, supports nested levels with dot notation.
func (g group) Name(value string) group {
	return g.set("name", value)
}

// OnEvent sets the event action configuration.
func (g group) OnEvent(value any) group {
	return g.set("onEvent", value)
}

// Placeholder sets the placeholder text.
func (g group) Placeholder(value string) group {
	return g.set("placeholder", value)
}

// ReadOnly sets whether the group is read-only.
func (g group) ReadOnly(value bool) group {
	return g.set("readOnly", value)
}

// ReadOnlyOn sets the expression to determine if the group is read-only.
func (g group) ReadOnlyOn(value string) group {
	return g.set("readOnlyOn", value)
}

// Remark sets a small icon with a tooltip.
func (g group) Remark(value string) group {
	return g.set("remark", value)
}

// Required sets whether the group is required.
func (g group) Required(value bool) group {
	return g.set("required", value)
}

// Row sets the row value.
func (g group) Row(value string) group {
	return g.set("row", value)
}

// SaveImmediately sets whether to save immediately (used in TableColumn).
func (g group) SaveImmediately(value bool) group {
	return g.set("saveImmediately", value)
}

// Size sets the size of the form item, options are: xs, sm, md, lg, full.
func (g group) Size(value string) group {
	return g.set("size", value)
}

// Static sets whether the group is displayed statically.
func (g group) Static(value bool) group {
	return g.set("static", value)
}

// StaticClassName sets the CSS class name for static display.
func (g group) StaticClassName(value string) group {
	return g.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input value.
func (g group) StaticInputClassName(value string) group {
	return g.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label.
func (g group) StaticLabelClassName(value string) group {
	return g.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the group is displayed statically.
func (g group) StaticOn(value string) group {
	return g.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display.
func (g group) StaticPlaceholder(value string) group {
	return g.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema.
func (g group) StaticSchema(value string) group {
	return g.set("staticSchema", value)
}

// Style sets the component style.
func (g group) Style(value any) group {
	return g.set("style", value)
}

// SubFormHorizontal sets the horizontal layout configuration for sub-forms.
func (g group) SubFormHorizontal(value string) group {
	return g.set("subFormHorizontal", value)
}

// SubFormMode sets the default display mode for sub-form items, options are: normal, inline, horizontal.
func (g group) SubFormMode(value string) group {
	return g.set("subFormMode", value)
}

// SubmitOnChange sets whether to submit the form when changed.
func (g group) SubmitOnChange(value bool) group {
	return g.set("submitOnChange", value)
}

// TestIdBuilder sets the test ID builder.
func (g group) TestIdBuilder(value string) group {
	return g.set("testIdBuilder", value)
}

// UseMobileUI sets whether to disable mobile UI styles.
func (g group) UseMobileUI(value bool) group {
	return g.set("useMobileUI", value)
}

// ValidateApi sets the remote validation API for the form item.
func (g group) ValidateApi(value string) group {
	return g.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on change after form submission.
func (g group) ValidateOnChange(value bool) group {
	return g.set("validateOnChange", value)
}

// ValidationErrors sets the validation error messages.
func (g group) ValidationErrors(value string) group {
	return g.set("validationErrors", value)
}

// Validations sets the validation rules.
func (g group) Validations(value string) group {
	return g.set("validations", value)
}

// Value sets the default value, must be static and not support variables.
func (g group) Value(value string) group {
	return g.set("value", value)
}

// Visible sets whether the group is visible.
func (g group) Visible(value bool) group {
	return g.set("visible", value)
}

// VisibleOn sets the expression to determine if the group is visible.
func (g group) VisibleOn(value string) group {
	return g.set("visibleOn", value)
}

// Width sets the width in a table.
func (g group) Width(value string) group {
	return g.set("width", value)
}
