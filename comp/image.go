package comp

// image 图片展示控件
type image schema

// Image 创建一个新的 Image 实例
func Image() image {
	return make(image).set("type", "image")
}

func (i image) set(key string, value any) image {
	i[key] = value
	return i
}

// Alt 图片无法显示时的替换文本
func (i image) Alt(value string) image {
	return i.set("alt", value)
}

// Blank 是否新窗口打开
func (i image) Blank(value bool) image {
	return i.set("blank", value)
}

// Caption 图片说明文字
func (i image) Caption(value string) image {
	return i.set("caption", value)
}

// ClassName 外层 css 类名
func (i image) ClassName(value string) image {
	return i.set("className", value)
}

// DefaultImage 默认图片地址
func (i image) DefaultImage(value string) image {
	return i.set("defaultImage", value)
}

// Disabled 是否禁用
func (i image) Disabled(value bool) image {
	return i.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (i image) DisabledOn(value string) image {
	return i.set("disabledOn", value)
}

// EditorSetting 编辑器配置
func (i image) EditorSetting(value string) image {
	return i.set("editorSetting", value)
}

// EnlargeAble 是否启动放大功能
func (i image) EnlargeAble(value bool) image {
	return i.set("enlargeAble", value)
}

// EnlargeWithGallary 放大时是否显示图片集
func (i image) EnlargeWithGallary(value bool) image {
	return i.set("enlargeWithGallary", value)
}

// Height 高度
func (i image) Height(value string) image {
	return i.set("height", value)
}

// Hidden 是否隐藏
func (i image) Hidden(value bool) image {
	return i.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (i image) HiddenOn(value string) image {
	return i.set("hiddenOn", value)
}

// Href 链接地址
func (i image) Href(value string) image {
	return i.set("href", value)
}

// HtmlTarget 链接的 target
func (i image) HtmlTarget(value string) image {
	return i.set("htmlTarget", value)
}

// ID 组件唯一 id
func (i image) ID(value string) image {
	return i.set("id", value)
}

// ImageCaption 图片描述信息
func (i image) ImageCaption(value string) image {
	return i.set("imageCaption", value)
}

// ImageClassName 图片 css 类名
func (i image) ImageClassName(value string) image {
	return i.set("imageClassName", value)
}

// ImageGallaryClassName 放大详情图 CSS 类名
func (i image) ImageGallaryClassName(value string) image {
	return i.set("imageGallaryClassName", value)
}

// ImageMode 图片展示模式
func (i image) ImageMode(value string) image {
	return i.set("imageMode", value)
}

// InnerClassName 组件内层 css 类名
func (i image) InnerClassName(value string) image {
	return i.set("innerClassName", value)
}

// Name 关联字段名
func (i image) Name(value string) image {
	return i.set("name", value)
}

// OnEvent 事件动作配置
func (i image) OnEvent(value any) image {
	return i.set("onEvent", value)
}

// OriginalSrc 大图地址
func (i image) OriginalSrc(value string) image {
	return i.set("originalSrc", value)
}

// ShowToolbar 是否展示图片工具栏
func (i image) ShowToolbar(value bool) image {
	return i.set("showToolbar", value)
}

// Src 图片地址
func (i image) Src(value string) image {
	return i.set("src", value)
}

// Static 是否静态展示
func (i image) Static(value bool) image {
	return i.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (i image) StaticClassName(value string) image {
	return i.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (i image) StaticInputClassName(value string) image {
	return i.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (i image) StaticLabelClassName(value string) image {
	return i.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (i image) StaticOn(value string) image {
	return i.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (i image) StaticPlaceholder(value string) image {
	return i.set("staticPlaceholder", value)
}

// StaticSchema 静态展示模式的 schema
func (i image) StaticSchema(value string) image {
	return i.set("staticSchema", value)
}

// Style 组件样式
func (i image) Style(value any) image {
	return i.set("style", value)
}

// TestIdBuilder
func (i image) TestIdBuilder(value string) image {
	return i.set("testIdBuilder", value)
}

// Testid
func (i image) Testid(value string) image {
	return i.set("testid", value)
}

// ThumbClassName 图片缩略图外层 css 类名
func (i image) ThumbClassName(value string) image {
	return i.set("thumbClassName", value)
}

// ThumbMode 预览图模式
func (i image) ThumbMode(value string) image {
	return i.set("thumbMode", value)
}

// ThumbRatio 预览图比率
func (i image) ThumbRatio(value string) image {
	return i.set("thumbRatio", value)
}

// Title 图片标题
func (i image) Title(value any) image {
	return i.set("title", value)
}

// ToolbarActions 工具栏配置
func (i image) ToolbarActions(value string) image {
	return i.set("toolbarActions", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (i image) UseMobileUI(value bool) image {
	return i.set("useMobileUI", value)
}

// Visible 是否显示
func (i image) Visible(value bool) image {
	return i.set("visible", value)
}

// VisibleOn 是否显示表达式
func (i image) VisibleOn(value string) image {
	return i.set("visibleOn", value)
}

// Width 宽度
func (i image) Width(value string) image {
	return i.set("width", value)
}
