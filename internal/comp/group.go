package comp

import "github.com/zrcoder/amisgo/schema"

// Group represents a form Group renderer that allows multiple forms to be displayed in a row.
type Group schema.Schema

func NewGroup() Group {
	return Group{"type": "group"}
}

func (g Group) set(key string, value any) Group {
	g[key] = value
	return g
}

// AutoFill sets the autoFill property to automatically fill values when an option is selected.
func (g Group) AutoFill(value string) Group {
	return g.set("autoFill", value)
}

// Body sets the body property with form items.
func (g Group) Body(value ...any) Group {
	return g.set("body", value)
}

// ClassName sets the CSS class name for the container.
func (g Group) ClassName(value string) Group {
	return g.set("className", value)
}

// ClearValueOnHidden sets whether to remove the form item value when hidden.
func (g Group) ClearValueOnHidden(value bool) Group {
	return g.set("clearValueOnHidden", value)
}

// Desc sets the description.
func (g Group) Desc(value string) Group {
	return g.set("desc", value)
}

// Description sets the description content, supporting HTML fragments.
func (g Group) Description(value string) Group {
	return g.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description.
func (g Group) DescriptionClassName(value string) Group {
	return g.set("descriptionClassName", value)
}

// Direction sets the layout direction, either horizontal or vertical.
func (g Group) Direction(value string) Group {
	return g.set("direction", value)
}

// Disabled sets whether the group is disabled.
func (g Group) Disabled(value bool) Group {
	return g.set("disabled", value)
}

// DisabledOn sets the expression to determine if the group is disabled.
func (g Group) DisabledOn(value string) Group {
	return g.set("disabledOn", value)
}

// EditorSetting sets the editor configuration, ignored at runtime.
func (g Group) EditorSetting(value string) Group {
	return g.set("editorSetting", value)
}

// ExtraName sets an additional field name for range components.
func (g Group) ExtraName(value string) Group {
	return g.set("extraName", value)
}

// Gap sets the gap size, options are: xs, sm, normal.
func (g Group) Gap(value string) Group {
	return g.set("gap", value)
}

// Hidden sets whether the group is hidden.
func (g Group) Hidden(value bool) Group {
	return g.set("hidden", value)
}

// HiddenOn sets the expression to determine if the group is hidden.
func (g Group) HiddenOn(value string) Group {
	return g.set("hiddenOn", value)
}

// Hint sets the input hint, displayed when focused.
func (g Group) Hint(value string) Group {
	return g.set("hint", value)
}

// Horizontal sets the horizontal layout configuration.
func (g Group) Horizontal(value string) Group {
	return g.set("horizontal", value)
}

// ID sets the unique component ID, mainly for logging.
func (g Group) ID(value string) Group {
	return g.set("id", value)
}

// InitAutoFill sets the initial autoFill value.
func (g Group) InitAutoFill(value string) Group {
	return g.set("initAutoFill", value)
}

// Inline sets whether the form control is in inline mode.
func (g Group) Inline(value bool) Group {
	return g.set("inline", value)
}

// InputClassName sets the CSS class name for the input.
func (g Group) InputClassName(value string) Group {
	return g.set("inputClassName", value)
}

// Label sets the label text.
func (g Group) Label(value string) Group {
	return g.set("label", value)
}

// LabelAlign sets the label alignment, options are: right, left, top, inherit.
func (g Group) LabelAlign(value string) Group {
	return g.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label.
func (g Group) LabelClassName(value string) Group {
	return g.set("labelClassName", value)
}

// LabelRemark sets a small icon with a tooltip next to the label.
func (g Group) LabelRemark(value string) Group {
	return g.set("labelRemark", value)
}

// LabelWidth sets the custom width for the label, default unit is px.
func (g Group) LabelWidth(value string) Group {
	return g.set("labelWidth", value)
}

// Mode sets the display mode for the form item, options are: normal, inline, horizontal.
func (g Group) Mode(value string) Group {
	return g.set("mode", value)
}

// Name sets the field name for form submission, supports nested levels with dot notation.
func (g Group) Name(value string) Group {
	return g.set("name", value)
}

// OnEvent sets the event action configuration.
func (g Group) OnEvent(value any) Group {
	return g.set("onEvent", value)
}

// Placeholder sets the placeholder text.
func (g Group) Placeholder(value string) Group {
	return g.set("placeholder", value)
}

// ReadOnly sets whether the group is read-only.
func (g Group) ReadOnly(value bool) Group {
	return g.set("readOnly", value)
}

// ReadOnlyOn sets the expression to determine if the group is read-only.
func (g Group) ReadOnlyOn(value string) Group {
	return g.set("readOnlyOn", value)
}

// Remark sets a small icon with a tooltip.
func (g Group) Remark(value string) Group {
	return g.set("remark", value)
}

// Required sets whether the group is required.
func (g Group) Required(value bool) Group {
	return g.set("required", value)
}

// Row sets the row value.
func (g Group) Row(value string) Group {
	return g.set("row", value)
}

// SaveImmediately sets whether to save immediately (used in TableColumn).
func (g Group) SaveImmediately(value bool) Group {
	return g.set("saveImmediately", value)
}

// Size sets the size of the form item, options are: xs, sm, md, lg, full.
func (g Group) Size(value string) Group {
	return g.set("size", value)
}

// Static sets whether the group is displayed statically.
func (g Group) Static(value bool) Group {
	return g.set("static", value)
}

// StaticClassName sets the CSS class name for static display.
func (g Group) StaticClassName(value string) Group {
	return g.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input value.
func (g Group) StaticInputClassName(value string) Group {
	return g.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label.
func (g Group) StaticLabelClassName(value string) Group {
	return g.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the group is displayed statically.
func (g Group) StaticOn(value string) Group {
	return g.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display.
func (g Group) StaticPlaceholder(value string) Group {
	return g.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema.Schema.
func (g Group) StaticSchema(value string) Group {
	return g.set("staticSchema", value)
}

// Style sets the component style.
func (g Group) Style(value any) Group {
	return g.set("style", value)
}

// SubFormHorizontal sets the horizontal layout configuration for sub-forms.
func (g Group) SubFormHorizontal(value string) Group {
	return g.set("subFormHorizontal", value)
}

// SubFormMode sets the default display mode for sub-form items, options are: normal, inline, horizontal.
func (g Group) SubFormMode(value string) Group {
	return g.set("subFormMode", value)
}

// SubmitOnChange sets whether to submit the form when changed.
func (g Group) SubmitOnChange(value bool) Group {
	return g.set("submitOnChange", value)
}

// TestIdBuilder sets the test ID builder.
func (g Group) TestIdBuilder(value string) Group {
	return g.set("testIdBuilder", value)
}

// UseMobileUI sets whether to disable mobile UI styles.
func (g Group) UseMobileUI(value bool) Group {
	return g.set("useMobileUI", value)
}

// ValidateApi sets the remote validation API for the form item.
func (g Group) ValidateApi(value string) Group {
	return g.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on change after form submission.
func (g Group) ValidateOnChange(value bool) Group {
	return g.set("validateOnChange", value)
}

// ValidationErrors sets the validation error messages.
func (g Group) ValidationErrors(value string) Group {
	return g.set("validationErrors", value)
}

// Validations sets the validation rules.
func (g Group) Validations(value string) Group {
	return g.set("validations", value)
}

// Value sets the default value, must be static and not support variables.
func (g Group) Value(value string) Group {
	return g.set("value", value)
}

// Visible sets whether the group is visible.
func (g Group) Visible(value bool) Group {
	return g.set("visible", value)
}

// VisibleOn sets the expression to determine if the group is visible.
func (g Group) VisibleOn(value string) Group {
	return g.set("visibleOn", value)
}

// Width sets the width in a table.
func (g Group) Width(value string) Group {
	return g.set("width", value)
}
