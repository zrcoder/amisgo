package comp

// vBox 垂直布局控件 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/vbox
//
// @version 6.7.0
type vBox schema

// VBox 创建一个新的 VBox 实例
func VBox() vBox {
	return vBox{}.set("type", "vbox")
}

func (v vBox) set(key string, value interface{}) vBox {
	v[key] = value
	return v
}

// className 容器 css 类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (v vBox) ClassName(value string) vBox {
	return v.set("className", value)
}

// disabled 是否禁用
func (v vBox) Disabled(value bool) vBox {
	return v.set("disabled", value)
}

// disabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (v vBox) DisabledOn(value string) vBox {
	return v.set("disabledOn", value)
}

// editorSetting 编辑器配置，运行时可以忽略
func (v vBox) EditorSetting(value string) vBox {
	return v.set("editorSetting", value)
}

// hidden 是否隐藏
func (v vBox) Hidden(value bool) vBox {
	return v.set("hidden", value)
}

// hiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (v vBox) HiddenOn(value string) vBox {
	return v.set("hiddenOn", value)
}

// id 组件唯一 id，主要用于日志采集
func (v vBox) Id(value string) vBox {
	return v.set("id", value)
}

// onEvent 事件动作配置
func (v vBox) OnEvent(value string) vBox {
	return v.set("onEvent", value)
}

// rows 行集合
func (v vBox) Rows(value string) vBox {
	return v.set("rows", value)
}

// static 是否静态展示
func (v vBox) Static(value bool) vBox {
	return v.set("static", value)
}

// staticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (v vBox) StaticClassName(value string) vBox {
	return v.set("staticClassName", value)
}

// staticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (v vBox) StaticInputClassName(value string) vBox {
	return v.set("staticInputClassName", value)
}

// staticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (v vBox) StaticLabelClassName(value string) vBox {
	return v.set("staticLabelClassName", value)
}

// staticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (v vBox) StaticOn(value string) vBox {
	return v.set("staticOn", value)
}

// staticPlaceholder 静态展示空值占位
func (v vBox) StaticPlaceholder(value string) vBox {
	return v.set("staticPlaceholder", value)
}

// staticSchema
func (v vBox) StaticSchema(value string) vBox {
	return v.set("staticSchema", value)
}

// style 组件样式
func (v vBox) Style(value string) vBox {
	return v.set("style", value)
}

// testIdBuilder
func (v vBox) TestIdBuilder(value string) vBox {
	return v.set("testIdBuilder", value)
}

// testid
func (v vBox) Testid(value string) vBox {
	return v.set("testid", value)
}

// useMobileUI 可以组件级别用来关闭移动端样式
func (v vBox) UseMobileUI(value bool) vBox {
	return v.set("useMobileUI", value)
}

// visible 是否显示
func (v vBox) Visible(value bool) vBox {
	return v.set("visible", value)
}

// visibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (v vBox) VisibleOn(value string) vBox {
	return v.set("visibleOn", value)
}
