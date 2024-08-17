package comp

// CollapseGroup 折叠渲染器
type CollapseGroup BaseRenderer

// NewCollapseGroup 创建一个新的 CollapseGroup 实例
func NewCollapseGroup() CollapseGroup {
	res := CollapseGroup(make(BaseRenderer))
	return res.set("type", "collapse-group")
}

func (c CollapseGroup) set(key string, value interface{}) CollapseGroup {
	c[key] = value
	return c
}

// Accordion 设置手风琴模式
func (c CollapseGroup) Accordion(value bool) CollapseGroup {
	return c.set("accordion", value)
}

// ActiveKey 设置激活面板
func (c CollapseGroup) ActiveKey(value string) CollapseGroup {
	return c.set("activeKey", value)
}

// Body 设置内容区域
func (c CollapseGroup) Body(value ...interface{}) CollapseGroup {
	return c.set("body", value)
}

// ClassName 设置容器 css 类名
func (c CollapseGroup) ClassName(value string) CollapseGroup {
	return c.set("className", value)
}

// Disabled 设置是否禁用
func (c CollapseGroup) Disabled(value bool) CollapseGroup {
	return c.set("disabled", value)
}

// DisabledOn 设置是否禁用表达式
func (c CollapseGroup) DisabledOn(value string) CollapseGroup {
	return c.set("disabledOn", value)
}

// EditorSetting 设置编辑器配置
func (c CollapseGroup) EditorSetting(value string) CollapseGroup {
	return c.set("editorSetting", value)
}

// EnableFieldSetStyle 设置是否启用 FieldSet 样式
func (c CollapseGroup) EnableFieldSetStyle(value bool) CollapseGroup {
	return c.set("enableFieldSetStyle", value)
}

// ExpandIcon 设置自定义切换图标
func (c CollapseGroup) ExpandIcon(value string) CollapseGroup {
	return c.set("expandIcon", value)
}

// ExpandIconPosition 设置图标位置
func (c CollapseGroup) ExpandIconPosition(value string) CollapseGroup {
	return c.set("expandIconPosition", value)
}

// Hidden 设置是否隐藏
func (c CollapseGroup) Hidden(value bool) CollapseGroup {
	return c.set("hidden", value)
}

// HiddenOn 设置是否隐藏表达式
func (c CollapseGroup) HiddenOn(value string) CollapseGroup {
	return c.set("hiddenOn", value)
}

// ID 设置组件唯一 id
func (c CollapseGroup) ID(value string) CollapseGroup {
	return c.set("id", value)
}

// OnEvent 设置事件动作配置
func (c CollapseGroup) OnEvent(value string) CollapseGroup {
	return c.set("onEvent", value)
}

// Static 设置是否静态展示
func (c CollapseGroup) Static(value bool) CollapseGroup {
	return c.set("static", value)
}

// StaticClassName 设置静态展示表单项类名
func (c CollapseGroup) StaticClassName(value string) CollapseGroup {
	return c.set("staticClassName", value)
}

// StaticInputClassName 设置静态展示表单项Value类名
func (c CollapseGroup) StaticInputClassName(value string) CollapseGroup {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName 设置静态展示表单项Label类名
func (c CollapseGroup) StaticLabelClassName(value string) CollapseGroup {
	return c.set("staticLabelClassName", value)
}

// StaticOn 设置是否静态展示表达式
func (c CollapseGroup) StaticOn(value string) CollapseGroup {
	return c.set("staticOn", value)
}

// StaticPlaceholder 设置静态展示空值占位
func (c CollapseGroup) StaticPlaceholder(value string) CollapseGroup {
	return c.set("staticPlaceholder", value)
}

// StaticSchema 设置静态展示模式
func (c CollapseGroup) StaticSchema(value string) CollapseGroup {
	return c.set("staticSchema", value)
}

// Style 设置组件样式
func (c CollapseGroup) Style(value string) CollapseGroup {
	return c.set("style", value)
}

// TestIdBuilder 设置测试 ID 构建器
func (c CollapseGroup) TestIdBuilder(value string) CollapseGroup {
	return c.set("testIdBuilder", value)
}

// Testid 设置测试 ID
func (c CollapseGroup) Testid(value string) CollapseGroup {
	return c.set("testid", value)
}

// UseMobileUI 设置是否使用移动端样式
func (c CollapseGroup) UseMobileUI(value bool) CollapseGroup {
	return c.set("useMobileUI", value)
}

// Visible 设置是否显示
func (c CollapseGroup) Visible(value bool) CollapseGroup {
	return c.set("visible", value)
}

// VisibleOn 设置显示条件表达式
func (c CollapseGroup) VisibleOn(value string) CollapseGroup {
	return c.set("visibleOn", value)
}
