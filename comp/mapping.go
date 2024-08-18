package comp

// Mapping 映射展示控件
//
// @version 6.7.0
type Mapping Schema

// NewMapping 创建一个新的 Mapping 实例
func NewMapping() Mapping {
	return Mapping{}.set("type", "map")
}

func (m Mapping) set(key string, value interface{}) Mapping {
	m[key] = value
	return m
}

// ClassName 容器 css 类名
func (m Mapping) ClassName(value string) Mapping {
	m.set("className", value)
	return m
}

// Disabled 是否禁用
func (m Mapping) Disabled(value bool) Mapping {
	m.set("disabled", value)
	return m
}

// DisabledOn 是否禁用表达式
func (m Mapping) DisabledOn(value string) Mapping {
	m.set("disabledOn", value)
	return m
}

// EditorSetting 编辑器配置
func (m Mapping) EditorSetting(value string) Mapping {
	m.set("editorSetting", value)
	return m
}

// Hidden 是否隐藏
func (m Mapping) Hidden(value bool) Mapping {
	m.set("hidden", value)
	return m
}

// HiddenOn 是否隐藏表达式
func (m Mapping) HiddenOn(value string) Mapping {
	m.set("hiddenOn", value)
	return m
}

// ID 组件唯一 id
func (m Mapping) ID(value string) Mapping {
	m.set("id", value)
	return m
}

// ItemSchema 自定义渲染映射值，支持 html 或 schema
func (m Mapping) ItemSchema(value string) Mapping {
	m.set("itemSchema", value)
	return m
}

// LabelField map 或 source 为对象数组时，作为 label 值的字段名
func (m Mapping) LabelField(value string) Mapping {
	m.set("labelField", value)
	return m
}

// Map 配置映射规则
func (m Mapping) Map(value string) Mapping {
	m.set("map", value)
	return m
}

// Name 关联字段名
func (m Mapping) Name(value string) Mapping {
	m.set("name", value)
	return m
}

// OnEvent 事件动作配置
func (m Mapping) OnEvent(value string) Mapping {
	m.set("onEvent", value)
	return m
}

// Placeholder 占位符
func (m Mapping) Placeholder(value string) Mapping {
	m.set("placeholder", value)
	return m
}

// Source 如果想远程拉取字典，请配置 source 为接口
func (m Mapping) Source(value string) Mapping {
	m.set("source", value)
	return m
}

// Static 是否静态展示
func (m Mapping) Static(value bool) Mapping {
	m.set("static", value)
	return m
}

// StaticClassName 静态展示表单项类名
func (m Mapping) StaticClassName(value string) Mapping {
	m.set("staticClassName", value)
	return m
}

// StaticInputClassName 静态展示表单项 Value 类名
func (m Mapping) StaticInputClassName(value string) Mapping {
	m.set("staticInputClassName", value)
	return m
}

// StaticLabelClassName 静态展示表单项 Label 类名
func (m Mapping) StaticLabelClassName(value string) Mapping {
	m.set("staticLabelClassName", value)
	return m
}

// StaticOn 是否静态展示表达式
func (m Mapping) StaticOn(value string) Mapping {
	m.set("staticOn", value)
	return m
}

// StaticPlaceholder 静态展示空值占位
func (m Mapping) StaticPlaceholder(value string) Mapping {
	m.set("staticPlaceholder", value)
	return m
}

// StaticSchema 静态展示 schema
func (m Mapping) StaticSchema(value string) Mapping {
	m.set("staticSchema", value)
	return m
}

// Style 组件样式
func (m Mapping) Style(value string) Mapping {
	m.set("style", value)
	return m
}

// TestIdBuilder testIdBuilder
func (m Mapping) TestIdBuilder(value string) Mapping {
	m.set("testIdBuilder", value)
	return m
}

// Testid testid
func (m Mapping) Testid(value string) Mapping {
	m.set("testid", value)
	return m
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (m Mapping) UseMobileUI(value bool) Mapping {
	m.set("useMobileUI", value)
	return m
}

// ValueField map 或 source 为对象数组时，作为 value 值的字段名
func (m Mapping) ValueField(value string) Mapping {
	m.set("valueField", value)
	return m
}

// Visible 是否显示
func (m Mapping) Visible(value bool) Mapping {
	m.set("visible", value)
	return m
}

// VisibleOn 是否显示表达式
func (m Mapping) VisibleOn(value string) Mapping {
	m.set("visibleOn", value)
	return m
}
