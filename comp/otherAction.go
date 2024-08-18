package comp

// OtherAction 表示其他动作的组件
//
// @version 6.7.0
type OtherAction Schema

// NewOtherAction 创建一个新的 OtherAction 实例
func NewOtherAction() OtherAction {
	return OtherAction{}.set("type", "button").set("actionType", "prev")
}

// set 设置字段值
func (oa OtherAction) set(key string, value interface{}) OtherAction {
	oa[key] = value
	return oa
}

// ActionType 可选值: prev | next | cancel | close | submit | confirm | add | reset | reset-and-submit
func (oa OtherAction) ActionType(value string) OtherAction {
	return oa.set("actionType", value)
}

// ActiveClassName 激活状态时的类名
func (oa OtherAction) ActiveClassName(value string) OtherAction {
	return oa.set("activeClassName", value)
}

// ActiveLevel 激活状态时的样式
func (oa OtherAction) ActiveLevel(value string) OtherAction {
	return oa.set("activeLevel", value)
}

// Badge 角标 (Badge 角标)
func (oa OtherAction) Badge(value string) OtherAction {
	return oa.set("badge", value)
}

// Block 是否为块状展示，默认为内联
func (oa OtherAction) Block(value bool) OtherAction {
	return oa.set("block", value)
}

// Body 子内容
func (oa OtherAction) Body(value ...interface{}) OtherAction {
	return oa.set("body", value)
}

// ClassName 容器 css 类名
func (oa OtherAction) ClassName(value string) OtherAction {
	return oa.set("className", value)
}

// Close 如果按钮在弹框中，可以配置这个动作完成后是否关闭弹窗，或者指定关闭目标弹框
func (oa OtherAction) Close(value string) OtherAction {
	return oa.set("close", value)
}

// ConfirmText 提示文字，配置了操作前会要求用户确认
func (oa OtherAction) ConfirmText(value string) OtherAction {
	return oa.set("confirmText", value)
}

// ConfirmTitle 确认弹窗标题
func (oa OtherAction) ConfirmTitle(value string) OtherAction {
	return oa.set("confirmTitle", value)
}

// CountDown 点击后的禁止倒计时（秒）
func (oa OtherAction) CountDown(value string) OtherAction {
	return oa.set("countDown", value)
}

// CountDownTpl 倒计时文字自定义
func (oa OtherAction) CountDownTpl(value string) OtherAction {
	return oa.set("countDownTpl", value)
}

