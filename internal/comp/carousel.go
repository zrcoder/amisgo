package comp

import "github.com/zrcoder/amisgo/schema"

// Carousel represents a Carousel component renderer
type Carousel schema.Schema

func NewCarousel() Carousel {
	return Carousel{"type": "carousel"}
}

func (c Carousel) set(key string, value any) Carousel {
	c[key] = value
	return c
}

// AlwaysShowArrow enables or disables persistent navigation arrows
func (c Carousel) AlwaysShowArrow(value bool) Carousel {
	return c.set("alwaysShowArrow", value)
}

// Animation sets the transition animation style
func (c Carousel) Animation(value string) Carousel {
	return c.set("animation", value)
}

// Auto enables or disables automatic sliding
func (c Carousel) Auto(value bool) Carousel {
	return c.set("auto", value)
}

// ClassName sets the CSS class name for the carousel container
func (c Carousel) ClassName(value string) Carousel {
	return c.set("className", value)
}

// Controls configures the navigation control style
func (c Carousel) Controls(value string) Carousel {
	return c.set("controls", value)
}

// ControlsTheme sets the visual theme for navigation controls
func (c Carousel) ControlsTheme(value string) Carousel {
	return c.set("controlsTheme", value)
}

// Disabled enables or disables the carousel component
func (c Carousel) Disabled(value bool) Carousel {
	return c.set("disabled", value)
}

// DisabledOn sets a conditional expression for disabling the carousel
func (c Carousel) DisabledOn(value string) Carousel {
	return c.set("disabledOn", value)
}

// Duration sets the transition duration for slide changes
func (c Carousel) Duration(value string) Carousel {
	return c.set("duration", value)
}

// EditorSetting configures editor-specific settings
func (c Carousel) EditorSetting(value string) Carousel {
	return c.set("editorSetting", value)
}

// Height sets the fixed height of the carousel
func (c Carousel) Height(value string) Carousel {
	return c.set("height", value)
}

// Hidden controls the visibility of the carousel
func (c Carousel) Hidden(value bool) Carousel {
	return c.set("hidden", value)
}

// HiddenOn sets a conditional expression for hiding the carousel
func (c Carousel) HiddenOn(value string) Carousel {
	return c.set("hiddenOn", value)
}

// Icons configures custom navigation icons
func (c Carousel) Icons(value string) Carousel {
	return c.set("icons", value)
}

// ID sets a unique identifier for the carousel component
func (c Carousel) ID(value string) Carousel {
	return c.set("id", value)
}

// Interval sets the time between automatic slide transitions
func (c Carousel) Interval(value string) Carousel {
	return c.set("interval", value)
}

// ItemSchema defines the schema.Schema for individual carousel items
func (c Carousel) ItemSchema(value string) Carousel {
	return c.set("itemSchema", value)
}

// Multiple configures multiple item display mode
func (c Carousel) Multiple(value any) Carousel {
	return c.set("multiple", value)
}

// Name sets the name attribute for the carousel
func (c Carousel) Name(value string) Carousel {
	return c.set("name", value)
}

// OnEvent configures event-driven actions
func (c Carousel) OnEvent(value Event) Carousel {
	return c.set("onEvent", value)
}

// Options sets the collection of carousel items
func (c Carousel) Options(value ...CarouselOption) Carousel {
	return c.set("options", value)
}

// Placeholder sets the content to display when no items are present
func (c Carousel) Placeholder(value string) Carousel {
	return c.set("placeholder", value)
}

// Static determines if the carousel is statically displayed
func (c Carousel) Static(value bool) Carousel {
	return c.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (c Carousel) StaticClassName(value string) Carousel {
	return c.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (c Carousel) StaticInputClassName(value string) Carousel {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (c Carousel) StaticLabelClassName(value string) Carousel {
	return c.set("staticLabelClassName", value)
}

// StaticOn sets a conditional expression for static display
func (c Carousel) StaticOn(value string) Carousel {
	return c.set("staticOn", value)
}

// StaticPlaceholder defines a placeholder for empty static values
func (c Carousel) StaticPlaceholder(value string) Carousel {
	return c.set("staticPlaceholder", value)
}

// StaticSchema sets the schema.Schema for static display
func (c Carousel) StaticSchema(value string) Carousel {
	return c.set("staticSchema", value)
}

// Style sets custom inline styles
func (c Carousel) Style(value any) Carousel {
	return c.set("style", value)
}

// TestIdBuilder configures test ID generation
func (c Carousel) TestIdBuilder(value string) Carousel {
	return c.set("testIdBuilder", value)
}

// TestID sets a specific test identifier
func (c Carousel) TestID(value string) Carousel {
	return c.set("testid", value)
}

// ThumbMode sets the thumbnail display mode (contain | cover)
func (c Carousel) ThumbMode(value string) Carousel {
	return c.set("thumbMode", value)
}

// UseMobileUI enables or disables mobile UI styling
func (c Carousel) UseMobileUI(value bool) Carousel {
	return c.set("useMobileUI", value)
}

// Visible controls the overall visibility of the carousel
func (c Carousel) Visible(value bool) Carousel {
	return c.set("visible", value)
}

// VisibleOn sets a conditional expression for carousel visibility
func (c Carousel) VisibleOn(value string) Carousel {
	return c.set("visibleOn", value)
}

// Width sets the fixed width of the carousel
func (c Carousel) Width(value string) Carousel {
	return c.set("width", value)
}
