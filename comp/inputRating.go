package comp

import "github.com/zrcoder/amisgo/model"

// inputRating documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/form/rating

type inputRating model.Schema

func InputRating() inputRating {
	return inputRating{}.set("type", "input-rating")
}

func (rc inputRating) set(key string, value any) inputRating {
	rc[key] = value
	return rc
}

// AllowClear allows clearing the rating by clicking again
func (rc inputRating) AllowClear(value bool) inputRating {
	return rc.set("allowClear", value)
}

// AutoFill auto-fills other form values when selected
func (rc inputRating) AutoFill(value string) inputRating {
	return rc.set("autoFill", value)
}

// Char custom character
func (rc inputRating) Char(value string) inputRating {
	return rc.set("char", value)
}

// CharClassName custom character class name
func (rc inputRating) CharClassName(value string) inputRating {
	return rc.set("charClassName", value)
}

// ClassName container CSS class name
func (rc inputRating) ClassName(value string) inputRating {
	return rc.set("className", value)
}

// ClearValueOnHidden removes value from form when hidden
func (rc inputRating) ClearValueOnHidden(value bool) inputRating {
	return rc.set("clearValueOnHidden", value)
}

// Colors selected star color
func (rc inputRating) Colors(value string) inputRating {
	return rc.set("colors", value)
}

// Count number of stars
func (rc inputRating) Count(value any) inputRating {
	return rc.set("count", value)
}

// Desc description
func (rc inputRating) Desc(value string) inputRating {
	return rc.set("desc", value)
}

// Description HTML description content
func (rc inputRating) Description(value string) inputRating {
	return rc.set("description", value)
}

// DescriptionClassName description class name
func (rc inputRating) DescriptionClassName(value string) inputRating {
	return rc.set("descriptionClassName", value)
}

// Disabled disables the rating
func (rc inputRating) Disabled(value bool) inputRating {
	return rc.set("disabled", value)
}

// DisabledOn disable condition expression
func (rc inputRating) DisabledOn(value string) inputRating {
	return rc.set("disabledOn", value)
}

// EditorSetting editor configuration, ignored at runtime
func (rc inputRating) EditorSetting(value string) inputRating {
	return rc.set("editorSetting", value)
}

// ExtraName extra field name
func (rc inputRating) ExtraName(value string) inputRating {
	return rc.set("extraName", value)
}

// Half allows half-star rating
func (rc inputRating) Half(value bool) inputRating {
	return rc.set("half", value)
}

// Hidden hides the rating
func (rc inputRating) Hidden(value bool) inputRating {
	return rc.set("hidden", value)
}

// HiddenOn hide condition expression
func (rc inputRating) HiddenOn(value string) inputRating {
	return rc.set("hiddenOn", value)
}

// Hint input hint displayed on focus
func (rc inputRating) Hint(value string) inputRating {
	return rc.set("hint", value)
}

// Horizontal horizontal layout configuration
func (rc inputRating) Horizontal(value string) inputRating {
	return rc.set("horizontal", value)
}

// ID unique component ID
func (rc inputRating) ID(value string) inputRating {
	return rc.set("id", value)
}

// InactiveColor unselected star color
func (rc inputRating) InactiveColor(value string) inputRating {
	return rc.set("inactiveColor", value)
}

// InitAutoFill initial auto-fill value
func (rc inputRating) InitAutoFill(value string) inputRating {
	return rc.set("initAutoFill", value)
}

// Inline inline mode for form control
func (rc inputRating) Inline(value bool) inputRating {
	return rc.set("inline", value)
}

// InputClassName input class name
func (rc inputRating) InputClassName(value string) inputRating {
	return rc.set("inputClassName", value)
}

// Label label text
func (rc inputRating) Label(value string) inputRating {
	return rc.set("label", value)
}

// LabelAlign label alignment
func (rc inputRating) LabelAlign(value string) inputRating {
	return rc.set("labelAlign", value)
}

// LabelClassName label class name
func (rc inputRating) LabelClassName(value string) inputRating {
	return rc.set("labelClassName", value)
}

// LabelRemark label remark tooltip
func (rc inputRating) LabelRemark(value string) inputRating {
	return rc.set("labelRemark", value)
}

// LabelWidth custom label width
func (rc inputRating) LabelWidth(value string) inputRating {
	return rc.set("labelWidth", value)
}

