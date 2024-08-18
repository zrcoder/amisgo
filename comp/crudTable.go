package comp

// CRUDTable
type CRUDTable Schema

// NewCRUDTable 创建一个新的 CRUDTable 实例
func NewCRUDTable() CRUDTable {
	return make(CRUDTable).set("type", "crud")
}

func (c CRUDTable) set(key string, value interface{}) CRUDTable {
	c[key] = value
	return c
}

// AffixFooter 设置 affixFooter 属性
func (c CRUDTable) AffixFooter(value bool) CRUDTable {
	return c.set("affixFooter", value)
}

// AffixHeader 设置 affixHeader 属性
func (c CRUDTable) AffixHeader(value bool) CRUDTable {
	return c.set("affixHeader", value)
}

// AffixRow 设置 affixRow 属性
func (c CRUDTable) AffixRow(value string) CRUDTable {
	return c.set("affixRow", value)
}

// AlwaysShowPagination 设置 alwaysShowPagination 属性
func (c CRUDTable) AlwaysShowPagination(value bool) CRUDTable {
	return c.set("alwaysShowPagination", value)
}

// Api 设置 api 属性
func (c CRUDTable) Api(value string) CRUDTable {
	return c.set("api", value)
}

// AutoFillHeight 设置 autoFillHeight 属性
func (c CRUDTable) AutoFillHeight(value string) CRUDTable {
	return c.set("autoFillHeight", value)
}

// AutoGenerateFilter 设置 autoGenerateFilter 属性
func (c CRUDTable) AutoGenerateFilter(value string) CRUDTable {
	return c.set("autoGenerateFilter", value)
}

// AutoJumpToTopOnPagerChange 设置 autoJumpToTopOnPagerChange 属性
func (c CRUDTable) AutoJumpToTopOnPagerChange(value bool) CRUDTable {
	return c.set("autoJumpToTopOnPagerChange", value)
}

// BulkActions 设置 bulkActions 属性
func (c CRUDTable) BulkActions(value string) CRUDTable {
	return c.set("bulkActions", value)
}

// CanAccessSuperData 设置 canAccessSuperData 属性
func (c CRUDTable) CanAccessSuperData(value bool) CRUDTable {
	return c.set("canAccessSuperData", value)
}

// ClassName 设置 className 属性
func (c CRUDTable) ClassName(value string) CRUDTable {
	return c.set("className", value)
}

// Columns 设置 columns 属性
func (c CRUDTable) Columns(value string) CRUDTable {
	return c.set("columns", value)
}

// ColumnsTogglable 设置 columnsTogglable 属性
func (c CRUDTable) ColumnsTogglable(value bool) CRUDTable {
	return c.set("columnsTogglable", value)
}

// CombineFromIndex 设置 combineFromIndex 属性
func (c CRUDTable) CombineFromIndex(value string) CRUDTable {
	return c.set("combineFromIndex", value)
}

// CombineNum 设置 combineNum 属性
func (c CRUDTable) CombineNum(value string) CRUDTable {
	return c.set("combineNum", value)
}

// Data 设置 data 属性
func (c CRUDTable) Data(value string) CRUDTable {
	return c.set("data", value)
}

// DefaultParams 设置 defaultParams 属性
func (c CRUDTable) DefaultParams(value string) CRUDTable {
	return c.set("defaultParams", value)
}

// DeferApi 设置 deferApi 属性
func (c CRUDTable) DeferApi(value string) CRUDTable {
	return c.set("deferApi", value)
}

// Disabled 设置 disabled 属性
func (c CRUDTable) Disabled(value bool) CRUDTable {
	return c.set("disabled", value)
}

// DisabledOn 设置 disabledOn 属性
func (c CRUDTable) DisabledOn(value string) CRUDTable {
	return c.set("disabledOn", value)
}

// Draggable 设置 draggable 属性
func (c CRUDTable) Draggable(value bool) CRUDTable {
	return c.set("draggable", value)
}

// DraggableOn 设置 draggableOn 属性
func (c CRUDTable) DraggableOn(value string) CRUDTable {
	return c.set("draggableOn", value)
}

// EditorSetting 设置 editorSetting 属性
func (c CRUDTable) EditorSetting(value string) CRUDTable {
	return c.set("editorSetting", value)
}

// ExpandConfig 设置 expandConfig 属性
func (c CRUDTable) ExpandConfig(value string) CRUDTable {
	return c.set("expandConfig", value)
}

// Filter 设置 filter 属性
func (c CRUDTable) Filter(value string) CRUDTable {
	return c.set("filter", value)
}

// FilterDefaultVisible 设置 filterDefaultVisible 属性
func (c CRUDTable) FilterDefaultVisible(value bool) CRUDTable {
	return c.set("filterDefaultVisible", value)
}

// FilterTogglable 设置 filterTogglable 属性
func (c CRUDTable) FilterTogglable(value bool) CRUDTable {
	return c.set("filterTogglable", value)
}

