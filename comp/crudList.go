package comp

// crudList 表示 CRUD 列表渲染器
type crudList schema

// CrudList 创建新的 CrudList 实例
func CrudList() crudList {
	return make(crudList).set("mode", "list").set("type", "crud")
}

// Set 设置键值对，并返回当前实例
func (c crudList) set(key string, value interface{}) crudList {
	c[key] = value
	return c
}

// AffixFooter 设置是否固底
func (c crudList) AffixFooter(value bool) crudList {
	return c.set("affixFooter", value)
}

// AffixHeader 设置是否固顶
func (c crudList) AffixHeader(value bool) crudList {
	return c.set("affixHeader", value)
}

// AlwaysShowPagination 默认只显示分页数大于 1 时才显示，设置为 true 总是显示
func (c crudList) AlwaysShowPagination(value bool) crudList {
	return c.set("alwaysShowPagination", value)
}

// API 设置初始化数据 API (整合 CRUD2List 和 CRUDList 的相同方法)
func (c crudList) API(value string) crudList {
	return c.set("api", value)
}

// AutoFillHeight 设置内容区域占满屏幕剩余空间
func (c crudList) AutoFillHeight(value bool) crudList {
	return c.set("autoFillHeight", value)
}

// AutoGenerateFilter 开启查询区域，会根据列元素的searchable属性值，自动生成查询条件表单
func (c crudList) AutoGenerateFilter(value string) crudList {
	return c.set("autoGenerateFilter", value)
}

// AutoJumpToTopOnPagerChange 设置是否自动跳顶部，当切分页的时候
func (c crudList) AutoJumpToTopOnPagerChange(value bool) crudList {
	return c.set("autoJumpToTopOnPagerChange", value)
}

// BulkActions 设置批量操作
func (c crudList) BulkActions(value string) crudList {
	return c.set("bulkActions", value)
}

// CheckOnItemClick 设置点击列表单行时是否选择
func (c crudList) CheckOnItemClick(value bool) crudList {
	return c.set("checkOnItemClick", value)
}

// ClassName 设置容器 CSS 类名
func (c crudList) ClassName(value string) crudList {
	return c.set("className", value)
}

// DefaultParams 设置默认初始参数
func (c crudList) DefaultParams(value string) crudList {
	return c.set("defaultParams", value)
}

// DeferApi 设置懒加载 API
func (c crudList) DeferApi(value string) crudList {
	return c.set("deferApi", value)
}

// Disabled 设置是否禁用
func (c crudList) Disabled(value bool) crudList {
	return c.set("disabled", value)
}

// DisabledOn 设置是否禁用表达式
func (c crudList) DisabledOn(value string) crudList {
	return c.set("disabledOn", value)
}

// Draggable 设置是否可通过拖拽排序
func (c crudList) Draggable(value bool) crudList {
	return c.set("draggable", value)
}

// DraggableOn 设置可拖拽排序表达式
func (c crudList) DraggableOn(value string) crudList {
	return c.set("draggableOn", value)
}

// EditorSetting 设置编辑器配置
func (c crudList) EditorSetting(value string) crudList {
	return c.set("editorSetting", value)
}

// ExpandConfig 设置内嵌模式的默认展开选项
func (c crudList) ExpandConfig(value string) crudList {
	return c.set("expandConfig", value)
}

// Filter 设置过滤器表单
func (c crudList) Filter(value string) crudList {
	return c.set("filter", value)
}

// FilterDefaultVisible 设置过滤器默认是否可见
func (c crudList) FilterDefaultVisible(value bool) crudList {
	return c.set("filterDefaultVisible", value)
}

// FilterTogglable 设置过滤器是否可切换
func (c crudList) FilterTogglable(value bool) crudList {
	return c.set("filterTogglable", value)
}

// Footer 设置底部区域
func (c crudList) Footer(value interface{}) crudList {
	return c.set("footer", value)
}

// FooterClassName 设置底部区域 CSS 类名
func (c crudList) FooterClassName(value string) crudList {
	return c.set("footerClassName", value)
}

