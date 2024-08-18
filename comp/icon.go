package comp

// Icon 图标渲染器
type Icon Schema

// NewIcon 创建一个新的 Icon 实例，并设置默认的 type
func NewIcon() Icon {
	return make(Icon).set("type", "icon")
}

func (i Icon) set(key string, value interface{}) Icon {
	i[key] = value
	return i
}

// Badge 角标 (Badge 角标)
func (i Icon) Badge(value string) Icon {
	return i.set("badge", value)
}

// ClassName 容器 css 类名
func (i Icon) ClassName(value string) Icon {
	return i.set("className", value)
}

// Disabled 是否禁用
func (i Icon) Disabled(value bool) Icon {
	return i.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (i Icon) DisabledOn(value string) Icon {
	return i.set("disabledOn", value)
}

// EditorSetting 编辑器配置
func (i Icon) EditorSetting(value string) Icon {
	return i.set("editorSetting", value)
}

// Hidden 是否隐藏
func (i Icon) Hidden(value bool) Icon {
	return i.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (i Icon) HiddenOn(value string) Icon {
	return i.set("hiddenOn", value)
}

// Icon 图标类型
func (i Icon) Icon(value string) Icon {
	return i.set("icon", value)
}

// ID 组件唯一 id
func (i Icon) ID(value string) Icon {
	return i.set("id", value)
}

// OnEvent 事件动作配置
func (i Icon) OnEvent(value string) Icon {
	return i.set("onEvent", value)
}

// Static 是否静态展示
func (i Icon) Static(value bool) Icon {
	return i.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (i Icon) StaticClassName(value string) Icon {
	return i.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (i Icon) StaticInputClassName(value string) Icon {
	return i.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (i Icon) StaticLabelClassName(value string) Icon {
	return i.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (i Icon) StaticOn(value string) Icon {
	return i.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (i Icon) StaticPlaceholder(value string) Icon {
	return i.set("staticPlaceholder", value)
}

// StaticSchema 静态展示 schema
func (i Icon) StaticSchema(value string) Icon {
	return i.set("staticSchema", value)
}

// Style 组件样式
func (i Icon) Style(value string) Icon {
	return i.set("style", value)
}

// TestIdBuilder 测试 id 构建器
func (i Icon) TestIdBuilder(value string) Icon {
	return i.set("testIdBuilder", value)
}

// Testid 测试 id
func (i Icon) Testid(value string) Icon {
	return i.set("testid", value)
}

// UseMobileUI 组件级别用来关闭移动端样式
func (i Icon) UseMobileUI(value bool) Icon {
	return i.set("useMobileUI", value)
}

// Vendor 可选值: iconfont | fa |
func (i Icon) Vendor(value string) Icon {
	return i.set("vendor", value)
}

// Visible 是否显示
func (i Icon) Visible(value bool) Icon {
	return i.set("visible", value)
}

// VisibleOn 是否显示表达式
func (i Icon) VisibleOn(value string) Icon {
	return i.set("visibleOn", value)
}
