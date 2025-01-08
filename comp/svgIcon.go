package comp

import "github.com/zrcoder/amisgo/model"

// svgIcon represents an SVG icon renderer.

type svgIcon model.Schema

// SvgIcon creates a new SvgIcon instance.
func SvgIcon() svgIcon {
	return svgIcon{}.set("type", "custom-svg-icon")
}

func (s svgIcon) set(key string, value any) svgIcon {
	s[key] = value
	return s
}

// ClassName sets the class name of the icon.
func (s svgIcon) ClassName(value string) svgIcon {
	return s.set("className", value)
}

// Icon sets the name of the icon.
func (s svgIcon) Icon(value string) svgIcon {
	return s.set("icon", value)
}
