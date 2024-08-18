package comp

// buttonToolbar 代表按钮工具条控件渲染器
type buttonToolbar schema

// ButtonToolbar 创建一个新的 ButtonToolbar 实例
func ButtonToolbar() buttonToolbar {
	return make(buttonToolbar).set("type", "button-toolbar")
}

// Set 覆盖 BaseRenderer 的 Set 方法，返回 ButtonToolbar
func (bt buttonToolbar) set(key string, value interface{}) buttonToolbar {
	bt[key] = value
	return bt
}

// Buttons 设置按钮集合
func (bt buttonToolbar) Buttons(value string) buttonToolbar {
	return bt.set("buttons", value)
}

// ClassName 设置容器的 CSS 类名
func (bt buttonToolbar) ClassName(value string) buttonToolbar {
	return bt.set("className", value)
}

// Disabled 设置是否禁用
func (bt buttonToolbar) Disabled(value bool) buttonToolbar {
	return bt.set("disabled", value)
}

// DisabledOn 设置禁用表达式
func (bt buttonToolbar) DisabledOn(value string) buttonToolbar {
	return bt.set("disabledOn", value)
}

// EditorSetting 设置编辑器配置
func (bt buttonToolbar) EditorSetting(value string) buttonToolbar {
	return bt.set("editorSetting", value)
}

// Hidden 设置是否隐藏
func (bt buttonToolbar) Hidden(value bool) buttonToolbar {
	return bt.set("hidden", value)
}

// HiddenOn 设置隐藏表达式
func (bt buttonToolbar) HiddenOn(value string) buttonToolbar {
	return bt.set("hiddenOn", value)
}

// Id 设置组件唯一 ID
func (bt buttonToolbar) Id(value string) buttonToolbar {
	return bt.set("id", value)
}

// OnEvent 设置事件动作配置
func (bt buttonToolbar) OnEvent(value string) buttonToolbar {
	return bt.set("onEvent", value)
}

// Static 设置是否静态展示
func (bt buttonToolbar) Static(value bool) buttonToolbar {
	return bt.set("static", value)
}

// StaticClassName 设置静态展示表单项类名
func (bt buttonToolbar) StaticClassName(value string) buttonToolbar {
	return bt.set("staticClassName", value)
}

// StaticInputClassName 设置静态展示表单项的值类名
func (bt buttonToolbar) StaticInputClassName(value string) buttonToolbar {
	return bt.set("staticInputClassName", value)
}

// StaticLabelClassName 设置静态展示表单项的标签类名
func (bt buttonToolbar) StaticLabelClassName(value string) buttonToolbar {
	return bt.set("staticLabelClassName", value)
}

// StaticOn 设置静态展示表达式
func (bt buttonToolbar) StaticOn(value string) buttonToolbar {
	return bt.set("staticOn", value)
}

// StaticPlaceholder 设置静态展示空值占位
func (bt buttonToolbar) StaticPlaceholder(value string) buttonToolbar {
	return bt.set("staticPlaceholder", value)
}

// StaticSchema 设置静态展示模式
func (bt buttonToolbar) StaticSchema(value string) buttonToolbar {
	return bt.set("staticSchema", value)
}

// Style 设置组件样式
func (bt buttonToolbar) Style(value string) buttonToolbar {
	return bt.set("style", value)
}

// TestIdBuilder 设置测试 ID 构建器
func (bt buttonToolbar) TestIdBuilder(value string) buttonToolbar {
	return bt.set("testIdBuilder", value)
}

// Testid 设置测试 ID
func (bt buttonToolbar) Testid(value string) buttonToolbar {
	return bt.set("testid", value)
}

// UseMobileUI 设置是否使用移动端 UI
func (bt buttonToolbar) UseMobileUI(value bool) buttonToolbar {
	return bt.set("useMobileUI", value)
}

// Visible 设置是否显示
func (bt buttonToolbar) Visible(value bool) buttonToolbar {
	return bt.set("visible", value)
}

// VisibleOn 设置显示表达式
func (bt buttonToolbar) VisibleOn(value string) buttonToolbar {
	return bt.set("visibleOn", value)
}
