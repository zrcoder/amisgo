package comp

// Image 图片展示控件
type Image Schema

// NewImage 创建一个新的 Image 实例
func NewImage() Image {
	return make(Image).set("type", "image")
}

func (i Image) set(key string, value interface{}) Image {
	i[key] = value
	return i
}

// Alt 图片无法显示时的替换文本
func (i Image) Alt(value string) Image {
	return i.set("alt", value)
}

// Blank 是否新窗口打开
func (i Image) Blank(value bool) Image {
	return i.set("blank", value)
}

// Caption 图片说明文字
func (i Image) Caption(value string) Image {
	return i.set("caption", value)
}

// ClassName 外层 css 类名
func (i Image) ClassName(value string) Image {
	return i.set("className", value)
}

// DefaultImage 默认图片地址
func (i Image) DefaultImage(value string) Image {
	return i.set("defaultImage", value)
}

// Disabled 是否禁用
func (i Image) Disabled(value bool) Image {
	return i.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (i Image) DisabledOn(value string) Image {
	return i.set("disabledOn", value)
}

// EditorSetting 编辑器配置
func (i Image) EditorSetting(value string) Image {
	return i.set("editorSetting", value)
}

// EnlargeAble 是否启动放大功能
func (i Image) EnlargeAble(value bool) Image {
	return i.set("enlargeAble", value)
}

// EnlargeWithGallary 放大时是否显示图片集
func (i Image) EnlargeWithGallary(value bool) Image {
	return i.set("enlargeWithGallary", value)
}

// Height 高度
func (i Image) Height(value string) Image {
	return i.set("height", value)
}

// Hidden 是否隐藏
func (i Image) Hidden(value bool) Image {
	return i.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (i Image) HiddenOn(value string) Image {
	return i.set("hiddenOn", value)
}

// Href 链接地址
func (i Image) Href(value string) Image {
	return i.set("href", value)
}

// HtmlTarget 链接的 target
func (i Image) HtmlTarget(value string) Image {
	return i.set("htmlTarget", value)
}

// ID 组件唯一 id
func (i Image) ID(value string) Image {
	return i.set("id", value)
}

// ImageCaption 图片描述信息
func (i Image) ImageCaption(value string) Image {
	return i.set("imageCaption", value)
}

// ImageClassName 图片 css 类名
func (i Image) ImageClassName(value string) Image {
	return i.set("imageClassName", value)
}

// ImageGallaryClassName 放大详情图 CSS 类名
func (i Image) ImageGallaryClassName(value string) Image {
	return i.set("imageGallaryClassName", value)
}

// ImageMode 图片展示模式
func (i Image) ImageMode(value string) Image {
	return i.set("imageMode", value)
}

// InnerClassName 组件内层 css 类名
func (i Image) InnerClassName(value string) Image {
	return i.set("innerClassName", value)
}

// Name 关联字段名
func (i Image) Name(value string) Image {
	return i.set("name", value)
}

// OnEvent 事件动作配置
func (i Image) OnEvent(value string) Image {
	return i.set("onEvent", value)
}

// OriginalSrc 大图地址
func (i Image) OriginalSrc(value string) Image {
	return i.set("originalSrc", value)
}

// ShowToolbar 是否展示图片工具栏
func (i Image) ShowToolbar(value bool) Image {
	return i.set("showToolbar", value)
}

// Src 图片地址
func (i Image) Src(value string) Image {
	return i.set("src", value)
}

// Static 是否静态展示
func (i Image) Static(value bool) Image {
	return i.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (i Image) StaticClassName(value string) Image {
	return i.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (i Image) StaticInputClassName(value string) Image {
	return i.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (i Image) StaticLabelClassName(value string) Image {
	return i.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (i Image) StaticOn(value string) Image {
	return i.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (i Image) StaticPlaceholder(value string) Image {
	return i.set("staticPlaceholder", value)
}

// StaticSchema 静态展示模式的 schema
func (i Image) StaticSchema(value string) Image {
	return i.set("staticSchema", value)
}

// Style 组件样式
func (i Image) Style(value string) Image {
	return i.set("style", value)
}

// TestIdBuilder
func (i Image) TestIdBuilder(value string) Image {
	return i.set("testIdBuilder", value)
}

// Testid
func (i Image) Testid(value string) Image {
	return i.set("testid", value)
}

// ThumbClassName 图片缩略图外层 css 类名
func (i Image) ThumbClassName(value string) Image {
	return i.set("thumbClassName", value)
}

// ThumbMode 预览图模式
func (i Image) ThumbMode(value string) Image {
	return i.set("thumbMode", value)
}

// ThumbRatio 预览图比率
func (i Image) ThumbRatio(value string) Image {
	return i.set("thumbRatio", value)
}

// Title 图片标题
func (i Image) Title(value string) Image {
	return i.set("title", value)
}

// ToolbarActions 工具栏配置
func (i Image) ToolbarActions(value string) Image {
	return i.set("toolbarActions", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (i Image) UseMobileUI(value bool) Image {
	return i.set("useMobileUI", value)
}

// Visible 是否显示
func (i Image) Visible(value bool) Image {
	return i.set("visible", value)
}

// VisibleOn 是否显示表达式
func (i Image) VisibleOn(value string) Image {
	return i.set("visibleOn", value)
}

// Width 宽度
func (i Image) Width(value string) Image {
	return i.set("width", value)
}
