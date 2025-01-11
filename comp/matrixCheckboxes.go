package comp

import "github.com/zrcoder/amisgo/model"

// matrixCheckboxes represents a checkbox matrix control.

type matrixCheckboxes model.Schema

// MatrixCheckboxes creates a new MatrixControl instance.
func MatrixCheckboxes() matrixCheckboxes {
	return matrixCheckboxes{"type": "matrix-checkboxes"}
}

func (mc matrixCheckboxes) set(key string, value any) matrixCheckboxes {
	mc[key] = value
	return mc
}

// AutoFill sets autoFill value.
func (mc matrixCheckboxes) AutoFill(value string) matrixCheckboxes {
	mc.set("autoFill", value)
	return mc
}

// ClassName sets the container CSS class name.
func (mc matrixCheckboxes) ClassName(value string) matrixCheckboxes {
	mc.set("className", value)
	return mc
}

// ClearValueOnHidden sets whether to clear the value when hidden.
func (mc matrixCheckboxes) ClearValueOnHidden(value bool) matrixCheckboxes {
	mc.set("clearValueOnHidden", value)
	return mc
}

// Columns sets the columns.
func (mc matrixCheckboxes) Columns(value ...any) matrixCheckboxes {
	mc.set("columns", value)
	return mc
}

// Desc sets the description.
func (mc matrixCheckboxes) Desc(value string) matrixCheckboxes {
	mc.set("desc", value)
	return mc
}

// Description sets the description, supports HTML.
func (mc matrixCheckboxes) Description(value string) matrixCheckboxes {
	mc.set("description", value)
	return mc
}

// DescriptionClassName sets the description CSS class name.
func (mc matrixCheckboxes) DescriptionClassName(value string) matrixCheckboxes {
	mc.set("descriptionClassName", value)
	return mc
}

// Disabled sets whether the control is disabled.
func (mc matrixCheckboxes) Disabled(value bool) matrixCheckboxes {
	mc.set("disabled", value)
	return mc
}

// DisabledOn sets the disabled expression.
func (mc matrixCheckboxes) DisabledOn(value string) matrixCheckboxes {
	mc.set("disabledOn", value)
	return mc
}

// EditorSetting sets the editor configuration.
func (mc matrixCheckboxes) EditorSetting(value string) matrixCheckboxes {
	mc.set("editorSetting", value)
	return mc
}

// ExtraName sets the extra field name.
func (mc matrixCheckboxes) ExtraName(value string) matrixCheckboxes {
	mc.set("extraName", value)
	return mc
}

// Hidden sets whether the control is hidden.
func (mc matrixCheckboxes) Hidden(value bool) matrixCheckboxes {
	mc.set("hidden", value)
	return mc
}

// HiddenOn sets the hidden expression.
func (mc matrixCheckboxes) HiddenOn(value string) matrixCheckboxes {
	mc.set("hiddenOn", value)
	return mc
}

// Hint sets the input hint.
func (mc matrixCheckboxes) Hint(value string) matrixCheckboxes {
	mc.set("hint", value)
	return mc
}

// Horizontal sets the horizontal layout configuration.
func (mc matrixCheckboxes) Horizontal(value string) matrixCheckboxes {
	mc.set("horizontal", value)
	return mc
}

// ID sets the unique component ID.
func (mc matrixCheckboxes) ID(value string) matrixCheckboxes {
	mc.set("id", value)
	return mc
}

// InitAutoFill sets the initial autoFill value.
func (mc matrixCheckboxes) InitAutoFill(value string) matrixCheckboxes {
	mc.set("initAutoFill", value)
	return mc
}

// Inline sets whether the control is inline.
func (mc matrixCheckboxes) Inline(value bool) matrixCheckboxes {
	mc.set("inline", value)
	return mc
}

// InputClassName sets the input CSS class name.
func (mc matrixCheckboxes) InputClassName(value string) matrixCheckboxes {
	mc.set("inputClassName", value)
	return mc
}

// Label sets the label.
func (mc matrixCheckboxes) Label(value string) matrixCheckboxes {
	mc.set("label", value)
	return mc
}

// LabelAlign sets the label alignment.
func (mc matrixCheckboxes) LabelAlign(value string) matrixCheckboxes {
	mc.set("labelAlign", value)
	return mc
}

// LabelClassName sets the label CSS class name.
func (mc matrixCheckboxes) LabelClassName(value string) matrixCheckboxes {
	mc.set("labelClassName", value)
	return mc
}

// LabelRemark sets the label remark.
func (mc matrixCheckboxes) LabelRemark(value string) matrixCheckboxes {
	mc.set("labelRemark", value)
	return mc
}

// LabelWidth sets the label width.
func (mc matrixCheckboxes) LabelWidth(value string) matrixCheckboxes {
	mc.set("labelWidth", value)
	return mc
}

// Mode sets the display mode.
func (mc matrixCheckboxes) Mode(value string) matrixCheckboxes {
	mc.set("mode", value)
	return mc
}

// Multiple sets whether multiple selection is allowed.
func (mc matrixCheckboxes) Multiple(value bool) matrixCheckboxes {
	mc.set("multiple", value)
	return mc
}

// Name sets the field name.
func (mc matrixCheckboxes) Name(value string) matrixCheckboxes {
	mc.set("name", value)
	return mc
}

// OnEvent sets the event configuration.
func (mc matrixCheckboxes) OnEvent(value any) matrixCheckboxes {
	mc.set("onEvent", value)
	return mc
}

