package comp

import "github.com/zrcoder/amisgo/model"

// inputQuarter represents a quarter selection control
type inputQuarter model.Schema

func InputQuarter() inputQuarter {
	return inputQuarter{"type": "input-quarter"}
}

// AutoFill sets auto-fill value
func (q inputQuarter) AutoFill(value string) inputQuarter {
	return q.set("autoFill", value)
}

// BorderMode sets border mode: full, half, or none
func (q inputQuarter) BorderMode(value string) inputQuarter {
	return q.set("borderMode", value)
}

// ClassName sets container CSS class name
func (q inputQuarter) ClassName(value string) inputQuarter {
	return q.set("className", value)
}

// ClearValueOnHidden clears value when hidden
func (q inputQuarter) ClearValueOnHidden(value bool) inputQuarter {
	return q.set("clearValueOnHidden", value)
}

// Clearable shows clear button
func (q inputQuarter) Clearable(value bool) inputQuarter {
	return q.set("clearable", value)
}

// Desc sets description
func (q inputQuarter) Desc(value string) inputQuarter {
	return q.set("desc", value)
}

// Description sets HTML description
func (q inputQuarter) Description(value string) inputQuarter {
	return q.set("description", value)
}

// DescriptionClassName sets description CSS class name
func (q inputQuarter) DescriptionClassName(value string) inputQuarter {
	return q.set("descriptionClassName", value)
}

// Disabled disables the control
func (q inputQuarter) Disabled(value bool) inputQuarter {
	return q.set("disabled", value)
}

// DisabledDate sets function to disable specific dates
func (q inputQuarter) DisabledDate(value string) inputQuarter {
	return q.set("disabledDate", value)
}

// DisabledOn sets expression to disable the control
func (q inputQuarter) DisabledOn(value string) inputQuarter {
	return q.set("disabledOn", value)
}

// DisplayFormat sets date display format
func (q inputQuarter) DisplayFormat(value string) inputQuarter {
	return q.set("displayFormat", value)
}

// EditorSetting sets editor configuration
func (q inputQuarter) EditorSetting(value string) inputQuarter {
	return q.set("editorSetting", value)
}

// Emebed sets inline mode
func (q inputQuarter) Emebed(value bool) inputQuarter {
	return q.set("emebed", value)
}

// ExtraName sets extra field name for range components
func (q inputQuarter) ExtraName(value string) inputQuarter {
	return q.set("extraName", value)
}

// Format sets month storage format
func (q inputQuarter) Format(value string) inputQuarter {
	return q.set("format", value)
}

// Hidden hides the control
func (q inputQuarter) Hidden(value bool) inputQuarter {
	return q.set("hidden", value)
}

// HiddenOn sets expression to hide the control
func (q inputQuarter) HiddenOn(value string) inputQuarter {
	return q.set("hiddenOn", value)
}

// Hint sets input hint
func (q inputQuarter) Hint(value string) inputQuarter {
	return q.set("hint", value)
}

// Horizontal sets horizontal layout configuration
func (q inputQuarter) Horizontal(value string) inputQuarter {
	return q.set("horizontal", value)
}

// ID sets unique component ID
func (q inputQuarter) ID(value string) inputQuarter {
	return q.set("id", value)
}

// InitAutoFill sets initial auto-fill value
func (q inputQuarter) InitAutoFill(value string) inputQuarter {
	return q.set("initAutoFill", value)
}

// Inline sets inline mode
func (q inputQuarter) Inline(value bool) inputQuarter {
	return q.set("inline", value)
}

// InputClassName sets input CSS class name
func (q inputQuarter) InputClassName(value string) inputQuarter {
	return q.set("inputClassName", value)
}

// InputFormat sets month display format
func (q inputQuarter) InputFormat(value string) inputQuarter {
	return q.set("inputFormat", value)
}

// Label sets label text
func (q inputQuarter) Label(value string) inputQuarter {
	return q.set("label", value)
}

// LabelAlign sets label alignment: right, left, top, or inherit
func (q inputQuarter) LabelAlign(value string) inputQuarter {
	return q.set("labelAlign", value)
}

