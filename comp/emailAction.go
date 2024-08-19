package comp

// emailAction 表示电子邮件动作按钮
type emailAction schema

// EmailAction 创建一个新的 EmailAction 实例，并设置默认的 type 和 actionType
func EmailAction() emailAction {
	return make(emailAction).set("type", "button").set("actionType", "email")
}

func (ea emailAction) set(key string, value any) emailAction {
	ea[key] = value
	return ea
}

// ActiveClassName 激活状态时的类名
func (ea emailAction) ActiveClassName(value string) emailAction {
	return ea.set("activeClassName", value)
}

// ActiveLevel 激活状态时的样式
func (ea emailAction) ActiveLevel(value string) emailAction {
	return ea.set("activeLevel", value)
}

// Badge 角标 (Badge 角标。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/badge)
func (ea emailAction) Badge(value string) emailAction {
	return ea.set("badge", value)
}

// Bcc 匿名抄送邮箱
func (ea emailAction) Bcc(value string) emailAction {
	return ea.set("bcc", value)
}

// Block 是否为块状展示，默认为内联。
func (ea emailAction) Block(value bool) emailAction {
	return ea.set("block", value)
}

// Body 邮件正文
func (ea emailAction) Body(value ...any) emailAction {
	return ea.set("body", value)
}

// Cc 抄送邮箱
func (ea emailAction) Cc(value string) emailAction {
	return ea.set("cc", value)
}

// ClassName 容器 css 类名
func (ea emailAction) ClassName(value string) emailAction {
	return ea.set("className", value)
}

// Close 如果按钮在弹框中，可以配置这个动作完成后是否关闭弹窗，或者指定关闭目标弹框。
func (ea emailAction) Close(value string) emailAction {
	return ea.set("close", value)
}

// ConfirmText 提示文字，配置了操作前会要求用户确认。
func (ea emailAction) ConfirmText(value string) emailAction {
	return ea.set("confirmText", value)
}

// ConfirmTitle 确认弹窗标题
func (ea emailAction) ConfirmTitle(value any) emailAction {
	return ea.set("confirmTitle", value)
}

// CountDown 点击后的禁止倒计时（秒）
func (ea emailAction) CountDown(value string) emailAction {
	return ea.set("countDown", value)
}

// CountDownTpl 倒计时文字自定义
func (ea emailAction) CountDownTpl(value string) emailAction {
	return ea.set("countDownTpl", value)
}

