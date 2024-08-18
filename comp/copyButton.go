package comp

// copyButton 代表复制行为的按钮
type copyButton schema

func CopyButton() copyButton {
	return make(copyButton).set("type", "button").set("actionType", "copy")
}

func (c copyButton) set(key string, value interface{}) copyButton {
	c[key] = value
	return c
}

// ActiveClassName 激活状态时的类名
func (c copyButton) ActiveClassName(value string) copyButton {
	return c.set("activeClassName", value)
}

// ActiveLevel 激活状态时的样式
func (c copyButton) ActiveLevel(value string) copyButton {
	return c.set("activeLevel", value)
}

// Badge 角标
func (c copyButton) Badge(value string) copyButton {
	return c.set("badge", value)
}

// Block 是否为块状展示
func (c copyButton) Block(value bool) copyButton {
	return c.set("block", value)
}

// Body 子内容
func (c copyButton) Body(value ...interface{}) copyButton {
	return c.set("body", value)
}

// ClassName 容器 css 类名
func (c copyButton) ClassName(value string) copyButton {
	return c.set("className", value)
}

// Close 配置动作完成后是否关闭弹窗
func (c copyButton) Close(value string) copyButton {
	return c.set("close", value)
}

// ConfirmText 提示文字
func (c copyButton) ConfirmText(value string) copyButton {
	return c.set("confirmText", value)
}

// ConfirmTitle 确认弹窗标题
func (c copyButton) ConfirmTitle(value string) copyButton {
	return c.set("confirmTitle", value)
}

// Content 设置复制的内容
func (c copyButton) Content(value string) copyButton {
	return c.set("content", value)
}

// Copy 复制内容的配置
func (c copyButton) Copy(value string) copyButton {
	return c.set("copy", value)
}

// CountDown 点击后的禁止倒计时（秒）
func (c copyButton) CountDown(value string) copyButton {
	return c.set("countDown", value)
}

// CountDownTpl 倒计时文字自定义
func (c copyButton) CountDownTpl(value string) copyButton {
	return c.set("countDownTpl", value)
}

// Disabled 是否禁用
func (c copyButton) Disabled(value bool) copyButton {
	return c.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (c copyButton) DisabledOn(value string) copyButton {
	return c.set("disabledOn", value)
}

// DisabledTip 禁用时的文案提示
func (c copyButton) DisabledTip(value string) copyButton {
	return c.set("disabledTip", value)
}

// EditorSetting 编辑器配置
func (c copyButton) EditorSetting(value string) copyButton {
	return c.set("editorSetting", value)
}

// Hidden 是否隐藏
func (c copyButton) Hidden(value bool) copyButton {
	return c.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (c copyButton) HiddenOn(value string) copyButton {
	return c.set("hiddenOn", value)
}

// HotKey 键盘快捷键
func (c copyButton) HotKey(value string) copyButton {
	return c.set("hotKey", value)
}

// Icon 按钮图标
func (c copyButton) Icon(value string) copyButton {
	return c.set("icon", value)
}

// IconClassName icon 上的 css 类名
func (c copyButton) IconClassName(value string) copyButton {
	return c.set("iconClassName", value)
}

// ID 主要用于用户行为跟踪里区分是哪个按钮
func (c copyButton) ID(value string) copyButton {
	return c.set("id", value)
}

// Label 按钮文字
func (c copyButton) Label(value string) copyButton {
	return c.set("label", value)
}

// Level 按钮样式
func (c copyButton) Level(value string) copyButton {
	return c.set("level", value)
}

// LoadingClassName loading 上的 css 类名
func (c copyButton) LoadingClassName(value string) copyButton {
	return c.set("loadingClassName", value)
}

// LoadingOn 是否显示 loading 效果
func (c copyButton) LoadingOn(value string) copyButton {
	return c.set("loadingOn", value)
}

// MergeData 是否将弹框中数据 merge 到父级作用域
func (c copyButton) MergeData(value bool) copyButton {
	return c.set("mergeData", value)
}

// OnClick 自定义事件处理函数
func (c copyButton) OnClick(value string) copyButton {
	return c.set("onClick", value)
}

// OnEvent 事件动作配置
func (c copyButton) OnEvent(value string) copyButton {
	return c.set("onEvent", value)
}

// Primary 是否为主要按钮
func (c copyButton) Primary(value bool) copyButton {
	return c.set("primary", value)
}

// RequireSelected 是否需要勾选元素才能点击
func (c copyButton) RequireSelected(value bool) copyButton {
	return c.set("requireSelected", value)
}

// Required 配置字段验证
func (c copyButton) Required(value string) copyButton {
	return c.set("required", value)
}

// RightIcon 右侧按钮图标
func (c copyButton) RightIcon(value string) copyButton {
	return c.set("rightIcon", value)
}

// RightIconClassName 右侧 icon 上的 css 类名
func (c copyButton) RightIconClassName(value string) copyButton {
	return c.set("rightIconClassName", value)
}

// Size 按钮大小
func (c copyButton) Size(value string) copyButton {
	return c.set("size", value)
}

// Static 是否静态展示
func (c copyButton) Static(value bool) copyButton {
	return c.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (c copyButton) StaticClassName(value string) copyButton {
	return c.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (c copyButton) StaticInputClassName(value string) copyButton {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (c copyButton) StaticLabelClassName(value string) copyButton {
	return c.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (c copyButton) StaticOn(value string) copyButton {
	return c.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (c copyButton) StaticPlaceholder(value string) copyButton {
	return c.set("staticPlaceholder", value)
}

// StaticSchema 静态展示 schema
func (c copyButton) StaticSchema(value string) copyButton {
	return c.set("staticSchema", value)
}

// Style 组件样式
func (c copyButton) Style(value string) copyButton {
	return c.set("style", value)
}

// Target 指定触发动作的对象
func (c copyButton) Target(value string) copyButton {
	return c.set("target", value)
}

// TestIdBuilder 测试 ID 生成器
func (c copyButton) TestIdBuilder(value string) copyButton {
	return c.set("testIdBuilder", value)
}

// Testid 测试 ID
func (c copyButton) Testid(value string) copyButton {
	return c.set("testid", value)
}

// Tooltip 提示信息
func (c copyButton) Tooltip(value string) copyButton {
	return c.set("tooltip", value)
}

// TooltipPlacement 提示信息的位置
func (c copyButton) TooltipPlacement(value string) copyButton {
	return c.set("tooltipPlacement", value)
}

// UseMobileUI 关闭移动端样式
func (c copyButton) UseMobileUI(value bool) copyButton {
	return c.set("useMobileUI", value)
}

// Visible 是否显示
func (c copyButton) Visible(value bool) copyButton {
	return c.set("visible", value)
}

// VisibleOn 是否显示表达式
func (c copyButton) VisibleOn(value string) copyButton {
	return c.set("visibleOn", value)
}
