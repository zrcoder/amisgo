package comp

// iframe 渲染器 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/iframe
type iframe Schema

// Iframe 创建一个新的 Iframe 实例
func Iframe() iframe {
	return make(iframe).set("type", "iframe")
}

// Allow
func (i iframe) Allow(value string) iframe {
	return i.set("allow", value)
}

// ClassName 容器 css 类名
func (i iframe) ClassName(value string) iframe {
	return i.set("className", value)
}

// Disabled 是否禁用
func (i iframe) Disabled(value bool) iframe {
	return i.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (i iframe) DisabledOn(value string) iframe {
	return i.set("disabledOn", value)
}

// EditorSetting 编辑器配置
func (i iframe) EditorSetting(value string) iframe {
	return i.set("editorSetting", value)
}

// Events 事件响应
func (i iframe) Events(value string) iframe {
	return i.set("events", value)
}

// Height
func (i iframe) Height(value string) iframe {
	return i.set("height", value)
}

// Hidden 是否隐藏
func (i iframe) Hidden(value bool) iframe {
	return i.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (i iframe) HiddenOn(value string) iframe {
	return i.set("hiddenOn", value)
}

// Id 组件唯一 id
func (i iframe) ID(value string) iframe {
	return i.set("id", value)
}

// Name
func (i iframe) Name(value string) iframe {
	return i.set("name", value)
}

// OnEvent 事件动作配置
func (i iframe) OnEvent(value any) iframe {
	return i.set("onEvent", value)
}

// Referrerpolicy 可选值: no-referrer | no-referrer-when-downgrade | origin | origin-when-cross-origin | same-origin | strict-origin | strict-origin-when-cross-origin | unsafe-url
func (i iframe) Referrerpolicy(value string) iframe {
	return i.set("referrerpolicy", value)
}

// Sandbox
func (i iframe) Sandbox(value string) iframe {
	return i.set("sandbox", value)
}

// Src 页面地址
func (i iframe) Src(value string) iframe {
	return i.set("src", value)
}

// Static 是否静态展示
func (i iframe) Static(value bool) iframe {
	return i.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (i iframe) StaticClassName(value string) iframe {
	return i.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (i iframe) StaticInputClassName(value string) iframe {
	return i.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (i iframe) StaticLabelClassName(value string) iframe {
	return i.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (i iframe) StaticOn(value string) iframe {
	return i.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (i iframe) StaticPlaceholder(value string) iframe {
	return i.set("staticPlaceholder", value)
}

// StaticSchema
func (i iframe) StaticSchema(value string) iframe {
	return i.set("staticSchema", value)
}

// Style 组件样式
func (i iframe) Style(value any) iframe {
	return i.set("style", value)
}

// TestIdBuilder
func (i iframe) TestIdBuilder(value string) iframe {
	return i.set("testIdBuilder", value)
}

// Testid
func (i iframe) Testid(value string) iframe {
	return i.set("testid", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (i iframe) UseMobileUI(value bool) iframe {
	return i.set("useMobileUI", value)
}

// Visible 是否显示
func (i iframe) Visible(value bool) iframe {
	return i.set("visible", value)
}

// VisibleOn 是否显示表达式
func (i iframe) VisibleOn(value string) iframe {
	return i.set("visibleOn", value)
}

// Width
func (i iframe) Width(value string) iframe {
	return i.set("width", value)
}

func (i iframe) set(key string, value any) iframe {
	i[key] = value
	return i
}
