package comp

import "github.com/zrcoder/amisgo/schema"

// SearchBox represents a search box configuration.
type SearchBox schema.Schema

func NewSearchBox() SearchBox {
	return SearchBox{"type": "search-box"}
}

func (s SearchBox) set(key string, value any) SearchBox {
	s[key] = value
	return s
}

// ClassName sets the outer CSS class name.
func (s SearchBox) ClassName(value string) SearchBox {
	return s.set("className", value)
}

// ClearAndSubmit sets whether to clear content and immediately search.
func (s SearchBox) ClearAndSubmit(value bool) SearchBox {
	return s.set("clearAndSubmit", value)
}

// Clearable sets whether the content can be cleared.
func (s SearchBox) Clearable(value bool) SearchBox {
	return s.set("clearable", value)
}

// Disabled sets whether the component is disabled.
func (s SearchBox) Disabled(value bool) SearchBox {
	return s.set("disabled", value)
}

// DisabledOn sets the expression to disable the component.
func (s SearchBox) DisabledOn(value string) SearchBox {
	return s.set("disabledOn", value)
}

// EditorSetting sets the editor configuration.
func (s SearchBox) EditorSetting(value string) SearchBox {
	return s.set("editorSetting", value)
}

// Enhance sets whether it's an enhanced style.
func (s SearchBox) Enhance(value bool) SearchBox {
	return s.set("enhance", value)
}

// Hidden sets whether the component is hidden.
func (s SearchBox) Hidden(value bool) SearchBox {
	return s.set("hidden", value)
}

// HiddenOn sets the expression to hide the component.
func (s SearchBox) HiddenOn(value string) SearchBox {
	return s.set("hiddenOn", value)
}

// ID sets the unique component ID.
func (s SearchBox) ID(value string) SearchBox {
	return s.set("id", value)
}

// Loading sets whether the component is in a loading state.
func (s SearchBox) Loading(value bool) SearchBox {
	return s.set("loading", value)
}

// Mini sets whether it's a mini style.
func (s SearchBox) Mini(value bool) SearchBox {
	return s.set("mini", value)
}

// Name sets the keyword name.
func (s SearchBox) Name(value string) SearchBox {
	return s.set("name", value)
}

// OnEvent sets the event action configuration.
func (s SearchBox) OnEvent(value any) SearchBox {
	return s.set("onEvent", value)
}

// Placeholder sets the placeholder text.
func (s SearchBox) Placeholder(value string) SearchBox {
	return s.set("placeholder", value)
}

// SearchImediately sets whether to search immediately.
func (s SearchBox) SearchImediately(value bool) SearchBox {
	return s.set("searchImediately", value)
}

// Static sets whether it's a static display.
func (s SearchBox) Static(value bool) SearchBox {
	return s.set("static", value)
}

// StaticClassName sets the static display form item class name.
func (s SearchBox) StaticClassName(value string) SearchBox {
	return s.set("staticClassName", value)
}

// StaticInputClassName sets the static display form item value class name.
func (s SearchBox) StaticInputClassName(value string) SearchBox {
	return s.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static display form item label class name.
func (s SearchBox) StaticLabelClassName(value string) SearchBox {
	return s.set("staticLabelClassName", value)
}

// StaticOn sets the static display expression.
func (s SearchBox) StaticOn(value string) SearchBox {
	return s.set("staticOn", value)
}

// StaticPlaceholder sets the static display empty placeholder.
func (s SearchBox) StaticPlaceholder(value string) SearchBox {
	return s.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema.Schema.
func (s SearchBox) StaticSchema(value string) SearchBox {
	return s.set("staticSchema", value)
}

// Style sets the component style.
func (s SearchBox) Style(value any) SearchBox {
	return s.set("style", value)
}

// TestIdBuilder sets the test ID builder.
func (s SearchBox) TestIdBuilder(value string) SearchBox {
	return s.set("testIdBuilder", value)
}

// Testid sets the test ID.
func (s SearchBox) Testid(value string) SearchBox {
	return s.set("testid", value)
}

// UseMobileUI sets whether to use mobile UI at the component level.
func (s SearchBox) UseMobileUI(value bool) SearchBox {
	return s.set("useMobileUI", value)
}

// Visible sets whether the component is visible.
func (s SearchBox) Visible(value bool) SearchBox {
	return s.set("visible", value)
}

// VisibleOn sets the expression to make the component visible.
func (s SearchBox) VisibleOn(value string) SearchBox {
	return s.set("visibleOn", value)
}
