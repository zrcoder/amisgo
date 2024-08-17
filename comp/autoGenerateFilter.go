package comp

// AutoGenerateFilter 自动生成过滤器渲染器
type AutoGenerateFilter Schema

// NewAutoGenerateFilter 创建一个新的 AutoGenerateFilter 实例
func NewAutoGenerateFilter() AutoGenerateFilter {
	return make(AutoGenerateFilter)
}

// Set 设置属性
func (a AutoGenerateFilter) set(key string, value interface{}) AutoGenerateFilter {
	a[key] = value
	return a
}

// ColumnsNum 设置过滤条件单行列数
func (a AutoGenerateFilter) ColumnsNum(value string) AutoGenerateFilter {
	return a.set("columnsNum", value)
}

// DefaultCollapsed 设置是否默认收起
func (a AutoGenerateFilter) DefaultCollapsed(value bool) AutoGenerateFilter {
	return a.set("defaultCollapsed", value)
}

// ShowBtnToolbar 设置是否显示设置查询字段
func (a AutoGenerateFilter) ShowBtnToolbar(value bool) AutoGenerateFilter {
	return a.set("showBtnToolbar", value)
}
