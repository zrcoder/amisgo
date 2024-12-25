package comp

// dropdownButton 表示下拉按钮渲染器。
type dropdownButton Schema

// DropdownButton 创建一个新的 DropdownButton 实例，并设置默认的 type 为 'dropdown-button'
func DropdownButton() dropdownButton {
	return make(dropdownButton).set("type", "dropdown-button")
}

func (d dropdownButton) set(key string, value any) dropdownButton {
	d[key] = value
	return d
}

// Align 设置对齐方式，可选值: left | right
func (d dropdownButton) Align(value string) dropdownButton {
	return d.set("align", value)
}

// Block 是否独占一行 `display: block`
func (d dropdownButton) Block(value bool) dropdownButton {
	return d.set("block", value)
}

// Body 内容区域
func (d dropdownButton) Body(value ...any) dropdownButton {
	return d.set("body", value)
}

// BtnClassName 按钮的 CSS 类名
func (d dropdownButton) BtnClassName(value string) dropdownButton {
	return d.set("btnClassName", value)
}

// Buttons 按钮集合，支持分组
func (d dropdownButton) Buttons(value string) dropdownButton {
	return d.set("buttons", value)
}

// ClassName 容器 CSS 类名
func (d dropdownButton) ClassName(value string) dropdownButton {
	return d.set("className", value)
}

// CloseOnClick 点击内容是否关闭
func (d dropdownButton) CloseOnClick(value bool) dropdownButton {
	return d.set("closeOnClick", value)
}

// CloseOnOutside 点击外部是否关闭
func (d dropdownButton) CloseOnOutside(value bool) dropdownButton {
	return d.set("closeOnOutside", value)
}

// Disabled 是否禁用
func (d dropdownButton) Disabled(value bool) dropdownButton {
	return d.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (d dropdownButton) DisabledOn(value string) dropdownButton {
	return d.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (d dropdownButton) EditorSetting(value string) dropdownButton {
	return d.set("editorSetting", value)
}

// Hidden 是否隐藏
func (d dropdownButton) Hidden(value bool) dropdownButton {
	return d.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (d dropdownButton) HiddenOn(value string) dropdownButton {
	return d.set("hiddenOn", value)
}

// HideCaret 是否显示下拉按钮的箭头
func (d dropdownButton) HideCaret(value bool) dropdownButton {
	return d.set("hideCaret", value)
}

// IconOnly 是否只显示图标
func (d dropdownButton) IconOnly(value bool) dropdownButton {
	return d.set("iconOnly", value)
}

// ID 组件唯一 ID，主要用于日志采集
func (d dropdownButton) ID(value string) dropdownButton {
	return d.set("id", value)
}

// Label 按钮文字
func (d dropdownButton) Label(value string) dropdownButton {
	return d.set("label", value)
}

// Level 按钮级别，样式，可选值: info | success | danger | warning | primary | link
func (d dropdownButton) Level(value string) dropdownButton {
	return d.set("level", value)
}

// MenuClassName 菜单 CSS 样式
func (d dropdownButton) MenuClassName(value string) dropdownButton {
	return d.set("menuClassName", value)
}

// OnEvent 事件动作配置
func (d dropdownButton) OnEvent(value any) dropdownButton {
	return d.set("onEvent", value)
}

// OverlayPlacement 层叠位置
func (d dropdownButton) OverlayPlacement(value string) dropdownButton {
	return d.set("overlayPlacement", value)
}

// RightIcon 右侧图标
func (d dropdownButton) RightIcon(value string) dropdownButton {
	return d.set("rightIcon", value)
}

// Size 按钮大小，可选值: xs | sm | md | lg
func (d dropdownButton) Size(value string) dropdownButton {
	return d.set("size", value)
}

// Static 是否静态展示
func (d dropdownButton) Static(value bool) dropdownButton {
	return d.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (d dropdownButton) StaticClassName(value string) dropdownButton {
	return d.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项值类名
func (d dropdownButton) StaticInputClassName(value string) dropdownButton {
	return d.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项标签类名
func (d dropdownButton) StaticLabelClassName(value string) dropdownButton {
	return d.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (d dropdownButton) StaticOn(value string) dropdownButton {
	return d.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (d dropdownButton) StaticPlaceholder(value string) dropdownButton {
	return d.set("staticPlaceholder", value)
}

// StaticSchema 静态展示 schema
func (d dropdownButton) StaticSchema(value string) dropdownButton {
	return d.set("staticSchema", value)
}

// Style 组件样式
func (d dropdownButton) Style(value any) dropdownButton {
	return d.set("style", value)
}

// TestIdBuilder 测试 ID 构造函数
func (d dropdownButton) TestIdBuilder(value string) dropdownButton {
	return d.set("testIdBuilder", value)
}

// Testid 测试 ID
func (d dropdownButton) Testid(value string) dropdownButton {
	return d.set("testid", value)
}

// Trigger 触发条件，默认是 click，可选值: click | hover
func (d dropdownButton) Trigger(value string) dropdownButton {
	return d.set("trigger", value)
}

// UseMobileUI 组件级别用来关闭移动端样式
func (d dropdownButton) UseMobileUI(value bool) dropdownButton {
	return d.set("useMobileUI", value)
}

// Visible 是否显示
func (d dropdownButton) Visible(value bool) dropdownButton {
	return d.set("visible", value)
}

// VisibleOn 是否显示表达式
func (d dropdownButton) VisibleOn(value string) dropdownButton {
	return d.set("visibleOn", value)
}
