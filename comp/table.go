package comp

import "github.com/zrcoder/amisgo/model"

// table is a table renderer. Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/table

type table model.Schema

// Table creates a new Table instance
func Table() table {
	return table{}.set("type", "table")
}

func Table2() table {
	return table{}.set("type", "table2")
}

func (t table) set(key string, value any) table {
	t[key] = value
	return t
}

// AffixFooter sets whether the footer is fixed
func (t table) AffixFooter(value bool) table {
	return t.set("affixFooter", value)
}

// AffixHeader sets whether the header is fixed
func (t table) AffixHeader(value bool) table {
	return t.set("affixHeader", value)
}

// AffixRow sets the bottom summary row
func (t table) AffixRow(value string) table {
	return t.set("affixRow", value)
}

// AutoFillHeight sets the table to auto-calculate height
func (t table) AutoFillHeight(value string) table {
	return t.set("autoFillHeight", value)
}

// AutoGenerateFilter enables query area, auto-generates query form based on column's searchable attribute
func (t table) AutoGenerateFilter(value string) table {
	return t.set("autoGenerateFilter", value)
}

// CanAccessSuperData sets whether the table can access parent data domain values, default is false
func (t table) CanAccessSuperData(value bool) table {
	return t.set("canAccessSuperData", value)
}

// ClassName sets the container CSS class name
func (t table) ClassName(value string) table {
	return t.set("className", value)
}

// Columns sets the table's column information
func (t table) Columns(value ...any) table {
	return t.set("columns", value)
}

// ColumnsTogglable sets whether to show column toggle switch, auto if columns >= 5
func (t table) ColumnsTogglable(value bool) table {
	return t.set("columnsTogglable", value)
}

// CombineFromIndex sets the merge cell configuration, starting from which column
func (t table) CombineFromIndex(value string) table {
	return t.set("combineFromIndex", value)
}

// CombineNum sets the merge cell configuration, number indicates how many columns to merge from left to right
func (t table) CombineNum(value string) table {
	return t.set("combineNum", value)
}

// Data sets the data
func (t table) Data(value any) table {
	return t.set("data", value)
}

// DeferApi sets the lazy load API, used when row data is marked with defer: true
func (t table) DeferApi(value string) table {
	return t.set("deferApi", value)
}

// Disabled sets whether the table is disabled
func (t table) Disabled(value bool) table {
	return t.set("disabled", value)
}

// DisabledOn sets the expression for disabling the table
func (t table) DisabledOn(value string) table {
	return t.set("disabledOn", value)
}

// EditorSetting sets the editor configuration, can be ignored at runtime
func (t table) EditorSetting(value string) table {
	return t.set("editorSetting", value)
}

// Footable sets whether to enable the footer display feature, suitable for mobile display
func (t table) Footable(value bool) table {
	return t.set("footable", value)
}

// FooterClassName sets the CSS class name for the footer container
func (t table) FooterClassName(value string) table {
	return t.set("footerClassName", value)
}

// HeaderClassName sets the CSS class name for the header container
func (t table) HeaderClassName(value string) table {
	return t.set("headerClassName", value)
}

// Hidden sets whether the table is hidden
func (t table) Hidden(value bool) table {
	return t.set("hidden", value)
}

// HiddenOn sets the expression for hiding the table
func (t table) HiddenOn(value string) table {
	return t.set("hiddenOn", value)
}

// ID sets the unique component ID, mainly for log collection
func (t table) ID(value string) table {
	return t.set("id", value)
}

// ItemBadge sets the row badge (Badge)
func (t table) ItemBadge(value string) table {
	return t.set("itemBadge", value)
}

// OnEvent sets the event action configuration
func (t table) OnEvent(value any) table {
	return t.set("onEvent", value)
}

// Placeholder sets the placeholder
func (t table) Placeholder(value string) table {
	return t.set("placeholder", value)
}

// PrefixRow sets the top summary row
func (t table) PrefixRow(value string) table {
	return t.set("prefixRow", value)
}

// Resizable sets whether the column width is adjustable
func (t table) Resizable(value bool) table {
	return t.set("resizable", value)
}

// RowClassNameExpr sets the row style expression
func (t table) RowClassNameExpr(value string) table {
	return t.set("rowClassNameExpr", value)
}

// ShowFooter sets whether to show the footer
func (t table) ShowFooter(value bool) table {
	return t.set("showFooter", value)
}

// ShowHeader sets whether to show the header
func (t table) ShowHeader(value bool) table {
	return t.set("showHeader", value)
}

// Source sets the data source: bind to the current environment variable
func (t table) Source(value string) table {
	return t.set("source", value)
}

// Static sets whether to display statically
func (t table) Static(value bool) table {
	return t.set("static", value)
}

// StaticClassName sets the static display form item class name
func (t table) StaticClassName(value string) table {
	return t.set("staticClassName", value)
}

