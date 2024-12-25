package comp

type MPageItem Schema

func PageItem() MPageItem {
	return MPageItem{}
}

func (p MPageItem) set(key string, value any) MPageItem {
	p[key] = value
	return p
}

// Label 菜单名称。
func (p MPageItem) Label(value string) MPageItem {
	return p.set("label", value)
}

// Icon 菜单图标，比如：fa fa-file.
func (p MPageItem) Icon(value string) MPageItem {
	return p.set("icon", value)
}

// url 页面路由路径，当路由命中该路径时，启用当前页面。
// 当路径不是 / 打头时，会连接父级路径。比如：父级的路径为 folder，而此时配置 pageA, 那么当页面地址为 /folder/pageA 时才会命中此页面。
// 当路径是 / 开头如： /crud/list 时，则不会拼接父级路径。
// 另外还支持 /crud/view/:id 这类带参数的路由，页面中可以通过 ${params.id} 取到此值。
func (p MPageItem) Url(value string) MPageItem {
	return p.set("url", value)
}

// Schema 页面的配置，具体配置为 comp.Page
func (p MPageItem) Schema(value any) MPageItem {
	return p.set("schema", value)
}

// SchemaApi 如果想通过接口拉取，请配置。返回路径为 json>data。schema 和 SchemaApi 只能二选一。
func (p MPageItem) SchemaApi(value string) MPageItem {
	return p.set("schemaApi", value)
}

// Link 如果想配置个外部链接菜单，只需要配置 link 即可。
func (p MPageItem) Link(value string) MPageItem {
	return p.set("link", value)
}

// Redirect 跳转，当命中当前页面时，跳转到目标页面。
func (p MPageItem) Redirect(value string) MPageItem {
	return p.set("redirect", value)
}

// Rewrite 改成渲染其他路径的页面，这个方式页面地址不会发生修改。
func (p MPageItem) Rewrite(value string) MPageItem {
	return p.set("rewrite", value)
}

// IsDefaultPage 当你需要自定义 404 页面的时候有用，不要出现多个这样的页面，因为只有第一个才会有用。
func (p MPageItem) IsDefaultPage(value bool) MPageItem {
	return p.set("isDefaultPage", value)
}

// Visible 有些页面可能不想出现在菜单中，可以配置成 false，另外带参数的路由无需配置，直接就是不可见的。
func (p MPageItem) Visible(value bool) MPageItem {
	return p.set("visible", value)
}

// ClassName 菜单类名。
func (p MPageItem) ClassName(value string) MPageItem {
	return p.set("className", value)
}

// Children 子页面
func (p MPageItem) Children(value ...MPageItem) MPageItem {
	return p.set("children", value)
}
