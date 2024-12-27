package comp

// crud 定义了 CRUD 表格和列表的渲染器
type crud Schema

// Crud
func Crud() crud {
	return CrudTable()
}

// CrudTable 创建一个新的 CRUDTable 实例
func CrudTable() crud {
	return crud{}.set("type", "crud").Mode("table")
}

// CrudList
func CrudList() crud {
	return crud{}.set("type", "crud").Mode("list")
}

// CrudCards
func CrudCards() crud {
	return crud{}.set("type", "crud").Mode("cards")
}

// Set 设置键值对，并返回当前实例
func (c crud) set(key string, value any) crud {
	c[key] = value
	return c
}

// 常见配置字段

// Api 设置初始化数据 API
func (c crud) Api(value string) crud {
	return c.set("api", value)
}

// FetchData 设置获取数据的 api 实现方法
func (c crud) FetchData(getter func() (any, error)) crud {
	return c.Api(serveData(getter))
}

// AutoFillHeight 设置内容区域是否占满屏幕剩余空间
func (c crud) AutoFillHeight(value any) crud {
	return c.set("autoFillHeight", value)
}

// Columns 设置表格列配置
func (c crud) Columns(value ...any) crud {
	return c.set("columns", value)
}

// Disabled 设置是否禁用
func (c crud) Disabled(value bool) crud {
	return c.set("disabled", value)
}

// Expandable 设置表格行可展开配置
func (c crud) Expandable(value bool) crud {
	return c.set("expandable", value)
}

// Footer 设置指定表尾
func (c crud) Footer(value string) crud {
	return c.set("footer", value)
}

// HeaderToolbar 设置顶部区域
func (c crud) HeaderToolbar(value ...any) crud {
	return c.set("headerToolbar", value)
}

// Hidden 设置是否隐藏
func (c crud) Hidden(value bool) crud {
	return c.set("hidden", value)
}

// Interval 设置自动刷新时间
func (c crud) Interval(value string) crud {
	return c.set("interval", value)
}

// Loading 设置加载中状态
func (c crud) Loading(value string) crud {
	return c.set("loading", value)
}

// Name 设置组件名字
func (c crud) Name(value string) crud {
	return c.set("name", value)
}

// Pagination 设置分页配置
func (c crud) Pagination(value any) crud {
	return c.set("pagination", value)
}

// Placeholder 设置列表空数据的提示
func (c crud) Placeholder(value string) crud {
	return c.set("placeholder", value)
}

// RowClassName 设置行 CSS 类名
func (c crud) RowClassName(value string) crud {
	return c.set("rowClassName", value)
}

// Sticky 设置是否粘性头部
func (c crud) Sticky(value bool) crud {
	return c.set("sticky", value)
}

// Style 设置组件样式
func (c crud) Style(value any) crud {
	return c.set("style", value)
}

// TableLayout 设置表格布局
func (c crud) TableLayout(value string) crud {
	return c.set("tableLayout", value)
}

// Title 设置表格标题
func (c crud) Title(value any) crud {
	return c.set("title", value)
}

// Width 设置宽度
func (c crud) Width(value string) crud {
	return c.set("width", value)
}

// CRUD2Table 特有字段

// Actions 设置操作列配置
func (c crud) Actions(value string) crud {
	return c.set("actions", value)
}

// AutoJumpToTopOnPagerChange 设置是否自动跳到顶部
func (c crud) AutoJumpToTopOnPagerChange(value bool) crud {
	return c.set("autoJumpToTopOnPagerChange", value)
}

// Bordered 设置是否展示边框
func (c crud) Bordered(value bool) crud {
	return c.set("bordered", value)
}

// CanAccessSuperData 设置表格是否可以获取父级数据域值
func (c crud) CanAccessSuperData(value bool) crud {
	return c.set("canAccessSuperData", value)
}

// ChildrenColumnName 设置数据源嵌套自定义字段名
func (c crud) ChildrenColumnName(value string) crud {
	return c.set("childrenColumnName", value)
}

// ClassName 设置容器 css 类名
func (c crud) ClassName(value string) crud {
	return c.set("className", value)
}

// ColumnsTogglable 设置表格是否可自定义列
func (c crud) ColumnsTogglable(value bool) crud {
	return c.set("columnsTogglable", value)
}

