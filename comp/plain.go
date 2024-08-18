package comp

// Plain 纯文本渲染器
//
// @version 6.7.0
type Plain Schema

// NewPlain 创建一个新的 Plain 实例
func NewPlain() Plain {
	return Plain{}.set("type", "plain")
}

func (p Plain) set(key string, value interface{}) Plain {
	p[key] = value
	return p
}

// ClassName 容器 css 类名
func (p Plain) ClassName(value string) Plain {
	return p.set("className", value)
}

// Disabled 是否禁用
func (p Plain) Disabled(value bool) Plain {
	return p.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (p Plain) DisabledOn(value string) Plain {
	return p.set("disabledOn", value)
}

// EditorSetting 编辑器配置
func (p Plain) EditorSetting(value string) Plain {
	return p.set("editorSetting", value)
}

// Hidden 是否隐藏
func (p Plain) Hidden(value bool) Plain {
	return p.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (p Plain) HiddenOn(value string) Plain {
	return p.set("hiddenOn", value)
}

// ID 组件唯一 id
func (p Plain) ID(value string) Plain {
	return p.set("id", value)
}

// Inline 是否内联显示
func (p Plain) Inline(value bool) Plain {
	return p.set("inline", value)
}

// OnEvent 事件动作配置
func (p Plain) OnEvent(value string) Plain {
	return p.set("onEvent", value)
}

// Placeholder 占位符
func (p Plain) Placeholder(value string) Plain {
	return p.set("placeholder", value)
}

// Static 是否静态展示
func (p Plain) Static(value bool) Plain {
	return p.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (p Plain) StaticClassName(value string) Plain {
	return p.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (p Plain) StaticInputClassName(value string) Plain {
	return p.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (p Plain) StaticLabelClassName(value string) Plain {
	return p.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (p Plain) StaticOn(value string) Plain {
	return p.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (p Plain) StaticPlaceholder(value string) Plain {
	return p.set("staticPlaceholder", value)
}

// StaticSchema 静态展示 schema
func (p Plain) StaticSchema(value string) Plain {
	return p.set("staticSchema", value)
}

// Style 组件样式
func (p Plain) Style(value string) Plain {
	return p.set("style", value)
}

// TestIdBuilder 测试 id 构造器
func (p Plain) TestIdBuilder(value string) Plain {
	return p.set("testIdBuilder", value)
}

// Testid 测试 id
func (p Plain) Testid(value string) Plain {
	return p.set("testid", value)
}

// Text 支持两种语法
func (p Plain) Text(value string) Plain {
	return p.set("text", value)
}

// Tpl 支持两种语法
func (p Plain) Tpl(value string) Plain {
	return p.set("tpl", value)
}

// UseMobileUI 组件级别关闭移动端样式
func (p Plain) UseMobileUI(value bool) Plain {
	return p.set("useMobileUI", value)
}

// Visible 是否显示
func (p Plain) Visible(value bool) Plain {
	return p.set("visible", value)
}

// VisibleOn 是否显示表达式
func (p Plain) VisibleOn(value string) Plain {
	return p.set("visibleOn", value)
}
