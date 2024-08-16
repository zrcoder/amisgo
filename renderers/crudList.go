package renderers

// CRUDList 表示 CRUD 列表渲染器
type CRUDList struct {
	*BaseRenderer
}

// NewCRUDList 创建新的 CRUDList 实例
func NewCRUDList() *CRUDList {
	c := &CRUDList{
		BaseRenderer: NewBaseRenderer(),
	}
	c.set("mode", "list")
	c.set("type", "crud")
	return c
}

// Set 设置键值对，并返回当前实例
func (c *CRUDList) set(key string, value interface{}) *CRUDList {
	c.BaseRenderer.set(key, value)
	return c
}

// AffixFooter 设置是否固底
func (c *CRUDList) AffixFooter(value bool) *CRUDList {
	return c.set("affixFooter", value)
}

// AffixHeader 设置是否固顶
func (c *CRUDList) AffixHeader(value bool) *CRUDList {
	return c.set("affixHeader", value)
}

// AlwaysShowPagination 默认只显示分页数大于 1 时才显示，设置为 true 总是显示
func (c *CRUDList) AlwaysShowPagination(value bool) *CRUDList {
	return c.set("alwaysShowPagination", value)
}

// Api 设置初始化数据 API
func (c *CRUDList) Api(value string) *CRUDList {
	return c.set("api", value)
}

// AutoFillHeight 设置内容区域是否占满屏幕剩余空间
func (c *CRUDList) AutoFillHeight(value string) *CRUDList {
	return c.set("autoFillHeight", value)
}

// AutoGenerateFilter 开启查询区域，会根据列元素的searchable属性值，自动生成查询条件表单
func (c *CRUDList) AutoGenerateFilter(value string) *CRUDList {
	return c.set("autoGenerateFilter", value)
}

// AutoJumpToTopOnPagerChange 设置分页时是否自动跳到顶部
func (c *CRUDList) AutoJumpToTopOnPagerChange(value bool) *CRUDList {
	return c.set("autoJumpToTopOnPagerChange", value)
}

// BulkActions 设置批量操作
func (c *CRUDList) BulkActions(value string) *CRUDList {
	return c.set("bulkActions", value)
}

// CheckOnItemClick 设置点击列表单行时是否选择
func (c *CRUDList) CheckOnItemClick(value bool) *CRUDList {
	return c.set("checkOnItemClick", value)
}

// ClassName 设置容器 CSS 类名
func (c *CRUDList) ClassName(value string) *CRUDList {
	return c.set("className", value)
}

// DefaultParams 设置默认初始参数
func (c *CRUDList) DefaultParams(value string) *CRUDList {
	return c.set("defaultParams", value)
}

// DeferApi 设置懒加载 API
func (c *CRUDList) DeferApi(value string) *CRUDList {
	return c.set("deferApi", value)
}

// Disabled 设置是否禁用
func (c *CRUDList) Disabled(value bool) *CRUDList {
	return c.set("disabled", value)
}

// DisabledOn 设置禁用表达式
func (c *CRUDList) DisabledOn(value string) *CRUDList {
	return c.set("disabledOn", value)
}

// Draggable 设置是否可通过拖拽排序
func (c *CRUDList) Draggable(value bool) *CRUDList {
	return c.set("draggable", value)
}

// DraggableOn 设置可拖拽排序表达式
func (c *CRUDList) DraggableOn(value string) *CRUDList {
	return c.set("draggableOn", value)
}

// EditorSetting 设置编辑器配置
func (c *CRUDList) EditorSetting(value string) *CRUDList {
	return c.set("editorSetting", value)
}

// ExpandConfig 设置内嵌模式的默认展开选项
func (c *CRUDList) ExpandConfig(value string) *CRUDList {
	return c.set("expandConfig", value)
}

// Filter 设置过滤器表单
func (c *CRUDList) Filter(value string) *CRUDList {
	return c.set("filter", value)
}

// FilterDefaultVisible 设置过滤器默认是否可见
func (c *CRUDList) FilterDefaultVisible(value bool) *CRUDList {
	return c.set("filterDefaultVisible", value)
}

