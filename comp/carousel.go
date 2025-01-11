package comp

import "github.com/zrcoder/amisgo/model"

// Carousel represents a carousel component renderer
type carousel model.Schema

// Carousel creates a new Carousel instance
func Carousel() carousel {
	return carousel{"type": "carousel"}
}

func (c carousel) set(key string, value any) carousel {
	c[key] = value
	return c
}

// AlwaysShowArrow enables or disables persistent navigation arrows
func (c carousel) AlwaysShowArrow(value bool) carousel {
	return c.set("alwaysShowArrow", value)
}

// Animation sets the transition animation style
func (c carousel) Animation(value string) carousel {
	return c.set("animation", value)
}

// Auto enables or disables automatic sliding
func (c carousel) Auto(value bool) carousel {
	return c.set("auto", value)
}

// ClassName sets the CSS class name for the carousel container
func (c carousel) ClassName(value string) carousel {
	return c.set("className", value)
}

// Controls configures the navigation control style
func (c carousel) Controls(value string) carousel {
	return c.set("controls", value)
}

// ControlsTheme sets the visual theme for navigation controls
func (c carousel) ControlsTheme(value string) carousel {
	return c.set("controlsTheme", value)
}

// Disabled enables or disables the carousel component
func (c carousel) Disabled(value bool) carousel {
	return c.set("disabled", value)
}

// DisabledOn sets a conditional expression for disabling the carousel
func (c carousel) DisabledOn(value string) carousel {
	return c.set("disabledOn", value)
}

// Duration sets the transition duration for slide changes
func (c carousel) Duration(value string) carousel {
	return c.set("duration", value)
}

// EditorSetting configures editor-specific settings
func (c carousel) EditorSetting(value string) carousel {
	return c.set("editorSetting", value)
}

// Height sets the fixed height of the carousel
func (c carousel) Height(value string) carousel {
	return c.set("height", value)
}

// Hidden controls the visibility of the carousel
func (c carousel) Hidden(value bool) carousel {
	return c.set("hidden", value)
}

// HiddenOn sets a conditional expression for hiding the carousel
func (c carousel) HiddenOn(value string) carousel {
	return c.set("hiddenOn", value)
}

// Icons configures custom navigation icons
func (c carousel) Icons(value string) carousel {
	return c.set("icons", value)
}

// ID sets a unique identifier for the carousel component
func (c carousel) ID(value string) carousel {
	return c.set("id", value)
}

// Interval sets the time between automatic slide transitions
func (c carousel) Interval(value string) carousel {
	return c.set("interval", value)
}

// ItemSchema defines the schema for individual carousel items
func (c carousel) ItemSchema(value string) carousel {
	return c.set("itemSchema", value)
}

// Multiple configures multiple item display mode
func (c carousel) Multiple(value any) carousel {
	return c.set("multiple", value)
}

// Name sets the name attribute for the carousel
func (c carousel) Name(value string) carousel {
	return c.set("name", value)
}

// OnEvent configures event-driven actions
func (c carousel) OnEvent(value any) carousel {
	return c.set("onEvent", value)
}

// Options sets the collection of carousel items
func (c carousel) Options(value ...any) carousel {
	return c.set("options", value)
}

// Placeholder sets the content to display when no items are present
func (c carousel) Placeholder(value string) carousel {
	return c.set("placeholder", value)
}

// Static determines if the carousel is statically displayed
func (c carousel) Static(value bool) carousel {
	return c.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (c carousel) StaticClassName(value string) carousel {
	return c.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (c carousel) StaticInputClassName(value string) carousel {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (c carousel) StaticLabelClassName(value string) carousel {
	return c.set("staticLabelClassName", value)
}

// StaticOn sets a conditional expression for static display
func (c carousel) StaticOn(value string) carousel {
	return c.set("staticOn", value)
}

// StaticPlaceholder defines a placeholder for empty static values
func (c carousel) StaticPlaceholder(value string) carousel {
	return c.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (c carousel) StaticSchema(value string) carousel {
	return c.set("staticSchema", value)
}

// Style sets custom inline styles
func (c carousel) Style(value any) carousel {
	return c.set("style", value)
}

// TestIdBuilder configures test ID generation
func (c carousel) TestIdBuilder(value string) carousel {
	return c.set("testIdBuilder", value)
}

// TestID sets a specific test identifier
func (c carousel) TestID(value string) carousel {
	return c.set("testid", value)
}

// ThumbMode sets the thumbnail display mode (contain | cover)
func (c carousel) ThumbMode(value string) carousel {
	return c.set("thumbMode", value)
}

// UseMobileUI enables or disables mobile UI styling
func (c carousel) UseMobileUI(value bool) carousel {
	return c.set("useMobileUI", value)
}

// Visible controls the overall visibility of the carousel
func (c carousel) Visible(value bool) carousel {
	return c.set("visible", value)
}

// VisibleOn sets a conditional expression for carousel visibility
func (c carousel) VisibleOn(value string) carousel {
	return c.set("visibleOn", value)
}

// Width sets the fixed width of the carousel
func (c carousel) Width(value string) carousel {
	return c.set("width", value)
}
