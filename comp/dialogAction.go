package comp

// dialogAction 代表一个对话框动作按钮的配置。
type dialogAction Schema

// DialogAction 创建一个新的 DialogAction 实例，并初始化默认设置。
func DialogAction() dialogAction {
	return make(dialogAction).set("type", "button").set("actionType", "dialog")
}

func (da dialogAction) set(key string, value any) dialogAction {
	da[key] = value
	return da
}

// ActiveClassName 激活状态时的类名
func (da dialogAction) ActiveClassName(value string) dialogAction {
	return da.set("activeClassName", value)
}

// ActiveLevel 激活状态时的样式
func (da dialogAction) ActiveLevel(value string) dialogAction {
	return da.set("activeLevel", value)
}

// Badge 角标 (Badge 角标。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/badge)
func (da dialogAction) Badge(value string) dialogAction {
	return da.set("badge", value)
}

// Block 是否为块状展示，默认为内联。
func (da dialogAction) Block(value bool) dialogAction {
	return da.set("block", value)
}

// Body 子内容 (子内容)
func (da dialogAction) Body(value ...any) dialogAction {
	return da.set("body", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (da dialogAction) ClassName(value string) dialogAction {
	return da.set("className", value)
}

// Close 如果按钮在弹框中，可以配置这个动作完成后是否关闭弹窗，或者指定关闭目标弹框。
func (da dialogAction) Close(value string) dialogAction {
	return da.set("close", value)
}

// ConfirmText 提示文字，配置了操作前会要求用户确认。
func (da dialogAction) ConfirmText(value string) dialogAction {
	return da.set("confirmText", value)
}

// ConfirmTitle 确认弹窗标题
func (da dialogAction) ConfirmTitle(value any) dialogAction {
	return da.set("confirmTitle", value)
}

// CountDown 点击后的禁止倒计时（秒）
func (da dialogAction) CountDown(value string) dialogAction {
	return da.set("countDown", value)
}

// CountDownTpl 倒计时文字自定义
func (da dialogAction) CountDownTpl(value string) dialogAction {
	return da.set("countDownTpl", value)
}

// Data 数据映射
func (da dialogAction) Data(value Data) dialogAction {
	return da.set("data", value)
}

// Dialog 弹框详情 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/dialog (弹框详情 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/dialog)
func (da dialogAction) Dialog(value string) dialogAction {
	return da.set("dialog", value)
}

// Disabled 是否禁用
func (da dialogAction) Disabled(value bool) dialogAction {
	return da.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (da dialogAction) DisabledOn(value string) dialogAction {
	return da.set("disabledOn", value)
}

// DisabledTip 禁用时的文案提示。
func (da dialogAction) DisabledTip(value string) dialogAction {
	return da.set("disabledTip", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (da dialogAction) EditorSetting(value string) dialogAction {
	return da.set("editorSetting", value)
}

// Hidden 是否隐藏
func (da dialogAction) Hidden(value bool) dialogAction {
	return da.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (da dialogAction) HiddenOn(value string) dialogAction {
	return da.set("hiddenOn", value)
}

// HotKey 键盘快捷键
func (da dialogAction) HotKey(value string) dialogAction {
	return da.set("hotKey", value)
}

// Icon 按钮图标， iconfont 的类名 (iconfont 里面的类名。)
func (da dialogAction) Icon(value string) dialogAction {
	return da.set("icon", value)
}

// IconClassName icon 上的css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (da dialogAction) IconClassName(value string) dialogAction {
	return da.set("iconClassName", value)
}

// ID 主要用于用户行为跟踪里区分是哪个按钮
func (da dialogAction) ID(value string) dialogAction {
	return da.set("id", value)
}

// Label 按钮文字
func (da dialogAction) Label(value string) dialogAction {
	return da.set("label", value)
}

// Level 按钮样式 可选值: info | success | warning | danger | link | primary | dark | light | secondary
func (da dialogAction) Level(value string) dialogAction {
	return da.set("level", value)
}

// LoadingClassName loading 上的css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (da dialogAction) LoadingClassName(value string) dialogAction {
	return da.set("loadingClassName", value)
}

// LoadingOn 是否显示loading效果
func (da dialogAction) LoadingOn(value string) dialogAction {
	return da.set("loadingOn", value)
}

// MergeData 是否将弹框中数据 merge 到父级作用域。
func (da dialogAction) MergeData(value bool) dialogAction {
	return da.set("mergeData", value)
}

// NextCondition 是否有下一个的表达式，正常可以不用配置，如果想要刷掉某些数据可以配置这个。 (表达式，语法 `data.xxx > 5`。)
func (da dialogAction) NextCondition(value string) dialogAction {
	return da.set("nextCondition", value)
}

// OnClick 自定义事件处理函数
func (da dialogAction) OnClick(value string) dialogAction {
	return da.set("onClick", value)
}

// OnEvent 事件动作配置
func (da dialogAction) OnEvent(value any) dialogAction {
	return da.set("onEvent", value)
}

// Primary
func (da dialogAction) Primary(value bool) dialogAction {
	return da.set("primary", value)
}

// Redirect
func (da dialogAction) Redirect(value string) dialogAction {
	return da.set("redirect", value)
}

// Reload 配置刷新动作，这个动作通常在完成渲染器本省的固定动作后出发。一般用来配置目标组件的 name 属性。多个目标可以用逗号隔开。当目标是 windows 时表示刷新整个页面。刷新目标的同时还支持传递参数如： `foo?a=${a}&b=${b},boo?c=${c}`
func (da dialogAction) Reload(value string) dialogAction {
	return da.set("reload", value)
}

// RequireSelected 当按钮时批量操作按钮时，默认必须有勾选元素才能可点击，如果此属性配置成 false，则没有点选成员也能点击。
func (da dialogAction) RequireSelected(value bool) dialogAction {
	return da.set("requireSelected", value)
}

// Required 如果按钮在form中，配置此属性会要求用户把指定的字段通过验证后才会触发行为。
func (da dialogAction) Required(value string) dialogAction {
	return da.set("required", value)
}

// RightIcon 右侧按钮图标， iconfont 的类名 (iconfont 里面的类名。)
func (da dialogAction) RightIcon(value string) dialogAction {
	return da.set("rightIcon", value)
}

// RightIconClassName 右侧 icon 上的 css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (da dialogAction) RightIconClassName(value string) dialogAction {
	return da.set("rightIconClassName", value)
}

// Size 按钮大小 可选值: xs | sm | md | lg
func (da dialogAction) Size(value string) dialogAction {
	return da.set("size", value)
}

// Static 是否静态展示
func (da dialogAction) Static(value bool) dialogAction {
	return da.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (da dialogAction) StaticClassName(value string) dialogAction {
	return da.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (da dialogAction) StaticInputClassName(value string) dialogAction {
	return da.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (da dialogAction) StaticLabelClassName(value string) dialogAction {
	return da.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (da dialogAction) StaticOn(value string) dialogAction {
	return da.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (da dialogAction) StaticPlaceholder(value string) dialogAction {
	return da.set("staticPlaceholder", value)
}

// StaticSchema
func (da dialogAction) StaticSchema(value string) dialogAction {
	return da.set("staticSchema", value)
}

// Style 组件样式
func (da dialogAction) Style(value any) dialogAction {
	return da.set("style", value)
}

// Target 可以指定让谁来触发这个动作。
func (da dialogAction) Target(value string) dialogAction {
	return da.set("target", value)
}

// TestIdBuilder
func (da dialogAction) TestIdBuilder(value string) dialogAction {
	return da.set("testIdBuilder", value)
}

// Testid
func (da dialogAction) Testid(value string) dialogAction {
	return da.set("testid", value)
}

// Tooltip
func (da dialogAction) Tooltip(value string) dialogAction {
	return da.set("tooltip", value)
}

// TooltipPlacement 可选值: top | right | bottom | left
func (da dialogAction) TooltipPlacement(value string) dialogAction {
	return da.set("tooltipPlacement", value)
}

// Type 指定按钮类型，支持 button、submit或者reset三种类型。 可选值: button | submit | reset
func (da dialogAction) Type(value string) dialogAction {
	return da.set("type", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (da dialogAction) UseMobileUI(value bool) dialogAction {
	return da.set("useMobileUI", value)
}

// Visible 是否显示
func (da dialogAction) Visible(value bool) dialogAction {
	return da.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (da dialogAction) VisibleOn(value string) dialogAction {
	return da.set("visibleOn", value)
}
