package comp

import "github.com/zrcoder/amisgo/model"

// Table is a Table renderer. Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/Table
type Table model.Schema

func NewTable() Table {
	return Table{"type": "table"}
}

func NewTable2() Table {
	return Table{"type": "table2"}
}

func (t Table) set(key string, value any) Table {
	t[key] = value
	return t
}

// AffixFooter sets whether the footer is fixed
func (t Table) AffixFooter(value bool) Table {
	return t.set("affixFooter", value)
}

// AffixHeader sets whether the header is fixed
func (t Table) AffixHeader(value bool) Table {
	return t.set("affixHeader", value)
}

// AffixRow sets the bottom summary row
func (t Table) AffixRow(value string) Table {
	return t.set("affixRow", value)
}

// AutoFillHeight sets the table to auto-calculate height
func (t Table) AutoFillHeight(value string) Table {
	return t.set("autoFillHeight", value)
}

// AutoGenerateFilter enables query area, auto-generates query form based on column's searchable attribute
func (t Table) AutoGenerateFilter(value string) Table {
	return t.set("autoGenerateFilter", value)
}

// CanAccessSuperData sets whether the table can access parent data domain values, default is false
func (t Table) CanAccessSuperData(value bool) Table {
	return t.set("canAccessSuperData", value)
}

// ClassName sets the container CSS class name
func (t Table) ClassName(value string) Table {
	return t.set("className", value)
}

// Columns sets the table's column information
func (t Table) Columns(value ...any) Table {
	return t.set("columns", value)
}

// ColumnsTogglable sets whether to show column toggle switch, auto if columns >= 5
func (t Table) ColumnsTogglable(value bool) Table {
	return t.set("columnsTogglable", value)
}

// CombineFromIndex sets the merge cell configuration, starting from which column
func (t Table) CombineFromIndex(value string) Table {
	return t.set("combineFromIndex", value)
}

// CombineNum sets the merge cell configuration, number indicates how many columns to merge from left to right
func (t Table) CombineNum(value string) Table {
	return t.set("combineNum", value)
}

// Data sets the data
func (t Table) Data(value any) Table {
	return t.set("data", value)
}

// DeferApi sets the lazy load API, used when row data is marked with defer: true
func (t Table) DeferApi(value string) Table {
	return t.set("deferApi", value)
}

// Disabled sets whether the table is disabled
func (t Table) Disabled(value bool) Table {
	return t.set("disabled", value)
}

// DisabledOn sets the expression for disabling the table
func (t Table) DisabledOn(value string) Table {
	return t.set("disabledOn", value)
}

// EditorSetting sets the editor configuration, can be ignored at runtime
func (t Table) EditorSetting(value string) Table {
	return t.set("editorSetting", value)
}

// Footable sets whether to enable the footer display feature, suitable for mobile display
func (t Table) Footable(value bool) Table {
	return t.set("footable", value)
}

// FooterClassName sets the CSS class name for the footer container
func (t Table) FooterClassName(value string) Table {
	return t.set("footerClassName", value)
}

// HeaderClassName sets the CSS class name for the header container
func (t Table) HeaderClassName(value string) Table {
	return t.set("headerClassName", value)
}

// Hidden sets whether the table is hidden
func (t Table) Hidden(value bool) Table {
	return t.set("hidden", value)
}

// HiddenOn sets the expression for hiding the table
func (t Table) HiddenOn(value string) Table {
	return t.set("hiddenOn", value)
}

// ID sets the unique component ID, mainly for log collection
func (t Table) ID(value string) Table {
	return t.set("id", value)
}

// ItemBadge sets the row badge (Badge)
func (t Table) ItemBadge(value string) Table {
	return t.set("itemBadge", value)
}

// OnEvent sets the event action configuration
func (t Table) OnEvent(value any) Table {
	return t.set("onEvent", value)
}

// Placeholder sets the placeholder
func (t Table) Placeholder(value string) Table {
	return t.set("placeholder", value)
}

// PrefixRow sets the top summary row
func (t Table) PrefixRow(value string) Table {
	return t.set("prefixRow", value)
}

// Resizable sets whether the column width is adjustable
func (t Table) Resizable(value bool) Table {
	return t.set("resizable", value)
}

// RowClassNameExpr sets the row style expression
func (t Table) RowClassNameExpr(value string) Table {
	return t.set("rowClassNameExpr", value)
}

// ShowFooter sets whether to show the footer
func (t Table) ShowFooter(value bool) Table {
	return t.set("showFooter", value)
}

// ShowHeader sets whether to show the header
func (t Table) ShowHeader(value bool) Table {
	return t.set("showHeader", value)
}

// Source sets the data source: bind to the current environment variable
func (t Table) Source(value string) Table {
	return t.set("source", value)
}

// Static sets whether to display statically
func (t Table) Static(value bool) Table {
	return t.set("static", value)
}

// StaticClassName sets the static display form item class name
func (t Table) StaticClassName(value string) Table {
	return t.set("staticClassName", value)
}

