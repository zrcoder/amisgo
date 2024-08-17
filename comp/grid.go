package comp

// Grid 格子布局渲染器
type Grid Schema

// NewGrid 创建一个新的 Grid 实例
func NewGrid() Grid {
	return make(Grid).set("type", "grid")
}

// set 设置属性并返回当前对象
func (g Grid) set(key string, value interface{}) Grid {
	g[key] = value
	return g
}

// Align 水平对齐方式
func (g Grid) Align(value string) Grid {
	return g.set("align", value)
}

// ClassName 容器 css 类名
func (g Grid) ClassName(value string) Grid {
	return g.set("className", value)
}

// Columns 列集合
func (g Grid) Columns(value string) Grid {
	return g.set("columns", value)
}

// Disabled 是否禁用
func (g Grid) Disabled(value bool) Grid {
	return g.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (g Grid) DisabledOn(value string) Grid {
	return g.set("disabledOn", value)
}

// EditorSetting 编辑器配置
func (g Grid) EditorSetting(value string) Grid {
	return g.set("editorSetting", value)
}

// Gap 水平间距
func (g Grid) Gap(value string) Grid {
	return g.set("gap", value)
}

// Hidden 是否隐藏
func (g Grid) Hidden(value bool) Grid {
	return g.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (g Grid) HiddenOn(value string) Grid {
	return g.set("hiddenOn", value)
}

// ID 组件唯一 id
func (g Grid) ID(value string) Grid {
	return g.set("id", value)
}

// OnEvent 事件动作配置
func (g Grid) OnEvent(value string) Grid {
	return g.set("onEvent", value)
}

// Static 是否静态展示
func (g Grid) Static(value bool) Grid {
	return g.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (g Grid) StaticClassName(value string) Grid {
	return g.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (g Grid) StaticInputClassName(value string) Grid {
	return g.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (g Grid) StaticLabelClassName(value string) Grid {
	return g.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (g Grid) StaticOn(value string) Grid {
	return g.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (g Grid) StaticPlaceholder(value string) Grid {
	return g.set("staticPlaceholder", value)
}

// StaticSchema 静态展示模式的 schema
func (g Grid) StaticSchema(value string) Grid {
	return g.set("staticSchema", value)
}

// Style 组件样式
func (g Grid) Style(value string) Grid {
	return g.set("style", value)
}

// TestIdBuilder 测试 id 构建器
func (g Grid) TestIdBuilder(value string) Grid {
	return g.set("testIdBuilder", value)
}

// Testid 测试 id
func (g Grid) Testid(value string) Grid {
	return g.set("testid", value)
}

// Type 指定为 Grid 格子布局渲染器
func (g Grid) Type(value string) Grid {
	return g.set("type", value)
}

// UseMobileUI 组件级别关闭移动端样式
func (g Grid) UseMobileUI(value bool) Grid {
	return g.set("useMobileUI", value)
}

// Valign 垂直对齐方式
func (g Grid) Valign(value string) Grid {
	return g.set("valign", value)
}

// Visible 是否显示
func (g Grid) Visible(value bool) Grid {
	return g.set("visible", value)
}

// VisibleOn 是否显示表达式
func (g Grid) VisibleOn(value string) Grid {
	return g.set("visibleOn", value)
}
