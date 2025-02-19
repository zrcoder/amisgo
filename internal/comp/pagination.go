package comp

import "github.com/zrcoder/amisgo/schema"

// Pagination represents the amis Pagination renderer
type Pagination schema.Schema

func NewPagination() Pagination {
	return Pagination{"type": "pagination"}
}

// set sets a field value
func (p Pagination) set(key string, value any) Pagination {
	p[key] = value
	return p
}

// ActivePage sets the current page number
func (p Pagination) ActivePage(value int) Pagination {
	return p.set("activePage", value)
}

// ClassName sets the container CSS class name
func (p Pagination) ClassName(value string) Pagination {
	return p.set("className", value)
}

// Disabled sets whether the pagination is disabled
func (p Pagination) Disabled(value bool) Pagination {
	return p.set("disabled", value)
}

// DisabledOn sets the expression to disable pagination
func (p Pagination) DisabledOn(value string) Pagination {
	return p.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (p Pagination) EditorSetting(value string) Pagination {
	return p.set("editorSetting", value)
}

// HasNext sets whether there is a next page
func (p Pagination) HasNext(value bool) Pagination {
	return p.set("hasNext", value)
}

// Hidden sets whether the pagination is hidden
func (p Pagination) Hidden(value bool) Pagination {
	return p.set("hidden", value)
}

// HiddenOn sets the expression to hide pagination
func (p Pagination) HiddenOn(value string) Pagination {
	return p.set("hiddenOn", value)
}

// ID sets the unique component ID
func (p Pagination) ID(value string) Pagination {
	return p.set("id", value)
}

// Layout sets the pagination layout
func (p Pagination) Layout(value []string) Pagination {
	return p.set("layout", value)
}

// MaxButtons sets the maximum number of pagination buttons
func (p Pagination) MaxButtons(value int) Pagination {
	return p.set("maxButtons", value)
}

// Mode sets the pagination mode
func (p Pagination) Mode(value string) Pagination {
	return p.set("mode", value)
}

// OnEvent sets the event action configuration
func (p Pagination) OnEvent(value any) Pagination {
	return p.set("onEvent", value)
}

// PerPage sets the number of items per page
func (p Pagination) PerPage(value int) Pagination {
	return p.set("perPage", value)
}

// PerPageAvailable sets whether per page options are available
func (p Pagination) PerPageAvailable(value bool) Pagination {
	return p.set("perPageAvailable", value)
}

// PopOverContainerSelector sets the popover container selector
func (p Pagination) PopOverContainerSelector(value string) Pagination {
	return p.set("popOverContainerSelector", value)
}

// ShowPageInput sets whether to show the quick jump input box
func (p Pagination) ShowPageInput(value bool) Pagination {
	return p.set("showPageInput", value)
}

// ShowPerPage sets whether to show per page options
func (p Pagination) ShowPerPage(value bool) Pagination {
	return p.set("showPerPage", value)
}

// Static sets whether the pagination is static
func (p Pagination) Static(value bool) Pagination {
	return p.set("static", value)
}

// StaticClassName sets the static display form item class name
func (p Pagination) StaticClassName(value string) Pagination {
	return p.set("staticClassName", value)
}

// StaticInputClassName sets the static display form item value class name
func (p Pagination) StaticInputClassName(value string) Pagination {
	return p.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static display form item label class name
func (p Pagination) StaticLabelClassName(value string) Pagination {
	return p.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display
func (p Pagination) StaticOn(value string) Pagination {
	return p.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (p Pagination) StaticPlaceholder(value string) Pagination {
	return p.set("staticPlaceholder", value)
}

// StaticSchema sets the schema.Schema for static display mode
func (p Pagination) StaticSchema(value string) Pagination {
	return p.set("staticSchema", value)
}

// Style sets the component style
func (p Pagination) Style(value any) Pagination {
	return p.set("style", value)
}

// TestIdBuilder sets the custom test ID builder
func (p Pagination) TestIdBuilder(value string) Pagination {
	return p.set("testIdBuilder", value)
}

// Testid sets the test ID
func (p Pagination) Testid(value string) Pagination {
	return p.set("testid", value)
}

// Total sets the total number of items
func (p Pagination) Total(value int) Pagination {
	return p.set("total", value)
}

// UseMobileUI sets whether to use mobile UI
func (p Pagination) UseMobileUI(value bool) Pagination {
	return p.set("useMobileUI", value)
}

// Visible sets whether the pagination is visible
func (p Pagination) Visible(value bool) Pagination {
	return p.set("visible", value)
}

// VisibleOn sets the expression to show pagination
func (p Pagination) VisibleOn(value string) Pagination {
	return p.set("visibleOn", value)
}

// Size sets the size of the pagination
func (p Pagination) Size(value string) Pagination {
	return p.set("size", value)
}

// EllipsisPageGap sets the page gap for ellipsis
func (p Pagination) EllipsisPageGap(value int) Pagination {
	return p.set("ellipsisPageGap", value)
}
