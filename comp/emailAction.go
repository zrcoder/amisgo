package comp

// EmailAction 表示电子邮件动作按钮
type EmailAction Schema

// NewEmailAction 创建一个新的 EmailAction 实例，并设置默认的 type 和 actionType
func NewEmailAction() EmailAction {
	return make(EmailAction).set("type", "button").set("actionType", "email")
}

func (ea EmailAction) set(key string, value interface{}) EmailAction {
	ea[key] = value
	return ea
}

// ActionType 指定为打开邮箱行为
func (ea EmailAction) ActionType(value string) EmailAction {
	return ea.set("actionType", value)
}

// ActiveClassName 激活状态时的类名
func (ea EmailAction) ActiveClassName(value string) EmailAction {
	return ea.set("activeClassName", value)
}

// ActiveLevel 激活状态时的样式
func (ea EmailAction) ActiveLevel(value string) EmailAction {
	return ea.set("activeLevel", value)
}

// Badge 角标 (Badge 角标。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/badge)
func (ea EmailAction) Badge(value string) EmailAction {
	return ea.set("badge", value)
}

// Bcc 匿名抄送邮箱
func (ea EmailAction) Bcc(value string) EmailAction {
	return ea.set("bcc", value)
}

// Block 是否为块状展示，默认为内联。
func (ea EmailAction) Block(value bool) EmailAction {
	return ea.set("block", value)
}

// Body 邮件正文
func (ea EmailAction) Body(value ...interface{}) EmailAction {
	return ea.set("body", value)
}

// Cc 抄送邮箱
func (ea EmailAction) Cc(value string) EmailAction {
	return ea.set("cc", value)
}

// ClassName 容器 css 类名
func (ea EmailAction) ClassName(value string) EmailAction {
	return ea.set("className", value)
}

// Close 如果按钮在弹框中，可以配置这个动作完成后是否关闭弹窗，或者指定关闭目标弹框。
func (ea EmailAction) Close(value string) EmailAction {
	return ea.set("close", value)
}

// ConfirmText 提示文字，配置了操作前会要求用户确认。
func (ea EmailAction) ConfirmText(value string) EmailAction {
	return ea.set("confirmText", value)
}

// ConfirmTitle 确认弹窗标题
func (ea EmailAction) ConfirmTitle(value string) EmailAction {
	return ea.set("confirmTitle", value)
}

// CountDown 点击后的禁止倒计时（秒）
func (ea EmailAction) CountDown(value string) EmailAction {
	return ea.set("countDown", value)
}

// CountDownTpl 倒计时文字自定义
func (ea EmailAction) CountDownTpl(value string) EmailAction {
	return ea.set("countDownTpl", value)
}

