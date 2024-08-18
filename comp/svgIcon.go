package comp

// SvgIcon 图标渲染器
//
// @version 6.7.0
type SvgIcon Schema

// NewSvgIcon 创建一个新的 SvgIcon 实例
func NewSvgIcon() SvgIcon {
	return SvgIcon{}.set("type", "custom-svg-icon")
}

func (s SvgIcon) set(key string, value interface{}) SvgIcon {
	s[key] = value
	return s
}

// ClassName 设置图标的类名
func (s SvgIcon) ClassName(value string) SvgIcon {
	return s.set("className", value)
}

// Icon 设置图标的名称
func (s SvgIcon) Icon(value string) SvgIcon {
	return s.set("icon", value)
}
