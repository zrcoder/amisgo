package comp

import "github.com/zrcoder/amisgo/model"

// images represents an image gallery component.
type images model.Schema

// Images creates a new Images instance with default type.
func Images() images {
	return images{"type": "images"}
}

func (i images) set(key string, value any) images {
	i[key] = value
	return i
}

// ClassName sets the outer CSS class name.
func (i images) ClassName(value string) images {
	return i.set("className", value)
}

// DefaultImage sets the default image URL.
func (i images) DefaultImage(value string) images {
	return i.set("defaultImage", value)
}

// Delimiter sets the delimiter for values.
func (i images) Delimiter(value string) images {
	return i.set("delimiter", value)
}

// Disabled sets whether the component is disabled.
func (i images) Disabled(value bool) images {
	return i.set("disabled", value)
}

// DisabledOn sets the expression to determine if the component is disabled.
func (i images) DisabledOn(value string) images {
	return i.set("disabledOn", value)
}

// EditorSetting sets the editor configuration.
func (i images) EditorSetting(value string) images {
	return i.set("editorSetting", value)
}

// EnlargeAble sets whether the enlarge feature is enabled.
func (i images) EnlargeAble(value bool) images {
	return i.set("enlargeAble", value)
}

// EnlargetWithImages sets whether to show the image gallery when enlarged.
func (i images) EnlargetWithImages(value bool) images {
	return i.set("enlargetWithImages", value)
}

// Hidden sets whether the component is hidden.
func (i images) Hidden(value bool) images {
	return i.set("hidden", value)
}

// HiddenOn sets the expression to determine if the component is hidden.
func (i images) HiddenOn(value string) images {
	return i.set("hiddenOn", value)
}

// ID sets the unique component ID.
func (i images) ID(value string) images {
	return i.set("id", value)
}

// ImageGallaryClassName sets the CSS class name for the enlarged image gallery.
func (i images) ImageGallaryClassName(value string) images {
	return i.set("imageGallaryClassName", value)
}

// ListClassName sets the CSS class name for the list.
func (i images) ListClassName(value string) images {
	return i.set("listClassName", value)
}

// Name sets the associated field name or src.
func (i images) Name(value string) images {
	return i.set("name", value)
}

// OnEvent sets the event action configuration.
func (i images) OnEvent(value any) images {
	return i.set("onEvent", value)
}

// Options sets the options.
func (i images) Options(value ...any) images {
	return i.set("options", value)
}

// OriginalSrc sets the URL for the original image.
func (i images) OriginalSrc(value string) images {
	return i.set("originalSrc", value)
}

// Placeholder sets the placeholder text when the list is empty.
func (i images) Placeholder(value string) images {
	return i.set("placeholder", value)
}

// ShowDimensions sets whether to display dimensions.
func (i images) ShowDimensions(value bool) images {
	return i.set("showDimensions", value)
}

// ShowToolbar sets whether to display the image toolbar.
func (i images) ShowToolbar(value bool) images {
	return i.set("showToolbar", value)
}

// Source sets the source.
func (i images) Source(value string) images {
	return i.set("source", value)
}

// Src sets the image URL.
func (i images) Src(value string) images {
	return i.set("src", value)
}

// Static sets whether the component is displayed statically.
func (i images) Static(value bool) images {
	return i.set("static", value)
}

// StaticClassName sets the CSS class name for static display.
func (i images) StaticClassName(value string) images {
	return i.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display.
func (i images) StaticInputClassName(value string) images {
	return i.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display.
func (i images) StaticLabelClassName(value string) images {
	return i.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the component is displayed statically.
func (i images) StaticOn(value string) images {
	return i.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display.
func (i images) StaticPlaceholder(value string) images {
	return i.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema.
func (i images) StaticSchema(value string) images {
	return i.set("staticSchema", value)
}

// Style sets the component style.
func (i images) Style(value any) images {
	return i.set("style", value)
}

// TestIdBuilder sets the test ID builder.
func (i images) TestIdBuilder(value string) images {
	return i.set("testIdBuilder", value)
}

// Testid sets the test ID.
func (i images) Testid(value string) images {
	return i.set("testid", value)
}

// ThumbMode sets the thumbnail mode.
func (i images) ThumbMode(value string) images {
	return i.set("thumbMode", value)
}

// ThumbRatio sets the thumbnail ratio.
func (i images) ThumbRatio(value string) images {
	return i.set("thumbRatio", value)
}

// ToolbarActions sets the toolbar actions.
func (i images) ToolbarActions(value string) images {
	return i.set("toolbarActions", value)
}

// Type sets the renderer type.
func (i images) Type(value string) images {
	return i.set("type", value)
}

// UseMobileUI sets whether to use mobile UI.
func (i images) UseMobileUI(value bool) images {
	return i.set("useMobileUI", value)
}

// Value sets the value.
func (i images) Value(value string) images {
	return i.set("value", value)
}

// Visible sets whether the component is visible.
func (i images) Visible(value bool) images {
	return i.set("visible", value)
}

// VisibleOn sets the expression to determine if the component is visible.
func (i images) VisibleOn(value string) images {
	return i.set("visibleOn", value)
}
