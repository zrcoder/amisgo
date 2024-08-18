package comp

// Wrapper 表示一个容器渲染器
//
// 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/wrapper
// @author  slowlyo
// @version 6.7.0
type Wrapper Schema

// NewWrapper 创建一个新的 Wrapper 实例
func NewWrapper() Wrapper {
	return Wrapper{}.set("type", "wrapper")
}

func (w Wrapper) set(key string, value interface{}) Wrapper {
	w[key] = value
	return w
}

// Body 内容 (内容)
func (w Wrapper) Body(value ...interface{}) Wrapper {
	return w.set("body", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (w Wrapper) ClassName(value string) Wrapper {
	return w.set("className", value)
}

// Disabled 是否禁用
func (w Wrapper) Disabled(value bool) Wrapper {
	return w.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (w Wrapper) DisabledOn(value string) Wrapper {
	return w.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (w Wrapper) EditorSetting(value string) Wrapper {
	return w.set("editorSetting", value)
}

// Hidden 是否隐藏
func (w Wrapper) Hidden(value bool) Wrapper {
	return w.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (w Wrapper) HiddenOn(value string) Wrapper {
	return w.set("hiddenOn", value)
}

// Id 组件唯一 id，主要用于日志采集
func (w Wrapper) Id(value string) Wrapper {
	return w.set("id", value)
}

// OnEvent 事件动作配置
func (w Wrapper) OnEvent(value string) Wrapper {
	return w.set("onEvent", value)
}

// Size 可选值: xs | sm | md | lg | none
func (w Wrapper) Size(value string) Wrapper {
	return w.set("size", value)
}

// Static 是否静态展示
func (w Wrapper) Static(value bool) Wrapper {
	return w.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (w Wrapper) StaticClassName(value string) Wrapper {
	return w.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (w Wrapper) StaticInputClassName(value string) Wrapper {
	return w.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (w Wrapper) StaticLabelClassName(value string) Wrapper {
	return w.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (w Wrapper) StaticOn(value string) Wrapper {
	return w.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (w Wrapper) StaticPlaceholder(value string) Wrapper {
	return w.set("staticPlaceholder", value)
}

// StaticSchema
func (w Wrapper) StaticSchema(value string) Wrapper {
	return w.set("staticSchema", value)
}

// Style 自定义样式
func (w Wrapper) Style(value string) Wrapper {
	return w.set("style", value)
}

// TestIdBuilder
func (w Wrapper) TestIdBuilder(value string) Wrapper {
	return w.set("testIdBuilder", value)
}

// Testid
func (w Wrapper) Testid(value string) Wrapper {
	return w.set("testid", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (w Wrapper) UseMobileUI(value bool) Wrapper {
	return w.set("useMobileUI", value)
}

// Visible 是否显示
func (w Wrapper) Visible(value bool) Wrapper {
	return w.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (w Wrapper) VisibleOn(value string) Wrapper {
	return w.set("visibleOn", value)
}

// Wrap
func (w Wrapper) Wrap(value bool) Wrapper {
	return w.set("wrap", value)
}
