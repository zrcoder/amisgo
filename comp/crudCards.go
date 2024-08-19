package comp

// crudCards 定义了 crudCards 渲染器
type crudCards schema

// CrudCards 创建一个新的 CrudCards 实例
func CrudCards() crudCards {
	return make(crudCards).set("mode", "cards").set("type", "crud")
}

func (c crudCards) set(key string, value any) crudCards {
	c[key] = value
	return c
}

// AffixFooter 设置是否固底
func (c crudCards) AffixFooter(value bool) crudCards {
	return c.set("affixFooter", value)
}

// AffixHeader 设置是否固顶
func (c crudCards) AffixHeader(value bool) crudCards {
	return c.set("affixHeader", value)
}

// AlwaysShowPagination 设置是否总是显示分页
func (c crudCards) AlwaysShowPagination(value bool) crudCards {
	return c.set("alwaysShowPagination", value)
}

// Api 设置初始化数据 API
func (c crudCards) Api(value string) crudCards {
	return c.set("api", value)
}

// AutoFillHeight 设置内容区域是否占满屏幕剩余空间
func (c crudCards) AutoFillHeight(value string) crudCards {
	return c.set("autoFillHeight", value)
}

// AutoGenerateFilter 开启自动生成查询条件表单
func (c crudCards) AutoGenerateFilter(value string) crudCards {
	return c.set("autoGenerateFilter", value)
}

// AutoJumpToTopOnPagerChange 设置是否自动跳到顶部
func (c crudCards) AutoJumpToTopOnPagerChange(value bool) crudCards {
	return c.set("autoJumpToTopOnPagerChange", value)
}

// BulkActions 设置批量操作
func (c crudCards) BulkActions(value string) crudCards {
	return c.set("bulkActions", value)
}

// Card 设置卡片配置
func (c crudCards) Card(value any) crudCards {
	return c.set("card", value)
}

// CheckOnItemClick 设置点击卡片时是否勾选卡片
func (c crudCards) CheckOnItemClick(value bool) crudCards {
	return c.set("checkOnItemClick", value)
}

// ClassName 设置容器 CSS 类名
func (c crudCards) ClassName(value string) crudCards {
	return c.set("className", value)
}

// DefaultParams 设置默认参数
func (c crudCards) DefaultParams(value string) crudCards {
	return c.set("defaultParams", value)
}

// DeferApi 设置懒加载 API
func (c crudCards) DeferApi(value string) crudCards {
	return c.set("deferApi", value)
}

// Disabled 设置是否禁用
func (c crudCards) Disabled(value bool) crudCards {
	return c.set("disabled", value)
}

// DisabledOn 设置禁用表达式
func (c crudCards) DisabledOn(value string) crudCards {
	return c.set("disabledOn", value)
}

// Draggable 设置是否可通过拖拽排序
func (c crudCards) Draggable(value bool) crudCards {
	return c.set("draggable", value)
}

// DraggableOn 设置拖拽排序表达式
func (c crudCards) DraggableOn(value string) crudCards {
	return c.set("draggableOn", value)
}

// EditorSetting 设置编辑器配置
func (c crudCards) EditorSetting(value string) crudCards {
	return c.set("editorSetting", value)
}

// ExpandConfig 设置展开配置
func (c crudCards) ExpandConfig(value string) crudCards {
	return c.set("expandConfig", value)
}

// Filter 设置过滤器表单
func (c crudCards) Filter(value string) crudCards {
	return c.set("filter", value)
}

// FilterDefaultVisible 设置过滤器默认是否可见
func (c crudCards) FilterDefaultVisible(value bool) crudCards {
	return c.set("filterDefaultVisible", value)
}

// FilterTogglable 设置过滤器是否可切换
func (c crudCards) FilterTogglable(value bool) crudCards {
	return c.set("filterTogglable", value)
}

// Footer 设置底部区域
func (c crudCards) Footer(value any) crudCards {
	return c.set("footer", value)
}

// FooterClassName 设置底部 CSS 类名
func (c crudCards) FooterClassName(value string) crudCards {
	return c.set("footerClassName", value)
}

// FooterToolbar 设置底部工具栏
func (c crudCards) FooterToolbar(value any) crudCards {
	return c.set("footerToolbar", value)
}

// FooterToolbarClassName 设置底部工具栏 CSS 类名
func (c crudCards) FooterToolbarClassName(value string) crudCards {
	return c.set("footerToolbarClassName", value)
}

