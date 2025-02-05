package comp

import "github.com/zrcoder/amisgo/model"

// Audio represents an Audio renderer.
type Audio model.Schema

func NewAudio() Audio {
	return Audio{"type": "audio"}
}

func (a Audio) set(key string, value any) Audio {
	a[key] = value
	return a
}

// AutoPlay sets whether the audio should play automatically.
func (a Audio) AutoPlay(value bool) Audio {
	return a.set("autoPlay", value)
}

// ClassName sets the CSS class name for the container.
func (a Audio) ClassName(value string) Audio {
	return a.set("className", value)
}

// Controls sets the controls configuration.
func (a Audio) Controls(value string) Audio {
	return a.set("controls", value)
}

// Disabled sets whether the audio is disabled.
func (a Audio) Disabled(value bool) Audio {
	return a.set("disabled", value)
}

// DisabledOn sets the expression to determine whether the audio is disabled.
func (a Audio) DisabledOn(value string) Audio {
	return a.set("disabledOn", value)
}

// EditorSetting sets the editor configuration.
func (a Audio) EditorSetting(value string) Audio {
	return a.set("editorSetting", value)
}

// Hidden sets whether the audio is hidden.
func (a Audio) Hidden(value bool) Audio {
	return a.set("hidden", value)
}

// HiddenOn sets the expression to determine whether the audio is hidden.
func (a Audio) HiddenOn(value string) Audio {
	return a.set("hiddenOn", value)
}

// ID sets the unique ID for the component.
func (a Audio) ID(value string) Audio {
	return a.set("id", value)
}

// Inline sets whether the audio is displayed inline.
func (a Audio) Inline(value bool) Audio {
	return a.set("inline", value)
}

// Loop sets whether the audio should loop.
func (a Audio) Loop(value bool) Audio {
	return a.set("loop", value)
}

// OnEvent sets the event action configuration.
func (a Audio) OnEvent(value any) Audio {
	return a.set("onEvent", value)
}

// Rates sets the available playback speeds.
func (a Audio) Rates(value string) Audio {
	return a.set("rates", value)
}

// Src sets the audio source URL.
func (a Audio) Src(value string) Audio {
	return a.set("src", value)
}

// Static sets whether the audio is displayed statically.
func (a Audio) Static(value bool) Audio {
	return a.set("static", value)
}

// StaticClassName sets the CSS class name for static display.
func (a Audio) StaticClassName(value string) Audio {
	return a.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for the static input.
func (a Audio) StaticInputClassName(value string) Audio {
	return a.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for the static label.
func (a Audio) StaticLabelClassName(value string) Audio {
	return a.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine whether the audio is displayed statically.
func (a Audio) StaticOn(value string) Audio {
	return a.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display.
func (a Audio) StaticPlaceholder(value string) Audio {
	return a.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display.
func (a Audio) StaticSchema(value string) Audio {
	return a.set("staticSchema", value)
}

// Style sets the style for the component.
func (a Audio) Style(value any) Audio {
	return a.set("style", value)
}

// TestIdBuilder sets the test ID builder.
func (a Audio) TestIdBuilder(value string) Audio {
	return a.set("testIdBuilder", value)
}

// Testid sets the test ID.
func (a Audio) Testid(value string) Audio {
	return a.set("testid", value)
}

// UseMobileUI sets whether to use mobile UI.
func (a Audio) UseMobileUI(value bool) Audio {
	return a.set("useMobileUI", value)
}

// Visible sets whether the audio is visible.
func (a Audio) Visible(value bool) Audio {
	return a.set("visible", value)
}

// VisibleOn sets the expression to determine whether the audio is visible.
func (a Audio) VisibleOn(value string) Audio {
	return a.set("visibleOn", value)
}
