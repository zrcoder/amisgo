package comp

// tpl 渲染器
type tpl schema

// Tpl 创建一个新的 Tpl 实例
func Tpl() tpl {
	return tpl{}.set("type", "tpl")
}

func (t tpl) set(key string, value interface{}) tpl {
	t[key] = value
	return t
}

// Badge 角标 (Badge 角标。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/badge)
func (t tpl) Badge(value string) tpl {
	return t.set("badge", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (t tpl) ClassName(value string) tpl {
	return t.set("className", value)
}

// Disabled 是否禁用
func (t tpl) Disabled(value bool) tpl {
	return t.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (t tpl) DisabledOn(value string) tpl {
	return t.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (t tpl) EditorSetting(value string) tpl {
	return t.set("editorSetting", value)
}

// Hidden 是否隐藏
func (t tpl) Hidden(value bool) tpl {
	return t.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (t tpl) HiddenOn(value string) tpl {
	return t.set("hiddenOn", value)
}

// Html 支持两种语法，但是不能混着用。分别是：1. `${xxx}` 或者 `${xxx|upperCase}` 2. `<%= data.xxx %>` 更多文档：https://aisuda.bce.baidu.com/amis/zh-CN/docs/concepts/template
func (t tpl) Html(value string) tpl {
	return t.set("html", value)
}

// Id 组件唯一 id，主要用于日志采集
func (t tpl) Id(value string) tpl {
	return t.set("id", value)
}

// Inline 是否内联显示？
func (t tpl) Inline(value bool) tpl {
	return t.set("inline", value)
}

// OnEvent 事件动作配置
func (t tpl) OnEvent(value string) tpl {
	return t.set("onEvent", value)
}

// Raw
func (t tpl) Raw(value string) tpl {
	return t.set("raw", value)
}

// Static 是否静态展示
func (t tpl) Static(value bool) tpl {
	return t.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (t tpl) StaticClassName(value string) tpl {
	return t.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (t tpl) StaticInputClassName(value string) tpl {
	return t.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (t tpl) StaticLabelClassName(value string) tpl {
	return t.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (t tpl) StaticOn(value string) tpl {
	return t.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (t tpl) StaticPlaceholder(value string) tpl {
	return t.set("staticPlaceholder", value)
}

// StaticSchema
func (t tpl) StaticSchema(value string) tpl {
	return t.set("staticSchema", value)
}

// Style 自定义样式
func (t tpl) Style(value string) tpl {
	return t.set("style", value)
}

// TestIdBuilder
func (t tpl) TestIdBuilder(value string) tpl {
	return t.set("testIdBuilder", value)
}

// Testid
func (t tpl) Testid(value string) tpl {
	return t.set("testid", value)
}

// Text 支持两种语法，但是不能混着用。分别是：1. `${xxx}` 或者 `${xxx|upperCase}` 2. `<%= data.xxx %>` 更多文档：https://aisuda.bce.baidu.com/amis/zh-CN/docs/concepts/template
func (t tpl) Text(value string) tpl {
	return t.set("text", value)
}

// Tpl 支持两种语法，但是不能混着用。分别是：1. `${xxx}` 或者 `${xxx|upperCase}` 2. `<%= data.xxx %>` 更多文档：https://aisuda.bce.baidu.com/amis/zh-CN/docs/concepts/template
func (t tpl) Tpl(value string) tpl {
	return t.set("tpl", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (t tpl) UseMobileUI(value bool) tpl {
	return t.set("useMobileUI", value)
}

// Visible 是否显示
func (t tpl) Visible(value bool) tpl {
	return t.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (t tpl) VisibleOn(value string) tpl {
	return t.set("visibleOn", value)
}

// WrapperComponent 标签类型
func (t tpl) WrapperComponent(value string) tpl {
	return t.set("wrapperComponent", value)
}
