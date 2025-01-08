package comp

import (
	"github.com/zrcoder/amisgo/internal/servermux"
	"github.com/zrcoder/amisgo/model"
)

// Crud represents a CRUD (Create, Read, Update, Delete) table and list renderer
type crud model.Schema

// Crud creates a default CRUD table
func Crud() crud {
	return CrudTable()
}

// CrudTable creates a new CRUDTable instance
func CrudTable() crud {
	return crud{}.set("type", "crud").Mode("table")
}

// CrudList creates a new CRUD list instance
func CrudList() crud {
	return crud{}.set("type", "crud").Mode("list")
}

// CrudCards creates a new CRUD cards instance
func CrudCards() crud {
	return crud{}.set("type", "crud").Mode("cards")
}

// Set sets a key-value pair and returns the current instance
func (c crud) set(key string, value any) crud {
	c[key] = value
	return c
}

// Api sets the initial data retrieval API
func (c crud) Api(value string) crud {
	return c.set("api", value)
}

// FetchData sets the API implementation method for retrieving data
func (c crud) FetchData(getter func() (any, error)) crud {
	return c.Api(servermux.ServeData(getter))
}

// AutoFillHeight determines whether the content area should occupy the remaining screen space
func (c crud) AutoFillHeight(value any) crud {
	return c.set("autoFillHeight", value)
}

// Columns configures the table columns
func (c crud) Columns(value ...any) crud {
	return c.set("columns", value)
}

// Disabled enables or disables the component
func (c crud) Disabled(value bool) crud {
	return c.set("disabled", value)
}

// Expandable configures whether table rows can be expanded
func (c crud) Expandable(value bool) crud {
	return c.set("expandable", value)
}

// Footer sets the table footer
func (c crud) Footer(value string) crud {
	return c.set("footer", value)
}

// HeaderToolbar configures the top area of the component
func (c crud) HeaderToolbar(value ...any) crud {
	return c.set("headerToolbar", value)
}

// Hidden controls the overall visibility of the component
func (c crud) Hidden(value bool) crud {
	return c.set("hidden", value)
}

// Interval sets the auto-refresh time
func (c crud) Interval(value string) crud {
	return c.set("interval", value)
}

// Loading sets the loading state
func (c crud) Loading(value string) crud {
	return c.set("loading", value)
}

// Name sets the component name
func (c crud) Name(value string) crud {
	return c.set("name", value)
}

// Pagination configures pagination settings
func (c crud) Pagination(value any) crud {
	return c.set("pagination", value)
}

// Placeholder sets the hint for empty list data
func (c crud) Placeholder(value string) crud {
	return c.set("placeholder", value)
}

// RowClassName sets the CSS class name for rows
func (c crud) RowClassName(value string) crud {
	return c.set("rowClassName", value)
}

// Sticky determines whether the header is sticky
func (c crud) Sticky(value bool) crud {
	return c.set("sticky", value)
}

// Style sets custom inline styles for the component
func (c crud) Style(value any) crud {
	return c.set("style", value)
}

// TableLayout sets the table layout
func (c crud) TableLayout(value string) crud {
	return c.set("tableLayout", value)
}

// Title sets the table title
func (c crud) Title(value any) crud {
	return c.set("title", value)
}

// Width sets the component width
func (c crud) Width(value string) crud {
	return c.set("width", value)
}

// Actions configures the operation column
func (c crud) Actions(value string) crud {
	return c.set("actions", value)
}

// AutoJumpToTopOnPagerChange determines whether to automatically scroll to the top when changing pages
func (c crud) AutoJumpToTopOnPagerChange(value bool) crud {
	return c.set("autoJumpToTopOnPagerChange", value)
}

// Bordered enables or disables table border display
func (c crud) Bordered(value bool) crud {
	return c.set("bordered", value)
}

// CanAccessSuperData determines whether the table can access parent-level data domain values
func (c crud) CanAccessSuperData(value bool) crud {
	return c.set("canAccessSuperData", value)
}

// ChildrenColumnName sets a custom field name for nested data sources
func (c crud) ChildrenColumnName(value string) crud {
	return c.set("childrenColumnName", value)
}

// ClassName sets the CSS class name for the container
func (c crud) ClassName(value string) crud {
	return c.set("className", value)
}

