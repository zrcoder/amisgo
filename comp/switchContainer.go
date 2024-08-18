package comp

// switchContainer 状态容器渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/state-container
//
// @version 6.7.0
type switchContainer schema

// SwitchContainer 创建一个新的 SwitchContainer 实例
func SwitchContainer() switchContainer {
	return switchContainer{}.set("type", "switch-container")
}

func (s switchContainer) set(key string, value interface{}) switchContainer {
	s[key] = value
	return s
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (s switchContainer) ClassName(value string) switchContainer {
	return s.set("className", value)
}

// Disabled 是否禁用
func (s switchContainer) Disabled(value bool) switchContainer {
	return s.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (s switchContainer) DisabledOn(value string) switchContainer {
	return s.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (s switchContainer) EditorSetting(value string) switchContainer {
	return s.set("editorSetting", value)
}

// Hidden 是否隐藏
func (s switchContainer) Hidden(value bool) switchContainer {
	return s.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (s switchContainer) HiddenOn(value string) switchContainer {
	return s.set("hiddenOn", value)
}

// Id 组件唯一 id，主要用于日志采集
func (s switchContainer) Id(value string) switchContainer {
	return s.set("id", value)
}

// Items 状态项列表
func (s switchContainer) Items(value string) switchContainer {
	return s.set("items", value)
}

// OnEvent 事件动作配置
func (s switchContainer) OnEvent(value string) switchContainer {
	return s.set("onEvent", value)
}

// Static 是否静态展示
func (s switchContainer) Static(value bool) switchContainer {
	return s.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (s switchContainer) StaticClassName(value string) switchContainer {
	return s.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (s switchContainer) StaticInputClassName(value string) switchContainer {
	return s.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (s switchContainer) StaticLabelClassName(value string) switchContainer {
	return s.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (s switchContainer) StaticOn(value string) switchContainer {
	return s.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (s switchContainer) StaticPlaceholder(value string) switchContainer {
	return s.set("staticPlaceholder", value)
}

// StaticSchema
func (s switchContainer) StaticSchema(value string) switchContainer {
	return s.set("staticSchema", value)
}

// Style 自定义样式
func (s switchContainer) Style(value string) switchContainer {
	return s.set("style", value)
}

// TestIdBuilder
func (s switchContainer) TestIdBuilder(value string) switchContainer {
	return s.set("testIdBuilder", value)
}

// Testid
func (s switchContainer) Testid(value string) switchContainer {
	return s.set("testid", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (s switchContainer) UseMobileUI(value bool) switchContainer {
	return s.set("useMobileUI", value)
}

// Visible 是否显示
func (s switchContainer) Visible(value bool) switchContainer {
	return s.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (s switchContainer) VisibleOn(value string) switchContainer {
	return s.set("visibleOn", value)
}
