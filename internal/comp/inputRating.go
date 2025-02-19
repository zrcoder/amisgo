package comp

import "github.com/zrcoder/amisgo/schema"

// InputRating documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/form/rating
type InputRating schema.Schema

func NewInputRating() InputRating {
	return InputRating{"type": "input-rating"}
}

func (rc InputRating) set(key string, value any) InputRating {
	rc[key] = value
	return rc
}

// AllowClear allows clearing the rating by clicking again
func (rc InputRating) AllowClear(value bool) InputRating {
	return rc.set("allowClear", value)
}

// AutoFill auto-fills other form values when selected
func (rc InputRating) AutoFill(value string) InputRating {
	return rc.set("autoFill", value)
}

// Char custom character
func (rc InputRating) Char(value string) InputRating {
	return rc.set("char", value)
}

// CharClassName custom character class name
func (rc InputRating) CharClassName(value string) InputRating {
	return rc.set("charClassName", value)
}

// ClassName container CSS class name
func (rc InputRating) ClassName(value string) InputRating {
	return rc.set("className", value)
}

// ClearValueOnHidden removes value from form when hidden
func (rc InputRating) ClearValueOnHidden(value bool) InputRating {
	return rc.set("clearValueOnHidden", value)
}

// Colors selected star color
func (rc InputRating) Colors(value string) InputRating {
	return rc.set("colors", value)
}

// Count number of stars
func (rc InputRating) Count(value any) InputRating {
	return rc.set("count", value)
}

// Desc description
func (rc InputRating) Desc(value string) InputRating {
	return rc.set("desc", value)
}

// Description HTML description content
func (rc InputRating) Description(value string) InputRating {
	return rc.set("description", value)
}

// DescriptionClassName description class name
func (rc InputRating) DescriptionClassName(value string) InputRating {
	return rc.set("descriptionClassName", value)
}

// Disabled disables the rating
func (rc InputRating) Disabled(value bool) InputRating {
	return rc.set("disabled", value)
}

// DisabledOn disable condition expression
func (rc InputRating) DisabledOn(value string) InputRating {
	return rc.set("disabledOn", value)
}

// EditorSetting editor configuration, ignored at runtime
func (rc InputRating) EditorSetting(value string) InputRating {
	return rc.set("editorSetting", value)
}

// ExtraName extra field name
func (rc InputRating) ExtraName(value string) InputRating {
	return rc.set("extraName", value)
}

// Half allows half-star rating
func (rc InputRating) Half(value bool) InputRating {
	return rc.set("half", value)
}

// Hidden hides the rating
func (rc InputRating) Hidden(value bool) InputRating {
	return rc.set("hidden", value)
}

// HiddenOn hide condition expression
func (rc InputRating) HiddenOn(value string) InputRating {
	return rc.set("hiddenOn", value)
}

// Hint input hint displayed on focus
func (rc InputRating) Hint(value string) InputRating {
	return rc.set("hint", value)
}

// Horizontal horizontal layout configuration
func (rc InputRating) Horizontal(value string) InputRating {
	return rc.set("horizontal", value)
}

// ID unique component ID
func (rc InputRating) ID(value string) InputRating {
	return rc.set("id", value)
}

// InactiveColor unselected star color
func (rc InputRating) InactiveColor(value string) InputRating {
	return rc.set("inactiveColor", value)
}

// InitAutoFill initial auto-fill value
func (rc InputRating) InitAutoFill(value string) InputRating {
	return rc.set("initAutoFill", value)
}

// Inline inline mode for form control
func (rc InputRating) Inline(value bool) InputRating {
	return rc.set("inline", value)
}

// InputClassName input class name
func (rc InputRating) InputClassName(value string) InputRating {
	return rc.set("inputClassName", value)
}

// Label label text
func (rc InputRating) Label(value string) InputRating {
	return rc.set("label", value)
}

// LabelAlign label alignment
func (rc InputRating) LabelAlign(value string) InputRating {
	return rc.set("labelAlign", value)
}

// LabelClassName label class name
func (rc InputRating) LabelClassName(value string) InputRating {
	return rc.set("labelClassName", value)
}

// LabelRemark label remark tooltip
func (rc InputRating) LabelRemark(value string) InputRating {
	return rc.set("labelRemark", value)
}

// LabelWidth custom label width
func (rc InputRating) LabelWidth(value string) InputRating {
	return rc.set("labelWidth", value)
}

