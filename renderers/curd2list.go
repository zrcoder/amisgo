package renderers

// CRUD2List 是 CRUD2 列表渲染器
type CRUD2List struct {
	*BaseRenderer
}

// NewCRUD2List 创建一个新的 CRUD2List 实例
func NewCRUD2List() *CRUD2List {
	list := &CRUD2List{BaseRenderer: NewBaseRenderer()}
	list.set("mode", "list")
	list.set("type", "crud2")
	return list
}

// Set 设置属性
func (l *CRUD2List) set(key string, value interface{}) *CRUD2List {
	l.BaseRenderer.set(key, value)
	return l
}

// AffixFooter 设置是否固底
func (l *CRUD2List) AffixFooter(value bool) *CRUD2List {
	return l.set("affixFooter", value)
}

// AffixHeader 设置是否固顶
func (l *CRUD2List) AffixHeader(value bool) *CRUD2List {
	return l.set("affixHeader", value)
}

// API 设置初始化数据 API
func (l *CRUD2List) API(value string) *CRUD2List {
	return l.set("api", value)
}

// AutoFillHeight 设置内容区域占满屏幕剩余空间
func (l *CRUD2List) AutoFillHeight(value bool) *CRUD2List {
	return l.set("autoFillHeight", value)
}

// AutoJumpToTopOnPagerChange 设置是否自动跳顶部，当切分页的时候
func (l *CRUD2List) AutoJumpToTopOnPagerChange(value bool) *CRUD2List {
	return l.set("autoJumpToTopOnPagerChange", value)
}

// CheckOnItemClick 设置点击列表单行时是否选择
func (l *CRUD2List) CheckOnItemClick(value bool) *CRUD2List {
	return l.set("checkOnItemClick", value)
}

// ClassName 设置容器 CSS 类名
func (l *CRUD2List) ClassName(value string) *CRUD2List {
	return l.set("className", value)
}

// Disabled 设置是否禁用
func (l *CRUD2List) Disabled(value bool) *CRUD2List {
	return l.set("disabled", value)
}

// DisabledOn 设置是否禁用表达式
func (l *CRUD2List) DisabledOn(value string) *CRUD2List {
	return l.set("disabledOn", value)
}

// EditorSetting 设置编辑器配置
func (l *CRUD2List) EditorSetting(value string) *CRUD2List {
	return l.set("editorSetting", value)
}

// Footer 设置底部区域
func (l *CRUD2List) Footer(value interface{}) *CRUD2List {
	return l.set("footer", value)
}

// FooterClassName 设置底部区域 CSS 类名
func (l *CRUD2List) FooterClassName(value string) *CRUD2List {
	return l.set("footerClassName", value)
}

// FooterToolbar 设置底部工具栏区域
func (l *CRUD2List) FooterToolbar(value interface{}) *CRUD2List {
	return l.set("footerToolbar", value)
}

// FooterToolbarClassName 设置底部工具栏 CSS 类名
func (l *CRUD2List) FooterToolbarClassName(value string) *CRUD2List {
	return l.set("footerToolbarClassName", value)
}

// Header 设置顶部区域
func (l *CRUD2List) Header(value interface{}) *CRUD2List {
	return l.set("header", value)
}

// HeaderClassName 设置顶部区域 CSS 类名
func (l *CRUD2List) HeaderClassName(value string) *CRUD2List {
	return l.set("headerClassName", value)
}

// HeaderToolbar 设置顶部工具栏区域
func (l *CRUD2List) HeaderToolbar(value interface{}) *CRUD2List {
	return l.set("headerToolbar", value)
}

// HeaderToolbarClassName 设置顶部工具栏 CSS 类名
func (l *CRUD2List) HeaderToolbarClassName(value string) *CRUD2List {
	return l.set("headerToolbarClassName", value)
}

// Hidden 设置是否隐藏
func (l *CRUD2List) Hidden(value bool) *CRUD2List {
	return l.set("hidden", value)
}

// HiddenOn 设置是否隐藏表达式
func (l *CRUD2List) HiddenOn(value string) *CRUD2List {
	return l.set("hiddenOn", value)
}

