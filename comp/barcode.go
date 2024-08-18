package comp

// barcode 条形码渲染器
type barcode schema

// Barcode 创建一个新的 Barcode 实例
func Barcode() barcode {
	return make(barcode).set("type", "barcode")
}

func (b barcode) set(key string, value interface{}) barcode {
	b[key] = value
	return b
}

// ClassName 设置外层类名
func (b barcode) ClassName(value string) barcode {
	return b.set("className", value)
}
