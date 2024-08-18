package comp

// ajaxAction 表示一个 Ajax 行为按钮
type ajaxAction schema

func AjaxAction() ajaxAction {
	return make(ajaxAction).set("type", "button").set("actionType", "ajax")
}

// ActiveClassName 激活状态时的类名
func (a ajaxAction) ActiveClassName(value string) ajaxAction {
	return a.set("activeClassName", value)
}

// ActiveLevel 激活状态时的样式
func (a ajaxAction) ActiveLevel(value string) ajaxAction {
	return a.set("activeLevel", value)
}

// Api 配置 ajax 发送地址
func (a ajaxAction) Api(value string) ajaxAction {
	return a.set("api", value)
}

// Badge 角标 (Badge 角标。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/badge)
func (a ajaxAction) Badge(value string) ajaxAction {
	return a.set("badge", value)
}

// Block 是否为块状展示，默认为内联。
func (a ajaxAction) Block(value bool) ajaxAction {
	return a.set("block", value)
}

// Body 子内容
func (a ajaxAction) Body(value ...interface{}) ajaxAction {
	return a.set("body", value)
}

// ClassName 容器 css 类名
func (a ajaxAction) ClassName(value string) ajaxAction {
	return a.set("className", value)
}

// Close 如果按钮在弹框中，可以配置这个动作完成后是否关闭弹窗，或者指定关闭目标弹框。
func (a ajaxAction) Close(value string) ajaxAction {
	return a.set("close", value)
}

// ConfirmText 提示文字，配置了操作前会要求用户确认。
func (a ajaxAction) ConfirmText(value string) ajaxAction {
	return a.set("confirmText", value)
}

// ConfirmTitle 确认弹窗标题
func (a ajaxAction) ConfirmTitle(value string) ajaxAction {
	return a.set("confirmTitle", value)
}

// CountDown 点击后的禁止倒计时（秒）
func (a ajaxAction) CountDown(value string) ajaxAction {
	return a.set("countDown", value)
}

// CountDownTpl 倒计时文字自定义
func (a ajaxAction) CountDownTpl(value string) ajaxAction {
	return a.set("countDownTpl", value)
}