// Placeholder sets the placeholder.
func (mc matrixCheckboxes) Placeholder(value string) matrixCheckboxes {
	mc.set("placeholder", value)
	return mc
}

// ReadOnly sets whether the control is read-only.
func (mc matrixCheckboxes) ReadOnly(value bool) matrixCheckboxes {
	mc.set("readOnly", value)
	return mc
}

// ReadOnlyOn sets the read-only expression.
func (mc matrixCheckboxes) ReadOnlyOn(value string) matrixCheckboxes {
	mc.set("readOnlyOn", value)
	return mc
}

// Remark sets the remark.
func (mc matrixCheckboxes) Remark(value string) matrixCheckboxes {
	mc.set("remark", value)
	return mc
}

// Required sets whether the control is required.
func (mc matrixCheckboxes) Required(value bool) matrixCheckboxes {
	mc.set("required", value)
	return mc
}

// Row sets the row.
func (mc matrixCheckboxes) Row(value string) matrixCheckboxes {
	mc.set("row", value)
	return mc
}

// RowLabel sets the row label.
func (mc matrixCheckboxes) RowLabel(value string) matrixCheckboxes {
	mc.set("rowLabel", value)
	return mc
}

// Rows sets the rows.
func (mc matrixCheckboxes) Rows(value string) matrixCheckboxes {
	mc.set("rows", value)
	return mc
}

// SaveImmediately sets whether to save immediately.
func (mc matrixCheckboxes) SaveImmediately(value bool) matrixCheckboxes {
	mc.set("saveImmediately", value)
	return mc
}

// SingleSelectMode sets the single select mode.
func (mc matrixCheckboxes) SingleSelectMode(value bool) matrixCheckboxes {
	mc.set("singleSelectMode", value)
	return mc
}

// Size sets the size.
func (mc matrixCheckboxes) Size(value string) matrixCheckboxes {
	mc.set("size", value)
	return mc
}

// Source sets the source for options.
func (mc matrixCheckboxes) Source(value string) matrixCheckboxes {
	mc.set("source", value)
	return mc
}

// Static sets whether the control is static.
func (mc matrixCheckboxes) Static(value bool) matrixCheckboxes {
	mc.set("static", value)
	return mc
}

// StaticClassName sets the static CSS class name.
func (mc matrixCheckboxes) StaticClassName(value string) matrixCheckboxes {
	mc.set("staticClassName", value)
	return mc
}

// StaticInputClassName sets the static input CSS class name.
func (mc matrixCheckboxes) StaticInputClassName(value string) matrixCheckboxes {
	mc.set("staticInputClassName", value)
	return mc
}

// StaticLabelClassName sets the static label CSS class name.
func (mc matrixCheckboxes) StaticLabelClassName(value string) matrixCheckboxes {
	mc.set("staticLabelClassName", value)
	return mc
}

// StaticOn sets the static expression.
func (mc matrixCheckboxes) StaticOn(value string) matrixCheckboxes {
	mc.set("staticOn", value)
	return mc
}

// StaticPlaceholder sets the static placeholder.
func (mc matrixCheckboxes) StaticPlaceholder(value string) matrixCheckboxes {
	mc.set("staticPlaceholder", value)
	return mc
}

// StaticSchema sets the static schema.
func (mc matrixCheckboxes) StaticSchema(value string) matrixCheckboxes {
	mc.set("staticSchema", value)
	return mc
}

// Style sets the component style.
func (mc matrixCheckboxes) Style(value any) matrixCheckboxes {
	mc.set("style", value)
	return mc
}

// SubmitOnChange sets whether to submit on change.
func (mc matrixCheckboxes) SubmitOnChange(value bool) matrixCheckboxes {
	mc.set("submitOnChange", value)
	return mc
}

// TestIdBuilder sets the test ID builder.
func (mc matrixCheckboxes) TestIdBuilder(value string) matrixCheckboxes {
	mc.set("testIdBuilder", value)
	return mc
}

// UseMobileUI sets whether to use mobile UI.
func (mc matrixCheckboxes) UseMobileUI(value bool) matrixCheckboxes {
	mc.set("useMobileUI", value)
	return mc
}

// ValidateApi sets the validation API.
func (mc matrixCheckboxes) ValidateApi(value string) matrixCheckboxes {
	mc.set("validateApi", value)
	return mc
}

// ValidateOnChange sets whether to validate on change.
func (mc matrixCheckboxes) ValidateOnChange(value bool) matrixCheckboxes {
	mc.set("validateOnChange", value)
	return mc
}

// ValidationErrors sets the validation error messages.
func (mc matrixCheckboxes) ValidationErrors(value string) matrixCheckboxes {
	mc.set("validationErrors", value)
	return mc
}

// Validations sets the validations.
func (mc matrixCheckboxes) Validations(value string) matrixCheckboxes {
	mc.set("validations", value)
	return mc
}

// Value sets the default value.
func (mc matrixCheckboxes) Value(value string) matrixCheckboxes {
	mc.set("value", value)
	return mc
}

// Visible sets whether the control is visible.
func (mc matrixCheckboxes) Visible(value bool) matrixCheckboxes {
	mc.set("visible", value)
	return mc
}

// VisibleOn sets the visibility expression.
func (mc matrixCheckboxes) VisibleOn(value string) matrixCheckboxes {
	mc.set("visibleOn", value)
	return mc
}

// Width sets the width in a table.
func (mc matrixCheckboxes) Width(value string) matrixCheckboxes {
	mc.set("width", value)
	return mc
}