// ColumnsTogglable enables or disables custom column configuration
func (c crud) ColumnsTogglable(value bool) crud {
	return c.set("columnsTogglable", value)
}

// EditorSetting configures editor-specific settings
func (c crud) EditorSetting(value string) crud {
	return c.set("editorSetting", value)
}

// FooterToolbar configures the bottom area of the component
func (c crud) FooterToolbar(value ...any) crud {
	return c.set("footerToolbar", value)
}

// FooterToolbarClassName sets the CSS class name for the footer toolbar
func (c crud) FooterToolbarClassName(value string) crud {
	return c.set("footerToolbarClassName", value)
}

// HeaderToolbarClassName sets the CSS class name for the top toolbar area
func (c crud) HeaderToolbarClassName(value string) crud {
	return c.set("headerToolbarClassName", value)
}

// HideQuickSaveBtn controls the visibility of the quick save button
func (c crud) HideQuickSaveBtn(value bool) crud {
	return c.set("hideQuickSaveBtn", value)
}

// KeyField sets the ID field name for multi-select and nested expansion records
func (c crud) KeyField(value string) crud {
	return c.set("keyField", value)
}

// LazyRenderAfter sets the number of columns to render at once
func (c crud) LazyRenderAfter(value string) crud {
	return c.set("lazyRenderAfter", value)
}

// LineHeight determines whether to fix the content row height
func (c crud) LineHeight(value string) crud {
	return c.set("lineHeight", value)
}

// LoadDataOnce enables or disables front-end single-load mode
func (c crud) LoadDataOnce(value bool) crud {
	return c.set("loadDataOnce", value)
}

// LoadType sets the data display mode
func (c crud) LoadType(value string) crud {
	return c.set("loadType", value)
}

// LoadingConfig configures the loading state
func (c crud) LoadingConfig(value string) crud {
	return c.set("loadingConfig", value)
}

// MaxKeepItemSelectionLength sets the maximum limit for batch operations
func (c crud) MaxKeepItemSelectionLength(value string) crud {
	return c.set("maxKeepItemSelectionLength", value)
}

// Messages configures interface error message settings
func (c crud) Messages(value string) crud {
	return c.set("messages", value)
}

// Mode specifies the content area display mode (table | list | cards, default is table)
func (c crud) Mode(value string) crud {
	return c.set("mode", value)
}

// Multiple enables or disables multi-select data
func (c crud) Multiple(value bool) crud {
	return c.set("multiple", value)
}

// OnEvent configures event-driven actions
func (c crud) OnEvent(value any) crud {
	return c.set("onEvent", value)
}

// PageField sets the page number field name for pagination
func (c crud) PageField(value string) crud {
	return c.set("pageField", value)
}

// ParsePrimitiveQuery enables or disables Query information conversion
func (c crud) ParsePrimitiveQuery(value string) crud {
	return c.set("parsePrimitiveQuery", value)
}

// PerPage sets the number of items loaded per page in infinite scroll mode
func (c crud) PerPage(value int) crud {
	return c.set("perPage", value)
}

// PerPageField sets the field name for specifying items per page
func (c crud) PerPageField(value string) crud {
	return c.set("perPageField", value)
}

// PopOverContainer specifies the target DOM for mounting
func (c crud) PopOverContainer(value string) crud {
	return c.set("popOverContainer", value)
}

// PrimaryField sets the row identifier
func (c crud) PrimaryField(value string) crud {
	return c.set("primaryField", value)
}

// QuickSaveApi sets the API for batch saving after quick editing
func (c crud) QuickSaveApi(value string) crud {
	return c.set("quickSaveApi", value)
}

// QuickSaveItemApi sets the API for instant saving during quick editing
func (c crud) QuickSaveItemApi(value string) crud {
	return c.set("quickSaveItemApi", value)
}

// Reload sets the component name to be reloaded
func (c crud) Reload(value string) crud {
	return c.set("reload", value)
}

// RowClassNameExpr sets custom row styling
func (c crud) RowClassNameExpr(value string) crud {
	return c.set("rowClassNameExpr", value)
}

// RowSelection configures table row selection settings
func (c crud) RowSelection(value string) crud {
	return c.set("rowSelection", value)
}