// Disabled 是否禁用
func (a ajaxAction) Disabled(value bool) ajaxAction {
	return a.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (a ajaxAction) DisabledOn(value string) ajaxAction {
	return a.set("disabledOn", value)
}

// DisabledTip 禁用时的文案提示。
func (a ajaxAction) DisabledTip(value string) ajaxAction {
	return a.set("disabledTip", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (a ajaxAction) EditorSetting(value string) ajaxAction {
	return a.set("editorSetting", value)
}

// Feedback 自定义反馈内容
func (a ajaxAction) Feedback(value string) ajaxAction {
	return a.set("feedback", value)
}

// Hidden 是否隐藏
func (a ajaxAction) Hidden(value bool) ajaxAction {
	return a.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (a ajaxAction) HiddenOn(value string) ajaxAction {
	return a.set("hiddenOn", value)
}

// HotKey 键盘快捷键
func (a ajaxAction) HotKey(value string) ajaxAction {
	return a.set("hotKey", value)
}

// Icon 按钮图标， iconfont 的类名
func (a ajaxAction) Icon(value string) ajaxAction {
	return a.set("icon", value)
}

// IconClassName icon 上的 css 类名
func (a ajaxAction) IconClassName(value string) ajaxAction {
	return a.set("iconClassName", value)
}

// Id 主要用于用户行为跟踪里区分是哪个按钮
func (a ajaxAction) Id(value string) ajaxAction {
	return a.set("id", value)
}

// IgnoreConfirm 是否忽略确认
func (a ajaxAction) IgnoreConfirm(value bool) ajaxAction {
	return a.set("ignoreConfirm", value)
}

// IsolateScope 是否开启请求隔离
func (a ajaxAction) IsolateScope(value bool) ajaxAction {
	return a.set("isolateScope", value)
}

// Label 按钮文字
func (a ajaxAction) Label(value string) ajaxAction {
	return a.set("label", value)
}

// Level 按钮样式
func (a ajaxAction) Level(value string) ajaxAction {
	return a.set("level", value)
}

// LoadingClassName loading 上的 css 类名
func (a ajaxAction) LoadingClassName(value string) ajaxAction {
	return a.set("loadingClassName", value)
}

// LoadingOn 是否显示 loading 效果
func (a ajaxAction) LoadingOn(value string) ajaxAction {
	return a.set("loadingOn", value)
}

// MergeData 是否将弹框中数据 merge 到父级作用域
func (a ajaxAction) MergeData(value bool) ajaxAction {
	return a.set("mergeData", value)
}

// OnClick 自定义事件处理函数
func (a ajaxAction) OnClick(value string) ajaxAction {
	return a.set("onClick", value)
}

// OnEvent 事件动作配置
func (a ajaxAction) OnEvent(value string) ajaxAction {
	return a.set("onEvent", value)
}

// Primary 是否主要按钮
func (a ajaxAction) Primary(value bool) ajaxAction {
	return a.set("primary", value)
}

// Redirect 配置重定向
func (a ajaxAction) Redirect(value string) ajaxAction {
	return a.set("redirect", value)
}

// Reload 配置刷新动作
func (a ajaxAction) Reload(value string) ajaxAction {
	return a.set("reload", value)
}

// RequireSelected 是否要求选中
func (a ajaxAction) RequireSelected(value bool) ajaxAction {
	return a.set("requireSelected", value)
}

// Required 配置字段验证
func (a ajaxAction) Required(value string) ajaxAction {
	return a.set("required", value)
}

// RightIcon 右侧按钮图标
func (a ajaxAction) RightIcon(value string) ajaxAction {
	return a.set("rightIcon", value)
}

// RightIconClassName 右侧 icon 上的 css 类名
func (a ajaxAction) RightIconClassName(value string) ajaxAction {
	return a.set("rightIconClassName", value)
}

// Size 按钮大小
func (a ajaxAction) Size(value string) ajaxAction {
	return a.set("size", value)
}

// Static 是否静态展示
func (a ajaxAction) Static(value bool) ajaxAction {
	return a.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (a ajaxAction) StaticClassName(value string) ajaxAction {
	return a.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (a ajaxAction) StaticInputClassName(value string) ajaxAction {
	return a.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (a ajaxAction) StaticLabelClassName(value string) ajaxAction {
	return a.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (a ajaxAction) StaticOn(value string) ajaxAction {
	return a.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (a ajaxAction) StaticPlaceholder(value string) ajaxAction {
	return a.set("staticPlaceholder", value)
}

// StaticSchema 静态展示的 schema
func (a ajaxAction) StaticSchema(value string) ajaxAction {
	return a.set("staticSchema", value)
}

// Style 组件样式
func (a ajaxAction) Style(value string) ajaxAction {
	return a.set("style", value)
}

// Target 可以指定让谁来触发这个动作
func (a ajaxAction) Target(value string) ajaxAction {
	return a.set("target", value)
}

// TestIdBuilder 测试 ID 构建器
func (a ajaxAction) TestIdBuilder(value string) ajaxAction {
	return a.set("testIdBuilder", value)
}

// Testid 测试 ID
func (a ajaxAction) Testid(value string) ajaxAction {
	return a.set("testid", value)
}

// Tooltip 提示文本
func (a ajaxAction) Tooltip(value string) ajaxAction {
	return a.set("tooltip", value)
}

// TooltipPlacement 提示文本位置
func (a ajaxAction) TooltipPlacement(value string) ajaxAction {
	return a.set("tooltipPlacement", value)
}

// 指定按钮类型，支持 button、submit或者reset三种类型。 可选值: button | submit | reset
func (a ajaxAction) Type(value string) ajaxAction {
	return a.set("type", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (a ajaxAction) UseMobileUI(value bool) ajaxAction {
	return a.set("useMobileUI", value)
}

// Visible 是否显示
func (a ajaxAction) Visible(value bool) ajaxAction {
	return a.set("visible", value)
}

// VisibleOn 是否显示表达式
func (a ajaxAction) VisibleOn(value string) ajaxAction {
	return a.set("visibleOn", value)
}

func (a ajaxAction) set(key string, value interface{}) ajaxAction {
	a[key] = value
	return a
}