// StaticInputClassName sets the static display form item value class name
func (t table) StaticInputClassName(value string) table {
	return t.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static display form item label class name
func (t table) StaticLabelClassName(value string) table {
	return t.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display
func (t table) StaticOn(value string) table {
	return t.set("staticOn", value)
}

// StaticPlaceholder sets the static display empty value placeholder
func (t table) StaticPlaceholder(value string) table {
	return t.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema
func (t table) StaticSchema(value string) table {
	return t.set("staticSchema", value)
}

// Style sets the component style
func (t table) Style(value any) table {
	return t.set("style", value)
}

// TableClassName sets the table CSS class name
func (t table) TableClassName(value string) table {
	return t.set("tableClassName", value)
}

// TableLayout sets the table layout, options: fixed | auto
func (t table) TableLayout(value string) table {
	return t.set("tableLayout", value)
}

// TestIdBuilder sets the test ID builder
func (t table) TestIdBuilder(value string) table {
	return t.set("testIdBuilder", value)
}

// Testid sets the test ID
func (t table) Testid(value string) table {
	return t.set("testid", value)
}

// Title sets the title
func (t table) Title(value any) table {
	return t.set("title", value)
}

// ToolbarClassName sets the toolbar CSS class name
func (t table) ToolbarClassName(value string) table {
	return t.set("toolbarClassName", value)
}

// UseMobileUI sets whether to disable mobile styles at the component level
func (t table) UseMobileUI(value bool) table {
	return t.set("useMobileUI", value)
}

// Visible sets whether the table is visible
func (t table) Visible(value bool) table {
	return t.set("visible", value)
}

// VisibleOn sets the expression for visibility
func (t table) VisibleOn(value string) table {
	return t.set("visibleOn", value)
}

// Actions sets the action column configuration
func (t table) Actions(value string) table {
	return t.set("actions", value)
}

// Bordered sets whether to show borders
func (t table) Bordered(value bool) table {
	return t.set("bordered", value)
}

// ChildrenColumnName sets the custom field name for nested data source
func (t table) ChildrenColumnName(value string) table {
	return t.set("childrenColumnName", value)
}

// Expandable sets whether the table row is expandable
func (t table) Expandable(value bool) table {
	return t.set("expandable", value)
}

// Footer sets the footer
func (t table) Footer(value string) table {
	return t.set("footer", value)
}

// KeepItemSelectionOnPageChange sets whether to keep data on page change
func (t table) KeepItemSelectionOnPageChange(value bool) table {
	return t.set("keepItemSelectionOnPageChange", value)
}

// KeyField sets the ID field name for multi-select and nested expansion, default is id
func (t table) KeyField(value string) table {
	return t.set("keyField", value)
}

// LazyRenderAfter sets the number of columns to render at once, default is 100, can improve table rendering performance
func (t table) LazyRenderAfter(value string) table {
	return t.set("lazyRenderAfter", value)
}

// LineHeight sets whether to fix the content row height
func (t table) LineHeight(value string) table {
	return t.set("lineHeight", value)
}

// Loading sets the loading state
func (t table) Loading(value string) table {
	return t.set("loading", value)
}

// MaxKeepItemSelectionLength sets the maximum limit for batch operations
func (t table) MaxKeepItemSelectionLength(value string) table {
	return t.set("maxKeepItemSelectionLength", value)
}

// Messages sets the interface error message configuration (message text configuration, remember this has the lowest priority, if your interface returns msg, the interface return takes precedence)
func (t table) Messages(value string) table {
	return t.set("messages", value)
}

// Multiple sets whether multiple selection is allowed, same as rowSelection.type, compatible with original CRUD attribute, default is multi-select, only works when selectable is set
func (t table) Multiple(value bool) table {
	return t.set("multiple", value)
}

// PopOverContainer sets the specified mount DOM
func (t table) PopOverContainer(value string) table {
	return t.set("popOverContainer", value)
}

// PrimaryField sets the ID field name, same as keyField, compatible with original CRUD attribute
func (t table) PrimaryField(value string) table {
	return t.set("primaryField", value)
}

// QuickSaveApi sets the API for batch saving after quick editing
func (t table) QuickSaveApi(value string) table {
	return t.set("quickSaveApi", value)
}

// QuickSaveItemApi sets the API for immediate saving when quick editing is configured
func (t table) QuickSaveItemApi(value string) table {
	return t.set("quickSaveItemApi", value)
}

// Reload sets the name of the component to reload
func (t table) Reload(value string) table {
	return t.set("reload", value)
}

// RowSelection sets the table selection configuration
func (t table) RowSelection(value string) table {
	return t.set("rowSelection", value)
}

// Selectable sets whether the table is selectable, same as rowSelection, compatible with original CRUD attribute, default is multi-select
func (t table) Selectable(value bool) table {
	return t.set("selectable", value)
}

// ShowBadge sets whether to show row badges
func (t table) ShowBadge(value bool) table {
	return t.set("showBadge", value)
}

// Sticky sets whether the header is sticky
func (t table) Sticky(value bool) table {
	return t.set("sticky", value)
}

// Size sets the table size, supports large, default, small, default is medium size
func (t table) Size(value string) table {
	return t.set("size", value)
}