// SaveOrderApi sets the API for saving order
func (c crud) SaveOrderApi(value string) crud {
	return c.set("saveOrderApi", value)
}

// Selectable enables or disables table data selection
func (c crud) Selectable(value bool) crud {
	return c.set("selectable", value)
}

// ShowBadge controls the display of badges
func (c crud) ShowBadge(value bool) crud {
	return c.set("showBadge", value)
}

// ShowHeader determines whether to display the header
func (c crud) ShowHeader(value bool) crud {
	return c.set("showHeader", value)
}

// ShowPagination controls the visibility of pagination
func (c crud) ShowPagination(value bool) crud {
	return c.set("showPagination", value)
}

// ShowQuickSaveBtn determines whether to display the quick save button
func (c crud) ShowQuickSaveBtn(value bool) crud {
	return c.set("showQuickSaveBtn", value)
}

// ShowToolbar controls the visibility of the toolbar
func (c crud) ShowToolbar(value bool) crud {
	return c.set("showToolbar", value)
}

// Source sets the data source
func (c crud) Source(value string) crud {
	return c.set("source", value)
}

// Toolbar specifies the table header toolbar
func (c crud) Toolbar(value string) crud {
	return c.set("toolbar", value)
}

// ToolbarClassName sets the CSS class name for the toolbar
func (c crud) ToolbarClassName(value string) crud {
	return c.set("toolbarClassName", value)
}

// Unfoldable enables or disables the ability to expand or collapse
func (c crud) Unfoldable(value bool) crud {
	return c.set("unfoldable", value)
}

// Validate sets validation configuration
func (c crud) Validate(value string) crud {
	return c.set("validate", value)
}

// ValueField sets the field name for table data
func (c crud) ValueField(value string) crud {
	return c.set("valueField", value)
}

// WidthMode sets the column width mode
func (c crud) WidthMode(value string) crud {
	return c.set("widthMode", value)
}

// AffixFooter determines whether to fix the footer at the bottom of the viewport
func (c crud) AffixFooter(value bool) crud {
	return c.set("affixFooter", value)
}

// AffixHeader determines whether to fix the header at the top of the viewport
func (c crud) AffixHeader(value bool) crud {
	return c.set("affixHeader", value)
}

// AlwaysShowPagination controls whether pagination is always visible
func (c crud) AlwaysShowPagination(value bool) crud {
	return c.set("alwaysShowPagination", value)
}

// AutoGenerateFilter enables automatic generation of query condition form
func (c crud) AutoGenerateFilter(value any) crud {
	return c.set("autoGenerateFilter", value)
}

// BulkActions configures batch operation actions
func (c crud) BulkActions(value ...any) crud {
	return c.set("bulkActions", value)
}

// Card sets the card configuration
func (c crud) Card(value any) crud {
	return c.set("card", value)
}

// CheckOnItemClick determines whether to select a card when clicking on it
func (c crud) CheckOnItemClick(value bool) crud {
	return c.set("checkOnItemClick", value)
}

// DefaultParams sets default parameters for the component
func (c crud) DefaultParams(value string) crud {
	return c.set("defaultParams", value)
}

// DeferApi sets the lazy-loading API
func (c crud) DeferApi(value string) crud {
	return c.set("deferApi", value)
}

// DisabledOn sets a conditional expression for disabling the component
func (c crud) DisabledOn(value string) crud {
	return c.set("disabledOn", value)
}

// Draggable enables or disables sorting through drag and drop
func (c crud) Draggable(value bool) crud {
	return c.set("draggable", value)
}

// DraggableOn sets a conditional expression for drag and drop sorting
func (c crud) DraggableOn(value string) crud {
	return c.set("draggableOn", value)
}

// ExpandConfig sets the configuration for row expansion
func (c crud) ExpandConfig(value string) crud {
	return c.set("expandConfig", value)
}

// Filter sets the filter form configuration
func (c crud) Filter(value form) crud {
	return c.set("filter", value)
}

// FilterDefaultVisible determines the default visibility of the filter
func (c crud) FilterDefaultVisible(value bool) crud {
	return c.set("filterDefaultVisible", value)
}

// FilterTogglable configures the filter toggle behavior
// Supports boolean or a configuration object with labels and icons
// Default is false
func (c crud) FilterTogglable(value any) crud {
	return c.set("filterTogglable", value)
}

