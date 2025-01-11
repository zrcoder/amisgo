package comp

import "github.com/zrcoder/amisgo/model"

// fieldSet represents a collection of form items
type fieldSet model.Schema

func FieldSet() fieldSet {
	return fieldSet{"type": "fieldset", "titlePosition": "top"}
}

func (f fieldSet) set(key string, value any) fieldSet {
	f[key] = value
	return f
}

// AutoFill sets the autoFill value
func (f fieldSet) AutoFill(value string) fieldSet {
	return f.set("autoFill", value)
}

// Body sets the body content
func (f fieldSet) Body(value ...any) fieldSet {
	return f.set("body", value)
}

// BodyClassName sets the className for the body container
func (f fieldSet) BodyClassName(value string) fieldSet {
	return f.set("bodyClassName", value)
}

// ClassName sets the CSS class name for the container
func (f fieldSet) ClassName(value string) fieldSet {
	return f.set("className", value)
}

// ClearValueOnHidden sets whether to clear the value when the form item is hidden
func (f fieldSet) ClearValueOnHidden(value bool) fieldSet {
	return f.set("clearValueOnHidden", value)
}

// Collapsable sets whether the fieldset is collapsable
func (f fieldSet) Collapsable(value bool) fieldSet {
	return f.set("collapsable", value)
}

// CollapseHeader sets the collapse header
func (f fieldSet) CollapseHeader(value string) fieldSet {
	return f.set("collapseHeader", value)
}

// CollapseTitle sets the collapse title
func (f fieldSet) CollapseTitle(value any) fieldSet {
	return f.set("collapseTitle", value)
}

// Collapsed sets whether the fieldset is collapsed by default
func (f fieldSet) Collapsed(value bool) fieldSet {
	return f.set("collapsed", value)
}

// Desc sets the description
func (f fieldSet) Desc(value string) fieldSet {
	return f.set("desc", value)
}

// Description sets the description content, supports HTML
func (f fieldSet) Description(value string) fieldSet {
	return f.set("description", value)
}

// DescriptionClassName sets the className for the description
func (f fieldSet) DescriptionClassName(value string) fieldSet {
	return f.set("descriptionClassName", value)
}

// Disabled sets whether the fieldset is disabled
func (f fieldSet) Disabled(value bool) fieldSet {
	return f.set("disabled", value)
}

// DisabledOn sets the expression to determine if the fieldset is disabled
func (f fieldSet) DisabledOn(value string) fieldSet {
	return f.set("disabledOn", value)
}

// DivideLine sets whether to show a dividing line
func (f fieldSet) DivideLine(value bool) fieldSet {
	return f.set("divideLine", value)
}

// EditorSetting sets the editor configuration
func (f fieldSet) EditorSetting(value string) fieldSet {
	return f.set("editorSetting", value)
}

// ExpandIcon sets the custom toggle icon
func (f fieldSet) ExpandIcon(value string) fieldSet {
	return f.set("expandIcon", value)
}

// ExtraName sets an additional field name
func (f fieldSet) ExtraName(value string) fieldSet {
	return f.set("extraName", value)
}

// Header sets the header
func (f fieldSet) Header(value string) fieldSet {
	return f.set("header", value)
}

// HeaderPosition sets the header position
func (f fieldSet) HeaderPosition(value string) fieldSet {
	return f.set("headerPosition", value)
}

// HeadingClassName sets the className for the heading
func (f fieldSet) HeadingClassName(value string) fieldSet {
	return f.set("headingClassName", value)
}

// Hidden sets whether the fieldset is hidden
func (f fieldSet) Hidden(value bool) fieldSet {
	return f.set("hidden", value)
}

// HiddenOn sets the expression to determine if the fieldset is hidden
func (f fieldSet) HiddenOn(value string) fieldSet {
	return f.set("hiddenOn", value)
}

// Hint sets the input hint
func (f fieldSet) Hint(value string) fieldSet {
	return f.set("hint", value)
}

// Horizontal sets the horizontal layout
func (f fieldSet) Horizontal(value string) fieldSet {
	return f.set("horizontal", value)
}

// ID sets the unique ID for the component
func (f fieldSet) ID(value string) fieldSet {
	return f.set("id", value)
}

