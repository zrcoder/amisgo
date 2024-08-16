package renderers

// ButtonGroup 代表按钮组渲染器
type ButtonGroup struct {
	*BaseRenderer
}

// NewButtonGroup 创建一个新的 ButtonGroup 实例
func NewButtonGroup() *ButtonGroup {
	bg := &ButtonGroup{BaseRenderer: NewBaseRenderer()}
	bg.set("type", "button-group")
	return bg
}

func (br *ButtonGroup) set(key string, value interface{}) *ButtonGroup {
	br.BaseRenderer.set(key, value)
	return br
}

// BtnActiveClassName 设置按钮激活状态的类名
func (bg *ButtonGroup) BtnActiveClassName(value string) *ButtonGroup {
	return bg.set("btnActiveClassName", value)
}

// BtnActiveLevel 设置按钮选中的样式级别
func (bg *ButtonGroup) BtnActiveLevel(value string) *ButtonGroup {
	return bg.set("btnActiveLevel", value)
}

// BtnClassName 设置按钮的 CSS 类名
func (bg *ButtonGroup) BtnClassName(value string) *ButtonGroup {
	return bg.set("btnClassName", value)
}

// BtnLevel 设置按钮样式级别
func (bg *ButtonGroup) BtnLevel(value string) *ButtonGroup {
	return bg.set("btnLevel", value)
}

// Buttons 设置按钮集合
func (bg *ButtonGroup) Buttons(value string) *ButtonGroup {
	return bg.set("buttons", value)
}

// ClassName 设置容器的 CSS 类名
func (bg *ButtonGroup) ClassName(value string) *ButtonGroup {
	return bg.set("className", value)
}

// Disabled 设置是否禁用
func (bg *ButtonGroup) Disabled(value bool) *ButtonGroup {
	return bg.set("disabled", value)
}

// DisabledOn 设置通过 JS 表达式来配置禁用状态
func (bg *ButtonGroup) DisabledOn(value string) *ButtonGroup {
	return bg.set("disabledOn", value)
}

// EditorSetting 设置编辑器配置
func (bg *ButtonGroup) EditorSetting(value string) *ButtonGroup {
	return bg.set("editorSetting", value)
}

// Hidden 设置是否隐藏
func (bg *ButtonGroup) Hidden(value bool) *ButtonGroup {
	return bg.set("hidden", value)
}

// HiddenOn 设置是否隐藏的表达式
func (bg *ButtonGroup) HiddenOn(value string) *ButtonGroup {
	return bg.set("hiddenOn", value)
}

// Id 设置组件唯一 ID
func (bg *ButtonGroup) Id(value string) *ButtonGroup {
	return bg.set("id", value)
}

// OnEvent 设置事件动作配置
func (bg *ButtonGroup) OnEvent(value string) *ButtonGroup {
	return bg.set("onEvent", value)
}

// Size 设置按钮大小
func (bg *ButtonGroup) Size(value string) *ButtonGroup {
	return bg.set("size", value)
}

// Static 设置是否静态展示
func (bg *ButtonGroup) Static(value bool) *ButtonGroup {
	return bg.set("static", value)
}

// StaticClassName 设置静态展示表单项的 CSS 类名
func (bg *ButtonGroup) StaticClassName(value string) *ButtonGroup {
	return bg.set("staticClassName", value)
}

// StaticInputClassName 设置静态展示表单项的 Value 类名
func (bg *ButtonGroup) StaticInputClassName(value string) *ButtonGroup {
	return bg.set("staticInputClassName", value)
}

// StaticLabelClassName 设置静态展示表单项的 Label 类名
func (bg *ButtonGroup) StaticLabelClassName(value string) *ButtonGroup {
	return bg.set("staticLabelClassName", value)
}

// StaticOn 设置静态展示的表达式
func (bg *ButtonGroup) StaticOn(value string) *ButtonGroup {
	return bg.set("staticOn", value)
}

// StaticPlaceholder 设置静态展示的空值占位
func (bg *ButtonGroup) StaticPlaceholder(value string) *ButtonGroup {
	return bg.set("staticPlaceholder", value)
}

// StaticSchema 设置静态展示的 Schema
func (bg *ButtonGroup) StaticSchema(value string) *ButtonGroup {
	return bg.set("staticSchema", value)
}

// Style 设置组件样式
func (bg *ButtonGroup) Style(value string) *ButtonGroup {
	return bg.set("style", value)
}

// TestIdBuilder 设置测试 ID 构建器
func (bg *ButtonGroup) TestIdBuilder(value string) *ButtonGroup {
	return bg.set("testIdBuilder", value)
}

// Testid 设置测试 ID
func (bg *ButtonGroup) Testid(value string) *ButtonGroup {
	return bg.set("testid", value)
}

// Tiled 设置平铺展示
func (bg *ButtonGroup) Tiled(value bool) *ButtonGroup {
	return bg.set("tiled", value)
}

// UseMobileUI 设置是否使用移动端 UI 样式
func (bg *ButtonGroup) UseMobileUI(value bool) *ButtonGroup {
	return bg.set("useMobileUI", value)
}

// Vertical 设置是否垂直展示
func (bg *ButtonGroup) Vertical(value bool) *ButtonGroup {
	return bg.set("vertical", value)
}

// Visible 设置是否显示
func (bg *ButtonGroup) Visible(value bool) *ButtonGroup {
	return bg.set("visible", value)
}

// VisibleOn 设置通过 JS 表达式来配置当前表单项是否显示
func (bg *ButtonGroup) VisibleOn(value string) *ButtonGroup {
	return bg.set("visibleOn", value)
}
