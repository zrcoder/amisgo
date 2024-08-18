package comp

// Steps 文档
//
// @version 6.7.0
type Steps Schema

// NewSteps 创建一个新的 Steps 实例
func NewSteps() Steps {
	return Steps{}.set("type", "steps")
}

func (s Steps) set(key string, value interface{}) Steps {
	s[key] = value
	return s
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (s Steps) ClassName(value string) Steps {
	return s.set("className", value)
}

// Disabled 是否禁用
func (s Steps) Disabled(value bool) Steps {
	return s.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (s Steps) DisabledOn(value string) Steps {
	return s.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (s Steps) EditorSetting(value string) Steps {
	return s.set("editorSetting", value)
}

// Hidden 是否隐藏
func (s Steps) Hidden(value bool) Steps {
	return s.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (s Steps) HiddenOn(value string) Steps {
	return s.set("hiddenOn", value)
}

// Id 组件唯一 id，主要用于日志采集
func (s Steps) Id(value string) Steps {
	return s.set("id", value)
}

// LabelPlacement 标签放置位置 可选值: horizontal | vertical
func (s Steps) LabelPlacement(value string) Steps {
	return s.set("labelPlacement", value)
}

// Mode 展示模式 可选值: horizontal | vertical
func (s Steps) Mode(value string) Steps {
	return s.set("mode", value)
}

// Name 变量映射
func (s Steps) Name(value string) Steps {
	return s.set("name", value)
}

// OnEvent 事件动作配置
func (s Steps) OnEvent(value string) Steps {
	return s.set("onEvent", value)
}

// ProgressDot 点状步骤条
func (s Steps) ProgressDot(value bool) Steps {
	return s.set("progressDot", value)
}

// Source API 或 数据映射
func (s Steps) Source(value string) Steps {
	return s.set("source", value)
}

// Static 是否静态展示
func (s Steps) Static(value bool) Steps {
	return s.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (s Steps) StaticClassName(value string) Steps {
	return s.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (s Steps) StaticInputClassName(value string) Steps {
	return s.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (s Steps) StaticLabelClassName(value string) Steps {
	return s.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (s Steps) StaticOn(value string) Steps {
	return s.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (s Steps) StaticPlaceholder(value string) Steps {
	return s.set("staticPlaceholder", value)
}

// StaticSchema
func (s Steps) StaticSchema(value string) Steps {
	return s.set("staticSchema", value)
}

// Status
func (s Steps) Status(value string) Steps {
	return s.set("status", value)
}

// Steps 步骤
func (s Steps) Steps(value string) Steps {
	return s.set("steps", value)
}

// Style 组件样式
func (s Steps) Style(value string) Steps {
	return s.set("style", value)
}

// TestIdBuilder
func (s Steps) TestIdBuilder(value string) Steps {
	return s.set("testIdBuilder", value)
}

// Testid
func (s Steps) Testid(value string) Steps {
	return s.set("testid", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (s Steps) UseMobileUI(value bool) Steps {
	return s.set("useMobileUI", value)
}

// Value 指定当前步骤
func (s Steps) Value(value string) Steps {
	return s.set("value", value)
}

// Visible 是否显示
func (s Steps) Visible(value bool) Steps {
	return s.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (s Steps) VisibleOn(value string) Steps {
	return s.set("visibleOn", value)
}
