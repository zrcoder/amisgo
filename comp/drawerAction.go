package comp

import "github.com/zrcoder/amisgo/model"

// drawerAction 表示一个抽出式弹框中的操作按钮。
type drawerAction schema

// DrawerAction 创建一个新的 DrawerAction 实例，并设置默认的 type 和 actionType
func DrawerAction() drawerAction {
	return make(drawerAction).set("type", "button").set("actionType", "drawer")
}

func (d drawerAction) set(key string, value any) drawerAction {
	d[key] = value
	return d
}

// ActiveClassName 激活状态时的类名
func (d drawerAction) ActiveClassName(value string) drawerAction {
	return d.set("activeClassName", value)
}

// ActiveLevel 激活状态时的样式
func (d drawerAction) ActiveLevel(value string) drawerAction {
	return d.set("activeLevel", value)
}

// Align 设置对齐方式
func (d drawerAction) Align(value string) drawerAction {
	return d.set("align", value)
}

// Badge 角标
func (d drawerAction) Badge(value string) drawerAction {
	return d.set("badge", value)
}

// Block 是否为块状展示，默认为内联
func (d drawerAction) Block(value bool) drawerAction {
	return d.set("block", value)
}

// Body 子内容
func (d drawerAction) Body(value ...any) drawerAction {
	return d.set("body", value)
}

// ClassName 容器 css 类名
func (d drawerAction) ClassName(value string) drawerAction {
	return d.set("className", value)
}

// Close 如果按钮在弹框中，可以配置这个动作完成后是否关闭弹窗，或者指定关闭目标弹框
func (d drawerAction) Close(value string) drawerAction {
	return d.set("close", value)
}

// ConfirmText 提示文字
func (d drawerAction) ConfirmText(value string) drawerAction {
	return d.set("confirmText", value)
}

// ConfirmTitle 确认弹窗标题
func (d drawerAction) ConfirmTitle(value any) drawerAction {
	return d.set("confirmTitle", value)
}

// CountDown 点击后的禁止倒计时（秒）
func (d drawerAction) CountDown(value string) drawerAction {
	return d.set("countDown", value)
}

// CountDownTpl 倒计时文字自定义
func (d drawerAction) CountDownTpl(value string) drawerAction {
	return d.set("countDownTpl", value)
}

// Data 数据映射
func (d drawerAction) Data(value model.Data) drawerAction {
	return d.set("data", value)
}

