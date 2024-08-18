package comp

// sparkLine
//
// @version 6.7.0
type sparkLine schema

// SparkLine 创建一个新的 SparkLine 实例
func SparkLine() sparkLine {
	return sparkLine{}.set("type", "sparkline")
}

func (s sparkLine) set(key string, value interface{}) sparkLine {
	s[key] = value
	return s
}

// ClassName css 类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (s sparkLine) ClassName(value string) sparkLine {
	return s.set("className", value)
}

// ClickAction 点击行为 (点击行为)
func (s sparkLine) ClickAction(value string) sparkLine {
	return s.set("clickAction", value)
}

// Disabled 是否禁用
func (s sparkLine) Disabled(value bool) sparkLine {
	return s.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (s sparkLine) DisabledOn(value string) sparkLine {
	return s.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (s sparkLine) EditorSetting(value string) sparkLine {
	return s.set("editorSetting", value)
}

// Height 高度
func (s sparkLine) Height(value string) sparkLine {
	return s.set("height", value)
}

// Hidden 是否隐藏
func (s sparkLine) Hidden(value bool) sparkLine {
	return s.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (s sparkLine) HiddenOn(value string) sparkLine {
	return s.set("hiddenOn", value)
}

// Id 组件唯一 id，主要用于日志采集
func (s sparkLine) Id(value string) sparkLine {
	return s.set("id", value)
}

// Name 关联数据变量。
func (s sparkLine) Name(value string) sparkLine {
	return s.set("name", value)
}

// OnEvent 事件动作配置
func (s sparkLine) OnEvent(value string) sparkLine {
	return s.set("onEvent", value)
}

// Placeholder 空数据时显示的内容
func (s sparkLine) Placeholder(value string) sparkLine {
	return s.set("placeholder", value)
}

// Static 是否静态展示
func (s sparkLine) Static(value bool) sparkLine {
	return s.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (s sparkLine) StaticClassName(value string) sparkLine {
	return s.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (s sparkLine) StaticInputClassName(value string) sparkLine {
	return s.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (s sparkLine) StaticLabelClassName(value string) sparkLine {
	return s.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (s sparkLine) StaticOn(value string) sparkLine {
	return s.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (s sparkLine) StaticPlaceholder(value string) sparkLine {
	return s.set("staticPlaceholder", value)
}

// StaticSchema
func (s sparkLine) StaticSchema(value string) sparkLine {
	return s.set("staticSchema", value)
}

// Style 组件样式
func (s sparkLine) Style(value string) sparkLine {
	return s.set("style", value)
}

// TestIdBuilder
func (s sparkLine) TestIdBuilder(value string) sparkLine {
	return s.set("testIdBuilder", value)
}

// Testid
func (s sparkLine) Testid(value string) sparkLine {
	return s.set("testid", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (s sparkLine) UseMobileUI(value bool) sparkLine {
	return s.set("useMobileUI", value)
}

// Value
func (s sparkLine) Value(value string) sparkLine {
	return s.set("value", value)
}

// Visible 是否显示
func (s sparkLine) Visible(value bool) sparkLine {
	return s.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (s sparkLine) VisibleOn(value string) sparkLine {
	return s.set("visibleOn", value)
}

// Width 宽度
func (s sparkLine) Width(value string) sparkLine {
	return s.set("width", value)
}
