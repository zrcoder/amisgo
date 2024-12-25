package comp

// icon 图标渲染器
type icon Schema

// Icon 创建一个新的 Icon 实例，并设置默认的 type
func Icon() icon {
	return make(icon).set("type", "icon")
}

func (i icon) set(key string, value any) icon {
	i[key] = value
	return i
}

// AddOnclassName
func (i icon) AddOnclassName(value string) icon {
	return i.set("addOnclassName", value)
}

// Badge 角标 (Badge 角标)
func (i icon) Badge(value string) icon {
	return i.set("badge", value)
}

// ClassName 容器 css 类名
func (i icon) ClassName(value string) icon {
	return i.set("className", value)
}

// Disabled 是否禁用
func (i icon) Disabled(value bool) icon {
	return i.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (i icon) DisabledOn(value string) icon {
	return i.set("disabledOn", value)
}

// EditorSetting 编辑器配置
func (i icon) EditorSetting(value string) icon {
	return i.set("editorSetting", value)
}

// Hidden 是否隐藏
func (i icon) Hidden(value bool) icon {
	return i.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (i icon) HiddenOn(value string) icon {
	return i.set("hiddenOn", value)
}

// Icon 图标类型
func (i icon) Icon(value string) icon {
	return i.set("icon", value)
}

// ID 组件唯一 id
func (i icon) ID(value string) icon {
	return i.set("id", value)
}

// OnEvent 事件动作配置
func (i icon) OnEvent(value any) icon {
	return i.set("onEvent", value)
}

// Static 是否静态展示
func (i icon) Static(value bool) icon {
	return i.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (i icon) StaticClassName(value string) icon {
	return i.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (i icon) StaticInputClassName(value string) icon {
	return i.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (i icon) StaticLabelClassName(value string) icon {
	return i.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (i icon) StaticOn(value string) icon {
	return i.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (i icon) StaticPlaceholder(value string) icon {
	return i.set("staticPlaceholder", value)
}

// StaticSchema 静态展示 schema
func (i icon) StaticSchema(value string) icon {
	return i.set("staticSchema", value)
}

// Style 组件样式
func (i icon) Style(value any) icon {
	return i.set("style", value)
}

// TestIdBuilder 测试 id 构建器
func (i icon) TestIdBuilder(value string) icon {
	return i.set("testIdBuilder", value)
}

// Testid 测试 id
func (i icon) Testid(value string) icon {
	return i.set("testid", value)
}

// UseMobileUI 组件级别用来关闭移动端样式
func (i icon) UseMobileUI(value bool) icon {
	return i.set("useMobileUI", value)
}

// Vendor 可选值: iconfont | fa |
func (i icon) Vendor(value string) icon {
	return i.set("vendor", value)
}

// Visible 是否显示
func (i icon) Visible(value bool) icon {
	return i.set("visible", value)
}

// VisibleOn 是否显示表达式
func (i icon) VisibleOn(value string) icon {
	return i.set("visibleOn", value)
}
