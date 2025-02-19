package comp

import "github.com/zrcoder/amisgo/schema"

// Progress
type Progress schema.Schema

func NewProgress() Progress {
	return Progress{"type": "progress", "mode": "line"}
}

func (p Progress) set(key string, value any) Progress {
	p[key] = value
	return p
}

// Animate sets animation visibility
func (p Progress) Animate(value bool) Progress {
	return p.set("animate", value)
}

// ClassName sets the container CSS class name
func (p Progress) ClassName(value string) Progress {
	return p.set("className", value)
}

// Disabled sets the disabled state
func (p Progress) Disabled(value bool) Progress {
	return p.set("disabled", value)
}

// DisabledOn sets the disabled expression
func (p Progress) DisabledOn(value string) Progress {
	return p.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (p Progress) EditorSetting(value string) Progress {
	return p.set("editorSetting", value)
}

// GapDegree sets the gap angle for the dashboard progress bar
func (p Progress) GapDegree(value string) Progress {
	return p.set("gapDegree", value)
}

// GapPosition sets the gap position for the dashboard progress bar
func (p Progress) GapPosition(value string) Progress {
	return p.set("gapPosition", value)
}

// Hidden sets the hidden state
func (p Progress) Hidden(value bool) Progress {
	return p.set("hidden", value)
}

// HiddenOn sets the hidden expression
func (p Progress) HiddenOn(value string) Progress {
	return p.set("hiddenOn", value)
}

// ID sets the unique component ID
func (p Progress) ID(value string) Progress {
	return p.set("id", value)
}

// Map sets the value segment configuration
func (p Progress) Map(value string) Progress {
	return p.set("map", value)
}

// Mode sets the progress bar type
func (p Progress) Mode(value string) Progress {
	return p.set("mode", value)
}

// Name sets the associated field name
func (p Progress) Name(value string) Progress {
	return p.set("name", value)
}

// OnEvent sets the event action configuration
func (p Progress) OnEvent(value any) Progress {
	return p.set("onEvent", value)
}

// Placeholder sets the placeholder
func (p Progress) Placeholder(value string) Progress {
	return p.set("placeholder", value)
}

// ProgressClassName sets the progress bar CSS class name
func (p Progress) ProgressClassName(value string) Progress {
	return p.set("progressClassName", value)
}

// ShowLabel sets the visibility of the value
func (p Progress) ShowLabel(value bool) Progress {
	return p.set("showLabel", value)
}

// ShowThresholdText sets the visibility of the threshold value
func (p Progress) ShowThresholdText(value bool) Progress {
	return p.set("showThresholdText", value)
}

// Static sets the static display state
func (p Progress) Static(value bool) Progress {
	return p.set("static", value)
}

// StaticClassName sets the static display form item class name
func (p Progress) StaticClassName(value string) Progress {
	return p.set("staticClassName", value)
}

// StaticInputClassName sets the static display form item value class name
func (p Progress) StaticInputClassName(value string) Progress {
	return p.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static display form item label class name
func (p Progress) StaticLabelClassName(value string) Progress {
	return p.set("staticLabelClassName", value)
}

// StaticOn sets the static display expression
func (p Progress) StaticOn(value string) Progress {
	return p.set("staticOn", value)
}

// StaticPlaceholder sets the static display placeholder for empty values
func (p Progress) StaticPlaceholder(value string) Progress {
	return p.set("staticPlaceholder", value)
}

// StaticSchema sets the static display schema.Schema
func (p Progress) StaticSchema(value string) Progress {
	return p.set("staticSchema", value)
}

// Stripe sets the visibility of the background stripe
func (p Progress) Stripe(value bool) Progress {
	return p.set("stripe", value)
}

// StrokeWidth sets the width of the progress bar line
func (p Progress) StrokeWidth(value string) Progress {
	return p.set("strokeWidth", value)
}

// Style sets the component style
func (p Progress) Style(value any) Progress {
	return p.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (p Progress) TestIdBuilder(value string) Progress {
	return p.set("testIdBuilder", value)
}

// Testid sets the test ID
func (p Progress) Testid(value string) Progress {
	return p.set("testid", value)
}

// Threshold sets the threshold value
func (p Progress) Threshold(value string) Progress {
	return p.set("threshold", value)
}

// UseMobileUI sets the mobile UI usage
func (p Progress) UseMobileUI(value bool) Progress {
	return p.set("useMobileUI", value)
}

// Value sets the progress value
func (p Progress) Value(value string) Progress {
	return p.set("value", value)
}

// ValueTpl sets the content template function
func (p Progress) ValueTpl(value string) Progress {
	return p.set("valueTpl", value)
}

// Visible sets the visibility
func (p Progress) Visible(value bool) Progress {
	return p.set("visible", value)
}

// VisibleOn sets the visibility expression
func (p Progress) VisibleOn(value string) Progress {
	return p.set("visibleOn", value)
}
