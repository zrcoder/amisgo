package comp

// crudTable 定义了 CRUD 表格的渲染器
type crudTable schema

// CrudTable 创建一个新的 CRUDTable 实例
func CrudTable() crudTable {
	return make(crudTable).set("type", "crud")
}

func (c crudTable) set(key string, value interface{}) crudTable {
	c[key] = value
	return c
}

// 常见配置字段

// Api 设置初始化数据 API
func (c crudTable) Api(value string) crudTable {
	return c.set("api", value)
}

// AutoFillHeight 设置内容区域是否占满屏幕剩余空间
func (c crudTable) AutoFillHeight(value interface{}) crudTable {
	return c.set("autoFillHeight", value)
}

// Columns 设置表格列配置
func (c crudTable) Columns(value string) crudTable {
	return c.set("columns", value)
}

// Disabled 设置是否禁用
func (c crudTable) Disabled(value bool) crudTable {
	return c.set("disabled", value)
}

// Expandable 设置表格行可展开配置
func (c crudTable) Expandable(value bool) crudTable {
	return c.set("expandable", value)
}

// Footer 设置指定表尾
func (c crudTable) Footer(value string) crudTable {
	return c.set("footer", value)
}

// HeaderToolbar 设置顶部区域
func (c crudTable) HeaderToolbar(value string) crudTable {
	return c.set("headerToolbar", value)
}

// Hidden 设置是否隐藏
func (c crudTable) Hidden(value bool) crudTable {
	return c.set("hidden", value)
}

// Id 设置组件唯一 id
func (c crudTable) Id(value string) crudTable {
	return c.set("id", value)
}

// Interval 设置自动刷新时间
func (c crudTable) Interval(value string) crudTable {
	return c.set("interval", value)
}

// Loading 设置加载中状态
func (c crudTable) Loading(value string) crudTable {
	return c.set("loading", value)
}

// Name 设置组件名字
func (c crudTable) Name(value string) crudTable {
	return c.set("name", value)
}

// Pagination 设置分页配置
func (c crudTable) Pagination(value interface{}) crudTable {
	return c.set("pagination", value)
}

// Placeholder 设置列表空数据的提示
func (c crudTable) Placeholder(value string) crudTable {
	return c.set("placeholder", value)
}

// RowClassName 设置行 CSS 类名
func (c crudTable) RowClassName(value string) crudTable {
	return c.set("rowClassName", value)
}

// Sticky 设置是否粘性头部
func (c crudTable) Sticky(value bool) crudTable {
	return c.set("sticky", value)
}

// Style 设置组件样式
func (c crudTable) Style(value string) crudTable {
	return c.set("style", value)
}

// TableLayout 设置表格布局
func (c crudTable) TableLayout(value string) crudTable {
	return c.set("tableLayout", value)
}

// Title 设置表格标题
func (c crudTable) Title(value string) crudTable {
	return c.set("title", value)
}

// Width 设置宽度
func (c crudTable) Width(value string) crudTable {
	return c.set("width", value)
}

// CRUD2Table 特有字段

// Actions 设置操作列配置
func (c crudTable) Actions(value string) crudTable {
	return c.set("actions", value)
}

// AutoJumpToTopOnPagerChange 设置是否自动跳到顶部
func (c crudTable) AutoJumpToTopOnPagerChange(value bool) crudTable {
	return c.set("autoJumpToTopOnPagerChange", value)
}

// Bordered 设置是否展示边框
func (c crudTable) Bordered(value bool) crudTable {
	return c.set("bordered", value)
}

// CanAccessSuperData 设置表格是否可以获取父级数据域值
func (c crudTable) CanAccessSuperData(value bool) crudTable {
	return c.set("canAccessSuperData", value)
}

// ChildrenColumnName 设置数据源嵌套自定义字段名
func (c crudTable) ChildrenColumnName(value string) crudTable {
	return c.set("childrenColumnName", value)
}

// ClassName 设置容器 css 类名
func (c crudTable) ClassName(value string) crudTable {
	return c.set("className", value)
}

// ColumnsTogglable 设置表格是否可自定义列
func (c crudTable) ColumnsTogglable(value bool) crudTable {
	return c.set("columnsTogglable", value)
}

// EditorSetting 设置编辑器配置
func (c crudTable) EditorSetting(value string) crudTable {
	return c.set("editorSetting", value)
}

// FooterToolbar 设置底部区域
func (c crudTable) FooterToolbar(value string) crudTable {
	return c.set("footerToolbar", value)
}

// FooterToolbarClassName 设置底部区域 CSS 类名
func (c crudTable) FooterToolbarClassName(value string) crudTable {
	return c.set("footerToolbarClassName", value)
}

// HeaderToolbarClassName 设置顶部区域 CSS 类名
func (c crudTable) HeaderToolbarClassName(value string) crudTable {
	return c.set("headerToolbarClassName", value)
}

// HideQuickSaveBtn 设置是否隐藏快速保存按钮
func (c crudTable) HideQuickSaveBtn(value bool) crudTable {
	return c.set("hideQuickSaveBtn", value)
}

// KeyField 设置多选、嵌套展开记录的 ID 字段名
func (c crudTable) KeyField(value string) crudTable {
	return c.set("keyField", value)
}

