package comp

// steps 文档

type steps Schema

// Steps 创建一个新的 Steps 实例
func Steps() steps {
	return steps{}.set("type", "steps")
}

func (s steps) set(key string, value any) steps {
	s[key] = value
	return s
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (s steps) ClassName(value string) steps {
	return s.set("className", value)
}

// Disabled 是否禁用
func (s steps) Disabled(value bool) steps {
	return s.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (s steps) DisabledOn(value string) steps {
	return s.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (s steps) EditorSetting(value string) steps {
	return s.set("editorSetting", value)
}

// Hidden 是否隐藏
func (s steps) Hidden(value bool) steps {
	return s.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (s steps) HiddenOn(value string) steps {
	return s.set("hiddenOn", value)
}

// Id 组件唯一 id，主要用于日志采集
func (s steps) ID(value string) steps {
	return s.set("id", value)
}

// LabelPlacement 标签放置位置 可选值: horizontal | vertical
func (s steps) LabelPlacement(value string) steps {
	return s.set("labelPlacement", value)
}

// Mode 展示模式 可选值: horizontal | vertical
func (s steps) Mode(value string) steps {
	return s.set("mode", value)
}

// Name 变量映射
func (s steps) Name(value string) steps {
	return s.set("name", value)
}

// OnEvent 事件动作配置
func (s steps) OnEvent(value any) steps {
	return s.set("onEvent", value)
}

// ProgressDot 点状步骤条
func (s steps) ProgressDot(value bool) steps {
	return s.set("progressDot", value)
}

// Source API 或 数据映射
func (s steps) Source(value string) steps {
	return s.set("source", value)
}

// Static 是否静态展示
func (s steps) Static(value bool) steps {
	return s.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (s steps) StaticClassName(value string) steps {
	return s.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (s steps) StaticInputClassName(value string) steps {
	return s.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (s steps) StaticLabelClassName(value string) steps {
	return s.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (s steps) StaticOn(value string) steps {
	return s.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (s steps) StaticPlaceholder(value string) steps {
	return s.set("staticPlaceholder", value)
}

// StaticSchema
func (s steps) StaticSchema(value string) steps {
	return s.set("staticSchema", value)
}

// Status
func (s steps) Status(value string) steps {
	return s.set("status", value)
}

// Steps 步骤
func (s steps) Steps(value string) steps {
	return s.set("steps", value)
}

// Style 组件样式
func (s steps) Style(value any) steps {
	return s.set("style", value)
}

// TestIdBuilder
func (s steps) TestIdBuilder(value string) steps {
	return s.set("testIdBuilder", value)
}

// Testid
func (s steps) Testid(value string) steps {
	return s.set("testid", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (s steps) UseMobileUI(value bool) steps {
	return s.set("useMobileUI", value)
}

// Value 指定当前步骤
func (s steps) Value(value string) steps {
	return s.set("value", value)
}

// Visible 是否显示
func (s steps) Visible(value bool) steps {
	return s.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (s steps) VisibleOn(value string) steps {
	return s.set("visibleOn", value)
}