// Header 设置顶部区域
func (c crudCards) Header(value any) crudCards {
	return c.set("header", value)
}

// HeaderClassName 设置顶部 CSS 类名
func (c crudCards) HeaderClassName(value string) crudCards {
	return c.set("headerClassName", value)
}

// HeaderToolbar 设置顶部工具栏
func (c crudCards) HeaderToolbar(value any) crudCards {
	return c.set("headerToolbar", value)
}

// HeaderToolbarClassName 设置顶部工具栏 CSS 类名
func (c crudCards) HeaderToolbarClassName(value string) crudCards {
	return c.set("headerToolbarClassName", value)
}

// Hidden 设置是否隐藏
func (c crudCards) Hidden(value bool) crudCards {
	return c.set("hidden", value)
}

// HiddenOn 设置隐藏表达式
func (c crudCards) HiddenOn(value string) crudCards {
	return c.set("hiddenOn", value)
}

// HideCheckToggler 设置是否隐藏勾选框
func (c crudCards) HideCheckToggler(value bool) crudCards {
	return c.set("hideCheckToggler", value)
}

// HideQuickSaveBtn 设置是否隐藏快速保存按钮
func (c crudCards) HideQuickSaveBtn(value bool) crudCards {
	return c.set("hideQuickSaveBtn", value)
}

// Id 设置组件唯一 ID
func (c crudCards) Id(value string) crudCards {
	return c.set("id", value)
}

// InitFetch 设置初始是否拉取
func (c crudCards) InitFetch(value bool) crudCards {
	return c.set("initFetch", value)
}

// InitFetchOn 设置初始拉取表达式
func (c crudCards) InitFetchOn(value string) crudCards {
	return c.set("initFetchOn", value)
}

// InnerClassName 设置内部 DOM 的 CSS 类名
func (c crudCards) InnerClassName(value string) crudCards {
	return c.set("innerClassName", value)
}

// Interval 设置自动刷新时间
func (c crudCards) Interval(value string) crudCards {
	return c.set("interval", value)
}

// ItemActions 设置单条操作
func (c crudCards) ItemActions(value string) crudCards {
	return c.set("itemActions", value)
}

// ItemCheckableOn 设置约束批量操作
func (c crudCards) ItemCheckableOn(value string) crudCards {
	return c.set("itemCheckableOn", value)
}

// ItemClassName 设置卡片 CSS 类名
func (c crudCards) ItemClassName(value string) crudCards {
	return c.set("itemClassName", value)
}

// ItemDraggableOn 设置项是否可拖拽排序
func (c crudCards) ItemDraggableOn(value string) crudCards {
	return c.set("itemDraggableOn", value)
}

// KeepItemSelectionOnPageChange 设置分页时是否保留用户选择
func (c crudCards) KeepItemSelectionOnPageChange(value bool) crudCards {
	return c.set("keepItemSelectionOnPageChange", value)
}

// LabelTpl 设置已勾选项的文案
func (c crudCards) LabelTpl(value string) crudCards {
	return c.set("labelTpl", value)
}

// LoadDataOnce 设置是否为前端单次加载模式
func (c crudCards) LoadDataOnce(value bool) crudCards {
	return c.set("loadDataOnce", value)
}

// LoadDataOnceFetchOnFilter 设置 loadDataOnce 时的过滤条件
func (c crudCards) LoadDataOnceFetchOnFilter(value bool) crudCards {
	return c.set("loadDataOnceFetchOnFilter", value)
}

// LoadingConfig 设置加载配置
func (c crudCards) LoadingConfig(value string) crudCards {
	return c.set("loadingConfig", value)
}

// MasonryLayout 设置是否为瀑布流布局
func (c crudCards) MasonryLayout(value bool) crudCards {
	return c.set("masonryLayout", value)
}

// MatchFunc 设置自定义搜索匹配函数
func (c crudCards) MatchFunc(value string) crudCards {
	return c.set("matchFunc", value)
}

// Messages 设置消息文案配置
func (c crudCards) Messages(value string) crudCards {
	return c.set("messages", value)
}

// Mode 设置展示模式
func (c crudCards) Mode(value string) crudCards {
	return c.set("mode", value)
}

// Name 设置组件名字
func (c crudCards) Name(value string) crudCards {
	return c.set("name", value)
}

// OnEvent 设置事件动作配置
func (c crudCards) OnEvent(value any) crudCards {
	return c.set("onEvent", value)
}