// Mode display mode
func (rc inputRating) Mode(value string) inputRating {
	return rc.set("mode", value)
}

// Name field name
func (rc inputRating) Name(value string) inputRating {
	return rc.set("name", value)
}

// OnEvent event action configuration
func (rc inputRating) OnEvent(value any) inputRating {
	return rc.set("onEvent", value)
}

// Placeholder placeholder text
func (rc inputRating) Placeholder(value string) inputRating {
	return rc.set("placeholder", value)
}

// ReadOnly read-only mode
func (rc inputRating) ReadOnly(value bool) inputRating {
	return rc.set("readOnly", value)
}

// ReadOnlyOn read-only condition expression
func (rc inputRating) ReadOnlyOn(value string) inputRating {
	return rc.set("readOnlyOn", value)
}

// Remark remark tooltip
func (rc inputRating) Remark(value string) inputRating {
	return rc.set("remark", value)
}

// Required required field
func (rc inputRating) Required(value bool) inputRating {
	return rc.set("required", value)
}

// Row row configuration
func (rc inputRating) Row(value string) inputRating {
	return rc.set("row", value)
}

// SaveImmediately save immediately on change
func (rc inputRating) SaveImmediately(value bool) inputRating {
	return rc.set("saveImmediately", value)
}

// Size form item size
func (rc inputRating) Size(value string) inputRating {
	return rc.set("size", value)
}

// Static static display mode
func (rc inputRating) Static(value bool) inputRating {
	return rc.set("static", value)
}

// StaticClassName static display class name
func (rc inputRating) StaticClassName(value string) inputRating {
	return rc.set("staticClassName", value)
}

// StaticInputClassName static input value class name
func (rc inputRating) StaticInputClassName(value string) inputRating {
	return rc.set("staticInputClassName", value)
}

// StaticLabelClassName static label class name
func (rc inputRating) StaticLabelClassName(value string) inputRating {
	return rc.set("staticLabelClassName", value)
}

// StaticOn static display condition expression
func (rc inputRating) StaticOn(value string) inputRating {
	return rc.set("staticOn", value)
}

// StaticPlaceholder static display placeholder
func (rc inputRating) StaticPlaceholder(value string) inputRating {
	return rc.set("staticPlaceholder", value)
}

// StaticSchema static schema
func (rc inputRating) StaticSchema(value string) inputRating {
	return rc.set("staticSchema", value)
}

// Style component style
func (rc inputRating) Style(value any) inputRating {
	return rc.set("style", value)
}

// SubmitOnChange submit form on change
func (rc inputRating) SubmitOnChange(value bool) inputRating {
	return rc.set("submitOnChange", value)
}

// TestIdBuilder test ID builder
func (rc inputRating) TestIdBuilder(value string) inputRating {
	return rc.set("testIdBuilder", value)
}

// TextClassName custom text class name
func (rc inputRating) TextClassName(value string) inputRating {
	return rc.set("textClassName", value)
}

// TextPosition text position
func (rc inputRating) TextPosition(value string) inputRating {
	return rc.set("textPosition", value)
}

// Texts selected star text
func (rc inputRating) Texts(value string) inputRating {
	return rc.set("texts", value)
}

// UseMobileUI disable mobile UI styles
func (rc inputRating) UseMobileUI(value bool) inputRating {
	return rc.set("useMobileUI", value)
}

// ValidateApi remote validation API
func (rc inputRating) ValidateApi(value string) inputRating {
	return rc.set("validateApi", value)
}

// ValidateOnChange validate on change
func (rc inputRating) ValidateOnChange(value bool) inputRating {
	return rc.set("validateOnChange", value)
}

// ValidationErrors validation error messages
func (rc inputRating) ValidationErrors(value string) inputRating {
	return rc.set("validationErrors", value)
}

// Validations validation rules
func (rc inputRating) Validations(value string) inputRating {
	return rc.set("validations", value)
}

// Value default value
func (rc inputRating) Value(value any) inputRating {
	return rc.set("value", value)
}

// Visible visibility
func (rc inputRating) Visible(value bool) inputRating {
	return rc.set("visible", value)
}

// VisibleOn visibility condition expression
func (rc inputRating) VisibleOn(value string) inputRating {
	return rc.set("visibleOn", value)
}

// Width width adjustment in table
func (rc inputRating) Width(value string) inputRating {
	return rc.set("width", value)
}
