package comp

// CRUD2Table 定义了 CRUD2 表格的渲染器
type CRUD2Table Schema

// NewCRUD2Table 创建一个新的 CRUD2Table 实例
func NewCRUD2Table() CRUD2Table {
	return make(CRUD2Table).set("type", "crud2")
}

func (t CRUD2Table) set(key string, value interface{}) CRUD2Table {
	t[key] = value
	return t
}

// Actions 设置操作列配置
func (t CRUD2Table) Actions(value string) CRUD2Table {
	return t.set("actions", value)
}

// Api 设置初始化数据 API
func (t CRUD2Table) Api(value string) CRUD2Table {
	return t.set("api", value)
}

// AutoFillHeight 设置内容区域是否占满屏幕剩余空间
func (t CRUD2Table) AutoFillHeight(value bool) CRUD2Table {
	return t.set("autoFillHeight", value)
}

// AutoJumpToTopOnPagerChange 设置是否自动跳到顶部
func (t CRUD2Table) AutoJumpToTopOnPagerChange(value bool) CRUD2Table {
	return t.set("autoJumpToTopOnPagerChange", value)
}

// Bordered 设置是否展示边框
func (t CRUD2Table) Bordered(value bool) CRUD2Table {
	return t.set("bordered", value)
}

// CanAccessSuperData 设置表格是否可以获取父级数据域值
func (t CRUD2Table) CanAccessSuperData(value bool) CRUD2Table {
	return t.set("canAccessSuperData", value)
}

// ChildrenColumnName 设置数据源嵌套自定义字段名
func (t CRUD2Table) ChildrenColumnName(value string) CRUD2Table {
	return t.set("childrenColumnName", value)
}

// ClassName 设置容器 css 类名
func (t CRUD2Table) ClassName(value string) CRUD2Table {
	return t.set("className", value)
}

// Columns 设置表格列配置
func (t CRUD2Table) Columns(value string) CRUD2Table {
	return t.set("columns", value)
}

// ColumnsTogglable 设置表格是否可自定义列
func (t CRUD2Table) ColumnsTogglable(value bool) CRUD2Table {
	return t.set("columnsTogglable", value)
}

// Disabled 设置是否禁用
func (t CRUD2Table) Disabled(value bool) CRUD2Table {
	return t.set("disabled", value)
}

// DisabledOn 设置禁用表达式
func (t CRUD2Table) DisabledOn(value string) CRUD2Table {
	return t.set("disabledOn", value)
}

// EditorSetting 设置编辑器配置
func (t CRUD2Table) EditorSetting(value string) CRUD2Table {
	return t.set("editorSetting", value)
}

// Expandable 设置表格行可展开配置
func (t CRUD2Table) Expandable(value bool) CRUD2Table {
	return t.set("expandable", value)
}

// Footer 设置指定表尾
func (t CRUD2Table) Footer(value string) CRUD2Table {
	return t.set("footer", value)
}

// FooterToolbar 设置底部区域
func (t CRUD2Table) FooterToolbar(value string) CRUD2Table {
	return t.set("footerToolbar", value)
}

// FooterToolbarClassName 设置底部区域 CSS 类名
func (t CRUD2Table) FooterToolbarClassName(value string) CRUD2Table {
	return t.set("footerToolbarClassName", value)
}

// HeaderToolbar 设置顶部区域
func (t CRUD2Table) HeaderToolbar(value string) CRUD2Table {
	return t.set("headerToolbar", value)
}

// HeaderToolbarClassName 设置顶部区域 CSS 类名
func (t CRUD2Table) HeaderToolbarClassName(value string) CRUD2Table {
	return t.set("headerToolbarClassName", value)
}

// Hidden 设置是否隐藏
func (t CRUD2Table) Hidden(value bool) CRUD2Table {
	return t.set("hidden", value)
}

// HiddenOn 设置隐藏表达式
func (t CRUD2Table) HiddenOn(value string) CRUD2Table {
	return t.set("hiddenOn", value)
}

// HideQuickSaveBtn 设置是否隐藏快速保存按钮
func (t CRUD2Table) HideQuickSaveBtn(value bool) CRUD2Table {
	return t.set("hideQuickSaveBtn", value)
}

// Id 设置组件唯一 id
func (t CRUD2Table) Id(value string) CRUD2Table {
	return t.set("id", value)
}

// Interval 设置自动刷新时间
func (t CRUD2Table) Interval(value string) CRUD2Table {
	return t.set("interval", value)
}

