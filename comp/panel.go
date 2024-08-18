package comp

// Panel 代表 amis Panel 渲染器
//
// @version 6.7.0
type Panel Schema

// NewPanel 创建一个新的 Panel 实例
func NewPanel() Panel {
	return Panel{}.set("type", "panel")
}

// set 用于设置字段值
func (p Panel) set(key string, value interface{}) Panel {
	p[key] = value
	return p
}

// Actions 按钮集合
func (p Panel) Actions(value string) Panel {
	return p.set("actions", value)
}

// ActionsClassName 按钮集合外层类名
func (p Panel) ActionsClassName(value string) Panel {
	return p.set("actionsClassName", value)
}

// ActionsControlClassName 按钮控制的类名
func (p Panel) ActionsControlClassName(value string) Panel {
	return p.set("actionsControlClassName", value)
}

// AffixFooter 固定底部
func (p Panel) AffixFooter(value string) Panel {
	return p.set("affixFooter", value)
}

// Body 内容区域
func (p Panel) Body(value ...interface{}) Panel {
	return p.set("body", value)
}

// BodyClassName 配置 Body 容器 className
func (p Panel) BodyClassName(value string) Panel {
	return p.set("bodyClassName", value)
}

// BodyControlClassName Body 控制的类名
func (p Panel) BodyControlClassName(value string) Panel {
	return p.set("bodyControlClassName", value)
}

// ClassName 容器 css 类名
func (p Panel) ClassName(value string) Panel {
	return p.set("className", value)
}

// Disabled 是否禁用
func (p Panel) Disabled(value bool) Panel {
	return p.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (p Panel) DisabledOn(value string) Panel {
	return p.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (p Panel) EditorSetting(value string) Panel {
	return p.set("editorSetting", value)
}

// Footer 底部内容区域
func (p Panel) Footer(value string) Panel {
	return p.set("footer", value)
}

// FooterClassName 配置 footer 容器 className
func (p Panel) FooterClassName(value string) Panel {
	return p.set("footerClassName", value)
}

// FooterWrapClassName footer 和 actions 外层 div 类名
func (p Panel) FooterWrapClassName(value string) Panel {
	return p.set("footerWrapClassName", value)
}

// Header 头部内容, 和 title 二选一
func (p Panel) Header(value string) Panel {
	return p.set("header", value)
}

// HeaderClassName 配置 header 容器 className
func (p Panel) HeaderClassName(value string) Panel {
	return p.set("headerClassName", value)
}

// HeaderControlClassName 头部控制的类名
func (p Panel) HeaderControlClassName(value string) Panel {
	return p.set("headerControlClassName", value)
}

// Hidden 是否隐藏
func (p Panel) Hidden(value bool) Panel {
	return p.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (p Panel) HiddenOn(value string) Panel {
	return p.set("hiddenOn", value)
}

// ID 组件唯一 id
func (p Panel) ID(value string) Panel {
	return p.set("id", value)
}

// OnEvent 事件动作配置
func (p Panel) OnEvent(value string) Panel {
	return p.set("onEvent", value)
}

// Static 是否静态展示
func (p Panel) Static(value bool) Panel {
	return p.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (p Panel) StaticClassName(value string) Panel {
	return p.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项 Value 类名
func (p Panel) StaticInputClassName(value string) Panel {
	return p.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项 Label 类名
func (p Panel) StaticLabelClassName(value string) Panel {
	return p.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (p Panel) StaticOn(value string) Panel {
	return p.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (p Panel) StaticPlaceholder(value string) Panel {
	return p.set("staticPlaceholder", value)
}

// StaticSchema 静态展示模式的 schema
func (p Panel) StaticSchema(value string) Panel {
	return p.set("staticSchema", value)
}

// Style 组件样式
func (p Panel) Style(value string) Panel {
	return p.set("style", value)
}

// SubFormHorizontal 如果是水平排版，这个属性可以细化水平排版的左右宽度占比
func (p Panel) SubFormHorizontal(value string) Panel {
	return p.set("subFormHorizontal", value)
}

// SubFormMode 配置子表单项默认的展示方式
func (p Panel) SubFormMode(value string) Panel {
	return p.set("subFormMode", value)
}

// TestIdBuilder 自定义测试 ID 构建器
func (p Panel) TestIdBuilder(value string) Panel {
	return p.set("testIdBuilder", value)
}

// Testid 测试 ID
func (p Panel) Testid(value string) Panel {
	return p.set("testid", value)
}

// Title Panel 标题
func (p Panel) Title(value string) Panel {
	return p.set("title", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (p Panel) UseMobileUI(value bool) Panel {
	return p.set("useMobileUI", value)
}

// Visible 是否显示
func (p Panel) Visible(value bool) Panel {
	return p.set("visible", value)
}

// VisibleOn 是否显示表达式
func (p Panel) VisibleOn(value string) Panel {
	return p.set("visibleOn", value)
}
