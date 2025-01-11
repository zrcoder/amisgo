package comp

import "github.com/zrcoder/amisgo/model"

// video Video player component

type video model.Schema

// Video creates a new Video instance
func Video() video {
	return video{"type": "video"}
}

func (v video) set(key string, value any) video {
	v[key] = value
	return v
}

// AspectRatio sets the video aspect ratio
func (v video) AspectRatio(value string) video {
	return v.set("aspectRatio", value)
}

// AutoPlay sets whether the video should autoplay
func (v video) AutoPlay(value bool) video {
	return v.set("autoPlay", value)
}

// ClassName sets the CSS class name for the container
func (v video) ClassName(value string) video {
	return v.set("className", value)
}

// ColumnsCount sets the number of frames per row
func (v video) ColumnsCount(value string) video {
	return v.set("columnsCount", value)
}

// Disabled sets whether the video is disabled
func (v video) Disabled(value bool) video {
	return v.set("disabled", value)
}

// DisabledOn sets the expression to disable the video
func (v video) DisabledOn(value string) video {
	return v.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (v video) EditorSetting(value string) video {
	return v.set("editorSetting", value)
}

// Frames sets the frames for the video
func (v video) Frames(value string) video {
	return v.set("frames", value)
}

// FramesClassName sets the CSS class name for the frames container
func (v video) FramesClassName(value string) video {
	return v.set("framesClassName", value)
}

// Hidden sets whether the video is hidden
func (v video) Hidden(value bool) video {
	return v.set("hidden", value)
}

// HiddenOn sets the expression to hide the video
func (v video) HiddenOn(value string) video {
	return v.set("hiddenOn", value)
}

// ID sets the unique ID for the component
func (v video) ID(value string) video {
	return v.set("id", value)
}

// IsLive sets whether the video is live
func (v video) IsLive(value bool) video {
	return v.set("isLive", value)
}

// JumpBufferDuration sets the buffer duration for frame jumps
func (v video) JumpBufferDuration(value string) video {
	return v.set("jumpBufferDuration", value)
}

// JumpFrame sets whether to jump to the frame on click
func (v video) JumpFrame(value bool) video {
	return v.set("jumpFrame", value)
}

// Loop sets whether the video should loop
func (v video) Loop(value bool) video {
	return v.set("loop", value)
}

// Muted sets whether the video is muted initially
func (v video) Muted(value bool) video {
	return v.set("muted", value)
}

// OnEvent sets the event configuration
func (v video) OnEvent(value any) video {
	return v.set("onEvent", value)
}

// PlayerClassName sets the CSS class name for the player
func (v video) PlayerClassName(value string) video {
	return v.set("playerClassName", value)
}

// Poster sets the video poster URL
func (v video) Poster(value string) video {
	return v.set("poster", value)
}

// Rates sets the video playback rates
func (v video) Rates(value string) video {
	return v.set("rates", value)
}

// SplitPoster sets whether to display the video and poster separately
func (v video) SplitPoster(value bool) video {
	return v.set("splitPoster", value)
}

// Src sets the video source URL
func (v video) Src(value string) video {
	return v.set("src", value)
}

// Static sets whether the video is displayed statically
func (v video) Static(value bool) video {
	return v.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (v video) StaticClassName(value string) video {
	return v.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (v video) StaticInputClassName(value string) video {
	return v.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (v video) StaticLabelClassName(value string) video {
	return v.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display
func (v video) StaticOn(value string) video {
	return v.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (v video) StaticPlaceholder(value string) video {
	return v.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (v video) StaticSchema(value string) video {
	return v.set("staticSchema", value)
}

// StopOnNextFrame sets whether to stop on the next frame
func (v video) StopOnNextFrame(value bool) video {
	return v.set("stopOnNextFrame", value)
}

// Style sets the component style
func (v video) Style(value any) video {
	return v.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (v video) TestIdBuilder(value string) video {
	return v.set("testIdBuilder", value)
}

// Testid sets the test ID
func (v video) Testid(value string) video {
	return v.set("testid", value)
}

// UseMobileUI sets whether to use mobile UI
func (v video) UseMobileUI(value bool) video {
	return v.set("useMobileUI", value)
}

// VideoType sets the video type (e.g., video/x-flv)
func (v video) VideoType(value string) video {
	return v.set("videoType", value)
}

// Visible sets whether the video is visible
func (v video) Visible(value bool) video {
	return v.set("visible", value)
}

// VisibleOn sets the expression to show the video
func (v video) VisibleOn(value string) video {
	return v.set("visibleOn", value)
}
