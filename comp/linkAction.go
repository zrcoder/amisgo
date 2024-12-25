package comp

// linkAction 链接动作控件

type linkAction Schema

// LinkAction 创建一个新的 LinkAction 实例，并设置默认的 type 和 actionType
func LinkAction() linkAction {
	return make(linkAction).set("type", "button").set("actionType", "link")
}

// ActiveClassName 激活状态时的类名
func (la linkAction) ActiveClassName(value string) linkAction {
	la.set("activeClassName", value)
	return la
}

// ActiveLevel 激活状态时的样式
func (la linkAction) ActiveLevel(value string) linkAction {
	la.set("activeLevel", value)
	return la
}

// Badge 角标
func (la linkAction) Badge(value string) linkAction {
	la.set("badge", value)
	return la
}

// Block 是否为块状展示
func (la linkAction) Block(value bool) linkAction {
	la.set("block", value)
	return la
}

// Body 子内容
func (la linkAction) Body(value ...any) linkAction {
	la.set("body", value)
	return la
}

// ClassName 容器 css 类名
func (la linkAction) ClassName(value string) linkAction {
	la.set("className", value)
	return la
}

// Close 如果按钮在弹框中，配置这个动作完成后是否关闭弹窗
func (la linkAction) Close(value string) linkAction {
	la.set("close", value)
	return la
}

// ConfirmText 提示文字
func (la linkAction) ConfirmText(value string) linkAction {
	la.set("confirmText", value)
	return la
}

// ConfirmTitle 确认弹窗标题
func (la linkAction) ConfirmTitle(value any) linkAction {
	la.set("confirmTitle", value)
	return la
}

// CountDown 点击后的禁止倒计时（秒）
func (la linkAction) CountDown(value string) linkAction {
	la.set("countDown", value)
	return la
}

// CountDownTpl 倒计时文字自定义
func (la linkAction) CountDownTpl(value string) linkAction {
	la.set("countDownTpl", value)
	return la
}

// Disabled 是否禁用
func (la linkAction) Disabled(value bool) linkAction {
	la.set("disabled", value)
	return la
}

// DisabledOn 是否禁用表达式
func (la linkAction) DisabledOn(value string) linkAction {
	la.set("disabledOn", value)
	return la
}

// DisabledTip 禁用时的文案提示
func (la linkAction) DisabledTip(value string) linkAction {
	la.set("disabledTip", value)
	return la
}

// EditorSetting 编辑器配置，运行时可以忽略
func (la linkAction) EditorSetting(value string) linkAction {
	la.set("editorSetting", value)
	return la
}

// Hidden 是否隐藏
func (la linkAction) Hidden(value bool) linkAction {
	la.set("hidden", value)
	return la
}

// HiddenOn 是否隐藏表达式
func (la linkAction) HiddenOn(value string) linkAction {
	la.set("hiddenOn", value)
	return la
}

// HotKey 键盘快捷键
func (la linkAction) HotKey(value string) linkAction {
	la.set("hotKey", value)
	return la
}

// Icon 按钮图标
func (la linkAction) Icon(value string) linkAction {
	la.set("icon", value)
	return la
}

// IconClassName icon 上的 css 类名
func (la linkAction) IconClassName(value string) linkAction {
	la.set("iconClassName", value)
	return la
}

// ID 主要用于用户行为跟踪里区分是哪个按钮
func (la linkAction) ID(value string) linkAction {
	la.set("id", value)
	return la
}

// Label 按钮文字
func (la linkAction) Label(value string) linkAction {
	la.set("label", value)
	return la
}

// Level 按钮样式
func (la linkAction) Level(value string) linkAction {
	la.set("level", value)
	return la
}

// Link 跳转到哪
func (la linkAction) Link(value string) linkAction {
	la.set("link", value)
	return la
}

// LoadingClassName loading 上的 css 类名
func (la linkAction) LoadingClassName(value string) linkAction {
	la.set("loadingClassName", value)
	return la
}