// LabelClassName sets label CSS class name
func (q inputQuarter) LabelClassName(value string) inputQuarter {
	return q.set("labelClassName", value)
}

// LabelRemark sets label remark
func (q inputQuarter) LabelRemark(value string) inputQuarter {
	return q.set("labelRemark", value)
}

// LabelWidth sets custom label width in px
func (q inputQuarter) LabelWidth(value string) inputQuarter {
	return q.set("labelWidth", value)
}

// Mode sets display mode: normal, inline, or horizontal
func (q inputQuarter) Mode(value string) inputQuarter {
	return q.set("mode", value)
}

// Name sets field name for form submission
func (q inputQuarter) Name(value string) inputQuarter {
	return q.set("name", value)
}

// OnEvent sets event actions
func (q inputQuarter) OnEvent(value any) inputQuarter {
	return q.set("onEvent", value)
}

// Placeholder sets placeholder text
func (q inputQuarter) Placeholder(value string) inputQuarter {
	return q.set("placeholder", value)
}

// ReadOnly sets read-only mode
func (q inputQuarter) ReadOnly(value bool) inputQuarter {
	return q.set("readOnly", value)
}

// ReadOnlyOn sets expression for read-only mode
func (q inputQuarter) ReadOnlyOn(value string) inputQuarter {
	return q.set("readOnlyOn", value)
}

// Remark sets remark text
func (q inputQuarter) Remark(value string) inputQuarter {
	return q.set("remark", value)
}

// Required sets required field
func (q inputQuarter) Required(value bool) inputQuarter {
	return q.set("required", value)
}

// Row sets row value
func (q inputQuarter) Row(value string) inputQuarter {
	return q.set("row", value)
}

// SaveImmediately sets immediate save for TableColumn
func (q inputQuarter) SaveImmediately(value bool) inputQuarter {
	return q.set("saveImmediately", value)
}

// Shortcuts sets date shortcuts
func (q inputQuarter) Shortcuts(value string) inputQuarter {
	return q.set("shortcuts", value)
}

// Size sets form item size: xs, sm, md, lg, or full
func (q inputQuarter) Size(value string) inputQuarter {
	return q.set("size", value)
}

// Static sets static display mode
func (q inputQuarter) Static(value bool) inputQuarter {
	return q.set("static", value)
}

// StaticClassName sets static display CSS class name
func (q inputQuarter) StaticClassName(value string) inputQuarter {
	return q.set("staticClassName", value)
}

// StaticInputClassName sets static input CSS class name
func (q inputQuarter) StaticInputClassName(value string) inputQuarter {
	return q.set("staticInputClassName", value)
}

// StaticLabelClassName sets static label CSS class name
func (q inputQuarter) StaticLabelClassName(value string) inputQuarter {
	return q.set("staticLabelClassName", value)
}

// Template sets template value
func (q inputQuarter) Template(value string) inputQuarter {
	return q.set("template", value)
}

// TimeFormat sets time display format
func (q inputQuarter) TimeFormat(value string) inputQuarter {
	return q.set("timeFormat", value)
}

// Validation sets validation message
func (q inputQuarter) Validation(value string) inputQuarter {
	return q.set("validation", value)
}

// Value sets default value
func (q inputQuarter) Value(value string) inputQuarter {
	return q.set("value", value)
}

// VerticalAlign sets vertical alignment: top, middle, or bottom
func (q inputQuarter) VerticalAlign(value string) inputQuarter {
	return q.set("verticalAlign", value)
}

// VisibleOn sets expression for visibility
func (q inputQuarter) VisibleOn(value string) inputQuarter {
	return q.set("visibleOn", value)
}

// Width sets width value
func (q inputQuarter) Width(value string) inputQuarter {
	return q.set("width", value)
}

// set sets a property value
func (q inputQuarter) set(key string, value any) inputQuarter {
	// This method should set the property `key` with the given `value` on the `QuarterControl` instance.
	// Assuming we have a way to store and manage these properties.
	return q
}
