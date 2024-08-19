package comp

// paginationWrapper 代表 amis paginationWrapper 渲染器
//
// @version 6.7.0
type paginationWrapper schema

// PaginationWrapper 创建一个新的 PaginationWrapper 实例
func PaginationWrapper() paginationWrapper {
	return paginationWrapper{}.set("type", "pagination-wrapper")
}

// set 用于设置字段值
func (pw paginationWrapper) set(key string, value any) paginationWrapper {
	pw[key] = value
	return pw
}

// Body 内容区域
func (pw paginationWrapper) Body(value ...any) paginationWrapper {
	return pw.set("body", value)
}

// ClassName 容器 css 类名
func (pw paginationWrapper) ClassName(value string) paginationWrapper {
	return pw.set("className", value)
}

// Disabled 是否禁用
func (pw paginationWrapper) Disabled(value bool) paginationWrapper {
	return pw.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (pw paginationWrapper) DisabledOn(value string) paginationWrapper {
	return pw.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (pw paginationWrapper) EditorSetting(value string) paginationWrapper {
	return pw.set("editorSetting", value)
}

// Hidden 是否隐藏
func (pw paginationWrapper) Hidden(value bool) paginationWrapper {
	return pw.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (pw paginationWrapper) HiddenOn(value string) paginationWrapper {
	return pw.set("hiddenOn", value)
}

// ID 组件唯一 id
func (pw paginationWrapper) ID(value string) paginationWrapper {
	return pw.set("id", value)
}

// InputName 输入字段名
func (pw paginationWrapper) InputName(value string) paginationWrapper {
	return pw.set("inputName", value)
}

// MaxButtons 最多显示多少个分页按钮
func (pw paginationWrapper) MaxButtons(value string) paginationWrapper {
	return pw.set("maxButtons", value)
}

// OnEvent 事件动作配置
func (pw paginationWrapper) OnEvent(value any) paginationWrapper {
	return pw.set("onEvent", value)
}

// OutputName 输出字段名
func (pw paginationWrapper) OutputName(value string) paginationWrapper {
	return pw.set("outputName", value)
}

// PerPage 每页显示多条数据
func (pw paginationWrapper) PerPage(value string) paginationWrapper {
	return pw.set("perPage", value)
}

// Position 分页显示位置
func (pw paginationWrapper) Position(value string) paginationWrapper {
	return pw.set("position", value)
}

// ShowPageInput 是否显示快速跳转输入框
func (pw paginationWrapper) ShowPageInput(value bool) paginationWrapper {
	return pw.set("showPageInput", value)
}

// Static 是否静态展示
func (pw paginationWrapper) Static(value bool) paginationWrapper {
	return pw.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (pw paginationWrapper) StaticClassName(value string) paginationWrapper {
	return pw.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项 Value 类名
func (pw paginationWrapper) StaticInputClassName(value string) paginationWrapper {
	return pw.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项 Label 类名
func (pw paginationWrapper) StaticLabelClassName(value string) paginationWrapper {
	return pw.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (pw paginationWrapper) StaticOn(value string) paginationWrapper {
	return pw.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (pw paginationWrapper) StaticPlaceholder(value string) paginationWrapper {
	return pw.set("staticPlaceholder", value)
}

// StaticSchema 静态展示模式的 schema
func (pw paginationWrapper) StaticSchema(value string) paginationWrapper {
	return pw.set("staticSchema", value)
}

// Style 组件样式
func (pw paginationWrapper) Style(value any) paginationWrapper {
	return pw.set("style", value)
}

// TestIdBuilder 自定义测试 ID 构建器
func (pw paginationWrapper) TestIdBuilder(value string) paginationWrapper {
	return pw.set("testIdBuilder", value)
}

// Testid 测试 ID
func (pw paginationWrapper) Testid(value string) paginationWrapper {
	return pw.set("testid", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (pw paginationWrapper) UseMobileUI(value bool) paginationWrapper {
	return pw.set("useMobileUI", value)
}

// Visible 是否显示
func (pw paginationWrapper) Visible(value bool) paginationWrapper {
	return pw.set("visible", value)
}

// VisibleOn 是否显示表达式
func (pw paginationWrapper) VisibleOn(value string) paginationWrapper {
	return pw.set("visibleOn", value)
}
