package comp

// plain 纯文本渲染器
//
// @version 6.7.0
type plain schema

// Plain 创建一个新的 Plain 实例
func Plain() plain {
	return plain{}.set("type", "plain")
}

func (p plain) set(key string, value interface{}) plain {
	p[key] = value
	return p
}

// ClassName 容器 css 类名
func (p plain) ClassName(value string) plain {
	return p.set("className", value)
}

// Disabled 是否禁用
func (p plain) Disabled(value bool) plain {
	return p.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (p plain) DisabledOn(value string) plain {
	return p.set("disabledOn", value)
}

// EditorSetting 编辑器配置
func (p plain) EditorSetting(value string) plain {
	return p.set("editorSetting", value)
}

// Hidden 是否隐藏
func (p plain) Hidden(value bool) plain {
	return p.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (p plain) HiddenOn(value string) plain {
	return p.set("hiddenOn", value)
}

// ID 组件唯一 id
func (p plain) ID(value string) plain {
	return p.set("id", value)
}

// Inline 是否内联显示
func (p plain) Inline(value bool) plain {
	return p.set("inline", value)
}

// OnEvent 事件动作配置
func (p plain) OnEvent(value string) plain {
	return p.set("onEvent", value)
}

// Placeholder 占位符
func (p plain) Placeholder(value string) plain {
	return p.set("placeholder", value)
}

// Static 是否静态展示
func (p plain) Static(value bool) plain {
	return p.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (p plain) StaticClassName(value string) plain {
	return p.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (p plain) StaticInputClassName(value string) plain {
	return p.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (p plain) StaticLabelClassName(value string) plain {
	return p.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (p plain) StaticOn(value string) plain {
	return p.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (p plain) StaticPlaceholder(value string) plain {
	return p.set("staticPlaceholder", value)
}

// StaticSchema 静态展示 schema
func (p plain) StaticSchema(value string) plain {
	return p.set("staticSchema", value)
}

// Style 组件样式
func (p plain) Style(value string) plain {
	return p.set("style", value)
}

// TestIdBuilder 测试 id 构造器
func (p plain) TestIdBuilder(value string) plain {
	return p.set("testIdBuilder", value)
}

// Testid 测试 id
func (p plain) Testid(value string) plain {
	return p.set("testid", value)
}

// Text 支持两种语法
func (p plain) Text(value string) plain {
	return p.set("text", value)
}

// Tpl 支持两种语法
func (p plain) Tpl(value string) plain {
	return p.set("tpl", value)
}

// UseMobileUI 组件级别关闭移动端样式
func (p plain) UseMobileUI(value bool) plain {
	return p.set("useMobileUI", value)
}

// Visible 是否显示
func (p plain) Visible(value bool) plain {
	return p.set("visible", value)
}

// VisibleOn 是否显示表达式
func (p plain) VisibleOn(value string) plain {
	return p.set("visibleOn", value)
}
