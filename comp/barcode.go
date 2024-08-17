package comp

// Barcode 条形码渲染器
type Barcode Schema

// NewBarcode 创建一个新的 Barcode 实例
func NewBarcode() Barcode {
	return make(Barcode).set("type", "barcode")
}

func (b Barcode) set(key string, value interface{}) Barcode {
	b[key] = value
	return b
}

// ClassName 设置外层类名
func (b Barcode) ClassName(value string) Barcode {
	return b.set("className", value)
}