// HideCheckToggler 设置是否隐藏勾选框
func (l *CRUD2List) HideCheckToggler(value bool) *CRUD2List {
	return l.set("hideCheckToggler", value)
}

// HideQuickSaveBtn 设置是否隐藏快速编辑按钮
func (l *CRUD2List) HideQuickSaveBtn(value bool) *CRUD2List {
	return l.set("hideQuickSaveBtn", value)
}

// ID 设置组件唯一 ID
func (l *CRUD2List) ID(value string) *CRUD2List {
	return l.set("id", value)
}

// Interval 设置自动刷新时间
func (l *CRUD2List) Interval(value string) *CRUD2List {
	return l.set("interval", value)
}

// ItemAction 设置点击列表项的行为
func (l *CRUD2List) ItemAction(value string) *CRUD2List {
	return l.set("itemAction", value)
}

// ItemCheckableOn 设置约束批量操作
func (l *CRUD2List) ItemCheckableOn(value string) *CRUD2List {
	return l.set("itemCheckableOn", value)
}

// ItemDraggableOn 配置某项是否可拖拽排序
func (l *CRUD2List) ItemDraggableOn(value string) *CRUD2List {
	return l.set("itemDraggableOn", value)
}

// KeepItemSelectionOnPageChange 翻页时是否保留用户已选的数据
func (l *CRUD2List) KeepItemSelectionOnPageChange(value bool) *CRUD2List {
	return l.set("keepItemSelectionOnPageChange", value)
}

// ListItem 单条数据展示内容配置
func (l *CRUD2List) ListItem(value interface{}) *CRUD2List {
	return l.set("listItem", value)
}

// LoadDataOnce 是否为前端单次加载模式
func (l *CRUD2List) LoadDataOnce(value bool) *CRUD2List {
	return l.set("loadDataOnce", value)
}

// LoadType 数据展示模式
func (l *CRUD2List) LoadType(value string) *CRUD2List {
	return l.set("loadType", value)
}

// LoadingConfig 设置加载配置
func (l *CRUD2List) LoadingConfig(value string) *CRUD2List {
	return l.set("loadingConfig", value)
}

// Mode 设置指定内容区的展示模式
func (l *CRUD2List) Mode(value string) *CRUD2List {
	return l.set("mode", value)
}

// Multiple 设置是否可以多选数据
func (l *CRUD2List) Multiple(value bool) *CRUD2List {
	return l.set("multiple", value)
}

// Name 设置组件名字
func (l *CRUD2List) Name(value string) *CRUD2List {
	return l.set("name", value)
}

// OnEvent 设置事件动作配置
func (l *CRUD2List) OnEvent(value string) *CRUD2List {
	return l.set("onEvent", value)
}

// PageField 设置分页页码字段名
func (l *CRUD2List) PageField(value string) *CRUD2List {
	return l.set("pageField", value)
}

// ParsePrimitiveQuery 设置是否开启 Query 信息转换
func (l *CRUD2List) ParsePrimitiveQuery(value string) *CRUD2List {
	return l.set("parsePrimitiveQuery", value)
}

// PerPage 设置无限加载时每页加载数量
func (l *CRUD2List) PerPage(value string) *CRUD2List {
	return l.set("perPage", value)
}

// PerPageField 设置分页每页显示的字段名
func (l *CRUD2List) PerPageField(value string) *CRUD2List {
	return l.set("perPageField", value)
}

// Placeholder 设置无数据提示
func (l *CRUD2List) Placeholder(value string) *CRUD2List {
	return l.set("placeholder", value)
}

// PrimaryField 设置行标识符
func (l *CRUD2List) PrimaryField(value string) *CRUD2List {
	return l.set("primaryField", value)
}

// QuickSaveApi 设置快速编辑后批量保存的 API
func (l *CRUD2List) QuickSaveApi(value string) *CRUD2List {
	return l.set("quickSaveApi", value)
}

// QuickSaveItemApi 设置快速编辑配置成及时保存时使用的 API
func (l *CRUD2List) QuickSaveItemApi(value string) *CRUD2List {
	return l.set("quickSaveItemApi", value)
}

// SaveOrderApi 设置保存排序的 API
func (l *CRUD2List) SaveOrderApi(value string) *CRUD2List {
	return l.set("saveOrderApi", value)
}

