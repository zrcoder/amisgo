package comp

import (
	"net/http"

	"github.com/zrcoder/amisgo/internal/servemux"
	"github.com/zrcoder/amisgo/schema"
)

// Crud represents a CRUD (Create, Read, Update, Delete) table and list renderer
type Crud schema.Schema

// NewCrud creates a default CRUD table
func NewCrud(mux *http.ServeMux) Crud {
	return NewCrudTable(mux)
}

// NewCrudTable creates a new CRUDTable instance
func NewCrudTable(mux *http.ServeMux) Crud {
	return Crud{"type": "crud", servemux.Key: mux, "mode": "table"}
}

// NewCrudList creates a new CRUD list instance
func NewCrudList(mux *http.ServeMux) Crud {
	return NewCrudTable(mux).Mode("list")
}

// NewCrudCards creates a new CRUD cards instance
func NewCrudCards(mux *http.ServeMux) Crud {
	return NewCrudTable(mux).Mode("cards")
}

// Set sets a key-value pair and returns the current instance
func (c Crud) set(key string, value any) Crud {
	c[key] = value
	return c
}

// Api sets the initial data retrieval API
func (c Crud) Api(value string) Crud {
	return c.set("api", value)
}

// FetchData sets the API implementation method for retrieving data
func (c Crud) FetchData(getter func() (any, error)) Crud {
	return c.Api(servemux.ServeData(c.mux(), getter))
}

// AutoFillHeight determines whether the content area should occupy the remaining screen space
func (c Crud) AutoFillHeight(value any) Crud {
	return c.set("autoFillHeight", value)
}

// Columns configures the table columns
func (c Crud) Columns(value ...any) Crud {
	return c.set("columns", value)
}

// Disabled enables or disables the component
func (c Crud) Disabled(value bool) Crud {
	return c.set("disabled", value)
}

// Expandable configures whether table rows can be expanded
func (c Crud) Expandable(value bool) Crud {
	return c.set("expandable", value)
}

// Footer sets the table footer
func (c Crud) Footer(value string) Crud {
	return c.set("footer", value)
}

// HeaderToolbar configures the top area of the component
func (c Crud) HeaderToolbar(value ...any) Crud {
	return c.set("headerToolbar", value)
}

// Hidden controls the overall visibility of the component
func (c Crud) Hidden(value bool) Crud {
	return c.set("hidden", value)
}

// Interval sets the auto-refresh time
func (c Crud) Interval(value string) Crud {
	return c.set("interval", value)
}

// Loading sets the loading state
func (c Crud) Loading(value string) Crud {
	return c.set("loading", value)
}

// Name sets the component name
func (c Crud) Name(value string) Crud {
	return c.set("name", value)
}

// Pagination configures pagination settings
func (c Crud) Pagination(value any) Crud {
	return c.set("pagination", value)
}

// Placeholder sets the hint for empty list data
func (c Crud) Placeholder(value string) Crud {
	return c.set("placeholder", value)
}

// RowClassName sets the CSS class name for rows
func (c Crud) RowClassName(value string) Crud {
	return c.set("rowClassName", value)
}

// Sticky determines whether the header is sticky
func (c Crud) Sticky(value bool) Crud {
	return c.set("sticky", value)
}

// Style sets custom inline styles for the component
func (c Crud) Style(value any) Crud {
	return c.set("style", value)
}

// TableLayout sets the table layout
func (c Crud) TableLayout(value string) Crud {
	return c.set("tableLayout", value)
}

// Title sets the table title
func (c Crud) Title(value any) Crud {
	return c.set("title", value)
}

// Width sets the component width
func (c Crud) Width(value string) Crud {
	return c.set("width", value)
}

// Actions configures the operation column
func (c Crud) Actions(value Action) Crud {
	return c.set("actions", value)
}

// AutoJumpToTopOnPagerChange determines whether to automatically scroll to the top when changing pages
func (c Crud) AutoJumpToTopOnPagerChange(value bool) Crud {
	return c.set("autoJumpToTopOnPagerChange", value)
}

// Bordered enables or disables table border display
func (c Crud) Bordered(value bool) Crud {
	return c.set("bordered", value)
}

// CanAccessSuperData determines whether the table can access parent-level data domain values
func (c Crud) CanAccessSuperData(value bool) Crud {
	return c.set("canAccessSuperData", value)
}

// ChildrenColumnName sets a custom field name for nested data sources
func (c Crud) ChildrenColumnName(value string) Crud {
	return c.set("childrenColumnName", value)
}

