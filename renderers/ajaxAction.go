package renderers

// AjaxAction 表示一个 Ajax 行为按钮
type AjaxAction struct {
	*BaseRenderer
}

// NewAjaxAction 创建一个新的 AjaxAction 实例
func NewAjaxAction() *AjaxAction {
	a := &AjaxAction{
		BaseRenderer: NewBaseRenderer(),
	}
	a.set("type", "button")
	a.set("actionType", "ajax")
	return a
}

// ActionType 指定为发送 ajax 的行为。
func (a *AjaxAction) ActionType(value string) *AjaxAction {
	a.set("actionType", value)
	return a
}

// ActiveClassName 激活状态时的类名
func (a *AjaxAction) ActiveClassName(value string) *AjaxAction {
	a.set("activeClassName", value)
	return a
}

// ActiveLevel 激活状态时的样式
func (a *AjaxAction) ActiveLevel(value string) *AjaxAction {
	a.set("activeLevel", value)
	return a
}

// Api 配置 ajax 发送地址
func (a *AjaxAction) Api(value string) *AjaxAction {
	a.set("api", value)
	return a
}

// Badge 角标 (Badge 角标。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/badge)
func (a *AjaxAction) Badge(value string) *AjaxAction {
	a.set("badge", value)
	return a
}

// Block 是否为块状展示，默认为内联。
func (a *AjaxAction) Block(value bool) *AjaxAction {
	a.set("block", value)
	return a
}

// Body 子内容
func (a *AjaxAction) Body(value string) *AjaxAction {
	a.set("body", value)
	return a
}

// ClassName 容器 css 类名
func (a *AjaxAction) ClassName(value string) *AjaxAction {
	a.set("className", value)
	return a
}

// Close 如果按钮在弹框中，可以配置这个动作完成后是否关闭弹窗，或者指定关闭目标弹框。
func (a *AjaxAction) Close(value string) *AjaxAction {
	a.set("close", value)
	return a
}

// ConfirmText 提示文字，配置了操作前会要求用户确认。
func (a *AjaxAction) ConfirmText(value string) *AjaxAction {
	a.set("confirmText", value)
	return a
}

// ConfirmTitle 确认弹窗标题
func (a *AjaxAction) ConfirmTitle(value string) *AjaxAction {
	a.set("confirmTitle", value)
	return a
}

// CountDown 点击后的禁止倒计时（秒）
func (a *AjaxAction) CountDown(value string) *AjaxAction {
	a.set("countDown", value)
	return a
}

// CountDownTpl 倒计时文字自定义
func (a *AjaxAction) CountDownTpl(value string) *AjaxAction {
	a.set("countDownTpl", value)
	return a
}

