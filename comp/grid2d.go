package comp

// Grid2D 二维布局渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/grid-2d
type Grid2D BaseRenderer

// NewGrid2D 创建一个新的 Grid2D 实例
func NewGrid2D() Grid2D {
	return Grid2D(make(BaseRenderer)).set("type", "grid-2d")
}

// set 设置属性并返回当前对象
func (g Grid2D) set(key string, value interface{}) Grid2D {
	g[key] = value
	return g
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (g Grid2D) ClassName(value string) Grid2D {
	return g.set("className", value)
}

// Cols 列数量，默认是 12
func (g Grid2D) Cols(value string) Grid2D {
	return g.set("cols", value)
}

// Disabled 是否禁用
func (g Grid2D) Disabled(value bool) Grid2D {
	return g.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (g Grid2D) DisabledOn(value string) Grid2D {
	return g.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (g Grid2D) EditorSetting(value string) Grid2D {
	return g.set("editorSetting", value)
}

// Gap 格子间距，默认 0，包含行和列
func (g Grid2D) Gap(value string) Grid2D {
	return g.set("gap", value)
}

// GapRow 格子行级别的间距，如果不设置就和 gap 一样
func (g Grid2D) GapRow(value string) Grid2D {
	return g.set("gapRow", value)
}

// Grids 每个格子的配置
func (g Grid2D) Grids(value string) Grid2D {
	return g.set("grids", value)
}

// Hidden 是否隐藏
func (g Grid2D) Hidden(value bool) Grid2D {
	return g.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (g Grid2D) HiddenOn(value string) Grid2D {
	return g.set("hiddenOn", value)
}

// ID 组件唯一 id，主要用于日志采集
func (g Grid2D) ID(value string) Grid2D {
	return g.set("id", value)
}

// OnEvent 事件动作配置
func (g Grid2D) OnEvent(value string) Grid2D {
	return g.set("onEvent", value)
}

// RowHeight 单位行高度，默认 50 px
func (g Grid2D) RowHeight(value string) Grid2D {
	return g.set("rowHeight", value)
}

// Static 是否静态展示
func (g Grid2D) Static(value bool) Grid2D {
	return g.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (g Grid2D) StaticClassName(value string) Grid2D {
	return g.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (g Grid2D) StaticInputClassName(value string) Grid2D {
	return g.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (g Grid2D) StaticLabelClassName(value string) Grid2D {
	return g.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (g Grid2D) StaticOn(value string) Grid2D {
	return g.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (g Grid2D) StaticPlaceholder(value string) Grid2D {
	return g.set("staticPlaceholder", value)
}

// StaticSchema 静态展示模式的 schema
func (g Grid2D) StaticSchema(value string) Grid2D {
	return g.set("staticSchema", value)
}

// Style 组件样式
func (g Grid2D) Style(value string) Grid2D {
	return g.set("style", value)
}

// TestIdBuilder 测试 id 构建器
func (g Grid2D) TestIdBuilder(value string) Grid2D {
	return g.set("testIdBuilder", value)
}

// Testid 测试 id
func (g Grid2D) Testid(value string) Grid2D {
	return g.set("testid", value)
}

// Type 指定为 grid-2d 展示类型
func (g Grid2D) Type(value string) Grid2D {
	return g.set("type", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (g Grid2D) UseMobileUI(value bool) Grid2D {
	return g.set("useMobileUI", value)
}

// Visible 是否显示
func (g Grid2D) Visible(value bool) Grid2D {
	return g.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (g Grid2D) VisibleOn(value string) Grid2D {
	return g.set("visibleOn", value)
}

// Width grid 2d 容器宽度，默认是 auto
func (g Grid2D) Width(value string) Grid2D {
	return g.set("width", value)
}
