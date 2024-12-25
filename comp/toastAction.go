package comp

// toastAction

type toastAction Schema

// ToastAction 创建一个新的 ToastAction 实例
func ToastAction() toastAction {
	return toastAction{}.set("type", "button").set("actionType", "toast")
}

func (ta toastAction) set(key string, value any) toastAction {
	ta[key] = value
	return ta
}

// ActiveClassName 激活状态时的类名
func (ta toastAction) ActiveClassName(value string) toastAction {
	return ta.set("activeClassName", value)
}

// ActiveLevel 激活状态时的样式
func (ta toastAction) ActiveLevel(value string) toastAction {
	return ta.set("activeLevel", value)
}

// Badge 角标 (Badge 角标。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/badge)
func (ta toastAction) Badge(value string) toastAction {
	return ta.set("badge", value)
}

// Block 是否为块状展示，默认为内联。
func (ta toastAction) Block(value bool) toastAction {
	return ta.set("block", value)
}

// Body 子内容 (子内容)
func (ta toastAction) Body(value ...any) toastAction {
	return ta.set("body", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: {"red": "data.progress > 80", "blue": "data.progress > 60"})
func (ta toastAction) ClassName(value string) toastAction {
	return ta.set("className", value)
}

// Close 如果按钮在弹框中，可以配置这个动作完成后是否关闭弹窗，或者指定关闭目标弹框。
func (ta toastAction) Close(value string) toastAction {
	return ta.set("close", value)
}

// ConfirmText 提示文字，配置了操作前会要求用户确认。
func (ta toastAction) ConfirmText(value string) toastAction {
	return ta.set("confirmText", value)
}

// ConfirmTitle 确认弹窗标题
func (ta toastAction) ConfirmTitle(value any) toastAction {
	return ta.set("confirmTitle", value)
}

// CountDown 点击后的禁止倒计时（秒）
func (ta toastAction) CountDown(value string) toastAction {
	return ta.set("countDown", value)
}

// CountDownTpl 倒计时文字自定义
func (ta toastAction) CountDownTpl(value string) toastAction {
	return ta.set("countDownTpl", value)
}

// Disabled 是否禁用
func (ta toastAction) Disabled(value bool) toastAction {
	return ta.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (ta toastAction) DisabledOn(value string) toastAction {
	return ta.set("disabledOn", value)
}

// DisabledTip 禁用时的文案提示。
func (ta toastAction) DisabledTip(value string) toastAction {
	return ta.set("disabledTip", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (ta toastAction) EditorSetting(value string) toastAction {
	return ta.set("editorSetting", value)
}

// Hidden 是否隐藏
func (ta toastAction) Hidden(value bool) toastAction {
	return ta.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (ta toastAction) HiddenOn(value string) toastAction {
	return ta.set("hiddenOn", value)
}

// HotKey 键盘快捷键
func (ta toastAction) HotKey(value string) toastAction {
	return ta.set("hotKey", value)
}

// Icon 按钮图标， iconfont 的类名 (iconfont 里面的类名。)
func (ta toastAction) Icon(value string) toastAction {
	return ta.set("icon", value)
}

// IconClassName icon 上的css 类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: {"red": "data.progress > 80", "blue": "data.progress > 60"})
func (ta toastAction) IconClassName(value string) toastAction {
	return ta.set("iconClassName", value)
}

// Id 主要用于用户行为跟踪里区分是哪个按钮
func (ta toastAction) ID(value string) toastAction {
	return ta.set("id", value)
}

// Label 按钮文字
func (ta toastAction) Label(value string) toastAction {
	return ta.set("label", value)
}

// Level 按钮样式 可选值: info | success | warning | danger | link | primary | dark | light | secondary
func (ta toastAction) Level(value string) toastAction {
	return ta.set("level", value)
}

// LoadingClassName loading 上的css 类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: {"red": "data.progress > 80", "blue": "data.progress > 60"})
func (ta toastAction) LoadingClassName(value string) toastAction {
	return ta.set("loadingClassName", value)
}

// LoadingOn 是否显示loading效果
func (ta toastAction) LoadingOn(value string) toastAction {
	return ta.set("loadingOn", value)
}

// MergeData 是否将弹框中数据 merge 到父级作用域。
func (ta toastAction) MergeData(value bool) toastAction {
	return ta.set("mergeData", value)
}

// OnClick 自定义事件处理函数
func (ta toastAction) OnClick(value string) toastAction {
	return ta.set("onClick", value)
}

// OnEvent 事件动作配置
func (ta toastAction) OnEvent(value any) toastAction {
	return ta.set("onEvent", value)
}

// Primary
func (ta toastAction) Primary(value bool) toastAction {
	return ta.set("primary", value)
}

// RequireSelected 当按钮时批量操作按钮时，默认必须有勾选元素才能可点击，如果此属性配置成 false，则没有点选成员也能点击。
func (ta toastAction) RequireSelected(value bool) toastAction {
	return ta.set("requireSelected", value)
}

// Required 如果按钮在form中，配置此属性会要求用户把指定的字段通过验证后才会触发行为。
func (ta toastAction) Required(value string) toastAction {
	return ta.set("required", value)
}

// RightIcon 右侧按钮图标， iconfont 的类名 (iconfont 里面的类名。)
func (ta toastAction) RightIcon(value string) toastAction {
	return ta.set("rightIcon", value)
}

// RightIconClassName 右侧 icon 上的 css 类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: {"red": "data.progress > 80", "blue": "data.progress > 60"})
func (ta toastAction) RightIconClassName(value string) toastAction {
	return ta.set("rightIconClassName", value)
}

// Size 按钮大小 可选值: xs | sm | md | lg
func (ta toastAction) Size(value string) toastAction {
	return ta.set("size", value)
}

// Static 是否静态展示
func (ta toastAction) Static(value bool) toastAction {
	return ta.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: {"red": "data.progress > 80", "blue": "data.progress > 60"})
func (ta toastAction) StaticClassName(value string) toastAction {
	return ta.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: {"red": "data.progress > 80", "blue": "data.progress > 60"})
func (ta toastAction) StaticInputClassName(value string) toastAction {
	return ta.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: {"red": "data.progress > 80", "blue": "data.progress > 60"})
func (ta toastAction) StaticLabelClassName(value string) toastAction {
	return ta.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (ta toastAction) StaticOn(value string) toastAction {
	return ta.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (ta toastAction) StaticPlaceholder(value string) toastAction {
	return ta.set("staticPlaceholder", value)
}

// StaticSchema
func (ta toastAction) StaticSchema(value string) toastAction {
	return ta.set("staticSchema", value)
}

// Style 组件样式
func (ta toastAction) Style(value any) toastAction {
	return ta.set("style", value)
}

// Target 可以指定让谁来触发这个动作。
func (ta toastAction) Target(value string) toastAction {
	return ta.set("target", value)
}

// TestIdBuilder
func (ta toastAction) TestIdBuilder(value string) toastAction {
	return ta.set("testIdBuilder", value)
}

// Testid
func (ta toastAction) Testid(value string) toastAction {
	return ta.set("testid", value)
}

// Toast 轻提示详情 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/toast (轻提示详情 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/toast)
func (ta toastAction) Toast(value string) toastAction {
	return ta.set("toast", value)
}

// Tooltip
func (ta toastAction) Tooltip(value string) toastAction {
	return ta.set("tooltip", value)
}

// TooltipPlacement 可选值: top | right | bottom | left
func (ta toastAction) TooltipPlacement(value string) toastAction {
	return ta.set("tooltipPlacement", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (ta toastAction) UseMobileUI(value bool) toastAction {
	return ta.set("useMobileUI", value)
}

// Visible 是否显示
func (ta toastAction) Visible(value bool) toastAction {
	return ta.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (ta toastAction) VisibleOn(value string) toastAction {
	return ta.set("visibleOn", value)
}
