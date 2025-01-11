package comp

import "github.com/zrcoder/amisgo/model"

// Card represents a card component renderer
type card model.Schema

// Card creates a new Card instance
func Card() card {
	return card{"type": "card"}
}

func (c card) set(key string, value any) card {
	c[key] = value
	return c
}

// Actions configures the card's action buttons or menu
func (c card) Actions(value string) card {
	return c.set("actions", value)
}

// Body sets the main content of the card
func (c card) Body(value ...any) card {
	return c.set("body", value)
}

// BodyClassName sets the CSS class name for the card body
func (c card) BodyClassName(value string) card {
	return c.set("bodyClassName", value)
}

// CheckOnItemClick enables or disables item selection on click
func (c card) CheckOnItemClick(value bool) card {
	return c.set("checkOnItemClick", value)
}

// ClassName sets the CSS class name for the card container
func (c card) ClassName(value string) card {
	return c.set("className", value)
}

// Disabled enables or disables the card component
func (c card) Disabled(value bool) card {
	return c.set("disabled", value)
}

// DisabledOn sets a conditional expression for disabling the card
func (c card) DisabledOn(value string) card {
	return c.set("disabledOn", value)
}

// EditorSetting configures editor-specific settings
func (c card) EditorSetting(value string) card {
	return c.set("editorSetting", value)
}

// Header sets the card's header content
func (c card) Header(value string) card {
	return c.set("header", value)
}

// Hidden controls the visibility of the card
func (c card) Hidden(value bool) card {
	return c.set("hidden", value)
}

// HiddenOn sets a conditional expression for hiding the card
func (c card) HiddenOn(value string) card {
	return c.set("hiddenOn", value)
}

// HideCheckToggler enables or disables the checkbox toggle visibility
func (c card) HideCheckToggler(value bool) card {
	return c.set("hideCheckToggler", value)
}

// ID sets a unique identifier for the card component
func (c card) ID(value string) card {
	return c.set("id", value)
}

// ItemAction configures the action for individual card items
func (c card) ItemAction(value string) card {
	return c.set("itemAction", value)
}

// Media sets the media content for the card
func (c card) Media(value string) card {
	return c.set("media", value)
}

// OnEvent configures event-driven actions
func (c card) OnEvent(value any) card {
	return c.set("onEvent", value)
}

// Secondary sets secondary content or information for the card
func (c card) Secondary(value string) card {
	return c.set("secondary", value)
}

// Static determines if the card is statically displayed
func (c card) Static(value bool) card {
	return c.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (c card) StaticClassName(value string) card {
	return c.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (c card) StaticInputClassName(value string) card {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (c card) StaticLabelClassName(value string) card {
	return c.set("staticLabelClassName", value)
}

// StaticOn sets a conditional expression for static display
func (c card) StaticOn(value string) card {
	return c.set("staticOn", value)
}

// StaticPlaceholder defines a placeholder for empty static values
func (c card) StaticPlaceholder(value string) card {
	return c.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (c card) StaticSchema(value string) card {
	return c.set("staticSchema", value)
}

// Style sets custom inline styles
func (c card) Style(value any) card {
	return c.set("style", value)
}

// TestIdBuilder configures test ID generation
func (c card) TestIdBuilder(value string) card {
	return c.set("testIdBuilder", value)
}

// TestID sets a specific test identifier
func (c card) TestID(value string) card {
	return c.set("testid", value)
}

// Toolbar configures the card's toolbar
func (c card) Toolbar(value string) card {
	return c.set("toolbar", value)
}

// UseCardLabel enables or disables card label functionality
func (c card) UseCardLabel(value bool) card {
	return c.set("useCardLabel", value)
}

// UseMobileUI enables or disables mobile UI styling
func (c card) UseMobileUI(value bool) card {
	return c.set("useMobileUI", value)
}

// Visible controls the overall visibility of the card
func (c card) Visible(value bool) card {
	return c.set("visible", value)
}

// VisibleOn sets a conditional expression for card visibility
func (c card) VisibleOn(value string) card {
	return c.set("visibleOn", value)
}

// WrapperComponent sets the wrapper component for the card
func (c card) WrapperComponent(value string) card {
	return c.set("wrapperComponent", value)
}
