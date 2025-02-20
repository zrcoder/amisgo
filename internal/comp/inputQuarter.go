package comp

import "github.com/zrcoder/amisgo/schema"

// InputQuarter represents a quarter selection control
type InputQuarter schema.Schema

func NewInputQuarter() InputQuarter {
	return InputQuarter{"type": "input-quarter"}
}

// AutoFill sets auto-fill value
func (q InputQuarter) AutoFill(value string) InputQuarter {
	return q.set("autoFill", value)
}

// BorderMode sets border mode: full, half, or none
func (q InputQuarter) BorderMode(value string) InputQuarter {
	return q.set("borderMode", value)
}

// ClassName sets container CSS class name
func (q InputQuarter) ClassName(value string) InputQuarter {
	return q.set("className", value)
}

// ClearValueOnHidden clears value when hidden
func (q InputQuarter) ClearValueOnHidden(value bool) InputQuarter {
	return q.set("clearValueOnHidden", value)
}

// Clearable shows clear button
func (q InputQuarter) Clearable(value bool) InputQuarter {
	return q.set("clearable", value)
}

// Desc sets description
func (q InputQuarter) Desc(value string) InputQuarter {
	return q.set("desc", value)
}

// Description sets HTML description
func (q InputQuarter) Description(value string) InputQuarter {
	return q.set("description", value)
}

// DescriptionClassName sets description CSS class name
func (q InputQuarter) DescriptionClassName(value string) InputQuarter {
	return q.set("descriptionClassName", value)
}

// Disabled disables the control
func (q InputQuarter) Disabled(value bool) InputQuarter {
	return q.set("disabled", value)
}

// DisabledDate sets function to disable specific dates
func (q InputQuarter) DisabledDate(value string) InputQuarter {
	return q.set("disabledDate", value)
}

// DisabledOn sets expression to disable the control
func (q InputQuarter) DisabledOn(value string) InputQuarter {
	return q.set("disabledOn", value)
}

// DisplayFormat sets date display format
func (q InputQuarter) DisplayFormat(value string) InputQuarter {
	return q.set("displayFormat", value)
}

// EditorSetting sets editor configuration
func (q InputQuarter) EditorSetting(value string) InputQuarter {
	return q.set("editorSetting", value)
}

// Emebed sets inline mode
func (q InputQuarter) Emebed(value bool) InputQuarter {
	return q.set("emebed", value)
}

// ExtraName sets extra field name for range components
func (q InputQuarter) ExtraName(value string) InputQuarter {
	return q.set("extraName", value)
}

// Format sets month storage format
func (q InputQuarter) Format(value string) InputQuarter {
	return q.set("format", value)
}

// Hidden hides the control
func (q InputQuarter) Hidden(value bool) InputQuarter {
	return q.set("hidden", value)
}

// HiddenOn sets expression to hide the control
func (q InputQuarter) HiddenOn(value string) InputQuarter {
	return q.set("hiddenOn", value)
}

// Hint sets input hint
func (q InputQuarter) Hint(value string) InputQuarter {
	return q.set("hint", value)
}

// Horizontal sets horizontal layout configuration
func (q InputQuarter) Horizontal(value string) InputQuarter {
	return q.set("horizontal", value)
}

// ID sets unique component ID
func (q InputQuarter) ID(value string) InputQuarter {
	return q.set("id", value)
}

// InitAutoFill sets initial auto-fill value
func (q InputQuarter) InitAutoFill(value string) InputQuarter {
	return q.set("initAutoFill", value)
}

// Inline sets inline mode
func (q InputQuarter) Inline(value bool) InputQuarter {
	return q.set("inline", value)
}

// InputClassName sets input CSS class name
func (q InputQuarter) InputClassName(value string) InputQuarter {
	return q.set("inputClassName", value)
}

// InputFormat sets month display format
func (q InputQuarter) InputFormat(value string) InputQuarter {
	return q.set("inputFormat", value)
}

// Label sets label text
func (q InputQuarter) Label(value string) InputQuarter {
	return q.set("label", value)
}

