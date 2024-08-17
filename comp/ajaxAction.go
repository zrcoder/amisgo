package comp

// AjaxAction 表示一个 Ajax 行为按钮
type AjaxAction Schema

// NewAjaxAction 创建一个新的 AjaxAction 实例
func NewAjaxAction() AjaxAction {
	return make(AjaxAction).set("type", "button").set("actionType", "ajax")
}

// ActionType 指定为发送 ajax 的行为。
func (a AjaxAction) ActionType(value string) AjaxAction {
	return a.set("actionType", value)
}

// ActiveClassName 激活状态时的类名
func (a AjaxAction) ActiveClassName(value string) AjaxAction {
	return a.set("activeClassName", value)
}

// ActiveLevel 激活状态时的样式
func (a AjaxAction) ActiveLevel(value string) AjaxAction {
	return a.set("activeLevel", value)
}

// Api 配置 ajax 发送地址
func (a AjaxAction) Api(value string) AjaxAction {
	return a.set("api", value)
}

// Badge 角标 (Badge 角标。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/badge)
func (a AjaxAction) Badge(value string) AjaxAction {
	return a.set("badge", value)
}

// Block 是否为块状展示，默认为内联。
func (a AjaxAction) Block(value bool) AjaxAction {
	return a.set("block", value)
}

// Body 子内容
func (a AjaxAction) Body(value ...interface{}) AjaxAction {
	return a.set("body", value)
}

// ClassName 容器 css 类名
func (a AjaxAction) ClassName(value string) AjaxAction {
	return a.set("className", value)
}

// Close 如果按钮在弹框中，可以配置这个动作完成后是否关闭弹窗，或者指定关闭目标弹框。
func (a AjaxAction) Close(value string) AjaxAction {
	return a.set("close", value)
}

// ConfirmText 提示文字，配置了操作前会要求用户确认。
func (a AjaxAction) ConfirmText(value string) AjaxAction {
	return a.set("confirmText", value)
}

// ConfirmTitle 确认弹窗标题
func (a AjaxAction) ConfirmTitle(value string) AjaxAction {
	return a.set("confirmTitle", value)
}

// CountDown 点击后的禁止倒计时（秒）
func (a AjaxAction) CountDown(value string) AjaxAction {
	return a.set("countDown", value)
}

// CountDownTpl 倒计时文字自定义
func (a AjaxAction) CountDownTpl(value string) AjaxAction {
	return a.set("countDownTpl", value)
}

