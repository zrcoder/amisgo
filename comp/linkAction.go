package comp

// LinkAction 链接动作控件
//
// @version 6.7.0
type LinkAction Schema

// NewLinkAction 创建一个新的 LinkAction 实例，并设置默认的 type 和 actionType
func NewLinkAction() LinkAction {
	return make(LinkAction).set("type", "button").set("actionType", "link")
}

// ActiveClassName 激活状态时的类名
func (la LinkAction) ActiveClassName(value string) LinkAction {
	la.set("activeClassName", value)
	return la
}

// ActiveLevel 激活状态时的样式
func (la LinkAction) ActiveLevel(value string) LinkAction {
	la.set("activeLevel", value)
	return la
}

// Badge 角标
func (la LinkAction) Badge(value string) LinkAction {
	la.set("badge", value)
	return la
}

// Block 是否为块状展示
func (la LinkAction) Block(value bool) LinkAction {
	la.set("block", value)
	return la
}

// Body 子内容
func (la LinkAction) Body(value ...interface{}) LinkAction {
	la.set("body", value)
	return la
}

// ClassName 容器 css 类名
func (la LinkAction) ClassName(value string) LinkAction {
	la.set("className", value)
	return la
}

// Close 如果按钮在弹框中，配置这个动作完成后是否关闭弹窗
func (la LinkAction) Close(value string) LinkAction {
	la.set("close", value)
	return la
}

// ConfirmText 提示文字
func (la LinkAction) ConfirmText(value string) LinkAction {
	la.set("confirmText", value)
	return la
}

// ConfirmTitle 确认弹窗标题
func (la LinkAction) ConfirmTitle(value string) LinkAction {
	la.set("confirmTitle", value)
	return la
}

// CountDown 点击后的禁止倒计时（秒）
func (la LinkAction) CountDown(value string) LinkAction {
	la.set("countDown", value)
	return la
}

// CountDownTpl 倒计时文字自定义
func (la LinkAction) CountDownTpl(value string) LinkAction {
	la.set("countDownTpl", value)
	return la
}

// Disabled 是否禁用
func (la LinkAction) Disabled(value bool) LinkAction {
	la.set("disabled", value)
	return la
}

// DisabledOn 是否禁用表达式
func (la LinkAction) DisabledOn(value string) LinkAction {
	la.set("disabledOn", value)
	return la
}

// DisabledTip 禁用时的文案提示
func (la LinkAction) DisabledTip(value string) LinkAction {
	la.set("disabledTip", value)
	return la
}

// EditorSetting 编辑器配置，运行时可以忽略
func (la LinkAction) EditorSetting(value string) LinkAction {
	la.set("editorSetting", value)
	return la
}

// Hidden 是否隐藏
func (la LinkAction) Hidden(value bool) LinkAction {
	la.set("hidden", value)
	return la
}

// HiddenOn 是否隐藏表达式
func (la LinkAction) HiddenOn(value string) LinkAction {
	la.set("hiddenOn", value)
	return la
}

// HotKey 键盘快捷键
func (la LinkAction) HotKey(value string) LinkAction {
	la.set("hotKey", value)
	return la
}

// Icon 按钮图标
func (la LinkAction) Icon(value string) LinkAction {
	la.set("icon", value)
	return la
}

// IconClassName icon 上的 css 类名
func (la LinkAction) IconClassName(value string) LinkAction {
	la.set("iconClassName", value)
	return la
}

// ID 主要用于用户行为跟踪里区分是哪个按钮
func (la LinkAction) ID(value string) LinkAction {
	la.set("id", value)
	return la
}

// Label 按钮文字
func (la LinkAction) Label(value string) LinkAction {
	la.set("label", value)
	return la
}

// Level 按钮样式
func (la LinkAction) Level(value string) LinkAction {
	la.set("level", value)
	return la
}

// Link 跳转到哪
func (la LinkAction) Link(value string) LinkAction {
	la.set("link", value)
	return la
}

// LoadingClassName loading 上的 css 类名
func (la LinkAction) LoadingClassName(value string) LinkAction {
	la.set("loadingClassName", value)
	return la
}