// Disabled 是否禁用
func (ea EmailAction) Disabled(value bool) EmailAction {
	return ea.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (ea EmailAction) DisabledOn(value string) EmailAction {
	return ea.set("disabledOn", value)
}

// DisabledTip 禁用时的文案提示。
func (ea EmailAction) DisabledTip(value string) EmailAction {
	return ea.set("disabledTip", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (ea EmailAction) EditorSetting(value string) EmailAction {
	return ea.set("editorSetting", value)
}

// Hidden 是否隐藏
func (ea EmailAction) Hidden(value bool) EmailAction {
	return ea.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (ea EmailAction) HiddenOn(value string) EmailAction {
	return ea.set("hiddenOn", value)
}

// HotKey 键盘快捷键
func (ea EmailAction) HotKey(value string) EmailAction {
	return ea.set("hotKey", value)
}

// Icon 按钮图标， iconfont 的类名
func (ea EmailAction) Icon(value string) EmailAction {
	return ea.set("icon", value)
}

// IconClassName icon 上的 css 类名
func (ea EmailAction) IconClassName(value string) EmailAction {
	return ea.set("iconClassName", value)
}

// ID 主要用于用户行为跟踪里区分是哪个按钮
func (ea EmailAction) ID(value string) EmailAction {
	return ea.set("id", value)
}

// Label 按钮文字
func (ea EmailAction) Label(value string) EmailAction {
	return ea.set("label", value)
}

// Level 按钮样式
func (ea EmailAction) Level(value string) EmailAction {
	return ea.set("level", value)
}

// LoadingClassName loading 上的 css 类名
func (ea EmailAction) LoadingClassName(value string) EmailAction {
	return ea.set("loadingClassName", value)
}

// LoadingOn 是否显示 loading 效果
func (ea EmailAction) LoadingOn(value string) EmailAction {
	return ea.set("loadingOn", value)
}

// MergeData 是否将弹框中数据 merge 到父级作用域
func (ea EmailAction) MergeData(value bool) EmailAction {
	return ea.set("mergeData", value)
}

// OnClick 自定义事件处理函数
func (ea EmailAction) OnClick(value string) EmailAction {
	return ea.set("onClick", value)
}

// OnEvent 事件动作配置
func (ea EmailAction) OnEvent(value string) EmailAction {
	return ea.set("onEvent", value)
}

// Primary 设置为主要按钮
func (ea EmailAction) Primary(value bool) EmailAction {
	return ea.set("primary", value)
}

// RequireSelected 当按钮时批量操作按钮时，默认必须有勾选元素才能可点击
func (ea EmailAction) RequireSelected(value bool) EmailAction {
	return ea.set("requireSelected", value)
}

// Required 如果按钮在 form 中，配置此属性会要求用户把指定的字段通过验证后才会触发行为。
func (ea EmailAction) Required(value string) EmailAction {
	return ea.set("required", value)
}

// RightIcon 右侧按钮图标， iconfont 的类名
func (ea EmailAction) RightIcon(value string) EmailAction {
	return ea.set("rightIcon", value)
}

// RightIconClassName 右侧 icon 上的 css 类名
func (ea EmailAction) RightIconClassName(value string) EmailAction {
	return ea.set("rightIconClassName", value)
}

// Size 按钮大小
func (ea EmailAction) Size(value string) EmailAction {
	return ea.set("size", value)
}

// Static 是否静态展示
func (ea EmailAction) Static(value bool) EmailAction {
	return ea.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (ea EmailAction) StaticClassName(value string) EmailAction {
	return ea.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项 Value 类名
func (ea EmailAction) StaticInputClassName(value string) EmailAction {
	return ea.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项 Label 类名
func (ea EmailAction) StaticLabelClassName(value string) EmailAction {
	return ea.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (ea EmailAction) StaticOn(value string) EmailAction {
	return ea.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (ea EmailAction) StaticPlaceholder(value string) EmailAction {
	return ea.set("staticPlaceholder", value)
}

// StaticSchema 静态展示 schema
func (ea EmailAction) StaticSchema(value string) EmailAction {
	return ea.set("staticSchema", value)
}

// Style 组件样式
func (ea EmailAction) Style(value string) EmailAction {
	return ea.set("style", value)
}

// Subject 邮件主题
func (ea EmailAction) Subject(value string) EmailAction {
	return ea.set("subject", value)
}

// Target 可以指定让谁来触发这个动作
func (ea EmailAction) Target(value string) EmailAction {
	return ea.set("target", value)
}

// TestIdBuilder 测试 ID 构建器
func (ea EmailAction) TestIdBuilder(value string) EmailAction {
	return ea.set("testIdBuilder", value)
}

// Testid 测试 ID
func (ea EmailAction) Testid(value string) EmailAction {
	return ea.set("testid", value)
}

// To 收件人邮箱
func (ea EmailAction) To(value string) EmailAction {
	return ea.set("to", value)
}

// Tooltip 提示文字
func (ea EmailAction) Tooltip(value string) EmailAction {
	return ea.set("tooltip", value)
}

// TooltipPlacement 提示文字位置
func (ea EmailAction) TooltipPlacement(value string) EmailAction {
	return ea.set("tooltipPlacement", value)
}

// Type 指定按钮类型
func (ea EmailAction) Type(value string) EmailAction {
	return ea.set("type", value)
}

// UseMobileUI 组件级别关闭移动端样式
func (ea EmailAction) UseMobileUI(value bool) EmailAction {
	return ea.set("useMobileUI", value)
}

// Visible 是否显示
func (ea EmailAction) Visible(value bool) EmailAction {
	return ea.set("visible", value)
}

// VisibleOn 是否显示表达式
func (ea EmailAction) VisibleOn(value string) EmailAction {
	return ea.set("visibleOn", value)
}
