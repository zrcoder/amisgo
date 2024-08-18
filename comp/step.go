package comp

// Step 文档
//
// @version 6.7.0
type Step Schema

// NewStep 创建一个新的 Step 实例
func NewStep() Step {
	return Step{}
}

func (s Step) set(key string, value interface{}) Step {
	s[key] = value
	return s
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (s Step) ClassName(value string) Step {
	return s.set("className", value)
}

// Description 描述
func (s Step) Description(value string) Step {
	return s.set("description", value)
}

// Disabled 是否禁用
func (s Step) Disabled(value bool) Step {
	return s.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (s Step) DisabledOn(value string) Step {
	return s.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (s Step) EditorSetting(value string) Step {
	return s.set("editorSetting", value)
}

// Hidden 是否隐藏
func (s Step) Hidden(value bool) Step {
	return s.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (s Step) HiddenOn(value string) Step {
	return s.set("hiddenOn", value)
}

// Icon 图标
func (s Step) Icon(value string) Step {
	return s.set("icon", value)
}

// Id 组件唯一 id，主要用于日志采集
func (s Step) Id(value string) Step {
	return s.set("id", value)
}

// OnEvent 事件动作配置
func (s Step) OnEvent(value string) Step {
	return s.set("onEvent", value)
}

// Static 是否静态展示
func (s Step) Static(value bool) Step {
	return s.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (s Step) StaticClassName(value string) Step {
	return s.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (s Step) StaticInputClassName(value string) Step {
	return s.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (s Step) StaticLabelClassName(value string) Step {
	return s.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (s Step) StaticOn(value string) Step {
	return s.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (s Step) StaticPlaceholder(value string) Step {
	return s.set("staticPlaceholder", value)
}

// StaticSchema
func (s Step) StaticSchema(value string) Step {
	return s.set("staticSchema", value)
}

// Style 组件样式
func (s Step) Style(value string) Step {
	return s.set("style", value)
}

// SubTitle 子标题
func (s Step) SubTitle(value string) Step {
	return s.set("subTitle", value)
}

// TestIdBuilder
func (s Step) TestIdBuilder(value string) Step {
	return s.set("testIdBuilder", value)
}

// Testid
func (s Step) Testid(value string) Step {
	return s.set("testid", value)
}

// Title 标题
func (s Step) Title(value string) Step {
	return s.set("title", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (s Step) UseMobileUI(value bool) Step {
	return s.set("useMobileUI", value)
}

// Value
func (s Step) Value(value string) Step {
	return s.set("value", value)
}

// Visible 是否显示
func (s Step) Visible(value bool) Step {
	return s.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (s Step) VisibleOn(value string) Step {
	return s.set("visibleOn", value)
}
