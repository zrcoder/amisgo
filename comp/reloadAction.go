package comp

// reloadAction 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/button
//
// @version 6.7.0
type reloadAction schema

// ReloadAction 创建一个新的 ReloadAction 实例
func ReloadAction() reloadAction {
	return reloadAction{}.set("type", "button").set("actionType", "reload")
}

func (ra reloadAction) set(key string, value interface{}) reloadAction {
	ra[key] = value
	return ra
}

// ActiveClassName 激活状态时的类名
func (ra reloadAction) ActiveClassName(value string) reloadAction {
	return ra.set("activeClassName", value)
}

// ActiveLevel 激活状态时的样式
func (ra reloadAction) ActiveLevel(value string) reloadAction {
	return ra.set("activeLevel", value)
}

// Badge 角标 (Badge 角标。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/badge)
func (ra reloadAction) Badge(value string) reloadAction {
	return ra.set("badge", value)
}

// Block 是否为块状展示，默认为内联
func (ra reloadAction) Block(value bool) reloadAction {
	return ra.set("block", value)
}

// Body 子内容
func (ra reloadAction) Body(value ...interface{}) reloadAction {
	return ra.set("body", value)
}

// ClassName 容器 css 类名
func (ra reloadAction) ClassName(value string) reloadAction {
	return ra.set("className", value)
}

// Close 如果按钮在弹框中，可以配置这个动作完成后是否关闭弹窗，或者指定关闭目标弹框
func (ra reloadAction) Close(value string) reloadAction {
	return ra.set("close", value)
}

// ConfirmText 提示文字，配置了操作前会要求用户确认
func (ra reloadAction) ConfirmText(value string) reloadAction {
	return ra.set("confirmText", value)
}

// ConfirmTitle 确认弹窗标题
func (ra reloadAction) ConfirmTitle(value string) reloadAction {
	return ra.set("confirmTitle", value)
}

// CountDown 点击后的禁止倒计时（秒）
func (ra reloadAction) CountDown(value string) reloadAction {
	return ra.set("countDown", value)
}

// CountDownTpl 倒计时文字自定义
func (ra reloadAction) CountDownTpl(value string) reloadAction {
	return ra.set("countDownTpl", value)
}

// Disabled 是否禁用
func (ra reloadAction) Disabled(value bool) reloadAction {
	return ra.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`)
func (ra reloadAction) DisabledOn(value string) reloadAction {
	return ra.set("disabledOn", value)
}

// DisabledTip 禁用时的文案提示
func (ra reloadAction) DisabledTip(value string) reloadAction {
	return ra.set("disabledTip", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (ra reloadAction) EditorSetting(value string) reloadAction {
	return ra.set("editorSetting", value)
}

// Hidden 是否隐藏
func (ra reloadAction) Hidden(value bool) reloadAction {
	return ra.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`)
func (ra reloadAction) HiddenOn(value string) reloadAction {
	return ra.set("hiddenOn", value)
}

// HotKey 键盘快捷键
func (ra reloadAction) HotKey(value string) reloadAction {
	return ra.set("hotKey", value)
}

// Icon 按钮图标， iconfont 的类名 (iconfont 里面的类名。)
func (ra reloadAction) Icon(value string) reloadAction {
	return ra.set("icon", value)
}

// IconClassName icon 上的css 类名
func (ra reloadAction) IconClassName(value string) reloadAction {
	return ra.set("iconClassName", value)
}

// Id 主要用于用户行为跟踪里区分是哪个按钮
func (ra reloadAction) Id(value string) reloadAction {
	return ra.set("id", value)
}

// Label 按钮文字
func (ra reloadAction) Label(value string) reloadAction {
	return ra.set("label", value)
}

// Level 按钮样式 可选值: info | success | warning | danger | link | primary | dark | light | secondary
func (ra reloadAction) Level(value string) reloadAction {
	return ra.set("level", value)
}

// LoadingClassName loading 上的css 类名
func (ra reloadAction) LoadingClassName(value string) reloadAction {
	return ra.set("loadingClassName", value)
}

