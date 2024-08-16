package renderers

// Barcode 条形码渲染器
type Barcode struct {
	*BaseRenderer
}

// NewBarcode 创建一个新的 Barcode 实例
func NewBarcode() *Barcode {
	barcode := &Barcode{BaseRenderer: NewBaseRenderer()}
	barcode.set("type", "barcode")
	return barcode
}

// Set 设置属性
func (b *Barcode) set(key string, value interface{}) *Barcode {
	b.BaseRenderer.set(key, value)
	return b
}

// ClassName 设置外层类名
func (b *Barcode) ClassName(value string) *Barcode {
	return b.set("className", value)
}