// Footable 设置 footable 属性
func (c CRUDTable) Footable(value bool) CRUDTable {
	return c.set("footable", value)
}

// FormItemMode 设置 formItemMode 属性
func (c CRUDTable) FormItemMode(value string) CRUDTable {
	return c.set("formItemMode", value)
}

// HeaderTitle 设置 headerTitle 属性
func (c CRUDTable) HeaderTitle(value string) CRUDTable {
	return c.set("headerTitle", value)
}

// HoverAction 设置 hoverAction 属性
func (c CRUDTable) HoverAction(value bool) CRUDTable {
	return c.set("hoverAction", value)
}

// HoverActionType 设置 hoverActionType 属性
func (c CRUDTable) HoverActionType(value string) CRUDTable {
	return c.set("hoverActionType", value)
}

// ItemActions 设置 itemActions 属性
func (c CRUDTable) ItemActions(value string) CRUDTable {
	return c.set("itemActions", value)
}

// LabelTpl 设置 labelTpl 属性
func (c CRUDTable) LabelTpl(value string) CRUDTable {
	return c.set("labelTpl", value)
}

// Loading设置 loading 属性
func (c CRUDTable) Loading(value bool) CRUDTable {
	return c.set("loading", value)
}

// MaxHeight 设置 maxHeight 属性
func (c CRUDTable) MaxHeight(value string) CRUDTable {
	return c.set("maxHeight", value)
}

// MaxIndex 设置 maxIndex 属性
func (c CRUDTable) MaxIndex(value int) CRUDTable {
	return c.set("maxIndex", value)
}

// Name 设置 name 属性
func (c CRUDTable) Name(value string) CRUDTable {
	return c.set("name", value)
}

// Pagination 设置 pagination 属性
func (c CRUDTable) Pagination(value bool) CRUDTable {
	return c.set("pagination", value)
}

// PageField 设置 pageField 属性
func (c CRUDTable) PageField(value string) CRUDTable {
	return c.set("pageField", value)
}

// PerPage 设置 perPage 属性
func (c CRUDTable) PerPage(value int) CRUDTable {
	return c.set("perPage", value)
}

// Placeholder 设置 placeholder 属性
func (c CRUDTable) Placeholder(value string) CRUDTable {
	return c.set("placeholder", value)
}

// RowClassName 设置 rowClassName 属性
func (c CRUDTable) RowClassName(value string) CRUDTable {
	return c.set("rowClassName", value)
}

// SaveOrder 设置 saveOrder 属性
func (c CRUDTable) SaveOrder(value bool) CRUDTable {
	return c.set("saveOrder", value)
}

// SearchSettings 设置 searchSettings 属性
func (c CRUDTable) SearchSettings(value string) CRUDTable {
	return c.set("searchSettings", value)
}

// Searchable 设置 searchable 属性
func (c CRUDTable) Searchable(value bool) CRUDTable {
	return c.set("searchable", value)
}

// ShowSetting 设置 showSetting 属性
func (c CRUDTable) ShowSetting(value bool) CRUDTable {
	return c.set("showSetting", value)
}

// Size 设置 size 属性
func (c CRUDTable) Size(value string) CRUDTable {
	return c.set("size", value)
}

// Sticky 设置 sticky 属性
func (c CRUDTable) Sticky(value bool) CRUDTable {
	return c.set("sticky", value)
}

// StickyHeader 设置 stickyHeader 属性
func (c CRUDTable) StickyHeader(value bool) CRUDTable {
	return c.set("stickyHeader", value)
}

// Striped 设置 striped 属性
func (c CRUDTable) Striped(value bool) CRUDTable {
	return c.set("striped", value)
}

// ToggledBySetting 设置 toggledBySetting 属性
func (c CRUDTable) ToggledBySetting(value bool) CRUDTable {
	return c.set("toggledBySetting", value)
}

// TableWidth 设置 tableWidth 属性
func (c CRUDTable) TableWidth(value string) CRUDTable {
	return c.set("tableWidth", value)
}

// Title 设置 title 属性
func (c CRUDTable) Title(value string) CRUDTable {
	return c.set("title", value)
}

// UnmountOnHidden 设置 unmountOnHidden 属性
func (c CRUDTable) UnmountOnHidden(value bool) CRUDTable {
	return c.set("unmountOnHidden", value)
}

// UniqueId 设置 uniqueId 属性
func (c CRUDTable) UniqueId(value string) CRUDTable {
	return c.set("uniqueId", value)
}

// VerticalAlign 设置 verticalAlign 属性
func (c CRUDTable) VerticalAlign(value string) CRUDTable {
	return c.set("verticalAlign", value)
}

// Width 设置 width 属性
func (c CRUDTable) Width(value string) CRUDTable {
	return c.set("width", value)
}

// Wrap 设置 wrap 属性
func (c CRUDTable) Wrap(value bool) CRUDTable {
	return c.set("wrap", value)
}
