package comp

// operation 操作栏渲染器
//
// @version 6.7.0
type operation schema

// Operation 创建一个新的 Operation 实例
func Operation() operation {
	return operation{}.set("type", "operation")
}

// set 设置字段值
func (o operation) set(key string, value any) operation {
	o[key] = value
	return o
}

// Buttons 设置按钮
func (o operation) Buttons(value string) operation {
	return o.set("buttons", value)
}

// ClassName 设置容器 css 类名
func (o operation) ClassName(value string) operation {
	return o.set("className", value)
}

// Disabled 是否禁用
func (o operation) Disabled(value bool) operation {
	return o.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (o operation) DisabledOn(value string) operation {
	return o.set("disabledOn", value)
}

// EditorSetting 编辑器配置
func (o operation) EditorSetting(value string) operation {
	return o.set("editorSetting", value)
}

// Fixed 固定列
func (o operation) Fixed(value string) operation {
	return o.set("fixed", value)
}

// Hidden 是否隐藏
func (o operation) Hidden(value bool) operation {
	return o.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (o operation) HiddenOn(value string) operation {
	return o.set("hiddenOn", value)
}

// Id 组件唯一 id
func (o operation) Id(value string) operation {
	return o.set("id", value)
}

// Label 设置 label
func (o operation) Label(value string) operation {
	return o.set("label", value)
}

// OnEvent 事件动作配置
func (o operation) OnEvent(value any) operation {
	return o.set("onEvent", value)
}

// Placeholder 占位符
func (o operation) Placeholder(value string) operation {
	return o.set("placeholder", value)
}

// Static 是否静态展示
func (o operation) Static(value bool) operation {
	return o.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (o operation) StaticClassName(value string) operation {
	return o.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项 Value 类名
func (o operation) StaticInputClassName(value string) operation {
	return o.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项 Label 类名
func (o operation) StaticLabelClassName(value string) operation {
	return o.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (o operation) StaticOn(value string) operation {
	return o.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (o operation) StaticPlaceholder(value string) operation {
	return o.set("staticPlaceholder", value)
}

// StaticSchema 静态模式的 schema
func (o operation) StaticSchema(value string) operation {
	return o.set("staticSchema", value)
}

// Style 组件样式
func (o operation) Style(value any) operation {
	return o.set("style", value)
}

// TestIdBuilder 测试 ID 构建器
func (o operation) TestIdBuilder(value string) operation {
	return o.set("testIdBuilder", value)
}

// Testid 测试 ID
func (o operation) Testid(value string) operation {
	return o.set("testid", value)
}

// UseMobileUI 关闭移动端样式
func (o operation) UseMobileUI(value bool) operation {
	return o.set("useMobileUI", value)
}

// Visible 是否显示
func (o operation) Visible(value bool) operation {
	return o.set("visible", value)
}

// VisibleOn 是否显示表达式
func (o operation) VisibleOn(value string) operation {
	return o.set("visibleOn", value)
}
