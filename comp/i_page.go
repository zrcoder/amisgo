package comp

import "github.com/zrcoder/amisgo/model"

type MPageItem model.Schema

func PageItem() MPageItem {
	return MPageItem{}
}

func (p MPageItem) set(key string, value any) MPageItem {
	p[key] = value
	return p
}

// Label sets the menu name.
func (p MPageItem) Label(value string) MPageItem {
	return p.set("label", value)
}

// Icon sets the menu icon, e.g., fa fa-file.
func (p MPageItem) Icon(value string) MPageItem {
	return p.set("icon", value)
}

// Url sets the page route path. When the route matches this path, the current page is activated.
// If the path does not start with /, it will be concatenated with the parent path. For example, if the parent path is folder and the current path is pageA,
// the page will be activated when the address is /folder/pageA.
// If the path starts with /, such as /crud/list, the parent path will not be concatenated.
// It also supports routes with parameters like /crud/view/:id, and the value can be accessed in the page using ${params.id}.
func (p MPageItem) Url(value string) MPageItem {
	return p.set("url", value)
}

// model.Schema sets the page configuration.
func (p MPageItem) Schema(value any) MPageItem {
	return p.set("schema", value)
}

// model.SchemaApi sets the API to fetch the schema.
func (p MPageItem) SchemaApi(value string) MPageItem {
	return p.set("schemaApi", value)
}

// Link sets an external link for the menu.
func (p MPageItem) Link(value string) MPageItem {
	return p.set("link", value)
}

// Redirect sets the redirect target when the current page is hit.
func (p MPageItem) Redirect(value string) MPageItem {
	return p.set("redirect", value)
}

// Rewrite sets the path to render another page without changing the URL.
func (p MPageItem) Rewrite(value string) MPageItem {
	return p.set("rewrite", value)
}

// IsDefaultPage sets the page as the custom 404 page.
func (p MPageItem) IsDefaultPage(value bool) MPageItem {
	return p.set("isDefaultPage", value)
}

// Visible sets the visibility of the page in the menu.
func (p MPageItem) Visible(value bool) MPageItem {
	return p.set("visible", value)
}

// ClassName sets the menu class name.
func (p MPageItem) ClassName(value string) MPageItem {
	return p.set("className", value)
}

// Children sets the subpages.
func (p MPageItem) Children(value ...MPageItem) MPageItem {
	return p.set("children", value)
}
