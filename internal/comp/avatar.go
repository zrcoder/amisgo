package comp

import "github.com/zrcoder/amisgo/schema"

// Avatar represents an Avatar renderer
type Avatar schema.Schema

func NewAvatar() Avatar {
	return Avatar{"type": "avatar", "crossOrigin": "anonymous"}
}

func (a Avatar) set(key string, value any) Avatar {
	a[key] = value
	return a
}

// Alt sets the alternative text for when the image cannot be displayed
func (a Avatar) Alt(value string) Avatar {
	return a.set("alt", value)
}

// Badge configures the badge for the avatar
func (a Avatar) Badge(value string) Avatar {
	return a.set("badge", value)
}

// ClassName sets the CSS class name
func (a Avatar) ClassName(value string) Avatar {
	return a.set("className", value)
}

// CrossOrigin sets the CORS attribute for the image
func (a Avatar) CrossOrigin(value string) Avatar {
	return a.set("crossOrigin", value)
}

// DefaultAvatar sets the default avatar image
func (a Avatar) DefaultAvatar(value string) Avatar {
	return a.set("defaultAvatar", value)
}

// Disabled enables or disables the avatar component
func (a Avatar) Disabled(value bool) Avatar {
	return a.set("disabled", value)
}

// DisabledOn sets a conditional expression for disabling the avatar
func (a Avatar) DisabledOn(value string) Avatar {
	return a.set("disabledOn", value)
}

// Draggable allows or prevents image dragging
func (a Avatar) Draggable(value bool) Avatar {
	return a.set("draggable", value)
}

// EditorSetting configures editor-specific settings (ignored during runtime)
func (a Avatar) EditorSetting(value string) Avatar {
	return a.set("editorSetting", value)
}

// Fit sets the image scaling method relative to the container
func (a Avatar) Fit(value string) Avatar {
	return a.set("fit", value)
}

// Gap sets the pixel distance for character-type avatars from left and right boundaries
func (a Avatar) Gap(value string) Avatar {
	return a.set("gap", value)
}

// Hidden controls the visibility of the avatar
func (a Avatar) Hidden(value bool) Avatar {
	return a.set("hidden", value)
}

// HiddenOn sets a conditional expression for hiding the avatar
func (a Avatar) HiddenOn(value string) Avatar {
	return a.set("hiddenOn", value)
}

// Icon sets the icon for the avatar
func (a Avatar) Icon(value string) Avatar {
	return a.set("icon", value)
}

// ID sets a unique identifier for the component
func (a Avatar) ID(value string) Avatar {
	return a.set("id", value)
}

// OnError configures error handling for image loading failures
func (a Avatar) OnError(value string) Avatar {
	return a.set("onError", value)
}

// OnEvent configures event-driven actions
func (a Avatar) OnEvent(value any) Avatar {
	return a.set("onEvent", value)
}

// Shape sets the avatar shape
func (a Avatar) Shape(value string) Avatar {
	return a.set("shape", value)
}

// Size sets the avatar size
func (a Avatar) Size(value string) Avatar {
	return a.set("size", value)
}

// Src sets the image source URL
func (a Avatar) Src(value string) Avatar {
	return a.set("src", value)
}

// Static determines if the avatar is statically displayed
func (a Avatar) Static(value bool) Avatar {
	return a.set("static", value)
}

// StaticClassName sets the CSS class for static display
func (a Avatar) StaticClassName(value string) Avatar {
	return a.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class for static input display
func (a Avatar) StaticInputClassName(value string) Avatar {
	return a.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class for static label display
func (a Avatar) StaticLabelClassName(value string) Avatar {
	return a.set("staticLabelClassName", value)
}

// StaticOn sets a conditional expression for static display
func (a Avatar) StaticOn(value string) Avatar {
	return a.set("staticOn", value)
}

// StaticPlaceholder defines a placeholder for empty static values
func (a Avatar) StaticPlaceholder(value string) Avatar {
	return a.set("staticPlaceholder", value)
}

// StaticSchema sets the schema.Schema for static display
func (a Avatar) StaticSchema(value string) Avatar {
	return a.set("staticSchema", value)
}

// Style sets custom inline styles
func (a Avatar) Style(value any) Avatar {
	return a.set("style", value)
}

// TestIdBuilder configures test ID generation
func (a Avatar) TestIdBuilder(value string) Avatar {
	return a.set("testIdBuilder", value)
}

// Testid sets a specific test identifier
func (a Avatar) Testid(value string) Avatar {
	return a.set("testid", value)
}

// Text sets the text content for the avatar
func (a Avatar) Text(value string) Avatar {
	return a.set("text", value)
}

// UseMobileUI enables or disables mobile-specific styling
func (a Avatar) UseMobileUI(value bool) Avatar {
	return a.set("useMobileUI", value)
}

// Visible controls the overall visibility of the avatar
func (a Avatar) Visible(value bool) Avatar {
	return a.set("visible", value)
}

// VisibleOn sets a conditional expression for avatar visibility
func (a Avatar) VisibleOn(value string) Avatar {
	return a.set("visibleOn", value)
}
