package comp

// autoGenerateFilter 自动生成过滤器渲染器
type autoGenerateFilter schema

// AutoGenerateFilter 创建一个新的 AutoGenerateFilter 实例
func AutoGenerateFilter() autoGenerateFilter {
	return make(autoGenerateFilter)
}

func (a autoGenerateFilter) set(key string, value interface{}) autoGenerateFilter {
	a[key] = value
	return a
}

// ColumnsNum 设置过滤条件单行列数
func (a autoGenerateFilter) ColumnsNum(value string) autoGenerateFilter {
	return a.set("columnsNum", value)
}

// DefaultCollapsed 设置是否默认收起
func (a autoGenerateFilter) DefaultCollapsed(value bool) autoGenerateFilter {
	return a.set("defaultCollapsed", value)
}

// ShowBtnToolbar 设置是否显示设置查询字段
func (a autoGenerateFilter) ShowBtnToolbar(value bool) autoGenerateFilter {
	return a.set("showBtnToolbar", value)
}