// FooterClassName sets the CSS class name for the footer
func (c crud) FooterClassName(value string) crud {
	return c.set("footerClassName", value)
}

// Header configures the top area of the component
func (c crud) Header(value any) crud {
	return c.set("header", value)
}

// HeaderClassName sets the CSS class name for the header
func (c crud) HeaderClassName(value string) crud {
	return c.set("headerClassName", value)
}

// HiddenOn sets a conditional expression for hiding the component
func (c crud) HiddenOn(value string) crud {
	return c.set("hiddenOn", value)
}

// HideCheckToggler controls the visibility of the checkbox toggler
func (c crud) HideCheckToggler(value bool) crud {
	return c.set("hideCheckToggler", value)
}

// InitFetch determines whether to fetch data initially
func (c crud) InitFetch(value bool) crud {
	return c.set("initFetch", value)
}

// InitFetchOn sets a conditional expression for initial data fetching
func (c crud) InitFetchOn(value string) crud {
	return c.set("initFetchOn", value)
}

// InnerClassName sets the CSS class name for the internal DOM
func (c crud) InnerClassName(value string) crud {
	return c.set("innerClassName", value)
}

// ItemActions configures single-item actions
func (c crud) ItemActions(value string) crud {
	return c.set("itemActions", value)
}

// ItemCheckableOn sets constraints for batch operations
func (c crud) ItemCheckableOn(value string) crud {
	return c.set("itemCheckableOn", value)
}

// ItemClassName sets the CSS class name for cards
func (c crud) ItemClassName(value string) crud {
	return c.set("itemClassName", value)
}

// ItemDraggableOn determines whether items can be drag-sorted
func (c crud) ItemDraggableOn(value string) crud {
	return c.set("itemDraggableOn", value)
}

// KeepItemSelectionOnPageChange preserves user selection when changing pages
func (c crud) KeepItemSelectionOnPageChange(value bool) crud {
	return c.set("keepItemSelectionOnPageChange", value)
}

// LabelTpl sets the text for selected items
func (c crud) LabelTpl(value string) crud {
	return c.set("labelTpl", value)
}

// LoadDataOnceFetchOnFilter configures filtering when loading data once
func (c crud) LoadDataOnceFetchOnFilter(value bool) crud {
	return c.set("loadDataOnceFetchOnFilter", value)
}

// MasonryLayout enables or disables masonry (waterfall) layout
func (c crud) MasonryLayout(value bool) crud {
	return c.set("masonryLayout", value)
}

// MatchFunc sets a custom search matching function
func (c crud) MatchFunc(value string) crud {
	return c.set("matchFunc", value)
}

// OrderBy sets the default sorting field
func (c crud) OrderBy(value string) crud {
	return c.set("orderBy", value)
}

// Query sets the query fields
func (c crud) Query(value string) crud {
	return c.set("query", value)
}

// Render determines whether to render the component
func (c crud) Render(value bool) crud {
	return c.set("render", value)
}

// ResetPage controls whether to reset pagination
func (c crud) ResetPage(value bool) crud {
	return c.set("resetPage", value)
}

// Searchable enables or disables search functionality
func (c crud) Searchable(value bool) crud {
	return c.set("searchable", value)
}

// SearchPlaceholder sets the search input placeholder text
func (c crud) SearchPlaceholder(value string) crud {
	return c.set("searchPlaceholder", value)
}

// ShowErrorMsg controls the visibility of error messages
func (c crud) ShowErrorMsg(value bool) crud {
	return c.set("showErrorMsg", value)
}

// ShowSearch controls the visibility of the search input
func (c crud) ShowSearch(value bool) crud {
	return c.set("showSearch", value)
}

// ShowSort controls the visibility of sorting options
func (c crud) ShowSort(value bool) crud {
	return c.set("showSort", value)
}

// ShowSwitch controls the visibility of the switch button
func (c crud) ShowSwitch(value bool) crud {
	return c.set("showSwitch", value)
}

// Size sets the component size
func (c crud) Size(value string) crud {
	return c.set("size", value)
}

// Sortable enables or disables sorting functionality
func (c crud) Sortable(value bool) crud {
	return c.set("sortable", value)
}

