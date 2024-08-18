package comp

// svgIcon 图标渲染器
//
// @version 6.7.0
type svgIcon schema

// SvgIcon 创建一个新的 SvgIcon 实例
func SvgIcon() svgIcon {
	return svgIcon{}.set("type", "custom-svg-icon")
}

func (s svgIcon) set(key string, value interface{}) svgIcon {
	s[key] = value
	return s
}

// ClassName 设置图标的类名
func (s svgIcon) ClassName(value string) svgIcon {
	return s.set("className", value)
}

// Icon 设置图标的名称
func (s svgIcon) Icon(value string) svgIcon {
	return s.set("icon", value)
}