// EditorSetting 设置编辑器配置
func (c crud) EditorSetting(value string) crud {
	return c.set("editorSetting", value)
}

// FooterToolbar 设置底部区域
func (c crud) FooterToolbar(value ...any) crud {
	return c.set("footerToolbar", value)
}

// FooterToolbarClassName 设置底部区域 CSS 类名
func (c crud) FooterToolbarClassName(value string) crud {
	return c.set("footerToolbarClassName", value)
}

// HeaderToolbarClassName 设置顶部区域 CSS 类名
func (c crud) HeaderToolbarClassName(value string) crud {
	return c.set("headerToolbarClassName", value)
}

// HideQuickSaveBtn 设置是否隐藏快速保存按钮
func (c crud) HideQuickSaveBtn(value bool) crud {
	return c.set("hideQuickSaveBtn", value)
}

// KeyField 设置多选、嵌套展开记录的 ID 字段名
func (c crud) KeyField(value string) crud {
	return c.set("keyField", value)
}

// LazyRenderAfter 设置一次性渲染的列数
func (c crud) LazyRenderAfter(value string) crud {
	return c.set("lazyRenderAfter", value)
}

// LineHeight 设置是否固定内容行高度
func (c crud) LineHeight(value string) crud {
	return c.set("lineHeight", value)
}

// LoadDataOnce 设置是否为前端单次加载模式
func (c crud) LoadDataOnce(value bool) crud {
	return c.set("loadDataOnce", value)
}

// LoadType 设置数据展示模式
func (c crud) LoadType(value string) crud {
	return c.set("loadType", value)
}

// LoadingConfig 设置加载中配置
func (c crud) LoadingConfig(value string) crud {
	return c.set("loadingConfig", value)
}

// MaxKeepItemSelectionLength 设置批量操作最大限制数
func (c crud) MaxKeepItemSelectionLength(value string) crud {
	return c.set("maxKeepItemSelectionLength", value)
}

// Messages 设置接口报错信息配置
func (c crud) Messages(value string) crud {
	return c.set("messages", value)
}

// Mode 设置指定内容区的展示模式 table | list | cards 默认 table
func (c crud) Mode(value string) crud {
	return c.set("mode", value)
}

// Multiple 设置是否可以多选数据
func (c crud) Multiple(value bool) crud {
	return c.set("multiple", value)
}

// OnEvent 设置事件动作配置
func (c crud) OnEvent(value any) crud {
	return c.set("onEvent", value)
}

// PageField 设置分页页码字段名
func (c crud) PageField(value string) crud {
	return c.set("pageField", value)
}

// ParsePrimitiveQuery 设置是否开启 Query 信息转换
func (c crud) ParsePrimitiveQuery(value string) crud {
	return c.set("parsePrimitiveQuery", value)
}

// PerPage 设置无限加载时每页加载数量
func (c crud) PerPage(value int) crud {
	return c.set("perPage", value)
}

// PerPageField 设置分页每页显示多少条数据的字段名
func (c crud) PerPageField(value string) crud {
	return c.set("perPageField", value)
}

// PopOverContainer 设置指定挂载 DOM
func (c crud) PopOverContainer(value string) crud {
	return c.set("popOverContainer", value)
}

// PrimaryField 设置行标识符
func (c crud) PrimaryField(value string) crud {
	return c.set("primaryField", value)
}

// QuickSaveApi 设置快速编辑后批量保存的 API
func (c crud) QuickSaveApi(value string) crud {
	return c.set("quickSaveApi", value)
}

// QuickSaveItemApi 设置快速编辑即时保存的 API
func (c crud) QuickSaveItemApi(value string) crud {
	return c.set("quickSaveItemApi", value)
}

// Reload 设置重新加载的组件名称
func (c crud) Reload(value string) crud {
	return c.set("reload", value)
}

// RowClassNameExpr 设置自定义行样式
func (c crud) RowClassNameExpr(value string) crud {
	return c.set("rowClassNameExpr", value)
}

// RowSelection 设置表格可选择配置
func (c crud) RowSelection(value string) crud {
	return c.set("rowSelection", value)
}

