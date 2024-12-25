package comp

// table 表格渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/table

type table Schema

// Table 创建一个新的 Table 实例
func Table() table {
	return table{}.set("type", "table")
}

func Table2() table {
	return table{}.set("type", "table2")
}

func (t table) set(key string, value any) table {
	t[key] = value
	return t
}

// AffixFooter 是否固底
func (t table) AffixFooter(value bool) table {
	return t.set("affixFooter", value)
}

// AffixHeader 是否固定表头
func (t table) AffixHeader(value bool) table {
	return t.set("affixHeader", value)
}

// AffixRow 底部总结行
func (t table) AffixRow(value string) table {
	return t.set("affixRow", value)
}

// AutoFillHeight 表格自动计算高度
func (t table) AutoFillHeight(value string) table {
	return t.set("autoFillHeight", value)
}

// AutoGenerateFilter 开启查询区域，会根据列元素的searchable属性值，自动生成查询条件表单
func (t table) AutoGenerateFilter(value string) table {
	return t.set("autoGenerateFilter", value)
}

// CanAccessSuperData 表格是否可以获取父级数据域值，默认为false
func (t table) CanAccessSuperData(value bool) table {
	return t.set("canAccessSuperData", value)
}

// ClassName 容器 css 类名
func (t table) ClassName(value string) table {
	return t.set("className", value)
}

// Columns 表格的列信息
func (t table) Columns(value ...any) table {
	return t.set("columns", value)
}

// ColumnsTogglable 展示列显示开关，自动即：列数量大于或等于5个时自动开启
func (t table) ColumnsTogglable(value bool) table {
	return t.set("columnsTogglable", value)
}

// CombineFromIndex 合并单元格配置，配置从第几列开始合并
func (t table) CombineFromIndex(value string) table {
	return t.set("combineFromIndex", value)
}

// CombineNum 合并单元格配置，配置数字表示从左到右的多少列自动合并单元格
func (t table) CombineNum(value string) table {
	return t.set("combineNum", value)
}

// Data 设置数据
func (t table) Data(value Data) table {
	return t.set("data", value)
}

// DeferApi 懒加载 API，当行数据中用 defer: true 标记了，则其孩子节点将会用这个 API 来拉取数据
func (t table) DeferApi(value string) table {
	return t.set("deferApi", value)
}

