package comp

// words 表示一个显示文字的组件

// @version 6.7.0
type words schema

// Words 创建一个新的 Words 实例
func Words() words {
	return words{}.set("type", "words")
}

func (w words) set(key string, value interface{}) words {
	w[key] = value
	return w
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (w words) ClassName(value string) words {
	return w.set("className", value)
}

// CollapseButton 展示文字 (展示文字)
func (w words) CollapseButton(value string) words {
	return w.set("collapseButton", value)
}

// CollapseButtonText 收起文字
func (w words) CollapseButtonText(value string) words {
	return w.set("collapseButtonText", value)
}

// Delimiter 分割符
func (w words) Delimiter(value string) words {
	return w.set("delimiter", value)
}

// Disabled 是否禁用
func (w words) Disabled(value bool) words {
	return w.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (w words) DisabledOn(value string) words {
	return w.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (w words) EditorSetting(value string) words {
	return w.set("editorSetting", value)
}

// ExpendButton 展示文字 (展示文字)
func (w words) ExpendButton(value string) words {
	return w.set("expendButton", value)
}

// ExpendButtonText 展示文字
func (w words) ExpendButtonText(value string) words {
	return w.set("expendButtonText", value)
}

// Hidden 是否隐藏
func (w words) Hidden(value bool) words {
	return w.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (w words) HiddenOn(value string) words {
	return w.set("hiddenOn", value)
}

// Id 组件唯一 id，主要用于日志采集
func (w words) Id(value string) words {
	return w.set("id", value)
}

// InTag 当数据是数组时，是否使用tag的方式展示
func (w words) InTag(value string) words {
	return w.set("inTag", value)
}

// Limit 展示限制, 为0时也无限制
func (w words) Limit(value string) words {
	return w.set("limit", value)
}

// OnEvent 事件动作配置
func (w words) OnEvent(value string) words {
	return w.set("onEvent", value)
}

// Static 是否静态展示
func (w words) Static(value bool) words {
	return w.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (w words) StaticClassName(value string) words {
	return w.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (w words) StaticInputClassName(value string) words {
	return w.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (w words) StaticLabelClassName(value string) words {
	return w.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (w words) StaticOn(value string) words {
	return w.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (w words) StaticPlaceholder(value string) words {
	return w.set("staticPlaceholder", value)
}

// StaticSchema
func (w words) StaticSchema(value string) words {
	return w.set("staticSchema", value)
}

// Style 组件样式
func (w words) Style(value string) words {
	return w.set("style", value)
}

// TestIdBuilder
func (w words) TestIdBuilder(value string) words {
	return w.set("testIdBuilder", value)
}

// Testid
func (w words) Testid(value string) words {
	return w.set("testid", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (w words) UseMobileUI(value bool) words {
	return w.set("useMobileUI", value)
}

// Visible 是否显示
func (w words) Visible(value bool) words {
	return w.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (w words) VisibleOn(value string) words {
	return w.set("visibleOn", value)
}

// Words tags数据
func (w words) Words(value string) words {
	return w.set("words", value)
}
