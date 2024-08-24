package comp

// vanillaAction 代表动作配置。
type vanillaAction schema

// VanillaAction 创建一个新的 VanillaAction 实例。
func VanillaAction() vanillaAction {
	return vanillaAction{}.set("type", "button")
}

func (v vanillaAction) set(key string, value any) vanillaAction {
	v[key] = value
	return v
}

// ActionType 设置动作类型。
func (v vanillaAction) ActionType(value string) vanillaAction {
	return v.set("actionType", value)
}

// ActiveClassName 激活状态时的类名
func (v vanillaAction) ActiveClassName(value string) vanillaAction {
	return v.set("activeClassName", value)
}

// ActiveLevel 激活状态时的样式
func (v vanillaAction) ActiveLevel(value string) vanillaAction {
	return v.set("activeLevel", value)
}

// Badge 角标 (Badge 角标。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/badge)
func (v vanillaAction) Badge(value string) vanillaAction {
	return v.set("badge", value)
}

// Block 是否为块状展示，默认为内联。
func (v vanillaAction) Block(value bool) vanillaAction {
	return v.set("block", value)
}

// Body 子内容 (子内容)
func (v vanillaAction) Body(value ...any) vanillaAction {
	return v.set("body", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (v vanillaAction) ClassName(value string) vanillaAction {
	return v.set("className", value)
}

// Close 如果按钮在弹框中，可以配置这个动作完成后是否关闭弹窗，或者指定关闭目标弹框。
func (v vanillaAction) Close(value string) vanillaAction {
	return v.set("close", value)
}

// ConfirmText 提示文字，配置了操作前会要求用户确认。
func (v vanillaAction) ConfirmText(value string) vanillaAction {
	return v.set("confirmText", value)
}

// ConfirmTitle 确认弹窗标题
func (v vanillaAction) ConfirmTitle(value any) vanillaAction {
	return v.set("confirmTitle", value)
}

// CountDown 点击后的禁止倒计时（秒）
func (v vanillaAction) CountDown(value string) vanillaAction {
	return v.set("countDown", value)
}

// CountDownTpl 倒计时文字自定义
func (v vanillaAction) CountDownTpl(value string) vanillaAction {
	return v.set("countDownTpl", value)
}

// Disabled 是否禁用
func (v vanillaAction) Disabled(value bool) vanillaAction {
	return v.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (v vanillaAction) DisabledOn(value string) vanillaAction {
	return v.set("disabledOn", value)
}

// DisabledTip 禁用时的文案提示。
func (v vanillaAction) DisabledTip(value string) vanillaAction {
	return v.set("disabledTip", value)
}

// DownloadFileName
func (v vanillaAction) DownloadFileName(value string) vanillaAction {
	return v.set("downloadFileName", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (v vanillaAction) EditorSetting(value string) vanillaAction {
	return v.set("editorSetting", value)
}

// Hidden 是否隐藏
func (v vanillaAction) Hidden(value bool) vanillaAction {
	return v.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (v vanillaAction) HiddenOn(value string) vanillaAction {
	return v.set("hiddenOn", value)
}

// HotKey 键盘快捷键
func (v vanillaAction) HotKey(value string) vanillaAction {
	return v.set("hotKey", value)
}

// Icon 按钮图标， iconfont 的类名 (iconfont 里面的类名。)
func (v vanillaAction) Icon(value string) vanillaAction {
	return v.set("icon", value)
}

// IconClassName icon 上的css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (v vanillaAction) IconClassName(value string) vanillaAction {
	return v.set("iconClassName", value)
}

// Id 主要用于用户行为跟踪里区分是哪个按钮
func (v vanillaAction) Id(value string) vanillaAction {
	return v.set("id", value)
}

// Label 按钮文字
func (v vanillaAction) Label(value string) vanillaAction {
	return v.set("label", value)
}

// Url 按钮链接
func (v vanillaAction) Url(value string) vanillaAction {
	return v.set("url", value)
}

// Level 按钮样式 可选值: info | success | warning | danger | link | primary | dark | light | secondary
func (v vanillaAction) Level(value string) vanillaAction {
	return v.set("level", value)
}

// LoadingClassName loading 上的css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (v vanillaAction) LoadingClassName(value string) vanillaAction {
	return v.set("loadingClassName", value)
}

// LoadingOn 是否显示loading效果
func (v vanillaAction) LoadingOn(value string) vanillaAction {
	return v.set("loadingOn", value)
}

// MergeData 是否将弹框中数据 merge 到父级作用域。
func (v vanillaAction) MergeData(value bool) vanillaAction {
	return v.set("mergeData", value)
}

// OnClick 自定义事件处理函数
func (v vanillaAction) OnClick(value string) vanillaAction {
	return v.set("onClick", value)
}

// OnEvent 事件动作配置
func (v vanillaAction) OnEvent(value any) vanillaAction {
	return v.set("onEvent", value)
}

// Primary
func (v vanillaAction) Primary(value bool) vanillaAction {
	return v.set("primary", value)
}

// RequireSelected 当按钮时批量操作按钮时，默认必须有勾选元素才能可点击，如果此属性配置成 false，则没有点选成员也能点击。
func (v vanillaAction) RequireSelected(value bool) vanillaAction {
	return v.set("requireSelected", value)
}

// Required 如果按钮在form中，配置此属性会要求用户把指定的字段通过验证后才会触发行为。
func (v vanillaAction) Required(value string) vanillaAction {
	return v.set("required", value)
}

// RightIcon 右侧按钮图标， iconfont 的类名 (iconfont 里面的类名。)
func (v vanillaAction) RightIcon(value string) vanillaAction {
	return v.set("rightIcon", value)
}

// RightIconClassName 右侧 icon 上的 css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (v vanillaAction) RightIconClassName(value string) vanillaAction {
	return v.set("rightIconClassName", value)
}

// Size 按钮大小 可选值: xs | sm | md | lg
func (v vanillaAction) Size(value string) vanillaAction {
	return v.set("size", value)
}

// Static 是否静态展示
func (v vanillaAction) Static(value bool) vanillaAction {
	return v.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (v vanillaAction) StaticClassName(value string) vanillaAction {
	return v.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (v vanillaAction) StaticInputClassName(value string) vanillaAction {
	return v.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (v vanillaAction) StaticLabelClassName(value string) vanillaAction {
	return v.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (v vanillaAction) StaticOn(value string) vanillaAction {
	return v.set("staticOn", value)
}

// StaticPlaceholder 静态展示为空时的占位符。
func (v vanillaAction) StaticPlaceholder(value string) vanillaAction {
	return v.set("staticPlaceholder", value)
}

// StaticSchema 静态展示模式的 schema
func (v vanillaAction) StaticSchema(value string) vanillaAction {
	return v.set("staticSchema", value)
}

// Style 组件样式
func (v vanillaAction) Style(value any) vanillaAction {
	return v.set("style", value)
}

// Target 谁能触发这个动作。
func (v vanillaAction) Target(value string) vanillaAction {
	return v.set("target", value)
}

// TestIdBuilder 设置测试ID生成函数
func (v vanillaAction) TestIdBuilder(value string) vanillaAction {
	return v.set("testIdBuilder", value)
}

// Testid 测试ID
func (v vanillaAction) Testid(value string) vanillaAction {
	return v.set("testid", value)
}

// Tooltip 提示文字
func (v vanillaAction) Tooltip(value string) vanillaAction {
	return v.set("tooltip", value)
}

// TooltipPlacement 提示文字的显示位置 可选值: top | right | bottom | left
func (v vanillaAction) TooltipPlacement(value string) vanillaAction {
	return v.set("tooltipPlacement", value)
}

// UseMobileUI 是否禁用移动端样式
func (v vanillaAction) UseMobileUI(value bool) vanillaAction {
	return v.set("useMobileUI", value)
}

// Visible 是否可见
func (v vanillaAction) Visible(value bool) vanillaAction {
	return v.set("visible", value)
}

// VisibleOn 是否可见表达式 (表达式，语法 `data.xxx > 5`。)
func (v vanillaAction) VisibleOn(value string) vanillaAction {
	return v.set("visibleOn", value)
}
