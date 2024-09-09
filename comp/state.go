package comp

// state

type state schema

// State 创建一个新的 State 实例
func State() state {
	return state{}
}

func (s state) set(key string, value any) state {
	s[key] = value
	return s
}

// Body 内容 (内容)
func (s state) Body(value ...any) state {
	return s.set("body", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (s state) ClassName(value string) state {
	return s.set("className", value)
}

// Disabled 是否禁用
func (s state) Disabled(value bool) state {
	return s.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (s state) DisabledOn(value string) state {
	return s.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (s state) EditorSetting(value string) state {
	return s.set("editorSetting", value)
}

// Hidden 是否隐藏
func (s state) Hidden(value bool) state {
	return s.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (s state) HiddenOn(value string) state {
	return s.set("hiddenOn", value)
}

// Id 组件唯一 id，主要用于日志采集
func (s state) Id(value string) state {
	return s.set("id", value)
}

// OnEvent 事件动作配置
func (s state) OnEvent(value any) state {
	return s.set("onEvent", value)
}

// Static 是否静态展示
func (s state) Static(value bool) state {
	return s.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (s state) StaticClassName(value string) state {
	return s.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (s state) StaticInputClassName(value string) state {
	return s.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (s state) StaticLabelClassName(value string) state {
	return s.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (s state) StaticOn(value string) state {
	return s.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (s state) StaticPlaceholder(value string) state {
	return s.set("staticPlaceholder", value)
}

// StaticSchema
func (s state) StaticSchema(value string) state {
	return s.set("staticSchema", value)
}

// Style 组件样式
func (s state) Style(value any) state {
	return s.set("style", value)
}

// TestIdBuilder
func (s state) TestIdBuilder(value string) state {
	return s.set("testIdBuilder", value)
}

// Testid
func (s state) Testid(value string) state {
	return s.set("testid", value)
}

// Title 状态标题
func (s state) Title(value any) state {
	return s.set("title", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (s state) UseMobileUI(value bool) state {
	return s.set("useMobileUI", value)
}

// Visible 是否显示
func (s state) Visible(value bool) state {
	return s.set("visible", value)
}

// VisibleOn 显示条件
func (s state) VisibleOn(value string) state {
	return s.set("visibleOn", value)
}