// OrderBy 设置默认排序字段
func (c crudCards) OrderBy(value string) crudCards {
	return c.set("orderBy", value)
}

// PageField 设置分页字段
func (c crudCards) PageField(value string) crudCards {
	return c.set("pageField", value)
}

// Pagination 设置分页配置
func (c crudCards) Pagination(value string) crudCards {
	return c.set("pagination", value)
}

// PerPage 设置每页条数
func (c crudCards) PerPage(value int) crudCards {
	return c.set("perPage", value)
}

// Query 设置查询字段
func (c crudCards) Query(value string) crudCards {
	return c.set("query", value)
}

// Reload 设置是否需要重新加载
func (c crudCards) Reload(value bool) crudCards {
	return c.set("reload", value)
}

// Render设置是否渲染组件
func (c crudCards) Render(value bool) crudCards {
	return c.set("render", value)
}

// ResetPage 设置是否重置分页
func (c crudCards) ResetPage(value bool) crudCards {
	return c.set("resetPage", value)
}

// RowClassName 设置行 CSS 类名
func (c crudCards) RowClassName(value string) crudCards {
	return c.set("rowClassName", value)
}

// Searchable 设置是否可搜索
func (c crudCards) Searchable(value bool) crudCards {
	return c.set("searchable", value)
}

// SearchPlaceholder 设置搜索占位符
func (c crudCards) SearchPlaceholder(value string) crudCards {
	return c.set("searchPlaceholder", value)
}

// ShowErrorMsg 设置是否显示错误信息
func (c crudCards) ShowErrorMsg(value bool) crudCards {
	return c.set("showErrorMsg", value)
}

// ShowPagination 设置是否显示分页
func (c crudCards) ShowPagination(value bool) crudCards {
	return c.set("showPagination", value)
}

// ShowSearch 设置是否显示搜索框
func (c crudCards) ShowSearch(value bool) crudCards {
	return c.set("showSearch", value)
}

// ShowSort 设置是否显示排序
func (c crudCards) ShowSort(value bool) crudCards {
	return c.set("showSort", value)
}

// ShowSwitch 设置是否显示切换按钮
func (c crudCards) ShowSwitch(value bool) crudCards {
	return c.set("showSwitch", value)
}

// ShowToolbar 设置是否显示工具栏
func (c crudCards) ShowToolbar(value bool) crudCards {
	return c.set("showToolbar", value)
}

// Size 设置组件尺寸
func (c crudCards) Size(value string) crudCards {
	return c.set("size", value)
}

// Source 设置数据源
func (c crudCards) Source(value string) crudCards {
	return c.set("source", value)
}

// Sortable 设置是否可排序
func (c crudCards) Sortable(value bool) crudCards {
	return c.set("sortable", value)
}

// SubmitText 设置提交按钮文字
func (c crudCards) SubmitText(value string) crudCards {
	return c.set("submitText", value)
}

// SyncLocation 设置是否同步 URL
func (c crudCards) SyncLocation(value bool) crudCards {
	return c.set("syncLocation", value)
}

// Title 设置标题
func (c crudCards) Title(value any) crudCards {
	return c.set("title", value)
}

// Translations 设置多语言翻译
func (c crudCards) Translations(value map[string]string) crudCards {
	return c.set("translations", value)
}

// VerticalAlign 设置垂直对齐方式
func (c crudCards) VerticalAlign(value string) crudCards {
	return c.set("verticalAlign", value)
}

// Visible 设置是否可见
func (c crudCards) Visible(value bool) crudCards {
	return c.set("visible", value)
}

// VisibleOn 设置可见性表达式
func (c crudCards) VisibleOn(value string) crudCards {
	return c.set("visibleOn", value)
}

// Width 设置宽度
func (c crudCards) Width(value string) crudCards {
	return c.set("width", value)
}

// WrapItemClassName 设置包裹项的 CSS 类名
func (c crudCards) WrapItemClassName(value string) crudCards {
	return c.set("wrapItemClassName", value)
}

// ExtraAction 设置额外操作
func (c crudCards) ExtraAction(value any) crudCards {
	return c.set("extraAction", value)
}

// AutoSaveOnChange 设置是否自动保存更改
func (c crudCards) AutoSaveOnChange(value bool) crudCards {
	return c.set("autoSaveOnChange", value)
}

// // 这里可以继续添加 CRUD2Cards 中的其他方法...

// Helper 方法
func (c crudCards) merge(other crudCards) crudCards {
	for key, value := range other {
		c[key] = value
	}
	return c
}