// ClassName sets the CSS class name for the container
func (c Crud) ClassName(value string) Crud {
	return c.set("className", value)
}

// ColumnsTogglable enables or disables custom column configuration
func (c Crud) ColumnsTogglable(value bool) Crud {
	return c.set("columnsTogglable", value)
}

// EditorSetting configures editor-specific settings
func (c Crud) EditorSetting(value string) Crud {
	return c.set("editorSetting", value)
}

// FooterToolbar configures the bottom area of the component
func (c Crud) FooterToolbar(value ...any) Crud {
	return c.set("footerToolbar", value)
}

// FooterToolbarClassName sets the CSS class name for the footer toolbar
func (c Crud) FooterToolbarClassName(value string) Crud {
	return c.set("footerToolbarClassName", value)
}

// HeaderToolbarClassName sets the CSS class name for the top toolbar area
func (c Crud) HeaderToolbarClassName(value string) Crud {
	return c.set("headerToolbarClassName", value)
}

// HideQuickSaveBtn controls the visibility of the quick save button
func (c Crud) HideQuickSaveBtn(value bool) Crud {
	return c.set("hideQuickSaveBtn", value)
}

// KeyField sets the ID field name for multi-select and nested expansion records
func (c Crud) KeyField(value string) Crud {
	return c.set("keyField", value)
}

// LazyRenderAfter sets the number of columns to render at once
func (c Crud) LazyRenderAfter(value string) Crud {
	return c.set("lazyRenderAfter", value)
}

// LineHeight determines whether to fix the content row height
func (c Crud) LineHeight(value string) Crud {
	return c.set("lineHeight", value)
}

// LoadDataOnce enables or disables front-end single-load mode
func (c Crud) LoadDataOnce(value bool) Crud {
	return c.set("loadDataOnce", value)
}

// LoadType sets the data display mode
func (c Crud) LoadType(value string) Crud {
	return c.set("loadType", value)
}

// LoadingConfig configures the loading state
func (c Crud) LoadingConfig(value string) Crud {
	return c.set("loadingConfig", value)
}

// MaxKeepItemSelectionLength sets the maximum limit for batch operations
func (c Crud) MaxKeepItemSelectionLength(value string) Crud {
	return c.set("maxKeepItemSelectionLength", value)
}

// Messages configures interface error message settings
func (c Crud) Messages(value string) Crud {
	return c.set("messages", value)
}

// Mode specifies the content area display mode (table | list | cards, default is table)
func (c Crud) Mode(value string) Crud {
	return c.set("mode", value)
}

// Multiple enables or disables multi-select data
func (c Crud) Multiple(value bool) Crud {
	return c.set("multiple", value)
}

// OnEvent configures event-driven actions
func (c Crud) OnEvent(value Event) Crud {
	return c.set("onEvent", value)
}

// PageField sets the page number field name for pagination
func (c Crud) PageField(value string) Crud {
	return c.set("pageField", value)
}

// ParsePrimitiveQuery enables or disables Query information conversion
func (c Crud) ParsePrimitiveQuery(value string) Crud {
	return c.set("parsePrimitiveQuery", value)
}

// PerPage sets the number of items loaded per page in infinite scroll mode
func (c Crud) PerPage(value int) Crud {
	return c.set("perPage", value)
}

// PerPageField sets the field name for specifying items per page
func (c Crud) PerPageField(value string) Crud {
	return c.set("perPageField", value)
}

// PopOverContainer specifies the target DOM for mounting
func (c Crud) PopOverContainer(value string) Crud {
	return c.set("popOverContainer", value)
}

// PrimaryField sets the row identifier
func (c Crud) PrimaryField(value string) Crud {
	return c.set("primaryField", value)
}

// QuickSaveApi sets the API for batch saving after quick editing
func (c Crud) QuickSaveApi(value string) Crud {
	return c.set("quickSaveApi", value)
}

// QuickSaveItemApi sets the API for instant saving during quick editing
func (c Crud) QuickSaveItemApi(value string) Crud {
	return c.set("quickSaveItemApi", value)
}

// Reload sets the component name to be reloaded
func (c Crud) Reload(value string) Crud {
	return c.set("reload", value)
}

// RowClassNameExpr sets custom row styling
func (c Crud) RowClassNameExpr(value string) Crud {
	return c.set("rowClassNameExpr", value)
}

// RowSelection configures table row selection settings
func (c Crud) RowSelection(value string) Crud {
	return c.set("rowSelection", value)
}

