package comp

// CRUD2Cards 是 CRUD2 卡片渲染器
type CRUD2Cards Schema

// NewCRUD2Cards 创建一个新的 CRUD2Cards 实例
func NewCRUD2Cards() CRUD2Cards {
	return make(CRUD2Cards).set("mode", "cards").set("type", "crud2")
}

// Set 设置属性
func (c CRUD2Cards) set(key string, value interface{}) CRUD2Cards {
	c[key] = value
	return c
}

// AffixFooter 设置是否固底
func (c CRUD2Cards) AffixFooter(value bool) CRUD2Cards {
	return c.set("affixFooter", value)
}

// AffixHeader 设置是否固顶
func (c CRUD2Cards) AffixHeader(value bool) CRUD2Cards {
	return c.set("affixHeader", value)
}

// API 设置初始化数据 API
func (c CRUD2Cards) API(value string) CRUD2Cards {
	return c.set("api", value)
}

// AutoFillHeight 设置内容区域占满屏幕剩余空间
func (c CRUD2Cards) AutoFillHeight(value bool) CRUD2Cards {
	return c.set("autoFillHeight", value)
}

// AutoJumpToTopOnPagerChange 设置是否自动跳顶部，当切分页的时候
func (c CRUD2Cards) AutoJumpToTopOnPagerChange(value bool) CRUD2Cards {
	return c.set("autoJumpToTopOnPagerChange", value)
}

// Card 设置卡片配置
func (c CRUD2Cards) Card(value interface{}) CRUD2Cards {
	return c.set("card", value)
}

// CheckOnItemClick 设置点击卡片时是否勾选卡片
func (c CRUD2Cards) CheckOnItemClick(value bool) CRUD2Cards {
	return c.set("checkOnItemClick", value)
}

// ClassName 设置容器 CSS 类名
func (c CRUD2Cards) ClassName(value string) CRUD2Cards {
	return c.set("className", value)
}

// Disabled 设置是否禁用
func (c CRUD2Cards) Disabled(value bool) CRUD2Cards {
	return c.set("disabled", value)
}

// DisabledOn 设置是否禁用表达式
func (c CRUD2Cards) DisabledOn(value string) CRUD2Cards {
	return c.set("disabledOn", value)
}

// EditorSetting 设置编辑器配置
func (c CRUD2Cards) EditorSetting(value string) CRUD2Cards {
	return c.set("editorSetting", value)
}

// Footer 设置底部区域
func (c CRUD2Cards) Footer(value interface{}) CRUD2Cards {
	return c.set("footer", value)
}

// FooterClassName 设置底部 CSS 类名
func (c CRUD2Cards) FooterClassName(value string) CRUD2Cards {
	return c.set("footerClassName", value)
}

// FooterToolbar 设置底部工具栏区域
func (c CRUD2Cards) FooterToolbar(value interface{}) CRUD2Cards {
	return c.set("footerToolbar", value)
}

// FooterToolbarClassName 设置底部工具栏 CSS 类名
func (c CRUD2Cards) FooterToolbarClassName(value string) CRUD2Cards {
	return c.set("footerToolbarClassName", value)
}

// Header 设置顶部区域
func (c CRUD2Cards) Header(value interface{}) CRUD2Cards {
	return c.set("header", value)
}

// HeaderClassName 设置顶部 CSS 类名
func (c CRUD2Cards) HeaderClassName(value string) CRUD2Cards {
	return c.set("headerClassName", value)
}

// HeaderToolbar 设置顶部工具栏区域
func (c CRUD2Cards) HeaderToolbar(value interface{}) CRUD2Cards {
	return c.set("headerToolbar", value)
}

// HeaderToolbarClassName 设置顶部工具栏 CSS 类名
func (c CRUD2Cards) HeaderToolbarClassName(value string) CRUD2Cards {
	return c.set("headerToolbarClassName", value)
}

// Hidden 设置是否隐藏
func (c CRUD2Cards) Hidden(value bool) CRUD2Cards {
	return c.set("hidden", value)
}

// HiddenOn 设置是否隐藏表达式
func (c CRUD2Cards) HiddenOn(value string) CRUD2Cards {
	return c.set("hiddenOn", value)
}

// HideCheckToggler 设置是否隐藏勾选框
func (c CRUD2Cards) HideCheckToggler(value bool) CRUD2Cards {
	return c.set("hideCheckToggler", value)
}

// HideQuickSaveBtn 设置是否隐藏快速编辑按钮
func (c CRUD2Cards) HideQuickSaveBtn(value bool) CRUD2Cards {
	return c.set("hideQuickSaveBtn", value)
}

