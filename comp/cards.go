package comp

// Cards represents a collection of card components renderer
type cards Schema

// Cards creates a new Cards instance
func Cards() cards {
	return make(cards).set("type", "cards")
}

func (c cards) set(key string, value any) cards {
	c[key] = value
	return c
}

// AffixFooter enables or disables sticky footer positioning
func (c cards) AffixFooter(value bool) cards {
	return c.set("affixFooter", value)
}

// AffixHeader enables or disables sticky header positioning
func (c cards) AffixHeader(value bool) cards {
	return c.set("affixHeader", value)
}

// Card configures the individual card settings
func (c cards) Card(value string) cards {
	return c.set("card", value)
}

// CheckOnItemClick enables or disables item selection on click
func (c cards) CheckOnItemClick(value bool) cards {
	return c.set("checkOnItemClick", value)
}

// ClassName sets the CSS class name for the cards container
func (c cards) ClassName(value string) cards {
	return c.set("className", value)
}

// Disabled enables or disables the cards component
func (c cards) Disabled(value bool) cards {
	return c.set("disabled", value)
}

// DisabledOn sets a conditional expression for disabling the cards
func (c cards) DisabledOn(value string) cards {
	return c.set("disabledOn", value)
}

// EditorSetting configures editor-specific settings
func (c cards) EditorSetting(value string) cards {
	return c.set("editorSetting", value)
}

// Footer sets the footer content for the cards
func (c cards) Footer(value string) cards {
	return c.set("footer", value)
}

// FooterClassName sets the CSS class name for the footer
func (c cards) FooterClassName(value string) cards {
	return c.set("footerClassName", value)
}

// Header sets the header content for the cards
func (c cards) Header(value string) cards {
	return c.set("header", value)
}

// HeaderClassName sets the CSS class name for the header
func (c cards) HeaderClassName(value string) cards {
	return c.set("headerClassName", value)
}

// Hidden controls the visibility of the cards component
func (c cards) Hidden(value bool) cards {
	return c.set("hidden", value)
}

// HiddenOn sets a conditional expression for hiding the cards
func (c cards) HiddenOn(value string) cards {
	return c.set("hiddenOn", value)
}

// HideCheckToggler enables or disables the checkbox toggle visibility
func (c cards) HideCheckToggler(value bool) cards {
	return c.set("hideCheckToggler", value)
}

// ID sets a unique identifier for the cards component
func (c cards) ID(value string) cards {
	return c.set("id", value)
}

// ItemCheckableOn sets a conditional expression for item checkability
func (c cards) ItemCheckableOn(value string) cards {
	return c.set("itemCheckableOn", value)
}

// ItemClassName sets the CSS class name for individual card items
func (c cards) ItemClassName(value string) cards {
	return c.set("itemClassName", value)
}

// ItemDraggableOn sets a conditional expression for item draggability
func (c cards) ItemDraggableOn(value string) cards {
	return c.set("itemDraggableOn", value)
}

// LoadingConfig configures the loading state and appearance
func (c cards) LoadingConfig(value string) cards {
	return c.set("loadingConfig", value)
}

// MasonryLayout enables or disables masonry-style layout
func (c cards) MasonryLayout(value bool) cards {
	return c.set("masonryLayout", value)
}

// OnEvent configures event-driven actions
func (c cards) OnEvent(value any) cards {
	return c.set("onEvent", value)
}

// Placeholder sets the content to display when no cards are present
func (c cards) Placeholder(value string) cards {
	return c.set("placeholder", value)
}

// ShowFooter enables or disables footer display
func (c cards) ShowFooter(value bool) cards {
	return c.set("showFooter", value)
}

// ShowHeader enables or disables header display
func (c cards) ShowHeader(value bool) cards {
	return c.set("showHeader", value)
}

// Source sets the data source for populating cards
func (c cards) Source(value string) cards {
	return c.set("source", value)
}

// Static determines if the cards are statically displayed
func (c cards) Static(value bool) cards {
	return c.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (c cards) StaticClassName(value string) cards {
	return c.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (c cards) StaticInputClassName(value string) cards {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (c cards) StaticLabelClassName(value string) cards {
	return c.set("staticLabelClassName", value)
}

// StaticOn sets a conditional expression for static display
func (c cards) StaticOn(value string) cards {
	return c.set("staticOn", value)
}

// StaticPlaceholder defines a placeholder for empty static values
func (c cards) StaticPlaceholder(value string) cards {
	return c.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (c cards) StaticSchema(value string) cards {
	return c.set("staticSchema", value)
}

// Style sets custom inline styles
func (c cards) Style(value any) cards {
	return c.set("style", value)
}

// TestIdBuilder configures test ID generation
func (c cards) TestIdBuilder(value string) cards {
	return c.set("testIdBuilder", value)
}

// TestID sets a specific test identifier
func (c cards) TestID(value string) cards {
	return c.set("testid", value)
}

// Title sets the title for the cards component
func (c cards) Title(value any) cards {
	return c.set("title", value)
}

// UseMobileUI enables or disables mobile UI styling
func (c cards) UseMobileUI(value bool) cards {
	return c.set("useMobileUI", value)
}

// ValueField sets the field to use as the value for cards
func (c cards) ValueField(value string) cards {
	return c.set("valueField", value)
}

// Visible controls the overall visibility of the cards
func (c cards) Visible(value bool) cards {
	return c.set("visible", value)
}

// VisibleOn sets a conditional expression for cards visibility
func (c cards) VisibleOn(value string) cards {
	return c.set("visibleOn", value)
}
