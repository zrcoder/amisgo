package comp

// wrapper 表示一个容器渲染器
// 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/wrapper

type wrapper schema

// Wrapper 创建一个新的 Wrapper 实例
func Wrapper() wrapper {
	return wrapper{}.set("type", "wrapper")
}

func (w wrapper) set(key string, value any) wrapper {
	w[key] = value
	return w
}

// Body 内容 (内容)
func (w wrapper) Body(value ...any) wrapper {
	return w.set("body", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (w wrapper) ClassName(value string) wrapper {
	return w.set("className", value)
}

// Disabled 是否禁用
func (w wrapper) Disabled(value bool) wrapper {
	return w.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (w wrapper) DisabledOn(value string) wrapper {
	return w.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (w wrapper) EditorSetting(value string) wrapper {
	return w.set("editorSetting", value)
}

// Hidden 是否隐藏
func (w wrapper) Hidden(value bool) wrapper {
	return w.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (w wrapper) HiddenOn(value string) wrapper {
	return w.set("hiddenOn", value)
}

// Id 组件唯一 id，主要用于日志采集
func (w wrapper) Id(value string) wrapper {
	return w.set("id", value)
}

// OnEvent 事件动作配置
func (w wrapper) OnEvent(value any) wrapper {
	return w.set("onEvent", value)
}

// Size 可选值: xs | sm | md | lg | none
func (w wrapper) Size(value string) wrapper {
	return w.set("size", value)
}

// Static 是否静态展示
func (w wrapper) Static(value bool) wrapper {
	return w.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (w wrapper) StaticClassName(value string) wrapper {
	return w.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (w wrapper) StaticInputClassName(value string) wrapper {
	return w.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (w wrapper) StaticLabelClassName(value string) wrapper {
	return w.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (w wrapper) StaticOn(value string) wrapper {
	return w.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (w wrapper) StaticPlaceholder(value string) wrapper {
	return w.set("staticPlaceholder", value)
}

// StaticSchema
func (w wrapper) StaticSchema(value string) wrapper {
	return w.set("staticSchema", value)
}

// Style 自定义样式
func (w wrapper) Style(value any) wrapper {
	return w.set("style", value)
}

// TestIdBuilder
func (w wrapper) TestIdBuilder(value string) wrapper {
	return w.set("testIdBuilder", value)
}

// Testid
func (w wrapper) Testid(value string) wrapper {
	return w.set("testid", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (w wrapper) UseMobileUI(value bool) wrapper {
	return w.set("useMobileUI", value)
}

// Visible 是否显示
func (w wrapper) Visible(value bool) wrapper {
	return w.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (w wrapper) VisibleOn(value string) wrapper {
	return w.set("visibleOn", value)
}

// Wrap
func (w wrapper) Wrap(value bool) wrapper {
	return w.set("wrap", value)
}
