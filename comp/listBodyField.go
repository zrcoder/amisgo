package comp

// listBodyField 列表体字段控件

type listBodyField Schema

// ListBodyField 创建一个新的 ListBodyField 实例
func ListBodyField() listBodyField {
	return make(listBodyField)
}

// Copyable 配置点击复制功能
func (lb listBodyField) Copyable(value bool) listBodyField {
	lb.set("copyable", value)
	return lb
}

// InnerClassName 内层组件的 CSS 类名
func (lb listBodyField) InnerClassName(value string) listBodyField {
	lb.set("innerClassName", value)
	return lb
}

// Label 列标题
func (lb listBodyField) Label(value string) listBodyField {
	lb.set("label", value)
	return lb
}

// LabelClassName label 类名
func (lb listBodyField) LabelClassName(value string) listBodyField {
	lb.set("labelClassName", value)
	return lb
}

// Name 绑定字段名
func (lb listBodyField) Name(value string) listBodyField {
	lb.set("name", value)
	return lb
}

// PopOver 配置查看详情功能
func (lb listBodyField) PopOver(value string) listBodyField {
	lb.set("popOver", value)
	return lb
}

// QuickEdit 配置快速编辑功能
func (lb listBodyField) QuickEdit(value string) listBodyField {
	lb.set("quickEdit", value)
	return lb
}

func (lb listBodyField) set(key string, value any) listBodyField {
	lb[key] = value
	return lb
}