// SaveOrderApi sets the API for saving order
func (c Crud) SaveOrderApi(value string) Crud {
	return c.set("saveOrderApi", value)
}

// Selectable enables or disables table data selection
func (c Crud) Selectable(value bool) Crud {
	return c.set("selectable", value)
}

// ShowBadge controls the display of badges
func (c Crud) ShowBadge(value bool) Crud {
	return c.set("showBadge", value)
}

// ShowHeader determines whether to display the header
func (c Crud) ShowHeader(value bool) Crud {
	return c.set("showHeader", value)
}

// ShowPagination controls the visibility of pagination
func (c Crud) ShowPagination(value bool) Crud {
	return c.set("showPagination", value)
}

// ShowQuickSaveBtn determines whether to display the quick save button
func (c Crud) ShowQuickSaveBtn(value bool) Crud {
	return c.set("showQuickSaveBtn", value)
}

// ShowToolbar controls the visibility of the toolbar
func (c Crud) ShowToolbar(value bool) Crud {
	return c.set("showToolbar", value)
}

// Source sets the data source
func (c Crud) Source(value string) Crud {
	return c.set("source", value)
}

// Toolbar specifies the table header toolbar
func (c Crud) Toolbar(value string) Crud {
	return c.set("toolbar", value)
}

// ToolbarClassName sets the CSS class name for the toolbar
func (c Crud) ToolbarClassName(value string) Crud {
	return c.set("toolbarClassName", value)
}

// Unfoldable enables or disables the ability to expand or collapse
func (c Crud) Unfoldable(value bool) Crud {
	return c.set("unfoldable", value)
}

// Validate sets validation configuration
func (c Crud) Validate(value string) Crud {
	return c.set("validate", value)
}

// ValueField sets the field name for table data
func (c Crud) ValueField(value string) Crud {
	return c.set("valueField", value)
}

// WidthMode sets the column width mode
func (c Crud) WidthMode(value string) Crud {
	return c.set("widthMode", value)
}

// AffixFooter determines whether to fix the footer at the bottom of the viewport
func (c Crud) AffixFooter(value bool) Crud {
	return c.set("affixFooter", value)
}

// AffixHeader determines whether to fix the header at the top of the viewport
func (c Crud) AffixHeader(value bool) Crud {
	return c.set("affixHeader", value)
}

// AlwaysShowPagination controls whether pagination is always visible
func (c Crud) AlwaysShowPagination(value bool) Crud {
	return c.set("alwaysShowPagination", value)
}

// AutoGenerateFilter enables automatic generation of query condition form
func (c Crud) AutoGenerateFilter(value any) Crud {
	return c.set("autoGenerateFilter", value)
}

// BulkActions configures batch operation actions
func (c Crud) BulkActions(value ...Action) Crud {
	return c.set("bulkActions", value)
}

// Card sets the card configuration
func (c Crud) Card(value any) Crud {
	return c.set("card", value)
}

// CheckOnItemClick determines whether to select a card when clicking on it
func (c Crud) CheckOnItemClick(value bool) Crud {
	return c.set("checkOnItemClick", value)
}

// DefaultParams sets default parameters for the component
func (c Crud) DefaultParams(value string) Crud {
	return c.set("defaultParams", value)
}

// DeferApi sets the lazy-loading API
func (c Crud) DeferApi(value string) Crud {
	return c.set("deferApi", value)
}

// DisabledOn sets a conditional expression for disabling the component
func (c Crud) DisabledOn(value string) Crud {
	return c.set("disabledOn", value)
}

// Draggable enables or disables sorting through drag and drop
func (c Crud) Draggable(value bool) Crud {
	return c.set("draggable", value)
}

// DraggableOn sets a conditional expression for drag and drop sorting
func (c Crud) DraggableOn(value string) Crud {
	return c.set("draggableOn", value)
}

// ExpandConfig sets the configuration for row expansion
func (c Crud) ExpandConfig(value string) Crud {
	return c.set("expandConfig", value)
}

// Filter sets the filter form configuration
func (c Crud) Filter(value Form) Crud {
	return c.set("filter", value)
}

// FilterDefaultVisible determines the default visibility of the filter
func (c Crud) FilterDefaultVisible(value bool) Crud {
	return c.set("filterDefaultVisible", value)
}

// FilterTogglable configures the filter toggle behavior
// Supports boolean or a configuration object with labels and icons
// Default is false
func (c Crud) FilterTogglable(value any) Crud {
	return c.set("filterTogglable", value)
}

