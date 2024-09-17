package comp

type breadCrumbItem schema

func BreadcrumbItem() breadCrumbItem {
	return breadCrumbItem{}
}

func (bi breadCrumbItem) set(key string, value any) breadCrumbItem {
	bi[key] = value
	return bi
}

func (bi breadCrumbItem) Label(value string) breadCrumbItem {
	return bi.set("label", value)
}

func (bi breadCrumbItem) Href(value string) breadCrumbItem {
	return bi.set("href", value)
}

func (bi breadCrumbItem) Icon(value string) breadCrumbItem {
	return bi.set("icon", value)
}

func (bi breadCrumbItem) Dropdown(value ...breadCrumbItem) breadCrumbItem {
	return bi.set("dropdown", value)
}
