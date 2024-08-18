package comp

// State
//
// @version 6.7.0
type State Schema

// NewState 创建一个新的 State 实例
func NewState() State {
	return State{}
}

func (s State) set(key string, value interface{}) State {
	s[key] = value
	return s
}

// Body 内容 (内容)
func (s State) Body(value ...interface{}) State {
	return s.set("body", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (s State) ClassName(value string) State {
	return s.set("className", value)
}

// Disabled 是否禁用
func (s State) Disabled(value bool) State {
	return s.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (s State) DisabledOn(value string) State {
	return s.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (s State) EditorSetting(value string) State {
	return s.set("editorSetting", value)
}

// Hidden 是否隐藏
func (s State) Hidden(value bool) State {
	return s.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (s State) HiddenOn(value string) State {
	return s.set("hiddenOn", value)
}

// Id 组件唯一 id，主要用于日志采集
func (s State) Id(value string) State {
	return s.set("id", value)
}

// OnEvent 事件动作配置
func (s State) OnEvent(value string) State {
	return s.set("onEvent", value)
}

// Static 是否静态展示
func (s State) Static(value bool) State {
	return s.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (s State) StaticClassName(value string) State {
	return s.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (s State) StaticInputClassName(value string) State {
	return s.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (s State) StaticLabelClassName(value string) State {
	return s.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (s State) StaticOn(value string) State {
	return s.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (s State) StaticPlaceholder(value string) State {
	return s.set("staticPlaceholder", value)
}

// StaticSchema
func (s State) StaticSchema(value string) State {
	return s.set("staticSchema", value)
}

// Style 组件样式
func (s State) Style(value string) State {
	return s.set("style", value)
}

// TestIdBuilder
func (s State) TestIdBuilder(value string) State {
	return s.set("testIdBuilder", value)
}

// Testid
func (s State) Testid(value string) State {
	return s.set("testid", value)
}

// Title 状态标题
func (s State) Title(value string) State {
	return s.set("title", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (s State) UseMobileUI(value bool) State {
	return s.set("useMobileUI", value)
}

// Visible 是否显示
func (s State) Visible(value bool) State {
	return s.set("visible", value)
}

// VisibleOn 显示条件
func (s State) VisibleOn(value string) State {
	return s.set("visibleOn", value)
}
