package comp

// images 图片集展示控件。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/images
type images schema

// Images 创建一个新的 Images 实例，并设置默认的 type
func Images() images {
	return make(images).set("type", "images")
}

func (i images) set(key string, value any) images {
	i[key] = value
	return i
}

// ClassName 外层 CSS 类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (i images) ClassName(value string) images {
	return i.set("className", value)
}

// DefaultImage 默认图片地址 (默认图片地址)
func (i images) DefaultImage(value string) images {
	return i.set("defaultImage", value)
}

// Delimiter 配置值的连接符
func (i images) Delimiter(value string) images {
	return i.set("delimiter", value)
}

// Disabled 是否禁用
func (i images) Disabled(value bool) images {
	return i.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (i images) DisabledOn(value string) images {
	return i.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (i images) EditorSetting(value string) images {
	return i.set("editorSetting", value)
}

// EnlargeAble 是否启动放大功能。
func (i images) EnlargeAble(value bool) images {
	return i.set("enlargeAble", value)
}

// EnlargetWithImages 放大时是否显示图片集
func (i images) EnlargetWithImages(value bool) images {
	return i.set("enlargetWithImages", value)
}

// Hidden 是否隐藏
func (i images) Hidden(value bool) images {
	return i.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (i images) HiddenOn(value string) images {
	return i.set("hiddenOn", value)
}

// Id 组件唯一 id，主要用于日志采集
func (i images) ID(value string) images {
	return i.set("id", value)
}

// ImageGallaryClassName 放大详情图 CSS 类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (i images) ImageGallaryClassName(value string) images {
	return i.set("imageGallaryClassName", value)
}

// ListClassName 列表 CSS 类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (i images) ListClassName(value string) images {
	return i.set("listClassName", value)
}

// Name 关联字段名，也可以直接配置 src
func (i images) Name(value string) images {
	return i.set("name", value)
}

// OnEvent 事件动作配置
func (i images) OnEvent(value any) images {
	return i.set("onEvent", value)
}

// Options
func (i images) Options(value ...option) images {
	return i.set("options", value)
}

// OriginalSrc 大图地址，不设置用 src 属性，如果不是请配置，如：${imageOriginUrl}
func (i images) OriginalSrc(value string) images {
	return i.set("originalSrc", value)
}

// Placeholder 列表为空时显示
func (i images) Placeholder(value string) images {
	return i.set("placeholder", value)
}

// ShowDimensions 是否显示尺寸。
func (i images) ShowDimensions(value bool) images {
	return i.set("showDimensions", value)
}

// ShowToolbar 是否展示图片工具栏
func (i images) ShowToolbar(value bool) images {
	return i.set("showToolbar", value)
}

// Source
func (i images) Source(value string) images {
	return i.set("source", value)
}

// Src 图片地址，默认读取数据中的 image 属性，如果不是请配置 ,如  ${imageUrl}
func (i images) Src(value string) images {
	return i.set("src", value)
}

// Static 是否静态展示
func (i images) Static(value bool) images {
	return i.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (i images) StaticClassName(value string) images {
	return i.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (i images) StaticInputClassName(value string) images {
	return i.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (i images) StaticLabelClassName(value string) images {
	return i.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (i images) StaticOn(value string) images {
	return i.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (i images) StaticPlaceholder(value string) images {
	return i.set("staticPlaceholder", value)
}

// StaticSchema
func (i images) StaticSchema(value string) images {
	return i.set("staticSchema", value)
}

// Style 组件样式
func (i images) Style(value any) images {
	return i.set("style", value)
}

// TestIdBuilder
func (i images) TestIdBuilder(value string) images {
	return i.set("testIdBuilder", value)
}

// Testid
func (i images) Testid(value string) images {
	return i.set("testid", value)
}

// ThumbMode 预览图模式 可选值: w-full | h-full | contain | cover
func (i images) ThumbMode(value string) images {
	return i.set("thumbMode", value)
}

// ThumbRatio 预览图比率 可选值: 1:1 | 4:3 | 16:9
func (i images) ThumbRatio(value string) images {
	return i.set("thumbRatio", value)
}

// ToolbarActions 工具栏配置
func (i images) ToolbarActions(value string) images {
	return i.set("toolbarActions", value)
}

// Type 指定为图片集渲染器 可选值: images | static-images
func (i images) Type(value string) images {
	return i.set("type", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (i images) UseMobileUI(value bool) images {
	return i.set("useMobileUI", value)
}

// Value
func (i images) Value(value string) images {
	return i.set("value", value)
}

// Visible 是否显示
func (i images) Visible(value bool) images {
	return i.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (i images) VisibleOn(value string) images {
	return i.set("visibleOn", value)
}
