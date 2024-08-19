package comp

// flex 布局 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/flex
//
// @version 6.7.0
type flex schema

// Flex 创建一个新的 Flex 实例
func Flex() flex {
	return make(flex).set("type", "flex")
}

func (f flex) set(key string, value any) flex {
	f[key] = value
	return f
}

// AlignContent 设置多行情况下的垂直分布，可选值: normal | flex-start | flex-end | center | space-between | space-around | space-evenly | stretch
func (f flex) AlignContent(value string) flex {
	return f.set("alignContent", value)
}

// AlignItems 设置垂直布局，可选值: stretch | start | flex-start | flex-end | end | center | baseline
func (f flex) AlignItems(value ...any) flex {
	return f.set("alignItems", value)
}

// ClassName 设置容器 css 类名
func (f flex) ClassName(value string) flex {
	return f.set("className", value)
}

// Direction 设置方向，可选值: row | column | row-reverse | column-reverse
func (f flex) Direction(value string) flex {
	return f.set("direction", value)
}

// Disabled 设置是否禁用
func (f flex) Disabled(value bool) flex {
	return f.set("disabled", value)
}

// DisabledOn 设置是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (f flex) DisabledOn(value string) flex {
	return f.set("disabledOn", value)
}

// EditorSetting 设置编辑器配置，运行时可以忽略
func (f flex) EditorSetting(value string) flex {
	return f.set("editorSetting", value)
}

// Hidden 设置是否隐藏
func (f flex) Hidden(value bool) flex {
	return f.set("hidden", value)
}

// HiddenOn 设置是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (f flex) HiddenOn(value string) flex {
	return f.set("hiddenOn", value)
}

// ID 设置组件唯一 id，主要用于日志采集
func (f flex) ID(value string) flex {
	return f.set("id", value)
}

// Items 设置每个 flex 的设置 (每个 flex 的设置)
func (f flex) Items(value ...any) flex {
	return f.set("items", value)
}

// Justify 设置水平分布，可选值: start | flex-start | center | end | flex-end | space-around | space-between | space-evenly
func (f flex) Justify(value string) flex {
	return f.set("justify", value)
}

// OnEvent 设置事件动作配置
func (f flex) OnEvent(value any) flex {
	return f.set("onEvent", value)
}

// Static 设置是否静态展示
func (f flex) Static(value bool) flex {
	return f.set("static", value)
}

// StaticClassName 设置静态展示表单项类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (f flex) StaticClassName(value string) flex {
	return f.set("staticClassName", value)
}

// StaticInputClassName 设置静态展示表单项Value类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (f flex) StaticInputClassName(value string) flex {
	return f.set("staticInputClassName", value)
}

// StaticLabelClassName 设置静态展示表单项Label类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (f flex) StaticLabelClassName(value string) flex {
	return f.set("staticLabelClassName", value)
}

// StaticOn 设置是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (f flex) StaticOn(value string) flex {
	return f.set("staticOn", value)
}

// StaticPlaceholder 设置静态展示空值占位
func (f flex) StaticPlaceholder(value string) flex {
	return f.set("staticPlaceholder", value)
}

// StaticSchema 设置
func (f flex) StaticSchema(value string) flex {
	return f.set("staticSchema", value)
}

// Style 设置自定义样式
func (f flex) Style(value any) flex {
	return f.set("style", value)
}

// TestIdBuilder 设置
func (f flex) TestIdBuilder(value string) flex {
	return f.set("testIdBuilder", value)
}

// TestID 设置
func (f flex) TestID(value string) flex {
	return f.set("testid", value)
}

// UseMobileUI 设置可以组件级别用来关闭移动端样式
func (f flex) UseMobileUI(value bool) flex {
	return f.set("useMobileUI", value)
}

// Visible 设置是否显示
func (f flex) Visible(value bool) flex {
	return f.set("visible", value)
}

// VisibleOn 设置是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (f flex) VisibleOn(value string) flex {
	return f.set("visibleOn", value)
}
