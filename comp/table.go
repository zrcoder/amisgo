package comp

// Table 表格渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/table
//
// @version 6.7.0
type Table Schema

// NewTable 创建一个新的 Table 实例
func NewTable() Table {
	return Table{}.set("type", "table")
}

func (t Table) set(key string, value interface{}) Table {
	t[key] = value
	return t
}

// AffixFooter 是否固底
func (t Table) AffixFooter(value bool) Table {
	return t.set("affixFooter", value)
}

// AffixHeader 是否固定表头
func (t Table) AffixHeader(value bool) Table {
	return t.set("affixHeader", value)
}

// AffixRow 底部总结行
func (t Table) AffixRow(value string) Table {
	return t.set("affixRow", value)
}

// AutoFillHeight 表格自动计算高度
func (t Table) AutoFillHeight(value string) Table {
	return t.set("autoFillHeight", value)
}

// AutoGenerateFilter 开启查询区域，会根据列元素的searchable属性值，自动生成查询条件表单
func (t Table) AutoGenerateFilter(value string) Table {
	return t.set("autoGenerateFilter", value)
}

// CanAccessSuperData 表格是否可以获取父级数据域值，默认为false
func (t Table) CanAccessSuperData(value bool) Table {
	return t.set("canAccessSuperData", value)
}

// ClassName 容器 css 类名
func (t Table) ClassName(value string) Table {
	return t.set("className", value)
}

// Columns 表格的列信息
func (t Table) Columns(value string) Table {
	return t.set("columns", value)
}

// ColumnsTogglable 展示列显示开关，自动即：列数量大于或等于5个时自动开启
func (t Table) ColumnsTogglable(value bool) Table {
	return t.set("columnsTogglable", value)
}

// CombineFromIndex 合并单元格配置，配置从第几列开始合并
func (t Table) CombineFromIndex(value string) Table {
	return t.set("combineFromIndex", value)
}

// CombineNum 合并单元格配置，配置数字表示从左到右的多少列自动合并单元格
func (t Table) CombineNum(value string) Table {
	return t.set("combineNum", value)
}

// Data 设置数据
func (t Table) Data(value string) Table {
	return t.set("data", value)
}

// DeferApi 懒加载 API，当行数据中用 defer: true 标记了，则其孩子节点将会用这个 API 来拉取数据
func (t Table) DeferApi(value string) Table {
	return t.set("deferApi", value)
}

// Disabled 是否禁用
func (t Table) Disabled(value bool) Table {
	return t.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (t Table) DisabledOn(value string) Table {
	return t.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (t Table) EditorSetting(value string) Table {
	return t.set("editorSetting", value)
}

// Footable 是否开启底部展示功能，适合移动端展示
func (t Table) Footable(value bool) Table {
	return t.set("footable", value)
}

// FooterClassName 底部外层 CSS 类名
func (t Table) FooterClassName(value string) Table {
	return t.set("footerClassName", value)
}

// HeaderClassName 顶部外层 CSS 类名
func (t Table) HeaderClassName(value string) Table {
	return t.set("headerClassName", value)
}

// Hidden 是否隐藏
func (t Table) Hidden(value bool) Table {
	return t.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (t Table) HiddenOn(value string) Table {
	return t.set("hiddenOn", value)
}

// Id 组件唯一 id，主要用于日志采集
func (t Table) Id(value string) Table {
	return t.set("id", value)
}

// ItemBadge 行角标 (Badge 角标)
func (t Table) ItemBadge(value string) Table {
	return t.set("itemBadge", value)
}

// OnEvent 事件动作配置
func (t Table) OnEvent(value string) Table {
	return t.set("onEvent", value)
}

// Placeholder 占位符
func (t Table) Placeholder(value string) Table {
	return t.set("placeholder", value)
}

// PrefixRow 顶部总结行
func (t Table) PrefixRow(value string) Table {
	return t.set("prefixRow", value)
}

// Resizable 是否可调整列宽
func (t Table) Resizable(value bool) Table {
	return t.set("resizable", value)
}

// RowClassNameExpr 行样式表表达式
func (t Table) RowClassNameExpr(value string) Table {
	return t.set("rowClassNameExpr", value)
}

// ShowFooter 是否显示底部
func (t Table) ShowFooter(value bool) Table {
	return t.set("showFooter", value)
}

// ShowHeader 是否显示头部
func (t Table) ShowHeader(value bool) Table {
	return t.set("showHeader", value)
}

// Source 数据源：绑定当前环境变量
func (t Table) Source(value string) Table {
	return t.set("source", value)
}

// Static 是否静态展示
func (t Table) Static(value bool) Table {
	return t.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (t Table) StaticClassName(value string) Table {
	return t.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (t Table) StaticInputClassName(value string) Table {
	return t.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (t Table) StaticLabelClassName(value string) Table {
	return t.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (t Table) StaticOn(value string) Table {
	return t.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (t Table) StaticPlaceholder(value string) Table {
	return t.set("staticPlaceholder", value)
}

// StaticSchema
func (t Table) StaticSchema(value string) Table {
	return t.set("staticSchema", value)
}

// Style 组件样式
func (t Table) Style(value string) Table {
	return t.set("style", value)
}

// TableClassName 表格 CSS 类名
func (t Table) TableClassName(value string) Table {
	return t.set("tableClassName", value)
}

// TableLayout table layout 可选值: fixed | auto
func (t Table) TableLayout(value string) Table {
	return t.set("tableLayout", value)
}

// TestIdBuilder
func (t Table) TestIdBuilder(value string) Table {
	return t.set("testIdBuilder", value)
}

// Testid
func (t Table) Testid(value string) Table {
	return t.set("testid", value)
}

// Title 标题
func (t Table) Title(value string) Table {
	return t.set("title", value)
}

// ToolbarClassName 工具栏 CSS 类名
func (t Table) ToolbarClassName(value string) Table {
	return t.set("toolbarClassName", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (t Table) UseMobileUI(value bool) Table {
	return t.set("useMobileUI", value)
}

// Visible 是否显示
func (t Table) Visible(value bool) Table {
	return t.set("visible", value)
}

// VisibleOn 是否显示表达式
func (t Table) VisibleOn(value string) Table {
	return t.set("visibleOn", value)
}