// ItemBadge 设置行角标内容
func (t CRUD2Table) ItemBadge(value string) CRUD2Table {
	return t.set("itemBadge", value)
}

// ItemCheckableOn 设置约束批量操作
func (t CRUD2Table) ItemCheckableOn(value string) CRUD2Table {
	return t.set("itemCheckableOn", value)
}

// KeepItemSelectionOnPageChange 设置翻页时是否保留用户已选的数据
func (t CRUD2Table) KeepItemSelectionOnPageChange(value bool) CRUD2Table {
	return t.set("keepItemSelectionOnPageChange", value)
}

// KeyField 设置多选、嵌套展开记录的 ID 字段名
func (t CRUD2Table) KeyField(value string) CRUD2Table {
	return t.set("keyField", value)
}

// LazyRenderAfter 设置一次性渲染的列数
func (t CRUD2Table) LazyRenderAfter(value string) CRUD2Table {
	return t.set("lazyRenderAfter", value)
}

// LineHeight 设置是否固定内容行高度
func (t CRUD2Table) LineHeight(value string) CRUD2Table {
	return t.set("lineHeight", value)
}

// LoadDataOnce 设置是否为前端单次加载模式
func (t CRUD2Table) LoadDataOnce(value bool) CRUD2Table {
	return t.set("loadDataOnce", value)
}

// LoadType 设置数据展示模式
func (t CRUD2Table) LoadType(value string) CRUD2Table {
	return t.set("loadType", value)
}

// Loading 设置加载中状态
func (t CRUD2Table) Loading(value string) CRUD2Table {
	return t.set("loading", value)
}

// LoadingConfig 设置加载中配置
func (t CRUD2Table) LoadingConfig(value string) CRUD2Table {
	return t.set("loadingConfig", value)
}

// MaxKeepItemSelectionLength 设置批量操作最大限制数
func (t CRUD2Table) MaxKeepItemSelectionLength(value string) CRUD2Table {
	return t.set("maxKeepItemSelectionLength", value)
}

// Messages 设置接口报错信息配置
func (t CRUD2Table) Messages(value string) CRUD2Table {
	return t.set("messages", value)
}

// Mode 设置指定内容区的展示模式
func (t CRUD2Table) Mode(value string) CRUD2Table {
	return t.set("mode", value)
}

// Multiple 设置是否可以多选数据
func (t CRUD2Table) Multiple(value bool) CRUD2Table {
	return t.set("multiple", value)
}

// Name 设置组件名字
func (t CRUD2Table) Name(value string) CRUD2Table {
	return t.set("name", value)
}

// OnEvent 设置事件动作配置
func (t CRUD2Table) OnEvent(value string) CRUD2Table {
	return t.set("onEvent", value)
}

// PageField 设置分页页码字段名
func (t CRUD2Table) PageField(value string) CRUD2Table {
	return t.set("pageField", value)
}

// ParsePrimitiveQuery 设置是否开启 Query 信息转换
func (t CRUD2Table) ParsePrimitiveQuery(value string) CRUD2Table {
	return t.set("parsePrimitiveQuery", value)
}

// PerPage 设置无限加载时每页加载数量
func (t CRUD2Table) PerPage(value string) CRUD2Table {
	return t.set("perPage", value)
}

// PerPageField 设置分页每页显示多少条数据的字段名
func (t CRUD2Table) PerPageField(value string) CRUD2Table {
	return t.set("perPageField", value)
}

// PopOverContainer 设置指定挂载 DOM
func (t CRUD2Table) PopOverContainer(value string) CRUD2Table {
	return t.set("popOverContainer", value)
}

// PrimaryField 设置行标识符
func (t CRUD2Table) PrimaryField(value string) CRUD2Table {
	return t.set("primaryField", value)
}

// QuickSaveApi 设置快速编辑后批量保存的 API
func (t CRUD2Table) QuickSaveApi(value string) CRUD2Table {
	return t.set("quickSaveApi", value)
}

// QuickSaveItemApi 设置快速编辑即时保存的 API
func (t CRUD2Table) QuickSaveItemApi(value string) CRUD2Table {
	return t.set("quickSaveItemApi", value)
}

// Reload 设置重新加载的组件名称
func (t CRUD2Table) Reload(value string) CRUD2Table {
	return t.set("reload", value)
}

// RowClassNameExpr 设置自定义行样式
func (t CRUD2Table) RowClassNameExpr(value string) CRUD2Table {
	return t.set("rowClassNameExpr", value)
}

