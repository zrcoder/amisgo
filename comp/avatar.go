package comp

// Avatar represents an avatar renderer
type avatar Schema

// Avatar creates a new Avatar instance with default cross-origin settings
func Avatar() avatar {
	return make(avatar).set("type", "avatar").set("crossOrigin", "anonymous")
}

func (a avatar) set(key string, value any) avatar {
	a[key] = value
	return a
}

// Alt sets the alternative text for when the image cannot be displayed
func (a avatar) Alt(value string) avatar {
	return a.set("alt", value)
}

// Badge configures the badge for the avatar
func (a avatar) Badge(value string) avatar {
	return a.set("badge", value)
}

// ClassName sets the CSS class name
func (a avatar) ClassName(value string) avatar {
	return a.set("className", value)
}

// CrossOrigin sets the CORS attribute for the image
func (a avatar) CrossOrigin(value string) avatar {
	return a.set("crossOrigin", value)
}

// DefaultAvatar sets the default avatar image
func (a avatar) DefaultAvatar(value string) avatar {
	return a.set("defaultAvatar", value)
}

// Disabled enables or disables the avatar component
func (a avatar) Disabled(value bool) avatar {
	return a.set("disabled", value)
}

// DisabledOn sets a conditional expression for disabling the avatar
func (a avatar) DisabledOn(value string) avatar {
	return a.set("disabledOn", value)
}

// Draggable allows or prevents image dragging
func (a avatar) Draggable(value bool) avatar {
	return a.set("draggable", value)
}

// EditorSetting configures editor-specific settings (ignored during runtime)
func (a avatar) EditorSetting(value string) avatar {
	return a.set("editorSetting", value)
}

// Fit sets the image scaling method relative to the container
func (a avatar) Fit(value string) avatar {
	return a.set("fit", value)
}

// Gap sets the pixel distance for character-type avatars from left and right boundaries
func (a avatar) Gap(value string) avatar {
	return a.set("gap", value)
}

// Hidden controls the visibility of the avatar
func (a avatar) Hidden(value bool) avatar {
	return a.set("hidden", value)
}

// HiddenOn sets a conditional expression for hiding the avatar
func (a avatar) HiddenOn(value string) avatar {
	return a.set("hiddenOn", value)
}

// Icon sets the icon for the avatar
func (a avatar) Icon(value string) avatar {
	return a.set("icon", value)
}

// ID sets a unique identifier for the component
func (a avatar) ID(value string) avatar {
	return a.set("id", value)
}

// OnError configures error handling for image loading failures
func (a avatar) OnError(value string) avatar {
	return a.set("onError", value)
}

// OnEvent configures event-driven actions
func (a avatar) OnEvent(value any) avatar {
	return a.set("onEvent", value)
}

// Shape sets the avatar shape
func (a avatar) Shape(value string) avatar {
	return a.set("shape", value)
}

// Size sets the avatar size
func (a avatar) Size(value string) avatar {
	return a.set("size", value)
}

// Src sets the image source URL
func (a avatar) Src(value string) avatar {
	return a.set("src", value)
}

// Static determines if the avatar is statically displayed
func (a avatar) Static(value bool) avatar {
	return a.set("static", value)
}

// StaticClassName sets the CSS class for static display
func (a avatar) StaticClassName(value string) avatar {
	return a.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class for static input display
func (a avatar) StaticInputClassName(value string) avatar {
	return a.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class for static label display
func (a avatar) StaticLabelClassName(value string) avatar {
	return a.set("staticLabelClassName", value)
}

// StaticOn sets a conditional expression for static display
func (a avatar) StaticOn(value string) avatar {
	return a.set("staticOn", value)
}

// StaticPlaceholder defines a placeholder for empty static values
func (a avatar) StaticPlaceholder(value string) avatar {
	return a.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (a avatar) StaticSchema(value string) avatar {
	return a.set("staticSchema", value)
}

// Style sets custom inline styles
func (a avatar) Style(value any) avatar {
	return a.set("style", value)
}

// TestIdBuilder configures test ID generation
func (a avatar) TestIdBuilder(value string) avatar {
	return a.set("testIdBuilder", value)
}

// Testid sets a specific test identifier
func (a avatar) Testid(value string) avatar {
	return a.set("testid", value)
}

// Text sets the text content for the avatar
func (a avatar) Text(value string) avatar {
	return a.set("text", value)
}

// UseMobileUI enables or disables mobile-specific styling
func (a avatar) UseMobileUI(value bool) avatar {
	return a.set("useMobileUI", value)
}

// Visible controls the overall visibility of the avatar
func (a avatar) Visible(value bool) avatar {
	return a.set("visible", value)
}

// VisibleOn sets a conditional expression for avatar visibility
func (a avatar) VisibleOn(value string) avatar {
	return a.set("visibleOn", value)
}
