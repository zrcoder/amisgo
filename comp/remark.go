package comp

// Remark 提示渲染器，默认会显示个小图标，鼠标放上来的时候显示配置的内容.
//
// @version 6.7.0
type Remark Schema

// NewRemark 创建一个新的 Remark 实例
func NewRemark() Remark {
	return Remark{}.set("type", "remark")
}

func (rm Remark) set(key string, value interface{}) Remark {
	rm[key] = value
	return rm
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: {"red": "data.progress > 80", "blue": "data.progress > 60"})
func (rm Remark) ClassName(value string) Remark {
	return rm.set("className", value)
}

// Content 提示内容 (支持两种语法，但是不能混着用。分别是：1. `${xxx}` 或者 `${xxx|upperCase}` 2. `<%= data.xxx %>` 更多文档：https://aisuda.bce.baidu.com/amis/zh-CN/docs/concepts/template)
func (rm Remark) Content(value string) Remark {
	return rm.set("content", value)
}

// Disabled 是否禁用
func (rm Remark) Disabled(value bool) Remark {
	return rm.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (rm Remark) DisabledOn(value string) Remark {
	return rm.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (rm Remark) EditorSetting(value string) Remark {
	return rm.set("editorSetting", value)
}

// Hidden 是否隐藏
func (rm Remark) Hidden(value bool) Remark {
	return rm.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (rm Remark) HiddenOn(value string) Remark {
	return rm.set("hiddenOn", value)
}

// Icon iconfont 里面的类名。
func (rm Remark) Icon(value string) Remark {
	return rm.set("icon", value)
}

// Id 组件唯一 id，主要用于日志采集
func (rm Remark) Id(value string) Remark {
	return rm.set("id", value)
}

// Label
func (rm Remark) Label(value string) Remark {
	return rm.set("label", value)
}

// OnEvent 事件动作配置
func (rm Remark) OnEvent(value string) Remark {
	return rm.set("onEvent", value)
}

// Placement 显示位置 可选值: top | right | bottom | left
func (rm Remark) Placement(value string) Remark {
	return rm.set("placement", value)
}

// RootClose 点击其他内容时是否关闭弹框信息
func (rm Remark) RootClose(value bool) Remark {
	return rm.set("rootClose", value)
}

// Shape icon的形状 可选值: circle | square
func (rm Remark) Shape(value string) Remark {
	return rm.set("shape", value)
}

// Static 是否静态展示
func (rm Remark) Static(value bool) Remark {
	return rm.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: {"red": "data.progress > 80", "blue": "data.progress > 60"})
func (rm Remark) StaticClassName(value string) Remark {
	return rm.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: {"red": "data.progress > 80", "blue": "data.progress > 60"})
func (rm Remark) StaticInputClassName(value string) Remark {
	return rm.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: {"red": "data.progress > 80", "blue": "data.progress > 60"})
func (rm Remark) StaticLabelClassName(value string) Remark {
	return rm.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (rm Remark) StaticOn(value string) Remark {
	return rm.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (rm Remark) StaticPlaceholder(value string) Remark {
	return rm.set("staticPlaceholder", value)
}

// StaticSchema
func (rm Remark) StaticSchema(value string) Remark {
	return rm.set("staticSchema", value)
}

// Style 组件样式
func (rm Remark) Style(value string) Remark {
	return rm.set("style", value)
}

// TestIdBuilder
func (rm Remark) TestIdBuilder(value string) Remark {
	return rm.set("testIdBuilder", value)
}

// Testid
func (rm Remark) Testid(value string) Remark {
	return rm.set("testid", value)
}

// Title 提示标题
func (rm Remark) Title(value string) Remark {
	return rm.set("title", value)
}

// TooltipClassName css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: {"red": "data.progress > 80", "blue": "data.progress > 60"}
func (rm Remark) TooltipClassName(value string) Remark {
	return rm.set("tooltipClassName", value)
}

// Trigger 触发规则
func (rm Remark) Trigger(value string) Remark {
	return rm.set("trigger", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (rm Remark) UseMobileUI(value bool) Remark {
	return rm.set("useMobileUI", value)
}

// Visible 是否显示
func (rm Remark) Visible(value bool) Remark {
	return rm.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (rm Remark) VisibleOn(value string) Remark {
	return rm.set("visibleOn", value)
}
