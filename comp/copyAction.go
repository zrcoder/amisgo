package comp

// CopyAction 代表复制行为的按钮
type CopyAction Schema

// NewCopyAction 创建一个新的 CopyAction 实例，并设置默认的 actionType
func NewCopyAction() CopyAction {
	return make(CopyAction).set("type", "button").set("actionType", "copy")
}

func (c CopyAction) set(key string, value interface{}) CopyAction {
	c[key] = value
	return c
}

// ActionType 设置复制内容行为
func (c CopyAction) ActionType(value string) CopyAction {
	return c.set("actionType", value)
}

// ActiveClassName 激活状态时的类名
func (c CopyAction) ActiveClassName(value string) CopyAction {
	return c.set("activeClassName", value)
}

// ActiveLevel 激活状态时的样式
func (c CopyAction) ActiveLevel(value string) CopyAction {
	return c.set("activeLevel", value)
}

// Badge 角标
func (c CopyAction) Badge(value string) CopyAction {
	return c.set("badge", value)
}

// Block 是否为块状展示
func (c CopyAction) Block(value bool) CopyAction {
	return c.set("block", value)
}

// Body 子内容
func (c CopyAction) Body(value ...interface{}) CopyAction {
	return c.set("body", value)
}

// ClassName 容器 css 类名
func (c CopyAction) ClassName(value string) CopyAction {
	return c.set("className", value)
}

// Close 配置动作完成后是否关闭弹窗
func (c CopyAction) Close(value string) CopyAction {
	return c.set("close", value)
}

// ConfirmText 提示文字
func (c CopyAction) ConfirmText(value string) CopyAction {
	return c.set("confirmText", value)
}

// ConfirmTitle 确认弹窗标题
func (c CopyAction) ConfirmTitle(value string) CopyAction {
	return c.set("confirmTitle", value)
}

// Content 设置复制的内容
func (c CopyAction) Content(value string) CopyAction {
	return c.set("content", value)
}

// Copy 复制内容的配置
func (c CopyAction) Copy(value string) CopyAction {
	return c.set("copy", value)
}

// CountDown 点击后的禁止倒计时（秒）
func (c CopyAction) CountDown(value string) CopyAction {
	return c.set("countDown", value)
}

// CountDownTpl 倒计时文字自定义
func (c CopyAction) CountDownTpl(value string) CopyAction {
	return c.set("countDownTpl", value)
}

// Disabled 是否禁用
func (c CopyAction) Disabled(value bool) CopyAction {
	return c.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (c CopyAction) DisabledOn(value string) CopyAction {
	return c.set("disabledOn", value)
}

// DisabledTip 禁用时的文案提示
func (c CopyAction) DisabledTip(value string) CopyAction {
	return c.set("disabledTip", value)
}

// EditorSetting 编辑器配置
func (c CopyAction) EditorSetting(value string) CopyAction {
	return c.set("editorSetting", value)
}

// Hidden 是否隐藏
func (c CopyAction) Hidden(value bool) CopyAction {
	return c.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (c CopyAction) HiddenOn(value string) CopyAction {
	return c.set("hiddenOn", value)
}

// HotKey 键盘快捷键
func (c CopyAction) HotKey(value string) CopyAction {
	return c.set("hotKey", value)
}

// Icon 按钮图标
func (c CopyAction) Icon(value string) CopyAction {
	return c.set("icon", value)
}

// IconClassName icon 上的 css 类名
func (c CopyAction) IconClassName(value string) CopyAction {
	return c.set("iconClassName", value)
}

// ID 主要用于用户行为跟踪里区分是哪个按钮
func (c CopyAction) ID(value string) CopyAction {
	return c.set("id", value)
}

// Label 按钮文字
func (c CopyAction) Label(value string) CopyAction {
	return c.set("label", value)
}

// Level 按钮样式
func (c CopyAction) Level(value string) CopyAction {
	return c.set("level", value)
}

// LoadingClassName loading 上的 css 类名
func (c CopyAction) LoadingClassName(value string) CopyAction {
	return c.set("loadingClassName", value)
}

// LoadingOn 是否显示 loading 效果
func (c CopyAction) LoadingOn(value string) CopyAction {
	return c.set("loadingOn", value)
}

// MergeData 是否将弹框中数据 merge 到父级作用域
func (c CopyAction) MergeData(value bool) CopyAction {
	return c.set("mergeData", value)
}

// OnClick 自定义事件处理函数
func (c CopyAction) OnClick(value string) CopyAction {
	return c.set("onClick", value)
}

// OnEvent 事件动作配置
func (c CopyAction) OnEvent(value string) CopyAction {
	return c.set("onEvent", value)
}

// Primary 是否为主要按钮
func (c CopyAction) Primary(value bool) CopyAction {
	return c.set("primary", value)
}

// RequireSelected 是否需要勾选元素才能点击
func (c CopyAction) RequireSelected(value bool) CopyAction {
	return c.set("requireSelected", value)
}

// Required 配置字段验证
func (c CopyAction) Required(value string) CopyAction {
	return c.set("required", value)
}

// RightIcon 右侧按钮图标
func (c CopyAction) RightIcon(value string) CopyAction {
	return c.set("rightIcon", value)
}

// RightIconClassName 右侧 icon 上的 css 类名
func (c CopyAction) RightIconClassName(value string) CopyAction {
	return c.set("rightIconClassName", value)
}

// Size 按钮大小
func (c CopyAction) Size(value string) CopyAction {
	return c.set("size", value)
}

// Static 是否静态展示
func (c CopyAction) Static(value bool) CopyAction {
	return c.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (c CopyAction) StaticClassName(value string) CopyAction {
	return c.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (c CopyAction) StaticInputClassName(value string) CopyAction {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (c CopyAction) StaticLabelClassName(value string) CopyAction {
	return c.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (c CopyAction) StaticOn(value string) CopyAction {
	return c.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (c CopyAction) StaticPlaceholder(value string) CopyAction {
	return c.set("staticPlaceholder", value)
}

// StaticSchema 静态展示 schema
func (c CopyAction) StaticSchema(value string) CopyAction {
	return c.set("staticSchema", value)
}

// Style 组件样式
func (c CopyAction) Style(value string) CopyAction {
	return c.set("style", value)
}

// Target 指定触发动作的对象
func (c CopyAction) Target(value string) CopyAction {
	return c.set("target", value)
}

// TestIdBuilder 测试 ID 生成器
func (c CopyAction) TestIdBuilder(value string) CopyAction {
	return c.set("testIdBuilder", value)
}

// Testid 测试 ID
func (c CopyAction) Testid(value string) CopyAction {
	return c.set("testid", value)
}

// Tooltip 提示信息
func (c CopyAction) Tooltip(value string) CopyAction {
	return c.set("tooltip", value)
}

// TooltipPlacement 提示信息的位置
func (c CopyAction) TooltipPlacement(value string) CopyAction {
	return c.set("tooltipPlacement", value)
}

// UseMobileUI 关闭移动端样式
func (c CopyAction) UseMobileUI(value bool) CopyAction {
	return c.set("useMobileUI", value)
}

// Visible 是否显示
func (c CopyAction) Visible(value bool) CopyAction {
	return c.set("visible", value)
}

// VisibleOn 是否显示表达式
func (c CopyAction) VisibleOn(value string) CopyAction {
	return c.set("visibleOn", value)
}