// LabelAlign sets label alignment: right, left, top, or inherit
func (q InputQuarter) LabelAlign(value string) InputQuarter {
	return q.set("labelAlign", value)
}

// LabelClassName sets label CSS class name
func (q InputQuarter) LabelClassName(value string) InputQuarter {
	return q.set("labelClassName", value)
}

// LabelRemark sets label remark
func (q InputQuarter) LabelRemark(value string) InputQuarter {
	return q.set("labelRemark", value)
}

// LabelWidth sets custom label width in px
func (q InputQuarter) LabelWidth(value string) InputQuarter {
	return q.set("labelWidth", value)
}

// Mode sets display mode: normal, inline, or horizontal
func (q InputQuarter) Mode(value string) InputQuarter {
	return q.set("mode", value)
}

// Name sets field name for form submission
func (q InputQuarter) Name(value string) InputQuarter {
	return q.set("name", value)
}

// OnEvent sets event actions
func (q InputQuarter) OnEvent(value Event) InputQuarter {
	return q.set("onEvent", value)
}

// Placeholder sets placeholder text
func (q InputQuarter) Placeholder(value string) InputQuarter {
	return q.set("placeholder", value)
}

// ReadOnly sets read-only mode
func (q InputQuarter) ReadOnly(value bool) InputQuarter {
	return q.set("readOnly", value)
}

// ReadOnlyOn sets expression for read-only mode
func (q InputQuarter) ReadOnlyOn(value string) InputQuarter {
	return q.set("readOnlyOn", value)
}

// Remark sets remark text
func (q InputQuarter) Remark(value string) InputQuarter {
	return q.set("remark", value)
}

// Required sets required field
func (q InputQuarter) Required(value bool) InputQuarter {
	return q.set("required", value)
}

// Row sets row value
func (q InputQuarter) Row(value string) InputQuarter {
	return q.set("row", value)
}

// SaveImmediately sets immediate save for TableColumn
func (q InputQuarter) SaveImmediately(value bool) InputQuarter {
	return q.set("saveImmediately", value)
}

// Shortcuts sets date shortcuts
func (q InputQuarter) Shortcuts(value string) InputQuarter {
	return q.set("shortcuts", value)
}

// Size sets form item size: xs, sm, md, lg, or full
func (q InputQuarter) Size(value string) InputQuarter {
	return q.set("size", value)
}

// Static sets static display mode
func (q InputQuarter) Static(value bool) InputQuarter {
	return q.set("static", value)
}

// StaticClassName sets static display CSS class name
func (q InputQuarter) StaticClassName(value string) InputQuarter {
	return q.set("staticClassName", value)
}

// StaticInputClassName sets static input CSS class name
func (q InputQuarter) StaticInputClassName(value string) InputQuarter {
	return q.set("staticInputClassName", value)
}

// StaticLabelClassName sets static label CSS class name
func (q InputQuarter) StaticLabelClassName(value string) InputQuarter {
	return q.set("staticLabelClassName", value)
}

// Template sets template value
func (q InputQuarter) Template(value string) InputQuarter {
	return q.set("template", value)
}

// TimeFormat sets time display format
func (q InputQuarter) TimeFormat(value string) InputQuarter {
	return q.set("timeFormat", value)
}

// Validation sets validation message
func (q InputQuarter) Validation(value string) InputQuarter {
	return q.set("validation", value)
}

// Value sets default value
func (q InputQuarter) Value(value string) InputQuarter {
	return q.set("value", value)
}

// VerticalAlign sets vertical alignment: top, middle, or bottom
func (q InputQuarter) VerticalAlign(value string) InputQuarter {
	return q.set("verticalAlign", value)
}

// VisibleOn sets expression for visibility
func (q InputQuarter) VisibleOn(value string) InputQuarter {
	return q.set("visibleOn", value)
}

// Width sets width value
func (q InputQuarter) Width(value string) InputQuarter {
	return q.set("width", value)
}

// set sets a property value
func (q InputQuarter) set(key string, value any) InputQuarter {
	// This method should set the property `key` with the given `value` on the `QuarterControl` instance.
	// Assuming we have a way to store and manage these properties.
	return q
}
