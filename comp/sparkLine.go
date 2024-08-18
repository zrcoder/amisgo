package comp

// SparkLine
//
// @version 6.7.0
type SparkLine Schema

// NewSparkLine 创建一个新的 SparkLine 实例
func NewSparkLine() SparkLine {
	return SparkLine{}.set("type", "sparkline")
}

func (s SparkLine) set(key string, value interface{}) SparkLine {
	s[key] = value
	return s
}

// ClassName css 类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (s SparkLine) ClassName(value string) SparkLine {
	return s.set("className", value)
}

// ClickAction 点击行为 (点击行为)
func (s SparkLine) ClickAction(value string) SparkLine {
	return s.set("clickAction", value)
}

// Disabled 是否禁用
func (s SparkLine) Disabled(value bool) SparkLine {
	return s.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (s SparkLine) DisabledOn(value string) SparkLine {
	return s.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (s SparkLine) EditorSetting(value string) SparkLine {
	return s.set("editorSetting", value)
}

// Height 高度
func (s SparkLine) Height(value string) SparkLine {
	return s.set("height", value)
}

// Hidden 是否隐藏
func (s SparkLine) Hidden(value bool) SparkLine {
	return s.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (s SparkLine) HiddenOn(value string) SparkLine {
	return s.set("hiddenOn", value)
}

// Id 组件唯一 id，主要用于日志采集
func (s SparkLine) Id(value string) SparkLine {
	return s.set("id", value)
}

// Name 关联数据变量。
func (s SparkLine) Name(value string) SparkLine {
	return s.set("name", value)
}

// OnEvent 事件动作配置
func (s SparkLine) OnEvent(value string) SparkLine {
	return s.set("onEvent", value)
}

// Placeholder 空数据时显示的内容
func (s SparkLine) Placeholder(value string) SparkLine {
	return s.set("placeholder", value)
}

// Static 是否静态展示
func (s SparkLine) Static(value bool) SparkLine {
	return s.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (s SparkLine) StaticClassName(value string) SparkLine {
	return s.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (s SparkLine) StaticInputClassName(value string) SparkLine {
	return s.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (s SparkLine) StaticLabelClassName(value string) SparkLine {
	return s.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (s SparkLine) StaticOn(value string) SparkLine {
	return s.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (s SparkLine) StaticPlaceholder(value string) SparkLine {
	return s.set("staticPlaceholder", value)
}

// StaticSchema
func (s SparkLine) StaticSchema(value string) SparkLine {
	return s.set("staticSchema", value)
}

// Style 组件样式
func (s SparkLine) Style(value string) SparkLine {
	return s.set("style", value)
}

// TestIdBuilder
func (s SparkLine) TestIdBuilder(value string) SparkLine {
	return s.set("testIdBuilder", value)
}

// Testid
func (s SparkLine) Testid(value string) SparkLine {
	return s.set("testid", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (s SparkLine) UseMobileUI(value bool) SparkLine {
	return s.set("useMobileUI", value)
}

// Value
func (s SparkLine) Value(value string) SparkLine {
	return s.set("value", value)
}

// Visible 是否显示
func (s SparkLine) Visible(value bool) SparkLine {
	return s.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (s SparkLine) VisibleOn(value string) SparkLine {
	return s.set("visibleOn", value)
}

// Width 宽度
func (s SparkLine) Width(value string) SparkLine {
	return s.set("width", value)
}