// SaveOrderApi 设置排序保存 API
func (c crud) SaveOrderApi(value string) crud {
	return c.set("saveOrderApi", value)
}

// Selectable 设置是否可以选择表格数据
func (c crud) Selectable(value bool) crud {
	return c.set("selectable", value)
}

// ShowBadge 设置是否展示徽标
func (c crud) ShowBadge(value bool) crud {
	return c.set("showBadge", value)
}

// ShowHeader 设置是否展示头部
func (c crud) ShowHeader(value bool) crud {
	return c.set("showHeader", value)
}

// ShowPagination 设置是否展示分页
func (c crud) ShowPagination(value bool) crud {
	return c.set("showPagination", value)
}

// ShowQuickSaveBtn 设置是否展示快速保存按钮
func (c crud) ShowQuickSaveBtn(value bool) crud {
	return c.set("showQuickSaveBtn", value)
}

// ShowToolbar 设置是否展示工具栏
func (c crud) ShowToolbar(value bool) crud {
	return c.set("showToolbar", value)
}

// Source 设置数据源
func (c crud) Source(value string) crud {
	return c.set("source", value)
}

// Toolbar 设置指定表头工具栏
func (c crud) Toolbar(value string) crud {
	return c.set("toolbar", value)
}

// ToolbarClassName 设置工具栏 CSS 类名
func (c crud) ToolbarClassName(value string) crud {
	return c.set("toolbarClassName", value)
}

// Unfoldable 设置是否允许展开或收起
func (c crud) Unfoldable(value bool) crud {
	return c.set("unfoldable", value)
}

// Validate 设置校验配置
func (c crud) Validate(value string) crud {
	return c.set("validate", value)
}

// ValueField 设置表格数据的字段名称
func (c crud) ValueField(value string) crud {
	return c.set("valueField", value)
}

// WidthMode 设置列宽模式
func (c crud) WidthMode(value string) crud {
	return c.set("widthMode", value)
}

// AffixFooter 设置是否固底
func (c crud) AffixFooter(value bool) crud {
	return c.set("affixFooter", value)
}

// AffixHeader 设置是否固顶
func (c crud) AffixHeader(value bool) crud {
	return c.set("affixHeader", value)
}

// AlwaysShowPagination 设置是否总是显示分页
func (c crud) AlwaysShowPagination(value bool) crud {
	return c.set("alwaysShowPagination", value)
}

// AutoGenerateFilter 开启自动生成查询条件表单
func (c crud) AutoGenerateFilter(value any) crud {
	return c.set("autoGenerateFilter", value)
}

// BulkActions 设置批量操作
func (c crud) BulkActions(value ...any) crud {
	return c.set("bulkActions", value)
}

// Card 设置卡片配置
func (c crud) Card(value any) crud {
	return c.set("card", value)
}

// CheckOnItemClick 设置点击卡片时是否勾选卡片
func (c crud) CheckOnItemClick(value bool) crud {
	return c.set("checkOnItemClick", value)
}

// DefaultParams 设置默认参数
func (c crud) DefaultParams(value string) crud {
	return c.set("defaultParams", value)
}

// DeferApi 设置懒加载 API
func (c crud) DeferApi(value string) crud {
	return c.set("deferApi", value)
}

// DisabledOn 设置禁用表达式
func (c crud) DisabledOn(value string) crud {
	return c.set("disabledOn", value)
}

// Draggable 设置是否可通过拖拽排序
func (c crud) Draggable(value bool) crud {
	return c.set("draggable", value)
}

// DraggableOn 设置拖拽排序表达式
func (c crud) DraggableOn(value string) crud {
	return c.set("draggableOn", value)
}

// ExpandConfig 设置展开配置
func (c crud) ExpandConfig(value string) crud {
	return c.set("expandConfig", value)
}

// Filter 设置过滤器表单
func (c crud) Filter(value form) crud {
	return c.set("filter", value)
}

// FilterDefaultVisible 设置过滤器默认是否可见
func (c crud) FilterDefaultVisible(value bool) crud {
	return c.set("filterDefaultVisible", value)
}

// FilterTogglable 设置过滤器是否可切换
func (c crud) FilterTogglable(value bool) crud {
	return c.set("filterTogglable", value)
}

