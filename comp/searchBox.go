package comp

import "github.com/zrcoder/amisgo/model"

// searchBox represents a search box configuration.
type searchBox model.Schema

// SearchBox creates a new SearchBox instance.
func SearchBox() searchBox {
	return searchBox{}.set("type", "search-box")
}

func (s searchBox) set(key string, value any) searchBox {
	s[key] = value
	return s
}

// ClassName sets the outer CSS class name.
func (s searchBox) ClassName(value string) searchBox {
	return s.set("className", value)
}

// ClearAndSubmit sets whether to clear content and immediately search.
func (s searchBox) ClearAndSubmit(value bool) searchBox {
	return s.set("clearAndSubmit", value)
}

// Clearable sets whether the content can be cleared.
func (s searchBox) Clearable(value bool) searchBox {
	return s.set("clearable", value)
}

// Disabled sets whether the component is disabled.
func (s searchBox) Disabled(value bool) searchBox {
	return s.set("disabled", value)
}

// DisabledOn sets the expression to disable the component.
func (s searchBox) DisabledOn(value string) searchBox {
	return s.set("disabledOn", value)
}

// EditorSetting sets the editor configuration.
func (s searchBox) EditorSetting(value string) searchBox {
	return s.set("editorSetting", value)
}

// Enhance sets whether it's an enhanced style.
func (s searchBox) Enhance(value bool) searchBox {
	return s.set("enhance", value)
}

// Hidden sets whether the component is hidden.
func (s searchBox) Hidden(value bool) searchBox {
	return s.set("hidden", value)
}

// HiddenOn sets the expression to hide the component.
func (s searchBox) HiddenOn(value string) searchBox {
	return s.set("hiddenOn", value)
}

// ID sets the unique component ID.
func (s searchBox) ID(value string) searchBox {
	return s.set("id", value)
}

// Loading sets whether the component is in a loading state.
func (s searchBox) Loading(value bool) searchBox {
	return s.set("loading", value)
}

// Mini sets whether it's a mini style.
func (s searchBox) Mini(value bool) searchBox {
	return s.set("mini", value)
}

// Name sets the keyword name.
func (s searchBox) Name(value string) searchBox {
	return s.set("name", value)
}

// OnEvent sets the event action configuration.
func (s searchBox) OnEvent(value any) searchBox {
	return s.set("onEvent", value)
}

// Placeholder sets the placeholder text.
func (s searchBox) Placeholder(value string) searchBox {
	return s.set("placeholder", value)
}

// SearchImediately sets whether to search immediately.
func (s searchBox) SearchImediately(value bool) searchBox {
	return s.set("searchImediately", value)
}

// Static sets whether it's a static display.
func (s searchBox) Static(value bool) searchBox {
	return s.set("static", value)
}

// StaticClassName sets the static display form item class name.
func (s searchBox) StaticClassName(value string) searchBox {
	return s.set("staticClassName", value)
}

// StaticInputClassName sets the static display form item value class name.
func (s searchBox) StaticInputClassName(value string) searchBox {
	return s.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static display form item label class name.
func (s searchBox) StaticLabelClassName(value string) searchBox {
	return s.set("staticLabelClassName", value)
}

// StaticOn sets the static display expression.
func (s searchBox) StaticOn(value string) searchBox {
	return s.set("staticOn", value)
}

// StaticPlaceholder sets the static display empty placeholder.
func (s searchBox) StaticPlaceholder(value string) searchBox {
	return s.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema.
func (s searchBox) StaticSchema(value string) searchBox {
	return s.set("staticSchema", value)
}

// Style sets the component style.
func (s searchBox) Style(value any) searchBox {
	return s.set("style", value)
}

// TestIdBuilder sets the test ID builder.
func (s searchBox) TestIdBuilder(value string) searchBox {
	return s.set("testIdBuilder", value)
}

// Testid sets the test ID.
func (s searchBox) Testid(value string) searchBox {
	return s.set("testid", value)
}

// UseMobileUI sets whether to use mobile UI at the component level.
func (s searchBox) UseMobileUI(value bool) searchBox {
	return s.set("useMobileUI", value)
}

// Visible sets whether the component is visible.
func (s searchBox) Visible(value bool) searchBox {
	return s.set("visible", value)
}

// VisibleOn sets the expression to make the component visible.
func (s searchBox) VisibleOn(value string) searchBox {
	return s.set("visibleOn", value)
}