// LazyRenderAfter 设置一次性渲染的列数
func (c crudTable) LazyRenderAfter(value string) crudTable {
	return c.set("lazyRenderAfter", value)
}

// LineHeight 设置是否固定内容行高度
func (c crudTable) LineHeight(value string) crudTable {
	return c.set("lineHeight", value)
}

// LoadDataOnce 设置是否为前端单次加载模式
func (c crudTable) LoadDataOnce(value bool) crudTable {
	return c.set("loadDataOnce", value)
}

// LoadType 设置数据展示模式
func (c crudTable) LoadType(value string) crudTable {
	return c.set("loadType", value)
}

// LoadingConfig 设置加载中配置
func (c crudTable) LoadingConfig(value string) crudTable {
	return c.set("loadingConfig", value)
}

// MaxKeepItemSelectionLength 设置批量操作最大限制数
func (c crudTable) MaxKeepItemSelectionLength(value string) crudTable {
	return c.set("maxKeepItemSelectionLength", value)
}

// Messages 设置接口报错信息配置
func (c crudTable) Messages(value string) crudTable {
	return c.set("messages", value)
}

// Mode 设置指定内容区的展示模式
func (c crudTable) Mode(value string) crudTable {
	return c.set("mode", value)
}

// Multiple 设置是否可以多选数据
func (c crudTable) Multiple(value bool) crudTable {
	return c.set("multiple", value)
}

// OnEvent 设置事件动作配置
func (c crudTable) OnEvent(value string) crudTable {
	return c.set("onEvent", value)
}

// PageField 设置分页页码字段名
func (c crudTable) PageField(value string) crudTable {
	return c.set("pageField", value)
}

// ParsePrimitiveQuery 设置是否开启 Query 信息转换
func (c crudTable) ParsePrimitiveQuery(value string) crudTable {
	return c.set("parsePrimitiveQuery", value)
}

// PerPage 设置无限加载时每页加载数量
func (c crudTable) PerPage(value string) crudTable {
	return c.set("perPage", value)
}

// PerPageField 设置分页每页显示多少条数据的字段名
func (c crudTable) PerPageField(value string) crudTable {
	return c.set("perPageField", value)
}

// PopOverContainer 设置指定挂载 DOM
func (c crudTable) PopOverContainer(value string) crudTable {
	return c.set("popOverContainer", value)
}

// PrimaryField 设置行标识符
func (c crudTable) PrimaryField(value string) crudTable {
	return c.set("primaryField", value)
}

// QuickSaveApi 设置快速编辑后批量保存的 API
func (c crudTable) QuickSaveApi(value string) crudTable {
	return c.set("quickSaveApi", value)
}

// QuickSaveItemApi 设置快速编辑即时保存的 API
func (c crudTable) QuickSaveItemApi(value string) crudTable {
	return c.set("quickSaveItemApi", value)
}

// Reload 设置重新加载的组件名称
func (c crudTable) Reload(value string) crudTable {
	return c.set("reload", value)
}

// RowClassNameExpr 设置自定义行样式
func (c crudTable) RowClassNameExpr(value string) crudTable {
	return c.set("rowClassNameExpr", value)
}

// RowSelection 设置表格可选择配置
func (c crudTable) RowSelection(value string) crudTable {
	return c.set("rowSelection", value)
}

// SaveOrderApi 设置保存排序的 API
func (c crudTable) SaveOrderApi(value string) crudTable {
	return c.set("saveOrderApi", value)
}

// Selectable 设置是否可以选择数据
func (c crudTable) Selectable(value bool) crudTable {
	return c.set("selectable", value)
}

// ShowBadge 设置是否展示数据条数徽标
func (c crudTable) ShowBadge(value bool) crudTable {
	return c.set("showBadge", value)
}

// ShowHeader 设置是否显示表头
func (c crudTable) ShowHeader(value bool) crudTable {
	return c.set("showHeader", value)
}

// ShowPagination 设置是否显示分页
func (c crudTable) ShowPagination(value bool) crudTable {
	return c.set("showPagination", value)
}

// ShowQuickSaveBtn 设置是否展示快速保存按钮
func (c crudTable) ShowQuickSaveBtn(value bool) crudTable {
	return c.set("showQuickSaveBtn", value)
}

// ShowToolbar 设置是否显示工具栏
func (c crudTable) ShowToolbar(value bool) crudTable {
	return c.set("showToolbar", value)
}

// Source 设置表格数据源
func (c crudTable) Source(value string) crudTable {
	return c.set("source", value)
}

// Toolbar 设置指定表头工具栏
func (c crudTable) Toolbar(value string) crudTable {
	return c.set("toolbar", value)
}

// ToolbarClassName 设置工具栏 CSS 类名
func (c crudTable) ToolbarClassName(value string) crudTable {
	return c.set("toolbarClassName", value)
}

// Unfoldable 设置是否允许展开或收起
func (c crudTable) Unfoldable(value bool) crudTable {
	return c.set("unfoldable", value)
}

// Validate 设置校验配置
func (c crudTable) Validate(value string) crudTable {
	return c.set("validate", value)
}

// ValueField 设置表格数据的字段名称
func (c crudTable) ValueField(value string) crudTable {
	return c.set("valueField", value)
}

// WidthMode 设置列宽模式
func (c crudTable) WidthMode(value string) crudTable {
	return c.set("widthMode", value)
}
