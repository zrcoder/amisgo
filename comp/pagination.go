package comp

// Pagination 代表 amis Pagination 渲染器
//
// @version 6.7.0
type Pagination Schema

// NewPagination 创建一个新的 Pagination 实例
func NewPagination() Pagination {
	return Pagination{}.set("type", "pagination")
}

// set 用于设置字段值
func (p Pagination) set(key string, value interface{}) Pagination {
	p[key] = value
	return p
}

// ActivePage 当前页数
func (p Pagination) ActivePage(value string) Pagination {
	return p.set("activePage", value)
}

// ClassName 容器 css 类名
func (p Pagination) ClassName(value string) Pagination {
	return p.set("className", value)
}

// Disabled 是否禁用
func (p Pagination) Disabled(value bool) Pagination {
	return p.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (p Pagination) DisabledOn(value string) Pagination {
	return p.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (p Pagination) EditorSetting(value string) Pagination {
	return p.set("editorSetting", value)
}

// HasNext 是否有下一页
func (p Pagination) HasNext(value bool) Pagination {
	return p.set("hasNext", value)
}

// Hidden 是否隐藏
func (p Pagination) Hidden(value bool) Pagination {
	return p.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (p Pagination) HiddenOn(value string) Pagination {
	return p.set("hiddenOn", value)
}

// ID 组件唯一 id
func (p Pagination) ID(value string) Pagination {
	return p.set("id", value)
}

// Layout 通过控制layout属性的顺序，调整分页结构
func (p Pagination) Layout(value string) Pagination {
	return p.set("layout", value)
}

// MaxButtons 最多显示多少个分页按钮
func (p Pagination) MaxButtons(value string) Pagination {
	return p.set("maxButtons", value)
}

// Mode 模式，默认normal
func (p Pagination) Mode(value string) Pagination {
	return p.set("mode", value)
}

// OnEvent 事件动作配置
func (p Pagination) OnEvent(value string) Pagination {
	return p.set("onEvent", value)
}

// PerPage 每页显示条数
func (p Pagination) PerPage(value string) Pagination {
	return p.set("perPage", value)
}

// PerPageAvailable 指定每页可以显示多少条
func (p Pagination) PerPageAvailable(value bool) Pagination {
	return p.set("perPageAvailable", value)
}

// PopOverContainerSelector 弹层挂载节点
func (p Pagination) PopOverContainerSelector(value string) Pagination {
	return p.set("popOverContainerSelector", value)
}

// ShowPageInput 是否显示快速跳转输入框
func (p Pagination) ShowPageInput(value bool) Pagination {
	return p.set("showPageInput", value)
}

// ShowPerPage 是否展示分页切换
func (p Pagination) ShowPerPage(value bool) Pagination {
	return p.set("showPerPage", value)
}

// Static 是否静态展示
func (p Pagination) Static(value bool) Pagination {
	return p.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (p Pagination) StaticClassName(value string) Pagination {
	return p.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项 Value 类名
func (p Pagination) StaticInputClassName(value string) Pagination {
	return p.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项 Label 类名
func (p Pagination) StaticLabelClassName(value string) Pagination {
	return p.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (p Pagination) StaticOn(value string) Pagination {
	return p.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (p Pagination) StaticPlaceholder(value string) Pagination {
	return p.set("staticPlaceholder", value)
}

// StaticSchema 静态展示模式的 schema
func (p Pagination) StaticSchema(value string) Pagination {
	return p.set("staticSchema", value)
}

// Style 组件样式
func (p Pagination) Style(value string) Pagination {
	return p.set("style", value)
}

// TestIdBuilder 自定义测试 ID 构建器
func (p Pagination) TestIdBuilder(value string) Pagination {
	return p.set("testIdBuilder", value)
}

// Testid 测试 ID
func (p Pagination) Testid(value string) Pagination {
	return p.set("testid", value)
}

// Total 总条数
func (p Pagination) Total(value string) Pagination {
	return p.set("total", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (p Pagination) UseMobileUI(value bool) Pagination {
	return p.set("useMobileUI", value)
}

// Visible 是否显示
func (p Pagination) Visible(value bool) Pagination {
	return p.set("visible", value)
}

// VisibleOn 是否显示表达式
func (p Pagination) VisibleOn(value string) Pagination {
	return p.set("visibleOn", value)
}