// StaticInputClassName sets the static display form item value class name
func (t Table) StaticInputClassName(value string) Table {
	return t.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static display form item label class name
func (t Table) StaticLabelClassName(value string) Table {
	return t.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display
func (t Table) StaticOn(value string) Table {
	return t.set("staticOn", value)
}

// StaticPlaceholder sets the static display empty value placeholder
func (t Table) StaticPlaceholder(value string) Table {
	return t.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema
func (t Table) StaticSchema(value string) Table {
	return t.set("staticSchema", value)
}

// Style sets the component style
func (t Table) Style(value any) Table {
	return t.set("style", value)
}

// TableClassName sets the table CSS class name
func (t Table) TableClassName(value string) Table {
	return t.set("tableClassName", value)
}

// TableLayout sets the table layout, options: fixed | auto
func (t Table) TableLayout(value string) Table {
	return t.set("tableLayout", value)
}

// TestIdBuilder sets the test ID builder
func (t Table) TestIdBuilder(value string) Table {
	return t.set("testIdBuilder", value)
}

// Testid sets the test ID
func (t Table) Testid(value string) Table {
	return t.set("testid", value)
}

// Title sets the title
func (t Table) Title(value any) Table {
	return t.set("title", value)
}

// ToolbarClassName sets the toolbar CSS class name
func (t Table) ToolbarClassName(value string) Table {
	return t.set("toolbarClassName", value)
}

// UseMobileUI sets whether to disable mobile styles at the component level
func (t Table) UseMobileUI(value bool) Table {
	return t.set("useMobileUI", value)
}

// Visible sets whether the table is visible
func (t Table) Visible(value bool) Table {
	return t.set("visible", value)
}

// VisibleOn sets the expression for visibility
func (t Table) VisibleOn(value string) Table {
	return t.set("visibleOn", value)
}

// Actions sets the action column configuration
func (t Table) Actions(value string) Table {
	return t.set("actions", value)
}

// Bordered sets whether to show borders
func (t Table) Bordered(value bool) Table {
	return t.set("bordered", value)
}

// ChildrenColumnName sets the custom field name for nested data source
func (t Table) ChildrenColumnName(value string) Table {
	return t.set("childrenColumnName", value)
}

// Expandable sets whether the table row is expandable
func (t Table) Expandable(value bool) Table {
	return t.set("expandable", value)
}

// Footer sets the footer
func (t Table) Footer(value string) Table {
	return t.set("footer", value)
}

// KeepItemSelectionOnPageChange sets whether to keep data on page change
func (t Table) KeepItemSelectionOnPageChange(value bool) Table {
	return t.set("keepItemSelectionOnPageChange", value)
}

// KeyField sets the ID field name for multi-select and nested expansion, default is id
func (t Table) KeyField(value string) Table {
	return t.set("keyField", value)
}

// LazyRenderAfter sets the number of columns to render at once, default is 100, can improve table rendering performance
func (t Table) LazyRenderAfter(value string) Table {
	return t.set("lazyRenderAfter", value)
}

// LineHeight sets whether to fix the content row height
func (t Table) LineHeight(value string) Table {
	return t.set("lineHeight", value)
}

// Loading sets the loading state
func (t Table) Loading(value string) Table {
	return t.set("loading", value)
}

// MaxKeepItemSelectionLength sets the maximum limit for batch operations
func (t Table) MaxKeepItemSelectionLength(value string) Table {
	return t.set("maxKeepItemSelectionLength", value)
}

// Messages sets the interface error message configuration (message text configuration, remember this has the lowest priority, if your interface returns msg, the interface return takes precedence)
func (t Table) Messages(value string) Table {
	return t.set("messages", value)
}

// Multiple sets whether multiple selection is allowed, same as rowSelection.type, compatible with original CRUD attribute, default is multi-select, only works when selectable is set
func (t Table) Multiple(value bool) Table {
	return t.set("multiple", value)
}

// PopOverContainer sets the specified mount DOM
func (t Table) PopOverContainer(value string) Table {
	return t.set("popOverContainer", value)
}

// PrimaryField sets the ID field name, same as keyField, compatible with original CRUD attribute
func (t Table) PrimaryField(value string) Table {
	return t.set("primaryField", value)
}

// QuickSaveApi sets the API for batch saving after quick editing
func (t Table) QuickSaveApi(value string) Table {
	return t.set("quickSaveApi", value)
}

// QuickSaveItemApi sets the API for immediate saving when quick editing is configured
func (t Table) QuickSaveItemApi(value string) Table {
	return t.set("quickSaveItemApi", value)
}

// Reload sets the name of the component to reload
func (t Table) Reload(value string) Table {
	return t.set("reload", value)
}

// RowSelection sets the table selection configuration
func (t Table) RowSelection(value string) Table {
	return t.set("rowSelection", value)
}

// Selectable sets whether the table is selectable, same as rowSelection, compatible with original CRUD attribute, default is multi-select
func (t Table) Selectable(value bool) Table {
	return t.set("selectable", value)
}

// ShowBadge sets whether to show row badges
func (t Table) ShowBadge(value bool) Table {
	return t.set("showBadge", value)
}

// Sticky sets whether the header is sticky
func (t Table) Sticky(value bool) Table {
	return t.set("sticky", value)
}

// Size sets the table size, supports large, default, small, default is medium size
func (t Table) Size(value string) Table {
	return t.set("size", value)
}
