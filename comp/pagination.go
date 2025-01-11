package comp

import "github.com/zrcoder/amisgo/model"

// pagination represents the amis pagination renderer
type pagination model.Schema

// Pagination creates a new Pagination instance
func Pagination() pagination {
	return pagination{"type": "pagination"}
}

// set sets a field value
func (p pagination) set(key string, value any) pagination {
	p[key] = value
	return p
}

// ActivePage sets the current page number
func (p pagination) ActivePage(value int) pagination {
	return p.set("activePage", value)
}

// ClassName sets the container CSS class name
func (p pagination) ClassName(value string) pagination {
	return p.set("className", value)
}

// Disabled sets whether the pagination is disabled
func (p pagination) Disabled(value bool) pagination {
	return p.set("disabled", value)
}

// DisabledOn sets the expression to disable pagination
func (p pagination) DisabledOn(value string) pagination {
	return p.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (p pagination) EditorSetting(value string) pagination {
	return p.set("editorSetting", value)
}

// HasNext sets whether there is a next page
func (p pagination) HasNext(value bool) pagination {
	return p.set("hasNext", value)
}

// Hidden sets whether the pagination is hidden
func (p pagination) Hidden(value bool) pagination {
	return p.set("hidden", value)
}

// HiddenOn sets the expression to hide pagination
func (p pagination) HiddenOn(value string) pagination {
	return p.set("hiddenOn", value)
}

// ID sets the unique component ID
func (p pagination) ID(value string) pagination {
	return p.set("id", value)
}

// Layout sets the pagination layout
func (p pagination) Layout(value []string) pagination {
	return p.set("layout", value)
}

// MaxButtons sets the maximum number of pagination buttons
func (p pagination) MaxButtons(value int) pagination {
	return p.set("maxButtons", value)
}

// Mode sets the pagination mode
func (p pagination) Mode(value string) pagination {
	return p.set("mode", value)
}

// OnEvent sets the event action configuration
func (p pagination) OnEvent(value any) pagination {
	return p.set("onEvent", value)
}

// PerPage sets the number of items per page
func (p pagination) PerPage(value int) pagination {
	return p.set("perPage", value)
}

// PerPageAvailable sets whether per page options are available
func (p pagination) PerPageAvailable(value bool) pagination {
	return p.set("perPageAvailable", value)
}

// PopOverContainerSelector sets the popover container selector
func (p pagination) PopOverContainerSelector(value string) pagination {
	return p.set("popOverContainerSelector", value)
}

// ShowPageInput sets whether to show the quick jump input box
func (p pagination) ShowPageInput(value bool) pagination {
	return p.set("showPageInput", value)
}

// ShowPerPage sets whether to show per page options
func (p pagination) ShowPerPage(value bool) pagination {
	return p.set("showPerPage", value)
}

// Static sets whether the pagination is static
func (p pagination) Static(value bool) pagination {
	return p.set("static", value)
}

// StaticClassName sets the static display form item class name
func (p pagination) StaticClassName(value string) pagination {
	return p.set("staticClassName", value)
}

// StaticInputClassName sets the static display form item value class name
func (p pagination) StaticInputClassName(value string) pagination {
	return p.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static display form item label class name
func (p pagination) StaticLabelClassName(value string) pagination {
	return p.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display
func (p pagination) StaticOn(value string) pagination {
	return p.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (p pagination) StaticPlaceholder(value string) pagination {
	return p.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display mode
func (p pagination) StaticSchema(value string) pagination {
	return p.set("staticSchema", value)
}

// Style sets the component style
func (p pagination) Style(value any) pagination {
	return p.set("style", value)
}

// TestIdBuilder sets the custom test ID builder
func (p pagination) TestIdBuilder(value string) pagination {
	return p.set("testIdBuilder", value)
}

// Testid sets the test ID
func (p pagination) Testid(value string) pagination {
	return p.set("testid", value)
}

// Total sets the total number of items
func (p pagination) Total(value int) pagination {
	return p.set("total", value)
}

// UseMobileUI sets whether to use mobile UI
func (p pagination) UseMobileUI(value bool) pagination {
	return p.set("useMobileUI", value)
}

// Visible sets whether the pagination is visible
func (p pagination) Visible(value bool) pagination {
	return p.set("visible", value)
}

// VisibleOn sets the expression to show pagination
func (p pagination) VisibleOn(value string) pagination {
	return p.set("visibleOn", value)
}

// Size sets the size of the pagination
func (p pagination) Size(value string) pagination {
	return p.set("size", value)
}

// EllipsisPageGap sets the page gap for ellipsis
func (p pagination) EllipsisPageGap(value int) pagination {
	return p.set("ellipsisPageGap", value)
}
