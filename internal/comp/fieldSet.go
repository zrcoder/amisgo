package comp

import "github.com/zrcoder/amisgo/schema"

// FieldSet represents a collection of form items
type FieldSet schema.Schema

func NewFieldSet() FieldSet {
	return FieldSet{"type": "fieldset", "titlePosition": "top"}
}

func (f FieldSet) set(key string, value any) FieldSet {
	f[key] = value
	return f
}

// AutoFill sets the autoFill value
func (f FieldSet) AutoFill(value string) FieldSet {
	return f.set("autoFill", value)
}

// Body sets the body content
func (f FieldSet) Body(value ...any) FieldSet {
	return f.set("body", value)
}

// BodyClassName sets the className for the body container
func (f FieldSet) BodyClassName(value string) FieldSet {
	return f.set("bodyClassName", value)
}

// ClassName sets the CSS class name for the container
func (f FieldSet) ClassName(value string) FieldSet {
	return f.set("className", value)
}

// ClearValueOnHidden sets whether to clear the value when the form item is hidden
func (f FieldSet) ClearValueOnHidden(value bool) FieldSet {
	return f.set("clearValueOnHidden", value)
}

// Collapsable sets whether the fieldset is collapsable
func (f FieldSet) Collapsable(value bool) FieldSet {
	return f.set("collapsable", value)
}

// CollapseHeader sets the collapse header
func (f FieldSet) CollapseHeader(value string) FieldSet {
	return f.set("collapseHeader", value)
}

// CollapseTitle sets the collapse title
func (f FieldSet) CollapseTitle(value any) FieldSet {
	return f.set("collapseTitle", value)
}

// Collapsed sets whether the fieldset is collapsed by default
func (f FieldSet) Collapsed(value bool) FieldSet {
	return f.set("collapsed", value)
}

// Desc sets the description
func (f FieldSet) Desc(value string) FieldSet {
	return f.set("desc", value)
}

// Description sets the description content, supports HTML
func (f FieldSet) Description(value string) FieldSet {
	return f.set("description", value)
}

// DescriptionClassName sets the className for the description
func (f FieldSet) DescriptionClassName(value string) FieldSet {
	return f.set("descriptionClassName", value)
}

// Disabled sets whether the fieldset is disabled
func (f FieldSet) Disabled(value bool) FieldSet {
	return f.set("disabled", value)
}

// DisabledOn sets the expression to determine if the fieldset is disabled
func (f FieldSet) DisabledOn(value string) FieldSet {
	return f.set("disabledOn", value)
}

// DivideLine sets whether to show a dividing line
func (f FieldSet) DivideLine(value bool) FieldSet {
	return f.set("divideLine", value)
}

// EditorSetting sets the editor configuration
func (f FieldSet) EditorSetting(value string) FieldSet {
	return f.set("editorSetting", value)
}

// ExpandIcon sets the custom toggle icon
func (f FieldSet) ExpandIcon(value string) FieldSet {
	return f.set("expandIcon", value)
}

// ExtraName sets an additional field name
func (f FieldSet) ExtraName(value string) FieldSet {
	return f.set("extraName", value)
}

// Header sets the header
func (f FieldSet) Header(value string) FieldSet {
	return f.set("header", value)
}

// HeaderPosition sets the header position
func (f FieldSet) HeaderPosition(value string) FieldSet {
	return f.set("headerPosition", value)
}

// HeadingClassName sets the className for the heading
func (f FieldSet) HeadingClassName(value string) FieldSet {
	return f.set("headingClassName", value)
}

// Hidden sets whether the fieldset is hidden
func (f FieldSet) Hidden(value bool) FieldSet {
	return f.set("hidden", value)
}

// HiddenOn sets the expression to determine if the fieldset is hidden
func (f FieldSet) HiddenOn(value string) FieldSet {
	return f.set("hiddenOn", value)
}

// Hint sets the input hint
func (f FieldSet) Hint(value string) FieldSet {
	return f.set("hint", value)
}

// Horizontal sets the horizontal layout
func (f FieldSet) Horizontal(value string) FieldSet {
	return f.set("horizontal", value)
}

// ID sets the unique ID for the component
func (f FieldSet) ID(value string) FieldSet {
	return f.set("id", value)
}

