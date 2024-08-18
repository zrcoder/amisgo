package comp

// step 文档
//
// @version 6.7.0
type step schema

// Step 创建一个新的 Step 实例
func Step() step {
	return step{}
}

func (s step) set(key string, value interface{}) step {
	s[key] = value
	return s
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (s step) ClassName(value string) step {
	return s.set("className", value)
}

// Description 描述
func (s step) Description(value string) step {
	return s.set("description", value)
}

// Disabled 是否禁用
func (s step) Disabled(value bool) step {
	return s.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (s step) DisabledOn(value string) step {
	return s.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (s step) EditorSetting(value string) step {
	return s.set("editorSetting", value)
}

// Hidden 是否隐藏
func (s step) Hidden(value bool) step {
	return s.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (s step) HiddenOn(value string) step {
	return s.set("hiddenOn", value)
}

// Icon 图标
func (s step) Icon(value string) step {
	return s.set("icon", value)
}

// Id 组件唯一 id，主要用于日志采集
func (s step) Id(value string) step {
	return s.set("id", value)
}

// OnEvent 事件动作配置
func (s step) OnEvent(value string) step {
	return s.set("onEvent", value)
}

// Static 是否静态展示
func (s step) Static(value bool) step {
	return s.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (s step) StaticClassName(value string) step {
	return s.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (s step) StaticInputClassName(value string) step {
	return s.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (s step) StaticLabelClassName(value string) step {
	return s.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (s step) StaticOn(value string) step {
	return s.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (s step) StaticPlaceholder(value string) step {
	return s.set("staticPlaceholder", value)
}

// StaticSchema
func (s step) StaticSchema(value string) step {
	return s.set("staticSchema", value)
}

// Style 组件样式
func (s step) Style(value string) step {
	return s.set("style", value)
}

// SubTitle 子标题
func (s step) SubTitle(value string) step {
	return s.set("subTitle", value)
}

// TestIdBuilder
func (s step) TestIdBuilder(value string) step {
	return s.set("testIdBuilder", value)
}

// Testid
func (s step) Testid(value string) step {
	return s.set("testid", value)
}

// Title 标题
func (s step) Title(value string) step {
	return s.set("title", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (s step) UseMobileUI(value bool) step {
	return s.set("useMobileUI", value)
}

// Value
func (s step) Value(value string) step {
	return s.set("value", value)
}

// Visible 是否显示
func (s step) Visible(value bool) step {
	return s.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (s step) VisibleOn(value string) step {
	return s.set("visibleOn", value)
}
