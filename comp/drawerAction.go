package comp

// DrawerAction 表示一个抽出式弹框中的操作按钮。
type DrawerAction BaseRenderer

// NewDrawerAction 创建一个新的 DrawerAction 实例，并设置默认的 type 和 actionType
func NewDrawerAction() DrawerAction {
	d := DrawerAction(make(BaseRenderer))
	d.set("type", "button")
	d.set("actionType", "drawer")
	return d
}

func (d DrawerAction) set(key string, value interface{}) DrawerAction {
	d[key] = value
	return d
}

// ActionType 指定为打开弹窗，抽出式弹窗
func (d DrawerAction) ActionType(value string) DrawerAction {
	return d.set("actionType", value)
}

// ActiveClassName 激活状态时的类名
func (d DrawerAction) ActiveClassName(value string) DrawerAction {
	return d.set("activeClassName", value)
}

// ActiveLevel 激活状态时的样式
func (d DrawerAction) ActiveLevel(value string) DrawerAction {
	return d.set("activeLevel", value)
}

// Align 设置对齐方式
func (d DrawerAction) Align(value string) DrawerAction {
	return d.set("align", value)
}

// Badge 角标
func (d DrawerAction) Badge(value string) DrawerAction {
	return d.set("badge", value)
}

// Block 是否为块状展示，默认为内联
func (d DrawerAction) Block(value bool) DrawerAction {
	return d.set("block", value)
}

// Body 子内容
func (d DrawerAction) Body(value ...interface{}) DrawerAction {
	return d.set("body", value)
}

// ClassName 容器 css 类名
func (d DrawerAction) ClassName(value string) DrawerAction {
	return d.set("className", value)
}

// Close 如果按钮在弹框中，可以配置这个动作完成后是否关闭弹窗，或者指定关闭目标弹框
func (d DrawerAction) Close(value string) DrawerAction {
	return d.set("close", value)
}

// ConfirmText 提示文字
func (d DrawerAction) ConfirmText(value string) DrawerAction {
	return d.set("confirmText", value)
}

// ConfirmTitle 确认弹窗标题
func (d DrawerAction) ConfirmTitle(value string) DrawerAction {
	return d.set("confirmTitle", value)
}

// CountDown 点击后的禁止倒计时（秒）
func (d DrawerAction) CountDown(value string) DrawerAction {
	return d.set("countDown", value)
}

// CountDownTpl 倒计时文字自定义
func (d DrawerAction) CountDownTpl(value string) DrawerAction {
	return d.set("countDownTpl", value)
}

// Data 数据映射
func (d DrawerAction) Data(value string) DrawerAction {
	return d.set("data", value)
}