// InitAutoFill sets the initial autoFill value
func (f fieldSet) InitAutoFill(value string) fieldSet {
	return f.set("initAutoFill", value)
}

// Inline sets whether the form control is in inline mode
func (f fieldSet) Inline(value bool) fieldSet {
	return f.set("inline", value)
}

// InputClassName sets the className for the input
func (f fieldSet) InputClassName(value string) fieldSet {
	return f.set("inputClassName", value)
}

// Key sets the key
func (f fieldSet) Key(value string) fieldSet {
	return f.set("key", value)
}

// Label sets the label
func (f fieldSet) Label(value string) fieldSet {
	return f.set("label", value)
}

// LabelAlign sets the label alignment
func (f fieldSet) LabelAlign(value string) fieldSet {
	return f.set("labelAlign", value)
}

// LabelClassName sets the className for the label
func (f fieldSet) LabelClassName(value string) fieldSet {
	return f.set("labelClassName", value)
}

// LabelRemark sets the label remark
func (f fieldSet) LabelRemark(value string) fieldSet {
	return f.set("labelRemark", value)
}

// LabelWidth sets the custom width for the label
func (f fieldSet) LabelWidth(value string) fieldSet {
	return f.set("labelWidth", value)
}

// Mode sets the display mode for the form item
func (f fieldSet) Mode(value string) fieldSet {
	return f.set("mode", value)
}

// MountOnEnter sets whether to load content only when opened
func (f fieldSet) MountOnEnter(value bool) fieldSet {
	return f.set("mountOnEnter", value)
}

// Name sets the field name
func (f fieldSet) Name(value string) fieldSet {
	return f.set("name", value)
}

// OnEvent sets the event action configuration
func (f fieldSet) OnEvent(value any) fieldSet {
	return f.set("onEvent", value)
}

// Placeholder sets the placeholder
func (f fieldSet) Placeholder(value string) fieldSet {
	return f.set("placeholder", value)
}

// ReadOnly sets whether the fieldset is read-only
func (f fieldSet) ReadOnly(value bool) fieldSet {
	return f.set("readOnly", value)
}

// ReadOnlyOn sets the expression to determine if the fieldset is read-only
func (f fieldSet) ReadOnlyOn(value string) fieldSet {
	return f.set("readOnlyOn", value)
}

// Remark sets the remark
func (f fieldSet) Remark(value string) fieldSet {
	return f.set("remark", value)
}

// Required sets whether the fieldset is required
func (f fieldSet) Required(value bool) fieldSet {
	return f.set("required", value)
}

// Row sets the row
func (f fieldSet) Row(value string) fieldSet {
	return f.set("row", value)
}

// SaveImmediately sets whether to save immediately
func (f fieldSet) SaveImmediately(value bool) fieldSet {
	return f.set("saveImmediately", value)
}

// ShowArrow sets whether to show the arrow icon
func (f fieldSet) ShowArrow(value bool) fieldSet {
	return f.set("showArrow", value)
}

// Size sets the size of the control
func (f fieldSet) Size(value string) fieldSet {
	return f.set("size", value)
}

// Static sets whether to display statically
func (f fieldSet) Static(value bool) fieldSet {
	return f.set("static", value)
}

// StaticClassName sets the className for static display
func (f fieldSet) StaticClassName(value string) fieldSet {
	return f.set("staticClassName", value)
}

// StaticInputClassName sets the className for static input display
func (f fieldSet) StaticInputClassName(value string) fieldSet {
	return f.set("staticInputClassName", value)
}

// SubmitOnChange sets whether to submit on change
func (f fieldSet) SubmitOnChange(value bool) fieldSet {
	return f.set("submitOnChange", value)
}

// Title sets the title
func (f fieldSet) Title(value any) fieldSet {
	return f.set("title", value)
}

// TitlePosition sets the title position
func (f fieldSet) TitlePosition(value string) fieldSet {
	return f.set("titlePosition", value)
}

// Visible sets the visibility condition
func (f fieldSet) Visible(value bool) fieldSet {
	return f.set("visible", value)
}

// VisibleOn sets the expression to determine visibility
func (f fieldSet) VisibleOn(value string) fieldSet {
	return f.set("visibleOn", value)
}