// SubmitText sets the text for the submit button
func (c crud) SubmitText(value string) crud {
	return c.set("submitText", value)
}

// SyncLocation determines whether to synchronize with URL
func (c crud) SyncLocation(value bool) crud {
	return c.set("syncLocation", value)
}

// Translations sets multilingual translations
func (c crud) Translations(value map[string]string) crud {
	return c.set("translations", value)
}

// VerticalAlign sets the vertical alignment method
func (c crud) VerticalAlign(value string) crud {
	return c.set("verticalAlign", value)
}

// Visible controls the overall visibility of the component
func (c crud) Visible(value bool) crud {
	return c.set("visible", value)
}

// VisibleOn sets a conditional expression for component visibility
func (c crud) VisibleOn(value string) crud {
	return c.set("visibleOn", value)
}

// WrapItemClassName sets the CSS class name for wrap items
func (c crud) WrapItemClassName(value string) crud {
	return c.set("wrapItemClassName", value)
}

// ExtraAction sets extra actions
func (c crud) ExtraAction(value any) crud {
	return c.set("extraAction", value)
}

// AutoSaveOnChange sets whether to automatically save changes
func (c crud) AutoSaveOnChange(value bool) crud {
	return c.set("autoSaveOnChange", value)
}

// ID sets the component's unique ID
func (c crud) ID(value string) crud {
	return c.set("id", value)
}

// ItemAction sets the behavior of clicking on a list item
func (c crud) ItemAction(value string) crud {
	return c.set("itemAction", value)
}

// ListItem sets the configuration for displaying a single item
func (c crud) ListItem(value any) crud {
	return c.set("listItem", value)
}

// QuickSaveItemActions sets the configuration for quick save actions
func (c crud) QuickSaveItemActions(value string) crud {
	return c.set("quickSaveItemActions", value)
}

// QuickSaveSuccessMessage sets the success message for quick save
func (c crud) QuickSaveSuccessMessage(value string) crud {
	return c.set("quickSaveSuccessMessage", value)
}

// RefreshInterval sets the automatic refresh interval
func (c crud) RefreshInterval(value string) crud {
	return c.set("refreshInterval", value)
}

// RemoteFilter sets whether to support remote filtering
func (c crud) RemoteFilter(value bool) crud {
	return c.set("remoteFilter", value)
}

// RenderType sets the rendering type
func (c crud) RenderType(value string) crud {
	return c.set("renderType", value)
}

// ResetValueOnHidden sets whether to reset the value when hidden
func (c crud) ResetValueOnHidden(value bool) crud {
	return c.set("resetValueOnHidden", value)
}

// SearchFormClassName sets the CSS class name for the search form
func (c crud) SearchFormClassName(value string) crud {
	return c.set("searchFormClassName", value)
}

// SearchSchema sets the search form schema
func (c crud) SearchSchema(value any) crud {
	return c.set("searchSchema", value)
}

// SortSettings sets the sorting configuration
func (c crud) SortSettings(value string) crud {
	return c.set("sortSettings", value)
}

// SortFieldName sets the sorting field name
func (c crud) SortFieldName(value string) crud {
	return c.set("sortFieldName", value)
}

// Stateful sets whether to keep the state
func (c crud) Stateful(value bool) crud {
	return c.set("stateful", value)
}

// TableClassName sets the CSS class name for the table
func (c crud) TableClassName(value string) crud {
	return c.set("tableClassName", value)
}

// Tooltip sets the tooltip configuration
func (c crud) Tooltip(value string) crud {
	return c.set("tooltip", value)
}

// TransformItems sets the data transformation
func (c crud) TransformItems(value ...any) crud {
	return c.set("transformItems", value)
}

// Virtualized sets whether to use virtualization
func (c crud) Virtualized(value bool) crud {
	return c.set("virtualized", value)
}

// EmptyText sets the text for when there is no data
func (c crud) EmptyText(value string) crud {
	return c.set("emptyText", value)
}

// Components sets custom components
func (c crud) Components(value any) crud {
	return c.set("components", value)
}

// Other CRUD2List specific fields and methods
func (c crud) CustomField1(value string) crud {
	return c.set("customField1", value)
}

// Other crudTable specific fields and methods
func (c crud) CustomField2(value string) crud {
	return c.set("customField2", value)
}