// Disabled 是否禁用
func (d DrawerAction) Disabled(value bool) DrawerAction {
	return d.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (d DrawerAction) DisabledOn(value string) DrawerAction {
	return d.set("disabledOn", value)
}

// DisabledTip 禁用时的文案提示
func (d DrawerAction) DisabledTip(value string) DrawerAction {
	return d.set("disabledTip", value)
}

// Drawer 抽出式弹框详情
func (d DrawerAction) Drawer(value string) DrawerAction {
	return d.set("drawer", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (d DrawerAction) EditorSetting(value string) DrawerAction {
	return d.set("editorSetting", value)
}

// Hidden 是否隐藏
func (d DrawerAction) Hidden(value bool) DrawerAction {
	return d.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (d DrawerAction) HiddenOn(value string) DrawerAction {
	return d.set("hiddenOn", value)
}

// HotKey 键盘快捷键
func (d DrawerAction) HotKey(value string) DrawerAction {
	return d.set("hotKey", value)
}

// Icon 按钮图标
func (d DrawerAction) Icon(value string) DrawerAction {
	return d.set("icon", value)
}

// IconClassName icon 上的 css 类名
func (d DrawerAction) IconClassName(value string) DrawerAction {
	return d.set("iconClassName", value)
}

// ID 主要用于用户行为跟踪里区分是哪个按钮
func (d DrawerAction) ID(value string) DrawerAction {
	return d.set("id", value)
}

// Label 按钮文字
func (d DrawerAction) Label(value string) DrawerAction {
	return d.set("label", value)
}

// Level 按钮样式
func (d DrawerAction) Level(value string) DrawerAction {
	return d.set("level", value)
}

// LoadingClassName loading 上的 css 类名
func (d DrawerAction) LoadingClassName(value string) DrawerAction {
	return d.set("loadingClassName", value)
}

// LoadingOn 是否显示 loading 效果
func (d DrawerAction) LoadingOn(value string) DrawerAction {
	return d.set("loadingOn", value)
}

// MergeData 是否将弹框中数据 merge 到父级作用域
func (d DrawerAction) MergeData(value bool) DrawerAction {
	return d.set("mergeData", value)
}

// NextCondition 是否有下一个的表达式
func (d DrawerAction) NextCondition(value string) DrawerAction {
	return d.set("nextCondition", value)
}

// OnClick 自定义事件处理函数
func (d DrawerAction) OnClick(value string) DrawerAction {
	return d.set("onClick", value)
}

// OnEvent 事件动作配置
func (d DrawerAction) OnEvent(value string) DrawerAction {
	return d.set("onEvent", value)
}

// Primary 标识是否为主要按钮
func (d DrawerAction) Primary(value bool) DrawerAction {
	return d.set("primary", value)
}

// Redirect 重定向配置
func (d DrawerAction) Redirect(value string) DrawerAction {
	return d.set("redirect", value)
}

// Reload 配置刷新动作
func (d DrawerAction) Reload(value string) DrawerAction {
	return d.set("reload", value)
}

// RequireSelected 当按钮是批量操作按钮时，是否需要有勾选元素
func (d DrawerAction) RequireSelected(value bool) DrawerAction {
	return d.set("requireSelected", value)
}

// Required 如果按钮在 form 中，配置此属性会要求用户把指定的字段通过验证后才会触发行为
func (d DrawerAction) Required(value string) DrawerAction {
	return d.set("required", value)
}

// RightIcon 右侧按钮图标
func (d DrawerAction) RightIcon(value string) DrawerAction {
	return d.set("rightIcon", value)
}

// RightIconClassName 右侧 icon 上的 css 类名
func (d DrawerAction) RightIconClassName(value string) DrawerAction {
	return d.set("rightIconClassName", value)
}

// Size 按钮大小
func (d DrawerAction) Size(value string) DrawerAction {
	return d.set("size", value)
}

// Static 是否静态展示
func (d DrawerAction) Static(value bool) DrawerAction {
	return d.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (d DrawerAction) StaticClassName(value string) DrawerAction {
	return d.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (d DrawerAction) StaticInputClassName(value string) DrawerAction {
	return d.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (d DrawerAction) StaticLabelClassName(value string) DrawerAction {
	return d.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (d DrawerAction) StaticOn(value string) DrawerAction {
	return d.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (d DrawerAction) StaticPlaceholder(value string) DrawerAction {
	return d.set("staticPlaceholder", value)
}

// StaticSchema 静态展示 schema
func (d DrawerAction) StaticSchema(value string) DrawerAction {
	return d.set("staticSchema", value)
}

// Style 组件样式
func (d DrawerAction) Style(value string) DrawerAction {
	return d.set("style", value)
}

// Target 可以指定让谁来触发这个动作
func (d DrawerAction) Target(value string) DrawerAction {
	return d.set("target", value)
}

// TestIdBuilder 用于测试的 ID 构造函数
func (d DrawerAction) TestIdBuilder(value string) DrawerAction {
	return d.set("testIdBuilder", value)
}

// Testid 测试 ID
func (d DrawerAction) Testid(value string) DrawerAction {
	return d.set("testid", value)
}

// Tooltip 提示信息
func (d DrawerAction) Tooltip(value string) DrawerAction {
	return d.set("tooltip", value)
}

// TooltipPlacement 提示信息位置
func (d DrawerAction) TooltipPlacement(value string) DrawerAction {
	return d.set("tooltipPlacement", value)
}

// Type 指定按钮类型，支持 button、submit 或 reset 三种类型
func (d DrawerAction) Type(value string) DrawerAction {
	return d.set("type", value)
}

// UseMobileUI 组件级别用来关闭移动端样式
func (d DrawerAction) UseMobileUI(value bool) DrawerAction {
	return d.set("useMobileUI", value)
}

// Visible 是否显示
func (d DrawerAction) Visible(value bool) DrawerAction {
	return d.set("visible", value)
}

// VisibleOn 是否显示表达式
func (d DrawerAction) VisibleOn(value string) DrawerAction {
	return d.set("visibleOn", value)
}
