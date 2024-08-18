package comp

// Words 表示一个显示文字的组件
//
// @author  slowlyo
// @version 6.7.0
type Words Schema

// NewWords 创建一个新的 Words 实例
func NewWords() Words {
	return Words{}.set("type", "words")
}

func (w Words) set(key string, value interface{}) Words {
	w[key] = value
	return w
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (w Words) ClassName(value string) Words {
	return w.set("className", value)
}

// CollapseButton 展示文字 (展示文字)
func (w Words) CollapseButton(value string) Words {
	return w.set("collapseButton", value)
}

// CollapseButtonText 收起文字
func (w Words) CollapseButtonText(value string) Words {
	return w.set("collapseButtonText", value)
}

// Delimiter 分割符
func (w Words) Delimiter(value string) Words {
	return w.set("delimiter", value)
}

// Disabled 是否禁用
func (w Words) Disabled(value bool) Words {
	return w.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (w Words) DisabledOn(value string) Words {
	return w.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (w Words) EditorSetting(value string) Words {
	return w.set("editorSetting", value)
}

// ExpendButton 展示文字 (展示文字)
func (w Words) ExpendButton(value string) Words {
	return w.set("expendButton", value)
}

// ExpendButtonText 展示文字
func (w Words) ExpendButtonText(value string) Words {
	return w.set("expendButtonText", value)
}

// Hidden 是否隐藏
func (w Words) Hidden(value bool) Words {
	return w.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (w Words) HiddenOn(value string) Words {
	return w.set("hiddenOn", value)
}

// Id 组件唯一 id，主要用于日志采集
func (w Words) Id(value string) Words {
	return w.set("id", value)
}

// InTag 当数据是数组时，是否使用tag的方式展示
func (w Words) InTag(value string) Words {
	return w.set("inTag", value)
}

// Limit 展示限制, 为0时也无限制
func (w Words) Limit(value string) Words {
	return w.set("limit", value)
}

// OnEvent 事件动作配置
func (w Words) OnEvent(value string) Words {
	return w.set("onEvent", value)
}

// Static 是否静态展示
func (w Words) Static(value bool) Words {
	return w.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (w Words) StaticClassName(value string) Words {
	return w.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (w Words) StaticInputClassName(value string) Words {
	return w.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (w Words) StaticLabelClassName(value string) Words {
	return w.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (w Words) StaticOn(value string) Words {
	return w.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (w Words) StaticPlaceholder(value string) Words {
	return w.set("staticPlaceholder", value)
}

// StaticSchema
func (w Words) StaticSchema(value string) Words {
	return w.set("staticSchema", value)
}

// Style 组件样式
func (w Words) Style(value string) Words {
	return w.set("style", value)
}

// TestIdBuilder
func (w Words) TestIdBuilder(value string) Words {
	return w.set("testIdBuilder", value)
}

// Testid
func (w Words) Testid(value string) Words {
	return w.set("testid", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (w Words) UseMobileUI(value bool) Words {
	return w.set("useMobileUI", value)
}

// Visible 是否显示
func (w Words) Visible(value bool) Words {
	return w.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (w Words) VisibleOn(value string) Words {
	return w.set("visibleOn", value)
}

// Words tags数据
func (w Words) Words(value string) Words {
	return w.set("words", value)
}
