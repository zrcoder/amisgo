package comp

type breadcrumbItem schema

func BreadcrumbItem() breadcrumbItem {
	return breadcrumbItem{}
}

func (bi breadcrumbItem) set(key string, value any) breadcrumbItem {
	bi[key] = value
	return bi
}

func (bi breadcrumbItem) Label(value string) breadcrumbItem {
	return bi.set("label", value)
}

func (bi breadcrumbItem) Href(value string) breadcrumbItem {
	return bi.set("href", value)
}

func (bi breadcrumbItem) Icon(value string) breadcrumbItem {
	return bi.set("icon", value)
}

func (bi breadcrumbItem) Dropdown(value ...breadcrumbItem) breadcrumbItem {
	return bi.set("dropdown", value)
}
