package comp

// Action 行为按钮 https://aisuda.bce.baidu.com/amis/zh-CN/components/action
type Action BaseRenderer

// NewAction 创建一个新的 Action 实例
func NewAction() Action {
	a := Action(make(BaseRenderer))
	return a.set("type", "action")
}

// ActionType 【必填】这是 action 最核心的配置，来指定该 action 的作用类型 可选值: ajax | link | url | drawer | dialog | confirm | cancel | prev | next | copy | close
func (a Action) ActionType(value string) Action {
	return a.set("actionType", value)
}

// ActiveClassName 给按钮高亮添加类名。
func (a Action) ActiveClassName(value string) Action {
	return a.set("activeClassName", value)
}

// ActiveLevel 按钮高亮时的样式，配置支持同level。
func (a Action) ActiveLevel(value string) Action {
	return a.set("activeLevel", value)
}

// Api 设置 API 地址。
func (a Action) Api(value string) Action {
	return a.set("api", value)
}

// ClassName 添加类名。
func (a Action) ClassName(value string) Action {
	return a.set("className", value)
}

// Close 当action配置在dialog或drawer的actions中时，配置为true指定此次操作完后关闭当前dialog或drawer。当值为字符串，并且是祖先层弹框的名字的时候，会把祖先弹框关闭掉。
func (a Action) Close(value string) Action {
	return a.set("close", value)
}

// ConfirmText 当设置后，操作在开始前会询问用户。可用 '$[xxx]' 取值。
func (a Action) ConfirmText(value string) Action {
	return a.set("confirmText", value)
}

// ConfirmTitle 确认弹窗标题
func (a Action) ConfirmTitle(value string) Action {
	return a.set("confirmTitle", value)
}

// DisabledTip 被禁用后鼠标停留时弹出该段文字，也可以配置对象类型：字段为title和content。可用 '$[xxx]' 取值。
func (a Action) DisabledTip(value string) Action {
	return a.set("disabledTip", value)
}

// Icon 设置图标，例如fa fa-plus。
func (a Action) Icon(value string) Action {
	return a.set("icon", value)
}

// IconClassName 给图标上添加类名。
func (a Action) IconClassName(value string) Action {
	return a.set("iconClassName", value)
}

// Label 按钮文本。可用 '$[xxx]' 取值。
func (a Action) Label(value string) Action {
	return a.set("label", value)
}

// Level 按钮样式，支持：link /primary/secondary/info/success/warning/danger/light/dark/default。
func (a Action) Level(value string) Action {
	return a.set("level", value)
}

// Link 设置链接。
func (a Action) Link(value string) Action {
	return a.set("link", value)
}

// Reload 指定此次操作完后，需要刷新的目标组件名字（组件的name值，自己配置的），多个请用, 号隔开。
func (a Action) Reload(value string) Action {
	return a.set("reload", value)
}

// Required 配置字符串数组，指定在form中进行操作之前，需要指定的字段名的表单项通过验证
func (a Action) Required(value string) Action {
	return a.set("required", value)
}

// RightIcon 在按钮文本右侧设置图标，例如fa fa-plus。
func (a Action) RightIcon(value string) Action {
	return a.set("rightIcon", value)
}

// RightIconClassName 给右侧图标上添加类名。
func (a Action) RightIconClassName(value string) Action {
	return a.set("rightIconClassName", value)
}

// Size 按钮大小，支持：xs、sm、md、lg。 可选值: xs | sm | md | lg
func (a Action) Size(value string) Action {
	return a.set("size", value)
}

// Tooltip 鼠标停留时弹出该段文字，也可以配置对象类型：字段为title和content。可用 '$[xxx]' 取值。
func (a Action) Tooltip(value string) Action {
	return a.set("tooltip", value)
}

// TooltipPlacement 如果配置了tooltip或者disabledTip，指定提示信息位置，可配置top、bottom、left、right。
func (a Action) TooltipPlacement(value string) Action {
	return a.set("tooltipPlacement", value)
}

func (a Action) set(key string, value interface{}) Action {
	a[key] = value
	return a
}
