package comp

// mapping 映射展示控件
//
// @version 6.7.0
type mapping schema

// Mapping 创建一个新的 Mapping 实例
func Mapping() mapping {
	return mapping{}.set("type", "map")
}

func (m mapping) set(key string, value any) mapping {
	m[key] = value
	return m
}

// ClassName 容器 css 类名
func (m mapping) ClassName(value string) mapping {
	m.set("className", value)
	return m
}

// Disabled 是否禁用
func (m mapping) Disabled(value bool) mapping {
	m.set("disabled", value)
	return m
}

// DisabledOn 是否禁用表达式
func (m mapping) DisabledOn(value string) mapping {
	m.set("disabledOn", value)
	return m
}

// EditorSetting 编辑器配置
func (m mapping) EditorSetting(value string) mapping {
	m.set("editorSetting", value)
	return m
}

// Hidden 是否隐藏
func (m mapping) Hidden(value bool) mapping {
	m.set("hidden", value)
	return m
}

// HiddenOn 是否隐藏表达式
func (m mapping) HiddenOn(value string) mapping {
	m.set("hiddenOn", value)
	return m
}

// ID 组件唯一 id
func (m mapping) ID(value string) mapping {
	m.set("id", value)
	return m
}

// ItemSchema 自定义渲染映射值，支持 html 或 schema
func (m mapping) ItemSchema(value string) mapping {
	m.set("itemSchema", value)
	return m
}

// LabelField map 或 source 为对象数组时，作为 label 值的字段名
func (m mapping) LabelField(value string) mapping {
	m.set("labelField", value)
	return m
}

// Map 配置映射规则
func (m mapping) Map(value string) mapping {
	m.set("map", value)
	return m
}

// Name 关联字段名
func (m mapping) Name(value string) mapping {
	m.set("name", value)
	return m
}

// OnEvent 事件动作配置
func (m mapping) OnEvent(value any) mapping {
	m.set("onEvent", value)
	return m
}

// Placeholder 占位符
func (m mapping) Placeholder(value string) mapping {
	m.set("placeholder", value)
	return m
}

// Source 如果想远程拉取字典，请配置 source 为接口
func (m mapping) Source(value string) mapping {
	m.set("source", value)
	return m
}

// Static 是否静态展示
func (m mapping) Static(value bool) mapping {
	m.set("static", value)
	return m
}

// StaticClassName 静态展示表单项类名
func (m mapping) StaticClassName(value string) mapping {
	m.set("staticClassName", value)
	return m
}

// StaticInputClassName 静态展示表单项 Value 类名
func (m mapping) StaticInputClassName(value string) mapping {
	m.set("staticInputClassName", value)
	return m
}

// StaticLabelClassName 静态展示表单项 Label 类名
func (m mapping) StaticLabelClassName(value string) mapping {
	m.set("staticLabelClassName", value)
	return m
}

// StaticOn 是否静态展示表达式
func (m mapping) StaticOn(value string) mapping {
	m.set("staticOn", value)
	return m
}

// StaticPlaceholder 静态展示空值占位
func (m mapping) StaticPlaceholder(value string) mapping {
	m.set("staticPlaceholder", value)
	return m
}

// StaticSchema 静态展示 schema
func (m mapping) StaticSchema(value string) mapping {
	m.set("staticSchema", value)
	return m
}

// Style 组件样式
func (m mapping) Style(value any) mapping {
	m.set("style", value)
	return m
}

// TestIdBuilder testIdBuilder
func (m mapping) TestIdBuilder(value string) mapping {
	m.set("testIdBuilder", value)
	return m
}

// Testid testid
func (m mapping) Testid(value string) mapping {
	m.set("testid", value)
	return m
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (m mapping) UseMobileUI(value bool) mapping {
	m.set("useMobileUI", value)
	return m
}

// ValueField map 或 source 为对象数组时，作为 value 值的字段名
func (m mapping) ValueField(value string) mapping {
	m.set("valueField", value)
	return m
}

// Visible 是否显示
func (m mapping) Visible(value bool) mapping {
	m.set("visible", value)
	return m
}

// VisibleOn 是否显示表达式
func (m mapping) VisibleOn(value string) mapping {
	m.set("visibleOn", value)
	return m
}