// FooterClassName 设置底部 CSS 类名
func (c crud) FooterClassName(value string) crud {
	return c.set("footerClassName", value)
}

// Header 设置顶部区域
func (c crud) Header(value any) crud {
	return c.set("header", value)
}

// HeaderClassName 设置顶部 CSS 类名
func (c crud) HeaderClassName(value string) crud {
	return c.set("headerClassName", value)
}

// HiddenOn 设置隐藏表达式
func (c crud) HiddenOn(value string) crud {
	return c.set("hiddenOn", value)
}

// HideCheckToggler 设置是否隐藏勾选框
func (c crud) HideCheckToggler(value bool) crud {
	return c.set("hideCheckToggler", value)
}

// InitFetch 设置初始是否拉取
func (c crud) InitFetch(value bool) crud {
	return c.set("initFetch", value)
}

// InitFetchOn 设置初始拉取表达式
func (c crud) InitFetchOn(value string) crud {
	return c.set("initFetchOn", value)
}

// InnerClassName 设置内部 DOM 的 CSS 类名
func (c crud) InnerClassName(value string) crud {
	return c.set("innerClassName", value)
}

// ItemActions 设置单条操作
func (c crud) ItemActions(value string) crud {
	return c.set("itemActions", value)
}

// ItemCheckableOn 设置约束批量操作
func (c crud) ItemCheckableOn(value string) crud {
	return c.set("itemCheckableOn", value)
}

// ItemClassName 设置卡片 CSS 类名
func (c crud) ItemClassName(value string) crud {
	return c.set("itemClassName", value)
}

// ItemDraggableOn 设置项是否可拖拽排序
func (c crud) ItemDraggableOn(value string) crud {
	return c.set("itemDraggableOn", value)
}

// KeepItemSelectionOnPageChange 设置分页时是否保留用户选择
func (c crud) KeepItemSelectionOnPageChange(value bool) crud {
	return c.set("keepItemSelectionOnPageChange", value)
}

// LabelTpl 设置已勾选项的文案
func (c crud) LabelTpl(value string) crud {
	return c.set("labelTpl", value)
}

// LoadDataOnceFetchOnFilter 设置 loadDataOnce 时的过滤条件
func (c crud) LoadDataOnceFetchOnFilter(value bool) crud {
	return c.set("loadDataOnceFetchOnFilter", value)
}

// MasonryLayout 设置是否为瀑布流布局
func (c crud) MasonryLayout(value bool) crud {
	return c.set("masonryLayout", value)
}

// MatchFunc 设置自定义搜索匹配函数
func (c crud) MatchFunc(value string) crud {
	return c.set("matchFunc", value)
}

// OrderBy 设置默认排序字段
func (c crud) OrderBy(value string) crud {
	return c.set("orderBy", value)
}

// Query 设置查询字段
func (c crud) Query(value string) crud {
	return c.set("query", value)
}

// Render设置是否渲染组件
func (c crud) Render(value bool) crud {
	return c.set("render", value)
}

// ResetPage 设置是否重置分页
func (c crud) ResetPage(value bool) crud {
	return c.set("resetPage", value)
}

// Searchable 设置是否可搜索
func (c crud) Searchable(value bool) crud {
	return c.set("searchable", value)
}

// SearchPlaceholder 设置搜索占位符
func (c crud) SearchPlaceholder(value string) crud {
	return c.set("searchPlaceholder", value)
}

// ShowErrorMsg 设置是否显示错误信息
func (c crud) ShowErrorMsg(value bool) crud {
	return c.set("showErrorMsg", value)
}

// ShowSearch 设置是否显示搜索框
func (c crud) ShowSearch(value bool) crud {
	return c.set("showSearch", value)
}

// ShowSort 设置是否显示排序
func (c crud) ShowSort(value bool) crud {
	return c.set("showSort", value)
}

// ShowSwitch 设置是否显示切换按钮
func (c crud) ShowSwitch(value bool) crud {
	return c.set("showSwitch", value)
}

// Size 设置组件尺寸
func (c crud) Size(value string) crud {
	return c.set("size", value)
}

// Sortable 设置是否可排序
func (c crud) Sortable(value bool) crud {
	return c.set("sortable", value)
}

// SubmitText 设置提交按钮文字
func (c crud) SubmitText(value string) crud {
	return c.set("submitText", value)
}