// Disabled 是否禁用
func (a *AjaxAction) Disabled(value bool) *AjaxAction {
	a.set("disabled", value)
	return a
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (a *AjaxAction) DisabledOn(value string) *AjaxAction {
	a.set("disabledOn", value)
	return a
}

// DisabledTip 禁用时的文案提示。
func (a *AjaxAction) DisabledTip(value string) *AjaxAction {
	a.set("disabledTip", value)
	return a
}

// EditorSetting 编辑器配置，运行时可以忽略
func (a *AjaxAction) EditorSetting(value string) *AjaxAction {
	a.set("editorSetting", value)
	return a
}

// Feedback 自定义反馈内容
func (a *AjaxAction) Feedback(value string) *AjaxAction {
	a.set("feedback", value)
	return a
}

// Hidden 是否隐藏
func (a *AjaxAction) Hidden(value bool) *AjaxAction {
	a.set("hidden", value)
	return a
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (a *AjaxAction) HiddenOn(value string) *AjaxAction {
	a.set("hiddenOn", value)
	return a
}

// HotKey 键盘快捷键
func (a *AjaxAction) HotKey(value string) *AjaxAction {
	a.set("hotKey", value)
	return a
}

// Icon 按钮图标， iconfont 的类名
func (a *AjaxAction) Icon(value string) *AjaxAction {
	a.set("icon", value)
	return a
}

// IconClassName icon 上的 css 类名
func (a *AjaxAction) IconClassName(value string) *AjaxAction {
	a.set("iconClassName", value)
	return a
}

// Id 主要用于用户行为跟踪里区分是哪个按钮
func (a *AjaxAction) Id(value string) *AjaxAction {
	a.set("id", value)
	return a
}

// IgnoreConfirm 是否忽略确认
func (a *AjaxAction) IgnoreConfirm(value bool) *AjaxAction {
	a.set("ignoreConfirm", value)
	return a
}

// IsolateScope 是否开启请求隔离
func (a *AjaxAction) IsolateScope(value bool) *AjaxAction {
	a.set("isolateScope", value)
	return a
}

// Label 按钮文字
func (a *AjaxAction) Label(value string) *AjaxAction {
	a.set("label", value)
	return a
}

// Level 按钮样式
func (a *AjaxAction) Level(value string) *AjaxAction {
	a.set("level", value)
	return a
}

// LoadingClassName loading 上的 css 类名
func (a *AjaxAction) LoadingClassName(value string) *AjaxAction {
	a.set("loadingClassName", value)
	return a
}

// LoadingOn 是否显示 loading 效果
func (a *AjaxAction) LoadingOn(value string) *AjaxAction {
	a.set("loadingOn", value)
	return a
}

// MergeData 是否将弹框中数据 merge 到父级作用域
func (a *AjaxAction) MergeData(value bool) *AjaxAction {
	a.set("mergeData", value)
	return a
}

// OnClick 自定义事件处理函数
func (a *AjaxAction) OnClick(value string) *AjaxAction {
	a.set("onClick", value)
	return a
}

// OnEvent 事件动作配置
func (a *AjaxAction) OnEvent(value string) *AjaxAction {
	a.set("onEvent", value)
	return a
}

// Primary 是否主要按钮
func (a *AjaxAction) Primary(value bool) *AjaxAction {
	a.set("primary", value)
	return a
}

// Redirect 配置重定向
func (a *AjaxAction) Redirect(value string) *AjaxAction {
	a.set("redirect", value)
	return a
}

// Reload 配置刷新动作
func (a *AjaxAction) Reload(value string) *AjaxAction {
	a.set("reload", value)
	return a
}

// RequireSelected 是否要求选中
func (a *AjaxAction) RequireSelected(value bool) *AjaxAction {
	a.set("requireSelected", value)
	return a
}

// Required 配置字段验证
func (a *AjaxAction) Required(value string) *AjaxAction {
	a.set("required", value)
	return a
}

// RightIcon 右侧按钮图标
func (a *AjaxAction) RightIcon(value string) *AjaxAction {
	a.set("rightIcon", value)
	return a
}

// RightIconClassName 右侧 icon 上的 css 类名
func (a *AjaxAction) RightIconClassName(value string) *AjaxAction {
	a.set("rightIconClassName", value)
	return a
}

// Size 按钮大小
func (a *AjaxAction) Size(value string) *AjaxAction {
	a.set("size", value)
	return a
}

// Static 是否静态展示
func (a *AjaxAction) Static(value bool) *AjaxAction {
	a.set("static", value)
	return a
}

// StaticClassName 静态展示表单项类名
func (a *AjaxAction) StaticClassName(value string) *AjaxAction {
	a.set("staticClassName", value)
	return a
}

// StaticInputClassName 静态展示表单项Value类名
func (a *AjaxAction) StaticInputClassName(value string) *AjaxAction {
	a.set("staticInputClassName", value)
	return a
}

// StaticLabelClassName 静态展示表单项Label类名
func (a *AjaxAction) StaticLabelClassName(value string) *AjaxAction {
	a.set("staticLabelClassName", value)
	return a
}

// StaticOn 是否静态展示表达式
func (a *AjaxAction) StaticOn(value string) *AjaxAction {
	a.set("staticOn", value)
	return a
}

// StaticPlaceholder 静态展示空值占位
func (a *AjaxAction) StaticPlaceholder(value string) *AjaxAction {
	a.set("staticPlaceholder", value)
	return a
}

// StaticSchema 静态展示的 schema
func (a *AjaxAction) StaticSchema(value string) *AjaxAction {
	a.set("staticSchema", value)
	return a
}

// Style 组件样式
func (a *AjaxAction) Style(value string) *AjaxAction {
	a.set("style", value)
	return a
}

// Target 可以指定让谁来触发这个动作
func (a *AjaxAction) Target(value string) *AjaxAction {
	a.set("target", value)
	return a
}

// TestIdBuilder 测试 ID 构建器
func (a *AjaxAction) TestIdBuilder(value string) *AjaxAction {
	a.set("testIdBuilder", value)
	return a
}

// Testid 测试 ID
func (a *AjaxAction) Testid(value string) *AjaxAction {
	a.set("testid", value)
	return a
}

// Tooltip 提示文本
func (a *AjaxAction) Tooltip(value string) *AjaxAction {
	a.set("tooltip", value)
	return a
}

// TooltipPlacement 提示文本位置
func (a *AjaxAction) TooltipPlacement(value string) *AjaxAction {
	a.set("tooltipPlacement", value)
	return a
}

// 指定按钮类型，支持 button、submit或者reset三种类型。 可选值: button | submit | reset
func (a *AjaxAction) Type(value string) *AjaxAction {
	a.set("type", value)
	return a
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (a *AjaxAction) UseMobileUI(value bool) *AjaxAction {
	a.set("useMobileUI", value)
	return a
}

// Visible 是否显示
func (a *AjaxAction) Visible(value bool) *AjaxAction {
	a.set("visible", value)
	return a
}

// VisibleOn 是否显示表达式
func (a *AjaxAction) VisibleOn(value string) *AjaxAction {
	a.set("visibleOn", value)
	return a
}
