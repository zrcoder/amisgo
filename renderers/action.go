package renderers

// Action 行为按钮 https://aisuda.bce.baidu.com/amis/zh-CN/components/action
type Action struct {
	*BaseRenderer
}

// NewAction 创建一个新的 Action 实例
func NewAction() *Action {
	a := &Action{
		BaseRenderer: NewBaseRenderer(),
	}
	a.set("type", "action")
	return a
}

// ActionType 【必填】这是 action 最核心的配置，来指定该 action 的作用类型 可选值: ajax | link | url | drawer | dialog | confirm | cancel | prev | next | copy | close
func (a *Action) ActionType(value string) *Action {
	a.set("actionType", value)
	return a
}

// ActiveClassName 给按钮高亮添加类名。
func (a *Action) ActiveClassName(value string) *Action {
	a.set("activeClassName", value)
	return a
}

// ActiveLevel 按钮高亮时的样式，配置支持同level。
func (a *Action) ActiveLevel(value string) *Action {
	a.set("activeLevel", value)
	return a
}

// Api 设置 API 地址。
func (a *Action) Api(value string) *Action {
	a.set("api", value)
	return a
}

// ClassName 添加类名。
func (a *Action) ClassName(value string) *Action {
	a.set("className", value)
	return a
}

// Close 当action配置在dialog或drawer的actions中时，配置为true指定此次操作完后关闭当前dialog或drawer。当值为字符串，并且是祖先层弹框的名字的时候，会把祖先弹框关闭掉。
func (a *Action) Close(value string) *Action {
	a.set("close", value)
	return a
}

// ConfirmText 当设置后，操作在开始前会询问用户。可用 '$[xxx]' 取值。
func (a *Action) ConfirmText(value string) *Action {
	a.set("confirmText", value)
	return a
}

// ConfirmTitle 确认弹窗标题
func (a *Action) ConfirmTitle(value string) *Action {
	a.set("confirmTitle", value)
	return a
}

// DisabledTip 被禁用后鼠标停留时弹出该段文字，也可以配置对象类型：字段为title和content。可用 '$[xxx]' 取值。
func (a *Action) DisabledTip(value string) *Action {
	a.set("disabledTip", value)
	return a
}

// Icon 设置图标，例如fa fa-plus。
func (a *Action) Icon(value string) *Action {
	a.set("icon", value)
	return a
}

// IconClassName 给图标上添加类名。
func (a *Action) IconClassName(value string) *Action {
	a.set("iconClassName", value)
	return a
}

// Label 按钮文本。可用 '$[xxx]' 取值。
func (a *Action) Label(value string) *Action {
	a.set("label", value)
	return a
}

// Level 按钮样式，支持：link /primary/secondary/info/success/warning/danger/light/dark/default。
func (a *Action) Level(value string) *Action {
	a.set("level", value)
	return a
}

// Link 设置链接。
func (a *Action) Link(value string) *Action {
	a.set("link", value)
	return a
}

// Reload 指定此次操作完后，需要刷新的目标组件名字（组件的name值，自己配置的），多个请用, 号隔开。
func (a *Action) Reload(value string) *Action {
	a.set("reload", value)
	return a
}

// Required 配置字符串数组，指定在form中进行操作之前，需要指定的字段名的表单项通过验证
func (a *Action) Required(value string) *Action {
	a.set("required", value)
	return a
}

// RightIcon 在按钮文本右侧设置图标，例如fa fa-plus。
func (a *Action) RightIcon(value string) *Action {
	a.set("rightIcon", value)
	return a
}

// RightIconClassName 给右侧图标上添加类名。
func (a *Action) RightIconClassName(value string) *Action {
	a.set("rightIconClassName", value)
	return a
}

// Size 按钮大小，支持：xs、sm、md、lg。 可选值: xs | sm | md | lg
func (a *Action) Size(value string) *Action {
	a.set("size", value)
	return a
}

// Tooltip 鼠标停留时弹出该段文字，也可以配置对象类型：字段为title和content。可用 '$[xxx]' 取值。
func (a *Action) Tooltip(value string) *Action {
	a.set("tooltip", value)
	return a
}

// TooltipPlacement 如果配置了tooltip或者disabledTip，指定提示信息位置，可配置top、bottom、left、right。
func (a *Action) TooltipPlacement(value string) *Action {
	a.set("tooltipPlacement", value)
	return a
}
