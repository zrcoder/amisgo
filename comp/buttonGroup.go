package comp

// buttonGroup 代表按钮组渲染器
type buttonGroup Schema

// ButtonGroup 创建一个新的 ButtonGroup 实例
func ButtonGroup() buttonGroup {
	return make(buttonGroup).set("type", "button-group")
}

func (br buttonGroup) set(key string, value any) buttonGroup {
	br[key] = value
	return br
}

// BtnActiveClassName 设置按钮激活状态的类名
func (bg buttonGroup) BtnActiveClassName(value string) buttonGroup {
	return bg.set("btnActiveClassName", value)
}

// BtnActiveLevel 设置按钮选中的样式级别
// 'link' | 'primary' | 'secondary' | 'info'|'success' | 'warning' | 'danger' | 'light'| 'dark' | 'default'
func (bg buttonGroup) BtnActiveLevel(value string) buttonGroup {
	return bg.set("btnActiveLevel", value)
}

// BtnClassName 设置按钮的 CSS 类名
func (bg buttonGroup) BtnClassName(value string) buttonGroup {
	return bg.set("btnClassName", value)
}

// BtnLevel 设置按钮样式级别
// 'link' | 'primary' | 'secondary' | 'info'|'success' | 'warning' | 'danger' | 'light'| 'dark' | 'default'
func (bg buttonGroup) BtnLevel(value string) buttonGroup {
	return bg.set("btnLevel", value)
}

// Buttons 设置按钮集合
func (bg buttonGroup) Buttons(value ...any) buttonGroup {
	return bg.set("buttons", value)
}

// ClassName 设置容器的 CSS 类名
func (bg buttonGroup) ClassName(value string) buttonGroup {
	return bg.set("className", value)
}

// Disabled 设置是否禁用
func (bg buttonGroup) Disabled(value bool) buttonGroup {
	return bg.set("disabled", value)
}

// DisabledOn 设置通过 JS 表达式来配置禁用状态
func (bg buttonGroup) DisabledOn(value string) buttonGroup {
	return bg.set("disabledOn", value)
}

// EditorSetting 设置编辑器配置
func (bg buttonGroup) EditorSetting(value string) buttonGroup {
	return bg.set("editorSetting", value)
}

// Hidden 设置是否隐藏
func (bg buttonGroup) Hidden(value bool) buttonGroup {
	return bg.set("hidden", value)
}

// HiddenOn 设置是否隐藏的表达式
func (bg buttonGroup) HiddenOn(value string) buttonGroup {
	return bg.set("hiddenOn", value)
}

// Id 设置组件唯一 ID
func (bg buttonGroup) ID(value string) buttonGroup {
	return bg.set("id", value)
}

// OnEvent 设置事件动作配置
func (bg buttonGroup) OnEvent(value any) buttonGroup {
	return bg.set("onEvent", value)
}

// Size 设置按钮大小
func (bg buttonGroup) Size(value string) buttonGroup {
	return bg.set("size", value)
}

// Static 设置是否静态展示
func (bg buttonGroup) Static(value bool) buttonGroup {
	return bg.set("static", value)
}

// StaticClassName 设置静态展示表单项的 CSS 类名
func (bg buttonGroup) StaticClassName(value string) buttonGroup {
	return bg.set("staticClassName", value)
}

// StaticInputClassName 设置静态展示表单项的 Value 类名
func (bg buttonGroup) StaticInputClassName(value string) buttonGroup {
	return bg.set("staticInputClassName", value)
}

// StaticLabelClassName 设置静态展示表单项的 Label 类名
func (bg buttonGroup) StaticLabelClassName(value string) buttonGroup {
	return bg.set("staticLabelClassName", value)
}

// StaticOn 设置静态展示的表达式
func (bg buttonGroup) StaticOn(value string) buttonGroup {
	return bg.set("staticOn", value)
}

// StaticPlaceholder 设置静态展示的空值占位
func (bg buttonGroup) StaticPlaceholder(value string) buttonGroup {
	return bg.set("staticPlaceholder", value)
}

// StaticSchema 设置静态展示的 Schema
func (bg buttonGroup) StaticSchema(value string) buttonGroup {
	return bg.set("staticSchema", value)
}

// Style 设置组件样式
func (bg buttonGroup) Style(value any) buttonGroup {
	return bg.set("style", value)
}

// TestIdBuilder 设置测试 ID 构建器
func (bg buttonGroup) TestIdBuilder(value string) buttonGroup {
	return bg.set("testIdBuilder", value)
}

// Testid 设置测试 ID
func (bg buttonGroup) Testid(value string) buttonGroup {
	return bg.set("testid", value)
}

// Tiled 设置平铺展示
func (bg buttonGroup) Tiled(value bool) buttonGroup {
	return bg.set("tiled", value)
}

// UseMobileUI 设置是否使用移动端 UI 样式
func (bg buttonGroup) UseMobileUI(value bool) buttonGroup {
	return bg.set("useMobileUI", value)
}

// Vertical 设置是否垂直展示
func (bg buttonGroup) Vertical(value bool) buttonGroup {
	return bg.set("vertical", value)
}

// Visible 设置是否显示
func (bg buttonGroup) Visible(value bool) buttonGroup {
	return bg.set("visible", value)
}

// VisibleOn 设置通过 JS 表达式来配置当前表单项是否显示
func (bg buttonGroup) VisibleOn(value string) buttonGroup {
	return bg.set("visibleOn", value)
}
