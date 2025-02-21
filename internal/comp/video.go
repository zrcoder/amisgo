package comp

import "github.com/zrcoder/amisgo/schema"

// Video Video player component
type Video schema.Schema

func NewVideo() Video {
	return Video{"type": "video"}
}

func (v Video) set(key string, value any) Video {
	v[key] = value
	return v
}

// AspectRatio sets the video aspect ratio
func (v Video) AspectRatio(value string) Video {
	return v.set("aspectRatio", value)
}

// AutoPlay sets whether the video should autoplay
func (v Video) AutoPlay(value bool) Video {
	return v.set("autoPlay", value)
}

// ClassName sets the CSS class name for the container
func (v Video) ClassName(value string) Video {
	return v.set("className", value)
}

// ColumnsCount sets the number of frames per row
func (v Video) ColumnsCount(value string) Video {
	return v.set("columnsCount", value)
}

// Disabled sets whether the video is disabled
func (v Video) Disabled(value bool) Video {
	return v.set("disabled", value)
}

// DisabledOn sets the expression to disable the video
func (v Video) DisabledOn(value string) Video {
	return v.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (v Video) EditorSetting(value string) Video {
	return v.set("editorSetting", value)
}

// Frames sets the frames for the video
func (v Video) Frames(value string) Video {
	return v.set("frames", value)
}

// FramesClassName sets the CSS class name for the frames container
func (v Video) FramesClassName(value string) Video {
	return v.set("framesClassName", value)
}

// Hidden sets whether the video is hidden
func (v Video) Hidden(value bool) Video {
	return v.set("hidden", value)
}

// HiddenOn sets the expression to hide the video
func (v Video) HiddenOn(value string) Video {
	return v.set("hiddenOn", value)
}

// ID sets the unique ID for the component
func (v Video) ID(value string) Video {
	return v.set("id", value)
}

// IsLive sets whether the video is live
func (v Video) IsLive(value bool) Video {
	return v.set("isLive", value)
}

// JumpBufferDuration sets the buffer duration for frame jumps
func (v Video) JumpBufferDuration(value string) Video {
	return v.set("jumpBufferDuration", value)
}

// JumpFrame sets whether to jump to the frame on click
func (v Video) JumpFrame(value bool) Video {
	return v.set("jumpFrame", value)
}

// Loop sets whether the video should loop
func (v Video) Loop(value bool) Video {
	return v.set("loop", value)
}

// Muted sets whether the video is muted initially
func (v Video) Muted(value bool) Video {
	return v.set("muted", value)
}

// OnEvent sets the event configuration
func (v Video) OnEvent(value Event) Video {
	return v.set("onEvent", value)
}

// PlayerClassName sets the CSS class name for the player
func (v Video) PlayerClassName(value string) Video {
	return v.set("playerClassName", value)
}

// Poster sets the video poster URL
func (v Video) Poster(value string) Video {
	return v.set("poster", value)
}

// Rates sets the video playback rates
func (v Video) Rates(value string) Video {
	return v.set("rates", value)
}

// SplitPoster sets whether to display the video and poster separately
func (v Video) SplitPoster(value bool) Video {
	return v.set("splitPoster", value)
}

// Src sets the video source URL
func (v Video) Src(value string) Video {
	return v.set("src", value)
}

// Static sets whether the video is displayed statically
func (v Video) Static(value bool) Video {
	return v.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (v Video) StaticClassName(value string) Video {
	return v.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (v Video) StaticInputClassName(value string) Video {
	return v.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (v Video) StaticLabelClassName(value string) Video {
	return v.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display
func (v Video) StaticOn(value string) Video {
	return v.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (v Video) StaticPlaceholder(value string) Video {
	return v.set("staticPlaceholder", value)
}

// StaticSchema sets the schema.Schema for static display
func (v Video) StaticSchema(value string) Video {
	return v.set("staticSchema", value)
}

// StopOnNextFrame sets whether to stop on the next frame
func (v Video) StopOnNextFrame(value bool) Video {
	return v.set("stopOnNextFrame", value)
}

// Style sets the component style
func (v Video) Style(value any) Video {
	return v.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (v Video) TestIdBuilder(value string) Video {
	return v.set("testIdBuilder", value)
}

// Testid sets the test ID
func (v Video) Testid(value string) Video {
	return v.set("testid", value)
}

// UseMobileUI sets whether to use mobile UI
func (v Video) UseMobileUI(value bool) Video {
	return v.set("useMobileUI", value)
}

// VideoType sets the video type (e.g., video/x-flv)
func (v Video) VideoType(value string) Video {
	return v.set("videoType", value)
}

// Visible sets whether the video is visible
func (v Video) Visible(value bool) Video {
	return v.set("visible", value)
}

// VisibleOn sets the expression to show the video
func (v Video) VisibleOn(value string) Video {
	return v.set("visibleOn", value)
}
