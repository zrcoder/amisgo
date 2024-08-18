package comp

// Json JSON 数据展示控件
//
// 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/json
// @version 6.7.0
type Json Schema

// NewJson 创建一个新的 Json 实例，并设置默认的 type
func NewJson() Json {
	return make(Json).set("type", "json")
}

// ClassName 容器 css 类名
func (j Json) ClassName(value string) Json {
	j.set("className", value)
	return j
}

func (j Json) set(key string, value interface{}) Json {
	j[key] = value
	return j
}

// Disabled 是否禁用
func (j Json) Disabled(value bool) Json {
	j.set("disabled", value)
	return j
}

// DisabledOn 是否禁用表达式
func (j Json) DisabledOn(value string) Json {
	j.set("disabledOn", value)
	return j
}

// DisplayDataTypes 是否显示数据类型
func (j Json) DisplayDataTypes(value bool) Json {
	j.set("displayDataTypes", value)
	return j
}

// EditorSetting 编辑器配置，运行时可以忽略
func (j Json) EditorSetting(value string) Json {
	j.set("editorSetting", value)
	return j
}

// EllipsisThreshold 设置字符串的最大展示长度
func (j Json) EllipsisThreshold(value string) Json {
	j.set("ellipsisThreshold", value)
	return j
}

// EnableClipboard 是否可复制
func (j Json) EnableClipboard(value bool) Json {
	j.set("enableClipboard", value)
	return j
}

// Hidden 是否隐藏
func (j Json) Hidden(value bool) Json {
	j.set("hidden", value)
	return j
}

// HiddenOn 是否隐藏表达式
func (j Json) HiddenOn(value string) Json {
	j.set("hiddenOn", value)
	return j
}

// IconStyle 图标风格
func (j Json) IconStyle(value string) Json {
	j.set("iconStyle", value)
	return j
}

// ID 组件唯一 id
func (j Json) ID(value string) Json {
	j.set("id", value)
	return j
}

// LevelExpand 默认展开的级别
func (j Json) LevelExpand(value string) Json {
	j.set("levelExpand", value)
	return j
}

// Mutable 是否可修改
func (j Json) Mutable(value bool) Json {
	j.set("mutable", value)
	return j
}

// OnEvent 事件动作配置
func (j Json) OnEvent(value string) Json {
	j.set("onEvent", value)
	return j
}

// QuotesOnKeys 是否显示键的引号
func (j Json) QuotesOnKeys(value bool) Json {
	j.set("quotesOnKeys", value)
	return j
}

// SortKeys 是否为键排序
func (j Json) SortKeys(value bool) Json {
	j.set("sortKeys", value)
	return j
}

// Source 支持从数据链取值
func (j Json) Source(value string) Json {
	j.set("source", value)
	return j
}

// Static 是否静态展示
func (j Json) Static(value bool) Json {
	j.set("static", value)
	return j
}

// StaticClassName 静态展示表单项类名
func (j Json) StaticClassName(value string) Json {
	j.set("staticClassName", value)
	return j
}

// StaticInputClassName 静态展示表单项Value类名
func (j Json) StaticInputClassName(value string) Json {
	j.set("staticInputClassName", value)
	return j
}

// StaticLabelClassName 静态展示表单项Label类名
func (j Json) StaticLabelClassName(value string) Json {
	j.set("staticLabelClassName", value)
	return j
}

// StaticOn 是否静态展示表达式
func (j Json) StaticOn(value string) Json {
	j.set("staticOn", value)
	return j
}

// StaticPlaceholder 静态展示空值占位
func (j Json) StaticPlaceholder(value string) Json {
	j.set("staticPlaceholder", value)
	return j
}

// StaticSchema 静态展示 Schema
func (j Json) StaticSchema(value string) Json {
	j.set("staticSchema", value)
	return j
}

// Style 组件样式
func (j Json) Style(value string) Json {
	j.set("style", value)
	return j
}

// TestIdBuilder 测试 ID 构建器
func (j Json) TestIdBuilder(value string) Json {
	j.set("testIdBuilder", value)
	return j
}

// Testid 测试 id
func (j Json) Testid(value string) Json {
	j.set("testid", value)
	return j
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (j Json) UseMobileUI(value bool) Json {
	j.set("useMobileUI", value)
	return j
}

// Value 要展示的 JSON 数据
func (j Json) Value(value string) Json {
	j.set("value", value)
	return j
}

// Visible 是否显示
func (j Json) Visible(value bool) Json {
	j.set("visible", value)
	return j
}

// VisibleOn 是否显示表达式
func (j Json) VisibleOn(value string) Json {
	j.set("visibleOn", value)
	return j
}
