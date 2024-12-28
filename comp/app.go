package comp

type app Schema

// App 用于实现多页应用，适合于全屏模式，如果只是局部渲染请不要使用。
func App() app {
	return app{}.set("type", "app")
}

func (a app) set(key string, value any) app {
	a[key] = value
	return a
}

// Api 页面配置接口，如果你想远程拉取页面配置请配置。返回配置路径 json>data>pages，具体格式请参考 pages 属性。
func (a app) Api(value any) app {
	return a.set("api", value)
}

// BrandName 应用名称
func (a app) BrandName(value string) app {
	return a.set("brandName", value)
}

// Logo 应用logo， 支持图片地址，或者 svg内容
func (a app) Logo(value string) app {
	return a.set("logo", value)
}

// ClassName css 类名。
func (a app) ClassName(value string) app {
	return a.set("className", value)
}

// Header 顶部区域
func (a app) Header(value ...any) app {
	return a.set("header", value)
}

// AsideBefore 页面菜单上前面区域。
func (a app) AsideBefore(value ...any) app {
	return a.set("asideBefore", value)
}

// AsideAfter 页面菜单下前面区域。
func (a app) AsideAfter(value ...any) app {
	return a.set("asideAfter", value)
}

// Footer 页脚
func (a app) Footer(value ...any) app {
	return a.set("footer", value)
}

// Pages 具体的页面配置。 通常为数组，数组第一层为分组，
// 一般只需要配置 label 集合，如果你不想分组，直接不配置，
// 真正的页面请在第二层开始配置，即第一层的 children 中。
func (a app) Pages(value ...MPageItem) app {
	return a.set("pages", value)
}
