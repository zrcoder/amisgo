package comp

import "github.com/zrcoder/amisgo/schema"

// Card represents a Card component renderer
type Card  schema.Schema

func NewCard() Card {
	return Card{"type": "card"}
}

func (c Card) set(key string, value any) Card {
	c[key] = value
	return c
}

// Actions configures the card's action buttons or menu
func (c Card) Actions(value string) Card {
	return c.set("actions", value)
}

// Body sets the main content of the card
func (c Card) Body(value ...any) Card {
	return c.set("body", value)
}

// BodyClassName sets the CSS class name for the card body
func (c Card) BodyClassName(value string) Card {
	return c.set("bodyClassName", value)
}

// CheckOnItemClick enables or disables item selection on click
func (c Card) CheckOnItemClick(value bool) Card {
	return c.set("checkOnItemClick", value)
}

// ClassName sets the CSS class name for the card container
func (c Card) ClassName(value string) Card {
	return c.set("className", value)
}

// Disabled enables or disables the card component
func (c Card) Disabled(value bool) Card {
	return c.set("disabled", value)
}

// DisabledOn sets a conditional expression for disabling the card
func (c Card) DisabledOn(value string) Card {
	return c.set("disabledOn", value)
}

// EditorSetting configures editor-specific settings
func (c Card) EditorSetting(value string) Card {
	return c.set("editorSetting", value)
}

// Header sets the card's header content
func (c Card) Header(value string) Card {
	return c.set("header", value)
}

// Hidden controls the visibility of the card
func (c Card) Hidden(value bool) Card {
	return c.set("hidden", value)
}

// HiddenOn sets a conditional expression for hiding the card
func (c Card) HiddenOn(value string) Card {
	return c.set("hiddenOn", value)
}

// HideCheckToggler enables or disables the checkbox toggle visibility
func (c Card) HideCheckToggler(value bool) Card {
	return c.set("hideCheckToggler", value)
}

// ID sets a unique identifier for the card component
func (c Card) ID(value string) Card {
	return c.set("id", value)
}

// ItemAction configures the action for individual card items
func (c Card) ItemAction(value string) Card {
	return c.set("itemAction", value)
}

// Media sets the media content for the card
func (c Card) Media(value string) Card {
	return c.set("media", value)
}

// OnEvent configures event-driven actions
func (c Card) OnEvent(value any) Card {
	return c.set("onEvent", value)
}

// Secondary sets secondary content or information for the card
func (c Card) Secondary(value string) Card {
	return c.set("secondary", value)
}

// Static determines if the card is statically displayed
func (c Card) Static(value bool) Card {
	return c.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (c Card) StaticClassName(value string) Card {
	return c.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (c Card) StaticInputClassName(value string) Card {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (c Card) StaticLabelClassName(value string) Card {
	return c.set("staticLabelClassName", value)
}

// StaticOn sets a conditional expression for static display
func (c Card) StaticOn(value string) Card {
	return c.set("staticOn", value)
}

// StaticPlaceholder defines a placeholder for empty static values
func (c Card) StaticPlaceholder(value string) Card {
	return c.set("staticPlaceholder", value)
}

// StaticSchema sets the schema.Schema for static display
func (c Card) StaticSchema(value string) Card {
	return c.set("staticSchema", value)
}

// Style sets custom inline styles
func (c Card) Style(value any) Card {
	return c.set("style", value)
}

// TestIdBuilder configures test ID generation
func (c Card) TestIdBuilder(value string) Card {
	return c.set("testIdBuilder", value)
}

// TestID sets a specific test identifier
func (c Card) TestID(value string) Card {
	return c.set("testid", value)
}

// Toolbar configures the card's toolbar
func (c Card) Toolbar(value string) Card {
	return c.set("toolbar", value)
}

// UseCardLabel enables or disables card label functionality
func (c Card) UseCardLabel(value bool) Card {
	return c.set("useCardLabel", value)
}

// UseMobileUI enables or disables mobile UI styling
func (c Card) UseMobileUI(value bool) Card {
	return c.set("useMobileUI", value)
}

// Visible controls the overall visibility of the card
func (c Card) Visible(value bool) Card {
	return c.set("visible", value)
}

// VisibleOn sets a conditional expression for card visibility
func (c Card) VisibleOn(value string) Card {
	return c.set("visibleOn", value)
}

// WrapperComponent sets the wrapper component for the card
func (c Card) WrapperComponent(value string) Card {
	return c.set("wrapperComponent", value)
}
