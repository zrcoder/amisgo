package comp

// remark 提示渲染器，默认会显示个小图标，鼠标放上来的时候显示配置的内容.
//
// @version 6.7.0
type remark schema

// Remark 创建一个新的 Remark 实例
func Remark() remark {
	return remark{}.set("type", "remark")
}

func (rm remark) set(key string, value any) remark {
	rm[key] = value
	return rm
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: {"red": "data.progress > 80", "blue": "data.progress > 60"})
func (rm remark) ClassName(value string) remark {
	return rm.set("className", value)
}

// Content 提示内容 (支持两种语法，但是不能混着用。分别是：1. `${xxx}` 或者 `${xxx|upperCase}` 2. `<%= data.xxx %>` 更多文档：https://aisuda.bce.baidu.com/amis/zh-CN/docs/concepts/template)
func (rm remark) Content(value string) remark {
	return rm.set("content", value)
}

// Disabled 是否禁用
func (rm remark) Disabled(value bool) remark {
	return rm.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (rm remark) DisabledOn(value string) remark {
	return rm.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (rm remark) EditorSetting(value string) remark {
	return rm.set("editorSetting", value)
}

// Hidden 是否隐藏
func (rm remark) Hidden(value bool) remark {
	return rm.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (rm remark) HiddenOn(value string) remark {
	return rm.set("hiddenOn", value)
}

// Icon iconfont 里面的类名。
func (rm remark) Icon(value string) remark {
	return rm.set("icon", value)
}

// Id 组件唯一 id，主要用于日志采集
func (rm remark) Id(value string) remark {
	return rm.set("id", value)
}

// Label
func (rm remark) Label(value string) remark {
	return rm.set("label", value)
}

// OnEvent 事件动作配置
func (rm remark) OnEvent(value any) remark {
	return rm.set("onEvent", value)
}

// Placement 显示位置 可选值: top | right | bottom | left
func (rm remark) Placement(value string) remark {
	return rm.set("placement", value)
}

// RootClose 点击其他内容时是否关闭弹框信息
func (rm remark) RootClose(value bool) remark {
	return rm.set("rootClose", value)
}

// Shape icon的形状 可选值: circle | square
func (rm remark) Shape(value string) remark {
	return rm.set("shape", value)
}

// Static 是否静态展示
func (rm remark) Static(value bool) remark {
	return rm.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: {"red": "data.progress > 80", "blue": "data.progress > 60"})
func (rm remark) StaticClassName(value string) remark {
	return rm.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: {"red": "data.progress > 80", "blue": "data.progress > 60"})
func (rm remark) StaticInputClassName(value string) remark {
	return rm.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: {"red": "data.progress > 80", "blue": "data.progress > 60"})
func (rm remark) StaticLabelClassName(value string) remark {
	return rm.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (rm remark) StaticOn(value string) remark {
	return rm.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (rm remark) StaticPlaceholder(value string) remark {
	return rm.set("staticPlaceholder", value)
}

// StaticSchema
func (rm remark) StaticSchema(value string) remark {
	return rm.set("staticSchema", value)
}

// Style 组件样式
func (rm remark) Style(value any) remark {
	return rm.set("style", value)
}

// TestIdBuilder
func (rm remark) TestIdBuilder(value string) remark {
	return rm.set("testIdBuilder", value)
}

// Testid
func (rm remark) Testid(value string) remark {
	return rm.set("testid", value)
}

// Title 提示标题
func (rm remark) Title(value any) remark {
	return rm.set("title", value)
}

// TooltipClassName css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: {"red": "data.progress > 80", "blue": "data.progress > 60"}
func (rm remark) TooltipClassName(value string) remark {
	return rm.set("tooltipClassName", value)
}

// Trigger 触发规则
func (rm remark) Trigger(value string) remark {
	return rm.set("trigger", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (rm remark) UseMobileUI(value bool) remark {
	return rm.set("useMobileUI", value)
}

// Visible 是否显示
func (rm remark) Visible(value bool) remark {
	return rm.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (rm remark) VisibleOn(value string) remark {
	return rm.set("visibleOn", value)
}