// FilterTogglable 设置过滤器是否可切换
func (c *CRUDList) FilterTogglable(value bool) *CRUDList {
	return c.set("filterTogglable", value)
}

// Footer 设置底部区域
func (c *CRUDList) Footer(value string) *CRUDList {
	return c.set("footer", value)
}

// FooterClassName 设置底部区域类名
func (c *CRUDList) FooterClassName(value string) *CRUDList {
	return c.set("footerClassName", value)
}

// FooterToolbar 设置底部工具栏
func (c *CRUDList) FooterToolbar(value string) *CRUDList {
	return c.set("footerToolbar", value)
}

// Header 设置顶部区域
func (c *CRUDList) Header(value string) *CRUDList {
	return c.set("header", value)
}

// HeaderClassName 设置顶部区域类名
func (c *CRUDList) HeaderClassName(value string) *CRUDList {
	return c.set("headerClassName", value)
}

// HeaderToolbar 设置顶部工具栏
func (c *CRUDList) HeaderToolbar(value string) *CRUDList {
	return c.set("headerToolbar", value)
}

// Hidden 设置是否隐藏
func (c *CRUDList) Hidden(value bool) *CRUDList {
	return c.set("hidden", value)
}

// HiddenOn 设置隐藏表达式
func (c *CRUDList) HiddenOn(value string) *CRUDList {
	return c.set("hiddenOn", value)
}

// HideCheckToggler 设置是否隐藏勾选框
func (c *CRUDList) HideCheckToggler(value bool) *CRUDList {
	return c.set("hideCheckToggler", value)
}

// HideQuickSaveBtn 设置是否隐藏快速编辑的按钮
func (c *CRUDList) HideQuickSaveBtn(value bool) *CRUDList {
	return c.set("hideQuickSaveBtn", value)
}

// Id 设置组件唯一 ID
func (c *CRUDList) Id(value string) *CRUDList {
	return c.set("id", value)
}

// InitFetch 设置初始是否拉取
func (c *CRUDList) InitFetch(value bool) *CRUDList {
	return c.set("initFetch", value)
}

// InitFetchOn 设置初始是否拉取，用表达式来配置
func (c *CRUDList) InitFetchOn(value string) *CRUDList {
	return c.set("initFetchOn", value)
}

// InnerClassName 设置内部 DOM 的 className
func (c *CRUDList) InnerClassName(value string) *CRUDList {
	return c.set("innerClassName", value)
}

// Interval 设置自动刷新时间
func (c *CRUDList) Interval(value string) *CRUDList {
	return c.set("interval", value)
}

// ItemAction 设置点击列表项的行为
func (c *CRUDList) ItemAction(value string) *CRUDList {
	return c.set("itemAction", value)
}

// ItemActions 设置单条操作
func (c *CRUDList) ItemActions(value string) *CRUDList {
	return c.set("itemActions", value)
}

// ItemCheckableOn 设置约束批量操作
func (c *CRUDList) ItemCheckableOn(value string) *CRUDList {
	return c.set("itemCheckableOn", value)
}

// ItemDraggableOn 设置配置某项是否可拖拽排序
func (c *CRUDList) ItemDraggableOn(value string) *CRUDList {
	return c.set("itemDraggableOn", value)
}

// KeepItemSelectionOnPageChange 设置分页时是否保留用户选择
func (c *CRUDList) KeepItemSelectionOnPageChange(value bool) *CRUDList {
	return c.set("keepItemSelectionOnPageChange", value)
}

// LabelTpl 设置已勾选项的文案
func (c *CRUDList) LabelTpl(value string) *CRUDList {
	return c.set("labelTpl", value)
}

// ListItem 设置单条数据展示内容配置
func (c *CRUDList) ListItem(value string) *CRUDList {
	return c.set("listItem", value)
}

// LoadDataOnce 设置是否为前端单次加载模式
func (c *CRUDList) LoadDataOnce(value bool) *CRUDList {
	return c.set("loadDataOnce", value)
}

