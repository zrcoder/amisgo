package comp

// DropdownButton 表示下拉按钮渲染器。
type DropdownButton Schema

// NewDropdownButton 创建一个新的 DropdownButton 实例，并设置默认的 type 为 'dropdown-button'
func NewDropdownButton() DropdownButton {
	return make(DropdownButton).set("type", "dropdown-button")
}

func (d DropdownButton) set(key string, value interface{}) DropdownButton {
	d[key] = value
	return d
}

// Align 设置对齐方式，可选值: left | right
func (d DropdownButton) Align(value string) DropdownButton {
	return d.set("align", value)
}

// Block 是否独占一行 `display: block`
func (d DropdownButton) Block(value bool) DropdownButton {
	return d.set("block", value)
}

// Body 内容区域
func (d DropdownButton) Body(value ...interface{}) DropdownButton {
	return d.set("body", value)
}

// BtnClassName 按钮的 CSS 类名
func (d DropdownButton) BtnClassName(value string) DropdownButton {
	return d.set("btnClassName", value)
}

// Buttons 按钮集合，支持分组
func (d DropdownButton) Buttons(value string) DropdownButton {
	return d.set("buttons", value)
}

// ClassName 容器 CSS 类名
func (d DropdownButton) ClassName(value string) DropdownButton {
	return d.set("className", value)
}

// CloseOnClick 点击内容是否关闭
func (d DropdownButton) CloseOnClick(value bool) DropdownButton {
	return d.set("closeOnClick", value)
}

// CloseOnOutside 点击外部是否关闭
func (d DropdownButton) CloseOnOutside(value bool) DropdownButton {
	return d.set("closeOnOutside", value)
}

// Disabled 是否禁用
func (d DropdownButton) Disabled(value bool) DropdownButton {
	return d.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (d DropdownButton) DisabledOn(value string) DropdownButton {
	return d.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (d DropdownButton) EditorSetting(value string) DropdownButton {
	return d.set("editorSetting", value)
}

// Hidden 是否隐藏
func (d DropdownButton) Hidden(value bool) DropdownButton {
	return d.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (d DropdownButton) HiddenOn(value string) DropdownButton {
	return d.set("hiddenOn", value)
}

// HideCaret 是否显示下拉按钮的箭头
func (d DropdownButton) HideCaret(value bool) DropdownButton {
	return d.set("hideCaret", value)
}

// IconOnly 是否只显示图标
func (d DropdownButton) IconOnly(value bool) DropdownButton {
	return d.set("iconOnly", value)
}

// ID 组件唯一 ID，主要用于日志采集
func (d DropdownButton) ID(value string) DropdownButton {
	return d.set("id", value)
}

// Label 按钮文字
func (d DropdownButton) Label(value string) DropdownButton {
	return d.set("label", value)
}

// Level 按钮级别，样式，可选值: info | success | danger | warning | primary | link
func (d DropdownButton) Level(value string) DropdownButton {
	return d.set("level", value)
}

// MenuClassName 菜单 CSS 样式
func (d DropdownButton) MenuClassName(value string) DropdownButton {
	return d.set("menuClassName", value)
}

// OnEvent 事件动作配置
func (d DropdownButton) OnEvent(value string) DropdownButton {
	return d.set("onEvent", value)
}

// OverlayPlacement 层叠位置
func (d DropdownButton) OverlayPlacement(value string) DropdownButton {
	return d.set("overlayPlacement", value)
}

// RightIcon 右侧图标
func (d DropdownButton) RightIcon(value string) DropdownButton {
	return d.set("rightIcon", value)
}

// Size 按钮大小，可选值: xs | sm | md | lg
func (d DropdownButton) Size(value string) DropdownButton {
	return d.set("size", value)
}

// Static 是否静态展示
func (d DropdownButton) Static(value bool) DropdownButton {
	return d.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (d DropdownButton) StaticClassName(value string) DropdownButton {
	return d.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项值类名
func (d DropdownButton) StaticInputClassName(value string) DropdownButton {
	return d.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项标签类名
func (d DropdownButton) StaticLabelClassName(value string) DropdownButton {
	return d.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (d DropdownButton) StaticOn(value string) DropdownButton {
	return d.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (d DropdownButton) StaticPlaceholder(value string) DropdownButton {
	return d.set("staticPlaceholder", value)
}

// StaticSchema 静态展示 schema
func (d DropdownButton) StaticSchema(value string) DropdownButton {
	return d.set("staticSchema", value)
}

// Style 组件样式
func (d DropdownButton) Style(value string) DropdownButton {
	return d.set("style", value)
}

// TestIdBuilder 测试 ID 构造函数
func (d DropdownButton) TestIdBuilder(value string) DropdownButton {
	return d.set("testIdBuilder", value)
}

// Testid 测试 ID
func (d DropdownButton) Testid(value string) DropdownButton {
	return d.set("testid", value)
}

// Trigger 触发条件，默认是 click，可选值: click | hover
func (d DropdownButton) Trigger(value string) DropdownButton {
	return d.set("trigger", value)
}

// UseMobileUI 组件级别用来关闭移动端样式
func (d DropdownButton) UseMobileUI(value bool) DropdownButton {
	return d.set("useMobileUI", value)
}

// Visible 是否显示
func (d DropdownButton) Visible(value bool) DropdownButton {
	return d.set("visible", value)
}

// VisibleOn 是否显示表达式
func (d DropdownButton) VisibleOn(value string) DropdownButton {
	return d.set("visibleOn", value)
}
