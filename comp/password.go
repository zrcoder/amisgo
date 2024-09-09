package comp

// password 代表 amis password 渲染器

type password schema

// Password 创建一个新的 Password 实例
func Password() password {
	return password{}.set("type", "password")
}

// set 用于设置字段值
func (p password) set(key string, value any) password {
	p[key] = value
	return p
}

// ClassName 容器 css 类名
func (p password) ClassName(value string) password {
	return p.set("className", value)
}

// Disabled 是否禁用
func (p password) Disabled(value bool) password {
	return p.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (p password) DisabledOn(value string) password {
	return p.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (p password) EditorSetting(value string) password {
	return p.set("editorSetting", value)
}

// Hidden 是否隐藏
func (p password) Hidden(value bool) password {
	return p.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (p password) HiddenOn(value string) password {
	return p.set("hiddenOn", value)
}

// ID 组件唯一 id
func (p password) ID(value string) password {
	return p.set("id", value)
}

// MosaicText 打码模式的文本
func (p password) MosaicText(value string) password {
	return p.set("mosaicText", value)
}

// OnEvent 事件动作配置
func (p password) OnEvent(value any) password {
	return p.set("onEvent", value)
}

// Static 是否静态展示
func (p password) Static(value bool) password {
	return p.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (p password) StaticClassName(value string) password {
	return p.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项 Value 类名
func (p password) StaticInputClassName(value string) password {
	return p.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项 Label 类名
func (p password) StaticLabelClassName(value string) password {
	return p.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (p password) StaticOn(value string) password {
	return p.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (p password) StaticPlaceholder(value string) password {
	return p.set("staticPlaceholder", value)
}

// StaticSchema 静态展示模式的 schema
func (p password) StaticSchema(value string) password {
	return p.set("staticSchema", value)
}

// Style 组件样式
func (p password) Style(value any) password {
	return p.set("style", value)
}

// TestIdBuilder 自定义测试 ID 构建器
func (p password) TestIdBuilder(value string) password {
	return p.set("testIdBuilder", value)
}

// Testid 测试 ID
func (p password) Testid(value string) password {
	return p.set("testid", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (p password) UseMobileUI(value bool) password {
	return p.set("useMobileUI", value)
}

// Visible 是否显示
func (p password) Visible(value bool) password {
	return p.set("visible", value)
}

// VisibleOn 是否显示表达式
func (p password) VisibleOn(value string) password {
	return p.set("visibleOn", value)
}
