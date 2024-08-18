package comp

// grid 格子布局渲染器
type grid schema

// Grid 创建一个新的 Grid 实例
func Grid() grid {
	return make(grid).set("type", "grid")
}

func (g grid) set(key string, value interface{}) grid {
	g[key] = value
	return g
}

// Align 水平对齐方式
func (g grid) Align(value string) grid {
	return g.set("align", value)
}

// ClassName 容器 css 类名
func (g grid) ClassName(value string) grid {
	return g.set("className", value)
}

// Columns 列集合
func (g grid) Columns(value string) grid {
	return g.set("columns", value)
}

// Disabled 是否禁用
func (g grid) Disabled(value bool) grid {
	return g.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (g grid) DisabledOn(value string) grid {
	return g.set("disabledOn", value)
}

// EditorSetting 编辑器配置
func (g grid) EditorSetting(value string) grid {
	return g.set("editorSetting", value)
}

// Gap 水平间距
func (g grid) Gap(value string) grid {
	return g.set("gap", value)
}

// Hidden 是否隐藏
func (g grid) Hidden(value bool) grid {
	return g.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (g grid) HiddenOn(value string) grid {
	return g.set("hiddenOn", value)
}

// ID 组件唯一 id
func (g grid) ID(value string) grid {
	return g.set("id", value)
}

// OnEvent 事件动作配置
func (g grid) OnEvent(value string) grid {
	return g.set("onEvent", value)
}

// Static 是否静态展示
func (g grid) Static(value bool) grid {
	return g.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (g grid) StaticClassName(value string) grid {
	return g.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (g grid) StaticInputClassName(value string) grid {
	return g.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (g grid) StaticLabelClassName(value string) grid {
	return g.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (g grid) StaticOn(value string) grid {
	return g.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (g grid) StaticPlaceholder(value string) grid {
	return g.set("staticPlaceholder", value)
}

// StaticSchema 静态展示模式的 schema
func (g grid) StaticSchema(value string) grid {
	return g.set("staticSchema", value)
}

// Style 组件样式
func (g grid) Style(value string) grid {
	return g.set("style", value)
}

// TestIdBuilder 测试 id 构建器
func (g grid) TestIdBuilder(value string) grid {
	return g.set("testIdBuilder", value)
}

// Testid 测试 id
func (g grid) Testid(value string) grid {
	return g.set("testid", value)
}

// UseMobileUI 组件级别关闭移动端样式
func (g grid) UseMobileUI(value bool) grid {
	return g.set("useMobileUI", value)
}

// Valign 垂直对齐方式
func (g grid) Valign(value string) grid {
	return g.set("valign", value)
}

// Visible 是否显示
func (g grid) Visible(value bool) grid {
	return g.set("visible", value)
}

// VisibleOn 是否显示表达式
func (g grid) VisibleOn(value string) grid {
	return g.set("visibleOn", value)
}
