package comp

// DialogAction 代表一个对话框动作按钮的配置。
type DialogAction Schema

// NewDialogAction 创建一个新的 DialogAction 实例，并初始化默认设置。
func NewDialogAction() DialogAction {
	return make(DialogAction).set("type", "button").set("actionType", "dialog")
}

func (da DialogAction) set(key string, value interface{}) DialogAction {
	da[key] = value
	return da
}

// ActiveClassName 激活状态时的类名
func (da DialogAction) ActiveClassName(value string) DialogAction {
	return da.set("activeClassName", value)
}

// ActiveLevel 激活状态时的样式
func (da DialogAction) ActiveLevel(value string) DialogAction {
	return da.set("activeLevel", value)
}

// Badge 角标 (Badge 角标。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/badge)
func (da DialogAction) Badge(value string) DialogAction {
	return da.set("badge", value)
}

// Block 是否为块状展示，默认为内联。
func (da DialogAction) Block(value bool) DialogAction {
	return da.set("block", value)
}

// Body 子内容 (子内容)
func (da DialogAction) Body(value ...interface{}) DialogAction {
	return da.set("body", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (da DialogAction) ClassName(value string) DialogAction {
	return da.set("className", value)
}

// Close 如果按钮在弹框中，可以配置这个动作完成后是否关闭弹窗，或者指定关闭目标弹框。
func (da DialogAction) Close(value string) DialogAction {
	return da.set("close", value)
}

// ConfirmText 提示文字，配置了操作前会要求用户确认。
func (da DialogAction) ConfirmText(value string) DialogAction {
	return da.set("confirmText", value)
}

// ConfirmTitle 确认弹窗标题
func (da DialogAction) ConfirmTitle(value string) DialogAction {
	return da.set("confirmTitle", value)
}

// CountDown 点击后的禁止倒计时（秒）
func (da DialogAction) CountDown(value string) DialogAction {
	return da.set("countDown", value)
}

// CountDownTpl 倒计时文字自定义
func (da DialogAction) CountDownTpl(value string) DialogAction {
	return da.set("countDownTpl", value)
}

// Data 数据映射
func (da DialogAction) Data(value string) DialogAction {
	return da.set("data", value)
}

// Dialog 弹框详情 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/dialog (弹框详情 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/dialog)
func (da DialogAction) Dialog(value string) DialogAction {
	return da.set("dialog", value)
}

// Disabled 是否禁用
func (da DialogAction) Disabled(value bool) DialogAction {
	return da.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (da DialogAction) DisabledOn(value string) DialogAction {
	return da.set("disabledOn", value)
}

// DisabledTip 禁用时的文案提示。
func (da DialogAction) DisabledTip(value string) DialogAction {
	return da.set("disabledTip", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (da DialogAction) EditorSetting(value string) DialogAction {
	return da.set("editorSetting", value)
}

// Hidden 是否隐藏
func (da DialogAction) Hidden(value bool) DialogAction {
	return da.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (da DialogAction) HiddenOn(value string) DialogAction {
	return da.set("hiddenOn", value)
}

// HotKey 键盘快捷键
func (da DialogAction) HotKey(value string) DialogAction {
	return da.set("hotKey", value)
}

// Icon 按钮图标， iconfont 的类名 (iconfont 里面的类名。)
func (da DialogAction) Icon(value string) DialogAction {
	return da.set("icon", value)
}

// IconClassName icon 上的css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (da DialogAction) IconClassName(value string) DialogAction {
	return da.set("iconClassName", value)
}

// ID 主要用于用户行为跟踪里区分是哪个按钮
func (da DialogAction) ID(value string) DialogAction {
	return da.set("id", value)
}

// Label 按钮文字
func (da DialogAction) Label(value string) DialogAction {
	return da.set("label", value)
}

// Level 按钮样式 可选值: info | success | warning | danger | link | primary | dark | light | secondary
func (da DialogAction) Level(value string) DialogAction {
	return da.set("level", value)
}

// LoadingClassName loading 上的css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (da DialogAction) LoadingClassName(value string) DialogAction {
	return da.set("loadingClassName", value)
}

// LoadingOn 是否显示loading效果
func (da DialogAction) LoadingOn(value string) DialogAction {
	return da.set("loadingOn", value)
}

// MergeData 是否将弹框中数据 merge 到父级作用域。
func (da DialogAction) MergeData(value bool) DialogAction {
	return da.set("mergeData", value)
}

// NextCondition 是否有下一个的表达式，正常可以不用配置，如果想要刷掉某些数据可以配置这个。 (表达式，语法 `data.xxx > 5`。)
func (da DialogAction) NextCondition(value string) DialogAction {
	return da.set("nextCondition", value)
}

// OnClick 自定义事件处理函数
func (da DialogAction) OnClick(value string) DialogAction {
	return da.set("onClick", value)
}

// OnEvent 事件动作配置
func (da DialogAction) OnEvent(value string) DialogAction {
	return da.set("onEvent", value)
}

// Primary
func (da DialogAction) Primary(value bool) DialogAction {
	return da.set("primary", value)
}

// Redirect
func (da DialogAction) Redirect(value string) DialogAction {
	return da.set("redirect", value)
}

// Reload 配置刷新动作，这个动作通常在完成渲染器本省的固定动作后出发。一般用来配置目标组件的 name 属性。多个目标可以用逗号隔开。当目标是 windows 时表示刷新整个页面。刷新目标的同时还支持传递参数如： `foo?a=${a}&b=${b},boo?c=${c}`
func (da DialogAction) Reload(value string) DialogAction {
	return da.set("reload", value)
}

// RequireSelected 当按钮时批量操作按钮时，默认必须有勾选元素才能可点击，如果此属性配置成 false，则没有点选成员也能点击。
func (da DialogAction) RequireSelected(value bool) DialogAction {
	return da.set("requireSelected", value)
}

// Required 如果按钮在form中，配置此属性会要求用户把指定的字段通过验证后才会触发行为。
func (da DialogAction) Required(value string) DialogAction {
	return da.set("required", value)
}

// RightIcon 右侧按钮图标， iconfont 的类名 (iconfont 里面的类名。)
func (da DialogAction) RightIcon(value string) DialogAction {
	return da.set("rightIcon", value)
}

// RightIconClassName 右侧 icon 上的 css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (da DialogAction) RightIconClassName(value string) DialogAction {
	return da.set("rightIconClassName", value)
}

// Size 按钮大小 可选值: xs | sm | md | lg
func (da DialogAction) Size(value string) DialogAction {
	return da.set("size", value)
}

// Static 是否静态展示
func (da DialogAction) Static(value bool) DialogAction {
	return da.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (da DialogAction) StaticClassName(value string) DialogAction {
	return da.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (da DialogAction) StaticInputClassName(value string) DialogAction {
	return da.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (da DialogAction) StaticLabelClassName(value string) DialogAction {
	return da.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (da DialogAction) StaticOn(value string) DialogAction {
	return da.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (da DialogAction) StaticPlaceholder(value string) DialogAction {
	return da.set("staticPlaceholder", value)
}

// StaticSchema
func (da DialogAction) StaticSchema(value string) DialogAction {
	return da.set("staticSchema", value)
}

// Style 组件样式
func (da DialogAction) Style(value string) DialogAction {
	return da.set("style", value)
}

// Target 可以指定让谁来触发这个动作。
func (da DialogAction) Target(value string) DialogAction {
	return da.set("target", value)
}

// TestIdBuilder
func (da DialogAction) TestIdBuilder(value string) DialogAction {
	return da.set("testIdBuilder", value)
}

// Testid
func (da DialogAction) Testid(value string) DialogAction {
	return da.set("testid", value)
}

// Tooltip
func (da DialogAction) Tooltip(value string) DialogAction {
	return da.set("tooltip", value)
}

// TooltipPlacement 可选值: top | right | bottom | left
func (da DialogAction) TooltipPlacement(value string) DialogAction {
	return da.set("tooltipPlacement", value)
}

// Type 指定按钮类型，支持 button、submit或者reset三种类型。 可选值: button | submit | reset
func (da DialogAction) Type(value string) DialogAction {
	return da.set("type", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (da DialogAction) UseMobileUI(value bool) DialogAction {
	return da.set("useMobileUI", value)
}

// Visible 是否显示
func (da DialogAction) Visible(value bool) DialogAction {
	return da.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (da DialogAction) VisibleOn(value string) DialogAction {
	return da.set("visibleOn", value)
}
