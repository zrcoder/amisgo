package comp

// pagination 代表 amis pagination 渲染器
//
// @version 6.7.0
type pagination schema

// Pagination 创建一个新的 Pagination 实例
func Pagination() pagination {
	return pagination{}.set("type", "pagination")
}

// set 用于设置字段值
func (p pagination) set(key string, value any) pagination {
	p[key] = value
	return p
}

// ActivePage 当前页数
func (p pagination) ActivePage(value int) pagination {
	return p.set("activePage", value)
}

// ClassName 容器 css 类名
func (p pagination) ClassName(value string) pagination {
	return p.set("className", value)
}

// Disabled 是否禁用， 默认 false
func (p pagination) Disabled(value bool) pagination {
	return p.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (p pagination) DisabledOn(value string) pagination {
	return p.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (p pagination) EditorSetting(value string) pagination {
	return p.set("editorSetting", value)
}

// HasNext 是否有下一页
func (p pagination) HasNext(value bool) pagination {
	return p.set("hasNext", value)
}

// Hidden 是否隐藏
func (p pagination) Hidden(value bool) pagination {
	return p.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (p pagination) HiddenOn(value string) pagination {
	return p.set("hiddenOn", value)
}

// ID 组件唯一 id
func (p pagination) ID(value string) pagination {
	return p.set("id", value)
}

// Layout 默认 ["pager"]， 通过控制layout属性的顺序，调整分页结构
func (p pagination) Layout(value []string) pagination {
	return p.set("layout", value)
}

// MaxButtons 最多显示多少个分页按钮
func (p pagination) MaxButtons(value int) pagination {
	return p.set("maxButtons", value)
}

// Mode 模式，normal | simple, 默认  normal
func (p pagination) Mode(value string) pagination {
	return p.set("mode", value)
}

// OnEvent 事件动作配置
func (p pagination) OnEvent(value any) pagination {
	return p.set("onEvent", value)
}

// PerPage 每页显示条数
func (p pagination) PerPage(value int) pagination {
	return p.set("perPage", value)
}

// PerPageAvailable 指定每页可以显示多少条
func (p pagination) PerPageAvailable(value bool) pagination {
	return p.set("perPageAvailable", value)
}

// PopOverContainerSelector 弹层挂载节点
func (p pagination) PopOverContainerSelector(value string) pagination {
	return p.set("popOverContainerSelector", value)
}

// ShowPageInput 是否显示快速跳转输入框
func (p pagination) ShowPageInput(value bool) pagination {
	return p.set("showPageInput", value)
}

// ShowPerPage 是否展示分页切换
func (p pagination) ShowPerPage(value bool) pagination {
	return p.set("showPerPage", value)
}

// Static 是否静态展示
func (p pagination) Static(value bool) pagination {
	return p.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (p pagination) StaticClassName(value string) pagination {
	return p.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项 Value 类名
func (p pagination) StaticInputClassName(value string) pagination {
	return p.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项 Label 类名
func (p pagination) StaticLabelClassName(value string) pagination {
	return p.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (p pagination) StaticOn(value string) pagination {
	return p.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (p pagination) StaticPlaceholder(value string) pagination {
	return p.set("staticPlaceholder", value)
}

// StaticSchema 静态展示模式的 schema
func (p pagination) StaticSchema(value string) pagination {
	return p.set("staticSchema", value)
}

// Style 组件样式
func (p pagination) Style(value any) pagination {
	return p.set("style", value)
}

// TestIdBuilder 自定义测试 ID 构建器
func (p pagination) TestIdBuilder(value string) pagination {
	return p.set("testIdBuilder", value)
}

// Testid 测试 ID
func (p pagination) Testid(value string) pagination {
	return p.set("testid", value)
}

// Total 总条数
func (p pagination) Total(value int) pagination {
	return p.set("total", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (p pagination) UseMobileUI(value bool) pagination {
	return p.set("useMobileUI", value)
}

// Visible 是否显示
func (p pagination) Visible(value bool) pagination {
	return p.set("visible", value)
}

// VisibleOn 是否显示表达式
func (p pagination) VisibleOn(value string) pagination {
	return p.set("visibleOn", value)
}

// Size 尺寸 "sm" | "md"
func (p pagination) Size(value string) pagination {
	return p.set("size", value)
}

// EllipsisPageGap 多页跳转页数，页数较多出现...时点击省略号时每次前进/后退的页数，默认为5
func (p pagination) EllipsisPageGap(value int) pagination {
	return p.set("ellipsisPageGap", value)
}