// FooterToolbar 设置底部工具栏区域
func (c crudList) FooterToolbar(value interface{}) crudList {
	return c.set("footerToolbar", value)
}

// Header 设置顶部区域
func (c crudList) Header(value interface{}) crudList {
	return c.set("header", value)
}

// HeaderClassName 设置顶部区域 CSS 类名
func (c crudList) HeaderClassName(value string) crudList {
	return c.set("headerClassName", value)
}

// HeaderToolbar 设置顶部工具栏区域
func (c crudList) HeaderToolbar(value interface{}) crudList {
	return c.set("headerToolbar", value)
}

// Hidden 设置是否隐藏
func (c crudList) Hidden(value bool) crudList {
	return c.set("hidden", value)
}

// HiddenOn 设置是否隐藏表达式
func (c crudList) HiddenOn(value string) crudList {
	return c.set("hiddenOn", value)
}

// HideCheckToggler 设置是否隐藏勾选框
func (c crudList) HideCheckToggler(value bool) crudList {
	return c.set("hideCheckToggler", value)
}

// HideQuickSaveBtn 设置是否隐藏快速编辑按钮
func (c crudList) HideQuickSaveBtn(value bool) crudList {
	return c.set("hideQuickSaveBtn", value)
}

// ID 设置组件唯一 ID
func (c crudList) ID(value string) crudList {
	return c.set("id", value)
}

// InitFetch 设置初始是否拉取
func (c crudList) InitFetch(value bool) crudList {
	return c.set("initFetch", value)
}

// InitFetchOn 设置初始是否拉取，用表达式来配置
func (c crudList) InitFetchOn(value string) crudList {
	return c.set("initFetchOn", value)
}

// InnerClassName 设置内部 DOM 的 className
func (c crudList) InnerClassName(value string) crudList {
	return c.set("innerClassName", value)
}

// Interval 设置自动刷新时间
func (c crudList) Interval(value string) crudList {
	return c.set("interval", value)
}

// ItemAction 设置点击列表项的行为
func (c crudList) ItemAction(value string) crudList {
	return c.set("itemAction", value)
}

// ItemActions 设置单条操作
func (c crudList) ItemActions(value string) crudList {
	return c.set("itemActions", value)
}

// ItemCheckableOn 设置约束批量操作
func (c crudList) ItemCheckableOn(value string) crudList {
	return c.set("itemCheckableOn", value)
}

// ItemDraggableOn 配置某项是否可拖拽排序
func (c crudList) ItemDraggableOn(value string) crudList {
	return c.set("itemDraggableOn", value)
}

// KeepItemSelectionOnPageChange 翻页时是否保留用户已选的数据
func (c crudList) KeepItemSelectionOnPageChange(value bool) crudList {
	return c.set("keepItemSelectionOnPageChange", value)
}

// LabelTpl 设置已勾选项的文案
func (c crudList) LabelTpl(value string) crudList {
	return c.set("labelTpl", value)
}

// ListItem 单条数据展示内容配置
func (c crudList) ListItem(value interface{}) crudList {
	return c.set("listItem", value)
}

// LoadDataOnce 是否为前端单次加载模式
func (c crudList) LoadDataOnce(value bool) crudList {
	return c.set("loadDataOnce", value)
}

// LoadDataOnceFetchOnFilter 设置在开启 loadDataOnce 时，当修改过滤条件时是否重新请求 API
func (c crudList) LoadDataOnceFetchOnFilter(value bool) crudList {
	return c.set("loadDataOnceFetchOnFilter", value)
}

// LoadingConfig 设置加载配置
func (c crudList) LoadingConfig(value string) crudList {
	return c.set("loadingConfig", value)
}

// MatchFunc 设置自定义搜索匹配函数
func (c crudList) MatchFunc(value string) crudList {
	return c.set("matchFunc", value)
}

// Messages 设置消息文案配置
func (c crudList) Messages(value string) crudList {
	return c.set("messages", value)
}

