package comp

// Tpl 渲染器
type Tpl Schema

// NewTpl 创建一个新的 Tpl 实例
func NewTpl() Tpl {
	return make(Tpl).set("type", "tpl")
}

// Badge 角标
func (t Tpl) Badge(value string) Tpl {
	return t.set("badge", value)
}

// ClassName 容器 css 类名
func (t Tpl) ClassName(value string) Tpl {
	return t.set("className", value)
}

// Disabled 是否禁用
func (t Tpl) Disabled(value bool) Tpl {
	return t.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (t Tpl) DisabledOn(value string) Tpl {
	return t.set("disabledOn", value)
}

// EditorSetting 编辑器配置
func (t Tpl) EditorSetting(value string) Tpl {
	return t.set("editorSetting", value)
}

// Hidden 是否隐藏
func (t Tpl) Hidden(value bool) Tpl {
	return t.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (t Tpl) HiddenOn(value string) Tpl {
	return t.set("hiddenOn", value)
}

// Html 支持两种语法
func (t Tpl) Html(value string) Tpl {
	return t.set("html", value)
}

// Id 组件唯一 id
func (t Tpl) Id(value string) Tpl {
	return t.set("id", value)
}

// Inline 是否内联显示
func (t Tpl) Inline(value bool) Tpl {
	return t.set("inline", value)
}

// OnEvent 事件动作配置
func (t Tpl) OnEvent(value string) Tpl {
	return t.set("onEvent", value)
}

// Raw
func (t Tpl) Raw(value string) Tpl {
	return t.set("raw", value)
}

// Static 是否静态展示
func (t Tpl) Static(value bool) Tpl {
	return t.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (t Tpl) StaticClassName(value string) Tpl {
	return t.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (t Tpl) StaticInputClassName(value string) Tpl {
	return t.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (t Tpl) StaticLabelClassName(value string) Tpl {
	return t.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (t Tpl) StaticOn(value string) Tpl {
	return t.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (t Tpl) StaticPlaceholder(value string) Tpl {
	return t.set("staticPlaceholder", value)
}

// StaticSchema
func (t Tpl) StaticSchema(value string) Tpl {
	return t.set("staticSchema", value)
}

// Style 自定义样式
func (t Tpl) Style(value string) Tpl {
	return t.set("style", value)
}

// TestIdBuilder
func (t Tpl) TestIdBuilder(value string) Tpl {
	return t.set("testIdBuilder", value)
}

// Testid
func (t Tpl) Testid(value string) Tpl {
	return t.set("testid", value)
}

// Text 支持两种语法
func (t Tpl) Text(value string) Tpl {
	return t.set("text", value)
}

// Tpl 支持两种语法
func (t Tpl) Tpl(value string) Tpl {
	return t.set("tpl", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (t Tpl) UseMobileUI(value bool) Tpl {
	return t.set("useMobileUI", value)
}

// Visible 是否显示
func (t Tpl) Visible(value bool) Tpl {
	return t.set("visible", value)
}

// VisibleOn 是否显示表达式
func (t Tpl) VisibleOn(value string) Tpl {
	return t.set("visibleOn", value)
}

// WrapperComponent 标签类型
func (t Tpl) WrapperComponent(value string) Tpl {
	return t.set("wrapperComponent", value)
}

func (t Tpl) set(key string, value interface{}) Tpl {
	t[key] = value
	return t
}