// RowSelection 设置表格可选择配置
func (t CRUD2Table) RowSelection(value string) CRUD2Table {
	return t.set("rowSelection", value)
}

// SaveOrderApi 设置保存排序的 API
func (t CRUD2Table) SaveOrderApi(value string) CRUD2Table {
	return t.set("saveOrderApi", value)
}

// Selectable 设置是否可以选择数据
func (t CRUD2Table) Selectable(value bool) CRUD2Table {
	return t.set("selectable", value)
}

// ShowBadge 设置是否展示行角标
func (t CRUD2Table) ShowBadge(value bool) CRUD2Table {
	return t.set("showBadge", value)
}

// ShowHeader 设置是否展示表头
func (t CRUD2Table) ShowHeader(value bool) CRUD2Table {
	return t.set("showHeader", value)
}

// ShowSelection 设置是否展示已选数据区域
func (t CRUD2Table) ShowSelection(value bool) CRUD2Table {
	return t.set("showSelection", value)
}

// SilentPolling 设置静默拉取
func (t CRUD2Table) SilentPolling(value bool) CRUD2Table {
	return t.set("silentPolling", value)
}

// Source 设置数据源
func (t CRUD2Table) Source(value string) CRUD2Table {
	return t.set("source", value)
}

// Static 设置是否静态展示
func (t CRUD2Table) Static(value bool) CRUD2Table {
	return t.set("static", value)
}

// StaticClassName 设置静态展示表单项类名
func (t CRUD2Table) StaticClassName(value string) CRUD2Table {
	return t.set("staticClassName", value)
}

// StaticInputClassName 设置静态展示表单项 Value 类名
func (t CRUD2Table) StaticInputClassName(value string) CRUD2Table {
	return t.set("staticInputClassName", value)
}

// StaticLabelClassName 设置静态展示表单项 Label 类名
func (t CRUD2Table) StaticLabelClassName(value string) CRUD2Table {
	return t.set("staticLabelClassName", value)
}

// StaticOn 设置静态展示表达式
func (t CRUD2Table) StaticOn(value string) CRUD2Table {
	return t.set("staticOn", value)
}

// StaticPlaceholder 设置静态展示空值占位
func (t CRUD2Table) StaticPlaceholder(value string) CRUD2Table {
	return t.set("staticPlaceholder", value)
}

// StaticSchema 设置静态展示模式
func (t CRUD2Table) StaticSchema(value string) CRUD2Table {
	return t.set("staticSchema", value)
}

// Sticky 设置是否粘性头部
func (t CRUD2Table) Sticky(value bool) CRUD2Table {
	return t.set("sticky", value)
}

// StopAutoRefreshWhen 设置停止自动刷新条件
func (t CRUD2Table) StopAutoRefreshWhen(value string) CRUD2Table {
	return t.set("stopAutoRefreshWhen", value)
}

// Style 设置组件样式
func (t CRUD2Table) Style(value string) CRUD2Table {
	return t.set("style", value)
}

// SyncLocation 设置是否将过滤条件同步到地址栏
func (t CRUD2Table) SyncLocation(value bool) CRUD2Table {
	return t.set("syncLocation", value)
}

// SyncResponse2Query 设置是否将接口返回内容自动同步到地址栏
func (t CRUD2Table) SyncResponse2Query(value bool) CRUD2Table {
	return t.set("syncResponse2Query", value)
}

// TableLayout 设置表格布局
func (t CRUD2Table) TableLayout(value string) CRUD2Table {
	return t.set("tableLayout", value)
}

// TestIdBuilder 设置测试 ID 构建器
func (t CRUD2Table) TestIdBuilder(value string) CRUD2Table {
	return t.set("testIdBuilder", value)
}

// Testid 设置测试 ID
func (t CRUD2Table) Testid(value string) CRUD2Table {
	return t.set("testid", value)
}

// Title 设置表格标题
func (t CRUD2Table) Title(value string) CRUD2Table {
	return t.set("title", value)
}

// UseMobileUI 设置是否使用移动端样式
func (t CRUD2Table) UseMobileUI(value bool) CRUD2Table {
	return t.set("useMobileUI", value)
}

// Visible 设置是否显示
func (t CRUD2Table) Visible(value bool) CRUD2Table {
	return t.set("visible", value)
}

// VisibleOn 设置显示表达式
func (t CRUD2Table) VisibleOn(value string) CRUD2Table {
	return t.set("visibleOn", value)
}
