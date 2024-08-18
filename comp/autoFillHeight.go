package comp

// autoFillHeight 自动填充高度渲染器
type autoFillHeight schema

// AutoFillHeight 创建一个新的 AutoFillHeight 实例
func AutoFillHeight() autoFillHeight {
	return make(autoFillHeight)
}

func (a autoFillHeight) set(key string, value interface{}) autoFillHeight {
	a[key] = value
	return a
}

// Height 设置高度
func (a autoFillHeight) Height(value string) autoFillHeight {
	return a.set("height", value)
}

// MaxHeight 设置最大高度
func (a autoFillHeight) MaxHeight(value string) autoFillHeight {
	return a.set("maxHeight", value)
}
