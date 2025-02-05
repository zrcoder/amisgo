package comp

import "github.com/zrcoder/amisgo/model"

// MatrixCheckboxes represents a checkbox matrix control.
type MatrixCheckboxes model.Schema

func NewMatrixCheckboxes() MatrixCheckboxes {
	return MatrixCheckboxes{"type": "matrix-checkboxes"}
}

func (mc MatrixCheckboxes) set(key string, value any) MatrixCheckboxes {
	mc[key] = value
	return mc
}

// AutoFill sets autoFill value.
func (mc MatrixCheckboxes) AutoFill(value string) MatrixCheckboxes {
	mc.set("autoFill", value)
	return mc
}

// ClassName sets the container CSS class name.
func (mc MatrixCheckboxes) ClassName(value string) MatrixCheckboxes {
	mc.set("className", value)
	return mc
}

// ClearValueOnHidden sets whether to clear the value when hidden.
func (mc MatrixCheckboxes) ClearValueOnHidden(value bool) MatrixCheckboxes {
	mc.set("clearValueOnHidden", value)
	return mc
}

// Columns sets the columns.
func (mc MatrixCheckboxes) Columns(value ...any) MatrixCheckboxes {
	mc.set("columns", value)
	return mc
}

// Desc sets the description.
func (mc MatrixCheckboxes) Desc(value string) MatrixCheckboxes {
	mc.set("desc", value)
	return mc
}

// Description sets the description, supports HTML.
func (mc MatrixCheckboxes) Description(value string) MatrixCheckboxes {
	mc.set("description", value)
	return mc
}

// DescriptionClassName sets the description CSS class name.
func (mc MatrixCheckboxes) DescriptionClassName(value string) MatrixCheckboxes {
	mc.set("descriptionClassName", value)
	return mc
}

// Disabled sets whether the control is disabled.
func (mc MatrixCheckboxes) Disabled(value bool) MatrixCheckboxes {
	mc.set("disabled", value)
	return mc
}

// DisabledOn sets the disabled expression.
func (mc MatrixCheckboxes) DisabledOn(value string) MatrixCheckboxes {
	mc.set("disabledOn", value)
	return mc
}

// EditorSetting sets the editor configuration.
func (mc MatrixCheckboxes) EditorSetting(value string) MatrixCheckboxes {
	mc.set("editorSetting", value)
	return mc
}

// ExtraName sets the extra field name.
func (mc MatrixCheckboxes) ExtraName(value string) MatrixCheckboxes {
	mc.set("extraName", value)
	return mc
}

// Hidden sets whether the control is hidden.
func (mc MatrixCheckboxes) Hidden(value bool) MatrixCheckboxes {
	mc.set("hidden", value)
	return mc
}

// HiddenOn sets the hidden expression.
func (mc MatrixCheckboxes) HiddenOn(value string) MatrixCheckboxes {
	mc.set("hiddenOn", value)
	return mc
}

// Hint sets the input hint.
func (mc MatrixCheckboxes) Hint(value string) MatrixCheckboxes {
	mc.set("hint", value)
	return mc
}

// Horizontal sets the horizontal layout configuration.
func (mc MatrixCheckboxes) Horizontal(value string) MatrixCheckboxes {
	mc.set("horizontal", value)
	return mc
}

// ID sets the unique component ID.
func (mc MatrixCheckboxes) ID(value string) MatrixCheckboxes {
	mc.set("id", value)
	return mc
}

// InitAutoFill sets the initial autoFill value.
func (mc MatrixCheckboxes) InitAutoFill(value string) MatrixCheckboxes {
	mc.set("initAutoFill", value)
	return mc
}

// Inline sets whether the control is inline.
func (mc MatrixCheckboxes) Inline(value bool) MatrixCheckboxes {
	mc.set("inline", value)
	return mc
}

// InputClassName sets the input CSS class name.
func (mc MatrixCheckboxes) InputClassName(value string) MatrixCheckboxes {
	mc.set("inputClassName", value)
	return mc
}

// Label sets the label.
func (mc MatrixCheckboxes) Label(value string) MatrixCheckboxes {
	mc.set("label", value)
	return mc
}

// LabelAlign sets the label alignment.
func (mc MatrixCheckboxes) LabelAlign(value string) MatrixCheckboxes {
	mc.set("labelAlign", value)
	return mc
}

// LabelClassName sets the label CSS class name.
func (mc MatrixCheckboxes) LabelClassName(value string) MatrixCheckboxes {
	mc.set("labelClassName", value)
	return mc
}

// LabelRemark sets the label remark.
func (mc MatrixCheckboxes) LabelRemark(value string) MatrixCheckboxes {
	mc.set("labelRemark", value)
	return mc
}

// LabelWidth sets the label width.
func (mc MatrixCheckboxes) LabelWidth(value string) MatrixCheckboxes {
	mc.set("labelWidth", value)
	return mc
}

// Mode sets the display mode.
func (mc MatrixCheckboxes) Mode(value string) MatrixCheckboxes {
	mc.set("mode", value)
	return mc
}

// Multiple sets whether multiple selection is allowed.
func (mc MatrixCheckboxes) Multiple(value bool) MatrixCheckboxes {
	mc.set("multiple", value)
	return mc
}

// Name sets the field name.
func (mc MatrixCheckboxes) Name(value string) MatrixCheckboxes {
	mc.set("name", value)
	return mc
}

// OnEvent sets the event configuration.
func (mc MatrixCheckboxes) OnEvent(value any) MatrixCheckboxes {
	mc.set("onEvent", value)
	return mc
}

