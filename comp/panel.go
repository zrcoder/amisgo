package comp

// panel 代表 amis panel 渲染器
//
// @version 6.7.0
type panel schema

// Panel 创建一个新的 Panel 实例
func Panel() panel {
	return panel{}.set("type", "panel")
}

// set 用于设置字段值
func (p panel) set(key string, value interface{}) panel {
	p[key] = value
	return p
}

// Actions 按钮集合
func (p panel) Actions(value string) panel {
	return p.set("actions", value)
}

// ActionsClassName 按钮集合外层类名
func (p panel) ActionsClassName(value string) panel {
	return p.set("actionsClassName", value)
}

// ActionsControlClassName 按钮控制的类名
func (p panel) ActionsControlClassName(value string) panel {
	return p.set("actionsControlClassName", value)
}

// AffixFooter 固定底部
func (p panel) AffixFooter(value string) panel {
	return p.set("affixFooter", value)
}

// Body 内容区域
func (p panel) Body(value ...interface{}) panel {
	return p.set("body", value)
}

// BodyClassName 配置 Body 容器 className
func (p panel) BodyClassName(value string) panel {
	return p.set("bodyClassName", value)
}

// BodyControlClassName Body 控制的类名
func (p panel) BodyControlClassName(value string) panel {
	return p.set("bodyControlClassName", value)
}

// ClassName 容器 css 类名
func (p panel) ClassName(value string) panel {
	return p.set("className", value)
}

// Disabled 是否禁用
func (p panel) Disabled(value bool) panel {
	return p.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (p panel) DisabledOn(value string) panel {
	return p.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (p panel) EditorSetting(value string) panel {
	return p.set("editorSetting", value)
}

// Footer 底部内容区域
func (p panel) Footer(value string) panel {
	return p.set("footer", value)
}

// FooterClassName 配置 footer 容器 className
func (p panel) FooterClassName(value string) panel {
	return p.set("footerClassName", value)
}

// FooterWrapClassName footer 和 actions 外层 div 类名
func (p panel) FooterWrapClassName(value string) panel {
	return p.set("footerWrapClassName", value)
}

// Header 头部内容, 和 title 二选一
func (p panel) Header(value string) panel {
	return p.set("header", value)
}

// HeaderClassName 配置 header 容器 className
func (p panel) HeaderClassName(value string) panel {
	return p.set("headerClassName", value)
}

// HeaderControlClassName 头部控制的类名
func (p panel) HeaderControlClassName(value string) panel {
	return p.set("headerControlClassName", value)
}

// Hidden 是否隐藏
func (p panel) Hidden(value bool) panel {
	return p.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (p panel) HiddenOn(value string) panel {
	return p.set("hiddenOn", value)
}

// ID 组件唯一 id
func (p panel) ID(value string) panel {
	return p.set("id", value)
}

// OnEvent 事件动作配置
func (p panel) OnEvent(value string) panel {
	return p.set("onEvent", value)
}

// Static 是否静态展示
func (p panel) Static(value bool) panel {
	return p.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (p panel) StaticClassName(value string) panel {
	return p.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项 Value 类名
func (p panel) StaticInputClassName(value string) panel {
	return p.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项 Label 类名
func (p panel) StaticLabelClassName(value string) panel {
	return p.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (p panel) StaticOn(value string) panel {
	return p.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (p panel) StaticPlaceholder(value string) panel {
	return p.set("staticPlaceholder", value)
}

// StaticSchema 静态展示模式的 schema
func (p panel) StaticSchema(value string) panel {
	return p.set("staticSchema", value)
}

// Style 组件样式
func (p panel) Style(value string) panel {
	return p.set("style", value)
}

// SubFormHorizontal 如果是水平排版，这个属性可以细化水平排版的左右宽度占比
func (p panel) SubFormHorizontal(value string) panel {
	return p.set("subFormHorizontal", value)
}

// SubFormMode 配置子表单项默认的展示方式
func (p panel) SubFormMode(value string) panel {
	return p.set("subFormMode", value)
}

// TestIdBuilder 自定义测试 ID 构建器
func (p panel) TestIdBuilder(value string) panel {
	return p.set("testIdBuilder", value)
}

// Testid 测试 ID
func (p panel) Testid(value string) panel {
	return p.set("testid", value)
}

// Title Panel 标题
func (p panel) Title(value string) panel {
	return p.set("title", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (p panel) UseMobileUI(value bool) panel {
	return p.set("useMobileUI", value)
}

// Visible 是否显示
func (p panel) Visible(value bool) panel {
	return p.set("visible", value)
}

// VisibleOn 是否显示表达式
func (p panel) VisibleOn(value string) panel {
	return p.set("visibleOn", value)
}