// InitAutoFill sets the initial autoFill value
func (f FieldSet) InitAutoFill(value string) FieldSet {
	return f.set("initAutoFill", value)
}

// Inline sets whether the form control is in inline mode
func (f FieldSet) Inline(value bool) FieldSet {
	return f.set("inline", value)
}

// InputClassName sets the className for the input
func (f FieldSet) InputClassName(value string) FieldSet {
	return f.set("inputClassName", value)
}

// Key sets the key
func (f FieldSet) Key(value string) FieldSet {
	return f.set("key", value)
}

// Label sets the label
func (f FieldSet) Label(value string) FieldSet {
	return f.set("label", value)
}

// LabelAlign sets the label alignment
func (f FieldSet) LabelAlign(value string) FieldSet {
	return f.set("labelAlign", value)
}

// LabelClassName sets the className for the label
func (f FieldSet) LabelClassName(value string) FieldSet {
	return f.set("labelClassName", value)
}

// LabelRemark sets the label remark
func (f FieldSet) LabelRemark(value string) FieldSet {
	return f.set("labelRemark", value)
}

// LabelWidth sets the custom width for the label
func (f FieldSet) LabelWidth(value string) FieldSet {
	return f.set("labelWidth", value)
}

// Mode sets the display mode for the form item
func (f FieldSet) Mode(value string) FieldSet {
	return f.set("mode", value)
}

// MountOnEnter sets whether to load content only when opened
func (f FieldSet) MountOnEnter(value bool) FieldSet {
	return f.set("mountOnEnter", value)
}

// Name sets the field name
func (f FieldSet) Name(value string) FieldSet {
	return f.set("name", value)
}

// OnEvent sets the event action configuration
func (f FieldSet) OnEvent(value Event) FieldSet {
	return f.set("onEvent", value)
}

// Placeholder sets the placeholder
func (f FieldSet) Placeholder(value string) FieldSet {
	return f.set("placeholder", value)
}

// ReadOnly sets whether the fieldset is read-only
func (f FieldSet) ReadOnly(value bool) FieldSet {
	return f.set("readOnly", value)
}

// ReadOnlyOn sets the expression to determine if the fieldset is read-only
func (f FieldSet) ReadOnlyOn(value string) FieldSet {
	return f.set("readOnlyOn", value)
}

// Remark sets the remark
func (f FieldSet) Remark(value string) FieldSet {
	return f.set("remark", value)
}

// Required sets whether the fieldset is required
func (f FieldSet) Required(value bool) FieldSet {
	return f.set("required", value)
}

// Row sets the row
func (f FieldSet) Row(value string) FieldSet {
	return f.set("row", value)
}

// SaveImmediately sets whether to save immediately
func (f FieldSet) SaveImmediately(value bool) FieldSet {
	return f.set("saveImmediately", value)
}

// ShowArrow sets whether to show the arrow icon
func (f FieldSet) ShowArrow(value bool) FieldSet {
	return f.set("showArrow", value)
}

// Size sets the size of the control
func (f FieldSet) Size(value string) FieldSet {
	return f.set("size", value)
}

// Static sets whether to display statically
func (f FieldSet) Static(value bool) FieldSet {
	return f.set("static", value)
}

// StaticClassName sets the className for static display
func (f FieldSet) StaticClassName(value string) FieldSet {
	return f.set("staticClassName", value)
}

// StaticInputClassName sets the className for static input display
func (f FieldSet) StaticInputClassName(value string) FieldSet {
	return f.set("staticInputClassName", value)
}

// SubmitOnChange sets whether to submit on change
func (f FieldSet) SubmitOnChange(value bool) FieldSet {
	return f.set("submitOnChange", value)
}

// Title sets the title
func (f FieldSet) Title(value any) FieldSet {
	return f.set("title", value)
}

// TitlePosition sets the title position
func (f FieldSet) TitlePosition(value string) FieldSet {
	return f.set("titlePosition", value)
}

// Visible sets the visibility condition
func (f FieldSet) Visible(value bool) FieldSet {
	return f.set("visible", value)
}

// VisibleOn sets the expression to determine visibility
func (f FieldSet) VisibleOn(value string) FieldSet {
	return f.set("visibleOn", value)
}
