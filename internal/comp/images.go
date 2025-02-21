package comp

import "github.com/zrcoder/amisgo/schema"

// Images represents an image gallery component.
type Images schema.Schema

func NewImages() Images {
	return Images{"type": "images"}
}

func (i Images) set(key string, value any) Images {
	i[key] = value
	return i
}

// ClassName sets the outer CSS class name.
func (i Images) ClassName(value string) Images {
	return i.set("className", value)
}

// DefaultImage sets the default image URL.
func (i Images) DefaultImage(value string) Images {
	return i.set("defaultImage", value)
}

// Delimiter sets the delimiter for values.
func (i Images) Delimiter(value string) Images {
	return i.set("delimiter", value)
}

// Disabled sets whether the component is disabled.
func (i Images) Disabled(value bool) Images {
	return i.set("disabled", value)
}

// DisabledOn sets the expression to determine if the component is disabled.
func (i Images) DisabledOn(value string) Images {
	return i.set("disabledOn", value)
}

// EditorSetting sets the editor configuration.
func (i Images) EditorSetting(value string) Images {
	return i.set("editorSetting", value)
}

// EnlargeAble sets whether the enlarge feature is enabled.
func (i Images) EnlargeAble(value bool) Images {
	return i.set("enlargeAble", value)
}

// EnlargetWithImages sets whether to show the image gallery when enlarged.
func (i Images) EnlargetWithImages(value bool) Images {
	return i.set("enlargetWithImages", value)
}

// Hidden sets whether the component is hidden.
func (i Images) Hidden(value bool) Images {
	return i.set("hidden", value)
}

// HiddenOn sets the expression to determine if the component is hidden.
func (i Images) HiddenOn(value string) Images {
	return i.set("hiddenOn", value)
}

// ID sets the unique component ID.
func (i Images) ID(value string) Images {
	return i.set("id", value)
}

// ImageGallaryClassName sets the CSS class name for the enlarged image gallery.
func (i Images) ImageGallaryClassName(value string) Images {
	return i.set("imageGallaryClassName", value)
}

// ListClassName sets the CSS class name for the list.
func (i Images) ListClassName(value string) Images {
	return i.set("listClassName", value)
}

// Name sets the associated field name or src.
func (i Images) Name(value string) Images {
	return i.set("name", value)
}

// OnEvent sets the event action configuration.
func (i Images) OnEvent(value Event) Images {
	return i.set("onEvent", value)
}

// Options sets the options.
func (i Images) Options(value ...any) Images {
	return i.set("options", value)
}

// OriginalSrc sets the URL for the original image.
func (i Images) OriginalSrc(value string) Images {
	return i.set("originalSrc", value)
}

// Placeholder sets the placeholder text when the list is empty.
func (i Images) Placeholder(value string) Images {
	return i.set("placeholder", value)
}

// ShowDimensions sets whether to display dimensions.
func (i Images) ShowDimensions(value bool) Images {
	return i.set("showDimensions", value)
}

// ShowToolbar sets whether to display the image toolbar.
func (i Images) ShowToolbar(value bool) Images {
	return i.set("showToolbar", value)
}

// Source sets the source.
func (i Images) Source(value string) Images {
	return i.set("source", value)
}

// Src sets the image URL.
func (i Images) Src(value string) Images {
	return i.set("src", value)
}

// Static sets whether the component is displayed statically.
func (i Images) Static(value bool) Images {
	return i.set("static", value)
}

// StaticClassName sets the CSS class name for static display.
func (i Images) StaticClassName(value string) Images {
	return i.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display.
func (i Images) StaticInputClassName(value string) Images {
	return i.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display.
func (i Images) StaticLabelClassName(value string) Images {
	return i.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the component is displayed statically.
func (i Images) StaticOn(value string) Images {
	return i.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display.
func (i Images) StaticPlaceholder(value string) Images {
	return i.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema.Schema.
func (i Images) StaticSchema(value string) Images {
	return i.set("staticSchema", value)
}

// Style sets the component style.
func (i Images) Style(value any) Images {
	return i.set("style", value)
}

// TestIdBuilder sets the test ID builder.
func (i Images) TestIdBuilder(value string) Images {
	return i.set("testIdBuilder", value)
}

// Testid sets the test ID.
func (i Images) Testid(value string) Images {
	return i.set("testid", value)
}

// ThumbMode sets the thumbnail mode.
func (i Images) ThumbMode(value string) Images {
	return i.set("thumbMode", value)
}

// ThumbRatio sets the thumbnail ratio.
func (i Images) ThumbRatio(value string) Images {
	return i.set("thumbRatio", value)
}

// ToolbarActions sets the toolbar actions.
func (i Images) ToolbarActions(value ...Action) Images {
	return i.set("toolbarActions", value)
}

// Type sets the renderer type.
func (i Images) Type(value string) Images {
	return i.set("type", value)
}

// UseMobileUI sets whether to use mobile UI.
func (i Images) UseMobileUI(value bool) Images {
	return i.set("useMobileUI", value)
}

// Value sets the value.
func (i Images) Value(value string) Images {
	return i.set("value", value)
}

// Visible sets whether the component is visible.
func (i Images) Visible(value bool) Images {
	return i.set("visible", value)
}

// VisibleOn sets the expression to determine if the component is visible.
func (i Images) VisibleOn(value string) Images {
	return i.set("visibleOn", value)
}
