package comp

// CRUDCards 定义了 CRUDCards 渲染器
type CRUDCards Schema

// NewCRUDCards 创建一个新的 CRUDCards 实例
func NewCRUDCards() CRUDCards {
	return make(CRUDCards).set("mode", "cards").set("type", "crud")
}

func (c CRUDCards) set(key string, value interface{}) CRUDCards {
	c[key] = value
	return c
}

// AffixFooter 设置是否固底
func (c CRUDCards) AffixFooter(value bool) CRUDCards {
	return c.set("affixFooter", value)
}

// AffixHeader 设置是否固顶
func (c CRUDCards) AffixHeader(value bool) CRUDCards {
	return c.set("affixHeader", value)
}

// AlwaysShowPagination 设置是否总是显示分页
func (c CRUDCards) AlwaysShowPagination(value bool) CRUDCards {
	return c.set("alwaysShowPagination", value)
}

// Api 设置初始化数据 API
func (c CRUDCards) Api(value string) CRUDCards {
	return c.set("api", value)
}

// AutoFillHeight 设置内容区域是否占满屏幕剩余空间
func (c CRUDCards) AutoFillHeight(value string) CRUDCards {
	return c.set("autoFillHeight", value)
}

// AutoGenerateFilter 开启自动生成查询条件表单
func (c CRUDCards) AutoGenerateFilter(value string) CRUDCards {
	return c.set("autoGenerateFilter", value)
}

// AutoJumpToTopOnPagerChange 设置是否自动跳到顶部
func (c CRUDCards) AutoJumpToTopOnPagerChange(value bool) CRUDCards {
	return c.set("autoJumpToTopOnPagerChange", value)
}

// BulkActions 设置批量操作
func (c CRUDCards) BulkActions(value string) CRUDCards {
	return c.set("bulkActions", value)
}

// Card 设置卡片配置
func (c CRUDCards) Card(value string) CRUDCards {
	return c.set("card", value)
}

// CheckOnItemClick 设置点击卡片时是否勾选卡片
func (c CRUDCards) CheckOnItemClick(value bool) CRUDCards {
	return c.set("checkOnItemClick", value)
}

// ClassName 设置容器 css 类名
func (c CRUDCards) ClassName(value string) CRUDCards {
	return c.set("className", value)
}

// DefaultParams 设置默认参数
func (c CRUDCards) DefaultParams(value string) CRUDCards {
	return c.set("defaultParams", value)
}

// DeferApi 设置懒加载 API
func (c CRUDCards) DeferApi(value string) CRUDCards {
	return c.set("deferApi", value)
}

// Disabled 设置是否禁用
func (c CRUDCards) Disabled(value bool) CRUDCards {
	return c.set("disabled", value)
}

// DisabledOn 设置禁用表达式
func (c CRUDCards) DisabledOn(value string) CRUDCards {
	return c.set("disabledOn", value)
}

// Draggable 设置是否可通过拖拽排序
func (c CRUDCards) Draggable(value bool) CRUDCards {
	return c.set("draggable", value)
}

// DraggableOn 设置拖拽排序表达式
func (c CRUDCards) DraggableOn(value string) CRUDCards {
	return c.set("draggableOn", value)
}

// EditorSetting 设置编辑器配置
func (c CRUDCards) EditorSetting(value string) CRUDCards {
	return c.set("editorSetting", value)
}

// ExpandConfig 设置展开配置
func (c CRUDCards) ExpandConfig(value string) CRUDCards {
	return c.set("expandConfig", value)
}

// Filter 设置过滤器表单
func (c CRUDCards) Filter(value string) CRUDCards {
	return c.set("filter", value)
}

// FilterDefaultVisible 设置过滤器默认是否可见
func (c CRUDCards) FilterDefaultVisible(value bool) CRUDCards {
	return c.set("filterDefaultVisible", value)
}

// FilterTogglable 设置过滤器是否可切换
func (c CRUDCards) FilterTogglable(value bool) CRUDCards {
	return c.set("filterTogglable", value)
}

// Footer 设置底部区域
func (c CRUDCards) Footer(value string) CRUDCards {
	return c.set("footer", value)
}

