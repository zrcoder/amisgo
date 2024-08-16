package renderers

// ButtonToolbar 代表按钮工具条控件渲染器
type ButtonToolbar struct {
	*BaseRenderer
}

// NewButtonToolbar 创建一个新的 ButtonToolbar 实例
func NewButtonToolbar() *ButtonToolbar {
	bt := &ButtonToolbar{BaseRenderer: NewBaseRenderer()}
	bt.set("type", "button-toolbar")
	return bt
}

// Set 覆盖 BaseRenderer 的 Set 方法，返回 *ButtonToolbar
func (bt *ButtonToolbar) set(key string, value interface{}) *ButtonToolbar {
	bt.BaseRenderer.set(key, value)
	return bt
}

// Buttons 设置按钮集合
func (bt *ButtonToolbar) Buttons(value string) *ButtonToolbar {
	return bt.set("buttons", value)
}

// ClassName 设置容器的 CSS 类名
func (bt *ButtonToolbar) ClassName(value string) *ButtonToolbar {
	return bt.set("className", value)
}

// Disabled 设置是否禁用
func (bt *ButtonToolbar) Disabled(value bool) *ButtonToolbar {
	return bt.set("disabled", value)
}

// DisabledOn 设置禁用表达式
func (bt *ButtonToolbar) DisabledOn(value string) *ButtonToolbar {
	return bt.set("disabledOn", value)
}

// EditorSetting 设置编辑器配置
func (bt *ButtonToolbar) EditorSetting(value string) *ButtonToolbar {
	return bt.set("editorSetting", value)
}

// Hidden 设置是否隐藏
func (bt *ButtonToolbar) Hidden(value bool) *ButtonToolbar {
	return bt.set("hidden", value)
}

// HiddenOn 设置隐藏表达式
func (bt *ButtonToolbar) HiddenOn(value string) *ButtonToolbar {
	return bt.set("hiddenOn", value)
}

// Id 设置组件唯一 ID
func (bt *ButtonToolbar) Id(value string) *ButtonToolbar {
	return bt.set("id", value)
}

// OnEvent 设置事件动作配置
func (bt *ButtonToolbar) OnEvent(value string) *ButtonToolbar {
	return bt.set("onEvent", value)
}

// Static 设置是否静态展示
func (bt *ButtonToolbar) Static(value bool) *ButtonToolbar {
	return bt.set("static", value)
}

// StaticClassName 设置静态展示表单项类名
func (bt *ButtonToolbar) StaticClassName(value string) *ButtonToolbar {
	return bt.set("staticClassName", value)
}

// StaticInputClassName 设置静态展示表单项的值类名
func (bt *ButtonToolbar) StaticInputClassName(value string) *ButtonToolbar {
	return bt.set("staticInputClassName", value)
}

// StaticLabelClassName 设置静态展示表单项的标签类名
func (bt *ButtonToolbar) StaticLabelClassName(value string) *ButtonToolbar {
	return bt.set("staticLabelClassName", value)
}

// StaticOn 设置静态展示表达式
func (bt *ButtonToolbar) StaticOn(value string) *ButtonToolbar {
	return bt.set("staticOn", value)
}

// StaticPlaceholder 设置静态展示空值占位
func (bt *ButtonToolbar) StaticPlaceholder(value string) *ButtonToolbar {
	return bt.set("staticPlaceholder", value)
}

// StaticSchema 设置静态展示模式
func (bt *ButtonToolbar) StaticSchema(value string) *ButtonToolbar {
	return bt.set("staticSchema", value)
}

// Style 设置组件样式
func (bt *ButtonToolbar) Style(value string) *ButtonToolbar {
	return bt.set("style", value)
}

// TestIdBuilder 设置测试 ID 构建器
func (bt *ButtonToolbar) TestIdBuilder(value string) *ButtonToolbar {
	return bt.set("testIdBuilder", value)
}

// Testid 设置测试 ID
func (bt *ButtonToolbar) Testid(value string) *ButtonToolbar {
	return bt.set("testid", value)
}

// UseMobileUI 设置是否使用移动端 UI
func (bt *ButtonToolbar) UseMobileUI(value bool) *ButtonToolbar {
	return bt.set("useMobileUI", value)
}

// Visible 设置是否显示
func (bt *ButtonToolbar) Visible(value bool) *ButtonToolbar {
	return bt.set("visible", value)
}

// VisibleOn 设置显示表达式
func (bt *ButtonToolbar) VisibleOn(value string) *ButtonToolbar {
	return bt.set("visibleOn", value)
}
