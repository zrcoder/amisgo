package comp

// Images 图片集展示控件。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/images
type Images Schema

// NewImages 创建一个新的 Images 实例，并设置默认的 type
func NewImages() Images {
	return make(Images).set("type", "images")
}

func (i Images) set(key string, value interface{}) Images {
	i[key] = value
	return i
}

// ClassName 外层 CSS 类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (i Images) ClassName(value string) Images {
	return i.set("className", value)
}

// DefaultImage 默认图片地址 (默认图片地址)
func (i Images) DefaultImage(value string) Images {
	return i.set("defaultImage", value)
}

// Delimiter 配置值的连接符
func (i Images) Delimiter(value string) Images {
	return i.set("delimiter", value)
}

// Disabled 是否禁用
func (i Images) Disabled(value bool) Images {
	return i.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (i Images) DisabledOn(value string) Images {
	return i.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (i Images) EditorSetting(value string) Images {
	return i.set("editorSetting", value)
}

// EnlargeAble 是否启动放大功能。
func (i Images) EnlargeAble(value bool) Images {
	return i.set("enlargeAble", value)
}

// EnlargetWithImages 放大时是否显示图片集
func (i Images) EnlargetWithImages(value bool) Images {
	return i.set("enlargetWithImages", value)
}

// Hidden 是否隐藏
func (i Images) Hidden(value bool) Images {
	return i.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (i Images) HiddenOn(value string) Images {
	return i.set("hiddenOn", value)
}

// Id 组件唯一 id，主要用于日志采集
func (i Images) Id(value string) Images {
	return i.set("id", value)
}

// ImageGallaryClassName 放大详情图 CSS 类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (i Images) ImageGallaryClassName(value string) Images {
	return i.set("imageGallaryClassName", value)
}

// ListClassName 列表 CSS 类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (i Images) ListClassName(value string) Images {
	return i.set("listClassName", value)
}

// Name 关联字段名，也可以直接配置 src
func (i Images) Name(value string) Images {
	return i.set("name", value)
}

// OnEvent 事件动作配置
func (i Images) OnEvent(value string) Images {
	return i.set("onEvent", value)
}

// Options
func (i Images) Options(value string) Images {
	return i.set("options", value)
}

// OriginalSrc 大图地址，不设置用 src 属性，如果不是请配置，如：${imageOriginUrl}
func (i Images) OriginalSrc(value string) Images {
	return i.set("originalSrc", value)
}

// Placeholder 列表为空时显示
func (i Images) Placeholder(value string) Images {
	return i.set("placeholder", value)
}

// ShowDimensions 是否显示尺寸。
func (i Images) ShowDimensions(value bool) Images {
	return i.set("showDimensions", value)
}

// ShowToolbar 是否展示图片工具栏
func (i Images) ShowToolbar(value bool) Images {
	return i.set("showToolbar", value)
}

// Source
func (i Images) Source(value string) Images {
	return i.set("source", value)
}

// Src 图片地址，默认读取数据中的 image 属性，如果不是请配置 ,如  ${imageUrl}
func (i Images) Src(value string) Images {
	return i.set("src", value)
}

// Static 是否静态展示
func (i Images) Static(value bool) Images {
	return i.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (i Images) StaticClassName(value string) Images {
	return i.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (i Images) StaticInputClassName(value string) Images {
	return i.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (i Images) StaticLabelClassName(value string) Images {
	return i.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (i Images) StaticOn(value string) Images {
	return i.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (i Images) StaticPlaceholder(value string) Images {
	return i.set("staticPlaceholder", value)
}

// StaticSchema
func (i Images) StaticSchema(value string) Images {
	return i.set("staticSchema", value)
}

// Style 组件样式
func (i Images) Style(value string) Images {
	return i.set("style", value)
}

// TestIdBuilder
func (i Images) TestIdBuilder(value string) Images {
	return i.set("testIdBuilder", value)
}

// Testid
func (i Images) Testid(value string) Images {
	return i.set("testid", value)
}

// ThumbMode 预览图模式 可选值: w-full | h-full | contain | cover
func (i Images) ThumbMode(value string) Images {
	return i.set("thumbMode", value)
}

// ThumbRatio 预览图比率 可选值: 1:1 | 4:3 | 16:9
func (i Images) ThumbRatio(value string) Images {
	return i.set("thumbRatio", value)
}

// ToolbarActions 工具栏配置
func (i Images) ToolbarActions(value string) Images {
	return i.set("toolbarActions", value)
}

// Type 指定为图片集渲染器 可选值: images | static-images
func (i Images) Type(value string) Images {
	return i.set("type", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (i Images) UseMobileUI(value bool) Images {
	return i.set("useMobileUI", value)
}

// Value
func (i Images) Value(value string) Images {
	return i.set("value", value)
}

// Visible 是否显示
func (i Images) Visible(value bool) Images {
	return i.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (i Images) VisibleOn(value string) Images {
	return i.set("visibleOn", value)
}
