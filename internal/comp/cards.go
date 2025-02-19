package comp

import "github.com/zrcoder/amisgo/schema"

// Cards represents a collection of card components renderer
type Cards schema.Schema

func NewCards() Cards {
	return Cards{"type": "cards"}
}

func (c Cards) set(key string, value any) Cards {
	c[key] = value
	return c
}

// AffixFooter enables or disables sticky footer positioning
func (c Cards) AffixFooter(value bool) Cards {
	return c.set("affixFooter", value)
}

// AffixHeader enables or disables sticky header positioning
func (c Cards) AffixHeader(value bool) Cards {
	return c.set("affixHeader", value)
}

// Card configures the individual card settings
func (c Cards) Card(value string) Cards {
	return c.set("card", value)
}

// CheckOnItemClick enables or disables item selection on click
func (c Cards) CheckOnItemClick(value bool) Cards {
	return c.set("checkOnItemClick", value)
}

// ClassName sets the CSS class name for the cards container
func (c Cards) ClassName(value string) Cards {
	return c.set("className", value)
}

// Disabled enables or disables the cards component
func (c Cards) Disabled(value bool) Cards {
	return c.set("disabled", value)
}

// DisabledOn sets a conditional expression for disabling the cards
func (c Cards) DisabledOn(value string) Cards {
	return c.set("disabledOn", value)
}

// EditorSetting configures editor-specific settings
func (c Cards) EditorSetting(value string) Cards {
	return c.set("editorSetting", value)
}

// Footer sets the footer content for the cards
func (c Cards) Footer(value string) Cards {
	return c.set("footer", value)
}

// FooterClassName sets the CSS class name for the footer
func (c Cards) FooterClassName(value string) Cards {
	return c.set("footerClassName", value)
}

// Header sets the header content for the cards
func (c Cards) Header(value string) Cards {
	return c.set("header", value)
}

// HeaderClassName sets the CSS class name for the header
func (c Cards) HeaderClassName(value string) Cards {
	return c.set("headerClassName", value)
}

// Hidden controls the visibility of the cards component
func (c Cards) Hidden(value bool) Cards {
	return c.set("hidden", value)
}

// HiddenOn sets a conditional expression for hiding the cards
func (c Cards) HiddenOn(value string) Cards {
	return c.set("hiddenOn", value)
}

// HideCheckToggler enables or disables the checkbox toggle visibility
func (c Cards) HideCheckToggler(value bool) Cards {
	return c.set("hideCheckToggler", value)
}

// ID sets a unique identifier for the cards component
func (c Cards) ID(value string) Cards {
	return c.set("id", value)
}

// ItemCheckableOn sets a conditional expression for item checkability
func (c Cards) ItemCheckableOn(value string) Cards {
	return c.set("itemCheckableOn", value)
}

// ItemClassName sets the CSS class name for individual card items
func (c Cards) ItemClassName(value string) Cards {
	return c.set("itemClassName", value)
}

// ItemDraggableOn sets a conditional expression for item draggability
func (c Cards) ItemDraggableOn(value string) Cards {
	return c.set("itemDraggableOn", value)
}

// LoadingConfig configures the loading state and appearance
func (c Cards) LoadingConfig(value string) Cards {
	return c.set("loadingConfig", value)
}

// MasonryLayout enables or disables masonry-style layout
func (c Cards) MasonryLayout(value bool) Cards {
	return c.set("masonryLayout", value)
}

// OnEvent configures event-driven actions
func (c Cards) OnEvent(value any) Cards {
	return c.set("onEvent", value)
}

// Placeholder sets the content to display when no cards are present
func (c Cards) Placeholder(value string) Cards {
	return c.set("placeholder", value)
}

// ShowFooter enables or disables footer display
func (c Cards) ShowFooter(value bool) Cards {
	return c.set("showFooter", value)
}

// ShowHeader enables or disables header display
func (c Cards) ShowHeader(value bool) Cards {
	return c.set("showHeader", value)
}

// Source sets the data source for populating cards
func (c Cards) Source(value string) Cards {
	return c.set("source", value)
}

// Static determines if the cards are statically displayed
func (c Cards) Static(value bool) Cards {
	return c.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (c Cards) StaticClassName(value string) Cards {
	return c.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (c Cards) StaticInputClassName(value string) Cards {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (c Cards) StaticLabelClassName(value string) Cards {
	return c.set("staticLabelClassName", value)
}

// StaticOn sets a conditional expression for static display
func (c Cards) StaticOn(value string) Cards {
	return c.set("staticOn", value)
}

// StaticPlaceholder defines a placeholder for empty static values
func (c Cards) StaticPlaceholder(value string) Cards {
	return c.set("staticPlaceholder", value)
}

// StaticSchema sets the schema.Schema for static display
func (c Cards) StaticSchema(value string) Cards {
	return c.set("staticSchema", value)
}

// Style sets custom inline styles
func (c Cards) Style(value any) Cards {
	return c.set("style", value)
}

// TestIdBuilder configures test ID generation
func (c Cards) TestIdBuilder(value string) Cards {
	return c.set("testIdBuilder", value)
}

// TestID sets a specific test identifier
func (c Cards) TestID(value string) Cards {
	return c.set("testid", value)
}

// Title sets the title for the cards component
func (c Cards) Title(value any) Cards {
	return c.set("title", value)
}

// UseMobileUI enables or disables mobile UI styling
func (c Cards) UseMobileUI(value bool) Cards {
	return c.set("useMobileUI", value)
}

// ValueField sets the field to use as the value for cards
func (c Cards) ValueField(value string) Cards {
	return c.set("valueField", value)
}

// Visible controls the overall visibility of the cards
func (c Cards) Visible(value bool) Cards {
	return c.set("visible", value)
}

// VisibleOn sets a conditional expression for cards visibility
func (c Cards) VisibleOn(value string) Cards {
	return c.set("visibleOn", value)
}