// Mode display mode
func (rc InputRating) Mode(value string) InputRating {
	return rc.set("mode", value)
}

// Name field name
func (rc InputRating) Name(value string) InputRating {
	return rc.set("name", value)
}

// OnEvent event action configuration
func (rc InputRating) OnEvent(value any) InputRating {
	return rc.set("onEvent", value)
}

// Placeholder placeholder text
func (rc InputRating) Placeholder(value string) InputRating {
	return rc.set("placeholder", value)
}

// ReadOnly read-only mode
func (rc InputRating) ReadOnly(value bool) InputRating {
	return rc.set("readOnly", value)
}

// ReadOnlyOn read-only condition expression
func (rc InputRating) ReadOnlyOn(value string) InputRating {
	return rc.set("readOnlyOn", value)
}

// Remark remark tooltip
func (rc InputRating) Remark(value string) InputRating {
	return rc.set("remark", value)
}

// Required required field
func (rc InputRating) Required(value bool) InputRating {
	return rc.set("required", value)
}

// Row row configuration
func (rc InputRating) Row(value string) InputRating {
	return rc.set("row", value)
}

// SaveImmediately save immediately on change
func (rc InputRating) SaveImmediately(value bool) InputRating {
	return rc.set("saveImmediately", value)
}

// Size form item size
func (rc InputRating) Size(value string) InputRating {
	return rc.set("size", value)
}

// Static static display mode
func (rc InputRating) Static(value bool) InputRating {
	return rc.set("static", value)
}

// StaticClassName static display class name
func (rc InputRating) StaticClassName(value string) InputRating {
	return rc.set("staticClassName", value)
}

// StaticInputClassName static input value class name
func (rc InputRating) StaticInputClassName(value string) InputRating {
	return rc.set("staticInputClassName", value)
}

// StaticLabelClassName static label class name
func (rc InputRating) StaticLabelClassName(value string) InputRating {
	return rc.set("staticLabelClassName", value)
}

// StaticOn static display condition expression
func (rc InputRating) StaticOn(value string) InputRating {
	return rc.set("staticOn", value)
}

// StaticPlaceholder static display placeholder
func (rc InputRating) StaticPlaceholder(value string) InputRating {
	return rc.set("staticPlaceholder", value)
}

// StaticSchema static schema.Schema
func (rc InputRating) StaticSchema(value string) InputRating {
	return rc.set("staticSchema", value)
}

// Style component style
func (rc InputRating) Style(value any) InputRating {
	return rc.set("style", value)
}

// SubmitOnChange submit form on change
func (rc InputRating) SubmitOnChange(value bool) InputRating {
	return rc.set("submitOnChange", value)
}

// TestIdBuilder test ID builder
func (rc InputRating) TestIdBuilder(value string) InputRating {
	return rc.set("testIdBuilder", value)
}

// TextClassName custom text class name
func (rc InputRating) TextClassName(value string) InputRating {
	return rc.set("textClassName", value)
}

// TextPosition text position
func (rc InputRating) TextPosition(value string) InputRating {
	return rc.set("textPosition", value)
}

// Texts selected star text
func (rc InputRating) Texts(value string) InputRating {
	return rc.set("texts", value)
}

// UseMobileUI disable mobile UI styles
func (rc InputRating) UseMobileUI(value bool) InputRating {
	return rc.set("useMobileUI", value)
}

// ValidateApi remote validation API
func (rc InputRating) ValidateApi(value string) InputRating {
	return rc.set("validateApi", value)
}

// ValidateOnChange validate on change
func (rc InputRating) ValidateOnChange(value bool) InputRating {
	return rc.set("validateOnChange", value)
}

// ValidationErrors validation error messages
func (rc InputRating) ValidationErrors(value string) InputRating {
	return rc.set("validationErrors", value)
}

// Validations validation rules
func (rc InputRating) Validations(value string) InputRating {
	return rc.set("validations", value)
}

// Value default value
func (rc InputRating) Value(value any) InputRating {
	return rc.set("value", value)
}

// Visible visibility
func (rc InputRating) Visible(value bool) InputRating {
	return rc.set("visible", value)
}

// VisibleOn visibility condition expression
func (rc InputRating) VisibleOn(value string) InputRating {
	return rc.set("visibleOn", value)
}

// Width width adjustment in table
func (rc InputRating) Width(value string) InputRating {
	return rc.set("width", value)
}