// Mode 设置指定内容区的展示模式
func (c crudList) Mode(value string) crudList {
	return c.set("mode", value)
}

// Name 设置组件名字
func (c crudList) Name(value string) crudList {
	return c.set("name", value)
}

// OnEvent 事件绑定
func (c crudList) OnEvent(value string) crudList {
	return c.set("onEvent", value)
}

// Pagination 设置分页配置
func (c crudList) Pagination(value interface{}) crudList {
	return c.set("pagination", value)
}

// Placeholder 设置列表空数据的提示
func (c crudList) Placeholder(value string) crudList {
	return c.set("placeholder", value)
}

// QuickSaveApi 设置快速保存 API
func (c crudList) QuickSaveApi(value string) crudList {
	return c.set("quickSaveApi", value)
}

// QuickSaveItemActions 设置快速编辑按钮配置
func (c crudList) QuickSaveItemActions(value string) crudList {
	return c.set("quickSaveItemActions", value)
}

// QuickSaveSuccessMessage 设置快速保存成功的提示信息
func (c crudList) QuickSaveSuccessMessage(value string) crudList {
	return c.set("quickSaveSuccessMessage", value)
}

// RefreshInterval 设置自动刷新时间
func (c crudList) RefreshInterval(value string) crudList {
	return c.set("refreshInterval", value)
}

// RemoteFilter 支持远程筛选
func (c crudList) RemoteFilter(value bool) crudList {
	return c.set("remoteFilter", value)
}

// RenderType 设置渲染类型
func (c crudList) RenderType(value string) crudList {
	return c.set("renderType", value)
}

// ResetValueOnHidden 设置隐藏时是否重置值
func (c crudList) ResetValueOnHidden(value bool) crudList {
	return c.set("resetValueOnHidden", value)
}

// RowClassName 设置行 CSS 类名
func (c crudList) RowClassName(value string) crudList {
	return c.set("rowClassName", value)
}

// SearchFormClassName 设置搜索表单 CSS 类名
func (c crudList) SearchFormClassName(value string) crudList {
	return c.set("searchFormClassName", value)
}

// SearchSchema 设置搜索表单
func (c crudList) SearchSchema(value interface{}) crudList {
	return c.set("searchSchema", value)
}

// SortSettings 设置排序配置
func (c crudList) SortSettings(value string) crudList {
	return c.set("sortSettings", value)
}

// SortFieldName 设置排序字段名
func (c crudList) SortFieldName(value string) crudList {
	return c.set("sortFieldName", value)
}

// Source 设置数据来源
func (c crudList) Source(value string) crudList {
	return c.set("source", value)
}

// Stateful 设置是否保留状态
func (c crudList) Stateful(value bool) crudList {
	return c.set("stateful", value)
}

// TableClassName 设置表格 CSS 类名
func (c crudList) TableClassName(value string) crudList {
	return c.set("tableClassName", value)
}

// TableLayout 设置表格布局
func (c crudList) TableLayout(value string) crudList {
	return c.set("tableLayout", value)
}

// Tooltip设置 提示信息配置
func (c crudList) Tooltip(value string) crudList {
	return c.set("tooltip", value)
}

// TransformItems 设置数据转换
func (c crudList) TransformItems(value string) crudList {
	return c.set("transformItems", value)
}

// Virtualized 设置虚拟化
func (c crudList) Virtualized(value bool) crudList {
	return c.set("virtualized", value)
}

// Width 设置宽度
func (c crudList) Width(value string) crudList {
	return c.set("width", value)
}

// EmptyText 设置无数据时的提示
func (c crudList) EmptyText(value string) crudList {
	return c.set("emptyText", value)
}

// Components 设置自定义组件
func (c crudList) Components(value interface{}) crudList {
	return c.set("components", value)
}

// 其他 CRUD2List 特有字段和方法
func (c crudList) CustomField1(value string) crudList {
	return c.set("customField1", value)
}

// 其他 CRUDList 特有字段和方法
func (c crudList) CustomField2(value string) crudList {
	return c.set("customField2", value)
}

// ...