// Disabled 是否禁用
func (a AjaxAction) Disabled(value bool) AjaxAction {
	return a.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (a AjaxAction) DisabledOn(value string) AjaxAction {
	return a.set("disabledOn", value)
}

// DisabledTip 禁用时的文案提示。
func (a AjaxAction) DisabledTip(value string) AjaxAction {
	return a.set("disabledTip", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (a AjaxAction) EditorSetting(value string) AjaxAction {
	return a.set("editorSetting", value)
}

// Feedback 自定义反馈内容
func (a AjaxAction) Feedback(value string) AjaxAction {
	return a.set("feedback", value)
}

// Hidden 是否隐藏
func (a AjaxAction) Hidden(value bool) AjaxAction {
	return a.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (a AjaxAction) HiddenOn(value string) AjaxAction {
	return a.set("hiddenOn", value)
}

// HotKey 键盘快捷键
func (a AjaxAction) HotKey(value string) AjaxAction {
	return a.set("hotKey", value)
}

// Icon 按钮图标， iconfont 的类名
func (a AjaxAction) Icon(value string) AjaxAction {
	return a.set("icon", value)
}

// IconClassName icon 上的 css 类名
func (a AjaxAction) IconClassName(value string) AjaxAction {
	return a.set("iconClassName", value)
}

// Id 主要用于用户行为跟踪里区分是哪个按钮
func (a AjaxAction) Id(value string) AjaxAction {
	return a.set("id", value)
}

// IgnoreConfirm 是否忽略确认
func (a AjaxAction) IgnoreConfirm(value bool) AjaxAction {
	return a.set("ignoreConfirm", value)
}

// IsolateScope 是否开启请求隔离
func (a AjaxAction) IsolateScope(value bool) AjaxAction {
	return a.set("isolateScope", value)
}

// Label 按钮文字
func (a AjaxAction) Label(value string) AjaxAction {
	return a.set("label", value)
}

// Level 按钮样式
func (a AjaxAction) Level(value string) AjaxAction {
	return a.set("level", value)
}

// LoadingClassName loading 上的 css 类名
func (a AjaxAction) LoadingClassName(value string) AjaxAction {
	return a.set("loadingClassName", value)
}

// LoadingOn 是否显示 loading 效果
func (a AjaxAction) LoadingOn(value string) AjaxAction {
	return a.set("loadingOn", value)
}

// MergeData 是否将弹框中数据 merge 到父级作用域
func (a AjaxAction) MergeData(value bool) AjaxAction {
	return a.set("mergeData", value)
}

// OnClick 自定义事件处理函数
func (a AjaxAction) OnClick(value string) AjaxAction {
	return a.set("onClick", value)
}

// OnEvent 事件动作配置
func (a AjaxAction) OnEvent(value string) AjaxAction {
	return a.set("onEvent", value)
}

// Primary 是否主要按钮
func (a AjaxAction) Primary(value bool) AjaxAction {
	return a.set("primary", value)
}

// Redirect 配置重定向
func (a AjaxAction) Redirect(value string) AjaxAction {
	return a.set("redirect", value)
}

// Reload 配置刷新动作
func (a AjaxAction) Reload(value string) AjaxAction {
	return a.set("reload", value)
}

// RequireSelected 是否要求选中
func (a AjaxAction) RequireSelected(value bool) AjaxAction {
	return a.set("requireSelected", value)
}

// Required 配置字段验证
func (a AjaxAction) Required(value string) AjaxAction {
	return a.set("required", value)
}

// RightIcon 右侧按钮图标
func (a AjaxAction) RightIcon(value string) AjaxAction {
	return a.set("rightIcon", value)
}

// RightIconClassName 右侧 icon 上的 css 类名
func (a AjaxAction) RightIconClassName(value string) AjaxAction {
	return a.set("rightIconClassName", value)
}

// Size 按钮大小
func (a AjaxAction) Size(value string) AjaxAction {
	return a.set("size", value)
}

// Static 是否静态展示
func (a AjaxAction) Static(value bool) AjaxAction {
	return a.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (a AjaxAction) StaticClassName(value string) AjaxAction {
	return a.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (a AjaxAction) StaticInputClassName(value string) AjaxAction {
	return a.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (a AjaxAction) StaticLabelClassName(value string) AjaxAction {
	return a.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (a AjaxAction) StaticOn(value string) AjaxAction {
	return a.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (a AjaxAction) StaticPlaceholder(value string) AjaxAction {
	return a.set("staticPlaceholder", value)
}

// StaticSchema 静态展示的 schema
func (a AjaxAction) StaticSchema(value string) AjaxAction {
	return a.set("staticSchema", value)
}

// Style 组件样式
func (a AjaxAction) Style(value string) AjaxAction {
	return a.set("style", value)
}

// Target 可以指定让谁来触发这个动作
func (a AjaxAction) Target(value string) AjaxAction {
	return a.set("target", value)
}

// TestIdBuilder 测试 ID 构建器
func (a AjaxAction) TestIdBuilder(value string) AjaxAction {
	return a.set("testIdBuilder", value)
}

// Testid 测试 ID
func (a AjaxAction) Testid(value string) AjaxAction {
	return a.set("testid", value)
}

// Tooltip 提示文本
func (a AjaxAction) Tooltip(value string) AjaxAction {
	return a.set("tooltip", value)
}

// TooltipPlacement 提示文本位置
func (a AjaxAction) TooltipPlacement(value string) AjaxAction {
	return a.set("tooltipPlacement", value)
}

// 指定按钮类型，支持 button、submit或者reset三种类型。 可选值: button | submit | reset
func (a AjaxAction) Type(value string) AjaxAction {
	return a.set("type", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (a AjaxAction) UseMobileUI(value bool) AjaxAction {
	return a.set("useMobileUI", value)
}

// Visible 是否显示
func (a AjaxAction) Visible(value bool) AjaxAction {
	return a.set("visible", value)
}

// VisibleOn 是否显示表达式
func (a AjaxAction) VisibleOn(value string) AjaxAction {
	return a.set("visibleOn", value)
}

func (a AjaxAction) set(key string, value interface{}) AjaxAction {
	a[key] = value
	return a
}
