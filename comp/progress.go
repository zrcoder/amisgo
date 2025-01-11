package comp

import "github.com/zrcoder/amisgo/model"

// progress

type progress model.Schema

// Progress creates a new Progress instance
func Progress() progress {
	return progress{"type": "progress", "mode": "line"}
}

func (p progress) set(key string, value any) progress {
	p[key] = value
	return p
}

// Animate sets animation visibility
func (p progress) Animate(value bool) progress {
	return p.set("animate", value)
}

// ClassName sets the container CSS class name
func (p progress) ClassName(value string) progress {
	return p.set("className", value)
}

// Disabled sets the disabled state
func (p progress) Disabled(value bool) progress {
	return p.set("disabled", value)
}

// DisabledOn sets the disabled expression
func (p progress) DisabledOn(value string) progress {
	return p.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (p progress) EditorSetting(value string) progress {
	return p.set("editorSetting", value)
}

// GapDegree sets the gap angle for the dashboard progress bar
func (p progress) GapDegree(value string) progress {
	return p.set("gapDegree", value)
}

// GapPosition sets the gap position for the dashboard progress bar
func (p progress) GapPosition(value string) progress {
	return p.set("gapPosition", value)
}

// Hidden sets the hidden state
func (p progress) Hidden(value bool) progress {
	return p.set("hidden", value)
}

// HiddenOn sets the hidden expression
func (p progress) HiddenOn(value string) progress {
	return p.set("hiddenOn", value)
}

// ID sets the unique component ID
func (p progress) ID(value string) progress {
	return p.set("id", value)
}

// Map sets the value segment configuration
func (p progress) Map(value string) progress {
	return p.set("map", value)
}

// Mode sets the progress bar type
func (p progress) Mode(value string) progress {
	return p.set("mode", value)
}

// Name sets the associated field name
func (p progress) Name(value string) progress {
	return p.set("name", value)
}

// OnEvent sets the event action configuration
func (p progress) OnEvent(value any) progress {
	return p.set("onEvent", value)
}

// Placeholder sets the placeholder
func (p progress) Placeholder(value string) progress {
	return p.set("placeholder", value)
}

// ProgressClassName sets the progress bar CSS class name
func (p progress) ProgressClassName(value string) progress {
	return p.set("progressClassName", value)
}

// ShowLabel sets the visibility of the value
func (p progress) ShowLabel(value bool) progress {
	return p.set("showLabel", value)
}

// ShowThresholdText sets the visibility of the threshold value
func (p progress) ShowThresholdText(value bool) progress {
	return p.set("showThresholdText", value)
}

// Static sets the static display state
func (p progress) Static(value bool) progress {
	return p.set("static", value)
}

// StaticClassName sets the static display form item class name
func (p progress) StaticClassName(value string) progress {
	return p.set("staticClassName", value)
}

// StaticInputClassName sets the static display form item value class name
func (p progress) StaticInputClassName(value string) progress {
	return p.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static display form item label class name
func (p progress) StaticLabelClassName(value string) progress {
	return p.set("staticLabelClassName", value)
}

// StaticOn sets the static display expression
func (p progress) StaticOn(value string) progress {
	return p.set("staticOn", value)
}

// StaticPlaceholder sets the static display placeholder for empty values
func (p progress) StaticPlaceholder(value string) progress {
	return p.set("staticPlaceholder", value)
}

// StaticSchema sets the static display schema
func (p progress) StaticSchema(value string) progress {
	return p.set("staticSchema", value)
}

// Stripe sets the visibility of the background stripe
func (p progress) Stripe(value bool) progress {
	return p.set("stripe", value)
}

// StrokeWidth sets the width of the progress bar line
func (p progress) StrokeWidth(value string) progress {
	return p.set("strokeWidth", value)
}

// Style sets the component style
func (p progress) Style(value any) progress {
	return p.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (p progress) TestIdBuilder(value string) progress {
	return p.set("testIdBuilder", value)
}

// Testid sets the test ID
func (p progress) Testid(value string) progress {
	return p.set("testid", value)
}

// Threshold sets the threshold value
func (p progress) Threshold(value string) progress {
	return p.set("threshold", value)
}

// UseMobileUI sets the mobile UI usage
func (p progress) UseMobileUI(value bool) progress {
	return p.set("useMobileUI", value)
}

// Value sets the progress value
func (p progress) Value(value string) progress {
	return p.set("value", value)
}

// ValueTpl sets the content template function
func (p progress) ValueTpl(value string) progress {
	return p.set("valueTpl", value)
}

// Visible sets the visibility
func (p progress) Visible(value bool) progress {
	return p.set("visible", value)
}

// VisibleOn sets the visibility expression
func (p progress) VisibleOn(value string) progress {
	return p.set("visibleOn", value)
}
