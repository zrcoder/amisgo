package comp

// ListBodyField 列表体字段控件
//
// @version 6.7.0
type ListBodyField Schema

// NewListBodyField 创建一个新的 ListBodyField 实例
func NewListBodyField() ListBodyField {
	return make(ListBodyField)
}

// Copyable 配置点击复制功能
func (lb ListBodyField) Copyable(value bool) ListBodyField {
	lb.set("copyable", value)
	return lb
}

// InnerClassName 内层组件的 CSS 类名
func (lb ListBodyField) InnerClassName(value string) ListBodyField {
	lb.set("innerClassName", value)
	return lb
}

// Label 列标题
func (lb ListBodyField) Label(value string) ListBodyField {
	lb.set("label", value)
	return lb
}

// LabelClassName label 类名
func (lb ListBodyField) LabelClassName(value string) ListBodyField {
	lb.set("labelClassName", value)
	return lb
}

// Name 绑定字段名
func (lb ListBodyField) Name(value string) ListBodyField {
	lb.set("name", value)
	return lb
}

// PopOver 配置查看详情功能
func (lb ListBodyField) PopOver(value string) ListBodyField {
	lb.set("popOver", value)
	return lb
}

// QuickEdit 配置快速编辑功能
func (lb ListBodyField) QuickEdit(value string) ListBodyField {
	lb.set("quickEdit", value)
	return lb
}

func (lb ListBodyField) set(key string, value interface{}) ListBodyField {
	lb[key] = value
	return lb
}
