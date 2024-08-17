package comp

// Each 表示循环功能渲染器。
type Each BaseRenderer

// NewEach 创建一个新的 Each 实例，并设置默认的 type 为 'each'
func NewEach() Each {
	e := Each(make(BaseRenderer))
	return e.set("type", "each")
}

func (e Each) set(key string, value interface{}) Each {
	e[key] = value
	return e
}

// ClassName 容器 CSS 类名
func (e Each) ClassName(value string) Each {
	return e.set("className", value)
}

// Disabled 是否禁用
func (e Each) Disabled(value bool) Each {
	return e.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (e Each) DisabledOn(value string) Each {
	return e.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (e Each) EditorSetting(value string) Each {
	return e.set("editorSetting", value)
}

// Hidden 是否隐藏
func (e Each) Hidden(value bool) Each {
	return e.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (e Each) HiddenOn(value string) Each {
	return e.set("hiddenOn", value)
}

// ID 组件唯一 ID，主要用于日志采集
func (e Each) ID(value string) Each {
	return e.set("id", value)
}

// IndexKeyName 控制通过什么字段读取序号
func (e Each) IndexKeyName(value string) Each {
	return e.set("indexKeyName", value)
}

// ItemKeyName 控制通过什么字段读取成员数据
func (e Each) ItemKeyName(value string) Each {
	return e.set("itemKeyName", value)
}

// Items 数据项
func (e Each) Items(value string) Each {
	return e.set("items", value)
}

// Name 关联字段名
func (e Each) Name(value string) Each {
	return e.set("name", value)
}

// OnEvent 事件动作配置
func (e Each) OnEvent(value string) Each {
	return e.set("onEvent", value)
}

// Placeholder 占位符
func (e Each) Placeholder(value string) Each {
	return e.set("placeholder", value)
}

// Source 关联字段名，支持数据映射
func (e Each) Source(value string) Each {
	return e.set("source", value)
}

// Static 是否静态展示
func (e Each) Static(value bool) Each {
	return e.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (e Each) StaticClassName(value string) Each {
	return e.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项值类名
func (e Each) StaticInputClassName(value string) Each {
	return e.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项标签类名
func (e Each) StaticLabelClassName(value string) Each {
	return e.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (e Each) StaticOn(value string) Each {
	return e.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (e Each) StaticPlaceholder(value string) Each {
	return e.set("staticPlaceholder", value)
}

// StaticSchema 静态展示 schema
func (e Each) StaticSchema(value string) Each {
	return e.set("staticSchema", value)
}

// Style 组件样式
func (e Each) Style(value string) Each {
	return e.set("style", value)
}

// TestIdBuilder 测试 ID 构造函数
func (e Each) TestIdBuilder(value string) Each {
	return e.set("testIdBuilder", value)
}

// Testid 测试 ID
func (e Each) Testid(value string) Each {
	return e.set("testid", value)
}

// UseMobileUI 组件级别用来关闭移动端样式
func (e Each) UseMobileUI(value bool) Each {
	return e.set("useMobileUI", value)
}

// Visible 是否显示
func (e Each) Visible(value bool) Each {
	return e.set("visible", value)
}

// VisibleOn 是否显示表达式
func (e Each) VisibleOn(value string) Each {
	return e.set("visibleOn", value)
}
