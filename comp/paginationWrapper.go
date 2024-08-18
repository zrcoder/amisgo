package comp

// PaginationWrapper 代表 amis PaginationWrapper 渲染器
//
// @version 6.7.0
type PaginationWrapper Schema

// NewPaginationWrapper 创建一个新的 PaginationWrapper 实例
func NewPaginationWrapper() PaginationWrapper {
	return PaginationWrapper{}.set("type", "pagination-wrapper")
}

// set 用于设置字段值
func (pw PaginationWrapper) set(key string, value interface{}) PaginationWrapper {
	pw[key] = value
	return pw
}

// Body 内容区域
func (pw PaginationWrapper) Body(value ...interface{}) PaginationWrapper {
	return pw.set("body", value)
}

// ClassName 容器 css 类名
func (pw PaginationWrapper) ClassName(value string) PaginationWrapper {
	return pw.set("className", value)
}

// Disabled 是否禁用
func (pw PaginationWrapper) Disabled(value bool) PaginationWrapper {
	return pw.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (pw PaginationWrapper) DisabledOn(value string) PaginationWrapper {
	return pw.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (pw PaginationWrapper) EditorSetting(value string) PaginationWrapper {
	return pw.set("editorSetting", value)
}

// Hidden 是否隐藏
func (pw PaginationWrapper) Hidden(value bool) PaginationWrapper {
	return pw.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (pw PaginationWrapper) HiddenOn(value string) PaginationWrapper {
	return pw.set("hiddenOn", value)
}

// ID 组件唯一 id
func (pw PaginationWrapper) ID(value string) PaginationWrapper {
	return pw.set("id", value)
}

// InputName 输入字段名
func (pw PaginationWrapper) InputName(value string) PaginationWrapper {
	return pw.set("inputName", value)
}

// MaxButtons 最多显示多少个分页按钮
func (pw PaginationWrapper) MaxButtons(value string) PaginationWrapper {
	return pw.set("maxButtons", value)
}

// OnEvent 事件动作配置
func (pw PaginationWrapper) OnEvent(value string) PaginationWrapper {
	return pw.set("onEvent", value)
}

// OutputName 输出字段名
func (pw PaginationWrapper) OutputName(value string) PaginationWrapper {
	return pw.set("outputName", value)
}

// PerPage 每页显示多条数据
func (pw PaginationWrapper) PerPage(value string) PaginationWrapper {
	return pw.set("perPage", value)
}

// Position 分页显示位置
func (pw PaginationWrapper) Position(value string) PaginationWrapper {
	return pw.set("position", value)
}

// ShowPageInput 是否显示快速跳转输入框
func (pw PaginationWrapper) ShowPageInput(value bool) PaginationWrapper {
	return pw.set("showPageInput", value)
}

// Static 是否静态展示
func (pw PaginationWrapper) Static(value bool) PaginationWrapper {
	return pw.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (pw PaginationWrapper) StaticClassName(value string) PaginationWrapper {
	return pw.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项 Value 类名
func (pw PaginationWrapper) StaticInputClassName(value string) PaginationWrapper {
	return pw.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项 Label 类名
func (pw PaginationWrapper) StaticLabelClassName(value string) PaginationWrapper {
	return pw.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (pw PaginationWrapper) StaticOn(value string) PaginationWrapper {
	return pw.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (pw PaginationWrapper) StaticPlaceholder(value string) PaginationWrapper {
	return pw.set("staticPlaceholder", value)
}

// StaticSchema 静态展示模式的 schema
func (pw PaginationWrapper) StaticSchema(value string) PaginationWrapper {
	return pw.set("staticSchema", value)
}

// Style 组件样式
func (pw PaginationWrapper) Style(value string) PaginationWrapper {
	return pw.set("style", value)
}

// TestIdBuilder 自定义测试 ID 构建器
func (pw PaginationWrapper) TestIdBuilder(value string) PaginationWrapper {
	return pw.set("testIdBuilder", value)
}

// Testid 测试 ID
func (pw PaginationWrapper) Testid(value string) PaginationWrapper {
	return pw.set("testid", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (pw PaginationWrapper) UseMobileUI(value bool) PaginationWrapper {
	return pw.set("useMobileUI", value)
}

// Visible 是否显示
func (pw PaginationWrapper) Visible(value bool) PaginationWrapper {
	return pw.set("visible", value)
}

// VisibleOn 是否显示表达式
func (pw PaginationWrapper) VisibleOn(value string) PaginationWrapper {
	return pw.set("visibleOn", value)
}