// FooterClassName sets the CSS class name for the footer
func (c Crud) FooterClassName(value string) Crud {
	return c.set("footerClassName", value)
}

// Header configures the top area of the component
func (c Crud) Header(value any) Crud {
	return c.set("header", value)
}

// HeaderClassName sets the CSS class name for the header
func (c Crud) HeaderClassName(value string) Crud {
	return c.set("headerClassName", value)
}

// HiddenOn sets a conditional expression for hiding the component
func (c Crud) HiddenOn(value string) Crud {
	return c.set("hiddenOn", value)
}

// HideCheckToggler controls the visibility of the checkbox toggler
func (c Crud) HideCheckToggler(value bool) Crud {
	return c.set("hideCheckToggler", value)
}

// InitFetch determines whether to fetch data initially
func (c Crud) InitFetch(value bool) Crud {
	return c.set("initFetch", value)
}

// InitFetchOn sets a conditional expression for initial data fetching
func (c Crud) InitFetchOn(value string) Crud {
	return c.set("initFetchOn", value)
}

// InnerClassName sets the CSS class name for the internal DOM
func (c Crud) InnerClassName(value string) Crud {
	return c.set("innerClassName", value)
}

// ItemActions configures single-item actions
func (c Crud) ItemActions(value Action) Crud {
	return c.set("itemActions", value)
}

// ItemCheckableOn sets constraints for batch operations
func (c Crud) ItemCheckableOn(value string) Crud {
	return c.set("itemCheckableOn", value)
}

// ItemClassName sets the CSS class name for cards
func (c Crud) ItemClassName(value string) Crud {
	return c.set("itemClassName", value)
}

// ItemDraggableOn determines whether items can be drag-sorted
func (c Crud) ItemDraggableOn(value string) Crud {
	return c.set("itemDraggableOn", value)
}

// KeepItemSelectionOnPageChange preserves user selection when changing pages
func (c Crud) KeepItemSelectionOnPageChange(value bool) Crud {
	return c.set("keepItemSelectionOnPageChange", value)
}

// LabelTpl sets the text for selected items
func (c Crud) LabelTpl(value string) Crud {
	return c.set("labelTpl", value)
}

// LoadDataOnceFetchOnFilter configures filtering when loading data once
func (c Crud) LoadDataOnceFetchOnFilter(value bool) Crud {
	return c.set("loadDataOnceFetchOnFilter", value)
}

// MasonryLayout enables or disables masonry (waterfall) layout
func (c Crud) MasonryLayout(value bool) Crud {
	return c.set("masonryLayout", value)
}

// MatchFunc sets a custom search matching function
func (c Crud) MatchFunc(value string) Crud {
	return c.set("matchFunc", value)
}

// OrderBy sets the default sorting field
func (c Crud) OrderBy(value string) Crud {
	return c.set("orderBy", value)
}

// Query sets the query fields
func (c Crud) Query(value string) Crud {
	return c.set("query", value)
}

// Render determines whether to render the component
func (c Crud) Render(value bool) Crud {
	return c.set("render", value)
}

// ResetPage controls whether to reset pagination
func (c Crud) ResetPage(value bool) Crud {
	return c.set("resetPage", value)
}

// Searchable enables or disables search functionality
func (c Crud) Searchable(value bool) Crud {
	return c.set("searchable", value)
}

// SearchPlaceholder sets the search input placeholder text
func (c Crud) SearchPlaceholder(value string) Crud {
	return c.set("searchPlaceholder", value)
}

// ShowErrorMsg controls the visibility of error messages
func (c Crud) ShowErrorMsg(value bool) Crud {
	return c.set("showErrorMsg", value)
}

// ShowSearch controls the visibility of the search input
func (c Crud) ShowSearch(value bool) Crud {
	return c.set("showSearch", value)
}

// ShowSort controls the visibility of sorting options
func (c Crud) ShowSort(value bool) Crud {
	return c.set("showSort", value)
}

// ShowSwitch controls the visibility of the switch button
func (c Crud) ShowSwitch(value bool) Crud {
	return c.set("showSwitch", value)
}

// Size sets the component size
func (c Crud) Size(value string) Crud {
	return c.set("size", value)
}

// Sortable enables or disables sorting functionality
func (c Crud) Sortable(value bool) Crud {
	return c.set("sortable", value)
}

// SubmitText sets the text for the submit button
func (c Crud) SubmitText(value string) Crud {
	return c.set("submitText", value)
}