// ID 设置组件唯一 ID
func (c CRUD2Cards) ID(value string) CRUD2Cards {
	return c.set("id", value)
}

// Interval 设置自动刷新时间
func (c CRUD2Cards) Interval(value string) CRUD2Cards {
	return c.set("interval", value)
}

// ItemCheckableOn 设置约束批量操作
func (c CRUD2Cards) ItemCheckableOn(value string) CRUD2Cards {
	return c.set("itemCheckableOn", value)
}

// ItemClassName 设置卡片 CSS 类名
func (c CRUD2Cards) ItemClassName(value string) CRUD2Cards {
	return c.set("itemClassName", value)
}

// ItemDraggableOn 设置配置项是否可拖拽排序
func (c CRUD2Cards) ItemDraggableOn(value string) CRUD2Cards {
	return c.set("itemDraggableOn", value)
}

// KeepItemSelectionOnPageChange 设置翻页时是否保留用户已选数据
func (c CRUD2Cards) KeepItemSelectionOnPageChange(value bool) CRUD2Cards {
	return c.set("keepItemSelectionOnPageChange", value)
}

// LoadDataOnce 设置是否为前端单次加载模式
func (c CRUD2Cards) LoadDataOnce(value bool) CRUD2Cards {
	return c.set("loadDataOnce", value)
}

// LoadType 设置数据展示模式
func (c CRUD2Cards) LoadType(value string) CRUD2Cards {
	return c.set("loadType", value)
}

// LoadingConfig 设置加载配置
func (c CRUD2Cards) LoadingConfig(value string) CRUD2Cards {
	return c.set("loadingConfig", value)
}

// MasonryLayout 设置是否为瀑布流布局
func (c CRUD2Cards) MasonryLayout(value bool) CRUD2Cards {
	return c.set("masonryLayout", value)
}

// Mode 设置指定内容区的展示模式
func (c CRUD2Cards) Mode(value string) CRUD2Cards {
	return c.set("mode", value)
}

// Multiple 设置是否可以多选数据
func (c CRUD2Cards) Multiple(value bool) CRUD2Cards {
	return c.set("multiple", value)
}

// Name 设置组件名字
func (c CRUD2Cards) Name(value string) CRUD2Cards {
	return c.set("name", value)
}

// OnEvent 设置事件动作配置
func (c CRUD2Cards) OnEvent(value string) CRUD2Cards {
	return c.set("onEvent", value)
}

// PageField 设置分页页码字段名
func (c CRUD2Cards) PageField(value string) CRUD2Cards {
	return c.set("pageField", value)
}

// ParsePrimitiveQuery 设置是否开启 Query 信息转换
func (c CRUD2Cards) ParsePrimitiveQuery(value string) CRUD2Cards {
	return c.set("parsePrimitiveQuery", value)
}

// PerPage 设置无限加载时每页加载数量
func (c CRUD2Cards) PerPage(value string) CRUD2Cards {
	return c.set("perPage", value)
}

// PerPageField 设置分页每页显示的多少条数据的字段名
func (c CRUD2Cards) PerPageField(value string) CRUD2Cards {
	return c.set("perPageField", value)
}

// Placeholder 设置无数据提示
func (c CRUD2Cards) Placeholder(value string) CRUD2Cards {
	return c.set("placeholder", value)
}

// PrimaryField 设置行标识符
func (c CRUD2Cards) PrimaryField(value string) CRUD2Cards {
	return c.set("primaryField", value)
}

// QuickSaveApi 设置快速编辑后批量保存的 API
func (c CRUD2Cards) QuickSaveApi(value string) CRUD2Cards {
	return c.set("quickSaveApi", value)
}

// QuickSaveItemApi 设置快速编辑配置成及时保存时使用的 API
func (c CRUD2Cards) QuickSaveItemApi(value string) CRUD2Cards {
	return c.set("quickSaveItemApi", value)
}

// SaveOrderApi 设置保存排序的 API
func (c CRUD2Cards) SaveOrderApi(value string) CRUD2Cards {
	return c.set("saveOrderApi", value)
}

// Selectable 设置是否可以选择
func (c CRUD2Cards) Selectable(value bool) CRUD2Cards {
	return c.set("selectable", value)
}

// ShowQuickSaveBtn 设置是否显示快速编辑按钮
func (c CRUD2Cards) ShowQuickSaveBtn(value bool) CRUD2Cards {
	return c.set("showQuickSaveBtn", value)
}

// Source 设置数据来源
func (c CRUD2Cards) Source(value interface{}) CRUD2Cards {
	return c.set("source", value)
}

// Title 设置标题
func (c CRUD2Cards) Title(value string) CRUD2Cards {
	return c.set("title", value)
}
