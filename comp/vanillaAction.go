package comp

// VanillaAction 代表动作配置。
type VanillaAction Schema

// NewVanillaAction 创建一个新的 VanillaAction 实例。
func NewVanillaAction() VanillaAction {
	return VanillaAction{}.set("type", "button")
}

func (v VanillaAction) set(key string, value interface{}) VanillaAction {
	v[key] = value
	return v
}

// ActionType 设置动作类型。
func (v VanillaAction) ActionType(value string) VanillaAction {
	return v.set("actionType", value)
}

// ActiveClassName 激活状态时的类名
func (v VanillaAction) ActiveClassName(value string) VanillaAction {
	return v.set("activeClassName", value)
}

// ActiveLevel 激活状态时的样式
func (v VanillaAction) ActiveLevel(value string) VanillaAction {
	return v.set("activeLevel", value)
}

// Badge 角标 (Badge 角标。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/badge)
func (v VanillaAction) Badge(value string) VanillaAction {
	return v.set("badge", value)
}

// Block 是否为块状展示，默认为内联。
func (v VanillaAction) Block(value bool) VanillaAction {
	return v.set("block", value)
}

// Body 子内容 (子内容)
func (v VanillaAction) Body(value ...interface{}) VanillaAction {
	return v.set("body", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (v VanillaAction) ClassName(value string) VanillaAction {
	return v.set("className", value)
}

// Close 如果按钮在弹框中，可以配置这个动作完成后是否关闭弹窗，或者指定关闭目标弹框。
func (v VanillaAction) Close(value string) VanillaAction {
	return v.set("close", value)
}

// ConfirmText 提示文字，配置了操作前会要求用户确认。
func (v VanillaAction) ConfirmText(value string) VanillaAction {
	return v.set("confirmText", value)
}

// ConfirmTitle 确认弹窗标题
func (v VanillaAction) ConfirmTitle(value string) VanillaAction {
	return v.set("confirmTitle", value)
}

// CountDown 点击后的禁止倒计时（秒）
func (v VanillaAction) CountDown(value string) VanillaAction {
	return v.set("countDown", value)
}

// CountDownTpl 倒计时文字自定义
func (v VanillaAction) CountDownTpl(value string) VanillaAction {
	return v.set("countDownTpl", value)
}

// Disabled 是否禁用
func (v VanillaAction) Disabled(value bool) VanillaAction {
	return v.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (v VanillaAction) DisabledOn(value string) VanillaAction {
	return v.set("disabledOn", value)
}

// DisabledTip 禁用时的文案提示。
func (v VanillaAction) DisabledTip(value string) VanillaAction {
	return v.set("disabledTip", value)
}

// DownloadFileName
func (v VanillaAction) DownloadFileName(value string) VanillaAction {
	return v.set("downloadFileName", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (v VanillaAction) EditorSetting(value string) VanillaAction {
	return v.set("editorSetting", value)
}

// Hidden 是否隐藏
func (v VanillaAction) Hidden(value bool) VanillaAction {
	return v.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (v VanillaAction) HiddenOn(value string) VanillaAction {
	return v.set("hiddenOn", value)
}

// HotKey 键盘快捷键
func (v VanillaAction) HotKey(value string) VanillaAction {
	return v.set("hotKey", value)
}

// Icon 按钮图标， iconfont 的类名 (iconfont 里面的类名。)
func (v VanillaAction) Icon(value string) VanillaAction {
	return v.set("icon", value)
}

// IconClassName icon 上的css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (v VanillaAction) IconClassName(value string) VanillaAction {
	return v.set("iconClassName", value)
}

// Id 主要用于用户行为跟踪里区分是哪个按钮
func (v VanillaAction) Id(value string) VanillaAction {
	return v.set("id", value)
}

// Label 按钮文字
func (v VanillaAction) Label(value string) VanillaAction {
	return v.set("label", value)
}

// Level 按钮样式 可选值: info | success | warning | danger | link | primary | dark | light | secondary
func (v VanillaAction) Level(value string) VanillaAction {
	return v.set("level", value)
}

// LoadingClassName loading 上的css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (v VanillaAction) LoadingClassName(value string) VanillaAction {
	return v.set("loadingClassName", value)
}

// LoadingOn 是否显示loading效果
func (v VanillaAction) LoadingOn(value string) VanillaAction {
	return v.set("loadingOn", value)
}

// MergeData 是否将弹框中数据 merge 到父级作用域。
func (v VanillaAction) MergeData(value bool) VanillaAction {
	return v.set("mergeData", value)
}

// OnClick 自定义事件处理函数
func (v VanillaAction) OnClick(value string) VanillaAction {
	return v.set("onClick", value)
}

// OnEvent 事件动作配置
func (v VanillaAction) OnEvent(value string) VanillaAction {
	return v.set("onEvent", value)
}

// Primary
func (v VanillaAction) Primary(value bool) VanillaAction {
	return v.set("primary", value)
}

// RequireSelected 当按钮时批量操作按钮时，默认必须有勾选元素才能可点击，如果此属性配置成 false，则没有点选成员也能点击。
func (v VanillaAction) RequireSelected(value bool) VanillaAction {
	return v.set("requireSelected", value)
}

// Required 如果按钮在form中，配置此属性会要求用户把指定的字段通过验证后才会触发行为。
func (v VanillaAction) Required(value string) VanillaAction {
	return v.set("required", value)
}

// RightIcon 右侧按钮图标， iconfont 的类名 (iconfont 里面的类名。)
func (v VanillaAction) RightIcon(value string) VanillaAction {
	return v.set("rightIcon", value)
}

// RightIconClassName 右侧 icon 上的 css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (v VanillaAction) RightIconClassName(value string) VanillaAction {
	return v.set("rightIconClassName", value)
}

// Size 按钮大小 可选值: xs | sm | md | lg
func (v VanillaAction) Size(value string) VanillaAction {
	return v.set("size", value)
}

// Static 是否静态展示
func (v VanillaAction) Static(value bool) VanillaAction {
	return v.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (v VanillaAction) StaticClassName(value string) VanillaAction {
	return v.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (v VanillaAction) StaticInputClassName(value string) VanillaAction {
	return v.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (v VanillaAction) StaticLabelClassName(value string) VanillaAction {
	return v.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (v VanillaAction) StaticOn(value string) VanillaAction {
	return v.set("staticOn", value)
}

// StaticPlaceholder 静态展示为空时的占位符。
func (v VanillaAction) StaticPlaceholder(value string) VanillaAction {
	return v.set("staticPlaceholder", value)
}

// StaticSchema 静态展示模式的 schema
func (v VanillaAction) StaticSchema(value string) VanillaAction {
	return v.set("staticSchema", value)
}

// Style 组件样式
func (v VanillaAction) Style(value string) VanillaAction {
	return v.set("style", value)
}

// Target 谁能触发这个动作。
func (v VanillaAction) Target(value string) VanillaAction {
	return v.set("target", value)
}

// TestIdBuilder 设置测试ID生成函数
func (v VanillaAction) TestIdBuilder(value string) VanillaAction {
	return v.set("testIdBuilder", value)
}

// Testid 测试ID
func (v VanillaAction) Testid(value string) VanillaAction {
	return v.set("testid", value)
}

// Tooltip 提示文字
func (v VanillaAction) Tooltip(value string) VanillaAction {
	return v.set("tooltip", value)
}

// TooltipPlacement 提示文字的显示位置 可选值: top | right | bottom | left
func (v VanillaAction) TooltipPlacement(value string) VanillaAction {
	return v.set("tooltipPlacement", value)
}

// UseMobileUI 是否禁用移动端样式
func (v VanillaAction) UseMobileUI(value bool) VanillaAction {
	return v.set("useMobileUI", value)
}

// Visible 是否可见
func (v VanillaAction) Visible(value bool) VanillaAction {
	return v.set("visible", value)
}

// VisibleOn 是否可见表达式 (表达式，语法 `data.xxx > 5`。)
func (v VanillaAction) VisibleOn(value string) VanillaAction {
	return v.set("visibleOn", value)
}