// Disabled 是否禁用
func (ea emailAction) Disabled(value bool) emailAction {
	return ea.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (ea emailAction) DisabledOn(value string) emailAction {
	return ea.set("disabledOn", value)
}

// DisabledTip 禁用时的文案提示。
func (ea emailAction) DisabledTip(value string) emailAction {
	return ea.set("disabledTip", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (ea emailAction) EditorSetting(value string) emailAction {
	return ea.set("editorSetting", value)
}

// Hidden 是否隐藏
func (ea emailAction) Hidden(value bool) emailAction {
	return ea.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (ea emailAction) HiddenOn(value string) emailAction {
	return ea.set("hiddenOn", value)
}

// HotKey 键盘快捷键
func (ea emailAction) HotKey(value string) emailAction {
	return ea.set("hotKey", value)
}

// Icon 按钮图标， iconfont 的类名
func (ea emailAction) Icon(value string) emailAction {
	return ea.set("icon", value)
}

// IconClassName icon 上的 css 类名
func (ea emailAction) IconClassName(value string) emailAction {
	return ea.set("iconClassName", value)
}

// ID 主要用于用户行为跟踪里区分是哪个按钮
func (ea emailAction) ID(value string) emailAction {
	return ea.set("id", value)
}

// Label 按钮文字
func (ea emailAction) Label(value string) emailAction {
	return ea.set("label", value)
}

// Level 按钮样式
func (ea emailAction) Level(value string) emailAction {
	return ea.set("level", value)
}

// LoadingClassName loading 上的 css 类名
func (ea emailAction) LoadingClassName(value string) emailAction {
	return ea.set("loadingClassName", value)
}

// LoadingOn 是否显示 loading 效果
func (ea emailAction) LoadingOn(value string) emailAction {
	return ea.set("loadingOn", value)
}

// MergeData 是否将弹框中数据 merge 到父级作用域
func (ea emailAction) MergeData(value bool) emailAction {
	return ea.set("mergeData", value)
}

// OnClick 自定义事件处理函数
func (ea emailAction) OnClick(value string) emailAction {
	return ea.set("onClick", value)
}

// OnEvent 事件动作配置
func (ea emailAction) OnEvent(value any) emailAction {
	return ea.set("onEvent", value)
}

// Primary 设置为主要按钮
func (ea emailAction) Primary(value bool) emailAction {
	return ea.set("primary", value)
}

// RequireSelected 当按钮时批量操作按钮时，默认必须有勾选元素才能可点击
func (ea emailAction) RequireSelected(value bool) emailAction {
	return ea.set("requireSelected", value)
}

// Required 如果按钮在 form 中，配置此属性会要求用户把指定的字段通过验证后才会触发行为。
func (ea emailAction) Required(value string) emailAction {
	return ea.set("required", value)
}

// RightIcon 右侧按钮图标， iconfont 的类名
func (ea emailAction) RightIcon(value string) emailAction {
	return ea.set("rightIcon", value)
}

// RightIconClassName 右侧 icon 上的 css 类名
func (ea emailAction) RightIconClassName(value string) emailAction {
	return ea.set("rightIconClassName", value)
}

// Size 按钮大小
func (ea emailAction) Size(value string) emailAction {
	return ea.set("size", value)
}

// Static 是否静态展示
func (ea emailAction) Static(value bool) emailAction {
	return ea.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (ea emailAction) StaticClassName(value string) emailAction {
	return ea.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项 Value 类名
func (ea emailAction) StaticInputClassName(value string) emailAction {
	return ea.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项 Label 类名
func (ea emailAction) StaticLabelClassName(value string) emailAction {
	return ea.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (ea emailAction) StaticOn(value string) emailAction {
	return ea.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (ea emailAction) StaticPlaceholder(value string) emailAction {
	return ea.set("staticPlaceholder", value)
}

// StaticSchema 静态展示 schema
func (ea emailAction) StaticSchema(value string) emailAction {
	return ea.set("staticSchema", value)
}

// Style 组件样式
func (ea emailAction) Style(value any) emailAction {
	return ea.set("style", value)
}

// Subject 邮件主题
func (ea emailAction) Subject(value string) emailAction {
	return ea.set("subject", value)
}

// Target 可以指定让谁来触发这个动作
func (ea emailAction) Target(value string) emailAction {
	return ea.set("target", value)
}

// TestIdBuilder 测试 ID 构建器
func (ea emailAction) TestIdBuilder(value string) emailAction {
	return ea.set("testIdBuilder", value)
}

// Testid 测试 ID
func (ea emailAction) Testid(value string) emailAction {
	return ea.set("testid", value)
}

// To 收件人邮箱
func (ea emailAction) To(value string) emailAction {
	return ea.set("to", value)
}

// Tooltip 提示文字
func (ea emailAction) Tooltip(value string) emailAction {
	return ea.set("tooltip", value)
}

// TooltipPlacement 提示文字位置
func (ea emailAction) TooltipPlacement(value string) emailAction {
	return ea.set("tooltipPlacement", value)
}

// UseMobileUI 组件级别关闭移动端样式
func (ea emailAction) UseMobileUI(value bool) emailAction {
	return ea.set("useMobileUI", value)
}

// Visible 是否显示
func (ea emailAction) Visible(value bool) emailAction {
	return ea.set("visible", value)
}

// VisibleOn 是否显示表达式
func (ea emailAction) VisibleOn(value string) emailAction {
	return ea.set("visibleOn", value)
}
