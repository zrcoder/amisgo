package comp

// each 表示循环功能渲染器。
type each schema

// Each 创建一个新的 Each 实例，并设置默认的 type 为 'each'
func Each() each {
	return make(each).set("type", "each")
}

func (e each) set(key string, value interface{}) each {
	e[key] = value
	return e
}

// ClassName 容器 CSS 类名
func (e each) ClassName(value string) each {
	return e.set("className", value)
}

// Disabled 是否禁用
func (e each) Disabled(value bool) each {
	return e.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (e each) DisabledOn(value string) each {
	return e.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (e each) EditorSetting(value string) each {
	return e.set("editorSetting", value)
}

// Hidden 是否隐藏
func (e each) Hidden(value bool) each {
	return e.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (e each) HiddenOn(value string) each {
	return e.set("hiddenOn", value)
}

// ID 组件唯一 ID，主要用于日志采集
func (e each) ID(value string) each {
	return e.set("id", value)
}

// IndexKeyName 控制通过什么字段读取序号
func (e each) IndexKeyName(value string) each {
	return e.set("indexKeyName", value)
}

// ItemKeyName 控制通过什么字段读取成员数据
func (e each) ItemKeyName(value string) each {
	return e.set("itemKeyName", value)
}

// Items 数据项
func (e each) Items(value string) each {
	return e.set("items", value)
}

// Name 关联字段名
func (e each) Name(value string) each {
	return e.set("name", value)
}

// OnEvent 事件动作配置
func (e each) OnEvent(value string) each {
	return e.set("onEvent", value)
}

// Placeholder 占位符
func (e each) Placeholder(value string) each {
	return e.set("placeholder", value)
}

// Source 关联字段名，支持数据映射
func (e each) Source(value string) each {
	return e.set("source", value)
}

// Static 是否静态展示
func (e each) Static(value bool) each {
	return e.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (e each) StaticClassName(value string) each {
	return e.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项值类名
func (e each) StaticInputClassName(value string) each {
	return e.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项标签类名
func (e each) StaticLabelClassName(value string) each {
	return e.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (e each) StaticOn(value string) each {
	return e.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (e each) StaticPlaceholder(value string) each {
	return e.set("staticPlaceholder", value)
}

// StaticSchema 静态展示 schema
func (e each) StaticSchema(value string) each {
	return e.set("staticSchema", value)
}

// Style 组件样式
func (e each) Style(value string) each {
	return e.set("style", value)
}

// TestIdBuilder 测试 ID 构造函数
func (e each) TestIdBuilder(value string) each {
	return e.set("testIdBuilder", value)
}

// Testid 测试 ID
func (e each) Testid(value string) each {
	return e.set("testid", value)
}

// UseMobileUI 组件级别用来关闭移动端样式
func (e each) UseMobileUI(value bool) each {
	return e.set("useMobileUI", value)
}

// Visible 是否显示
func (e each) Visible(value bool) each {
	return e.set("visible", value)
}

// VisibleOn 是否显示表达式
func (e each) VisibleOn(value string) each {
	return e.set("visibleOn", value)
}