// Disabled 是否禁用
func (t table) Disabled(value bool) table {
	return t.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (t table) DisabledOn(value string) table {
	return t.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (t table) EditorSetting(value string) table {
	return t.set("editorSetting", value)
}

// Footable 是否开启底部展示功能，适合移动端展示
func (t table) Footable(value bool) table {
	return t.set("footable", value)
}

// FooterClassName 底部外层 CSS 类名
func (t table) FooterClassName(value string) table {
	return t.set("footerClassName", value)
}

// HeaderClassName 顶部外层 CSS 类名
func (t table) HeaderClassName(value string) table {
	return t.set("headerClassName", value)
}

// Hidden 是否隐藏
func (t table) Hidden(value bool) table {
	return t.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (t table) HiddenOn(value string) table {
	return t.set("hiddenOn", value)
}

// Id 组件唯一 id，主要用于日志采集
func (t table) ID(value string) table {
	return t.set("id", value)
}

// ItemBadge 行角标 (Badge 角标)
func (t table) ItemBadge(value string) table {
	return t.set("itemBadge", value)
}

// OnEvent 事件动作配置
func (t table) OnEvent(value any) table {
	return t.set("onEvent", value)
}

// Placeholder 占位符
func (t table) Placeholder(value string) table {
	return t.set("placeholder", value)
}

// PrefixRow 顶部总结行
func (t table) PrefixRow(value string) table {
	return t.set("prefixRow", value)
}

// Resizable 是否可调整列宽
func (t table) Resizable(value bool) table {
	return t.set("resizable", value)
}

// RowClassNameExpr 行样式表表达式
func (t table) RowClassNameExpr(value string) table {
	return t.set("rowClassNameExpr", value)
}

// ShowFooter 是否显示底部
func (t table) ShowFooter(value bool) table {
	return t.set("showFooter", value)
}

// ShowHeader 是否显示头部
func (t table) ShowHeader(value bool) table {
	return t.set("showHeader", value)
}

// Source 数据源：绑定当前环境变量
func (t table) Source(value string) table {
	return t.set("source", value)
}

// Static 是否静态展示
func (t table) Static(value bool) table {
	return t.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (t table) StaticClassName(value string) table {
	return t.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (t table) StaticInputClassName(value string) table {
	return t.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (t table) StaticLabelClassName(value string) table {
	return t.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (t table) StaticOn(value string) table {
	return t.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (t table) StaticPlaceholder(value string) table {
	return t.set("staticPlaceholder", value)
}

// StaticSchema
func (t table) StaticSchema(value string) table {
	return t.set("staticSchema", value)
}

// Style 组件样式
func (t table) Style(value any) table {
	return t.set("style", value)
}

// TableClassName 表格 CSS 类名
func (t table) TableClassName(value string) table {
	return t.set("tableClassName", value)
}

// TableLayout table layout 可选值: fixed | auto
func (t table) TableLayout(value string) table {
	return t.set("tableLayout", value)
}

// TestIdBuilder
func (t table) TestIdBuilder(value string) table {
	return t.set("testIdBuilder", value)
}

// Testid
func (t table) Testid(value string) table {
	return t.set("testid", value)
}

// Title 标题
func (t table) Title(value any) table {
	return t.set("title", value)
}

// ToolbarClassName 工具栏 CSS 类名
func (t table) ToolbarClassName(value string) table {
	return t.set("toolbarClassName", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (t table) UseMobileUI(value bool) table {
	return t.set("useMobileUI", value)
}

// Visible 是否显示
func (t table) Visible(value bool) table {
	return t.set("visible", value)
}

// VisibleOn 是否显示表达式
func (t table) VisibleOn(value string) table {
	return t.set("visibleOn", value)
}

// Actions 操作列配置
func (t table) Actions(value string) table {
	return t.set("actions", value)
}

// Bordered 是否展示边框
func (t table) Bordered(value bool) table {
	return t.set("bordered", value)
}

// ChildrenColumnName 数据源嵌套自定义字段名
func (t table) ChildrenColumnName(value string) table {
	return t.set("childrenColumnName", value)
}

// Expandable 表格行可展开配置 (表格行可展开配置)
func (t table) Expandable(value bool) table {
	return t.set("expandable", value)
}

// Footer 指定表尾
func (t table) Footer(value string) table {
	return t.set("footer", value)
}

// KeepItemSelectionOnPageChange 翻页是否保存数据
func (t table) KeepItemSelectionOnPageChange(value bool) table {
	return t.set("keepItemSelectionOnPageChange", value)
}

// KeyField 多选、嵌套展开记录的ID字段名 默认id
func (t table) KeyField(value string) table {
	return t.set("keyField", value)
}

// LazyRenderAfter 当一次性渲染太多列上有用，默认为 100，可以用来提升表格渲染性能
func (t table) LazyRenderAfter(value string) table {
	return t.set("lazyRenderAfter", value)
}

// LineHeight 是否固定内容行高度
func (t table) LineHeight(value string) table {
	return t.set("lineHeight", value)
}

// Loading 加载中
func (t table) Loading(value string) table {
	return t.set("loading", value)
}

// MaxKeepItemSelectionLength 批量操作最大限制数
func (t table) MaxKeepItemSelectionLength(value string) table {
	return t.set("maxKeepItemSelectionLength", value)
}

// Messages 接口报错信息配置 (消息文案配置，记住这个优先级是最低的，如果你的接口返回了 msg，接口返回的优先。)
func (t table) Messages(value string) table {
	return t.set("messages", value)
}

// Multiple 是否可多选 作用同rowSelection.type 兼容原CRUD属性 不设置认为是多选 仅设置selectable才起作用
func (t table) Multiple(value bool) table {
	return t.set("multiple", value)
}

// PopOverContainer 指定挂载dom
func (t table) PopOverContainer(value string) table {
	return t.set("popOverContainer", value)
}

// PrimaryField 设置ID字段名 作用同keyFiled 兼容原CURD属性
func (t table) PrimaryField(value string) table {
	return t.set("primaryField", value)
}

// QuickSaveApi 快速编辑后用来批量保存的 API (快速编辑后用来批量保存的 API)
func (t table) QuickSaveApi(value string) table {
	return t.set("quickSaveApi", value)
}

// QuickSaveItemApi 快速编辑配置成及时保存时使用的 API (快速编辑配置成及时保存时使用的 API)
func (t table) QuickSaveItemApi(value string) table {
	return t.set("quickSaveItemApi", value)
}

// Reload 重新加载的组件名称
func (t table) Reload(value string) table {
	return t.set("reload", value)
}

// RowSelection 表格可选择配置
func (t table) RowSelection(value string) table {
	return t.set("rowSelection", value)
}

// Selectable 是否可选择 作用同rowSelection 兼容原CRUD属性 默认多选
func (t table) Selectable(value bool) table {
	return t.set("selectable", value)
}

// ShowBadge 是否展示行角标
func (t table) ShowBadge(value bool) table {
	return t.set("showBadge", value)
}

// Sticky 粘性头部
func (t table) Sticky(value bool) table {
	return t.set("sticky", value)
}

// Size 表格尺寸，支持large、default、small，default是中等尺寸。
func (t table) Size(value string) table {
	return t.set("size", value)
}
