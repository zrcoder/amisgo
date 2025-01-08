package comp

import "github.com/zrcoder/amisgo/model"

type MBreadcrumbItem model.Schema

func BreadcrumbItem() MBreadcrumbItem {
	return MBreadcrumbItem{}
}

func (bi MBreadcrumbItem) set(key string, value any) MBreadcrumbItem {
	bi[key] = value
	return bi
}

func (bi MBreadcrumbItem) Label(value string) MBreadcrumbItem {
	return bi.set("label", value)
}

func (bi MBreadcrumbItem) Href(value string) MBreadcrumbItem {
	return bi.set("href", value)
}

func (bi MBreadcrumbItem) Icon(value string) MBreadcrumbItem {
	return bi.set("icon", value)
}

func (bi MBreadcrumbItem) Dropdown(value ...MBreadcrumbItem) MBreadcrumbItem {
	return bi.set("dropdown", value)
}
