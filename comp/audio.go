package comp

// audio represents an audio renderer.
type audio Schema

// Audio creates a new audio renderer.
func Audio() audio {
	return make(audio).set("type", "audio")
}

func (a audio) set(key string, value any) audio {
	a[key] = value
	return a
}

// AutoPlay sets whether the audio should play automatically.
func (a audio) AutoPlay(value bool) audio {
	return a.set("autoPlay", value)
}

// ClassName sets the CSS class name for the container.
func (a audio) ClassName(value string) audio {
	return a.set("className", value)
}

// Controls sets the controls configuration.
func (a audio) Controls(value string) audio {
	return a.set("controls", value)
}

// Disabled sets whether the audio is disabled.
func (a audio) Disabled(value bool) audio {
	return a.set("disabled", value)
}

// DisabledOn sets the expression to determine whether the audio is disabled.
func (a audio) DisabledOn(value string) audio {
	return a.set("disabledOn", value)
}

// EditorSetting sets the editor configuration.
func (a audio) EditorSetting(value string) audio {
	return a.set("editorSetting", value)
}

// Hidden sets whether the audio is hidden.
func (a audio) Hidden(value bool) audio {
	return a.set("hidden", value)
}

// HiddenOn sets the expression to determine whether the audio is hidden.
func (a audio) HiddenOn(value string) audio {
	return a.set("hiddenOn", value)
}

// ID sets the unique ID for the component.
func (a audio) ID(value string) audio {
	return a.set("id", value)
}

// Inline sets whether the audio is displayed inline.
func (a audio) Inline(value bool) audio {
	return a.set("inline", value)
}

// Loop sets whether the audio should loop.
func (a audio) Loop(value bool) audio {
	return a.set("loop", value)
}

// OnEvent sets the event action configuration.
func (a audio) OnEvent(value any) audio {
	return a.set("onEvent", value)
}

// Rates sets the available playback speeds.
func (a audio) Rates(value string) audio {
	return a.set("rates", value)
}

// Src sets the audio source URL.
func (a audio) Src(value string) audio {
	return a.set("src", value)
}

// Static sets whether the audio is displayed statically.
func (a audio) Static(value bool) audio {
	return a.set("static", value)
}

// StaticClassName sets the CSS class name for static display.
func (a audio) StaticClassName(value string) audio {
	return a.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for the static input.
func (a audio) StaticInputClassName(value string) audio {
	return a.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for the static label.
func (a audio) StaticLabelClassName(value string) audio {
	return a.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine whether the audio is displayed statically.
func (a audio) StaticOn(value string) audio {
	return a.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display.
func (a audio) StaticPlaceholder(value string) audio {
	return a.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display.
func (a audio) StaticSchema(value string) audio {
	return a.set("staticSchema", value)
}

// Style sets the style for the component.
func (a audio) Style(value any) audio {
	return a.set("style", value)
}

// TestIdBuilder sets the test ID builder.
func (a audio) TestIdBuilder(value string) audio {
	return a.set("testIdBuilder", value)
}

// Testid sets the test ID.
func (a audio) Testid(value string) audio {
	return a.set("testid", value)
}

// UseMobileUI sets whether to use mobile UI.
func (a audio) UseMobileUI(value bool) audio {
	return a.set("useMobileUI", value)
}

// Visible sets whether the audio is visible.
func (a audio) Visible(value bool) audio {
	return a.set("visible", value)
}

// VisibleOn sets the expression to determine whether the audio is visible.
func (a audio) VisibleOn(value string) audio {
	return a.set("visibleOn", value)
}