// LoadingOn 是否显示 loading 效果
func (la linkAction) LoadingOn(value string) linkAction {
	la.set("loadingOn", value)
	return la
}

// MergeData 是否将弹框中数据 merge 到父级作用域
func (la linkAction) MergeData(value bool) linkAction {
	la.set("mergeData", value)
	return la
}

// OnClick 自定义事件处理函数
func (la linkAction) OnClick(value string) linkAction {
	la.set("onClick", value)
	return la
}

// OnEvent 事件动作配置
func (la linkAction) OnEvent(value any) linkAction {
	la.set("onEvent", value)
	return la
}

// Primary 是否为主要按钮
func (la linkAction) Primary(value bool) linkAction {
	la.set("primary", value)
	return la
}

// RequireSelected 当按钮时批量操作按钮时，默认必须有勾选元素才能可点击
func (la linkAction) RequireSelected(value bool) linkAction {
	la.set("requireSelected", value)
	return la
}

// Required 如果按钮在 form 中，配置此属性会要求用户把指定的字段通过验证后才会触发行为
func (la linkAction) Required(value string) linkAction {
	la.set("required", value)
	return la
}

// RightIcon 右侧按钮图标
func (la linkAction) RightIcon(value string) linkAction {
	la.set("rightIcon", value)
	return la
}

// RightIconClassName 右侧 icon 上的 css 类名
func (la linkAction) RightIconClassName(value string) linkAction {
	la.set("rightIconClassName", value)
	return la
}

// Size 按钮大小
func (la linkAction) Size(value string) linkAction {
	la.set("size", value)
	return la
}

// Static 是否静态展示
func (la linkAction) Static(value bool) linkAction {
	la.set("static", value)
	return la
}

// StaticClassName 静态展示表单项类名
func (la linkAction) StaticClassName(value string) linkAction {
	la.set("staticClassName", value)
	return la
}

// StaticInputClassName 静态展示表单项 Value 类名
func (la linkAction) StaticInputClassName(value string) linkAction {
	la.set("staticInputClassName", value)
	return la
}

// StaticLabelClassName 静态展示表单项 Label 类名
func (la linkAction) StaticLabelClassName(value string) linkAction {
	la.set("staticLabelClassName", value)
	return la
}

// StaticOn 是否静态展示表达式
func (la linkAction) StaticOn(value string) linkAction {
	la.set("staticOn", value)
	return la
}

// StaticPlaceholder 静态展示空值占位
func (la linkAction) StaticPlaceholder(value string) linkAction {
	la.set("staticPlaceholder", value)
	return la
}

// StaticSchema 静态展示 Schema
func (la linkAction) StaticSchema(value string) linkAction {
	la.set("staticSchema", value)
	return la
}

// Style 组件样式
func (la linkAction) Style(value any) linkAction {
	la.set("style", value)
	return la
}

// Target 可以指定让谁来触发这个动作
func (la linkAction) Target(value string) linkAction {
	la.set("target", value)
	return la
}

// TestIdBuilder 测试 ID 构建器
func (la linkAction) TestIdBuilder(value string) linkAction {
	la.set("testIdBuilder", value)
	return la
}

// Testid 测试 id
func (la linkAction) Testid(value string) linkAction {
	la.set("testid", value)
	return la
}

// Tooltip 工具提示文字
func (la linkAction) Tooltip(value string) linkAction {
	la.set("tooltip", value)
	return la
}

// TooltipPlacement 工具提示位置
func (la linkAction) TooltipPlacement(value string) linkAction {
	la.set("tooltipPlacement", value)
	return la
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (la linkAction) UseMobileUI(value bool) linkAction {
	la.set("useMobileUI", value)
	return la
}

// Visible 是否显示
func (la linkAction) Visible(value bool) linkAction {
	la.set("visible", value)
	return la
}

// VisibleOn 是否显示表达式
func (la linkAction) VisibleOn(value string) linkAction {
	la.set("visibleOn", value)
	return la
}

// set 设置键值对
func (la linkAction) set(key string, value any) linkAction {
	la[key] = value
	return la
}