// LoadingOn 是否显示loading效果
func (ra reloadAction) LoadingOn(value string) reloadAction {
	return ra.set("loadingOn", value)
}

// MergeData 是否将弹框中数据 merge 到父级作用域
func (ra reloadAction) MergeData(value bool) reloadAction {
	return ra.set("mergeData", value)
}

// OnClick 自定义事件处理函数
func (ra reloadAction) OnClick(value string) reloadAction {
	return ra.set("onClick", value)
}

// OnEvent 事件动作配置
func (ra reloadAction) OnEvent(value string) reloadAction {
	return ra.set("onEvent", value)
}

// Primary
func (ra reloadAction) Primary(value bool) reloadAction {
	return ra.set("primary", value)
}

// RequireSelected 当按钮时批量操作按钮时，默认必须有勾选元素才能可点击，如果此属性配置成 false，则没有点选成员也能点击
func (ra reloadAction) RequireSelected(value bool) reloadAction {
	return ra.set("requireSelected", value)
}

// Required 如果按钮在form中，配置此属性会要求用户把指定的字段通过验证后才会触发行为
func (ra reloadAction) Required(value string) reloadAction {
	return ra.set("required", value)
}

// RightIcon 右侧按钮图标， iconfont 的类名 (iconfont 里面的类名。)
func (ra reloadAction) RightIcon(value string) reloadAction {
	return ra.set("rightIcon", value)
}

// RightIconClassName 右侧 icon 上的 css 类名
func (ra reloadAction) RightIconClassName(value string) reloadAction {
	return ra.set("rightIconClassName", value)
}

// Size 按钮大小 可选值: xs | sm | md | lg
func (ra reloadAction) Size(value string) reloadAction {
	return ra.set("size", value)
}

// Static 是否静态展示
func (ra reloadAction) Static(value bool) reloadAction {
	return ra.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (ra reloadAction) StaticClassName(value string) reloadAction {
	return ra.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (ra reloadAction) StaticInputClassName(value string) reloadAction {
	return ra.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (ra reloadAction) StaticLabelClassName(value string) reloadAction {
	return ra.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`)
func (ra reloadAction) StaticOn(value string) reloadAction {
	return ra.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (ra reloadAction) StaticPlaceholder(value string) reloadAction {
	return ra.set("staticPlaceholder", value)
}

// StaticSchema
func (ra reloadAction) StaticSchema(value string) reloadAction {
	return ra.set("staticSchema", value)
}

// Style 组件样式
func (ra reloadAction) Style(value string) reloadAction {
	return ra.set("style", value)
}

// Target 指定目标组件。 (配置刷新动作，这个动作通常在完成渲染器本省的固定动作后出发。一般用来配置目标组件的 name 属性。多个目标可以用逗号隔开。当目标是 windows 时表示刷新整个页面。刷新目标的同时还支持传递参数如： `foo?a=${a}&b=${b},boo?c=${c}`)
func (ra reloadAction) Target(value string) reloadAction {
	return ra.set("target", value)
}

// TestIdBuilder
func (ra reloadAction) TestIdBuilder(value string) reloadAction {
	return ra.set("testIdBuilder", value)
}

// Testid
func (ra reloadAction) Testid(value string) reloadAction {
	return ra.set("testid", value)
}

// Tooltip
func (ra reloadAction) Tooltip(value string) reloadAction {
	return ra.set("tooltip", value)
}

// TooltipPlacement 可选值: top | right | bottom | left
func (ra reloadAction) TooltipPlacement(value string) reloadAction {
	return ra.set("tooltipPlacement", value)
}

// Type 指定按钮类型，支持 button、submit或者reset三种类型。 可选值: button | submit | reset
func (ra reloadAction) Type(value string) reloadAction {
	return ra.set("type", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (ra reloadAction) UseMobileUI(value bool) reloadAction {
	return ra.set("useMobileUI", value)
}

// Visible 是否显示
func (ra reloadAction) Visible(value bool) reloadAction {
	return ra.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`)
func (ra reloadAction) VisibleOn(value string) reloadAction {
	return ra.set("visibleOn", value)
}
