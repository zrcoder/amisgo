package comp

import "github.com/zrcoder/amisgo/model"

// image represents an image display component
type image model.Schema

// Image creates a new Image instance
func Image() image {
	return make(image).set("type", "image")
}

func (i image) set(key string, value any) image {
	i[key] = value
	return i
}

// Alt sets the alt text for the image
func (i image) Alt(value string) image {
	return i.set("alt", value)
}

// Blank sets whether the image opens in a new window
func (i image) Blank(value bool) image {
	return i.set("blank", value)
}

// Caption sets the caption for the image
func (i image) Caption(value string) image {
	return i.set("caption", value)
}

// ClassName sets the outer CSS class name
func (i image) ClassName(value string) image {
	return i.set("className", value)
}

// DefaultImage sets the default image URL
func (i image) DefaultImage(value string) image {
	return i.set("defaultImage", value)
}

// Disabled sets whether the image is disabled
func (i image) Disabled(value bool) image {
	return i.set("disabled", value)
}

// DisabledOn sets the expression to disable the image
func (i image) DisabledOn(value string) image {
	return i.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (i image) EditorSetting(value string) image {
	return i.set("editorSetting", value)
}

// EnlargeAble sets whether the image can be enlarged
func (i image) EnlargeAble(value bool) image {
	return i.set("enlargeAble", value)
}

// EnlargeWithGallary sets whether to show the image gallery when enlarged
func (i image) EnlargeWithGallary(value bool) image {
	return i.set("enlargeWithGallary", value)
}

// Height sets the height of the image
func (i image) Height(value string) image {
	return i.set("height", value)
}

// Hidden sets whether the image is hidden
func (i image) Hidden(value bool) image {
	return i.set("hidden", value)
}

// HiddenOn sets the expression to hide the image
func (i image) HiddenOn(value string) image {
	return i.set("hiddenOn", value)
}

// Href sets the link URL for the image
func (i image) Href(value string) image {
	return i.set("href", value)
}

// HtmlTarget sets the target attribute for the link
func (i image) HtmlTarget(value string) image {
	return i.set("htmlTarget", value)
}

// ID sets the unique ID for the component
func (i image) ID(value string) image {
	return i.set("id", value)
}

// ImageCaption sets the caption for the image
func (i image) ImageCaption(value string) image {
	return i.set("imageCaption", value)
}

// ImageClassName sets the CSS class name for the image
func (i image) ImageClassName(value string) image {
	return i.set("imageClassName", value)
}

// ImageGallaryClassName sets the CSS class name for the image gallery
func (i image) ImageGallaryClassName(value string) image {
	return i.set("imageGallaryClassName", value)
}

// ImageMode sets the display mode for the image (thumb | original)
func (i image) ImageMode(value string) image {
	return i.set("imageMode", value)
}

// InnerClassName sets the inner CSS class name
func (i image) InnerClassName(value string) image {
	return i.set("innerClassName", value)
}

// Name sets the associated field name
func (i image) Name(value string) image {
	return i.set("name", value)
}

// OnEvent sets the event configuration
func (i image) OnEvent(value any) image {
	return i.set("onEvent", value)
}

// OriginalSrc sets the URL for the original image
func (i image) OriginalSrc(value string) image {
	return i.set("originalSrc", value)
}

// ShowToolbar sets whether to show the image toolbar
func (i image) ShowToolbar(value bool) image {
	return i.set("showToolbar", value)
}

// Src sets the URL for the image
func (i image) Src(value string) image {
	return i.set("src", value)
}

// Static sets whether the image is displayed statically
func (i image) Static(value bool) image {
	return i.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (i image) StaticClassName(value string) image {
	return i.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (i image) StaticInputClassName(value string) image {
	return i.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (i image) StaticLabelClassName(value string) image {
	return i.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display
func (i image) StaticOn(value string) image {
	return i.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (i image) StaticPlaceholder(value string) image {
	return i.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (i image) StaticSchema(value string) image {
	return i.set("staticSchema", value)
}

// Style sets the style for the component
func (i image) Style(value any) image {
	return i.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (i image) TestIdBuilder(value string) image {
	return i.set("testIdBuilder", value)
}

// Testid sets the test ID
func (i image) Testid(value string) image {
	return i.set("testid", value)
}

// ThumbClassName sets the CSS class name for the thumbnail
func (i image) ThumbClassName(value string) image {
	return i.set("thumbClassName", value)
}

// ThumbMode sets the mode for the thumbnail
func (i image) ThumbMode(value string) image {
	return i.set("thumbMode", value)
}

// ThumbRatio sets the ratio for the thumbnail
func (i image) ThumbRatio(value string) image {
	return i.set("thumbRatio", value)
}

// Title sets the title for the image
func (i image) Title(value any) image {
	return i.set("title", value)
}

// ToolbarActions sets the toolbar actions
func (i image) ToolbarActions(value string) image {
	return i.set("toolbarActions", value)
}

// UseMobileUI sets whether to use mobile UI
func (i image) UseMobileUI(value bool) image {
	return i.set("useMobileUI", value)
}

// Visible sets whether the image is visible
func (i image) Visible(value bool) image {
	return i.set("visible", value)
}

// VisibleOn sets the expression for visibility
func (i image) VisibleOn(value string) image {
	return i.set("visibleOn", value)
}

// Width sets the width of the image
func (i image) Width(value string) image {
	return i.set("width", value)
}