// LoadingOn 是否显示 loading 效果
func (la LinkAction) LoadingOn(value string) LinkAction {
	la.set("loadingOn", value)
	return la
}

// MergeData 是否将弹框中数据 merge 到父级作用域
func (la LinkAction) MergeData(value bool) LinkAction {
	la.set("mergeData", value)
	return la
}

// OnClick 自定义事件处理函数
func (la LinkAction) OnClick(value string) LinkAction {
	la.set("onClick", value)
	return la
}

// OnEvent 事件动作配置
func (la LinkAction) OnEvent(value string) LinkAction {
	la.set("onEvent", value)
	return la
}

// Primary 是否为主要按钮
func (la LinkAction) Primary(value bool) LinkAction {
	la.set("primary", value)
	return la
}

// RequireSelected 当按钮时批量操作按钮时，默认必须有勾选元素才能可点击
func (la LinkAction) RequireSelected(value bool) LinkAction {
	la.set("requireSelected", value)
	return la
}

// Required 如果按钮在 form 中，配置此属性会要求用户把指定的字段通过验证后才会触发行为
func (la LinkAction) Required(value string) LinkAction {
	la.set("required", value)
	return la
}

// RightIcon 右侧按钮图标
func (la LinkAction) RightIcon(value string) LinkAction {
	la.set("rightIcon", value)
	return la
}

// RightIconClassName 右侧 icon 上的 css 类名
func (la LinkAction) RightIconClassName(value string) LinkAction {
	la.set("rightIconClassName", value)
	return la
}

// Size 按钮大小
func (la LinkAction) Size(value string) LinkAction {
	la.set("size", value)
	return la
}

// Static 是否静态展示
func (la LinkAction) Static(value bool) LinkAction {
	la.set("static", value)
	return la
}

// StaticClassName 静态展示表单项类名
func (la LinkAction) StaticClassName(value string) LinkAction {
	la.set("staticClassName", value)
	return la
}

// StaticInputClassName 静态展示表单项 Value 类名
func (la LinkAction) StaticInputClassName(value string) LinkAction {
	la.set("staticInputClassName", value)
	return la
}

// StaticLabelClassName 静态展示表单项 Label 类名
func (la LinkAction) StaticLabelClassName(value string) LinkAction {
	la.set("staticLabelClassName", value)
	return la
}

// StaticOn 是否静态展示表达式
func (la LinkAction) StaticOn(value string) LinkAction {
	la.set("staticOn", value)
	return la
}

// StaticPlaceholder 静态展示空值占位
func (la LinkAction) StaticPlaceholder(value string) LinkAction {
	la.set("staticPlaceholder", value)
	return la
}

// StaticSchema 静态展示 Schema
func (la LinkAction) StaticSchema(value string) LinkAction {
	la.set("staticSchema", value)
	return la
}

// Style 组件样式
func (la LinkAction) Style(value string) LinkAction {
	la.set("style", value)
	return la
}

// Target 可以指定让谁来触发这个动作
func (la LinkAction) Target(value string) LinkAction {
	la.set("target", value)
	return la
}

// TestIdBuilder 测试 ID 构建器
func (la LinkAction) TestIdBuilder(value string) LinkAction {
	la.set("testIdBuilder", value)
	return la
}

// Testid 测试 id
func (la LinkAction) Testid(value string) LinkAction {
	la.set("testid", value)
	return la
}

// Tooltip 工具提示文字
func (la LinkAction) Tooltip(value string) LinkAction {
	la.set("tooltip", value)
	return la
}

// TooltipPlacement 工具提示位置
func (la LinkAction) TooltipPlacement(value string) LinkAction {
	la.set("tooltipPlacement", value)
	return la
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (la LinkAction) UseMobileUI(value bool) LinkAction {
	la.set("useMobileUI", value)
	return la
}

// Visible 是否显示
func (la LinkAction) Visible(value bool) LinkAction {
	la.set("visible", value)
	return la
}

// VisibleOn 是否显示表达式
func (la LinkAction) VisibleOn(value string) LinkAction {
	la.set("visibleOn", value)
	return la
}

// set 设置键值对
func (la LinkAction) set(key string, value interface{}) LinkAction {
	la[key] = value
	return la
}