// FooterClassName 设置底部 CSS 类名
func (c CRUDCards) FooterClassName(value string) CRUDCards {
	return c.set("footerClassName", value)
}

// FooterToolbar 设置底部工具栏
func (c CRUDCards) FooterToolbar(value string) CRUDCards {
	return c.set("footerToolbar", value)
}

// Header 设置顶部区域
func (c CRUDCards) Header(value string) CRUDCards {
	return c.set("header", value)
}

// HeaderClassName 设置顶部 CSS 类名
func (c CRUDCards) HeaderClassName(value string) CRUDCards {
	return c.set("headerClassName", value)
}

// HeaderToolbar 设置顶部工具栏
func (c CRUDCards) HeaderToolbar(value string) CRUDCards {
	return c.set("headerToolbar", value)
}

// Hidden 设置是否隐藏
func (c CRUDCards) Hidden(value bool) CRUDCards {
	return c.set("hidden", value)
}

// HiddenOn 设置隐藏表达式
func (c CRUDCards) HiddenOn(value string) CRUDCards {
	return c.set("hiddenOn", value)
}

// HideCheckToggler 设置是否隐藏勾选框
func (c CRUDCards) HideCheckToggler(value bool) CRUDCards {
	return c.set("hideCheckToggler", value)
}

// HideQuickSaveBtn 设置是否隐藏快速保存按钮
func (c CRUDCards) HideQuickSaveBtn(value bool) CRUDCards {
	return c.set("hideQuickSaveBtn", value)
}

// Id 设置组件唯一 ID
func (c CRUDCards) Id(value string) CRUDCards {
	return c.set("id", value)
}

// InitFetch 设置初始是否拉取
func (c CRUDCards) InitFetch(value bool) CRUDCards {
	return c.set("initFetch", value)
}

// InitFetchOn 设置初始拉取表达式
func (c CRUDCards) InitFetchOn(value string) CRUDCards {
	return c.set("initFetchOn", value)
}

// InnerClassName 设置内部 DOM 的 CSS 类名
func (c CRUDCards) InnerClassName(value string) CRUDCards {
	return c.set("innerClassName", value)
}

// Interval 设置自动刷新时间
func (c CRUDCards) Interval(value string) CRUDCards {
	return c.set("interval", value)
}

// ItemActions 设置单条操作
func (c CRUDCards) ItemActions(value string) CRUDCards {
	return c.set("itemActions", value)
}

// ItemCheckableOn 设置约束批量操作
func (c CRUDCards) ItemCheckableOn(value string) CRUDCards {
	return c.set("itemCheckableOn", value)
}

// ItemClassName 设置卡片 CSS 类名
func (c CRUDCards) ItemClassName(value string) CRUDCards {
	return c.set("itemClassName", value)
}

// ItemDraggableOn 设置项是否可拖拽排序
func (c CRUDCards) ItemDraggableOn(value string) CRUDCards {
	return c.set("itemDraggableOn", value)
}

// KeepItemSelectionOnPageChange 设置分页时是否保留用户选择
func (c CRUDCards) KeepItemSelectionOnPageChange(value bool) CRUDCards {
	return c.set("keepItemSelectionOnPageChange", value)
}

// LabelTpl 设置已勾选项的文案
func (c CRUDCards) LabelTpl(value string) CRUDCards {
	return c.set("labelTpl", value)
}

// LoadDataOnce 设置是否为前端单次加载模式
func (c CRUDCards) LoadDataOnce(value bool) CRUDCards {
	return c.set("loadDataOnce", value)
}

// LoadDataOnceFetchOnFilter 设置 loadDataOnce 时的过滤条件
func (c CRUDCards) LoadDataOnceFetchOnFilter(value bool) CRUDCards {
	return c.set("loadDataOnceFetchOnFilter", value)
}

// LoadingConfig 设置加载配置
func (c CRUDCards) LoadingConfig(value string) CRUDCards {
	return c.set("loadingConfig", value)
}

// MasonryLayout 设置是否为瀑布流布局
func (c CRUDCards) MasonryLayout(value bool) CRUDCards {
	return c.set("masonryLayout", value)
}

