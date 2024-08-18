package comp

// TableSchema2 表格模式2
//
// @version 6.7.0
type TableSchema2 Schema

// NewTableSchema2 创建一个新的 TableSchema2 实例
func NewTableSchema2() TableSchema2 {
	return TableSchema2{}.set("type", "table2")
}

func (ts2 TableSchema2) set(key string, value interface{}) TableSchema2 {
	ts2[key] = value
	return ts2
}

// Actions 操作列配置
func (ts2 TableSchema2) Actions(value string) TableSchema2 {
	return ts2.set("actions", value)
}

// AutoFillHeight 表格自动计算高度
func (ts2 TableSchema2) AutoFillHeight(value string) TableSchema2 {
	return ts2.set("autoFillHeight", value)
}

// Bordered 是否展示边框
func (ts2 TableSchema2) Bordered(value bool) TableSchema2 {
	return ts2.set("bordered", value)
}

// CanAccessSuperData 表格是否可以获取父级数据域值，默认为false
func (ts2 TableSchema2) CanAccessSuperData(value bool) TableSchema2 {
	return ts2.set("canAccessSuperData", value)
}

// ChildrenColumnName 数据源嵌套自定义字段名
func (ts2 TableSchema2) ChildrenColumnName(value string) TableSchema2 {
	return ts2.set("childrenColumnName", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (ts2 TableSchema2) ClassName(value string) TableSchema2 {
	return ts2.set("className", value)
}

// Columns 表格列配置
func (ts2 TableSchema2) Columns(value string) TableSchema2 {
	return ts2.set("columns", value)
}

// ColumnsTogglable 表格可自定义列
func (ts2 TableSchema2) ColumnsTogglable(value bool) TableSchema2 {
	return ts2.set("columnsTogglable", value)
}

// Disabled 是否禁用
func (ts2 TableSchema2) Disabled(value bool) TableSchema2 {
	return ts2.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (ts2 TableSchema2) DisabledOn(value string) TableSchema2 {
	return ts2.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (ts2 TableSchema2) EditorSetting(value string) TableSchema2 {
	return ts2.set("editorSetting", value)
}

// Expandable 表格行可展开配置 (表格行可展开配置)
func (ts2 TableSchema2) Expandable(value bool) TableSchema2 {
	return ts2.set("expandable", value)
}

// Footer 指定表尾
func (ts2 TableSchema2) Footer(value string) TableSchema2 {
	return ts2.set("footer", value)
}

// Hidden 是否隐藏
func (ts2 TableSchema2) Hidden(value bool) TableSchema2 {
	return ts2.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (ts2 TableSchema2) HiddenOn(value string) TableSchema2 {
	return ts2.set("hiddenOn", value)
}

// Id 组件唯一 id，主要用于日志采集
func (ts2 TableSchema2) Id(value string) TableSchema2 {
	return ts2.set("id", value)
}

// ItemBadge 行角标内容 (Badge 角标。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/badge)
func (ts2 TableSchema2) ItemBadge(value string) TableSchema2 {
	return ts2.set("itemBadge", value)
}

// KeepItemSelectionOnPageChange 翻页是否保存数据
func (ts2 TableSchema2) KeepItemSelectionOnPageChange(value bool) TableSchema2 {
	return ts2.set("keepItemSelectionOnPageChange", value)
}

// KeyField 多选、嵌套展开记录的ID字段名 默认id
func (ts2 TableSchema2) KeyField(value string) TableSchema2 {
	return ts2.set("keyField", value)
}

// LazyRenderAfter 当一次性渲染太多列上有用，默认为 100，可以用来提升表格渲染性能
func (ts2 TableSchema2) LazyRenderAfter(value string) TableSchema2 {
	return ts2.set("lazyRenderAfter", value)
}

// LineHeight 是否固定内容行高度
func (ts2 TableSchema2) LineHeight(value string) TableSchema2 {
	return ts2.set("lineHeight", value)
}

// Loading 加载中
func (ts2 TableSchema2) Loading(value string) TableSchema2 {
	return ts2.set("loading", value)
}

// MaxKeepItemSelectionLength 批量操作最大限制数
func (ts2 TableSchema2) MaxKeepItemSelectionLength(value string) TableSchema2 {
	return ts2.set("maxKeepItemSelectionLength", value)
}

// Messages 接口报错信息配置 (消息文案配置，记住这个优先级是最低的，如果你的接口返回了 msg，接口返回的优先。)
func (ts2 TableSchema2) Messages(value string) TableSchema2 {
	return ts2.set("messages", value)
}

// Multiple 是否可多选 作用同rowSelection.type 兼容原CRUD属性 不设置认为是多选 仅设置selectable才起作用
func (ts2 TableSchema2) Multiple(value bool) TableSchema2 {
	return ts2.set("multiple", value)
}

// OnEvent 事件动作配置
func (ts2 TableSchema2) OnEvent(value string) TableSchema2 {
	return ts2.set("onEvent", value)
}

// PopOverContainer 指定挂载dom
func (ts2 TableSchema2) PopOverContainer(value string) TableSchema2 {
	return ts2.set("popOverContainer", value)
}

// PrimaryField 设置ID字段名 作用同keyFiled 兼容原CURD属性
func (ts2 TableSchema2) PrimaryField(value string) TableSchema2 {
	return ts2.set("primaryField", value)
}

// QuickSaveApi 快速编辑后用来批量保存的 API (快速编辑后用来批量保存的 API)
func (ts2 TableSchema2) QuickSaveApi(value string) TableSchema2 {
	return ts2.set("quickSaveApi", value)
}

// QuickSaveItemApi 快速编辑配置成及时保存时使用的 API (快速编辑配置成及时保存时使用的 API)
func (ts2 TableSchema2) QuickSaveItemApi(value string) TableSchema2 {
	return ts2.set("quickSaveItemApi", value)
}

// Reload 重新加载的组件名称
func (ts2 TableSchema2) Reload(value string) TableSchema2 {
	return ts2.set("reload", value)
}

// RowClassNameExpr 自定义行样式
func (ts2 TableSchema2) RowClassNameExpr(value string) TableSchema2 {
	return ts2.set("rowClassNameExpr", value)
}

// RowSelection 表格可选择配置
func (ts2 TableSchema2) RowSelection(value string) TableSchema2 {
	return ts2.set("rowSelection", value)
}

// Selectable 是否可选择 作用同rowSelection 兼容原CRUD属性 默认多选
func (ts2 TableSchema2) Selectable(value bool) TableSchema2 {
	return ts2.set("selectable", value)
}

// ShowBadge 是否展示行角标
func (ts2 TableSchema2) ShowBadge(value bool) TableSchema2 {
	return ts2.set("showBadge", value)
}

// ShowHeader 是否展示表头
func (ts2 TableSchema2) ShowHeader(value bool) TableSchema2 {
	return ts2.set("showHeader", value)
}

// Source 表格数据源 (表格数据源)
func (ts2 TableSchema2) Source(value string) TableSchema2 {
	return ts2.set("source", value)
}

// Static 是否静态展示
func (ts2 TableSchema2) Static(value bool) TableSchema2 {
	return ts2.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (ts2 TableSchema2) StaticClassName(value string) TableSchema2 {
	return ts2.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (ts2 TableSchema2) StaticInputClassName(value string) TableSchema2 {
	return ts2.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (ts2 TableSchema2) StaticLabelClassName(value string) TableSchema2 {
	return ts2.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (ts2 TableSchema2) StaticOn(value string) TableSchema2 {
	return ts2.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (ts2 TableSchema2) StaticPlaceholder(value string) TableSchema2 {
	return ts2.set("staticPlaceholder", value)
}

// StaticSchema
func (ts2 TableSchema2) StaticSchema(value string) TableSchema2 {
	return ts2.set("staticSchema", value)
}

// Sticky 粘性头部
func (ts2 TableSchema2) Sticky(value bool) TableSchema2 {
	return ts2.set("sticky", value)
}

// Style 组件样式
func (ts2 TableSchema2) Style(value string) TableSchema2 {
	return ts2.set("style", value)
}

// TableLayout 可选值: fixed | auto
func (ts2 TableSchema2) TableLayout(value string) TableSchema2 {
	return ts2.set("tableLayout", value)
}

// TestIdBuilder
func (ts2 TableSchema2) TestIdBuilder(value string) TableSchema2 {
	return ts2.set("testIdBuilder", value)
}

// Testid
func (ts2 TableSchema2) Testid(value string) TableSchema2 {
	return ts2.set("testid", value)
}

// Title 表格标题
func (ts2 TableSchema2) Title(value string) TableSchema2 {
	return ts2.set("title", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (ts2 TableSchema2) UseMobileUI(value bool) TableSchema2 {
	return ts2.set("useMobileUI", value)
}

// Visible 是否显示
func (ts2 TableSchema2) Visible(value bool) TableSchema2 {
	return ts2.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (ts2 TableSchema2) VisibleOn(value string) TableSchema2 {
	return ts2.set("visibleOn", value)
}
