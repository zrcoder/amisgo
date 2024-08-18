package comp

// WebComponent Web Component
//
// @author  slowlyo
// @version 6.7.0
type WebComponent Schema

// NewWebComponent 创建一个新的 WebComponent 实例
func NewWebComponent() WebComponent {
	return WebComponent{}.set("type", "web-component")
}

func (wc WebComponent) set(key string, value interface{}) WebComponent {
	wc[key] = value
	return wc
}

// Body 子节点
func (wc WebComponent) Body(value ...interface{}) WebComponent {
	return wc.set("body", value)
}

// Props 标签上的属性
func (wc WebComponent) Props(value string) WebComponent {
	return wc.set("props", value)
}

// Tag 具体使用的 web-component 标签
func (wc WebComponent) Tag(value string) WebComponent {
	return wc.set("tag", value)
}
