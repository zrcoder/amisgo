package comp

import "github.com/zrcoder/amisgo/model"

// Words represents a text display component
type Words model.Schema

func NewWords() Words {
	return Words{"type": "words"}
}

func (w Words) set(key string, value any) Words {
	w[key] = value
	return w
}

// ClassName sets the CSS class name
func (w Words) ClassName(value string) Words {
	return w.set("className", value)
}

// CollapseButton sets the collapse button text
func (w Words) CollapseButton(value string) Words {
	return w.set("collapseButton", value)
}

// CollapseButtonText sets the collapse button text
func (w Words) CollapseButtonText(value string) Words {
	return w.set("collapseButtonText", value)
}

// Delimiter sets the delimiter
func (w Words) Delimiter(value string) Words {
	return w.set("delimiter", value)
}

// Disabled sets the disabled state
func (w Words) Disabled(value bool) Words {
	return w.set("disabled", value)
}

// DisabledOn sets the disabled expression
func (w Words) DisabledOn(value string) Words {
	return w.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (w Words) EditorSetting(value string) Words {
	return w.set("editorSetting", value)
}

// ExpendButton sets the expand button text
func (w Words) ExpendButton(value string) Words {
	return w.set("expendButton", value)
}

// ExpendButtonText sets the expand button text
func (w Words) ExpendButtonText(value string) Words {
	return w.set("expendButtonText", value)
}

// Hidden sets the hidden state
func (w Words) Hidden(value bool) Words {
	return w.set("hidden", value)
}

// HiddenOn sets the hidden expression
func (w Words) HiddenOn(value string) Words {
	return w.set("hiddenOn", value)
}

// ID sets the unique component ID
func (w Words) ID(value string) Words {
	return w.set("id", value)
}

// InTag sets the tag display mode for arrays
func (w Words) InTag(value string) Words {
	return w.set("inTag", value)
}

// Limit sets the display limit
func (w Words) Limit(value string) Words {
	return w.set("limit", value)
}

// OnEvent sets the event configuration
func (w Words) OnEvent(value any) Words {
	return w.set("onEvent", value)
}

// Static sets the static display state
func (w Words) Static(value bool) Words {
	return w.set("static", value)
}

// StaticClassName sets the static display CSS class name
func (w Words) StaticClassName(value string) Words {
	return w.set("staticClassName", value)
}

// StaticInputClassName sets the static input CSS class name
func (w Words) StaticInputClassName(value string) Words {
	return w.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static label CSS class name
func (w Words) StaticLabelClassName(value string) Words {
	return w.set("staticLabelClassName", value)
}

// StaticOn sets the static display expression
func (w Words) StaticOn(value string) Words {
	return w.set("staticOn", value)
}

// StaticPlaceholder sets the static placeholder
func (w Words) StaticPlaceholder(value string) Words {
	return w.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema
func (w Words) StaticSchema(value string) Words {
	return w.set("staticSchema", value)
}

// Style sets the component style
func (w Words) Style(value any) Words {
	return w.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (w Words) TestIdBuilder(value string) Words {
	return w.set("testIdBuilder", value)
}

// Testid sets the test ID
func (w Words) Testid(value string) Words {
	return w.set("testid", value)
}

// UseMobileUI sets the mobile UI usage
func (w Words) UseMobileUI(value bool) Words {
	return w.set("useMobileUI", value)
}

// Visible sets the visibility state
func (w Words) Visible(value bool) Words {
	return w.set("visible", value)
}

// VisibleOn sets the visibility expression
func (w Words) VisibleOn(value string) Words {
	return w.set("visibleOn", value)
}

// Words sets the tags data
func (w Words) Words(value string) Words {
	return w.set("words", value)
}
