package comp

import "github.com/zrcoder/amisgo/schema"

// Image represents an Image display component
type Image schema.Schema

func NewImage() Image {
	return Image{"type": "image"}
}

func (i Image) set(key string, value any) Image {
	i[key] = value
	return i
}

// Alt sets the alt text for the image
func (i Image) Alt(value string) Image {
	return i.set("alt", value)
}

// Blank sets whether the image opens in a new window
func (i Image) Blank(value bool) Image {
	return i.set("blank", value)
}

// Caption sets the caption for the image
func (i Image) Caption(value string) Image {
	return i.set("caption", value)
}

// ClassName sets the outer CSS class name
func (i Image) ClassName(value string) Image {
	return i.set("className", value)
}

// DefaultImage sets the default image URL
func (i Image) DefaultImage(value string) Image {
	return i.set("defaultImage", value)
}

// Disabled sets whether the image is disabled
func (i Image) Disabled(value bool) Image {
	return i.set("disabled", value)
}

// DisabledOn sets the expression to disable the image
func (i Image) DisabledOn(value string) Image {
	return i.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (i Image) EditorSetting(value string) Image {
	return i.set("editorSetting", value)
}

// EnlargeAble sets whether the image can be enlarged
func (i Image) EnlargeAble(value bool) Image {
	return i.set("enlargeAble", value)
}

// EnlargeWithGallary sets whether to show the image gallery when enlarged
func (i Image) EnlargeWithGallary(value bool) Image {
	return i.set("enlargeWithGallary", value)
}

// Height sets the height of the image
func (i Image) Height(value string) Image {
	return i.set("height", value)
}

// Hidden sets whether the image is hidden
func (i Image) Hidden(value bool) Image {
	return i.set("hidden", value)
}

// HiddenOn sets the expression to hide the image
func (i Image) HiddenOn(value string) Image {
	return i.set("hiddenOn", value)
}

// Href sets the link URL for the image
func (i Image) Href(value string) Image {
	return i.set("href", value)
}

// HtmlTarget sets the target attribute for the link
func (i Image) HtmlTarget(value string) Image {
	return i.set("htmlTarget", value)
}

// ID sets the unique ID for the component
func (i Image) ID(value string) Image {
	return i.set("id", value)
}

// ImageCaption sets the caption for the image
func (i Image) ImageCaption(value string) Image {
	return i.set("imageCaption", value)
}

// ImageClassName sets the CSS class name for the image
func (i Image) ImageClassName(value string) Image {
	return i.set("imageClassName", value)
}

// ImageGallaryClassName sets the CSS class name for the image gallery
func (i Image) ImageGallaryClassName(value string) Image {
	return i.set("imageGallaryClassName", value)
}

// ImageMode sets the display mode for the image (thumb | original)
func (i Image) ImageMode(value string) Image {
	return i.set("imageMode", value)
}

// InnerClassName sets the inner CSS class name
func (i Image) InnerClassName(value string) Image {
	return i.set("innerClassName", value)
}

// Name sets the associated field name
func (i Image) Name(value string) Image {
	return i.set("name", value)
}

// OnEvent sets the event configuration
func (i Image) OnEvent(value Event) Image {
	return i.set("onEvent", value)
}

// OriginalSrc sets the URL for the original image
func (i Image) OriginalSrc(value string) Image {
	return i.set("originalSrc", value)
}

// ShowToolbar sets whether to show the image toolbar
func (i Image) ShowToolbar(value bool) Image {
	return i.set("showToolbar", value)
}

// Src sets the URL for the image
func (i Image) Src(value string) Image {
	return i.set("src", value)
}

// Static sets whether the image is displayed statically
func (i Image) Static(value bool) Image {
	return i.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (i Image) StaticClassName(value string) Image {
	return i.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (i Image) StaticInputClassName(value string) Image {
	return i.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (i Image) StaticLabelClassName(value string) Image {
	return i.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display
func (i Image) StaticOn(value string) Image {
	return i.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (i Image) StaticPlaceholder(value string) Image {
	return i.set("staticPlaceholder", value)
}

// StaticSchema sets the schema.Schema for static display
func (i Image) StaticSchema(value string) Image {
	return i.set("staticSchema", value)
}

// Style sets the style for the component
func (i Image) Style(value any) Image {
	return i.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (i Image) TestIdBuilder(value string) Image {
	return i.set("testIdBuilder", value)
}

// Testid sets the test ID
func (i Image) Testid(value string) Image {
	return i.set("testid", value)
}

// ThumbClassName sets the CSS class name for the thumbnail
func (i Image) ThumbClassName(value string) Image {
	return i.set("thumbClassName", value)
}

// ThumbMode sets the mode for the thumbnail
func (i Image) ThumbMode(value string) Image {
	return i.set("thumbMode", value)
}

// ThumbRatio sets the ratio for the thumbnail
func (i Image) ThumbRatio(value string) Image {
	return i.set("thumbRatio", value)
}

// Title sets the title for the image
func (i Image) Title(value any) Image {
	return i.set("title", value)
}

// ToolbarActions sets the toolbar actions
func (i Image) ToolbarActions(value ...Action) Image {
	return i.set("toolbarActions", value)
}

// UseMobileUI sets whether to use mobile UI
func (i Image) UseMobileUI(value bool) Image {
	return i.set("useMobileUI", value)
}

// Visible sets whether the image is visible
func (i Image) Visible(value bool) Image {
	return i.set("visible", value)
}

// VisibleOn sets the expression for visibility
func (i Image) VisibleOn(value string) Image {
	return i.set("visibleOn", value)
}

// Width sets the width of the image
func (i Image) Width(value string) Image {
	return i.set("width", value)
}