// Placeholder sets the placeholder.
func (mc MatrixCheckboxes) Placeholder(value string) MatrixCheckboxes {
	mc.set("placeholder", value)
	return mc
}

// ReadOnly sets whether the control is read-only.
func (mc MatrixCheckboxes) ReadOnly(value bool) MatrixCheckboxes {
	mc.set("readOnly", value)
	return mc
}

// ReadOnlyOn sets the read-only expression.
func (mc MatrixCheckboxes) ReadOnlyOn(value string) MatrixCheckboxes {
	mc.set("readOnlyOn", value)
	return mc
}

// Remark sets the remark.
func (mc MatrixCheckboxes) Remark(value string) MatrixCheckboxes {
	mc.set("remark", value)
	return mc
}

// Required sets whether the control is required.
func (mc MatrixCheckboxes) Required(value bool) MatrixCheckboxes {
	mc.set("required", value)
	return mc
}

// Row sets the row.
func (mc MatrixCheckboxes) Row(value string) MatrixCheckboxes {
	mc.set("row", value)
	return mc
}

// RowLabel sets the row label.
func (mc MatrixCheckboxes) RowLabel(value string) MatrixCheckboxes {
	mc.set("rowLabel", value)
	return mc
}

// Rows sets the rows.
func (mc MatrixCheckboxes) Rows(value string) MatrixCheckboxes {
	mc.set("rows", value)
	return mc
}

// SaveImmediately sets whether to save immediately.
func (mc MatrixCheckboxes) SaveImmediately(value bool) MatrixCheckboxes {
	mc.set("saveImmediately", value)
	return mc
}

// SingleSelectMode sets the single select mode.
func (mc MatrixCheckboxes) SingleSelectMode(value bool) MatrixCheckboxes {
	mc.set("singleSelectMode", value)
	return mc
}

// Size sets the size.
func (mc MatrixCheckboxes) Size(value string) MatrixCheckboxes {
	mc.set("size", value)
	return mc
}

// Source sets the source for options.
func (mc MatrixCheckboxes) Source(value string) MatrixCheckboxes {
	mc.set("source", value)
	return mc
}

// Static sets whether the control is static.
func (mc MatrixCheckboxes) Static(value bool) MatrixCheckboxes {
	mc.set("static", value)
	return mc
}

// StaticClassName sets the static CSS class name.
func (mc MatrixCheckboxes) StaticClassName(value string) MatrixCheckboxes {
	mc.set("staticClassName", value)
	return mc
}

// StaticInputClassName sets the static input CSS class name.
func (mc MatrixCheckboxes) StaticInputClassName(value string) MatrixCheckboxes {
	mc.set("staticInputClassName", value)
	return mc
}

// StaticLabelClassName sets the static label CSS class name.
func (mc MatrixCheckboxes) StaticLabelClassName(value string) MatrixCheckboxes {
	mc.set("staticLabelClassName", value)
	return mc
}

// StaticOn sets the static expression.
func (mc MatrixCheckboxes) StaticOn(value string) MatrixCheckboxes {
	mc.set("staticOn", value)
	return mc
}

// StaticPlaceholder sets the static placeholder.
func (mc MatrixCheckboxes) StaticPlaceholder(value string) MatrixCheckboxes {
	mc.set("staticPlaceholder", value)
	return mc
}

// StaticSchema sets the static schema.
func (mc MatrixCheckboxes) StaticSchema(value string) MatrixCheckboxes {
	mc.set("staticSchema", value)
	return mc
}

// Style sets the component style.
func (mc MatrixCheckboxes) Style(value any) MatrixCheckboxes {
	mc.set("style", value)
	return mc
}

// SubmitOnChange sets whether to submit on change.
func (mc MatrixCheckboxes) SubmitOnChange(value bool) MatrixCheckboxes {
	mc.set("submitOnChange", value)
	return mc
}

// TestIdBuilder sets the test ID builder.
func (mc MatrixCheckboxes) TestIdBuilder(value string) MatrixCheckboxes {
	mc.set("testIdBuilder", value)
	return mc
}

// UseMobileUI sets whether to use mobile UI.
func (mc MatrixCheckboxes) UseMobileUI(value bool) MatrixCheckboxes {
	mc.set("useMobileUI", value)
	return mc
}

// ValidateApi sets the validation API.
func (mc MatrixCheckboxes) ValidateApi(value string) MatrixCheckboxes {
	mc.set("validateApi", value)
	return mc
}

// ValidateOnChange sets whether to validate on change.
func (mc MatrixCheckboxes) ValidateOnChange(value bool) MatrixCheckboxes {
	mc.set("validateOnChange", value)
	return mc
}

// ValidationErrors sets the validation error messages.
func (mc MatrixCheckboxes) ValidationErrors(value string) MatrixCheckboxes {
	mc.set("validationErrors", value)
	return mc
}

// Validations sets the validations.
func (mc MatrixCheckboxes) Validations(value string) MatrixCheckboxes {
	mc.set("validations", value)
	return mc
}

// Value sets the default value.
func (mc MatrixCheckboxes) Value(value string) MatrixCheckboxes {
	mc.set("value", value)
	return mc
}

// Visible sets whether the control is visible.
func (mc MatrixCheckboxes) Visible(value bool) MatrixCheckboxes {
	mc.set("visible", value)
	return mc
}

// VisibleOn sets the visibility expression.
func (mc MatrixCheckboxes) VisibleOn(value string) MatrixCheckboxes {
	mc.set("visibleOn", value)
	return mc
}

// Width sets the width in a table.
func (mc MatrixCheckboxes) Width(value string) MatrixCheckboxes {
	mc.set("width", value)
	return mc
}
