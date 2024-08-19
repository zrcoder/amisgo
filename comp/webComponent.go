package comp

// webComponent Web Component

// @version 6.7.0
type webComponent schema

// WebComponent 创建一个新的 WebComponent 实例
func WebComponent() webComponent {
	return webComponent{}.set("type", "web-component")
}

func (wc webComponent) set(key string, value any) webComponent {
	wc[key] = value
	return wc
}

// Body 子节点
func (wc webComponent) Body(value ...any) webComponent {
	return wc.set("body", value)
}

// Props 标签上的属性
func (wc webComponent) Props(value string) webComponent {
	return wc.set("props", value)
}

// Tag 具体使用的 web-component 标签
func (wc webComponent) Tag(value string) webComponent {
	return wc.set("tag", value)
}