// Disabled 是否禁用
func (d drawerAction) Disabled(value bool) drawerAction {
	return d.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (d drawerAction) DisabledOn(value string) drawerAction {
	return d.set("disabledOn", value)
}

// DisabledTip 禁用时的文案提示
func (d drawerAction) DisabledTip(value string) drawerAction {
	return d.set("disabledTip", value)
}

// Drawer 抽出式弹框详情
func (d drawerAction) Drawer(value string) drawerAction {
	return d.set("drawer", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (d drawerAction) EditorSetting(value string) drawerAction {
	return d.set("editorSetting", value)
}

// Hidden 是否隐藏
func (d drawerAction) Hidden(value bool) drawerAction {
	return d.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (d drawerAction) HiddenOn(value string) drawerAction {
	return d.set("hiddenOn", value)
}

// HotKey 键盘快捷键
func (d drawerAction) HotKey(value string) drawerAction {
	return d.set("hotKey", value)
}

// Icon 按钮图标
func (d drawerAction) Icon(value string) drawerAction {
	return d.set("icon", value)
}

// IconClassName icon 上的 css 类名
func (d drawerAction) IconClassName(value string) drawerAction {
	return d.set("iconClassName", value)
}

// ID 主要用于用户行为跟踪里区分是哪个按钮
func (d drawerAction) ID(value string) drawerAction {
	return d.set("id", value)
}

// Label 按钮文字
func (d drawerAction) Label(value string) drawerAction {
	return d.set("label", value)
}

// Level 按钮样式
func (d drawerAction) Level(value string) drawerAction {
	return d.set("level", value)
}

// LoadingClassName loading 上的 css 类名
func (d drawerAction) LoadingClassName(value string) drawerAction {
	return d.set("loadingClassName", value)
}

// LoadingOn 是否显示 loading 效果
func (d drawerAction) LoadingOn(value string) drawerAction {
	return d.set("loadingOn", value)
}

// MergeData 是否将弹框中数据 merge 到父级作用域
func (d drawerAction) MergeData(value bool) drawerAction {
	return d.set("mergeData", value)
}

// NextCondition 是否有下一个的表达式
func (d drawerAction) NextCondition(value string) drawerAction {
	return d.set("nextCondition", value)
}

// OnClick 自定义事件处理函数
func (d drawerAction) OnClick(value string) drawerAction {
	return d.set("onClick", value)
}

// OnEvent 事件动作配置
func (d drawerAction) OnEvent(value any) drawerAction {
	return d.set("onEvent", value)
}

// Primary 标识是否为主要按钮
func (d drawerAction) Primary(value bool) drawerAction {
	return d.set("primary", value)
}

// Redirect 重定向配置
func (d drawerAction) Redirect(value string) drawerAction {
	return d.set("redirect", value)
}

// Reload 配置刷新动作
func (d drawerAction) Reload(value string) drawerAction {
	return d.set("reload", value)
}

// RequireSelected 当按钮是批量操作按钮时，是否需要有勾选元素
func (d drawerAction) RequireSelected(value bool) drawerAction {
	return d.set("requireSelected", value)
}

// Required 如果按钮在 form 中，配置此属性会要求用户把指定的字段通过验证后才会触发行为
func (d drawerAction) Required(value string) drawerAction {
	return d.set("required", value)
}

// RightIcon 右侧按钮图标
func (d drawerAction) RightIcon(value string) drawerAction {
	return d.set("rightIcon", value)
}

// RightIconClassName 右侧 icon 上的 css 类名
func (d drawerAction) RightIconClassName(value string) drawerAction {
	return d.set("rightIconClassName", value)
}

// Size 按钮大小
func (d drawerAction) Size(value string) drawerAction {
	return d.set("size", value)
}

// Static 是否静态展示
func (d drawerAction) Static(value bool) drawerAction {
	return d.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (d drawerAction) StaticClassName(value string) drawerAction {
	return d.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (d drawerAction) StaticInputClassName(value string) drawerAction {
	return d.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (d drawerAction) StaticLabelClassName(value string) drawerAction {
	return d.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (d drawerAction) StaticOn(value string) drawerAction {
	return d.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (d drawerAction) StaticPlaceholder(value string) drawerAction {
	return d.set("staticPlaceholder", value)
}

// StaticSchema 静态展示 schema
func (d drawerAction) StaticSchema(value string) drawerAction {
	return d.set("staticSchema", value)
}

// Style 组件样式
func (d drawerAction) Style(value any) drawerAction {
	return d.set("style", value)
}

// Target 可以指定让谁来触发这个动作
func (d drawerAction) Target(value string) drawerAction {
	return d.set("target", value)
}

// TestIdBuilder 用于测试的 ID 构造函数
func (d drawerAction) TestIdBuilder(value string) drawerAction {
	return d.set("testIdBuilder", value)
}

// Testid 测试 ID
func (d drawerAction) Testid(value string) drawerAction {
	return d.set("testid", value)
}

// Tooltip 提示信息
func (d drawerAction) Tooltip(value string) drawerAction {
	return d.set("tooltip", value)
}

// TooltipPlacement 提示信息位置
func (d drawerAction) TooltipPlacement(value string) drawerAction {
	return d.set("tooltipPlacement", value)
}

// Type 指定按钮类型，支持 button、submit 或 reset 三种类型
func (d drawerAction) Type(value string) drawerAction {
	return d.set("type", value)
}

// UseMobileUI 组件级别用来关闭移动端样式
func (d drawerAction) UseMobileUI(value bool) drawerAction {
	return d.set("useMobileUI", value)
}

// Visible 是否显示
func (d drawerAction) Visible(value bool) drawerAction {
	return d.set("visible", value)
}

// VisibleOn 是否显示表达式
func (d drawerAction) VisibleOn(value string) drawerAction {
	return d.set("visibleOn", value)
}