// SyncLocation determines whether to synchronize with URL
func (c Crud) SyncLocation(value bool) Crud {
	return c.set("syncLocation", value)
}

// Translations sets multilingual translations
func (c Crud) Translations(value map[string]string) Crud {
	return c.set("translations", value)
}

// VerticalAlign sets the vertical alignment method
func (c Crud) VerticalAlign(value string) Crud {
	return c.set("verticalAlign", value)
}

// Visible controls the overall visibility of the component
func (c Crud) Visible(value bool) Crud {
	return c.set("visible", value)
}

// VisibleOn sets a conditional expression for component visibility
func (c Crud) VisibleOn(value string) Crud {
	return c.set("visibleOn", value)
}

// WrapItemClassName sets the CSS class name for wrap items
func (c Crud) WrapItemClassName(value string) Crud {
	return c.set("wrapItemClassName", value)
}

// ExtraAction sets extra actions
func (c Crud) ExtraAction(value any) Crud {
	return c.set("extraAction", value)
}

// AutoSaveOnChange sets whether to automatically save changes
func (c Crud) AutoSaveOnChange(value bool) Crud {
	return c.set("autoSaveOnChange", value)
}

// ID sets the component's unique ID
func (c Crud) ID(value string) Crud {
	return c.set("id", value)
}

// ItemAction sets the behavior of clicking on a list item
func (c Crud) ItemAction(value string) Crud {
	return c.set("itemAction", value)
}

// ListItem sets the configuration for displaying a single item
func (c Crud) ListItem(value any) Crud {
	return c.set("listItem", value)
}

// QuickSaveItemActions sets the configuration for quick save actions
func (c Crud) QuickSaveItemActions(value Action) Crud {
	return c.set("quickSaveItemActions", value)
}

// QuickSaveSuccessMessage sets the success message for quick save
func (c Crud) QuickSaveSuccessMessage(value string) Crud {
	return c.set("quickSaveSuccessMessage", value)
}

// RefreshInterval sets the automatic refresh interval
func (c Crud) RefreshInterval(value string) Crud {
	return c.set("refreshInterval", value)
}

// RemoteFilter sets whether to support remote filtering
func (c Crud) RemoteFilter(value bool) Crud {
	return c.set("remoteFilter", value)
}

// RenderType sets the rendering type
func (c Crud) RenderType(value string) Crud {
	return c.set("renderType", value)
}

// ResetValueOnHidden sets whether to reset the value when hidden
func (c Crud) ResetValueOnHidden(value bool) Crud {
	return c.set("resetValueOnHidden", value)
}

// SearchFormClassName sets the CSS class name for the search form
func (c Crud) SearchFormClassName(value string) Crud {
	return c.set("searchFormClassName", value)
}

// SearchSchema sets the search form schema.Schema
func (c Crud) SearchSchema(value any) Crud {
	return c.set("searchSchema", value)
}

// SortSettings sets the sorting configuration
func (c Crud) SortSettings(value string) Crud {
	return c.set("sortSettings", value)
}

// SortFieldName sets the sorting field name
func (c Crud) SortFieldName(value string) Crud {
	return c.set("sortFieldName", value)
}

// Stateful sets whether to keep the state
func (c Crud) Stateful(value bool) Crud {
	return c.set("stateful", value)
}

// TableClassName sets the CSS class name for the table
func (c Crud) TableClassName(value string) Crud {
	return c.set("tableClassName", value)
}

// Tooltip sets the tooltip configuration
func (c Crud) Tooltip(value string) Crud {
	return c.set("tooltip", value)
}

// TransformItems sets the data transformation
func (c Crud) TransformItems(value ...any) Crud {
	return c.set("transformItems", value)
}

// Virtualized sets whether to use virtualization
func (c Crud) Virtualized(value bool) Crud {
	return c.set("virtualized", value)
}

// EmptyText sets the text for when there is no data
func (c Crud) EmptyText(value string) Crud {
	return c.set("emptyText", value)
}

// Components sets custom components
func (c Crud) Components(value any) Crud {
	return c.set("components", value)
}

// Other CRUD2List specific fields and methods
func (c Crud) CustomField1(value string) Crud {
	return c.set("customField1", value)
}

// Other crudTable specific fields and methods
func (c Crud) CustomField2(value string) Crud {
	return c.set("customField2", value)
}

func (c Crud) mux() *http.ServeMux {
	return c[servemux.Key].(*http.ServeMux)
}
