package comp

// ToastAction
//
// @author slowlyo
// @version 6.7.0
type ToastAction Schema

// NewToastAction 创建一个新的 ToastAction 实例
func NewToastAction() ToastAction {
	return ToastAction{}.set("type", "button").set("actionType", "toast")
}

func (ta ToastAction) set(key string, value interface{}) ToastAction {
	ta[key] = value
	return ta
}

// ActionType 指定为打开弹窗，抽出式弹窗
func (ta ToastAction) ActionType(value string) ToastAction {
	return ta.set("actionType", value)
}

// ActiveClassName 激活状态时的类名
func (ta ToastAction) ActiveClassName(value string) ToastAction {
	return ta.set("activeClassName", value)
}

// ActiveLevel 激活状态时的样式
func (ta ToastAction) ActiveLevel(value string) ToastAction {
	return ta.set("activeLevel", value)
}

// Badge 角标 (Badge 角标。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/badge)
func (ta ToastAction) Badge(value string) ToastAction {
	return ta.set("badge", value)
}

// Block 是否为块状展示，默认为内联。
func (ta ToastAction) Block(value bool) ToastAction {
	return ta.set("block", value)
}

// Body 子内容 (子内容)
func (ta ToastAction) Body(value ...interface{}) ToastAction {
	return ta.set("body", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: {"red": "data.progress > 80", "blue": "data.progress > 60"})
func (ta ToastAction) ClassName(value string) ToastAction {
	return ta.set("className", value)
}

// Close 如果按钮在弹框中，可以配置这个动作完成后是否关闭弹窗，或者指定关闭目标弹框。
func (ta ToastAction) Close(value string) ToastAction {
	return ta.set("close", value)
}

// ConfirmText 提示文字，配置了操作前会要求用户确认。
func (ta ToastAction) ConfirmText(value string) ToastAction {
	return ta.set("confirmText", value)
}

// ConfirmTitle 确认弹窗标题
func (ta ToastAction) ConfirmTitle(value string) ToastAction {
	return ta.set("confirmTitle", value)
}

// CountDown 点击后的禁止倒计时（秒）
func (ta ToastAction) CountDown(value string) ToastAction {
	return ta.set("countDown", value)
}

// CountDownTpl 倒计时文字自定义
func (ta ToastAction) CountDownTpl(value string) ToastAction {
	return ta.set("countDownTpl", value)
}

// Disabled 是否禁用
func (ta ToastAction) Disabled(value bool) ToastAction {
	return ta.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (ta ToastAction) DisabledOn(value string) ToastAction {
	return ta.set("disabledOn", value)
}

// DisabledTip 禁用时的文案提示。
func (ta ToastAction) DisabledTip(value string) ToastAction {
	return ta.set("disabledTip", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (ta ToastAction) EditorSetting(value string) ToastAction {
	return ta.set("editorSetting", value)
}

// Hidden 是否隐藏
func (ta ToastAction) Hidden(value bool) ToastAction {
	return ta.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (ta ToastAction) HiddenOn(value string) ToastAction {
	return ta.set("hiddenOn", value)
}

// HotKey 键盘快捷键
func (ta ToastAction) HotKey(value string) ToastAction {
	return ta.set("hotKey", value)
}

// Icon 按钮图标， iconfont 的类名 (iconfont 里面的类名。)
func (ta ToastAction) Icon(value string) ToastAction {
	return ta.set("icon", value)
}

// IconClassName icon 上的css 类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: {"red": "data.progress > 80", "blue": "data.progress > 60"})
func (ta ToastAction) IconClassName(value string) ToastAction {
	return ta.set("iconClassName", value)
}

// Id 主要用于用户行为跟踪里区分是哪个按钮
func (ta ToastAction) Id(value string) ToastAction {
	return ta.set("id", value)
}

// Label 按钮文字
func (ta ToastAction) Label(value string) ToastAction {
	return ta.set("label", value)
}

// Level 按钮样式 可选值: info | success | warning | danger | link | primary | dark | light | secondary
func (ta ToastAction) Level(value string) ToastAction {
	return ta.set("level", value)
}

// LoadingClassName loading 上的css 类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: {"red": "data.progress > 80", "blue": "data.progress > 60"})
func (ta ToastAction) LoadingClassName(value string) ToastAction {
	return ta.set("loadingClassName", value)
}

// LoadingOn 是否显示loading效果
func (ta ToastAction) LoadingOn(value string) ToastAction {
	return ta.set("loadingOn", value)
}

// MergeData 是否将弹框中数据 merge 到父级作用域。
func (ta ToastAction) MergeData(value bool) ToastAction {
	return ta.set("mergeData", value)
}

// OnClick 自定义事件处理函数
func (ta ToastAction) OnClick(value string) ToastAction {
	return ta.set("onClick", value)
}

// OnEvent 事件动作配置
func (ta ToastAction) OnEvent(value string) ToastAction {
	return ta.set("onEvent", value)
}

// Primary
func (ta ToastAction) Primary(value bool) ToastAction {
	return ta.set("primary", value)
}

// RequireSelected 当按钮时批量操作按钮时，默认必须有勾选元素才能可点击，如果此属性配置成 false，则没有点选成员也能点击。
func (ta ToastAction) RequireSelected(value bool) ToastAction {
	return ta.set("requireSelected", value)
}

// Required 如果按钮在form中，配置此属性会要求用户把指定的字段通过验证后才会触发行为。
func (ta ToastAction) Required(value string) ToastAction {
	return ta.set("required", value)
}

// RightIcon 右侧按钮图标， iconfont 的类名 (iconfont 里面的类名。)
func (ta ToastAction) RightIcon(value string) ToastAction {
	return ta.set("rightIcon", value)
}

// RightIconClassName 右侧 icon 上的 css 类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: {"red": "data.progress > 80", "blue": "data.progress > 60"})
func (ta ToastAction) RightIconClassName(value string) ToastAction {
	return ta.set("rightIconClassName", value)
}

// Size 按钮大小 可选值: xs | sm | md | lg
func (ta ToastAction) Size(value string) ToastAction {
	return ta.set("size", value)
}

// Static 是否静态展示
func (ta ToastAction) Static(value bool) ToastAction {
	return ta.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: {"red": "data.progress > 80", "blue": "data.progress > 60"})
func (ta ToastAction) StaticClassName(value string) ToastAction {
	return ta.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: {"red": "data.progress > 80", "blue": "data.progress > 60"})
func (ta ToastAction) StaticInputClassName(value string) ToastAction {
	return ta.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: {"red": "data.progress > 80", "blue": "data.progress > 60"})
func (ta ToastAction) StaticLabelClassName(value string) ToastAction {
	return ta.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (ta ToastAction) StaticOn(value string) ToastAction {
	return ta.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (ta ToastAction) StaticPlaceholder(value string) ToastAction {
	return ta.set("staticPlaceholder", value)
}

// StaticSchema
func (ta ToastAction) StaticSchema(value string) ToastAction {
	return ta.set("staticSchema", value)
}

// Style 组件样式
func (ta ToastAction) Style(value string) ToastAction {
	return ta.set("style", value)
}

// Target 可以指定让谁来触发这个动作。
func (ta ToastAction) Target(value string) ToastAction {
	return ta.set("target", value)
}

// TestIdBuilder
func (ta ToastAction) TestIdBuilder(value string) ToastAction {
	return ta.set("testIdBuilder", value)
}

// Testid
func (ta ToastAction) Testid(value string) ToastAction {
	return ta.set("testid", value)
}

// Toast 轻提示详情 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/toast (轻提示详情 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/toast)
func (ta ToastAction) Toast(value string) ToastAction {
	return ta.set("toast", value)
}

// Tooltip
func (ta ToastAction) Tooltip(value string) ToastAction {
	return ta.set("tooltip", value)
}

// TooltipPlacement 可选值: top | right | bottom | left
func (ta ToastAction) TooltipPlacement(value string) ToastAction {
	return ta.set("tooltipPlacement", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (ta ToastAction) UseMobileUI(value bool) ToastAction {
	return ta.set("useMobileUI", value)
}

// Visible 是否显示
func (ta ToastAction) Visible(value bool) ToastAction {
	return ta.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (ta ToastAction) VisibleOn(value string) ToastAction {
	return ta.set("visibleOn", value)
}
