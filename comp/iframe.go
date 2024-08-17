package comp

// IFrame 渲染器 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/iframe
type IFrame BaseRenderer

// NewIFrame 创建一个新的 IFrame 实例
func NewIFrame() IFrame {
	i := IFrame(make(BaseRenderer))
	i.set("type", "iframe")
	return i
}

// Allow
func (i IFrame) Allow(value string) IFrame {
	return i.set("allow", value)
}

// ClassName 容器 css 类名
func (i IFrame) ClassName(value string) IFrame {
	return i.set("className", value)
}

// Disabled 是否禁用
func (i IFrame) Disabled(value bool) IFrame {
	return i.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (i IFrame) DisabledOn(value string) IFrame {
	return i.set("disabledOn", value)
}

// EditorSetting 编辑器配置
func (i IFrame) EditorSetting(value string) IFrame {
	return i.set("editorSetting", value)
}

// Events 事件响应
func (i IFrame) Events(value string) IFrame {
	return i.set("events", value)
}

// Height
func (i IFrame) Height(value string) IFrame {
	return i.set("height", value)
}

// Hidden 是否隐藏
func (i IFrame) Hidden(value bool) IFrame {
	return i.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (i IFrame) HiddenOn(value string) IFrame {
	return i.set("hiddenOn", value)
}

// Id 组件唯一 id
func (i IFrame) Id(value string) IFrame {
	return i.set("id", value)
}

// Name
func (i IFrame) Name(value string) IFrame {
	return i.set("name", value)
}

// OnEvent 事件动作配置
func (i IFrame) OnEvent(value string) IFrame {
	return i.set("onEvent", value)
}

// Referrerpolicy 可选值: no-referrer | no-referrer-when-downgrade | origin | origin-when-cross-origin | same-origin | strict-origin | strict-origin-when-cross-origin | unsafe-url
func (i IFrame) Referrerpolicy(value string) IFrame {
	return i.set("referrerpolicy", value)
}

// Sandbox
func (i IFrame) Sandbox(value string) IFrame {
	return i.set("sandbox", value)
}

// Src 页面地址
func (i IFrame) Src(value string) IFrame {
	return i.set("src", value)
}

// Static 是否静态展示
func (i IFrame) Static(value bool) IFrame {
	return i.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (i IFrame) StaticClassName(value string) IFrame {
	return i.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (i IFrame) StaticInputClassName(value string) IFrame {
	return i.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (i IFrame) StaticLabelClassName(value string) IFrame {
	return i.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (i IFrame) StaticOn(value string) IFrame {
	return i.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (i IFrame) StaticPlaceholder(value string) IFrame {
	return i.set("staticPlaceholder", value)
}

// StaticSchema
func (i IFrame) StaticSchema(value string) IFrame {
	return i.set("staticSchema", value)
}

// Style 组件样式
func (i IFrame) Style(value string) IFrame {
	return i.set("style", value)
}

// TestIdBuilder
func (i IFrame) TestIdBuilder(value string) IFrame {
	return i.set("testIdBuilder", value)
}

// Testid
func (i IFrame) Testid(value string) IFrame {
	return i.set("testid", value)
}

// Type 指定为 iframe 渲染器
func (i IFrame) Type(value string) IFrame {
	return i.set("type", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (i IFrame) UseMobileUI(value bool) IFrame {
	return i.set("useMobileUI", value)
}

// Visible 是否显示
func (i IFrame) Visible(value bool) IFrame {
	return i.set("visible", value)
}

// VisibleOn 是否显示表达式
func (i IFrame) VisibleOn(value string) IFrame {
	return i.set("visibleOn", value)
}

// Width
func (i IFrame) Width(value string) IFrame {
	return i.set("width", value)
}

// set 设置属性
func (i IFrame) set(key string, value interface{}) IFrame {
	i[key] = value
	return i
}