// SyncLocation 设置是否同步 URL
func (c crud) SyncLocation(value bool) crud {
	return c.set("syncLocation", value)
}

// Translations 设置多语言翻译
func (c crud) Translations(value map[string]string) crud {
	return c.set("translations", value)
}

// VerticalAlign 设置垂直对齐方式
func (c crud) VerticalAlign(value string) crud {
	return c.set("verticalAlign", value)
}

// Visible 设置是否可见
func (c crud) Visible(value bool) crud {
	return c.set("visible", value)
}

// VisibleOn 设置可见性表达式
func (c crud) VisibleOn(value string) crud {
	return c.set("visibleOn", value)
}

// WrapItemClassName 设置包裹项的 CSS 类名
func (c crud) WrapItemClassName(value string) crud {
	return c.set("wrapItemClassName", value)
}

// ExtraAction 设置额外操作
func (c crud) ExtraAction(value any) crud {
	return c.set("extraAction", value)
}

// AutoSaveOnChange 设置是否自动保存更改
func (c crud) AutoSaveOnChange(value bool) crud {
	return c.set("autoSaveOnChange", value)
}

// ID 设置组件唯一 ID
func (c crud) ID(value string) crud {
	return c.set("id", value)
}

// ItemAction 设置点击列表项的行为
func (c crud) ItemAction(value string) crud {
	return c.set("itemAction", value)
}

// ListItem 单条数据展示内容配置
func (c crud) ListItem(value any) crud {
	return c.set("listItem", value)
}

// QuickSaveItemActions 设置快速编辑按钮配置
func (c crud) QuickSaveItemActions(value string) crud {
	return c.set("quickSaveItemActions", value)
}

// QuickSaveSuccessMessage 设置快速保存成功的提示信息
func (c crud) QuickSaveSuccessMessage(value string) crud {
	return c.set("quickSaveSuccessMessage", value)
}

// RefreshInterval 设置自动刷新时间
func (c crud) RefreshInterval(value string) crud {
	return c.set("refreshInterval", value)
}

// RemoteFilter 支持远程筛选
func (c crud) RemoteFilter(value bool) crud {
	return c.set("remoteFilter", value)
}

// RenderType 设置渲染类型
func (c crud) RenderType(value string) crud {
	return c.set("renderType", value)
}

// ResetValueOnHidden 设置隐藏时是否重置值
func (c crud) ResetValueOnHidden(value bool) crud {
	return c.set("resetValueOnHidden", value)
}

// SearchFormClassName 设置搜索表单 CSS 类名
func (c crud) SearchFormClassName(value string) crud {
	return c.set("searchFormClassName", value)
}

// SearchSchema 设置搜索表单
func (c crud) SearchSchema(value any) crud {
	return c.set("searchSchema", value)
}

// SortSettings 设置排序配置
func (c crud) SortSettings(value string) crud {
	return c.set("sortSettings", value)
}

// SortFieldName 设置排序字段名
func (c crud) SortFieldName(value string) crud {
	return c.set("sortFieldName", value)
}

// Stateful 设置是否保留状态
func (c crud) Stateful(value bool) crud {
	return c.set("stateful", value)
}

// TableClassName 设置表格 CSS 类名
func (c crud) TableClassName(value string) crud {
	return c.set("tableClassName", value)
}

// Tooltip设置 提示信息配置
func (c crud) Tooltip(value string) crud {
	return c.set("tooltip", value)
}

// TransformItems 设置数据转换
func (c crud) TransformItems(value ...any) crud {
	return c.set("transformItems", value)
}

// Virtualized 设置虚拟化
func (c crud) Virtualized(value bool) crud {
	return c.set("virtualized", value)
}

// EmptyText 设置无数据时的提示
func (c crud) EmptyText(value string) crud {
	return c.set("emptyText", value)
}

// Components 设置自定义组件
func (c crud) Components(value any) crud {
	return c.set("components", value)
}

// 其他 CRUD2List 特有字段和方法
func (c crud) CustomField1(value string) crud {
	return c.set("customField1", value)
}

// 其他 crudTable 特有字段和方法
func (c crud) CustomField2(value string) crud {
	return c.set("customField2", value)
}
