package comp

import "github.com/zrcoder/amisgo/schema"

// SvgIcon represents an SVG icon renderer.
type SvgIcon schema.Schema

func NewSvgIcon() SvgIcon {
	return SvgIcon{"type": "custom-svg-icon"}
}

func (s SvgIcon) set(key string, value any) SvgIcon {
	s[key] = value
	return s
}

// ClassName sets the class name of the icon.
func (s SvgIcon) ClassName(value string) SvgIcon {
	return s.set("className", value)
}

// Icon sets the name of the icon.
func (s SvgIcon) Icon(value string) SvgIcon {
	return s.set("icon", value)
}
