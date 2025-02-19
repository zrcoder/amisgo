package comp

import "github.com/zrcoder/amisgo/schema"

type PageItem schema.Schema

func NewPageItem() PageItem {
	return PageItem{}
}

func (p PageItem) set(key string, value any) PageItem {
	p[key] = value
	return p
}

// Label sets the menu name.
func (p PageItem) Label(value string) PageItem {
	return p.set("label", value)
}

// Icon sets the menu icon, e.g., fa fa-file.
func (p PageItem) Icon(value string) PageItem {
	return p.set("icon", value)
}

// Url sets the page route path. When the route matches this path, the current page is activated.
// If the path does not start with /, it will be concatenated with the parent path. For example, if the parent path is folder and the current path is pageA,
// the page will be activated when the address is /folder/pageA.
// If the path starts with /, such as /crud/list, the parent path will not be concatenated.
// It also supports routes with parameters like /crud/view/:id, and the value can be accessed in the page using ${params.id}.
func (p PageItem) Url(value string) PageItem {
	return p.set("url", value)
}

// Schema sets the page configuration.
func (p PageItem) Schema(value any) PageItem {
	return p.set("schema", value)
}

// SchemaApi sets the API to fetch the schema.Schema.
func (p PageItem) SchemaApi(value string) PageItem {
	return p.set("schemaApi", value)
}

// Link sets an external link for the menu.
func (p PageItem) Link(value string) PageItem {
	return p.set("link", value)
}

// Redirect sets the redirect target when the current page is hit.
func (p PageItem) Redirect(value string) PageItem {
	return p.set("redirect", value)
}

// Rewrite sets the path to render another page without changing the URL.
func (p PageItem) Rewrite(value string) PageItem {
	return p.set("rewrite", value)
}

// IsDefaultPage sets the page as the custom 404 page.
func (p PageItem) IsDefaultPage(value bool) PageItem {
	return p.set("isDefaultPage", value)
}

// Visible sets the visibility of the page in the menu.
func (p PageItem) Visible(value bool) PageItem {
	return p.set("visible", value)
}

// ClassName sets the menu class name.
func (p PageItem) ClassName(value string) PageItem {
	return p.set("className", value)
}

// Children sets the subpages.
func (p PageItem) Children(value ...PageItem) PageItem {
	return p.set("children", value)
}