// MatchFunc 设置自定义搜索匹配函数
func (c CRUDCards) MatchFunc(value string) CRUDCards {
	return c.set("matchFunc", value)
}

// Messages 设置消息文案配置
func (c CRUDCards) Messages(value string) CRUDCards {
	return c.set("messages", value)
}

// Mode 设置展示模式
func (c CRUDCards) Mode(value string) CRUDCards {
	return c.set("mode", value)
}

// Name 设置组件名字
func (c CRUDCards) Name(value string) CRUDCards {
	return c.set("name", value)
}

// OnEvent 设置事件动作配置
func (c CRUDCards) OnEvent(value string) CRUDCards {
	return c.set("onEvent", value)
}

// OrderBy 设置默认排序字段
func (c CRUDCards) OrderBy(value string) CRUDCards {
	return c.set("orderBy", value)
}

// PageField 设置分页字段名
func (c CRUDCards) PageField(value string) CRUDCards {
	return c.set("pageField", value)
}

// Pagination 设置分页
func (c CRUDCards) Pagination(value string) CRUDCards {
	return c.set("pagination", value)
}

// PerPage 设置每页显示条数
func (c CRUDCards) PerPage(value int) CRUDCards {
	return c.set("perPage", value)
}

// QuickSaveApi 设置快速保存 API
func (c CRUDCards) QuickSaveApi(value string) CRUDCards {
	return c.set("quickSaveApi", value)
}

// ReadOnly 设置只读
func (c CRUDCards) ReadOnly(value bool) CRUDCards {
	return c.set("readOnly", value)
}

// ResetPageOnReload 设置重载时是否重置分页
func (c CRUDCards) ResetPageOnReload(value bool) CRUDCards {
	return c.set("resetPageOnReload", value)
}

// SearchSettings 设置搜索配置
func (c CRUDCards) SearchSettings(value string) CRUDCards {
	return c.set("searchSettings", value)
}

// Selectable 设置是否可选择
func (c CRUDCards) Selectable(value bool) CRUDCards {
	return c.set("selectable", value)
}

// ShowPagination 设置是否显示分页
func (c CRUDCards) ShowPagination(value bool) CRUDCards {
	return c.set("showPagination", value)
}

// ShowQuickSaveBtn 设置是否显示快速保存按钮
func (c CRUDCards) ShowQuickSaveBtn(value bool) CRUDCards {
	return c.set("showQuickSaveBtn", value)
}

// ShowTotal 设置是否显示总条数
func (c CRUDCards) ShowTotal(value bool) CRUDCards {
	return c.set("showTotal", value)
}

// Source 设置数据源
func (c CRUDCards) Source(value string) CRUDCards {
	return c.set("source", value)
}

// SubItemSettings 设置子项配置
func (c CRUDCards) SubItemSettings(value string) CRUDCards {
	return c.set("subItemSettings", value)
}

// SyncLocation 设置是否同步浏览器位置
func (c CRUDCards) SyncLocation(value bool) CRUDCards {
	return c.set("syncLocation", value)
}

// Title 设置标题
func (c CRUDCards) Title(value string) CRUDCards {
	return c.set("title", value)
}

// Toolbar 设置工具栏
func (c CRUDCards) Toolbar(value string) CRUDCards {
	return c.set("toolbar", value)
}

// ToolbarClassName 设置工具栏 CSS 类名
func (c CRUDCards) ToolbarClassName(value string) CRUDCards {
	return c.set("toolbarClassName", value)
}

// Tooltip 设置工具提示
func (c CRUDCards) Tooltip(value string) CRUDCards {
	return c.set("tooltip", value)
}

// UseMobile 设置是否使用移动端适配
func (c CRUDCards) UseMobile(value bool) CRUDCards {
	return c.set("useMobile", value)
}

// ValueField 设置值字段名
func (c CRUDCards) ValueField(value string) CRUDCards {
	return c.set("valueField", value)
}

// VisibleOn 设置可见表达式
func (c CRUDCards) VisibleOn(value string) CRUDCards {
	return c.set("visibleOn", value)
}
