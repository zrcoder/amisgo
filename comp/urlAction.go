package comp

// urlAction URL 操作按钮控件
//
// @version 6.7.0
type urlAction schema

// UrlAction 创建一个新的 UrlAction 实例
func UrlAction() urlAction {
	return urlAction{}.set("type", "button").set("actionType", "url")
}

func (ua urlAction) set(key string, value any) urlAction {
	ua[key] = value
	return ua
}

// ActiveClassName 激活状态时的类名
func (ua urlAction) ActiveClassName(value string) urlAction {
	return ua.set("activeClassName", value)
}

// ActiveLevel 激活状态时的样式
func (ua urlAction) ActiveLevel(value string) urlAction {
	return ua.set("activeLevel", value)
}

// Badge 角标 (Badge 角标。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/badge)
func (ua urlAction) Badge(value string) urlAction {
	return ua.set("badge", value)
}

// Blank 是否新窗口打开
func (ua urlAction) Blank(value bool) urlAction {
	return ua.set("blank", value)
}

// Block 是否为块状展示，默认为内联。
func (ua urlAction) Block(value bool) urlAction {
	return ua.set("block", value)
}

// Body 子内容 (子内容)
func (ua urlAction) Body(value ...any) urlAction {
	return ua.set("body", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (ua urlAction) ClassName(value string) urlAction {
	return ua.set("className", value)
}

// Close 如果按钮在弹框中，可以配置这个动作完成后是否关闭弹窗，或者指定关闭目标弹框。
func (ua urlAction) Close(value string) urlAction {
	return ua.set("close", value)
}

// ConfirmText 提示文字，配置了操作前会要求用户确认。
func (ua urlAction) ConfirmText(value string) urlAction {
	return ua.set("confirmText", value)
}

// ConfirmTitle 确认弹窗标题
func (ua urlAction) ConfirmTitle(value any) urlAction {
	return ua.set("confirmTitle", value)
}

// CountDown 点击后的禁止倒计时（秒）
func (ua urlAction) CountDown(value string) urlAction {
	return ua.set("countDown", value)
}

// CountDownTpl 倒计时文字自定义
func (ua urlAction) CountDownTpl(value string) urlAction {
	return ua.set("countDownTpl", value)
}

// Disabled 是否禁用
func (ua urlAction) Disabled(value bool) urlAction {
	return ua.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (ua urlAction) DisabledOn(value string) urlAction {
	return ua.set("disabledOn", value)
}

// DisabledTip 禁用时的文案提示。
func (ua urlAction) DisabledTip(value string) urlAction {
	return ua.set("disabledTip", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (ua urlAction) EditorSetting(value string) urlAction {
	return ua.set("editorSetting", value)
}

// Hidden 是否隐藏
func (ua urlAction) Hidden(value bool) urlAction {
	return ua.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (ua urlAction) HiddenOn(value string) urlAction {
	return ua.set("hiddenOn", value)
}

// HotKey 键盘快捷键
func (ua urlAction) HotKey(value string) urlAction {
	return ua.set("hotKey", value)
}

// Icon 按钮图标， iconfont 的类名 (iconfont 里面的类名。)
func (ua urlAction) Icon(value string) urlAction {
	return ua.set("icon", value)
}

// IconClassName icon 上的css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (ua urlAction) IconClassName(value string) urlAction {
	return ua.set("iconClassName", value)
}

// Id 主要用于用户行为跟踪里区分是哪个按钮
func (ua urlAction) Id(value string) urlAction {
	return ua.set("id", value)
}

// Label 按钮文字
func (ua urlAction) Label(value string) urlAction {
	return ua.set("label", value)
}

// Level 按钮样式 可选值: info | success | warning | danger | link | primary | dark | light | secondary
func (ua urlAction) Level(value string) urlAction {
	return ua.set("level", value)
}

// Link 设置链接
func (ua urlAction) Link(value string) urlAction {
	return ua.set("link", value)
}

// LoadingClassName loading 上的css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (ua urlAction) LoadingClassName(value string) urlAction {
	return ua.set("loadingClassName", value)
}

// LoadingOn 是否显示loading效果
func (ua urlAction) LoadingOn(value string) urlAction {
	return ua.set("loadingOn", value)
}

// MergeData 是否将弹框中数据 merge 到父级作用域。
func (ua urlAction) MergeData(value bool) urlAction {
	return ua.set("mergeData", value)
}

// OnClick 自定义事件处理函数
func (ua urlAction) OnClick(value string) urlAction {
	return ua.set("onClick", value)
}

// OnEvent 事件动作配置
func (ua urlAction) OnEvent(value any) urlAction {
	return ua.set("onEvent", value)
}

// Primary
func (ua urlAction) Primary(value bool) urlAction {
	return ua.set("primary", value)
}

// RequireSelected 当按钮时批量操作按钮时，默认必须有勾选元素才能可点击，如果此属性配置成 false，则没有点选成员也能点击。
func (ua urlAction) RequireSelected(value bool) urlAction {
	return ua.set("requireSelected", value)
}

// Required 如果按钮在form中，配置此属性会要求用户把指定的字段通过验证后才会触发行为。
func (ua urlAction) Required(value string) urlAction {
	return ua.set("required", value)
}

// RightIcon 右侧按钮图标， iconfont 的类名 (iconfont 里面的类名。)
func (ua urlAction) RightIcon(value string) urlAction {
	return ua.set("rightIcon", value)
}

// RightIconClassName 右侧 icon 上的 css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (ua urlAction) RightIconClassName(value string) urlAction {
	return ua.set("rightIconClassName", value)
}

// Size 按钮大小 可选值: xs | sm | md | lg
func (ua urlAction) Size(value string) urlAction {
	return ua.set("size", value)
}

// Static 是否静态展示
func (ua urlAction) Static(value bool) urlAction {
	return ua.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (ua urlAction) StaticClassName(value string) urlAction {
	return ua.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (ua urlAction) StaticInputClassName(value string) urlAction {
	return ua.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (ua urlAction) StaticLabelClassName(value string) urlAction {
	return ua.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (ua urlAction) StaticOn(value string) urlAction {
	return ua.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (ua urlAction) StaticPlaceholder(value string) urlAction {
	return ua.set("staticPlaceholder", value)
}

// StaticSchema
func (ua urlAction) StaticSchema(value string) urlAction {
	return ua.set("staticSchema", value)
}

// Style 组件样式
func (ua urlAction) Style(value any) urlAction {
	return ua.set("style", value)
}

// Target 可以指定让谁来触发这个动作。
func (ua urlAction) Target(value string) urlAction {
	return ua.set("target", value)
}

// TestIdBuilder
func (ua urlAction) TestIdBuilder(value string) urlAction {
	return ua.set("testIdBuilder", value)
}

// Testid
func (ua urlAction) Testid(value string) urlAction {
	return ua.set("testid", value)
}

// Tooltip
func (ua urlAction) Tooltip(value string) urlAction {
	return ua.set("tooltip", value)
}

// TooltipPlacement 可选值: top | right | bottom | left
func (ua urlAction) TooltipPlacement(value string) urlAction {
	return ua.set("tooltipPlacement", value)
}

// Type 指定按钮类型，支持 button、submit或者reset三种类型。 可选值: button | submit | reset
func (ua urlAction) Type(value string) urlAction {
	return ua.set("type", value)
}

// Url 打开的目标地址
func (ua urlAction) Url(value string) urlAction {
	return ua.set("url", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (ua urlAction) UseMobileUI(value bool) urlAction {
	return ua.set("useMobileUI", value)
}

// Visible 是否显示
func (ua urlAction) Visible(value bool) urlAction {
	return ua.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (ua urlAction) VisibleOn(value string) urlAction {
	return ua.set("visibleOn", value)
}