// Selectable 设置是否可以选择数据
func (l *CRUD2List) Selectable(value bool) *CRUD2List {
	return l.set("selectable", value)
}

// ShowFooter 设置是否显示底部
func (l *CRUD2List) ShowFooter(value bool) *CRUD2List {
	return l.set("showFooter", value)
}

// ShowHeader 设置是否显示头部
func (l *CRUD2List) ShowHeader(value bool) *CRUD2List {
	return l.set("showHeader", value)
}

// ShowSelection 设置是否展示已选数据区域
func (l *CRUD2List) ShowSelection(value bool) *CRUD2List {
	return l.set("showSelection", value)
}

// SilentPolling 设置静默拉取
func (l *CRUD2List) SilentPolling(value bool) *CRUD2List {
	return l.set("silentPolling", value)
}

// Size 设置大小
func (l *CRUD2List) Size(value string) *CRUD2List {
	return l.set("size", value)
}

// Source 设置数据来源
func (l *CRUD2List) Source(value string) *CRUD2List {
	return l.set("source", value)
}

// Static 设置是否静态展示
func (l *CRUD2List) Static(value bool) *CRUD2List {
	return l.set("static", value)
}

// StaticClassName 设置静态展示表单项类名
func (l *CRUD2List) StaticClassName(value string) *CRUD2List {
	return l.set("staticClassName", value)
}

// StaticInputClassName 设置静态展示表单项 Value 类名
func (l *CRUD2List) StaticInputClassName(value string) *CRUD2List {
	return l.set("staticInputClassName", value)
}

// StaticLabelClassName 设置静态展示表单项 Label 类名
func (l *CRUD2List) StaticLabelClassName(value string) *CRUD2List {
	return l.set("staticLabelClassName", value)
}

// StaticOn 设置是否静态展示表达式
func (l *CRUD2List) StaticOn(value string) *CRUD2List {
	return l.set("staticOn", value)
}

// StaticPlaceholder 设置静态展示空值占位
func (l *CRUD2List) StaticPlaceholder(value string) *CRUD2List {
	return l.set("staticPlaceholder", value)
}

// StaticSchema 设置静态展示 Schema
func (l *CRUD2List) StaticSchema(value string) *CRUD2List {
	return l.set("staticSchema", value)
}

// StopAutoRefreshWhen 设置停止自动刷新的条件
func (l *CRUD2List) StopAutoRefreshWhen(value string) *CRUD2List {
	return l.set("stopAutoRefreshWhen", value)
}

// Style 设置组件样式
func (l *CRUD2List) Style(value string) *CRUD2List {
	return l.set("style", value)
}

// SyncLocation 设置是否将过滤条件的参数同步到地址栏
func (l *CRUD2List) SyncLocation(value bool) *CRUD2List {
	return l.set("syncLocation", value)
}

// SyncResponse2Query 设置是否将接口返回的内容自动同步到地址栏
func (l *CRUD2List) SyncResponse2Query(value bool) *CRUD2List {
	return l.set("syncResponse2Query", value)
}

// TestIdBuilder 设置测试 ID 构建器
func (l *CRUD2List) TestIdBuilder(value string) *CRUD2List {
	return l.set("testIdBuilder", value)
}

// Testid 设置测试 ID
func (l *CRUD2List) Testid(value string) *CRUD2List {
	return l.set("testid", value)
}

// Title 设置标题
func (l *CRUD2List) Title(value string) *CRUD2List {
	return l.set("title", value)
}

// UseMobileUI 设置是否可以组件级别用来关闭移动端样式
func (l *CRUD2List) UseMobileUI(value bool) *CRUD2List {
	return l.set("useMobileUI", value)
}

// ValueField 设置可以用来作为值的字段
func (l *CRUD2List) ValueField(value string) *CRUD2List {
	return l.set("valueField", value)
}

// Visible 设置是否显示
func (l *CRUD2List) Visible(value bool) *CRUD2List {
	return l.set("visible", value)
}

// VisibleOn 设置是否显示表达式
func (l *CRUD2List) VisibleOn(value string) *CRUD2List {
	return l.set("visibleOn", value)
}
