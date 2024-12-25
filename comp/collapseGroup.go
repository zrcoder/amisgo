package comp

// collapseGroup 折叠渲染器
type collapseGroup Schema

// CollapseGroup 创建一个新的 CollapseGroup 实例
func CollapseGroup() collapseGroup {
	return make(collapseGroup).set("type", "collapse-group")
}

func (c collapseGroup) set(key string, value any) collapseGroup {
	c[key] = value
	return c
}

// Accordion 设置手风琴模式
func (c collapseGroup) Accordion(value bool) collapseGroup {
	return c.set("accordion", value)
}

// ActiveKey 设置激活面板
func (c collapseGroup) ActiveKey(value any) collapseGroup {
	return c.set("activeKey", value)
}

// Body 设置内容区域
func (c collapseGroup) Body(value ...any) collapseGroup {
	return c.set("body", value)
}

// ClassName 设置容器 css 类名
func (c collapseGroup) ClassName(value string) collapseGroup {
	return c.set("className", value)
}

// Disabled 设置是否禁用
func (c collapseGroup) Disabled(value bool) collapseGroup {
	return c.set("disabled", value)
}

// DisabledOn 设置是否禁用表达式
func (c collapseGroup) DisabledOn(value string) collapseGroup {
	return c.set("disabledOn", value)
}

// EditorSetting 设置编辑器配置
func (c collapseGroup) EditorSetting(value string) collapseGroup {
	return c.set("editorSetting", value)
}

// EnableFieldSetStyle 设置是否启用 FieldSet 样式
func (c collapseGroup) EnableFieldSetStyle(value bool) collapseGroup {
	return c.set("enableFieldSetStyle", value)
}

// ExpandIcon 设置自定义切换图标
func (c collapseGroup) ExpandIcon(value any) collapseGroup {
	return c.set("expandIcon", value)
}

// ExpandIconPosition 设置图标位置
func (c collapseGroup) ExpandIconPosition(value string) collapseGroup {
	return c.set("expandIconPosition", value)
}

// Hidden 设置是否隐藏
func (c collapseGroup) Hidden(value bool) collapseGroup {
	return c.set("hidden", value)
}

// HiddenOn 设置是否隐藏表达式
func (c collapseGroup) HiddenOn(value string) collapseGroup {
	return c.set("hiddenOn", value)
}

// ID 设置组件唯一 id
func (c collapseGroup) ID(value string) collapseGroup {
	return c.set("id", value)
}

// OnEvent 设置事件动作配置
func (c collapseGroup) OnEvent(value any) collapseGroup {
	return c.set("onEvent", value)
}

// Static 设置是否静态展示
func (c collapseGroup) Static(value bool) collapseGroup {
	return c.set("static", value)
}

// StaticClassName 设置静态展示表单项类名
func (c collapseGroup) StaticClassName(value string) collapseGroup {
	return c.set("staticClassName", value)
}

// StaticInputClassName 设置静态展示表单项Value类名
func (c collapseGroup) StaticInputClassName(value string) collapseGroup {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName 设置静态展示表单项Label类名
func (c collapseGroup) StaticLabelClassName(value string) collapseGroup {
	return c.set("staticLabelClassName", value)
}

// StaticOn 设置是否静态展示表达式
func (c collapseGroup) StaticOn(value string) collapseGroup {
	return c.set("staticOn", value)
}

// StaticPlaceholder 设置静态展示空值占位
func (c collapseGroup) StaticPlaceholder(value string) collapseGroup {
	return c.set("staticPlaceholder", value)
}

// StaticSchema 设置静态展示模式
func (c collapseGroup) StaticSchema(value string) collapseGroup {
	return c.set("staticSchema", value)
}

// Style 设置组件样式
func (c collapseGroup) Style(value any) collapseGroup {
	return c.set("style", value)
}

// TestIdBuilder 设置测试 ID 构建器
func (c collapseGroup) TestIdBuilder(value string) collapseGroup {
	return c.set("testIdBuilder", value)
}

// Testid 设置测试 ID
func (c collapseGroup) Testid(value string) collapseGroup {
	return c.set("testid", value)
}

// UseMobileUI 设置是否使用移动端样式
func (c collapseGroup) UseMobileUI(value bool) collapseGroup {
	return c.set("useMobileUI", value)
}

// Visible 设置是否显示
func (c collapseGroup) Visible(value bool) collapseGroup {
	return c.set("visible", value)
}

// VisibleOn 设置显示条件表达式
func (c collapseGroup) VisibleOn(value string) collapseGroup {
	return c.set("visibleOn", value)
}