// Disabled 是否禁用
func (oa OtherAction) Disabled(value bool) OtherAction {
	return oa.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (oa OtherAction) DisabledOn(value string) OtherAction {
	return oa.set("disabledOn", value)
}

// DisabledTip 禁用时的文案提示
func (oa OtherAction) DisabledTip(value string) OtherAction {
	return oa.set("disabledTip", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (oa OtherAction) EditorSetting(value string) OtherAction {
	return oa.set("editorSetting", value)
}

// Hidden 是否隐藏
func (oa OtherAction) Hidden(value bool) OtherAction {
	return oa.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (oa OtherAction) HiddenOn(value string) OtherAction {
	return oa.set("hiddenOn", value)
}

// HotKey 键盘快捷键
func (oa OtherAction) HotKey(value string) OtherAction {
	return oa.set("hotKey", value)
}

// Icon 按钮图标， iconfont 的类名
func (oa OtherAction) Icon(value string) OtherAction {
	return oa.set("icon", value)
}

// IconClassName icon 上的 css 类名
func (oa OtherAction) IconClassName(value string) OtherAction {
	return oa.set("iconClassName", value)
}

// ID 主要用于用户行为跟踪里区分是哪个按钮
func (oa OtherAction) ID(value string) OtherAction {
	return oa.set("id", value)
}

// Label 按钮文字
func (oa OtherAction) Label(value string) OtherAction {
	return oa.set("label", value)
}

// Level 按钮样式，可选值: info | success | warning | danger | link | primary | dark | light | secondary
func (oa OtherAction) Level(value string) OtherAction {
	return oa.set("level", value)
}

// LoadingClassName loading 上的 css 类名
func (oa OtherAction) LoadingClassName(value string) OtherAction {
	return oa.set("loadingClassName", value)
}

// LoadingOn 是否显示 loading 效果
func (oa OtherAction) LoadingOn(value string) OtherAction {
	return oa.set("loadingOn", value)
}

// MergeData 是否将弹框中数据 merge 到父级作用域
func (oa OtherAction) MergeData(value bool) OtherAction {
	return oa.set("mergeData", value)
}

// OnClick 自定义事件处理函数
func (oa OtherAction) OnClick(value string) OtherAction {
	return oa.set("onClick", value)
}

// OnEvent 事件动作配置
func (oa OtherAction) OnEvent(value string) OtherAction {
	return oa.set("onEvent", value)
}

// Primary 是否为主要按钮
func (oa OtherAction) Primary(value bool) OtherAction {
	return oa.set("primary", value)
}

// RequireSelected 当按钮时批量操作按钮时，默认必须有勾选元素才能可点击
func (oa OtherAction) RequireSelected(value bool) OtherAction {
	return oa.set("requireSelected", value)
}

// Required 如果按钮在 form 中，配置此属性会要求用户把指定的字段通过验证后才会触发行为
func (oa OtherAction) Required(value string) OtherAction {
	return oa.set("required", value)
}

// RightIcon 右侧按钮图标， iconfont 的类名
func (oa OtherAction) RightIcon(value string) OtherAction {
	return oa.set("rightIcon", value)
}

// RightIconClassName 右侧 icon 上的 css 类名
func (oa OtherAction) RightIconClassName(value string) OtherAction {
	return oa.set("rightIconClassName", value)
}

// Size 按钮大小，可选值: xs | sm | md | lg
func (oa OtherAction) Size(value string) OtherAction {
	return oa.set("size", value)
}

// Static 是否静态展示
func (oa OtherAction) Static(value bool) OtherAction {
	return oa.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (oa OtherAction) StaticClassName(value string) OtherAction {
	return oa.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项 Value 类名
func (oa OtherAction) StaticInputClassName(value string) OtherAction {
	return oa.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项 Label 类名
func (oa OtherAction) StaticLabelClassName(value string) OtherAction {
	return oa.set("staticLabelClassName", value)
}

// StaticOn 静态展示表达式
func (oa OtherAction) StaticOn(value string) OtherAction {
	return oa.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (oa OtherAction) StaticPlaceholder(value string) OtherAction {
	return oa.set("staticPlaceholder", value)
}

// StaticSchema 静态展示模式的 schema
func (oa OtherAction) StaticSchema(value string) OtherAction {
	return oa.set("staticSchema", value)
}

// Style 组件样式
func (oa OtherAction) Style(value string) OtherAction {
	return oa.set("style", value)
}

// Target 可以指定让谁来触发这个动作
func (oa OtherAction) Target(value string) OtherAction {
	return oa.set("target", value)
}

// TestIdBuilder 自定义测试 ID 构建器
func (oa OtherAction) TestIdBuilder(value string) OtherAction {
	return oa.set("testIdBuilder", value)
}

// TestID 测试 ID
func (oa OtherAction) TestID(value string) OtherAction {
	return oa.set("testid", value)
}

// Tooltip 提示信息
func (oa OtherAction) Tooltip(value string) OtherAction {
	return oa.set("tooltip", value)
}

// TooltipPlacement 提示信息位置，可选值: top | right | bottom | left
func (oa OtherAction) TooltipPlacement(value string) OtherAction {
	return oa.set("tooltipPlacement", value)
}

// Type 指定按钮类型，支持 button、submit 或者 reset 三种类型
func (oa OtherAction) Type(value string) OtherAction {
	return oa.set("type", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (oa OtherAction) UseMobileUI(value bool) OtherAction {
	return oa.set("useMobileUI", value)
}

// Visible 是否显示
func (oa OtherAction) Visible(value bool) OtherAction {
	return oa.set("visible", value)
}

// VisibleOn 是否显示表达式
func (oa OtherAction) VisibleOn(value string) OtherAction {
	return oa.set("visibleOn", value)
}
