package comp

// grid2d 二维布局渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/grid-2d
type grid2d Schema

// Grid2D 创建一个新的 Grid2D 实例
func Grid2D() grid2d {
	return make(grid2d).set("type", "grid-2d")
}

func (g grid2d) set(key string, value any) grid2d {
	g[key] = value
	return g
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (g grid2d) ClassName(value string) grid2d {
	return g.set("gridClassName", value)
}

// Cols 列数量，默认是 12
func (g grid2d) Cols(value string) grid2d {
	return g.set("cols", value)
}

// Disabled 是否禁用
func (g grid2d) Disabled(value bool) grid2d {
	return g.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (g grid2d) DisabledOn(value string) grid2d {
	return g.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (g grid2d) EditorSetting(value string) grid2d {
	return g.set("editorSetting", value)
}

// Gap 格子间距，默认 0，包含行和列
func (g grid2d) Gap(value string) grid2d {
	return g.set("gap", value)
}

// RowGap 格子垂直间距
func (g grid2d) RowGap(value string) grid2d {
	return g.set("rowGap", value)
}

// Grids 每个格子的配置
func (g grid2d) Grids(value ...MGridItem) grid2d {
	return g.set("grids", value)
}

// Hidden 是否隐藏
func (g grid2d) Hidden(value bool) grid2d {
	return g.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (g grid2d) HiddenOn(value string) grid2d {
	return g.set("hiddenOn", value)
}

// ID 组件唯一 id，主要用于日志采集
func (g grid2d) ID(value string) grid2d {
	return g.set("id", value)
}

// OnEvent 事件动作配置
func (g grid2d) OnEvent(value any) grid2d {
	return g.set("onEvent", value)
}

// RowHeight 单位行高度，默认 50 px
func (g grid2d) RowHeight(value string) grid2d {
	return g.set("rowHeight", value)
}

// Static 是否静态展示
func (g grid2d) Static(value bool) grid2d {
	return g.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (g grid2d) StaticClassName(value string) grid2d {
	return g.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (g grid2d) StaticInputClassName(value string) grid2d {
	return g.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (g grid2d) StaticLabelClassName(value string) grid2d {
	return g.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (g grid2d) StaticOn(value string) grid2d {
	return g.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (g grid2d) StaticPlaceholder(value string) grid2d {
	return g.set("staticPlaceholder", value)
}

// StaticSchema 静态展示模式的 schema
func (g grid2d) StaticSchema(value string) grid2d {
	return g.set("staticSchema", value)
}

// Style 组件样式
func (g grid2d) Style(value any) grid2d {
	return g.set("style", value)
}

// TestIdBuilder 测试 id 构建器
func (g grid2d) TestIdBuilder(value string) grid2d {
	return g.set("testIdBuilder", value)
}

// Testid 测试 id
func (g grid2d) Testid(value string) grid2d {
	return g.set("testid", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (g grid2d) UseMobileUI(value bool) grid2d {
	return g.set("useMobileUI", value)
}

// Visible 是否显示
func (g grid2d) Visible(value bool) grid2d {
	return g.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (g grid2d) VisibleOn(value string) grid2d {
	return g.set("visibleOn", value)
}

// Width grid 2d 容器宽度，默认是 auto
func (g grid2d) Width(value string) grid2d {
	return g.set("width", value)
}
