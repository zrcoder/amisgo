package comp

// Password 代表 amis Password 渲染器
//
// @version 6.7.0
type Password Schema

// NewPassword 创建一个新的 Password 实例
func NewPassword() Password {
	return Password{}.set("type", "password")
}

// set 用于设置字段值
func (p Password) set(key string, value interface{}) Password {
	p[key] = value
	return p
}

// ClassName 容器 css 类名
func (p Password) ClassName(value string) Password {
	return p.set("className", value)
}

// Disabled 是否禁用
func (p Password) Disabled(value bool) Password {
	return p.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (p Password) DisabledOn(value string) Password {
	return p.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (p Password) EditorSetting(value string) Password {
	return p.set("editorSetting", value)
}

// Hidden 是否隐藏
func (p Password) Hidden(value bool) Password {
	return p.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (p Password) HiddenOn(value string) Password {
	return p.set("hiddenOn", value)
}

// ID 组件唯一 id
func (p Password) ID(value string) Password {
	return p.set("id", value)
}

// MosaicText 打码模式的文本
func (p Password) MosaicText(value string) Password {
	return p.set("mosaicText", value)
}

// OnEvent 事件动作配置
func (p Password) OnEvent(value string) Password {
	return p.set("onEvent", value)
}

// Static 是否静态展示
func (p Password) Static(value bool) Password {
	return p.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (p Password) StaticClassName(value string) Password {
	return p.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项 Value 类名
func (p Password) StaticInputClassName(value string) Password {
	return p.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项 Label 类名
func (p Password) StaticLabelClassName(value string) Password {
	return p.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (p Password) StaticOn(value string) Password {
	return p.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (p Password) StaticPlaceholder(value string) Password {
	return p.set("staticPlaceholder", value)
}

// StaticSchema 静态展示模式的 schema
func (p Password) StaticSchema(value string) Password {
	return p.set("staticSchema", value)
}

// Style 组件样式
func (p Password) Style(value string) Password {
	return p.set("style", value)
}

// TestIdBuilder 自定义测试 ID 构建器
func (p Password) TestIdBuilder(value string) Password {
	return p.set("testIdBuilder", value)
}

// Testid 测试 ID
func (p Password) Testid(value string) Password {
	return p.set("testid", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (p Password) UseMobileUI(value bool) Password {
	return p.set("useMobileUI", value)
}

// Visible 是否显示
func (p Password) Visible(value bool) Password {
	return p.set("visible", value)
}

// VisibleOn 是否显示表达式
func (p Password) VisibleOn(value string) Password {
	return p.set("visibleOn", value)
}
