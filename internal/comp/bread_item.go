package comp

import "github.com/zrcoder/amisgo/schema"

type BreadcrumbItem schema.Schema

func NewBreadcrumbItem() BreadcrumbItem {
	return BreadcrumbItem{}
}

func (bi BreadcrumbItem) set(key string, value any) BreadcrumbItem {
	bi[key] = value
	return bi
}

func (bi BreadcrumbItem) Label(value string) BreadcrumbItem {
	return bi.set("label", value)
}

func (bi BreadcrumbItem) Href(value string) BreadcrumbItem {
	return bi.set("href", value)
}

func (bi BreadcrumbItem) Icon(value string) BreadcrumbItem {
	return bi.set("icon", value)
}

func (bi BreadcrumbItem) Dropdown(value ...BreadcrumbItem) BreadcrumbItem {
	return bi.set("dropdown", value)
}