// LoadDataOnceFetchOnFilter 设置在开启 loadDataOnce 时，当修改过滤条件时是否重新请求 API
func (c *CRUDList) LoadDataOnceFetchOnFilter(value bool) *CRUDList {
	return c.set("loadDataOnceFetchOnFilter", value)
}

// LoadingConfig 设置加载配置
func (c *CRUDList) LoadingConfig(value string) *CRUDList {
	return c.set("loadingConfig", value)
}

// MatchFunc 设置自定义搜索匹配函数
func (c *CRUDList) MatchFunc(value string) *CRUDList {
	return c.set("matchFunc", value)
}

// Messages 设置消息文案配置
func (c *CRUDList) Messages(value string) *CRUDList {
	return c.set("messages", value)
}

// Mode 设置内容区展示模式
func (c *CRUDList) Mode(value string) *CRUDList {
	return c.set("mode", value)
}

// Name 设置组件名字
func (c *CRUDList) Name(value string) *CRUDList {
	return c.set("name", value)
}

// OnEvent 设置事件动作配置
func (c *CRUDList) OnEvent(value string) *CRUDList {
	return c.set("onEvent", value)
}

// PageField 设置分页参数字段
func (c *CRUDList) PageField(value string) *CRUDList {
	return c.set("pageField", value)
}

// PageFieldName 设置分页字段名称
func (c *CRUDList) PageFieldName(value string) *CRUDList {
	return c.set("pageFieldName", value)
}

// PerPage 设置每页条数
func (c *CRUDList) PerPage(value int) *CRUDList {
	return c.set("perPage", value)
}

// QuickSave 设置快速编辑配置
func (c *CRUDList) QuickSave(value string) *CRUDList {
	return c.set("quickSave", value)
}

// RefreshInterval 设置自动刷新时间间隔
func (c *CRUDList) RefreshInterval(value string) *CRUDList {
	return c.set("refreshInterval", value)
}

// Removable 设置是否可删除
func (c *CRUDList) Removable(value bool) *CRUDList {
	return c.set("removable", value)
}

// SaveOrder 设置排序信息保存
func (c *CRUDList) SaveOrder(value bool) *CRUDList {
	return c.set("saveOrder", value)
}

// SaveOrderApi 设置排序保存 API
func (c *CRUDList) SaveOrderApi(value string) *CRUDList {
	return c.set("saveOrderApi", value)
}

// Searchable 设置是否可搜索
func (c *CRUDList) Searchable(value bool) *CRUDList {
	return c.set("searchable", value)
}

// SimpleMode 设置简洁模式
func (c *CRUDList) SimpleMode(value bool) *CRUDList {
	return c.set("simpleMode", value)
}

// Size 设置组件尺寸
func (c *CRUDList) Size(value string) *CRUDList {
	return c.set("size", value)
}

// Stick 设置是否固定
func (c *CRUDList) Stick(value bool) *CRUDList {
	return c.set("stick", value)
}

// StickOn 设置固定条件
func (c *CRUDList) StickOn(value string) *CRUDList {
	return c.set("stickOn", value)
}

// Toolbar 设置工具栏
func (c *CRUDList) Toolbar(value string) *CRUDList {
	return c.set("toolbar", value)
}

// ToolbarClassName 设置工具栏类名
func (c *CRUDList) ToolbarClassName(value string) *CRUDList {
	return c.set("toolbarClassName", value)
}

// ToolbarPosition 设置工具栏位置
func (c *CRUDList) ToolbarPosition(value string) *CRUDList {
	return c.set("toolbarPosition", value)
}

// Unstable 设置是否不稳定（不推荐）
func (c *CRUDList) Unstable(value bool) *CRUDList {
	return c.set("unstable", value)
}

// UseMobile 设置是否使用移动端样式
func (c *CRUDList) UseMobile(value bool) *CRUDList {
	return c.set("useMobile", value)
}

// ValueField 设置值字段
func (c *CRUDList) ValueField(value string) *CRUDList {
	return c.set("valueField", value)
}

// ValueFieldOn 设置值字段表达式
func (c *CRUDList) ValueFieldOn(value string) *CRUDList {
	return c.set("valueFieldOn", value)
}
