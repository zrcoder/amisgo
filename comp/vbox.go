package comp

// VBox 垂直布局控件 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/vbox
//
// @version 6.7.0
type VBox Schema

// NewVBox 创建一个新的 VBox 实例
func NewVBox() VBox {
	return VBox{}.set("type", "vbox")
}

func (v VBox) set(key string, value interface{}) VBox {
	v[key] = value
	return v
}

// className 容器 css 类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (v VBox) ClassName(value string) VBox {
	return v.set("className", value)
}

// disabled 是否禁用
func (v VBox) Disabled(value bool) VBox {
	return v.set("disabled", value)
}

// disabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (v VBox) DisabledOn(value string) VBox {
	return v.set("disabledOn", value)
}

// editorSetting 编辑器配置，运行时可以忽略
func (v VBox) EditorSetting(value string) VBox {
	return v.set("editorSetting", value)
}

// hidden 是否隐藏
func (v VBox) Hidden(value bool) VBox {
	return v.set("hidden", value)
}

// hiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (v VBox) HiddenOn(value string) VBox {
	return v.set("hiddenOn", value)
}

// id 组件唯一 id，主要用于日志采集
func (v VBox) Id(value string) VBox {
	return v.set("id", value)
}

// onEvent 事件动作配置
func (v VBox) OnEvent(value string) VBox {
	return v.set("onEvent", value)
}

// rows 行集合
func (v VBox) Rows(value string) VBox {
	return v.set("rows", value)
}

// static 是否静态展示
func (v VBox) Static(value bool) VBox {
	return v.set("static", value)
}

// staticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (v VBox) StaticClassName(value string) VBox {
	return v.set("staticClassName", value)
}

// staticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (v VBox) StaticInputClassName(value string) VBox {
	return v.set("staticInputClassName", value)
}

// staticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (v VBox) StaticLabelClassName(value string) VBox {
	return v.set("staticLabelClassName", value)
}

// staticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (v VBox) StaticOn(value string) VBox {
	return v.set("staticOn", value)
}

// staticPlaceholder 静态展示空值占位
func (v VBox) StaticPlaceholder(value string) VBox {
	return v.set("staticPlaceholder", value)
}

// staticSchema
func (v VBox) StaticSchema(value string) VBox {
	return v.set("staticSchema", value)
}

// style 组件样式
func (v VBox) Style(value string) VBox {
	return v.set("style", value)
}

// testIdBuilder
func (v VBox) TestIdBuilder(value string) VBox {
	return v.set("testIdBuilder", value)
}

// testid
func (v VBox) Testid(value string) VBox {
	return v.set("testid", value)
}

// type
func (v VBox) Type(value string) VBox {
	return v.set("type", value)
}

// useMobileUI 可以组件级别用来关闭移动端样式
func (v VBox) UseMobileUI(value bool) VBox {
	return v.set("useMobileUI", value)
}

// visible 是否显示
func (v VBox) Visible(value bool) VBox {
	return v.set("visible", value)
}

// visibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (v VBox) VisibleOn(value string) VBox {
	return v.set("visibleOn", value)
}
