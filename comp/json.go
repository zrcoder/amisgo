package comp

// json JSON 数据展示控件
// 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/json

type json schema

// Json 创建一个新的 Json 实例，并设置默认的 type
func Json() json {
	return make(json).set("type", "json")
}

// ClassName 容器 css 类名
func (j json) ClassName(value string) json {
	j.set("className", value)
	return j
}

func (j json) set(key string, value any) json {
	j[key] = value
	return j
}

// Disabled 是否禁用
func (j json) Disabled(value bool) json {
	j.set("disabled", value)
	return j
}

// DisabledOn 是否禁用表达式
func (j json) DisabledOn(value string) json {
	j.set("disabledOn", value)
	return j
}

// DisplayDataTypes 是否显示数据类型
func (j json) DisplayDataTypes(value bool) json {
	j.set("displayDataTypes", value)
	return j
}

// EditorSetting 编辑器配置，运行时可以忽略
func (j json) EditorSetting(value string) json {
	j.set("editorSetting", value)
	return j
}

// EllipsisThreshold 设置字符串的最大展示长度
func (j json) EllipsisThreshold(value string) json {
	j.set("ellipsisThreshold", value)
	return j
}

// EnableClipboard 是否可复制
func (j json) EnableClipboard(value bool) json {
	j.set("enableClipboard", value)
	return j
}

// Hidden 是否隐藏
func (j json) Hidden(value bool) json {
	j.set("hidden", value)
	return j
}

// HiddenOn 是否隐藏表达式
func (j json) HiddenOn(value string) json {
	j.set("hiddenOn", value)
	return j
}

// IconStyle 图标风格
func (j json) IconStyle(value any) json {
	j.set("iconStyle", value)
	return j
}

// ID 组件唯一 id
func (j json) ID(value string) json {
	j.set("id", value)
	return j
}

// LevelExpand 默认展开的级别
func (j json) LevelExpand(value string) json {
	j.set("levelExpand", value)
	return j
}

// Mutable 是否可修改
func (j json) Mutable(value bool) json {
	j.set("mutable", value)
	return j
}

// OnEvent 事件动作配置
func (j json) OnEvent(value any) json {
	j.set("onEvent", value)
	return j
}

// QuotesOnKeys 是否显示键的引号
func (j json) QuotesOnKeys(value bool) json {
	j.set("quotesOnKeys", value)
	return j
}

// SortKeys 是否为键排序
func (j json) SortKeys(value bool) json {
	j.set("sortKeys", value)
	return j
}

// Source 支持从数据链取值
func (j json) Source(value string) json {
	j.set("source", value)
	return j
}

// Static 是否静态展示
func (j json) Static(value bool) json {
	j.set("static", value)
	return j
}

// StaticClassName 静态展示表单项类名
func (j json) StaticClassName(value string) json {
	j.set("staticClassName", value)
	return j
}

// StaticInputClassName 静态展示表单项Value类名
func (j json) StaticInputClassName(value string) json {
	j.set("staticInputClassName", value)
	return j
}

// StaticLabelClassName 静态展示表单项Label类名
func (j json) StaticLabelClassName(value string) json {
	j.set("staticLabelClassName", value)
	return j
}

// StaticOn 是否静态展示表达式
func (j json) StaticOn(value string) json {
	j.set("staticOn", value)
	return j
}

// StaticPlaceholder 静态展示空值占位
func (j json) StaticPlaceholder(value string) json {
	j.set("staticPlaceholder", value)
	return j
}

// StaticSchema 静态展示 Schema
func (j json) StaticSchema(value string) json {
	j.set("staticSchema", value)
	return j
}

// Style 组件样式
func (j json) Style(value any) json {
	j.set("style", value)
	return j
}

// TestIdBuilder 测试 ID 构建器
func (j json) TestIdBuilder(value string) json {
	j.set("testIdBuilder", value)
	return j
}

// Testid 测试 id
func (j json) Testid(value string) json {
	j.set("testid", value)
	return j
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (j json) UseMobileUI(value bool) json {
	j.set("useMobileUI", value)
	return j
}

// Value 要展示的 JSON 数据
func (j json) Value(value string) json {
	j.set("value", value)
	return j
}

// Visible 是否显示
func (j json) Visible(value bool) json {
	j.set("visible", value)
	return j
}

// VisibleOn 是否显示表达式
func (j json) VisibleOn(value string) json {
	j.set("visibleOn", value)
	return j
}
