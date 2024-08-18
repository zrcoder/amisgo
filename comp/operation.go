package comp

// Operation 操作栏渲染器
//
// @version 6.7.0
type Operation Schema

// NewOperation 创建一个新的 Operation 实例
func NewOperation() Operation {
	return Operation{}.set("type", "operation")
}

// set 设置字段值
func (o Operation) set(key string, value interface{}) Operation {
	o[key] = value
	return o
}

// Buttons 设置按钮
func (o Operation) Buttons(value string) Operation {
	return o.set("buttons", value)
}

// ClassName 设置容器 css 类名
func (o Operation) ClassName(value string) Operation {
	return o.set("className", value)
}

// Disabled 是否禁用
func (o Operation) Disabled(value bool) Operation {
	return o.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (o Operation) DisabledOn(value string) Operation {
	return o.set("disabledOn", value)
}

// EditorSetting 编辑器配置
func (o Operation) EditorSetting(value string) Operation {
	return o.set("editorSetting", value)
}

// Fixed 固定列
func (o Operation) Fixed(value string) Operation {
	return o.set("fixed", value)
}

// Hidden 是否隐藏
func (o Operation) Hidden(value bool) Operation {
	return o.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (o Operation) HiddenOn(value string) Operation {
	return o.set("hiddenOn", value)
}

// Id 组件唯一 id
func (o Operation) Id(value string) Operation {
	return o.set("id", value)
}

// Label 设置 label
func (o Operation) Label(value string) Operation {
	return o.set("label", value)
}

// OnEvent 事件动作配置
func (o Operation) OnEvent(value string) Operation {
	return o.set("onEvent", value)
}

// Placeholder 占位符
func (o Operation) Placeholder(value string) Operation {
	return o.set("placeholder", value)
}

// Static 是否静态展示
func (o Operation) Static(value bool) Operation {
	return o.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (o Operation) StaticClassName(value string) Operation {
	return o.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项 Value 类名
func (o Operation) StaticInputClassName(value string) Operation {
	return o.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项 Label 类名
func (o Operation) StaticLabelClassName(value string) Operation {
	return o.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (o Operation) StaticOn(value string) Operation {
	return o.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (o Operation) StaticPlaceholder(value string) Operation {
	return o.set("staticPlaceholder", value)
}

// StaticSchema 静态模式的 schema
func (o Operation) StaticSchema(value string) Operation {
	return o.set("staticSchema", value)
}

// Style 组件样式
func (o Operation) Style(value string) Operation {
	return o.set("style", value)
}

// TestIdBuilder 测试 ID 构建器
func (o Operation) TestIdBuilder(value string) Operation {
	return o.set("testIdBuilder", value)
}

// Testid 测试 ID
func (o Operation) Testid(value string) Operation {
	return o.set("testid", value)
}

// UseMobileUI 关闭移动端样式
func (o Operation) UseMobileUI(value bool) Operation {
	return o.set("useMobileUI", value)
}

// Visible 是否显示
func (o Operation) Visible(value bool) Operation {
	return o.set("visible", value)
}

// VisibleOn 是否显示表达式
func (o Operation) VisibleOn(value string) Operation {
	return o.set("visibleOn", value)
}
